// package gomagic is a go wrapper around the libmagic
// magic number recognition library

package gomagic

// #cgo LDFLAGS: -lmagic
// #include <magic.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// NoneFlag signifies no flags
const NoneFlag = C.MAGIC_NONE

// DebugFlag turns on debugging
const DebugFlag = C.MAGIC_DEBUG

// SymlinkFlag follows symlinks
const SymlinkFlag = C.MAGIC_SYMLINK

// CompressFlag checks inside compressed files
const CompressFlag = C.MAGIC_COMPRESS

// DevicesFlag looks at the contents of devices
const DevicesFlag = C.MAGIC_DEVICES

// MimeTypeFlag returns the MIME type
const MimeTypeFlag = C.MAGIC_MIME_TYPE

// ContinueFlag returns all matches
const ContinueFlag = C.MAGIC_CONTINUE

// CheckFlag prints warnings to stderr
const CheckFlag = C.MAGIC_CHECK

// PreserveAtimeFlag restores access time on exit
const PreserveAtimeFlag = C.MAGIC_PRESERVE_ATIME

// RawFlag doesn't convert unprintable chars
const RawFlag = C.MAGIC_RAW

// ErrorFlag handles ENOENT etc as real errors
const ErrorFlag = C.MAGIC_ERROR

// MimeEncodingFlag returns the MIME encoding
const MimeEncodingFlag = C.MAGIC_MIME_ENCODING

// MimeFlag returns mime type and mime encoding
const MimeFlag = (MimeTypeFlag | MimeEncodingFlag)

// AppleFlag returns the Apple creator/type
const AppleFlag = C.MAGIC_APPLE

// ExtensionFlag returns a /-separated list of extensions
const ExtensionFlag = C.MAGIC_EXTENSION

// CompressTransFlag checks inside compressed files but does not report compression
const CompressTransFlag = C.MAGIC_COMPRESS_TRANSP

// NodescFlag returns extension, mime type, and Apple creator type
const NodescFlag = (ExtensionFlag | MimeFlag | AppleFlag)

// NoCheckCompressFlag disabled checking for compressed files
const NoCheckCompressFlag = C.MAGIC_NO_CHECK_COMPRESS

// NoCheckTarFlag disables checking for tar files
const NoCheckTarFlag = C.MAGIC_NO_CHECK_TAR

// NoCheckSoftFlag disables checking of magic entries
const NoCheckSoftFlag = C.MAGIC_NO_CHECK_SOFT

// NoCheckAppTypeFlag disables checking for application type
const NoCheckAppTypeFlag = C.MAGIC_NO_CHECK_APPTYPE

// NoCheckElfFlag disables checking for elf details
const NoCheckElfFlag = C.MAGIC_NO_CHECK_ELF

// NoCheckTextFlag disables checking for text files
const NoCheckTextFlag = C.MAGIC_NO_CHECK_TEXT

// NoCheckCDFFlag disables checking for cdf files
const NoCheckCDFFlag = C.MAGIC_NO_CHECK_CDF

// NoCheckTokensFlag disables checking for tokens
const NoCheckTokensFlag = C.MAGIC_NO_CHECK_TOKENS

// NoCheckEncodingFlag disables checking for text encodings
const NoCheckEncodingFlag = C.MAGIC_NO_CHECK_ENCODING

// ParamIndirMax is a Getparam/Setparam option to control
// how many levels of recursion will be followed for indirect magic entries.
const ParamIndirMax = C.MAGIC_PARAM_INDIR_MAX

// ParamNameMax is a Getparam/Setparam option to control how many levels of
// recursion will be followed for for name/use calls
const ParamNameMax = C.MAGIC_PARAM_NAME_MAX

// ParamElfNotesMax is a Getparam/Setparam option to control how
// many ELF notes will be processed
const ParamElfNotesMax = C.MAGIC_PARAM_ELF_NOTES_MAX

// ParamElfPhnumMax is a Getparam/Setparam option to control how
// many ELF program sections will be processed
const ParamElfPhnumMax = C.MAGIC_PARAM_ELF_PHNUM_MAX

// ParamElfShnumMax is a Getparam/Setparam option to control how
// many ELF sections will be processed
const ParamElfShnumMax = C.MAGIC_PARAM_ELF_SHNUM_MAX

// ParamRegexMax is a Getparam/Setparam option to control how
// many regexes will be processed
const ParamRegexMax = C.MAGIC_PARAM_REGEX_MAX

// ParamBytesMax is a Getparam/Setparam option to control how
// many bytes will be processed
const ParamBytesMax = C.MAGIC_PARAM_BYTES_MAX

// Magic provides the access hook for calling libmagic
type magic struct {
	flags  int
	cookie C.magic_t
}

// New creates a new Magic struct
func New(flags int) (magic, error) {
	magic := magic{}
	magic.flags = C.MAGIC_NONE | flags
	magic.cookie = C.magic_open(C.int(magic.flags))
	if magic.cookie == nil {
		return magic, fmt.Errorf("Failed to initialize magic library")
	}
	// initialize default database
	C.magic_load(magic.cookie, nil)
	return magic, nil
}

// Close closes the libmagic database and deallocates any resources
func (m magic) Close() {
	C.magic_close(m.cookie)
}

// Check checks the validity of entries in the colon separated database files passed in as filename,
// or the empty string for the default database.
func (m magic) Check(fileNames string) error {
	var status C.int
	if fileNames == "" {
		status = C.magic_check(m.cookie, nil)
		fileNames = "default database"
	} else {
		status = C.magic_check(m.cookie, C.CString(fileNames))
	}
	if status != 0 {
		return fmt.Errorf("Database check failed for: %s", fileNames)
	}
	return nil
}

// Compile compiles compile the colon separated list of database files passed
// in as filename, or nil for the default database. Compiled files created are named
// from the basename(1) of each file argument with ``.mgc'' appended to it.
func (m magic) Compile(fileNames string) error {
	var status C.int
	if fileNames == "" {
		status = C.magic_compile(m.cookie, nil)
	} else {
		status = C.magic_compile(m.cookie, C.CString(fileNames))
	}
	if status != 0 {
		return fmt.Errorf("Failed to compile database files: %s", fileNames)
	}
	return nil
}

// Errno returns the last operating system error number (errno(2)) that was
// encountered by a system call
func (m magic) Errno() int {
	return int(C.magic_errno(m.cookie))
}

// Error returns a textual description of the last error or nil if there was no error
func (m magic) Error() string {
	return C.GoString(C.magic_error(m.cookie))
}

// ExamineBuffer returns a string with the libmagic result for
// the provided buffer
func (m magic) ExamineBuffer(buffer []byte) (string, error) {
	result := C.magic_buffer(m.cookie, unsafe.Pointer(&buffer[0]), C.size_t(len(buffer)))
	if result == nil {
		return "", fmt.Errorf("Failed to determine type for file %s", buffer)
	}
	return C.GoString(result), nil
}

// ExamineFile returns a string with the libmagic result for
// the requested file
func (m magic) ExamineFile(fileName string) (string, error) {
	result := C.magic_file(m.cookie, C.CString(fileName))
	if result == nil {
		return "", fmt.Errorf("Failed to determine type for file %s", fileName)
	}
	return C.GoString(result), nil
}

// Getparam returns the current values for various limits related to the magic library
// (see Param* constants)
func (m magic) Getparam(param int) (int, error) {
	value := 0
	status := C.magic_getparam(m.cookie, C.int(param), unsafe.Pointer(&value))
	if status != 0 {
		return value, fmt.Errorf("Failed to set parameters")
	}
	return value, nil
}

// List dumps all magic entries in a human readable format, dumping first the
// entries that are matched against binary files and then the ones that match text files.
// It takes an optional filename argument which is a colon separated list of database files, or
// an empty string for the default database.
func (m magic) List(fileNames string) error {
	var status C.int
	if fileNames == "" {
		status = C.magic_list(m.cookie, nil)
		fileNames = "default database"
	} else {
		status = C.magic_list(m.cookie, C.CString(fileNames))
	}
	if status != 0 {
		return fmt.Errorf("Failed to list the content of database files: %s", fileNames)
	}
	return nil
}

// Load loads the list of colon separated database files. These will be used in the
// next query. If this function is not called the default database will be used.
func (m magic) Load(fileNames string) error {
	var status C.int
	if fileNames == "" {
		status = C.magic_load(m.cookie, nil)
		fileNames = "default database"
	} else {
		status = C.magic_load(m.cookie, C.CString(fileNames))
	}
	if status != 0 {
		return fmt.Errorf("Failed to load new database files: %s", fileNames)
	}
	return nil
}

// SetFlags sets the flags for the magic type determining how magic checking behaves.
// This function returns an error on systems that don't support utime(3), or utimes(2)
// when PreserveAtimeFlag is set
func (m magic) SetFlags(flags int) error {
	result := C.magic_setflags(m.cookie, C.int(flags))
	if result == -1 {
		return fmt.Errorf("Failed to set flags")
	}
	return nil
}

// Setparam sets various limits related to the magic library (see Param* constants)
func (m magic) Setparam(param, value int) error {
	status := C.magic_setparam(m.cookie, C.int(param), unsafe.Pointer(&value))
	if status != 0 {
		return fmt.Errorf("Failed to set parameters")
	}
	return nil
}

// Version returns the version of the magic library used
func (m magic) Version() int {
	return int(C.magic_version())
}
