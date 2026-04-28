// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInspectProgressResponseBody
	GetCode() *string
	SetData(v *GetInspectProgressResponseBodyData) *GetInspectProgressResponseBody
	GetData() *GetInspectProgressResponseBodyData
	SetMessage(v string) *GetInspectProgressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInspectProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInspectProgressResponseBody
	GetSuccess() *bool
}

type GetInspectProgressResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInspectProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInspectProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInspectProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetInspectProgressResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInspectProgressResponseBody) GetData() *GetInspectProgressResponseBodyData {
	return s.Data
}

func (s *GetInspectProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInspectProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInspectProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInspectProgressResponseBody) SetCode(v string) *GetInspectProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetInspectProgressResponseBody) SetData(v *GetInspectProgressResponseBodyData) *GetInspectProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetInspectProgressResponseBody) SetMessage(v string) *GetInspectProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetInspectProgressResponseBody) SetRequestId(v string) *GetInspectProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInspectProgressResponseBody) SetSuccess(v bool) *GetInspectProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetInspectProgressResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInspectProgressResponseBodyData struct {
	// example:
	//
	// 100
	AllSubtaskCount *int32 `json:"AllSubtaskCount,omitempty" xml:"AllSubtaskCount,omitempty"`
	// example:
	//
	// True
	Finish *bool `json:"Finish,omitempty" xml:"Finish,omitempty"`
	// example:
	//
	// 1
	FinishRate *float64 `json:"FinishRate,omitempty" xml:"FinishRate,omitempty"`
	// example:
	//
	// 1
	FinishSubtaskCount *int32 `json:"FinishSubtaskCount,omitempty" xml:"FinishSubtaskCount,omitempty"`
	// example:
	//
	// 111
	LastInspectDate *int64 `json:"LastInspectDate,omitempty" xml:"LastInspectDate,omitempty"`
	// example:
	//
	// 95***135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s GetInspectProgressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInspectProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInspectProgressResponseBodyData) GetAllSubtaskCount() *int32 {
	return s.AllSubtaskCount
}

func (s *GetInspectProgressResponseBodyData) GetFinish() *bool {
	return s.Finish
}

func (s *GetInspectProgressResponseBodyData) GetFinishRate() *float64 {
	return s.FinishRate
}

func (s *GetInspectProgressResponseBodyData) GetFinishSubtaskCount() *int32 {
	return s.FinishSubtaskCount
}

func (s *GetInspectProgressResponseBodyData) GetLastInspectDate() *int64 {
	return s.LastInspectDate
}

func (s *GetInspectProgressResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetInspectProgressResponseBodyData) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *GetInspectProgressResponseBodyData) SetAllSubtaskCount(v int32) *GetInspectProgressResponseBodyData {
	s.AllSubtaskCount = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetFinish(v bool) *GetInspectProgressResponseBodyData {
	s.Finish = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetFinishRate(v float64) *GetInspectProgressResponseBodyData {
	s.FinishRate = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetFinishSubtaskCount(v int32) *GetInspectProgressResponseBodyData {
	s.FinishSubtaskCount = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetLastInspectDate(v int64) *GetInspectProgressResponseBodyData {
	s.LastInspectDate = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetTaskId(v int64) *GetInspectProgressResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) SetUsedTime(v int64) *GetInspectProgressResponseBodyData {
	s.UsedTime = &v
	return s
}

func (s *GetInspectProgressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
