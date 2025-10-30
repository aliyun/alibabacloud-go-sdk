// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUninstallApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListUninstallApplicationsResponseBodyApplications) *ListUninstallApplicationsResponseBody
	GetApplications() []*ListUninstallApplicationsResponseBodyApplications
	SetRequestId(v string) *ListUninstallApplicationsResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListUninstallApplicationsResponseBody
	GetTotalNum() *int64
}

type ListUninstallApplicationsResponseBody struct {
	Applications []*ListUninstallApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListUninstallApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUninstallApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUninstallApplicationsResponseBody) GetApplications() []*ListUninstallApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListUninstallApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUninstallApplicationsResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListUninstallApplicationsResponseBody) SetApplications(v []*ListUninstallApplicationsResponseBodyApplications) *ListUninstallApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListUninstallApplicationsResponseBody) SetRequestId(v string) *ListUninstallApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUninstallApplicationsResponseBody) SetTotalNum(v int64) *ListUninstallApplicationsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListUninstallApplicationsResponseBody) Validate() error {
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

type ListUninstallApplicationsResponseBodyApplications struct {
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

func (s ListUninstallApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListUninstallApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetDepartment() *string {
	return s.Department
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetDevTag() *string {
	return s.DevTag
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetDevType() *string {
	return s.DevType
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetHostname() *string {
	return s.Hostname
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetIdpName() *string {
	return s.IdpName
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetIsUninstall() *bool {
	return s.IsUninstall
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetMac() *string {
	return s.Mac
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetReason() *string {
	return s.Reason
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetStatus() *string {
	return s.Status
}

func (s *ListUninstallApplicationsResponseBodyApplications) GetUsername() *string {
	return s.Username
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetApplicationId(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetCreateTime(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetDepartment(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.Department = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetDevTag(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.DevTag = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetDevType(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.DevType = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetHostname(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.Hostname = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetIdpName(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.IdpName = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetIsUninstall(v bool) *ListUninstallApplicationsResponseBodyApplications {
	s.IsUninstall = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetMac(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.Mac = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetReason(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.Reason = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetSaseUserId(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.SaseUserId = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetStatus(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) SetUsername(v string) *ListUninstallApplicationsResponseBodyApplications {
	s.Username = &v
	return s
}

func (s *ListUninstallApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
