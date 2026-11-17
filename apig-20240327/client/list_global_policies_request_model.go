// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGlobalPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceType(v string) *ListGlobalPoliciesRequest
	GetAttachResourceType() *string
	SetClassName(v string) *ListGlobalPoliciesRequest
	GetClassName() *string
	SetEnable(v bool) *ListGlobalPoliciesRequest
	GetEnable() *bool
	SetEnvironmentId(v string) *ListGlobalPoliciesRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *ListGlobalPoliciesRequest
	GetGatewayId() *string
	SetGlobalPolicyType(v string) *ListGlobalPoliciesRequest
	GetGlobalPolicyType() *string
	SetIpAccessControlContent(v string) *ListGlobalPoliciesRequest
	GetIpAccessControlContent() *string
	SetIpAccessControlProtocolLayer(v string) *ListGlobalPoliciesRequest
	GetIpAccessControlProtocolLayer() *string
	SetIpAccessControlResourceName(v string) *ListGlobalPoliciesRequest
	GetIpAccessControlResourceName() *string
	SetIpAccessControlType(v string) *ListGlobalPoliciesRequest
	GetIpAccessControlType() *string
	SetName(v string) *ListGlobalPoliciesRequest
	GetName() *string
	SetPageNumber(v int32) *ListGlobalPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGlobalPoliciesRequest
	GetPageSize() *int32
}

type ListGlobalPoliciesRequest struct {
	// Filters by mounted resource type. This parameter takes effect only for the BlackOrWhite policy type. By default, the Gateway, Domain, and Route types are queried.
	//
	// example:
	//
	// Gateway
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// The policy class name. You can specify multiple values separated by commas.
	//
	// example:
	//
	// JWTAuth,OIDCAuth,ExternalZAuth
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// Filters by enabling status.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// Filters by environment ID.
	//
	// example:
	//
	// env-xxxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Filters by gateway ID at the attachment level.
	//
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The global policy type.
	//
	// example:
	//
	// GlobalAuth
	GlobalPolicyType *string `json:"globalPolicyType,omitempty" xml:"globalPolicyType,omitempty"`
	// Filters by IP content by performing a LIKE query on the configuration.
	//
	// example:
	//
	// 1.2.3.4
	IpAccessControlContent *string `json:"ipAccessControlContent,omitempty" xml:"ipAccessControlContent,omitempty"`
	// Filters by protocol layer. If you specify L4, configurations containing L4 are queried. If you specify L7, configurations that do not contain L4 are queried.
	//
	// example:
	//
	// L7
	IpAccessControlProtocolLayer *string `json:"ipAccessControlProtocolLayer,omitempty" xml:"ipAccessControlProtocolLayer,omitempty"`
	// Filters by resource name by performing a LIKE query on the configuration.
	//
	// example:
	//
	// my-route
	IpAccessControlResourceName *string `json:"ipAccessControlResourceName,omitempty" xml:"ipAccessControlResourceName,omitempty"`
	// Filters by IP access control type by performing a LIKE query on the configuration.
	//
	// example:
	//
	// White / Black
	IpAccessControlType *string `json:"ipAccessControlType,omitempty" xml:"ipAccessControlType,omitempty"`
	// Filters by exact policy name.
	//
	// example:
	//
	// my-jwt-auth
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
}

func (s ListGlobalPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListGlobalPoliciesRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *ListGlobalPoliciesRequest) GetClassName() *string {
	return s.ClassName
}

func (s *ListGlobalPoliciesRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ListGlobalPoliciesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListGlobalPoliciesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGlobalPoliciesRequest) GetGlobalPolicyType() *string {
	return s.GlobalPolicyType
}

func (s *ListGlobalPoliciesRequest) GetIpAccessControlContent() *string {
	return s.IpAccessControlContent
}

func (s *ListGlobalPoliciesRequest) GetIpAccessControlProtocolLayer() *string {
	return s.IpAccessControlProtocolLayer
}

func (s *ListGlobalPoliciesRequest) GetIpAccessControlResourceName() *string {
	return s.IpAccessControlResourceName
}

func (s *ListGlobalPoliciesRequest) GetIpAccessControlType() *string {
	return s.IpAccessControlType
}

func (s *ListGlobalPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListGlobalPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGlobalPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGlobalPoliciesRequest) SetAttachResourceType(v string) *ListGlobalPoliciesRequest {
	s.AttachResourceType = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetClassName(v string) *ListGlobalPoliciesRequest {
	s.ClassName = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetEnable(v bool) *ListGlobalPoliciesRequest {
	s.Enable = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetEnvironmentId(v string) *ListGlobalPoliciesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetGatewayId(v string) *ListGlobalPoliciesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetGlobalPolicyType(v string) *ListGlobalPoliciesRequest {
	s.GlobalPolicyType = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetIpAccessControlContent(v string) *ListGlobalPoliciesRequest {
	s.IpAccessControlContent = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetIpAccessControlProtocolLayer(v string) *ListGlobalPoliciesRequest {
	s.IpAccessControlProtocolLayer = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetIpAccessControlResourceName(v string) *ListGlobalPoliciesRequest {
	s.IpAccessControlResourceName = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetIpAccessControlType(v string) *ListGlobalPoliciesRequest {
	s.IpAccessControlType = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetName(v string) *ListGlobalPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetPageNumber(v int32) *ListGlobalPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGlobalPoliciesRequest) SetPageSize(v int32) *ListGlobalPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGlobalPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
