// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrecisionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPrecisionTaskResponseBody
	GetCode() *string
	SetCount(v int32) *ListPrecisionTaskResponseBody
	GetCount() *int32
	SetData(v *ListPrecisionTaskResponseBodyData) *ListPrecisionTaskResponseBody
	GetData() *ListPrecisionTaskResponseBodyData
	SetMessage(v string) *ListPrecisionTaskResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListPrecisionTaskResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPrecisionTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListPrecisionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPrecisionTaskResponseBody
	GetSuccess() *bool
}

type ListPrecisionTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Count *int32                             `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  *ListPrecisionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPrecisionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPrecisionTaskResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListPrecisionTaskResponseBody) GetData() *ListPrecisionTaskResponseBodyData {
	return s.Data
}

func (s *ListPrecisionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrecisionTaskResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPrecisionTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPrecisionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrecisionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPrecisionTaskResponseBody) SetCode(v string) *ListPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetCount(v int32) *ListPrecisionTaskResponseBody {
	s.Count = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetData(v *ListPrecisionTaskResponseBodyData) *ListPrecisionTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetMessage(v string) *ListPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetPageNumber(v int32) *ListPrecisionTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetPageSize(v int32) *ListPrecisionTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetRequestId(v string) *ListPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) SetSuccess(v bool) *ListPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ListPrecisionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPrecisionTaskResponseBodyData struct {
	PrecisionTask []*ListPrecisionTaskResponseBodyDataPrecisionTask `json:"PrecisionTask,omitempty" xml:"PrecisionTask,omitempty" type:"Repeated"`
}

func (s ListPrecisionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyData) GetPrecisionTask() []*ListPrecisionTaskResponseBodyDataPrecisionTask {
	return s.PrecisionTask
}

func (s *ListPrecisionTaskResponseBodyData) SetPrecisionTask(v []*ListPrecisionTaskResponseBodyDataPrecisionTask) *ListPrecisionTaskResponseBodyData {
	s.PrecisionTask = v
	return s
}

func (s *ListPrecisionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListPrecisionTaskResponseBodyDataPrecisionTask struct {
	// example:
	//
	// 2020-03-10 20:26:29
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1212
	DataSetId   *int64  `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// example:
	//
	// 331311
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 32
	IncorrectWords *int32                                                    `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Name           *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Precisions     *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions `json:"Precisions,omitempty" xml:"Precisions,omitempty" type:"Struct"`
	// example:
	//
	// 3
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 7C1DEF5F-2C18-4D36-99C6-8C276F781796
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 21
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2020-03-10 20:26:29
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 3
	VerifiedCount *int32 `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTask) String() string {
	return dara.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTask) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetDataSetId() *int64 {
	return s.DataSetId
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetDataSetName() *string {
	return s.DataSetName
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetDuration() *int32 {
	return s.Duration
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetIncorrectWords() *int32 {
	return s.IncorrectWords
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetName() *string {
	return s.Name
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetPrecisions() *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions {
	return s.Precisions
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetSource() *int32 {
	return s.Source
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetStatus() *int32 {
	return s.Status
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetTaskId() *string {
	return s.TaskId
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) GetVerifiedCount() *int32 {
	return s.VerifiedCount
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetCreateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.CreateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDataSetId(v int64) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.DataSetId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDataSetName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.DataSetName = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetDuration(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Duration = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetIncorrectWords(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.IncorrectWords = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Name = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetPrecisions(v *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Precisions = v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetSource(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Source = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetStatus(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.Status = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetTaskId(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.TaskId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetTotalCount(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.TotalCount = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetUpdateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.UpdateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) SetVerifiedCount(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTask {
	s.VerifiedCount = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTask) Validate() error {
	return dara.Validate(s)
}

type ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions struct {
	Precision []*ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision `json:"Precision,omitempty" xml:"Precision,omitempty" type:"Repeated"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) String() string {
	return dara.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) GetPrecision() []*ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	return s.Precision
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) SetPrecision(v []*ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions {
	s.Precision = v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisions) Validate() error {
	return dara.Validate(s)
}

type ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision struct {
	// example:
	//
	// 2020-03-10 20:26:29
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2321
	ModelId   *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// 0.98
	Precision *float32 `json:"Precision,omitempty" xml:"Precision,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 7C1DEF5F-2C18-4D36-99C6-8C276F781796
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) String() string {
	return dara.Prettify(s)
}

func (s ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GoString() string {
	return s.String()
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GetModelId() *int64 {
	return s.ModelId
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GetModelName() *string {
	return s.ModelName
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GetPrecision() *float32 {
	return s.Precision
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GetStatus() *int32 {
	return s.Status
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) GetTaskId() *string {
	return s.TaskId
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetCreateTime(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.CreateTime = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetModelId(v int64) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.ModelId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetModelName(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.ModelName = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetPrecision(v float32) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.Precision = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetStatus(v int32) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.Status = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) SetTaskId(v string) *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision {
	s.TaskId = &v
	return s
}

func (s *ListPrecisionTaskResponseBodyDataPrecisionTaskPrecisionsPrecision) Validate() error {
	return dara.Validate(s)
}
