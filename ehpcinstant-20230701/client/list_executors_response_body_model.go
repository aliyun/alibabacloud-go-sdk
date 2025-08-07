// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutors(v []*ListExecutorsResponseBodyExecutors) *ListExecutorsResponseBody
	GetExecutors() []*ListExecutorsResponseBodyExecutors
	SetPageNumber(v int32) *ListExecutorsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExecutorsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListExecutorsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListExecutorsResponseBody
	GetTotalCount() *string
}

type ListExecutorsResponseBody struct {
	Executors []*ListExecutorsResponseBodyExecutors `json:"Executors,omitempty" xml:"Executors,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 40
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExecutorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBody) GetExecutors() []*ListExecutorsResponseBodyExecutors {
	return s.Executors
}

func (s *ListExecutorsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExecutorsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutorsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListExecutorsResponseBody) SetExecutors(v []*ListExecutorsResponseBodyExecutors) *ListExecutorsResponseBody {
	s.Executors = v
	return s
}

func (s *ListExecutorsResponseBody) SetPageNumber(v int32) *ListExecutorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsResponseBody) SetPageSize(v int32) *ListExecutorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListExecutorsResponseBody) SetRequestId(v string) *ListExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorsResponseBody) SetTotalCount(v string) *ListExecutorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExecutorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExecutorsResponseBodyExecutors struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 0
	ArrayIndex    *int32 `json:"ArrayIndex,omitempty" xml:"ArrayIndex,omitempty"`
	BlockDuration *int32 `json:"BlockDuration,omitempty" xml:"BlockDuration,omitempty"`
	// example:
	//
	// 2024-02-20 10:04:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-02-20 10:04:18
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// job-xxxx-task0-1
	ExecutorId        *string   `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	ExpirationTime    *string   `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	ExternalIpAddress []*string `json:"ExternalIpAddress,omitempty" xml:"ExternalIpAddress,omitempty" type:"Repeated"`
	HostName          []*string `json:"HostName,omitempty" xml:"HostName,omitempty" type:"Repeated"`
	Image             *string   `json:"Image,omitempty" xml:"Image,omitempty"`
	IpAddress         []*string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty" type:"Repeated"`
	// example:
	//
	// job-hy1nggvyukuvkr******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// testJob
	JobName      *string                                     `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Preemptible  *bool                                       `json:"Preemptible,omitempty" xml:"Preemptible,omitempty"`
	Resource     *ListExecutorsResponseBodyExecutorsResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	ResourceType *string                                     `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime    *string                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Succeeded to release executor resource
	StatusReason *string                                   `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	Tags         []*ListExecutorsResponseBodyExecutorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// task0
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskSustainable *bool   `json:"TaskSustainable,omitempty" xml:"TaskSustainable,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListExecutorsResponseBodyExecutors) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutors) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutors) GetAppName() *string {
	return s.AppName
}

func (s *ListExecutorsResponseBodyExecutors) GetArrayIndex() *int32 {
	return s.ArrayIndex
}

func (s *ListExecutorsResponseBodyExecutors) GetBlockDuration() *int32 {
	return s.BlockDuration
}

func (s *ListExecutorsResponseBodyExecutors) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListExecutorsResponseBodyExecutors) GetEndTime() *string {
	return s.EndTime
}

func (s *ListExecutorsResponseBodyExecutors) GetExecutorId() *string {
	return s.ExecutorId
}

func (s *ListExecutorsResponseBodyExecutors) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *ListExecutorsResponseBodyExecutors) GetExternalIpAddress() []*string {
	return s.ExternalIpAddress
}

func (s *ListExecutorsResponseBodyExecutors) GetHostName() []*string {
	return s.HostName
}

func (s *ListExecutorsResponseBodyExecutors) GetImage() *string {
	return s.Image
}

func (s *ListExecutorsResponseBodyExecutors) GetIpAddress() []*string {
	return s.IpAddress
}

func (s *ListExecutorsResponseBodyExecutors) GetJobId() *string {
	return s.JobId
}

func (s *ListExecutorsResponseBodyExecutors) GetJobName() *string {
	return s.JobName
}

func (s *ListExecutorsResponseBodyExecutors) GetPreemptible() *bool {
	return s.Preemptible
}

func (s *ListExecutorsResponseBodyExecutors) GetResource() *ListExecutorsResponseBodyExecutorsResource {
	return s.Resource
}

func (s *ListExecutorsResponseBodyExecutors) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListExecutorsResponseBodyExecutors) GetStartTime() *string {
	return s.StartTime
}

func (s *ListExecutorsResponseBodyExecutors) GetStatus() *string {
	return s.Status
}

func (s *ListExecutorsResponseBodyExecutors) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListExecutorsResponseBodyExecutors) GetTags() []*ListExecutorsResponseBodyExecutorsTags {
	return s.Tags
}

func (s *ListExecutorsResponseBodyExecutors) GetTaskName() *string {
	return s.TaskName
}

func (s *ListExecutorsResponseBodyExecutors) GetTaskSustainable() *bool {
	return s.TaskSustainable
}

func (s *ListExecutorsResponseBodyExecutors) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListExecutorsResponseBodyExecutors) SetAppName(v string) *ListExecutorsResponseBodyExecutors {
	s.AppName = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetArrayIndex(v int32) *ListExecutorsResponseBodyExecutors {
	s.ArrayIndex = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetBlockDuration(v int32) *ListExecutorsResponseBodyExecutors {
	s.BlockDuration = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetCreateTime(v string) *ListExecutorsResponseBodyExecutors {
	s.CreateTime = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetEndTime(v string) *ListExecutorsResponseBodyExecutors {
	s.EndTime = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetExecutorId(v string) *ListExecutorsResponseBodyExecutors {
	s.ExecutorId = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetExpirationTime(v string) *ListExecutorsResponseBodyExecutors {
	s.ExpirationTime = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetExternalIpAddress(v []*string) *ListExecutorsResponseBodyExecutors {
	s.ExternalIpAddress = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetHostName(v []*string) *ListExecutorsResponseBodyExecutors {
	s.HostName = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetImage(v string) *ListExecutorsResponseBodyExecutors {
	s.Image = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetIpAddress(v []*string) *ListExecutorsResponseBodyExecutors {
	s.IpAddress = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetJobId(v string) *ListExecutorsResponseBodyExecutors {
	s.JobId = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetJobName(v string) *ListExecutorsResponseBodyExecutors {
	s.JobName = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetPreemptible(v bool) *ListExecutorsResponseBodyExecutors {
	s.Preemptible = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetResource(v *ListExecutorsResponseBodyExecutorsResource) *ListExecutorsResponseBodyExecutors {
	s.Resource = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetResourceType(v string) *ListExecutorsResponseBodyExecutors {
	s.ResourceType = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetStartTime(v string) *ListExecutorsResponseBodyExecutors {
	s.StartTime = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetStatus(v string) *ListExecutorsResponseBodyExecutors {
	s.Status = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetStatusReason(v string) *ListExecutorsResponseBodyExecutors {
	s.StatusReason = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetTags(v []*ListExecutorsResponseBodyExecutorsTags) *ListExecutorsResponseBodyExecutors {
	s.Tags = v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetTaskName(v string) *ListExecutorsResponseBodyExecutors {
	s.TaskName = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetTaskSustainable(v bool) *ListExecutorsResponseBodyExecutors {
	s.TaskSustainable = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) SetVswitchId(v string) *ListExecutorsResponseBodyExecutors {
	s.VswitchId = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutors) Validate() error {
	return dara.Validate(s)
}

type ListExecutorsResponseBodyExecutorsResource struct {
	Cores        *float32                                           `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Disks        []*ListExecutorsResponseBodyExecutorsResourceDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	InstanceType *string                                            `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Memory       *float32                                           `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListExecutorsResponseBodyExecutorsResource) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutorsResource) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutorsResource) GetCores() *float32 {
	return s.Cores
}

func (s *ListExecutorsResponseBodyExecutorsResource) GetDisks() []*ListExecutorsResponseBodyExecutorsResourceDisks {
	return s.Disks
}

func (s *ListExecutorsResponseBodyExecutorsResource) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListExecutorsResponseBodyExecutorsResource) GetMemory() *float32 {
	return s.Memory
}

func (s *ListExecutorsResponseBodyExecutorsResource) SetCores(v float32) *ListExecutorsResponseBodyExecutorsResource {
	s.Cores = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResource) SetDisks(v []*ListExecutorsResponseBodyExecutorsResourceDisks) *ListExecutorsResponseBodyExecutorsResource {
	s.Disks = v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResource) SetInstanceType(v string) *ListExecutorsResponseBodyExecutorsResource {
	s.InstanceType = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResource) SetMemory(v float32) *ListExecutorsResponseBodyExecutorsResource {
	s.Memory = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResource) Validate() error {
	return dara.Validate(s)
}

type ListExecutorsResponseBodyExecutorsResourceDisks struct {
	Size *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListExecutorsResponseBodyExecutorsResourceDisks) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutorsResourceDisks) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutorsResourceDisks) GetSize() *int32 {
	return s.Size
}

func (s *ListExecutorsResponseBodyExecutorsResourceDisks) GetType() *string {
	return s.Type
}

func (s *ListExecutorsResponseBodyExecutorsResourceDisks) SetSize(v int32) *ListExecutorsResponseBodyExecutorsResourceDisks {
	s.Size = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResourceDisks) SetType(v string) *ListExecutorsResponseBodyExecutorsResourceDisks {
	s.Type = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsResourceDisks) Validate() error {
	return dara.Validate(s)
}

type ListExecutorsResponseBodyExecutorsTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListExecutorsResponseBodyExecutorsTags) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsResponseBodyExecutorsTags) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyExecutorsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListExecutorsResponseBodyExecutorsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListExecutorsResponseBodyExecutorsTags) SetTagKey(v string) *ListExecutorsResponseBodyExecutorsTags {
	s.TagKey = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsTags) SetTagValue(v string) *ListExecutorsResponseBodyExecutorsTags {
	s.TagValue = &v
	return s
}

func (s *ListExecutorsResponseBodyExecutorsTags) Validate() error {
	return dara.Validate(s)
}
