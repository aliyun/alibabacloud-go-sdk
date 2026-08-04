// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ListUserApplicationsRequest
	GetAddress() *string
	SetCurrentPage(v int32) *ListUserApplicationsRequest
	GetCurrentPage() *int32
	SetName(v string) *ListUserApplicationsRequest
	GetName() *string
	SetPageSize(v int32) *ListUserApplicationsRequest
	GetPageSize() *int32
	SetSaseUserId(v string) *ListUserApplicationsRequest
	GetSaseUserId() *string
}

type ListUserApplicationsRequest struct {
	// The address of the private access application. The address must be 1 to 128 characters in length. It supports IPv4 addresses, CIDR blocks, domain names, and wildcard domain names. Fuzzy search is supported.
	//
	// example:
	//
	// sase.alibaba.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the private access application. The name must be 1 to 128 characters in length. It can contain Chinese characters, letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page for pagination. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
}

func (s ListUserApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsRequest) GetAddress() *string {
	return s.Address
}

func (s *ListUserApplicationsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUserApplicationsRequest) GetName() *string {
	return s.Name
}

func (s *ListUserApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserApplicationsRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListUserApplicationsRequest) SetAddress(v string) *ListUserApplicationsRequest {
	s.Address = &v
	return s
}

func (s *ListUserApplicationsRequest) SetCurrentPage(v int32) *ListUserApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserApplicationsRequest) SetName(v string) *ListUserApplicationsRequest {
	s.Name = &v
	return s
}

func (s *ListUserApplicationsRequest) SetPageSize(v int32) *ListUserApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserApplicationsRequest) SetSaseUserId(v string) *ListUserApplicationsRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListUserApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
