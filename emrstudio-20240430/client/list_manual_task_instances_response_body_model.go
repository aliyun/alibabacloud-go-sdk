// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManualTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListManualTaskInstancesResponseBodyData) *ListManualTaskInstancesResponseBody
	GetData() []*ListManualTaskInstancesResponseBodyData
	SetNextToken(v string) *ListManualTaskInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListManualTaskInstancesResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *ListManualTaskInstancesResponseBody
	GetTotalSize() *int32
}

type ListManualTaskInstancesResponseBody struct {
	Data []*ListManualTaskInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListManualTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListManualTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListManualTaskInstancesResponseBody) GetData() []*ListManualTaskInstancesResponseBodyData {
	return s.Data
}

func (s *ListManualTaskInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListManualTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListManualTaskInstancesResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListManualTaskInstancesResponseBody) SetData(v []*ListManualTaskInstancesResponseBodyData) *ListManualTaskInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListManualTaskInstancesResponseBody) SetNextToken(v string) *ListManualTaskInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListManualTaskInstancesResponseBody) SetRequestId(v string) *ListManualTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBody) SetTotalSize(v int32) *ListManualTaskInstancesResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListManualTaskInstancesResponseBody) Validate() error {
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

type ListManualTaskInstancesResponseBodyData struct {
	// example:
	//
	// c-b933c5aac7f7***
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// application_123_***
	ExternalAppId *string `json:"ExternalAppId,omitempty" xml:"ExternalAppId,omitempty"`
	// example:
	//
	// mti-3q9jo749ne5****
	ManualTaskInstanceId *string `json:"ManualTaskInstanceId,omitempty" xml:"ManualTaskInstanceId,omitempty"`
	// example:
	//
	// test
	ManualTaskInstanceName *string `json:"ManualTaskInstanceName,omitempty" xml:"ManualTaskInstanceName,omitempty"`
	// example:
	//
	// wg-3q9jo749ne5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListManualTaskInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListManualTaskInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListManualTaskInstancesResponseBodyData) GetEmrClusterId() *string {
	return s.EmrClusterId
}

func (s *ListManualTaskInstancesResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListManualTaskInstancesResponseBodyData) GetExternalAppId() *string {
	return s.ExternalAppId
}

func (s *ListManualTaskInstancesResponseBodyData) GetManualTaskInstanceId() *string {
	return s.ManualTaskInstanceId
}

func (s *ListManualTaskInstancesResponseBodyData) GetManualTaskInstanceName() *string {
	return s.ManualTaskInstanceName
}

func (s *ListManualTaskInstancesResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListManualTaskInstancesResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListManualTaskInstancesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListManualTaskInstancesResponseBodyData) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListManualTaskInstancesResponseBodyData) GetTaskParams() *string {
	return s.TaskParams
}

func (s *ListManualTaskInstancesResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *ListManualTaskInstancesResponseBodyData) SetEmrClusterId(v string) *ListManualTaskInstancesResponseBodyData {
	s.EmrClusterId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetEndTime(v string) *ListManualTaskInstancesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetExternalAppId(v string) *ListManualTaskInstancesResponseBodyData {
	s.ExternalAppId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetManualTaskInstanceId(v string) *ListManualTaskInstancesResponseBodyData {
	s.ManualTaskInstanceId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetManualTaskInstanceName(v string) *ListManualTaskInstancesResponseBodyData {
	s.ManualTaskInstanceName = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetResourceGroupId(v string) *ListManualTaskInstancesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetStartTime(v string) *ListManualTaskInstancesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetStatus(v string) *ListManualTaskInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetSubmitTime(v string) *ListManualTaskInstancesResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetTaskParams(v string) *ListManualTaskInstancesResponseBodyData {
	s.TaskParams = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) SetTaskType(v string) *ListManualTaskInstancesResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListManualTaskInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
