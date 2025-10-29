// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody
	GetJobs() []*ListJobsResponseBodyJobs
	SetPageNumber(v int32) *ListJobsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListJobsResponseBody
	GetTotalCount() *int32
}

type ListJobsResponseBody struct {
	Jobs []*ListJobsResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 882304F9-6DB1-5593-A719-33473D082B9C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 11
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetJobs() []*ListJobsResponseBodyJobs {
	return s.Jobs
}

func (s *ListJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobsResponseBodyJobs struct {
	Config *ListJobsResponseBodyJobsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 2022-07-05T02:13:43Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ElapsedTime *int64  `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	ExecuteType *string `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// example:
	//
	// true
	IsPassAssertCheck *bool `json:"isPassAssertCheck,omitempty" xml:"isPassAssertCheck,omitempty"`
	// example:
	//
	// job-433aead756057fff9e4dca57b147c
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// Errored
	Status       *string                           `json:"status,omitempty" xml:"status,omitempty"`
	StatusDetail map[string]*JobsStatusDetailValue `json:"statusDetail,omitempty" xml:"statusDetail,omitempty"`
	// example:
	//
	// task-518855d9a058c1176866c2c3efb
	TaskId                   *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) GetConfig() *ListJobsResponseBodyJobsConfig {
	return s.Config
}

func (s *ListJobsResponseBodyJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListJobsResponseBodyJobs) GetDescription() *string {
	return s.Description
}

func (s *ListJobsResponseBodyJobs) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *ListJobsResponseBodyJobs) GetExecuteType() *string {
	return s.ExecuteType
}

func (s *ListJobsResponseBodyJobs) GetIsPassAssertCheck() *bool {
	return s.IsPassAssertCheck
}

func (s *ListJobsResponseBodyJobs) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListJobsResponseBodyJobs) GetStatusDetail() map[string]*JobsStatusDetailValue {
	return s.StatusDetail
}

func (s *ListJobsResponseBodyJobs) GetTaskId() *string {
	return s.TaskId
}

func (s *ListJobsResponseBodyJobs) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *ListJobsResponseBodyJobs) SetConfig(v *ListJobsResponseBodyJobsConfig) *ListJobsResponseBodyJobs {
	s.Config = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetCreateTime(v string) *ListJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetDescription(v string) *ListJobsResponseBodyJobs {
	s.Description = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetElapsedTime(v int64) *ListJobsResponseBodyJobs {
	s.ElapsedTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetExecuteType(v string) *ListJobsResponseBodyJobs {
	s.ExecuteType = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetIsPassAssertCheck(v bool) *ListJobsResponseBodyJobs {
	s.IsPassAssertCheck = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobId(v string) *ListJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatus(v string) *ListJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatusDetail(v map[string]*JobsStatusDetailValue) *ListJobsResponseBodyJobs {
	s.StatusDetail = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetTaskId(v string) *ListJobsResponseBodyJobs {
	s.TaskId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetTerraformProviderVersion(v string) *ListJobsResponseBodyJobs {
	s.TerraformProviderVersion = &v
	return s
}

func (s *ListJobsResponseBodyJobs) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyJobsConfig struct {
	IsDestroy         *bool   `json:"isDestroy,omitempty" xml:"isDestroy,omitempty"`
	ModuleDescription *string `json:"moduleDescription,omitempty" xml:"moduleDescription,omitempty"`
	// example:
	//
	// v4
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// {}
	ResourcesChanged *string `json:"resourcesChanged,omitempty" xml:"resourcesChanged,omitempty"`
	SubCommand       *string `json:"subCommand,omitempty" xml:"subCommand,omitempty"`
}

func (s ListJobsResponseBodyJobsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyJobsConfig) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsConfig) GetIsDestroy() *bool {
	return s.IsDestroy
}

func (s *ListJobsResponseBodyJobsConfig) GetModuleDescription() *string {
	return s.ModuleDescription
}

func (s *ListJobsResponseBodyJobsConfig) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *ListJobsResponseBodyJobsConfig) GetResourcesChanged() *string {
	return s.ResourcesChanged
}

func (s *ListJobsResponseBodyJobsConfig) GetSubCommand() *string {
	return s.SubCommand
}

func (s *ListJobsResponseBodyJobsConfig) SetIsDestroy(v bool) *ListJobsResponseBodyJobsConfig {
	s.IsDestroy = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetModuleDescription(v string) *ListJobsResponseBodyJobsConfig {
	s.ModuleDescription = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetModuleVersion(v string) *ListJobsResponseBodyJobsConfig {
	s.ModuleVersion = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetResourcesChanged(v string) *ListJobsResponseBodyJobsConfig {
	s.ResourcesChanged = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetSubCommand(v string) *ListJobsResponseBodyJobsConfig {
	s.SubCommand = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) Validate() error {
	return dara.Validate(s)
}
