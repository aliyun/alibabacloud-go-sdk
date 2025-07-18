// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExcessiveDeviceRegistrationApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) *ListExcessiveDeviceRegistrationApplicationsResponseBody
	GetApplications() []*ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications
	SetRequestId(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListExcessiveDeviceRegistrationApplicationsResponseBody
	GetTotalNum() *int64
}

type ListExcessiveDeviceRegistrationApplicationsResponseBody struct {
	Applications []*ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListExcessiveDeviceRegistrationApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExcessiveDeviceRegistrationApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) GetApplications() []*ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) SetApplications(v []*ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) *ListExcessiveDeviceRegistrationApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) SetRequestId(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) SetTotalNum(v int64) *ListExcessiveDeviceRegistrationApplicationsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications struct {
	// example:
	//
	// reg-application-0f4a127b7e78****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 2023-07-17 18:46:55
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// false
	IsUsed *bool `json:"IsUsed,omitempty" xml:"IsUsed,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// Approved
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetDepartment() *string {
	return s.Department
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetDescription() *string {
	return s.Description
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetHostname() *string {
	return s.Hostname
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetIsUsed() *bool {
	return s.IsUsed
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetMac() *string {
	return s.Mac
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetStatus() *string {
	return s.Status
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GetUsername() *string {
	return s.Username
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetApplicationId(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetCreateTime(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetDepartment(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Department = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetDescription(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetDeviceTag(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.DeviceTag = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetDeviceType(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.DeviceType = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetHostname(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Hostname = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetIsUsed(v bool) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.IsUsed = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetMac(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Mac = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetSaseUserId(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.SaseUserId = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetStatus(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetUsername(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Username = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
