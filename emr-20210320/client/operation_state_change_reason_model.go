// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationStateChangeReason interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OperationStateChangeReason
	GetCode() *string
	SetMessage(v string) *OperationStateChangeReason
	GetMessage() *string
}

type OperationStateChangeReason struct {
	// The operation status code. Valid values:
	//
	// - `UserRequest`: The operation was initiated by a user.
	//
	// - `OutOfStock`: The requested ECS instance type is out of stock.
	//
	// - `NotAuthorized`: You are not authorized to perform the operation.
	//
	// - `QuotaExceeded`: A resource quota was exceeded.
	//
	// - `OperationDenied`: The operation was denied.
	//
	// - `AccountException`: An account exception occurred.
	//
	// - `NodeFailure`: An ECS node failed.
	//
	// - `BootstrapFailure`: A bootstrap action failed.
	//
	// - `ValidationFail`: The business logic validation failed.
	//
	// - `ServiceFailure`: A dependent service failed.
	//
	// - `InternalError`: An internal error occurred.
	//
	// example:
	//
	// OutOfStock
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A human-readable message that provides details about the state change.
	//
	// example:
	//
	// The requested resource is sold out in the specified zone, try other types of resources or other regions and zones.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s OperationStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s OperationStateChangeReason) GoString() string {
	return s.String()
}

func (s *OperationStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *OperationStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *OperationStateChangeReason) SetCode(v string) *OperationStateChangeReason {
	s.Code = &v
	return s
}

func (s *OperationStateChangeReason) SetMessage(v string) *OperationStateChangeReason {
	s.Message = &v
	return s
}

func (s *OperationStateChangeReason) Validate() error {
	return dara.Validate(s)
}
