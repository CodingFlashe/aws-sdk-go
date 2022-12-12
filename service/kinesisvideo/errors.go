// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have required permissions to perform this operation.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAccountChannelLimitExceededException for service response error code
	// "AccountChannelLimitExceededException".
	//
	// You have reached the maximum limit of active signaling channels for this
	// Amazon Web Services account in this region.
	ErrCodeAccountChannelLimitExceededException = "AccountChannelLimitExceededException"

	// ErrCodeAccountStreamLimitExceededException for service response error code
	// "AccountStreamLimitExceededException".
	//
	// The number of streams created for the account is too high.
	ErrCodeAccountStreamLimitExceededException = "AccountStreamLimitExceededException"

	// ErrCodeClientLimitExceededException for service response error code
	// "ClientLimitExceededException".
	//
	// Kinesis Video Streams has throttled the request because you have exceeded
	// the limit of allowed client calls. Try making the call later.
	ErrCodeClientLimitExceededException = "ClientLimitExceededException"

	// ErrCodeDeviceStreamLimitExceededException for service response error code
	// "DeviceStreamLimitExceededException".
	//
	// Not implemented.
	ErrCodeDeviceStreamLimitExceededException = "DeviceStreamLimitExceededException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// The value for this input parameter is invalid.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeInvalidDeviceException for service response error code
	// "InvalidDeviceException".
	//
	// Not implemented.
	ErrCodeInvalidDeviceException = "InvalidDeviceException"

	// ErrCodeInvalidResourceFormatException for service response error code
	// "InvalidResourceFormatException".
	//
	// The format of the StreamARN is invalid.
	ErrCodeInvalidResourceFormatException = "InvalidResourceFormatException"

	// ErrCodeNoDataRetentionException for service response error code
	// "NoDataRetentionException".
	//
	// The Stream data retention in hours is equal to zero.
	ErrCodeNoDataRetentionException = "NoDataRetentionException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// The caller is not authorized to perform this operation.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The resource is currently not available for this operation. New resources
	// cannot be created with the same name as existing resources. Also, resources
	// cannot be updated or deleted unless they are in an ACTIVE state.
	//
	// If this exception is returned, do not use it to determine whether the requested
	// resource already exists. Instead, it is recommended you use the resource-specific
	// describe API, for example, DescribeStream for video streams.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Amazon Kinesis Video Streams can't find the stream that you specified.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeStreamEdgeConfigurationNotFoundException for service response error code
	// "StreamEdgeConfigurationNotFoundException".
	//
	// The Exception rendered when the Amazon Kinesis Video Stream can't find a
	// stream's edge configuration that you specified.
	ErrCodeStreamEdgeConfigurationNotFoundException = "StreamEdgeConfigurationNotFoundException"

	// ErrCodeTagsPerResourceExceededLimitException for service response error code
	// "TagsPerResourceExceededLimitException".
	//
	// You have exceeded the limit of tags that you can associate with the resource.
	// A Kinesis video stream can support up to 50 tags.
	ErrCodeTagsPerResourceExceededLimitException = "TagsPerResourceExceededLimitException"

	// ErrCodeVersionMismatchException for service response error code
	// "VersionMismatchException".
	//
	// The stream version that you specified is not the latest version. To get the
	// latest version, use the DescribeStream (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_DescribeStream.html)
	// API.
	ErrCodeVersionMismatchException = "VersionMismatchException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                    newErrorAccessDeniedException,
	"AccountChannelLimitExceededException":     newErrorAccountChannelLimitExceededException,
	"AccountStreamLimitExceededException":      newErrorAccountStreamLimitExceededException,
	"ClientLimitExceededException":             newErrorClientLimitExceededException,
	"DeviceStreamLimitExceededException":       newErrorDeviceStreamLimitExceededException,
	"InvalidArgumentException":                 newErrorInvalidArgumentException,
	"InvalidDeviceException":                   newErrorInvalidDeviceException,
	"InvalidResourceFormatException":           newErrorInvalidResourceFormatException,
	"NoDataRetentionException":                 newErrorNoDataRetentionException,
	"NotAuthorizedException":                   newErrorNotAuthorizedException,
	"ResourceInUseException":                   newErrorResourceInUseException,
	"ResourceNotFoundException":                newErrorResourceNotFoundException,
	"StreamEdgeConfigurationNotFoundException": newErrorStreamEdgeConfigurationNotFoundException,
	"TagsPerResourceExceededLimitException":    newErrorTagsPerResourceExceededLimitException,
	"VersionMismatchException":                 newErrorVersionMismatchException,
}
