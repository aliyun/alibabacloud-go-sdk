// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketSummaryReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTicketSummaryReportResponseBody
	GetCode() *string
	SetData(v *GetTicketSummaryReportResponseBodyData) *GetTicketSummaryReportResponseBody
	GetData() *GetTicketSummaryReportResponseBodyData
	SetHttpStatusCode(v int64) *GetTicketSummaryReportResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetTicketSummaryReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTicketSummaryReportResponseBody
	GetRequestId() *string
}

type GetTicketSummaryReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTicketSummaryReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EAF3C248-E123-441B-A545-B6CD02E98EED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTicketSummaryReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTicketSummaryReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketSummaryReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTicketSummaryReportResponseBody) GetData() *GetTicketSummaryReportResponseBodyData {
	return s.Data
}

func (s *GetTicketSummaryReportResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetTicketSummaryReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTicketSummaryReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTicketSummaryReportResponseBody) SetCode(v string) *GetTicketSummaryReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketSummaryReportResponseBody) SetData(v *GetTicketSummaryReportResponseBodyData) *GetTicketSummaryReportResponseBody {
	s.Data = v
	return s
}

func (s *GetTicketSummaryReportResponseBody) SetHttpStatusCode(v int64) *GetTicketSummaryReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTicketSummaryReportResponseBody) SetMessage(v string) *GetTicketSummaryReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketSummaryReportResponseBody) SetRequestId(v string) *GetTicketSummaryReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketSummaryReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTicketSummaryReportResponseBodyData struct {
	// example:
	//
	// 3
	TicketsAssigned *string `json:"TicketsAssigned,omitempty" xml:"TicketsAssigned,omitempty"`
	// example:
	//
	// 10
	TicketsCreated *string `json:"TicketsCreated,omitempty" xml:"TicketsCreated,omitempty"`
	// example:
	//
	// 5
	TicketsParticipated *string `json:"TicketsParticipated,omitempty" xml:"TicketsParticipated,omitempty"`
}

func (s GetTicketSummaryReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTicketSummaryReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTicketSummaryReportResponseBodyData) GetTicketsAssigned() *string {
	return s.TicketsAssigned
}

func (s *GetTicketSummaryReportResponseBodyData) GetTicketsCreated() *string {
	return s.TicketsCreated
}

func (s *GetTicketSummaryReportResponseBodyData) GetTicketsParticipated() *string {
	return s.TicketsParticipated
}

func (s *GetTicketSummaryReportResponseBodyData) SetTicketsAssigned(v string) *GetTicketSummaryReportResponseBodyData {
	s.TicketsAssigned = &v
	return s
}

func (s *GetTicketSummaryReportResponseBodyData) SetTicketsCreated(v string) *GetTicketSummaryReportResponseBodyData {
	s.TicketsCreated = &v
	return s
}

func (s *GetTicketSummaryReportResponseBodyData) SetTicketsParticipated(v string) *GetTicketSummaryReportResponseBodyData {
	s.TicketsParticipated = &v
	return s
}

func (s *GetTicketSummaryReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
