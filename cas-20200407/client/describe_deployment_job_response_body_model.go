// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasContacts(v []*DescribeDeploymentJobResponseBodyCasContacts) *DescribeDeploymentJobResponseBody
	GetCasContacts() []*DescribeDeploymentJobResponseBodyCasContacts
	SetCertDomain(v string) *DescribeDeploymentJobResponseBody
	GetCertDomain() *string
	SetCertType(v string) *DescribeDeploymentJobResponseBody
	GetCertType() *string
	SetConfig(v string) *DescribeDeploymentJobResponseBody
	GetConfig() *string
	SetDel(v int32) *DescribeDeploymentJobResponseBody
	GetDel() *int32
	SetEndTime(v string) *DescribeDeploymentJobResponseBody
	GetEndTime() *string
	SetGmtCreate(v string) *DescribeDeploymentJobResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *DescribeDeploymentJobResponseBody
	GetGmtModified() *string
	SetId(v int64) *DescribeDeploymentJobResponseBody
	GetId() *int64
	SetInstanceId(v string) *DescribeDeploymentJobResponseBody
	GetInstanceId() *string
	SetJobType(v string) *DescribeDeploymentJobResponseBody
	GetJobType() *string
	SetName(v string) *DescribeDeploymentJobResponseBody
	GetName() *string
	SetProductName(v string) *DescribeDeploymentJobResponseBody
	GetProductName() *string
	SetRequestId(v string) *DescribeDeploymentJobResponseBody
	GetRequestId() *string
	SetRollback(v int32) *DescribeDeploymentJobResponseBody
	GetRollback() *int32
	SetScheduleTime(v string) *DescribeDeploymentJobResponseBody
	GetScheduleTime() *string
	SetStartTime(v string) *DescribeDeploymentJobResponseBody
	GetStartTime() *string
	SetStatus(v string) *DescribeDeploymentJobResponseBody
	GetStatus() *string
	SetUserId(v int64) *DescribeDeploymentJobResponseBody
	GetUserId() *int64
}

type DescribeDeploymentJobResponseBody struct {
	// The information about the contact.
	CasContacts []*DescribeDeploymentJobResponseBodyCasContacts `json:"CasContacts,omitempty" xml:"CasContacts,omitempty" type:"Repeated"`
	// The domain names bound to the certificate of the deployment task.
	//
	// example:
	//
	// example.aliyundoc.com,demo.aliyundoc.com
	CertDomain *string `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **upload**: uploaded certificate
	//
	// 	- **buy**: purchased certificate
	//
	// 	- **free**: free certificate available only on the China site (aliyun.com)
	//
	// example:
	//
	// buy
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The configurations of the deployment task.
	//
	// example:
	//
	// {\\"shareCertIds\\":[],\\"certIds\\":[12342649,12304338,12067351,9957285]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Indicates whether the deployment job was deleted. Valid values:
	//
	// 	- **0**: not deleted
	//
	// 	- **1**: deleted
	//
	// example:
	//
	// 1
	Del *int32 `json:"Del,omitempty" xml:"Del,omitempty"`
	// The end time of the deployment job. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679541809000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the deployment job was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679541809000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the deployment job was last modified. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679541809000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the deployment job.
	//
	// example:
	//
	// 8888
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the deployment task.
	//
	// example:
	//
	// 14dcc8afc7578e1f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the deployment job. Valid values:
	//
	// 	- **cloud**: multi-cloud deployment job.
	//
	// 	- **trustee**: hosted deployment job available only on the China site (aliyun.com).
	//
	// 	- **user**: cloud service deployment job. The cloud server is not included.
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the deployment task.
	//
	// example:
	//
	// auto-test-AXX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cloud services included in the deployment task.
	//
	// example:
	//
	// CDN
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deployment job includes the rollback worker. For example, if a cloud service in a deployment job has been rolled back, **1*	- is returned. Valid values:
	//
	// 	- **0**: The rollback worker is not included.
	//
	// 	- **1**: The rollback worker is included.
	//
	// example:
	//
	// 1
	Rollback *int32 `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	// The time when the deployment job was scheduled. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1678083209335
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The start time of the deployment job. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679541809000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the deployment job. Valid values:
	//
	// 	- **pending**
	//
	// 	- **editing**
	//
	// 	- **scheduling**
	//
	// 	- **processing**
	//
	// 	- **error**
	//
	// 	- **success**
	//
	// example:
	//
	// editing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Alibaba Cloud account in which the deployment job is created.
	//
	// example:
	//
	// 166688437XXXX785
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeDeploymentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobResponseBody) GetCasContacts() []*DescribeDeploymentJobResponseBodyCasContacts {
	return s.CasContacts
}

func (s *DescribeDeploymentJobResponseBody) GetCertDomain() *string {
	return s.CertDomain
}

func (s *DescribeDeploymentJobResponseBody) GetCertType() *string {
	return s.CertType
}

func (s *DescribeDeploymentJobResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeDeploymentJobResponseBody) GetDel() *int32 {
	return s.Del
}

func (s *DescribeDeploymentJobResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDeploymentJobResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDeploymentJobResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDeploymentJobResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeDeploymentJobResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDeploymentJobResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDeploymentJobResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeDeploymentJobResponseBody) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeDeploymentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeploymentJobResponseBody) GetRollback() *int32 {
	return s.Rollback
}

func (s *DescribeDeploymentJobResponseBody) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *DescribeDeploymentJobResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDeploymentJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDeploymentJobResponseBody) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeDeploymentJobResponseBody) SetCasContacts(v []*DescribeDeploymentJobResponseBodyCasContacts) *DescribeDeploymentJobResponseBody {
	s.CasContacts = v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetCertDomain(v string) *DescribeDeploymentJobResponseBody {
	s.CertDomain = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetCertType(v string) *DescribeDeploymentJobResponseBody {
	s.CertType = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetConfig(v string) *DescribeDeploymentJobResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetDel(v int32) *DescribeDeploymentJobResponseBody {
	s.Del = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetEndTime(v string) *DescribeDeploymentJobResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetGmtCreate(v string) *DescribeDeploymentJobResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetGmtModified(v string) *DescribeDeploymentJobResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetId(v int64) *DescribeDeploymentJobResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetInstanceId(v string) *DescribeDeploymentJobResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetJobType(v string) *DescribeDeploymentJobResponseBody {
	s.JobType = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetName(v string) *DescribeDeploymentJobResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetProductName(v string) *DescribeDeploymentJobResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetRequestId(v string) *DescribeDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetRollback(v int32) *DescribeDeploymentJobResponseBody {
	s.Rollback = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetScheduleTime(v string) *DescribeDeploymentJobResponseBody {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetStartTime(v string) *DescribeDeploymentJobResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetStatus(v string) *DescribeDeploymentJobResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) SetUserId(v int64) *DescribeDeploymentJobResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeDeploymentJobResponseBody) Validate() error {
	if s.CasContacts != nil {
		for _, item := range s.CasContacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeploymentJobResponseBodyCasContacts struct {
	// The email address of the contact.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the contact.
	//
	// example:
	//
	// 3304
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The phone number of the contact.
	//
	// example:
	//
	// 139****0000
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The name of the contact.
	//
	// example:
	//
	// zhangsan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDeploymentJobResponseBodyCasContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentJobResponseBodyCasContacts) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) GetEmail() *string {
	return s.Email
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) GetId() *string {
	return s.Id
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) GetMobile() *string {
	return s.Mobile
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) GetName() *string {
	return s.Name
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) SetEmail(v string) *DescribeDeploymentJobResponseBodyCasContacts {
	s.Email = &v
	return s
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) SetId(v string) *DescribeDeploymentJobResponseBodyCasContacts {
	s.Id = &v
	return s
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) SetMobile(v string) *DescribeDeploymentJobResponseBodyCasContacts {
	s.Mobile = &v
	return s
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) SetName(v string) *DescribeDeploymentJobResponseBodyCasContacts {
	s.Name = &v
	return s
}

func (s *DescribeDeploymentJobResponseBodyCasContacts) Validate() error {
	return dara.Validate(s)
}
