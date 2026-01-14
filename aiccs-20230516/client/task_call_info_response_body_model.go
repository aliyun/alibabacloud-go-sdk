// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *TaskCallInfoResponseBody
	GetCode() *int64
	SetMessage(v string) *TaskCallInfoResponseBody
	GetMessage() *string
	SetModel(v *TaskCallInfoResponseBodyModel) *TaskCallInfoResponseBody
	GetModel() *TaskCallInfoResponseBodyModel
	SetRequestId(v string) *TaskCallInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TaskCallInfoResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *TaskCallInfoResponseBody
	GetTimestamp() *int64
}

type TaskCallInfoResponseBody struct {
	// example:
	//
	// 15
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *TaskCallInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 62
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s TaskCallInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TaskCallInfoResponseBody) GoString() string {
	return s.String()
}

func (s *TaskCallInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TaskCallInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TaskCallInfoResponseBody) GetModel() *TaskCallInfoResponseBodyModel {
	return s.Model
}

func (s *TaskCallInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskCallInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TaskCallInfoResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *TaskCallInfoResponseBody) SetCode(v int64) *TaskCallInfoResponseBody {
	s.Code = &v
	return s
}

func (s *TaskCallInfoResponseBody) SetMessage(v string) *TaskCallInfoResponseBody {
	s.Message = &v
	return s
}

func (s *TaskCallInfoResponseBody) SetModel(v *TaskCallInfoResponseBodyModel) *TaskCallInfoResponseBody {
	s.Model = v
	return s
}

func (s *TaskCallInfoResponseBody) SetRequestId(v string) *TaskCallInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskCallInfoResponseBody) SetSuccess(v bool) *TaskCallInfoResponseBody {
	s.Success = &v
	return s
}

func (s *TaskCallInfoResponseBody) SetTimestamp(v int64) *TaskCallInfoResponseBody {
	s.Timestamp = &v
	return s
}

func (s *TaskCallInfoResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TaskCallInfoResponseBodyModel struct {
	// example:
	//
	// 75
	FinishTotal *int64 `json:"FinishTotal,omitempty" xml:"FinishTotal,omitempty"`
	// example:
	//
	// 59
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 3
	UnCallTotal *int64 `json:"UnCallTotal,omitempty" xml:"UnCallTotal,omitempty"`
}

func (s TaskCallInfoResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s TaskCallInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *TaskCallInfoResponseBodyModel) GetFinishTotal() *int64 {
	return s.FinishTotal
}

func (s *TaskCallInfoResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *TaskCallInfoResponseBodyModel) GetUnCallTotal() *int64 {
	return s.UnCallTotal
}

func (s *TaskCallInfoResponseBodyModel) SetFinishTotal(v int64) *TaskCallInfoResponseBodyModel {
	s.FinishTotal = &v
	return s
}

func (s *TaskCallInfoResponseBodyModel) SetTotal(v int64) *TaskCallInfoResponseBodyModel {
	s.Total = &v
	return s
}

func (s *TaskCallInfoResponseBodyModel) SetUnCallTotal(v int64) *TaskCallInfoResponseBodyModel {
	s.UnCallTotal = &v
	return s
}

func (s *TaskCallInfoResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
