// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backupstorage

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeDataAlreadyExistsException for service response error code
	// "DataAlreadyExistsException".
	//
	// Non-retryable exception. Attempted to create already existing object or chunk.
	// This message contains a checksum of already presented data.
	ErrCodeDataAlreadyExistsException = "DataAlreadyExistsException"

	// ErrCodeIllegalArgumentException for service response error code
	// "IllegalArgumentException".
	//
	// Non-retryable exception, indicates client error (wrong argument passed to
	// API). See exception message for details.
	ErrCodeIllegalArgumentException = "IllegalArgumentException"

	// ErrCodeKMSInvalidKeyUsageException for service response error code
	// "KMSInvalidKeyUsageException".
	//
	// Non-retryable exception. Indicates the KMS key usage is incorrect. See exception
	// message for details.
	ErrCodeKMSInvalidKeyUsageException = "KMSInvalidKeyUsageException"

	// ErrCodeNotReadableInputStreamException for service response error code
	// "NotReadableInputStreamException".
	//
	// Retryalble exception. Indicated issues while reading an input stream due
	// to the networking issues or connection drop on the client side.
	ErrCodeNotReadableInputStreamException = "NotReadableInputStreamException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Non-retryable exception. Attempted to make an operation on non-existing or
	// expired resource.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeRetryableException for service response error code
	// "RetryableException".
	//
	// Retryable exception. In general indicates internal failure that can be fixed
	// by retry.
	ErrCodeRetryableException = "RetryableException"

	// ErrCodeServiceInternalException for service response error code
	// "ServiceInternalException".
	//
	// Deprecated. To be removed from the model.
	ErrCodeServiceInternalException = "ServiceInternalException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// Retryable exception, indicates internal server error.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Increased rate over throttling limits. Can be retried with exponential backoff.
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":           newErrorAccessDeniedException,
	"DataAlreadyExistsException":      newErrorDataAlreadyExistsException,
	"IllegalArgumentException":        newErrorIllegalArgumentException,
	"KMSInvalidKeyUsageException":     newErrorKMSInvalidKeyUsageException,
	"NotReadableInputStreamException": newErrorNotReadableInputStreamException,
	"ResourceNotFoundException":       newErrorResourceNotFoundException,
	"RetryableException":              newErrorRetryableException,
	"ServiceInternalException":        newErrorServiceInternalException,
	"ServiceUnavailableException":     newErrorServiceUnavailableException,
	"ThrottlingException":             newErrorThrottlingException,
}
