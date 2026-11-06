// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionReportDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySessionReportDataResponseBody
	GetCode() *string
	SetCount(v int32) *QuerySessionReportDataResponseBody
	GetCount() *int32
	SetMessage(v string) *QuerySessionReportDataResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *QuerySessionReportDataResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QuerySessionReportDataResponseBody
	GetPageSize() *int32
	SetReportDataModel(v []*QuerySessionReportDataResponseBodyReportDataModel) *QuerySessionReportDataResponseBody
	GetReportDataModel() []*QuerySessionReportDataResponseBodyReportDataModel
	SetRequestId(v string) *QuerySessionReportDataResponseBody
	GetRequestId() *string
}

type QuerySessionReportDataResponseBody struct {
	// The business status code. A value of `200` indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of records that meet the query conditions.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the request processing result. If the request is successful, `success` is usually returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the current response.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page in the current response.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of session report data.
	ReportDataModel []*QuerySessionReportDataResponseBodyReportDataModel `json:"ReportDataModel,omitempty" xml:"ReportDataModel,omitempty" type:"Repeated"`
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySessionReportDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionReportDataResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionReportDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySessionReportDataResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QuerySessionReportDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySessionReportDataResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QuerySessionReportDataResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySessionReportDataResponseBody) GetReportDataModel() []*QuerySessionReportDataResponseBodyReportDataModel {
	return s.ReportDataModel
}

func (s *QuerySessionReportDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySessionReportDataResponseBody) SetCode(v string) *QuerySessionReportDataResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySessionReportDataResponseBody) SetCount(v int32) *QuerySessionReportDataResponseBody {
	s.Count = &v
	return s
}

func (s *QuerySessionReportDataResponseBody) SetMessage(v string) *QuerySessionReportDataResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySessionReportDataResponseBody) SetPageNumber(v int32) *QuerySessionReportDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QuerySessionReportDataResponseBody) SetPageSize(v int32) *QuerySessionReportDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *QuerySessionReportDataResponseBody) SetReportDataModel(v []*QuerySessionReportDataResponseBodyReportDataModel) *QuerySessionReportDataResponseBody {
	s.ReportDataModel = v
	return s
}

func (s *QuerySessionReportDataResponseBody) SetRequestId(v string) *QuerySessionReportDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySessionReportDataResponseBody) Validate() error {
	if s.ReportDataModel != nil {
		for _, item := range s.ReportDataModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySessionReportDataResponseBodyReportDataModel struct {
	// The number of sessions in the connected state in this record.
	//
	// example:
	//
	// 30
	ConnectSessionCount *int64 `json:"ConnectSessionCount,omitempty" xml:"ConnectSessionCount,omitempty"`
	// The total number of sessions corresponding to this record.
	//
	// example:
	//
	// 36
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The statistical time. The `1d` granularity returns `yyyy-MM-dd`. The `1m`, `30m`, and `6h` granularities return `yyyy-MM-dd HH:mm`. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
	// The number of sessions established through streaming access in this record.
	//
	// example:
	//
	// 24
	StreamingSessionCount *int64 `json:"StreamingSessionCount,omitempty" xml:"StreamingSessionCount,omitempty"`
}

func (s QuerySessionReportDataResponseBodyReportDataModel) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionReportDataResponseBodyReportDataModel) GoString() string {
	return s.String()
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) GetConnectSessionCount() *int64 {
	return s.ConnectSessionCount
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) GetStatDate() *string {
	return s.StatDate
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) GetStreamingSessionCount() *int64 {
	return s.StreamingSessionCount
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) SetConnectSessionCount(v int64) *QuerySessionReportDataResponseBodyReportDataModel {
	s.ConnectSessionCount = &v
	return s
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) SetSessionCount(v int64) *QuerySessionReportDataResponseBodyReportDataModel {
	s.SessionCount = &v
	return s
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) SetStatDate(v string) *QuerySessionReportDataResponseBodyReportDataModel {
	s.StatDate = &v
	return s
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) SetStreamingSessionCount(v int64) *QuerySessionReportDataResponseBodyReportDataModel {
	s.StreamingSessionCount = &v
	return s
}

func (s *QuerySessionReportDataResponseBodyReportDataModel) Validate() error {
	return dara.Validate(s)
}
