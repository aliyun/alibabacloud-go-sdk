// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrecisionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPrecisionTaskResponseBody
	GetCode() *string
	SetData(v *GetPrecisionTaskResponseBodyData) *GetPrecisionTaskResponseBody
	GetData() *GetPrecisionTaskResponseBodyData
	SetMessage(v string) *GetPrecisionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPrecisionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPrecisionTaskResponseBody
	GetSuccess() *bool
}

type GetPrecisionTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPrecisionTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPrecisionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPrecisionTaskResponseBody) GetData() *GetPrecisionTaskResponseBodyData {
	return s.Data
}

func (s *GetPrecisionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPrecisionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPrecisionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPrecisionTaskResponseBody) SetCode(v string) *GetPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetData(v *GetPrecisionTaskResponseBodyData) *GetPrecisionTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetMessage(v string) *GetPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetRequestId(v string) *GetPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) SetSuccess(v bool) *GetPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetPrecisionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPrecisionTaskResponseBodyData struct {
	// example:
	//
	// 1212
	DataSetId   *int64  `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	DataSetName *string `json:"DataSetName,omitempty" xml:"DataSetName,omitempty"`
	// example:
	//
	// 3423
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 23
	IncorrectWords *int32                                      `json:"IncorrectWords,omitempty" xml:"IncorrectWords,omitempty"`
	Name           *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Precisions     *GetPrecisionTaskResponseBodyDataPrecisions `json:"Precisions,omitempty" xml:"Precisions,omitempty" type:"Struct"`
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
	// 7C1DEF5F-2C18-4D36-99C6-8C27*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2020-03-10 20:26:29
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 2
	VerifiedCount *int32 `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s GetPrecisionTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyData) GetDataSetId() *int64 {
	return s.DataSetId
}

func (s *GetPrecisionTaskResponseBodyData) GetDataSetName() *string {
	return s.DataSetName
}

func (s *GetPrecisionTaskResponseBodyData) GetDuration() *int32 {
	return s.Duration
}

func (s *GetPrecisionTaskResponseBodyData) GetIncorrectWords() *int32 {
	return s.IncorrectWords
}

func (s *GetPrecisionTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetPrecisionTaskResponseBodyData) GetPrecisions() *GetPrecisionTaskResponseBodyDataPrecisions {
	return s.Precisions
}

func (s *GetPrecisionTaskResponseBodyData) GetSource() *int32 {
	return s.Source
}

func (s *GetPrecisionTaskResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetPrecisionTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetPrecisionTaskResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetPrecisionTaskResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPrecisionTaskResponseBodyData) GetVerifiedCount() *int32 {
	return s.VerifiedCount
}

func (s *GetPrecisionTaskResponseBodyData) SetDataSetId(v int64) *GetPrecisionTaskResponseBodyData {
	s.DataSetId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetDataSetName(v string) *GetPrecisionTaskResponseBodyData {
	s.DataSetName = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetDuration(v int32) *GetPrecisionTaskResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetIncorrectWords(v int32) *GetPrecisionTaskResponseBodyData {
	s.IncorrectWords = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetName(v string) *GetPrecisionTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetPrecisions(v *GetPrecisionTaskResponseBodyDataPrecisions) *GetPrecisionTaskResponseBodyData {
	s.Precisions = v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetSource(v int32) *GetPrecisionTaskResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetStatus(v int32) *GetPrecisionTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetTaskId(v string) *GetPrecisionTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetTotalCount(v int32) *GetPrecisionTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetUpdateTime(v string) *GetPrecisionTaskResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) SetVerifiedCount(v int32) *GetPrecisionTaskResponseBodyData {
	s.VerifiedCount = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetPrecisionTaskResponseBodyDataPrecisions struct {
	Precision []*GetPrecisionTaskResponseBodyDataPrecisionsPrecision `json:"Precision,omitempty" xml:"Precision,omitempty" type:"Repeated"`
}

func (s GetPrecisionTaskResponseBodyDataPrecisions) String() string {
	return dara.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyDataPrecisions) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyDataPrecisions) GetPrecision() []*GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	return s.Precision
}

func (s *GetPrecisionTaskResponseBodyDataPrecisions) SetPrecision(v []*GetPrecisionTaskResponseBodyDataPrecisionsPrecision) *GetPrecisionTaskResponseBodyDataPrecisions {
	s.Precision = v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisions) Validate() error {
	return dara.Validate(s)
}

type GetPrecisionTaskResponseBodyDataPrecisionsPrecision struct {
	// example:
	//
	// 2311
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
	// 593A04C0-E6E9-4CDC-8C99-B120C******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetPrecisionTaskResponseBodyDataPrecisionsPrecision) String() string {
	return dara.Prettify(s)
}

func (s GetPrecisionTaskResponseBodyDataPrecisionsPrecision) GoString() string {
	return s.String()
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) GetModelId() *int64 {
	return s.ModelId
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) GetModelName() *string {
	return s.ModelName
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) GetPrecision() *float32 {
	return s.Precision
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) GetStatus() *int32 {
	return s.Status
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) GetTaskId() *string {
	return s.TaskId
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetModelId(v int64) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.ModelId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetModelName(v string) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.ModelName = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetPrecision(v float32) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.Precision = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetStatus(v int32) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.Status = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) SetTaskId(v string) *GetPrecisionTaskResponseBodyDataPrecisionsPrecision {
	s.TaskId = &v
	return s
}

func (s *GetPrecisionTaskResponseBodyDataPrecisionsPrecision) Validate() error {
	return dara.Validate(s)
}
