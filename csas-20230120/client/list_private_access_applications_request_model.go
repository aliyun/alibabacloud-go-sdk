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
	AccessModes *string `json:"AccessModes,omitempty" xml:"AccessModes,omitempty"`
	// example:
	//
	// 192.168.0.0/16
	Address        *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ConnectorId    *string   `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// pa-policy-54a7838a48bf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
