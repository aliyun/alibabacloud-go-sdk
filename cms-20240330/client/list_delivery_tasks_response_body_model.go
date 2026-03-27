// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryTasks(v []*ListDeliveryTasksResponseBodyDeliveryTasks) *ListDeliveryTasksResponseBody
	GetDeliveryTasks() []*ListDeliveryTasksResponseBodyDeliveryTasks
	SetMaxResults(v int32) *ListDeliveryTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDeliveryTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDeliveryTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDeliveryTasksResponseBody
	GetTotalCount() *int32
}

type ListDeliveryTasksResponseBody struct {
	DeliveryTasks []*ListDeliveryTasksResponseBodyDeliveryTasks `json:"deliveryTasks,omitempty" xml:"deliveryTasks,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 2-ba4d-4b9f-aa24-dcb067a30f1c
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 7D7DF334-B2F2-5453-AD51-A27B337E3191
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDeliveryTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksResponseBody) GetDeliveryTasks() []*ListDeliveryTasksResponseBodyDeliveryTasks {
	return s.DeliveryTasks
}

func (s *ListDeliveryTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDeliveryTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeliveryTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeliveryTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDeliveryTasksResponseBody) SetDeliveryTasks(v []*ListDeliveryTasksResponseBodyDeliveryTasks) *ListDeliveryTasksResponseBody {
	s.DeliveryTasks = v
	return s
}

func (s *ListDeliveryTasksResponseBody) SetMaxResults(v int32) *ListDeliveryTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDeliveryTasksResponseBody) SetNextToken(v string) *ListDeliveryTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDeliveryTasksResponseBody) SetRequestId(v string) *ListDeliveryTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeliveryTasksResponseBody) SetTotalCount(v int32) *ListDeliveryTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDeliveryTasksResponseBody) Validate() error {
	if s.DeliveryTasks != nil {
		for _, item := range s.DeliveryTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDeliveryTasksResponseBodyDeliveryTasks struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-22T09:02:01Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// rw-5f2b4c7e66342s
	DataSourceId   *string                                              `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	ExternalLabels map[string]*string                                   `json:"externalLabels,omitempty" xml:"externalLabels,omitempty"`
	ExtraInfo      *ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	LabelFilters   map[string]*string                                   `json:"labelFilters,omitempty" xml:"labelFilters,omitempty"`
	// example:
	//
	// Allow
	LabelFiltersType *string `json:"labelFiltersType,omitempty" xml:"labelFiltersType,omitempty"`
	// example:
	//
	// rg-acfm3gn5i6bigbi
	ResourceGroupId *string                                               `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	SinkList        []*ListDeliveryTasksResponseBodyDeliveryTasksSinkList `json:"sinkList,omitempty" xml:"sinkList,omitempty" type:"Repeated"`
	// example:
	//
	// Enable
	Status *string                                           `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListDeliveryTasksResponseBodyDeliveryTasksTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// my delivery task
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// db21f8a126d96953
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// test-task
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-22T09:02:01Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListDeliveryTasksResponseBodyDeliveryTasks) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksResponseBodyDeliveryTasks) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetExternalLabels() map[string]*string {
	return s.ExternalLabels
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetExtraInfo() *ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo {
	return s.ExtraInfo
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetLabelFilters() map[string]*string {
	return s.LabelFilters
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetLabelFiltersType() *string {
	return s.LabelFiltersType
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetSinkList() []*ListDeliveryTasksResponseBodyDeliveryTasksSinkList {
	return s.SinkList
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetStatus() *string {
	return s.Status
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetTags() []*ListDeliveryTasksResponseBodyDeliveryTasksTags {
	return s.Tags
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetCreateTime(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.CreateTime = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetDataSourceId(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.DataSourceId = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetExternalLabels(v map[string]*string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.ExternalLabels = v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetExtraInfo(v *ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.ExtraInfo = v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetLabelFilters(v map[string]*string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.LabelFilters = v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetLabelFiltersType(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.LabelFiltersType = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetResourceGroupId(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetSinkList(v []*ListDeliveryTasksResponseBodyDeliveryTasksSinkList) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.SinkList = v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetStatus(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.Status = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetTags(v []*ListDeliveryTasksResponseBodyDeliveryTasksTags) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.Tags = v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetTaskDescription(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.TaskDescription = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetTaskId(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.TaskId = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetTaskName(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.TaskName = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) SetUpdateTime(v string) *ListDeliveryTasksResponseBodyDeliveryTasks {
	s.UpdateTime = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasks) Validate() error {
	if s.ExtraInfo != nil {
		if err := s.ExtraInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SinkList != nil {
		for _, item := range s.SinkList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo struct {
	TaskNameList []*string `json:"taskNameList,omitempty" xml:"taskNameList,omitempty" type:"Repeated"`
}

func (s ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo) GetTaskNameList() []*string {
	return s.TaskNameList
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo) SetTaskNameList(v []*string) *ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo {
	s.TaskNameList = v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksExtraInfo) Validate() error {
	return dara.Validate(s)
}

type ListDeliveryTasksResponseBodyDeliveryTasksSinkList struct {
	SinkConfigs map[string]*string `json:"sinkConfigs,omitempty" xml:"sinkConfigs,omitempty"`
	// example:
	//
	// Prometheus
	SinkType *string `json:"sinkType,omitempty" xml:"sinkType,omitempty"`
}

func (s ListDeliveryTasksResponseBodyDeliveryTasksSinkList) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksResponseBodyDeliveryTasksSinkList) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksSinkList) GetSinkConfigs() map[string]*string {
	return s.SinkConfigs
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksSinkList) GetSinkType() *string {
	return s.SinkType
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksSinkList) SetSinkConfigs(v map[string]*string) *ListDeliveryTasksResponseBodyDeliveryTasksSinkList {
	s.SinkConfigs = v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksSinkList) SetSinkType(v string) *ListDeliveryTasksResponseBodyDeliveryTasksSinkList {
	s.SinkType = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksSinkList) Validate() error {
	return dara.Validate(s)
}

type ListDeliveryTasksResponseBodyDeliveryTasksTags struct {
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListDeliveryTasksResponseBodyDeliveryTasksTags) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryTasksResponseBodyDeliveryTasksTags) GoString() string {
	return s.String()
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksTags) GetKey() *string {
	return s.Key
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksTags) GetValue() *string {
	return s.Value
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksTags) SetKey(v string) *ListDeliveryTasksResponseBodyDeliveryTasksTags {
	s.Key = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksTags) SetValue(v string) *ListDeliveryTasksResponseBodyDeliveryTasksTags {
	s.Value = &v
	return s
}

func (s *ListDeliveryTasksResponseBodyDeliveryTasksTags) Validate() error {
	return dara.Validate(s)
}
