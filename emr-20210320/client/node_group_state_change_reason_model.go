// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeGroupStateChangeReason interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *NodeGroupStateChangeReason
	GetCode() *string
	SetMessage(v string) *NodeGroupStateChangeReason
	GetMessage() *string
}

type NodeGroupStateChangeReason struct {
	// The status code for the state change. Valid values include:
	//
	// - `UserRequest`: A user request triggered the change.
	//
	// - `OutOfStock`: The requested ECS instance type is out of stock.
	//
	// - `NotAuthorized`: The request lacks the required permissions.
	//
	// - `QuotaExceeded`: The request exceeds the resource quota.
	//
	// - `OperationDenied`: The system denied the operation.
	//
	// - `AccountException`: An account exception occurred.
	//
	// - `NodeFailure`: An ECS node failed.
	//
	// - `BootstrapFailure`: The bootstrap process failed.
	//
	// - `ValidationFail`: The request parameters failed validation.
	//
	// - `ServiceFailure`: A dependent service failed.
	//
	// - `InternalError`: An unexpected internal error occurred.
	//
	// example:
	//
	// UserRequest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A human-readable message that provides details about the state change.
	//
	// example:
	//
	// The instance type is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s NodeGroupStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s NodeGroupStateChangeReason) GoString() string {
	return s.String()
}

func (s *NodeGroupStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *NodeGroupStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *NodeGroupStateChangeReason) SetCode(v string) *NodeGroupStateChangeReason {
	s.Code = &v
	return s
}

func (s *NodeGroupStateChangeReason) SetMessage(v string) *NodeGroupStateChangeReason {
	s.Message = &v
	return s
}

func (s *NodeGroupStateChangeReason) Validate() error {
	return dara.Validate(s)
}
