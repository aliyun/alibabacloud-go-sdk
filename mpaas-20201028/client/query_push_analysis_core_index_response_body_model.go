// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisCoreIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPushAnalysisCoreIndexResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryPushAnalysisCoreIndexResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryPushAnalysisCoreIndexResponseBodyResultContent) *QueryPushAnalysisCoreIndexResponseBody
	GetResultContent() *QueryPushAnalysisCoreIndexResponseBodyResultContent
	SetResultMessage(v string) *QueryPushAnalysisCoreIndexResponseBody
	GetResultMessage() *string
}

type QueryPushAnalysisCoreIndexResponseBody struct {
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryPushAnalysisCoreIndexResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                              `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryPushAnalysisCoreIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisCoreIndexResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisCoreIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPushAnalysisCoreIndexResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryPushAnalysisCoreIndexResponseBody) GetResultContent() *QueryPushAnalysisCoreIndexResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryPushAnalysisCoreIndexResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryPushAnalysisCoreIndexResponseBody) SetRequestId(v string) *QueryPushAnalysisCoreIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBody) SetResultCode(v string) *QueryPushAnalysisCoreIndexResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBody) SetResultContent(v *QueryPushAnalysisCoreIndexResponseBodyResultContent) *QueryPushAnalysisCoreIndexResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBody) SetResultMessage(v string) *QueryPushAnalysisCoreIndexResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPushAnalysisCoreIndexResponseBodyResultContent struct {
	Data *QueryPushAnalysisCoreIndexResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryPushAnalysisCoreIndexResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisCoreIndexResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContent) GetData() *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	return s.Data
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContent) SetData(v *QueryPushAnalysisCoreIndexResponseBodyResultContentData) *QueryPushAnalysisCoreIndexResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryPushAnalysisCoreIndexResponseBodyResultContentData struct {
	ArrivalNum   *float32 `json:"ArrivalNum,omitempty" xml:"ArrivalNum,omitempty"`
	ArrivalRate  *float32 `json:"ArrivalRate,omitempty" xml:"ArrivalRate,omitempty"`
	IgnoreNum    *float32 `json:"IgnoreNum,omitempty" xml:"IgnoreNum,omitempty"`
	IgnoreRate   *float32 `json:"IgnoreRate,omitempty" xml:"IgnoreRate,omitempty"`
	OpenNum      *float32 `json:"OpenNum,omitempty" xml:"OpenNum,omitempty"`
	OpenRate     *float32 `json:"OpenRate,omitempty" xml:"OpenRate,omitempty"`
	PushNum      *float32 `json:"PushNum,omitempty" xml:"PushNum,omitempty"`
	PushTotalNum *float32 `json:"PushTotalNum,omitempty" xml:"PushTotalNum,omitempty"`
}

func (s QueryPushAnalysisCoreIndexResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisCoreIndexResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) GetArrivalNum() *float32 {
	return s.ArrivalNum
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) GetArrivalRate() *float32 {
	return s.ArrivalRate
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) GetIgnoreNum() *float32 {
	return s.IgnoreNum
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) GetIgnoreRate() *float32 {
	return s.IgnoreRate
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) GetOpenNum() *float32 {
	return s.OpenNum
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) GetOpenRate() *float32 {
	return s.OpenRate
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) GetPushNum() *float32 {
	return s.PushNum
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) GetPushTotalNum() *float32 {
	return s.PushTotalNum
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) SetArrivalNum(v float32) *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	s.ArrivalNum = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) SetArrivalRate(v float32) *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	s.ArrivalRate = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) SetIgnoreNum(v float32) *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	s.IgnoreNum = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) SetIgnoreRate(v float32) *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	s.IgnoreRate = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) SetOpenNum(v float32) *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	s.OpenNum = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) SetOpenRate(v float32) *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	s.OpenRate = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) SetPushNum(v float32) *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	s.PushNum = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) SetPushTotalNum(v float32) *QueryPushAnalysisCoreIndexResponseBodyResultContentData {
	s.PushTotalNum = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}
