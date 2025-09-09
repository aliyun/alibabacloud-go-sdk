// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventCountByThreatLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEventCountByThreatLevelResponseBody
	GetCode() *int32
	SetData(v *DescribeEventCountByThreatLevelResponseBodyData) *DescribeEventCountByThreatLevelResponseBody
	GetData() *DescribeEventCountByThreatLevelResponseBodyData
	SetMessage(v string) *DescribeEventCountByThreatLevelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEventCountByThreatLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEventCountByThreatLevelResponseBody
	GetSuccess() *bool
}

type DescribeEventCountByThreatLevelResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeEventCountByThreatLevelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEventCountByThreatLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEventCountByThreatLevelResponseBody) GetData() *DescribeEventCountByThreatLevelResponseBodyData {
	return s.Data
}

func (s *DescribeEventCountByThreatLevelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventCountByThreatLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventCountByThreatLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetCode(v int32) *DescribeEventCountByThreatLevelResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetData(v *DescribeEventCountByThreatLevelResponseBodyData) *DescribeEventCountByThreatLevelResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetMessage(v string) *DescribeEventCountByThreatLevelResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetRequestId(v string) *DescribeEventCountByThreatLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) SetSuccess(v bool) *DescribeEventCountByThreatLevelResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventCountByThreatLevelResponseBodyData struct {
	// The total number of events.
	//
	// example:
	//
	// 100
	EventNum *int64 `json:"EventNum,omitempty" xml:"EventNum,omitempty"`
	// The number of high-risk events.
	//
	// example:
	//
	// 20
	HighLevelEventNum *int64 `json:"HighLevelEventNum,omitempty" xml:"HighLevelEventNum,omitempty"`
	InfoLevelEventNum *int64 `json:"InfoLevelEventNum,omitempty" xml:"InfoLevelEventNum,omitempty"`
	// The number of low-risk events.
	//
	// example:
	//
	// 52
	LowLevelEventNum *int64 `json:"LowLevelEventNum,omitempty" xml:"LowLevelEventNum,omitempty"`
	// The number of medium-risk events.
	//
	// example:
	//
	// 3
	MediumLevelEventNum  *int64 `json:"MediumLevelEventNum,omitempty" xml:"MediumLevelEventNum,omitempty"`
	SeriousLevelEventNum *int64 `json:"SeriousLevelEventNum,omitempty" xml:"SeriousLevelEventNum,omitempty"`
	// The number of unhandled events.
	//
	// example:
	//
	// 75
	UndealEventNum *int64 `json:"UndealEventNum,omitempty" xml:"UndealEventNum,omitempty"`
}

func (s DescribeEventCountByThreatLevelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventCountByThreatLevelResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) GetEventNum() *int64 {
	return s.EventNum
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) GetHighLevelEventNum() *int64 {
	return s.HighLevelEventNum
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) GetInfoLevelEventNum() *int64 {
	return s.InfoLevelEventNum
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) GetLowLevelEventNum() *int64 {
	return s.LowLevelEventNum
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) GetMediumLevelEventNum() *int64 {
	return s.MediumLevelEventNum
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) GetSeriousLevelEventNum() *int64 {
	return s.SeriousLevelEventNum
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) GetUndealEventNum() *int64 {
	return s.UndealEventNum
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.EventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetHighLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.HighLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetInfoLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.InfoLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetLowLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.LowLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetMediumLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.MediumLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetSeriousLevelEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.SeriousLevelEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) SetUndealEventNum(v int64) *DescribeEventCountByThreatLevelResponseBodyData {
	s.UndealEventNum = &v
	return s
}

func (s *DescribeEventCountByThreatLevelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
