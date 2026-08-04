// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody
	GetApplications() []*UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications
	SetRequestId(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody
	GetRequestId() *string
}

type UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody struct {
	// List of device registration applications that exceed your quota.
	Applications []*UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// ID of the request.
	//
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) GetApplications() []*UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	return s.Applications
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) SetApplications(v []*UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody {
	s.Applications = v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) SetRequestId(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications struct {
	// ID of the device registration application.
	//
	// example:
	//
	// reg-application-0f4a127b7e78****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Time when the device registration application was created.
	//
	// example:
	//
	// 2023-07-17 18:46:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Department to which the user belongs.
	//
	// example:
	//
	// 测试部
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// This field indicates the reason for the excessive device registration request.
	//
	// example:
	//
	// 这是一条超额注册申请
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID of the endpoint device.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// Operating system of the endpoint device. Valid values:
	//
	// - **Windows**: Windows operating system.
	//
	// - **macOS**: macOS operating system.
	//
	// - **Linux**: Linux operating system.
	//
	// - **Android**: Android operating system.
	//
	// - **iOS**: iOS operating system.
	//
	// - **Windows_Wuying**: Alibaba Cloud Cloud Desktop operating system.
	//
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// Name of the endpoint device.
	//
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Indicates whether the device registration application has been used. Valid values:
	//
	// - **true**: Used.
	//
	// - **false**: Not used.
	//
	// example:
	//
	// false
	IsUsed *bool `json:"IsUsed,omitempty" xml:"IsUsed,omitempty"`
	// MAC address of the endpoint device.
	//
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// User ID.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// Status of the device registration application. Valid values:
	//
	// - **Pending**: Pending review.
	//
	// - **Approved**: Approved.
	//
	// - **Rejected**: Rejected.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Username.
	//
	// example:
	//
	// 王先生
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetDepartment() *string {
	return s.Department
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetDescription() *string {
	return s.Description
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetDeviceTag() *string {
	return s.DeviceTag
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetDeviceType() *string {
	return s.DeviceType
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetHostname() *string {
	return s.Hostname
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetIsUsed() *bool {
	return s.IsUsed
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetMac() *string {
	return s.Mac
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetStatus() *string {
	return s.Status
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GetUsername() *string {
	return s.Username
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetApplicationId(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetCreateTime(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetDepartment(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Department = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetDescription(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetDeviceTag(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.DeviceTag = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetDeviceType(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.DeviceType = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetHostname(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Hostname = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetIsUsed(v bool) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.IsUsed = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetMac(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Mac = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetSaseUserId(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.SaseUserId = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetStatus(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetUsername(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Username = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
