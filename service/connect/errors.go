// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient permissions to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeContactFlowNotPublishedException for service response error code
	// "ContactFlowNotPublishedException".
	//
	// The flow has not been published.
	ErrCodeContactFlowNotPublishedException = "ContactFlowNotPublishedException"

	// ErrCodeContactNotFoundException for service response error code
	// "ContactNotFoundException".
	//
	// The contact with the specified ID is not active or does not exist. Applies
	// to Voice calls only, not to Chat or Task contacts.
	ErrCodeContactNotFoundException = "ContactNotFoundException"

	// ErrCodeDestinationNotAllowedException for service response error code
	// "DestinationNotAllowedException".
	//
	// Outbound calls to the destination number are not allowed.
	ErrCodeDestinationNotAllowedException = "DestinationNotAllowedException"

	// ErrCodeDuplicateResourceException for service response error code
	// "DuplicateResourceException".
	//
	// A resource with the specified name already exists.
	ErrCodeDuplicateResourceException = "DuplicateResourceException"

	// ErrCodeIdempotencyException for service response error code
	// "IdempotencyException".
	//
	// An entity with the same name already exists.
	ErrCodeIdempotencyException = "IdempotencyException"

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	//
	// Request processing failed because of an error or failure with the service.
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidContactFlowException for service response error code
	// "InvalidContactFlowException".
	//
	// The flow is not valid.
	ErrCodeInvalidContactFlowException = "InvalidContactFlowException"

	// ErrCodeInvalidContactFlowModuleException for service response error code
	// "InvalidContactFlowModuleException".
	//
	// The problems with the module. Please fix before trying again.
	ErrCodeInvalidContactFlowModuleException = "InvalidContactFlowModuleException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// One or more of the specified parameters are not valid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The request is not valid.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The allowed limit for the resource has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMaximumResultReturnedException for service response error code
	// "MaximumResultReturnedException".
	//
	// Maximum number (1000) of tags have been returned with current request. Consider
	// changing request parameters to get more tags.
	ErrCodeMaximumResultReturnedException = "MaximumResultReturnedException"

	// ErrCodeOutboundContactNotPermittedException for service response error code
	// "OutboundContactNotPermittedException".
	//
	// The contact is not permitted.
	ErrCodeOutboundContactNotPermittedException = "OutboundContactNotPermittedException"

	// ErrCodeOutputTypeNotFoundException for service response error code
	// "OutputTypeNotFoundException".
	//
	// Thrown for analyzed content when requested OutputType was not enabled for
	// a given contact. For example, if an OutputType.Raw was requested for a contact
	// that had `RedactedOnly` Redaction policy set in Contact flow.
	ErrCodeOutputTypeNotFoundException = "OutputTypeNotFoundException"

	// ErrCodePropertyValidationException for service response error code
	// "PropertyValidationException".
	//
	// The property is not valid.
	ErrCodePropertyValidationException = "PropertyValidationException"

	// ErrCodeResourceConflictException for service response error code
	// "ResourceConflictException".
	//
	// A resource already has that name.
	ErrCodeResourceConflictException = "ResourceConflictException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// That resource is already in use. Please try another.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceNotReadyException for service response error code
	// "ResourceNotReadyException".
	//
	// The resource is not ready.
	ErrCodeResourceNotReadyException = "ResourceNotReadyException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The service quota has been exceeded.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The throttling limit has been exceeded.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Displayed when rate-related API limits are exceeded.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUserNotFoundException for service response error code
	// "UserNotFoundException".
	//
	// No user with the specified credentials was found in the Amazon Connect instance.
	ErrCodeUserNotFoundException = "UserNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                newErrorAccessDeniedException,
	"ContactFlowNotPublishedException":     newErrorContactFlowNotPublishedException,
	"ContactNotFoundException":             newErrorContactNotFoundException,
	"DestinationNotAllowedException":       newErrorDestinationNotAllowedException,
	"DuplicateResourceException":           newErrorDuplicateResourceException,
	"IdempotencyException":                 newErrorIdempotencyException,
	"InternalServiceException":             newErrorInternalServiceException,
	"InvalidContactFlowException":          newErrorInvalidContactFlowException,
	"InvalidContactFlowModuleException":    newErrorInvalidContactFlowModuleException,
	"InvalidParameterException":            newErrorInvalidParameterException,
	"InvalidRequestException":              newErrorInvalidRequestException,
	"LimitExceededException":               newErrorLimitExceededException,
	"MaximumResultReturnedException":       newErrorMaximumResultReturnedException,
	"OutboundContactNotPermittedException": newErrorOutboundContactNotPermittedException,
	"OutputTypeNotFoundException":          newErrorOutputTypeNotFoundException,
	"PropertyValidationException":          newErrorPropertyValidationException,
	"ResourceConflictException":            newErrorResourceConflictException,
	"ResourceInUseException":               newErrorResourceInUseException,
	"ResourceNotFoundException":            newErrorResourceNotFoundException,
	"ResourceNotReadyException":            newErrorResourceNotReadyException,
	"ServiceQuotaExceededException":        newErrorServiceQuotaExceededException,
	"ThrottlingException":                  newErrorThrottlingException,
	"TooManyRequestsException":             newErrorTooManyRequestsException,
	"UserNotFoundException":                newErrorUserNotFoundException,
}
