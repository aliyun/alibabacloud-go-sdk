// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoricalCallerReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHistoricalCallerReportResponseBody
	GetCode() *string
	SetData(v *GetHistoricalCallerReportResponseBodyData) *GetHistoricalCallerReportResponseBody
	GetData() *GetHistoricalCallerReportResponseBodyData
	SetHttpStatusCode(v int32) *GetHistoricalCallerReportResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHistoricalCallerReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHistoricalCallerReportResponseBody
	GetRequestId() *string
}

type GetHistoricalCallerReportResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetHistoricalCallerReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHistoricalCallerReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalCallerReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoricalCallerReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHistoricalCallerReportResponseBody) GetData() *GetHistoricalCallerReportResponseBodyData {
	return s.Data
}

func (s *GetHistoricalCallerReportResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHistoricalCallerReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHistoricalCallerReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHistoricalCallerReportResponseBody) SetCode(v string) *GetHistoricalCallerReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) SetData(v *GetHistoricalCallerReportResponseBodyData) *GetHistoricalCallerReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) SetHttpStatusCode(v int32) *GetHistoricalCallerReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) SetMessage(v string) *GetHistoricalCallerReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) SetRequestId(v string) *GetHistoricalCallerReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHistoricalCallerReportResponseBodyData struct {
	// example:
	//
	// 1646917200000
	LastCallingTime *int64 `json:"LastCallingTime,omitempty" xml:"LastCallingTime,omitempty"`
	// example:
	//
	// 10
	TotalCalls *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s GetHistoricalCallerReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHistoricalCallerReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoricalCallerReportResponseBodyData) GetLastCallingTime() *int64 {
	return s.LastCallingTime
}

func (s *GetHistoricalCallerReportResponseBodyData) GetTotalCalls() *int64 {
	return s.TotalCalls
}

func (s *GetHistoricalCallerReportResponseBodyData) SetLastCallingTime(v int64) *GetHistoricalCallerReportResponseBodyData {
	s.LastCallingTime = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBodyData) SetTotalCalls(v int64) *GetHistoricalCallerReportResponseBodyData {
	s.TotalCalls = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBodyData) Validate() error {
	return dara.Validate(s)
}
