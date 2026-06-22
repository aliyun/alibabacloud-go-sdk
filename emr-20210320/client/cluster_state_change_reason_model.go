// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterStateChangeReason interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ClusterStateChangeReason
	GetCode() *string
	SetMessage(v string) *ClusterStateChangeReason
	GetMessage() *string
}

type ClusterStateChangeReason struct {
	// The status code for the state change. Valid values:
	//
	// - UserRequest: A user request triggered the state change.
	//
	// - OutOfStock: The requested ECS instance type is out of stock.
	//
	// - NotAuthorized: The operation was denied due to insufficient permission.
	//
	// - QuotaExceeded: The request exceeded a service quota.
	//
	// - OperationDenied: The operation was denied.
	//
	// - AccountException: An account-related exception occurred.
	//
	// - NodeFailure: An ECS node failed.
	//
	// - BootstrapFailure: A bootstrap action failed.
	//
	// - ValidationFail: The request failed business logic validation.
	//
	// - ServiceFailure: A dependent service failed.
	//
	// - InternalError: An internal error occurred.
	//
	// example:
	//
	// OutOfStock
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A human-readable message detailing the cluster state change.
	//
	// example:
	//
	// The requested resource is sold out in the specified zone, try other types of resources or other regions and zones.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ClusterStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s ClusterStateChangeReason) GoString() string {
	return s.String()
}

func (s *ClusterStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *ClusterStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *ClusterStateChangeReason) SetCode(v string) *ClusterStateChangeReason {
	s.Code = &v
	return s
}

func (s *ClusterStateChangeReason) SetMessage(v string) *ClusterStateChangeReason {
	s.Message = &v
	return s
}

func (s *ClusterStateChangeReason) Validate() error {
	return dara.Validate(s)
}
