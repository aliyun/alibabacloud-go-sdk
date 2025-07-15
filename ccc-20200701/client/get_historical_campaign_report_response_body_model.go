// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalCampaignReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHistoricalCampaignReportResponseBody
	GetCode() *string
	SetData(v *GetHistoricalCampaignReportResponseBodyData) *GetHistoricalCampaignReportResponseBody
	GetData() *GetHistoricalCampaignReportResponseBodyData
	SetHttpStatusCode(v int32) *GetHistoricalCampaignReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHistoricalCampaignReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHistoricalCampaignReportResponseBody
	GetRequestId() *string
}

type GetHistoricalCampaignReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetHistoricalCampaignReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E3A847C1-9800-57DF-9172-2CDDC026388D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHistoricalCampaignReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalCampaignReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoricalCampaignReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHistoricalCampaignReportResponseBody) GetData() *GetHistoricalCampaignReportResponseBodyData {
	return s.Data
}

func (s *GetHistoricalCampaignReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHistoricalCampaignReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHistoricalCampaignReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHistoricalCampaignReportResponseBody) SetCode(v string) *GetHistoricalCampaignReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) SetData(v *GetHistoricalCampaignReportResponseBodyData) *GetHistoricalCampaignReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) SetHttpStatusCode(v int32) *GetHistoricalCampaignReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) SetMessage(v string) *GetHistoricalCampaignReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) SetRequestId(v string) *GetHistoricalCampaignReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHistoricalCampaignReportResponseBodyData struct {
	AbandonRate *float32 `json:"AbandonRate,omitempty" xml:"AbandonRate,omitempty"`
	// example:
	//
	// 0.10
	AbandonedRate *float32 `json:"AbandonedRate,omitempty" xml:"AbandonedRate,omitempty"`
	AnswerRate    *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// example:
	//
	// 5
	CallsAbandoned *int64 `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	// example:
	//
	// 50
	CallsConnected *int64 `json:"CallsConnected,omitempty" xml:"CallsConnected,omitempty"`
	// example:
	//
	// 100
	CallsDialed *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	// example:
	//
	// 0.50
	ConnectedRate *float32 `json:"ConnectedRate,omitempty" xml:"ConnectedRate,omitempty"`
	// example:
	//
	// 0.50
	OccupancyRate *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
}

func (s GetHistoricalCampaignReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalCampaignReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoricalCampaignReportResponseBodyData) GetAbandonRate() *float32 {
	return s.AbandonRate
}

func (s *GetHistoricalCampaignReportResponseBodyData) GetAbandonedRate() *float32 {
	return s.AbandonedRate
}

func (s *GetHistoricalCampaignReportResponseBodyData) GetAnswerRate() *float32 {
	return s.AnswerRate
}

func (s *GetHistoricalCampaignReportResponseBodyData) GetCallsAbandoned() *int64 {
	return s.CallsAbandoned
}

func (s *GetHistoricalCampaignReportResponseBodyData) GetCallsConnected() *int64 {
	return s.CallsConnected
}

func (s *GetHistoricalCampaignReportResponseBodyData) GetCallsDialed() *int64 {
	return s.CallsDialed
}

func (s *GetHistoricalCampaignReportResponseBodyData) GetConnectedRate() *float32 {
	return s.ConnectedRate
}

func (s *GetHistoricalCampaignReportResponseBodyData) GetOccupancyRate() *float32 {
	return s.OccupancyRate
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetAbandonRate(v float32) *GetHistoricalCampaignReportResponseBodyData {
	s.AbandonRate = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetAbandonedRate(v float32) *GetHistoricalCampaignReportResponseBodyData {
	s.AbandonedRate = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetAnswerRate(v float32) *GetHistoricalCampaignReportResponseBodyData {
	s.AnswerRate = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetCallsAbandoned(v int64) *GetHistoricalCampaignReportResponseBodyData {
	s.CallsAbandoned = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetCallsConnected(v int64) *GetHistoricalCampaignReportResponseBodyData {
	s.CallsConnected = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetCallsDialed(v int64) *GetHistoricalCampaignReportResponseBodyData {
	s.CallsDialed = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetConnectedRate(v float32) *GetHistoricalCampaignReportResponseBodyData {
	s.ConnectedRate = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetOccupancyRate(v float32) *GetHistoricalCampaignReportResponseBodyData {
	s.OccupancyRate = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
