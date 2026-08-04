// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessModes(v string) *ListPrivateAccessApplicationsRequest
	GetAccessModes() *string
	SetAddress(v string) *ListPrivateAccessApplicationsRequest
	GetAddress() *string
	SetApplicationIds(v []*string) *ListPrivateAccessApplicationsRequest
	GetApplicationIds() []*string
	SetConnectorId(v string) *ListPrivateAccessApplicationsRequest
	GetConnectorId() *string
	SetCurrentPage(v int32) *ListPrivateAccessApplicationsRequest
	GetCurrentPage() *int32
	SetName(v string) *ListPrivateAccessApplicationsRequest
	GetName() *string
	SetPageSize(v int32) *ListPrivateAccessApplicationsRequest
	GetPageSize() *int32
	SetPolicyId(v string) *ListPrivateAccessApplicationsRequest
	GetPolicyId() *string
	SetStatus(v string) *ListPrivateAccessApplicationsRequest
	GetStatus() *string
	SetTagId(v string) *ListPrivateAccessApplicationsRequest
	GetTagId() *string
}

type ListPrivateAccessApplicationsRequest struct {
	// The access mode. Valid values:
	//
	// - **app**: application access. Filters applications that support application access.
	//
	// - **browser**: browser access. Filters applications that support browser access.
	//
	// example:
	//
	// app
	AccessModes *string `json:"AccessModes,omitempty" xml:"AccessModes,omitempty"`
	// The address of the internal-facing access application. The address is 1 to 128 characters in length and supports IPv4 addresses, CIDR blocks, domain names, and wildcard domain names.
	//
	// example:
	//
	// 192.168.0.0/16
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The IDs of internal-facing access applications. You can specify up to 100 application IDs.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The connector ID. You can obtain the ID from the [ListConnectors](~~ListConnectors~~) operation.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The page number of the current page displayed in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the internal-facing access application. The name is 1 to 128 characters in length, supports Chinese and uppercase and lowercase letters, and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page in a paged query. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the internal-facing access policy. You can obtain the ID from the following operations:
	//
	// - [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): queries internal-facing access policies in batches.
	//
	// - [CreatePrivateAccessPolicy](~~CreatePrivateAccessPolicy~~): creates an internal-facing access policy.
	//
	// example:
	//
	// pa-policy-54a7838a48bf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The status of the internal-facing access application. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the internal-facing access tag. You can obtain the ID from the following operations:
	//
	// - [ListPrivateAccessTags](~~ListPrivateAccessTags~~): queries internal-facing access tags in batches.
	//
	// - [CreatePrivateAccessTag](~~CreatePrivateAccessTag~~): creates an internal-facing access tag.
	//
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListPrivateAccessApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsRequest) GetAccessModes() *string {
	return s.AccessModes
}

func (s *ListPrivateAccessApplicationsRequest) GetAddress() *string {
	return s.Address
}

func (s *ListPrivateAccessApplicationsRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListPrivateAccessApplicationsRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *ListPrivateAccessApplicationsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListPrivateAccessApplicationsRequest) GetName() *string {
	return s.Name
}

func (s *ListPrivateAccessApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPrivateAccessApplicationsRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPrivateAccessApplicationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListPrivateAccessApplicationsRequest) GetTagId() *string {
	return s.TagId
}

func (s *ListPrivateAccessApplicationsRequest) SetAccessModes(v string) *ListPrivateAccessApplicationsRequest {
	s.AccessModes = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetAddress(v string) *ListPrivateAccessApplicationsRequest {
	s.Address = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetApplicationIds(v []*string) *ListPrivateAccessApplicationsRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetConnectorId(v string) *ListPrivateAccessApplicationsRequest {
	s.ConnectorId = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetCurrentPage(v int32) *ListPrivateAccessApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetName(v string) *ListPrivateAccessApplicationsRequest {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetPageSize(v int32) *ListPrivateAccessApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetPolicyId(v string) *ListPrivateAccessApplicationsRequest {
	s.PolicyId = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetStatus(v string) *ListPrivateAccessApplicationsRequest {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetTagId(v string) *ListPrivateAccessApplicationsRequest {
	s.TagId = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
