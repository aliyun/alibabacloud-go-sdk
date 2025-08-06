// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayQoeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int64) *DescribePlayQoeListResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *DescribePlayQoeListResponseBody
	GetPageSize() *int64
	SetQoeInfoList(v []*DescribePlayQoeListResponseBodyQoeInfoList) *DescribePlayQoeListResponseBody
	GetQoeInfoList() []*DescribePlayQoeListResponseBodyQoeInfoList
	SetRequestId(v string) *DescribePlayQoeListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePlayQoeListResponseBody
	GetTotalCount() *int64
}

type DescribePlayQoeListResponseBody struct {
	PageNo      *int64                                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int64                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QoeInfoList []*DescribePlayQoeListResponseBodyQoeInfoList `json:"QoeInfoList,omitempty" xml:"QoeInfoList,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlayQoeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQoeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayQoeListResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribePlayQoeListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePlayQoeListResponseBody) GetQoeInfoList() []*DescribePlayQoeListResponseBodyQoeInfoList {
	return s.QoeInfoList
}

func (s *DescribePlayQoeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayQoeListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePlayQoeListResponseBody) SetPageNo(v int64) *DescribePlayQoeListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribePlayQoeListResponseBody) SetPageSize(v int64) *DescribePlayQoeListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePlayQoeListResponseBody) SetQoeInfoList(v []*DescribePlayQoeListResponseBodyQoeInfoList) *DescribePlayQoeListResponseBody {
	s.QoeInfoList = v
	return s
}

func (s *DescribePlayQoeListResponseBody) SetRequestId(v string) *DescribePlayQoeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayQoeListResponseBody) SetTotalCount(v int64) *DescribePlayQoeListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePlayQoeListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayQoeListResponseBodyQoeInfoList struct {
	QoeFinishedVV      *float32 `json:"QoeFinishedVV,omitempty" xml:"QoeFinishedVV,omitempty"`
	QoeFinishedVVRate  *float32 `json:"QoeFinishedVVRate,omitempty" xml:"QoeFinishedVVRate,omitempty"`
	QoeJumpRate5s      *float32 `json:"QoeJumpRate5s,omitempty" xml:"QoeJumpRate5s,omitempty"`
	QoeUFinishedVVTime *float32 `json:"QoeUFinishedVVTime,omitempty" xml:"QoeUFinishedVVTime,omitempty"`
	QoeUV              *float32 `json:"QoeUV,omitempty" xml:"QoeUV,omitempty"`
	QoeUVVDuration     *float32 `json:"QoeUVVDuration,omitempty" xml:"QoeUVVDuration,omitempty"`
	QoeUVVTime         *float32 `json:"QoeUVVTime,omitempty" xml:"QoeUVVTime,omitempty"`
	QoeVDuration       *float32 `json:"QoeVDuration,omitempty" xml:"QoeVDuration,omitempty"`
	QoeVVDuration      *float32 `json:"QoeVVDuration,omitempty" xml:"QoeVVDuration,omitempty"`
	TraceId            *string  `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribePlayQoeListResponseBodyQoeInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQoeListResponseBodyQoeInfoList) GoString() string {
	return s.String()
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeFinishedVV() *float32 {
	return s.QoeFinishedVV
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeFinishedVVRate() *float32 {
	return s.QoeFinishedVVRate
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeJumpRate5s() *float32 {
	return s.QoeJumpRate5s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeUFinishedVVTime() *float32 {
	return s.QoeUFinishedVVTime
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeUV() *float32 {
	return s.QoeUV
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeUVVDuration() *float32 {
	return s.QoeUVVDuration
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeUVVTime() *float32 {
	return s.QoeUVVTime
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeVDuration() *float32 {
	return s.QoeVDuration
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetQoeVVDuration() *float32 {
	return s.QoeVVDuration
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeFinishedVV(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeFinishedVV = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeFinishedVVRate(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeFinishedVVRate = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeJumpRate5s(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeJumpRate5s = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeUFinishedVVTime(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeUFinishedVVTime = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeUV(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeUV = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeUVVDuration(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeUVVDuration = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeUVVTime(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeUVVTime = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeVDuration(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeVDuration = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetQoeVVDuration(v float32) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.QoeVVDuration = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) SetTraceId(v string) *DescribePlayQoeListResponseBodyQoeInfoList {
	s.TraceId = &v
	return s
}

func (s *DescribePlayQoeListResponseBodyQoeInfoList) Validate() error {
	return dara.Validate(s)
}
