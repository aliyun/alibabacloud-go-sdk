// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisCoreIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAnalysisCoreIndexResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListAnalysisCoreIndexResponseBody
	GetResultCode() *string
	SetResultContent(v *ListAnalysisCoreIndexResponseBodyResultContent) *ListAnalysisCoreIndexResponseBody
	GetResultContent() *ListAnalysisCoreIndexResponseBodyResultContent
	SetResultMessage(v string) *ListAnalysisCoreIndexResponseBody
	GetResultMessage() *string
	SetSuccess(v bool) *ListAnalysisCoreIndexResponseBody
	GetSuccess() *bool
}

type ListAnalysisCoreIndexResponseBody struct {
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ResultCode    *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *ListAnalysisCoreIndexResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAnalysisCoreIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisCoreIndexResponseBody) GoString() string {
	return s.String()
}

func (s *ListAnalysisCoreIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAnalysisCoreIndexResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListAnalysisCoreIndexResponseBody) GetResultContent() *ListAnalysisCoreIndexResponseBodyResultContent {
	return s.ResultContent
}

func (s *ListAnalysisCoreIndexResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListAnalysisCoreIndexResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnalysisCoreIndexResponseBody) SetRequestId(v string) *ListAnalysisCoreIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBody) SetResultCode(v string) *ListAnalysisCoreIndexResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBody) SetResultContent(v *ListAnalysisCoreIndexResponseBodyResultContent) *ListAnalysisCoreIndexResponseBody {
	s.ResultContent = v
	return s
}

func (s *ListAnalysisCoreIndexResponseBody) SetResultMessage(v string) *ListAnalysisCoreIndexResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBody) SetSuccess(v bool) *ListAnalysisCoreIndexResponseBody {
	s.Success = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAnalysisCoreIndexResponseBodyResultContent struct {
	Data *ListAnalysisCoreIndexResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// “”
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAnalysisCoreIndexResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisCoreIndexResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *ListAnalysisCoreIndexResponseBodyResultContent) GetData() *ListAnalysisCoreIndexResponseBodyResultContentData {
	return s.Data
}

func (s *ListAnalysisCoreIndexResponseBodyResultContent) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListAnalysisCoreIndexResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *ListAnalysisCoreIndexResponseBodyResultContent) SetData(v *ListAnalysisCoreIndexResponseBodyResultContentData) *ListAnalysisCoreIndexResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContent) SetResultMsg(v string) *ListAnalysisCoreIndexResponseBodyResultContent {
	s.ResultMsg = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContent) SetSuccess(v bool) *ListAnalysisCoreIndexResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type ListAnalysisCoreIndexResponseBodyResultContentData struct {
	// example:
	//
	// 0
	ArrivalNum *string `json:"ArrivalNum,omitempty" xml:"ArrivalNum,omitempty"`
	// example:
	//
	// 0
	ArrivalRate *string `json:"ArrivalRate,omitempty" xml:"ArrivalRate,omitempty"`
	// example:
	//
	// 0
	IgnoreNum *string `json:"IgnoreNum,omitempty" xml:"IgnoreNum,omitempty"`
	// example:
	//
	// 0
	IgnoreRate *string `json:"IgnoreRate,omitempty" xml:"IgnoreRate,omitempty"`
	// example:
	//
	// 0
	OpenNum *string `json:"OpenNum,omitempty" xml:"OpenNum,omitempty"`
	// example:
	//
	// 0
	OpenRate *string `json:"OpenRate,omitempty" xml:"OpenRate,omitempty"`
	// example:
	//
	// 0
	PushNum *string `json:"PushNum,omitempty" xml:"PushNum,omitempty"`
	// example:
	//
	// 0
	PushTotalNum *string `json:"PushTotalNum,omitempty" xml:"PushTotalNum,omitempty"`
}

func (s ListAnalysisCoreIndexResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisCoreIndexResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) GetArrivalNum() *string {
	return s.ArrivalNum
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) GetArrivalRate() *string {
	return s.ArrivalRate
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) GetIgnoreNum() *string {
	return s.IgnoreNum
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) GetIgnoreRate() *string {
	return s.IgnoreRate
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) GetOpenNum() *string {
	return s.OpenNum
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) GetOpenRate() *string {
	return s.OpenRate
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) GetPushNum() *string {
	return s.PushNum
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) GetPushTotalNum() *string {
	return s.PushTotalNum
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) SetArrivalNum(v string) *ListAnalysisCoreIndexResponseBodyResultContentData {
	s.ArrivalNum = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) SetArrivalRate(v string) *ListAnalysisCoreIndexResponseBodyResultContentData {
	s.ArrivalRate = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) SetIgnoreNum(v string) *ListAnalysisCoreIndexResponseBodyResultContentData {
	s.IgnoreNum = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) SetIgnoreRate(v string) *ListAnalysisCoreIndexResponseBodyResultContentData {
	s.IgnoreRate = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) SetOpenNum(v string) *ListAnalysisCoreIndexResponseBodyResultContentData {
	s.OpenNum = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) SetOpenRate(v string) *ListAnalysisCoreIndexResponseBodyResultContentData {
	s.OpenRate = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) SetPushNum(v string) *ListAnalysisCoreIndexResponseBodyResultContentData {
	s.PushNum = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) SetPushTotalNum(v string) *ListAnalysisCoreIndexResponseBodyResultContentData {
	s.PushTotalNum = &v
	return s
}

func (s *ListAnalysisCoreIndexResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
