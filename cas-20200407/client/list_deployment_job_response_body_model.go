// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListDeploymentJobResponseBody
	GetCurrentPage() *int32
	SetData(v []*ListDeploymentJobResponseBodyData) *ListDeploymentJobResponseBody
	GetData() []*ListDeploymentJobResponseBodyData
	SetRequestId(v string) *ListDeploymentJobResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListDeploymentJobResponseBody
	GetShowSize() *int32
	SetTotal(v int64) *ListDeploymentJobResponseBody
	GetTotal() *int64
}

type ListDeploymentJobResponseBody struct {
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data returned by the API.
	Data []*ListDeploymentJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of deployment tasks per page. Default value: **50**.
	//
	// example:
	//
	// 50
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of deployment tasks returned.
	//
	// example:
	//
	// 7
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeploymentJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDeploymentJobResponseBody) GetData() []*ListDeploymentJobResponseBodyData {
	return s.Data
}

func (s *ListDeploymentJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentJobResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListDeploymentJobResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListDeploymentJobResponseBody) SetCurrentPage(v int32) *ListDeploymentJobResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListDeploymentJobResponseBody) SetData(v []*ListDeploymentJobResponseBodyData) *ListDeploymentJobResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentJobResponseBody) SetRequestId(v string) *ListDeploymentJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentJobResponseBody) SetShowSize(v int32) *ListDeploymentJobResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListDeploymentJobResponseBody) SetTotal(v int64) *ListDeploymentJobResponseBody {
	s.Total = &v
	return s
}

func (s *ListDeploymentJobResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeploymentJobResponseBodyData struct {
	// The domain names included in the certificate of the deployment task.
	//
	// example:
	//
	// aliyundoc1.com,aliyundoc2.com,aliyundoc3.com
	CertDomain *string `json:"CertDomain,omitempty" xml:"CertDomain,omitempty"`
	// The certificate type. Valid values:
	//
	// - **upload**: uploaded certificate
	//
	// - **buy**: purchased certificate
	//
	// - **free**: personal test certificate (supported only on the China site)
	//
	// example:
	//
	// upload
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The deletion flag. This parameter has no practical significance. Valid values:
	//
	// - **0**: not deleted.
	//
	// - **1**: deleted.
	//
	// example:
	//
	// 1
	Del *int32 `json:"Del,omitempty" xml:"Del,omitempty"`
	// The end time of the deployment task execution, in UNIX timestamp format (seconds).
	//
	// example:
	//
	// 1606482979000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the deployment task was created, in UNIX timestamp format (seconds).
	//
	// example:
	//
	// 1624343180000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the deployment task was last modified, in UNIX timestamp format (seconds).
	//
	// example:
	//
	// 1606482979000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The deployment task ID. You can use this ID to query the deployment task details and status.
	//
	// example:
	//
	// 19975
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the deployment task.
	//
	// example:
	//
	// cas-job-user-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task type. Valid values:
	//
	// - **cloud**: multi-cloud deployment task.
	//
	// - **user**: cloud service deployment task (excluding ECS instances).
	//
	// example:
	//
	// user
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the deployment task.
	//
	// example:
	//
	// job-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cloud services included in the deployment task resources.
	//
	// example:
	//
	// NLB
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Indicates whether the task contains a rollback worker. For example, if a cloud service in the task has been rolled back, the value **1*	- is returned. Valid values:
	//
	// - **0**: does not contain a rollback worker.
	//
	// - **1**: contains a rollback worker.
	//
	// example:
	//
	// 1
	Rollback *int32 `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
	// The scheduled time of the task, in UNIX timestamp format (seconds).
	//
	// example:
	//
	// 1606482979000
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The start time of the deployment task execution, in UNIX timestamp format (seconds).
	//
	// example:
	//
	// 1606482979000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The deployment task status. Valid values:
	//
	// - **pending**: pending execution.
	//
	// - **editing**: being edited.
	//
	// - **scheduling**: being scheduled.
	//
	// - **processing**: being deployed.
	//
	// - **error**: deployment failed.
	//
	// - **success**: deployment succeeded.
	//
	// example:
	//
	// scheduling
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user account ID (UID).
	//
	// example:
	//
	// 166688437*****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDeploymentJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResponseBodyData) GetCertDomain() *string {
	return s.CertDomain
}

func (s *ListDeploymentJobResponseBodyData) GetCertType() *string {
	return s.CertType
}

func (s *ListDeploymentJobResponseBodyData) GetDel() *int32 {
	return s.Del
}

func (s *ListDeploymentJobResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDeploymentJobResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDeploymentJobResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDeploymentJobResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListDeploymentJobResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDeploymentJobResponseBodyData) GetJobType() *string {
	return s.JobType
}

func (s *ListDeploymentJobResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListDeploymentJobResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *ListDeploymentJobResponseBodyData) GetRollback() *int32 {
	return s.Rollback
}

func (s *ListDeploymentJobResponseBodyData) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ListDeploymentJobResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDeploymentJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListDeploymentJobResponseBodyData) GetUserId() *int64 {
	return s.UserId
}

func (s *ListDeploymentJobResponseBodyData) SetCertDomain(v string) *ListDeploymentJobResponseBodyData {
	s.CertDomain = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetCertType(v string) *ListDeploymentJobResponseBodyData {
	s.CertType = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetDel(v int32) *ListDeploymentJobResponseBodyData {
	s.Del = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetEndTime(v string) *ListDeploymentJobResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetGmtCreate(v string) *ListDeploymentJobResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetGmtModified(v string) *ListDeploymentJobResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetId(v int64) *ListDeploymentJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetInstanceId(v string) *ListDeploymentJobResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetJobType(v string) *ListDeploymentJobResponseBodyData {
	s.JobType = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetName(v string) *ListDeploymentJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetProductName(v string) *ListDeploymentJobResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetRollback(v int32) *ListDeploymentJobResponseBodyData {
	s.Rollback = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetScheduleTime(v string) *ListDeploymentJobResponseBodyData {
	s.ScheduleTime = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetStartTime(v string) *ListDeploymentJobResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetStatus(v string) *ListDeploymentJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) SetUserId(v int64) *ListDeploymentJobResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListDeploymentJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
