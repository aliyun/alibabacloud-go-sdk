// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListJobsResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *ListJobsResponseBody
	GetCurrentPage() *int32
	SetData(v *ListJobsResponseBodyData) *ListJobsResponseBody
	GetData() *ListJobsResponseBodyData
	SetErrorCode(v string) *ListJobsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListJobsResponseBody
	GetMessage() *string
	SetPageSize(v int32) *ListJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobsResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListJobsResponseBody
	GetTotalSize() *int32
}

type ListJobsResponseBody struct {
	// The HTTP status code or POP error code. Valid values:
	//
	// - **2xx**: success
	//
	// - **3xx**: redirection
	//
	// - **4xx**: client error
	//
	// - **5xx**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of job templates.
	Data *ListJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Take note of the following items:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section of this topic.
	//
	// example:
	//
	// InvalidAppId.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the operation result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B4D805CA-926D-41B1-8E63-7AD0C1ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the job template list was queried successfully. Valid values:
	//
	// - **true**: The query was successful.
	//
	// - **false**: The query failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of job templates.
	//
	// example:
	//
	// 2
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListJobsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListJobsResponseBody) GetData() *ListJobsResponseBodyData {
	return s.Data
}

func (s *ListJobsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListJobsResponseBody) SetCode(v string) *ListJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobsResponseBody) SetCurrentPage(v int32) *ListJobsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListJobsResponseBody) SetData(v *ListJobsResponseBodyData) *ListJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListJobsResponseBody) SetErrorCode(v string) *ListJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobsResponseBody) SetMessage(v string) *ListJobsResponseBody {
	s.Message = &v
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

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalSize(v int32) *ListJobsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobsResponseBodyData struct {
	// The list of job templates.
	Applications []*ListJobsResponseBodyDataApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of job templates.
	//
	// example:
	//
	// 2
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyData) GetApplications() []*ListJobsResponseBodyDataApplications {
	return s.Applications
}

func (s *ListJobsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListJobsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListJobsResponseBodyData) SetApplications(v []*ListJobsResponseBodyDataApplications) *ListJobsResponseBodyData {
	s.Applications = v
	return s
}

func (s *ListJobsResponseBodyData) SetCurrentPage(v int32) *ListJobsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListJobsResponseBodyData) SetPageSize(v int32) *ListJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBodyData) SetTotalSize(v int32) *ListJobsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListJobsResponseBodyData) Validate() error {
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

type ListJobsResponseBodyDataApplications struct {
	// The number of running instances.
	//
	// example:
	//
	// 0
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the job template.
	//
	// example:
	//
	// description
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The ID of the job template.
	//
	// example:
	//
	// f7730764-d88f-4b9a-8d8e-cd8efbfe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the job template.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The completion time of the last job.
	//
	// example:
	//
	// 1657522839
	CompletionTime *int64 `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	// The CPU required for each instance. Unit: millicores. The value cannot be 0. Currently, only the following defined specifications are supported:
	//
	// - **500**
	//
	// - **1000**
	//
	// - **2000**
	//
	// - **4000**
	//
	// - **8000**
	//
	// - **16000**
	//
	// - **32000**
	//
	// example:
	//
	// 500
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of failed instances.
	//
	// example:
	//
	// 0
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The image URL.
	//
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/sae-dev-test/nginx:stable
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether the last change order was executed successfully. Valid values:
	//
	// - **0**: The change order failed.
	//
	// - **1**: The change order succeeded.
	//
	// example:
	//
	// 1
	LastChangeorderState *string `json:"LastChangeorderState,omitempty" xml:"LastChangeorderState,omitempty"`
	// The running status of the last job. Valid values:
	//
	// - **0**: not executed
	//
	// - **1**: executed successfully
	//
	// - **2**: execution failed
	//
	// - **3**: executing
	//
	// example:
	//
	// 0
	LastJobState *string `json:"LastJobState,omitempty" xml:"LastJobState,omitempty"`
	// The start time of the last job.
	//
	// example:
	//
	// 1657522800
	LastStartTime *int64 `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	// The memory required for each instance. Unit: MB. The value cannot be 0. The memory has a one-to-one mapping with the CPU. Currently, only the following defined specifications are supported:
	//
	// - **1024**: corresponds to 500 and 1000 millicores
	//
	// - **2048**: corresponds to 500, 1000, and 2000 millicores
	//
	// - **4096**: corresponds to 1000, 2000, and 4000 millicores
	//
	// - **8192**: corresponds to 2000, 4000, and 8000 millicores
	//
	// - **12288**: corresponds to 12000 millicores
	//
	// - **16384**: corresponds to 4000, 8000, and 16000 millicores
	//
	// - **24576**: corresponds to 12000 millicores
	//
	// - **32768**: corresponds to 16000 millicores
	//
	// - **65536**: corresponds to 8000, 16000, and 32000 millicores
	//
	// - **131072**: corresponds to 32000 millicores
	//
	// example:
	//
	// 1024
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The additional information about the operation result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:demo
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of successful instances.
	//
	// example:
	//
	// 3
	Succeeded *int64 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
	// Specifies whether the job template is paused.
	//
	// example:
	//
	// false
	Suspend *bool `json:"Suspend,omitempty" xml:"Suspend,omitempty"`
	// The tags of the job template.
	Tags          []*ListJobsResponseBodyDataApplicationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TriggerConfig *string                                     `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty"`
}

func (s ListJobsResponseBodyDataApplications) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataApplications) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataApplications) GetActive() *int64 {
	return s.Active
}

func (s *ListJobsResponseBodyDataApplications) GetAppDescription() *string {
	return s.AppDescription
}

func (s *ListJobsResponseBodyDataApplications) GetAppId() *string {
	return s.AppId
}

func (s *ListJobsResponseBodyDataApplications) GetAppName() *string {
	return s.AppName
}

func (s *ListJobsResponseBodyDataApplications) GetCompletionTime() *int64 {
	return s.CompletionTime
}

func (s *ListJobsResponseBodyDataApplications) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListJobsResponseBodyDataApplications) GetFailed() *int64 {
	return s.Failed
}

func (s *ListJobsResponseBodyDataApplications) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListJobsResponseBodyDataApplications) GetLastChangeorderState() *string {
	return s.LastChangeorderState
}

func (s *ListJobsResponseBodyDataApplications) GetLastJobState() *string {
	return s.LastJobState
}

func (s *ListJobsResponseBodyDataApplications) GetLastStartTime() *int64 {
	return s.LastStartTime
}

func (s *ListJobsResponseBodyDataApplications) GetMem() *int32 {
	return s.Mem
}

func (s *ListJobsResponseBodyDataApplications) GetMessage() *string {
	return s.Message
}

func (s *ListJobsResponseBodyDataApplications) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListJobsResponseBodyDataApplications) GetRegionId() *string {
	return s.RegionId
}

func (s *ListJobsResponseBodyDataApplications) GetSucceeded() *int64 {
	return s.Succeeded
}

func (s *ListJobsResponseBodyDataApplications) GetSuspend() *bool {
	return s.Suspend
}

func (s *ListJobsResponseBodyDataApplications) GetTags() []*ListJobsResponseBodyDataApplicationsTags {
	return s.Tags
}

func (s *ListJobsResponseBodyDataApplications) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *ListJobsResponseBodyDataApplications) SetActive(v int64) *ListJobsResponseBodyDataApplications {
	s.Active = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetAppDescription(v string) *ListJobsResponseBodyDataApplications {
	s.AppDescription = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetAppId(v string) *ListJobsResponseBodyDataApplications {
	s.AppId = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetAppName(v string) *ListJobsResponseBodyDataApplications {
	s.AppName = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetCompletionTime(v int64) *ListJobsResponseBodyDataApplications {
	s.CompletionTime = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetCpu(v int32) *ListJobsResponseBodyDataApplications {
	s.Cpu = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetFailed(v int64) *ListJobsResponseBodyDataApplications {
	s.Failed = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetImageUrl(v string) *ListJobsResponseBodyDataApplications {
	s.ImageUrl = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetLastChangeorderState(v string) *ListJobsResponseBodyDataApplications {
	s.LastChangeorderState = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetLastJobState(v string) *ListJobsResponseBodyDataApplications {
	s.LastJobState = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetLastStartTime(v int64) *ListJobsResponseBodyDataApplications {
	s.LastStartTime = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetMem(v int32) *ListJobsResponseBodyDataApplications {
	s.Mem = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetMessage(v string) *ListJobsResponseBodyDataApplications {
	s.Message = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetNamespaceId(v string) *ListJobsResponseBodyDataApplications {
	s.NamespaceId = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetRegionId(v string) *ListJobsResponseBodyDataApplications {
	s.RegionId = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetSucceeded(v int64) *ListJobsResponseBodyDataApplications {
	s.Succeeded = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetSuspend(v bool) *ListJobsResponseBodyDataApplications {
	s.Suspend = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetTags(v []*ListJobsResponseBodyDataApplicationsTags) *ListJobsResponseBodyDataApplications {
	s.Tags = v
	return s
}

func (s *ListJobsResponseBodyDataApplications) SetTriggerConfig(v string) *ListJobsResponseBodyDataApplications {
	s.TriggerConfig = &v
	return s
}

func (s *ListJobsResponseBodyDataApplications) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobsResponseBodyDataApplicationsTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListJobsResponseBodyDataApplicationsTags) String() string {
	return dara.Prettify(s)
}

func (s ListJobsResponseBodyDataApplicationsTags) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyDataApplicationsTags) GetKey() *string {
	return s.Key
}

func (s *ListJobsResponseBodyDataApplicationsTags) GetValue() *string {
	return s.Value
}

func (s *ListJobsResponseBodyDataApplicationsTags) SetKey(v string) *ListJobsResponseBodyDataApplicationsTags {
	s.Key = &v
	return s
}

func (s *ListJobsResponseBodyDataApplicationsTags) SetValue(v string) *ListJobsResponseBodyDataApplicationsTags {
	s.Value = &v
	return s
}

func (s *ListJobsResponseBodyDataApplicationsTags) Validate() error {
	return dara.Validate(s)
}
