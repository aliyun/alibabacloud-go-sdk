// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeScoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRangeScoreResponseBody
	GetCode() *string
	SetData(v []*GetRangeScoreResponseBodyData) *GetRangeScoreResponseBody
	GetData() []*GetRangeScoreResponseBodyData
	SetMessage(v string) *GetRangeScoreResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRangeScoreResponseBody
	GetRequestId() *string
	SetTotal(v float32) *GetRangeScoreResponseBody
	GetTotal() *float32
}

type GetRangeScoreResponseBody struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetRangeScoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 2
	Total *float32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetRangeScoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRangeScoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetRangeScoreResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRangeScoreResponseBody) GetData() []*GetRangeScoreResponseBodyData {
	return s.Data
}

func (s *GetRangeScoreResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRangeScoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRangeScoreResponseBody) GetTotal() *float32 {
	return s.Total
}

func (s *GetRangeScoreResponseBody) SetCode(v string) *GetRangeScoreResponseBody {
	s.Code = &v
	return s
}

func (s *GetRangeScoreResponseBody) SetData(v []*GetRangeScoreResponseBodyData) *GetRangeScoreResponseBody {
	s.Data = v
	return s
}

func (s *GetRangeScoreResponseBody) SetMessage(v string) *GetRangeScoreResponseBody {
	s.Message = &v
	return s
}

func (s *GetRangeScoreResponseBody) SetRequestId(v string) *GetRangeScoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRangeScoreResponseBody) SetTotal(v float32) *GetRangeScoreResponseBody {
	s.Total = &v
	return s
}

func (s *GetRangeScoreResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRangeScoreResponseBodyData struct {
	// example:
	//
	// 1725797727754
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// saturation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 100
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRangeScoreResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRangeScoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRangeScoreResponseBodyData) GetTime() *int64 {
	return s.Time
}

func (s *GetRangeScoreResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetRangeScoreResponseBodyData) GetValue() *int64 {
	return s.Value
}

func (s *GetRangeScoreResponseBodyData) SetTime(v int64) *GetRangeScoreResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetRangeScoreResponseBodyData) SetType(v string) *GetRangeScoreResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetRangeScoreResponseBodyData) SetValue(v int64) *GetRangeScoreResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetRangeScoreResponseBodyData) Validate() error {
	return dara.Validate(s)
}
