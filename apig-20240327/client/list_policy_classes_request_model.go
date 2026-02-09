// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyClassesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceId(v string) *ListPolicyClassesRequest
	GetAttachResourceId() *string
	SetAttachResourceType(v string) *ListPolicyClassesRequest
	GetAttachResourceType() *string
	SetDirection(v string) *ListPolicyClassesRequest
	GetDirection() *string
	SetGatewayId(v string) *ListPolicyClassesRequest
	GetGatewayId() *string
	SetPageNumber(v int32) *ListPolicyClassesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPolicyClassesRequest
	GetPageSize() *int32
	SetType(v string) *ListPolicyClassesRequest
	GetType() *string
}

type ListPolicyClassesRequest struct {
	// The resource ID to attach the policy
	//
	// example:
	//
	// route-001
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// The supported mount point type. Valid values:
	//
	// 	- HttpApi: an HTTP API
	//
	// 	- Operation: an operation in an HTTP API
	//
	// 	- GatewayRoute: a route
	//
	// 	- GatewayService: a service
	//
	// 	- GatewayServicePort: a service port
	//
	// 	- Domain: a domain name
	//
	// 	- Gateway: an instance
	//
	// example:
	//
	// Operation
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The direction of traffic on which the policy takes effect. Valid values:
	//
	// 	- OutBound
	//
	// 	- InBound
	//
	// 	- Both
	//
	// example:
	//
	// InBound
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The gateway ID
	//
	// example:
	//
	// gw-001
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The template type.
	//
	// example:
	//
	// FlowControl
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPolicyClassesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyClassesRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyClassesRequest) GetAttachResourceId() *string {
	return s.AttachResourceId
}

func (s *ListPolicyClassesRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *ListPolicyClassesRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListPolicyClassesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListPolicyClassesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPolicyClassesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPolicyClassesRequest) GetType() *string {
	return s.Type
}

func (s *ListPolicyClassesRequest) SetAttachResourceId(v string) *ListPolicyClassesRequest {
	s.AttachResourceId = &v
	return s
}

func (s *ListPolicyClassesRequest) SetAttachResourceType(v string) *ListPolicyClassesRequest {
	s.AttachResourceType = &v
	return s
}

func (s *ListPolicyClassesRequest) SetDirection(v string) *ListPolicyClassesRequest {
	s.Direction = &v
	return s
}

func (s *ListPolicyClassesRequest) SetGatewayId(v string) *ListPolicyClassesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListPolicyClassesRequest) SetPageNumber(v int32) *ListPolicyClassesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyClassesRequest) SetPageSize(v int32) *ListPolicyClassesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyClassesRequest) SetType(v string) *ListPolicyClassesRequest {
	s.Type = &v
	return s
}

func (s *ListPolicyClassesRequest) Validate() error {
	return dara.Validate(s)
}
