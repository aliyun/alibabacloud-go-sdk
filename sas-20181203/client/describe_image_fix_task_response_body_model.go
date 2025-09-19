// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageFixTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildTasks(v []*DescribeImageFixTaskResponseBodyBuildTasks) *DescribeImageFixTaskResponseBody
	GetBuildTasks() []*DescribeImageFixTaskResponseBodyBuildTasks
	SetPageInfo(v *DescribeImageFixTaskResponseBodyPageInfo) *DescribeImageFixTaskResponseBody
	GetPageInfo() *DescribeImageFixTaskResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageFixTaskResponseBody
	GetRequestId() *string
}

type DescribeImageFixTaskResponseBody struct {
	// The tasks returned.
	BuildTasks []*DescribeImageFixTaskResponseBodyBuildTasks `json:"BuildTasks,omitempty" xml:"BuildTasks,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageFixTaskResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8AC52BBA-85D3-5F64-9B48-D08437CAF916
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageFixTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFixTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageFixTaskResponseBody) GetBuildTasks() []*DescribeImageFixTaskResponseBodyBuildTasks {
	return s.BuildTasks
}

func (s *DescribeImageFixTaskResponseBody) GetPageInfo() *DescribeImageFixTaskResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageFixTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageFixTaskResponseBody) SetBuildTasks(v []*DescribeImageFixTaskResponseBodyBuildTasks) *DescribeImageFixTaskResponseBody {
	s.BuildTasks = v
	return s
}

func (s *DescribeImageFixTaskResponseBody) SetPageInfo(v *DescribeImageFixTaskResponseBodyPageInfo) *DescribeImageFixTaskResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageFixTaskResponseBody) SetRequestId(v string) *DescribeImageFixTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageFixTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageFixTaskResponseBodyBuildTasks struct {
	// The ID of the task.
	//
	// example:
	//
	// ivf-939536b5-c3ca-427b-8183-91007756
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The timestamp when the task starts. Unit: milliseconds.
	//
	// example:
	//
	// 2021-10-14 20:34:07
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The timestamp when the task ends. Unit: milliseconds.
	//
	// example:
	//
	// 2021-10-14 20:32:05
	FixTime *string `json:"FixTime,omitempty" xml:"FixTime,omitempty"`
	// The version of the image after image risks are fixed.
	//
	// example:
	//
	// redhat8-vault
	NewTag *string `json:"NewTag,omitempty" xml:"NewTag,omitempty"`
	// The UUID of the image after image risks are fixed.
	//
	// example:
	//
	// 2fa731681911ae8d1b5f11893ace****
	NewUuid *string `json:"NewUuid,omitempty" xml:"NewUuid,omitempty"`
	// The version of the image.
	//
	// example:
	//
	// centos8.1-ja
	OldTag *string `json:"OldTag,omitempty" xml:"OldTag,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// 2fa731681911ae8d1b5f11893ace****
	OldUuid *string `json:"OldUuid,omitempty" xml:"OldUuid,omitempty"`
	// The region of the image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test-redhat
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace of the image.
	//
	// example:
	//
	// name-002
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **1**: The task is running.
	//
	// 	- **2**: The task is successful.
	//
	// 	- **3**: The task failed.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the task. The value is fixed as IMAGE_REPAIR. The value indicates a task that fixes image risks.
	//
	// example:
	//
	// IMAGE_REPAIR
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The alias of the fixed vulnerability.
	//
	// example:
	//
	// CVE-2007-5686:rpath_linux Information Disclosure
	VulAlias *string `json:"VulAlias,omitempty" xml:"VulAlias,omitempty"`
}

func (s DescribeImageFixTaskResponseBodyBuildTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFixTaskResponseBodyBuildTasks) GoString() string {
	return s.String()
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetBuildTaskId() *string {
	return s.BuildTaskId
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetFixTime() *string {
	return s.FixTime
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetNewTag() *string {
	return s.NewTag
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetNewUuid() *string {
	return s.NewUuid
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetOldTag() *string {
	return s.OldTag
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetOldUuid() *string {
	return s.OldUuid
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetRepoName() *string {
	return s.RepoName
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) GetVulAlias() *string {
	return s.VulAlias
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetBuildTaskId(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.BuildTaskId = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetFinishTime(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.FinishTime = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetFixTime(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.FixTime = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetNewTag(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.NewTag = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetNewUuid(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.NewUuid = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetOldTag(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.OldTag = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetOldUuid(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.OldUuid = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetRegionId(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.RegionId = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetRepoName(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.RepoName = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetRepoNamespace(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetStatus(v int32) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.Status = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetTaskType(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) SetVulAlias(v string) *DescribeImageFixTaskResponseBodyBuildTasks {
	s.VulAlias = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyBuildTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeImageFixTaskResponseBodyPageInfo struct {
	// The number of tasks returned on the current page.
	//
	// example:
	//
	// 12
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page. Default value: **1**
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of tasks returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageFixTaskResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFixTaskResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) SetCount(v int32) *DescribeImageFixTaskResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageFixTaskResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageFixTaskResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageFixTaskResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageFixTaskResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
