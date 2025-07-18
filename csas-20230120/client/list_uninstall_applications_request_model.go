// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUninstallApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListUninstallApplicationsRequest
	GetApplicationIds() []*string
	SetCurrentPage(v int64) *ListUninstallApplicationsRequest
	GetCurrentPage() *int64
	SetDepartment(v string) *ListUninstallApplicationsRequest
	GetDepartment() *string
	SetHostname(v string) *ListUninstallApplicationsRequest
	GetHostname() *string
	SetMac(v string) *ListUninstallApplicationsRequest
	GetMac() *string
	SetPageSize(v int64) *ListUninstallApplicationsRequest
	GetPageSize() *int64
	SetStatuses(v []*string) *ListUninstallApplicationsRequest
	GetStatuses() []*string
	SetUsername(v string) *ListUninstallApplicationsRequest
	GetUsername() *string
}

type ListUninstallApplicationsRequest struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	Username *string   `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUninstallApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUninstallApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListUninstallApplicationsRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListUninstallApplicationsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListUninstallApplicationsRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListUninstallApplicationsRequest) GetHostname() *string {
	return s.Hostname
}

func (s *ListUninstallApplicationsRequest) GetMac() *string {
	return s.Mac
}

func (s *ListUninstallApplicationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUninstallApplicationsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListUninstallApplicationsRequest) GetUsername() *string {
	return s.Username
}

func (s *ListUninstallApplicationsRequest) SetApplicationIds(v []*string) *ListUninstallApplicationsRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListUninstallApplicationsRequest) SetCurrentPage(v int64) *ListUninstallApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUninstallApplicationsRequest) SetDepartment(v string) *ListUninstallApplicationsRequest {
	s.Department = &v
	return s
}

func (s *ListUninstallApplicationsRequest) SetHostname(v string) *ListUninstallApplicationsRequest {
	s.Hostname = &v
	return s
}

func (s *ListUninstallApplicationsRequest) SetMac(v string) *ListUninstallApplicationsRequest {
	s.Mac = &v
	return s
}

func (s *ListUninstallApplicationsRequest) SetPageSize(v int64) *ListUninstallApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUninstallApplicationsRequest) SetStatuses(v []*string) *ListUninstallApplicationsRequest {
	s.Statuses = v
	return s
}

func (s *ListUninstallApplicationsRequest) SetUsername(v string) *ListUninstallApplicationsRequest {
	s.Username = &v
	return s
}

func (s *ListUninstallApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
