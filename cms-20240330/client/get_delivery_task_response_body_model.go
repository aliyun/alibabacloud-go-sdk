// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryTask(v *GetDeliveryTaskResponseBodyDeliveryTask) *GetDeliveryTaskResponseBody
	GetDeliveryTask() *GetDeliveryTaskResponseBodyDeliveryTask
	SetRequestId(v string) *GetDeliveryTaskResponseBody
	GetRequestId() *string
}

type GetDeliveryTaskResponseBody struct {
	DeliveryTask *GetDeliveryTaskResponseBodyDeliveryTask `json:"deliveryTask,omitempty" xml:"deliveryTask,omitempty" type:"Struct"`
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeliveryTaskResponseBody) GetDeliveryTask() *GetDeliveryTaskResponseBodyDeliveryTask {
	return s.DeliveryTask
}

func (s *GetDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeliveryTaskResponseBody) SetDeliveryTask(v *GetDeliveryTaskResponseBodyDeliveryTask) *GetDeliveryTaskResponseBody {
	s.DeliveryTask = v
	return s
}

func (s *GetDeliveryTaskResponseBody) SetRequestId(v string) *GetDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeliveryTaskResponseBody) Validate() error {
	if s.DeliveryTask != nil {
		if err := s.DeliveryTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeliveryTaskResponseBodyDeliveryTask struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2025-07-24T02:08:27Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// rw-xxxxxx
	DataSourceId   *string                                           `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	ExternalLabels map[string]*string                                `json:"externalLabels,omitempty" xml:"externalLabels,omitempty"`
	ExtraInfo      *GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo `json:"extraInfo,omitempty" xml:"extraInfo,omitempty" type:"Struct"`
	LabelFilters   map[string]*string                                `json:"labelFilters,omitempty" xml:"labelFilters,omitempty"`
	// example:
	//
	// Allow
	LabelFiltersType *string `json:"labelFiltersType,omitempty" xml:"labelFiltersType,omitempty"`
	// example:
	//
	// cn-chengdu
	RegionId *string                                            `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SinkList []*GetDeliveryTaskResponseBodyDeliveryTaskSinkList `json:"sinkList,omitempty" xml:"sinkList,omitempty" type:"Repeated"`
	// example:
	//
	// Pending2Running
	Status *string                                        `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*GetDeliveryTaskResponseBodyDeliveryTaskTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// my delivery task
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// example:
	//
	// 8b07eeac8249866d
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// test-task
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-22T11:50:48Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetDeliveryTaskResponseBodyDeliveryTask) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryTaskResponseBodyDeliveryTask) GoString() string {
	return s.String()
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetExternalLabels() map[string]*string {
	return s.ExternalLabels
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetExtraInfo() *GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo {
	return s.ExtraInfo
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetLabelFilters() map[string]*string {
	return s.LabelFilters
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetLabelFiltersType() *string {
	return s.LabelFiltersType
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetSinkList() []*GetDeliveryTaskResponseBodyDeliveryTaskSinkList {
	return s.SinkList
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetStatus() *string {
	return s.Status
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetTags() []*GetDeliveryTaskResponseBodyDeliveryTaskTags {
	return s.Tags
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetTaskName() *string {
	return s.TaskName
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetCreateTime(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.CreateTime = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetDataSourceId(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.DataSourceId = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetExternalLabels(v map[string]*string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.ExternalLabels = v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetExtraInfo(v *GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.ExtraInfo = v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetLabelFilters(v map[string]*string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.LabelFilters = v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetLabelFiltersType(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.LabelFiltersType = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetRegionId(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.RegionId = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetSinkList(v []*GetDeliveryTaskResponseBodyDeliveryTaskSinkList) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.SinkList = v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetStatus(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.Status = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetTags(v []*GetDeliveryTaskResponseBodyDeliveryTaskTags) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.Tags = v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetTaskDescription(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.TaskDescription = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetTaskId(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.TaskId = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetTaskName(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.TaskName = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) SetUpdateTime(v string) *GetDeliveryTaskResponseBodyDeliveryTask {
	s.UpdateTime = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTask) Validate() error {
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

type GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo struct {
	TaskNameList []*string `json:"taskNameList,omitempty" xml:"taskNameList,omitempty" type:"Repeated"`
}

func (s GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo) GoString() string {
	return s.String()
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo) GetTaskNameList() []*string {
	return s.TaskNameList
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo) SetTaskNameList(v []*string) *GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo {
	s.TaskNameList = v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskExtraInfo) Validate() error {
	return dara.Validate(s)
}

type GetDeliveryTaskResponseBodyDeliveryTaskSinkList struct {
	SinkConfigs map[string]*string `json:"sinkConfigs,omitempty" xml:"sinkConfigs,omitempty"`
	// example:
	//
	// Prometheus
	SinkType *string `json:"sinkType,omitempty" xml:"sinkType,omitempty"`
}

func (s GetDeliveryTaskResponseBodyDeliveryTaskSinkList) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryTaskResponseBodyDeliveryTaskSinkList) GoString() string {
	return s.String()
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskSinkList) GetSinkConfigs() map[string]*string {
	return s.SinkConfigs
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskSinkList) GetSinkType() *string {
	return s.SinkType
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskSinkList) SetSinkConfigs(v map[string]*string) *GetDeliveryTaskResponseBodyDeliveryTaskSinkList {
	s.SinkConfigs = v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskSinkList) SetSinkType(v string) *GetDeliveryTaskResponseBodyDeliveryTaskSinkList {
	s.SinkType = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskSinkList) Validate() error {
	return dara.Validate(s)
}

type GetDeliveryTaskResponseBodyDeliveryTaskTags struct {
	// example:
	//
	// sourcetype
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// production
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetDeliveryTaskResponseBodyDeliveryTaskTags) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryTaskResponseBodyDeliveryTaskTags) GoString() string {
	return s.String()
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskTags) GetKey() *string {
	return s.Key
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskTags) GetValue() *string {
	return s.Value
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskTags) SetKey(v string) *GetDeliveryTaskResponseBodyDeliveryTaskTags {
	s.Key = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskTags) SetValue(v string) *GetDeliveryTaskResponseBodyDeliveryTaskTags {
	s.Value = &v
	return s
}

func (s *GetDeliveryTaskResponseBodyDeliveryTaskTags) Validate() error {
	return dara.Validate(s)
}
