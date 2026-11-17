// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceId(v string) *ListPoliciesRequest
	GetAttachResourceId() *string
	SetAttachResourceType(v string) *ListPoliciesRequest
	GetAttachResourceType() *string
	SetEnvironmentId(v string) *ListPoliciesRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *ListPoliciesRequest
	GetGatewayId() *string
	SetPageNumber(v int32) *ListPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPoliciesRequest
	GetPageSize() *int32
	SetWithAttachments(v bool) *ListPoliciesRequest
	GetWithAttachments() *bool
	SetWithSystemPolicy(v bool) *ListPoliciesRequest
	GetWithSystemPolicy() *bool
}

type ListPoliciesRequest struct {
	// The ID of the mount point.
	//
	// example:
	//
	// api-cq7l5s5lhtgi6qasrdc0
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// The type of mount points supported by the policy. Valid values:
	//
	// - HttpApi: HTTP API.
	//
	// - Operation: The operation of the HTTP API.
	//
	// - GatewayRoute: Gateway route.
	//
	// - GatewayService: Gateway service.
	//
	// - GatewayServicePort: Gateway service port.
	//
	// - Domain: Gateway domain name.
	//
	// - Gateway: Gateway.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The ID of the environment.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// gw-cq2fm65lhtgm9nrrl7ag
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Specifies whether to return mount information.
	//
	// example:
	//
	// false
	WithAttachments *bool `json:"withAttachments,omitempty" xml:"withAttachments,omitempty"`
	// Specifies whether the policy is a system policy.
	//
	// example:
	//
	// true
	WithSystemPolicy *bool `json:"withSystemPolicy,omitempty" xml:"withSystemPolicy,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) GetAttachResourceId() *string {
	return s.AttachResourceId
}

func (s *ListPoliciesRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *ListPoliciesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListPoliciesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPoliciesRequest) GetWithAttachments() *bool {
	return s.WithAttachments
}

func (s *ListPoliciesRequest) GetWithSystemPolicy() *bool {
	return s.WithSystemPolicy
}

func (s *ListPoliciesRequest) SetAttachResourceId(v string) *ListPoliciesRequest {
	s.AttachResourceId = &v
	return s
}

func (s *ListPoliciesRequest) SetAttachResourceType(v string) *ListPoliciesRequest {
	s.AttachResourceType = &v
	return s
}

func (s *ListPoliciesRequest) SetEnvironmentId(v string) *ListPoliciesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListPoliciesRequest) SetGatewayId(v string) *ListPoliciesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListPoliciesRequest) SetPageNumber(v int32) *ListPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesRequest) SetPageSize(v int32) *ListPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPoliciesRequest) SetWithAttachments(v bool) *ListPoliciesRequest {
	s.WithAttachments = &v
	return s
}

func (s *ListPoliciesRequest) SetWithSystemPolicy(v bool) *ListPoliciesRequest {
	s.WithSystemPolicy = &v
	return s
}

func (s *ListPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
