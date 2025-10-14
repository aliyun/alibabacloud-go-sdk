// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRplInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeRplInspectionTaskResponseBodyData) *DescribeRplInspectionTaskResponseBody
	GetData() *DescribeRplInspectionTaskResponseBodyData
	SetMessage(v string) *DescribeRplInspectionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeRplInspectionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRplInspectionTaskResponseBody
	GetSuccess() *bool
}

type DescribeRplInspectionTaskResponseBody struct {
	Data *DescribeRplInspectionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRplInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRplInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRplInspectionTaskResponseBody) GetData() *DescribeRplInspectionTaskResponseBodyData {
	return s.Data
}

func (s *DescribeRplInspectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRplInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRplInspectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRplInspectionTaskResponseBody) SetData(v *DescribeRplInspectionTaskResponseBodyData) *DescribeRplInspectionTaskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRplInspectionTaskResponseBody) SetMessage(v string) *DescribeRplInspectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBody) SetRequestId(v string) *DescribeRplInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBody) SetSuccess(v bool) *DescribeRplInspectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRplInspectionTaskResponseBodyData struct {
	InspectionTaskList []*DescribeRplInspectionTaskResponseBodyDataInspectionTaskList `json:"InspectionTaskList,omitempty" xml:"InspectionTaskList,omitempty" type:"Repeated"`
	// example:
	//
	// PRE_CHECK
	SlinkStage *string `json:"SlinkStage,omitempty" xml:"SlinkStage,omitempty"`
}

func (s DescribeRplInspectionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRplInspectionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRplInspectionTaskResponseBodyData) GetInspectionTaskList() []*DescribeRplInspectionTaskResponseBodyDataInspectionTaskList {
	return s.InspectionTaskList
}

func (s *DescribeRplInspectionTaskResponseBodyData) GetSlinkStage() *string {
	return s.SlinkStage
}

func (s *DescribeRplInspectionTaskResponseBodyData) SetInspectionTaskList(v []*DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) *DescribeRplInspectionTaskResponseBodyData {
	s.InspectionTaskList = v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyData) SetSlinkStage(v string) *DescribeRplInspectionTaskResponseBodyData {
	s.SlinkStage = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyData) Validate() error {
	if s.InspectionTaskList != nil {
		for _, item := range s.InspectionTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRplInspectionTaskResponseBodyDataInspectionTaskList struct {
	// example:
	//
	// 2025-09-25T02:36:20.000+0000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// For database: wms ,Found SIMPLE sequences in 1.0. PolarDB-X 2.0 does not support SIMPLE sequence any more， please use show sequence to check them.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10142
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// slinktaskid。
	//
	// example:
	//
	// etx-hzrez23btmb6aq
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
	// example:
	//
	// CONNECTIVITY
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2025-09-23T03:25:21.000+0000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) GoString() string {
	return s.String()
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) GetDescription() *string {
	return s.Description
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) GetId() *int64 {
	return s.Id
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) GetStage() *string {
	return s.Stage
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) GetStatus() *string {
	return s.Status
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) SetCreateTime(v int64) *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList {
	s.CreateTime = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) SetDescription(v string) *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList {
	s.Description = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) SetId(v int64) *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList {
	s.Id = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) SetSlinkTaskId(v string) *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList {
	s.SlinkTaskId = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) SetStage(v string) *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList {
	s.Stage = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) SetStatus(v string) *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList {
	s.Status = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) SetUpdateTime(v int64) *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList {
	s.UpdateTime = &v
	return s
}

func (s *DescribeRplInspectionTaskResponseBodyDataInspectionTaskList) Validate() error {
	return dara.Validate(s)
}
