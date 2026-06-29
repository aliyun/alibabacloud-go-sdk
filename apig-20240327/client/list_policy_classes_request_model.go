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
	// The ID of the attached resource.
	//
	// example:
	//
	// route-001
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// The type of resource that the policy can be attached to. Valid values:
	//
	// - HttpApi: HttpApi.
	//
	// - Operation: an operation of an HttpApi.
	//
	// - GatewayRoute: a gateway route.
	//
	// - GatewayService: a gateway service.
	//
	// - GatewayServicePort: a gateway service port.
	//
	// - Domain: a gateway domain name.
	//
	// - Gateway: a gateway.
	//
	// example:
	//
	// Operation
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The direction of the policy. Valid values:
	//
	// - OutBound: outbound.
	//
	// - InBound: inbound.
	//
	// - Both: both directions.
	//
	// example:
	//
	// InBound
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// gw-001
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The page number. Default value: 1.
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
	// The policy templatetype.
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
