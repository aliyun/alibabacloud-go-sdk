// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
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
