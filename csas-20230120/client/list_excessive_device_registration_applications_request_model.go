// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExcessiveDeviceRegistrationApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetApplicationIds() []*string
	SetCurrentPage(v int64) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetCurrentPage() *int64
	SetDepartment(v string) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetDepartment() *string
	SetDeviceTag(v string) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetDeviceTag() *string
	SetHostname(v string) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetHostname() *string
	SetMac(v string) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetMac() *string
	SetPageSize(v int64) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetPageSize() *int64
	SetSaseUserId(v string) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetSaseUserId() *string
	SetStatuses(v []*string) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetStatuses() []*string
	SetUsername(v string) *ListExcessiveDeviceRegistrationApplicationsRequest
	GetUsername() *string
}

type ListExcessiveDeviceRegistrationApplicationsRequest struct {
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
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
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
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string   `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Statuses   []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	Username   *string   `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListExcessiveDeviceRegistrationApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExcessiveDeviceRegistrationApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetHostname() *string {
	return s.Hostname
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetMac() *string {
	return s.Mac
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) GetUsername() *string {
	return s.Username
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetApplicationIds(v []*string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetCurrentPage(v int64) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetDepartment(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Department = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetDeviceTag(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.DeviceTag = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetHostname(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Hostname = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetMac(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Mac = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetPageSize(v int64) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetSaseUserId(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetStatuses(v []*string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Statuses = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetUsername(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Username = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
