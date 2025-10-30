// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUninstallApplicationsStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*UpdateUninstallApplicationsStatusResponseBodyApplications) *UpdateUninstallApplicationsStatusResponseBody
	GetApplications() []*UpdateUninstallApplicationsStatusResponseBodyApplications
	SetRequestId(v string) *UpdateUninstallApplicationsStatusResponseBody
	GetRequestId() *string
}

type UpdateUninstallApplicationsStatusResponseBody struct {
	Applications []*UpdateUninstallApplicationsStatusResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 9B769522-D50C-5978-8981-52BE800D6099
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUninstallApplicationsStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUninstallApplicationsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUninstallApplicationsStatusResponseBody) GetApplications() []*UpdateUninstallApplicationsStatusResponseBodyApplications {
	return s.Applications
}

func (s *UpdateUninstallApplicationsStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUninstallApplicationsStatusResponseBody) SetApplications(v []*UpdateUninstallApplicationsStatusResponseBodyApplications) *UpdateUninstallApplicationsStatusResponseBody {
	s.Applications = v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBody) SetRequestId(v string) *UpdateUninstallApplicationsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBody) Validate() error {
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

type UpdateUninstallApplicationsStatusResponseBodyApplications struct {
	// example:
	//
	// uninstall-app-6646831ac314****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 2023-07-17 18:46:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// example:
	//
	// Windows
	DevType *string `json:"DevType,omitempty" xml:"DevType,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IdpName  *string `json:"IdpName,omitempty" xml:"IdpName,omitempty"`
	// example:
	//
	// false
	IsUninstall *bool `json:"IsUninstall,omitempty" xml:"IsUninstall,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac    *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
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

func (s UpdateUninstallApplicationsStatusResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s UpdateUninstallApplicationsStatusResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetDepartment() *string {
	return s.Department
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetDevTag() *string {
	return s.DevTag
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetDevType() *string {
	return s.DevType
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetHostname() *string {
	return s.Hostname
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetIdpName() *string {
	return s.IdpName
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetIsUninstall() *bool {
	return s.IsUninstall
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetMac() *string {
	return s.Mac
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetReason() *string {
	return s.Reason
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetStatus() *string {
	return s.Status
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) GetUsername() *string {
	return s.Username
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetApplicationId(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetCreateTime(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetDepartment(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.Department = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetDevTag(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.DevTag = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetDevType(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.DevType = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetHostname(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.Hostname = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetIdpName(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.IdpName = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetIsUninstall(v bool) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.IsUninstall = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetMac(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.Mac = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetReason(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.Reason = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetSaseUserId(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.SaseUserId = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetStatus(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) SetUsername(v string) *UpdateUninstallApplicationsStatusResponseBodyApplications {
	s.Username = &v
	return s
}

func (s *UpdateUninstallApplicationsStatusResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
