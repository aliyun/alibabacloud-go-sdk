// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNodeReportDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryNodeReportDataResponseBody
	GetCode() *string
	SetCount(v int32) *QueryNodeReportDataResponseBody
	GetCount() *int32
	SetMessage(v string) *QueryNodeReportDataResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *QueryNodeReportDataResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryNodeReportDataResponseBody
	GetPageSize() *int32
	SetReportDataModel(v []*QueryNodeReportDataResponseBodyReportDataModel) *QueryNodeReportDataResponseBody
	GetReportDataModel() []*QueryNodeReportDataResponseBodyReportDataModel
	SetRequestId(v string) *QueryNodeReportDataResponseBody
	GetRequestId() *string
}

type QueryNodeReportDataResponseBody struct {
	// The business status code. A value of 200 indicates that the request is successful. Other values indicate that the request failed.
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
	// The description of the request processing result. If the request is successful, success is usually returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number corresponding to the current response.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page used in the current response.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of node report data.
	ReportDataModel []*QueryNodeReportDataResponseBodyReportDataModel `json:"ReportDataModel,omitempty" xml:"ReportDataModel,omitempty" type:"Repeated"`
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryNodeReportDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryNodeReportDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNodeReportDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryNodeReportDataResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QueryNodeReportDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryNodeReportDataResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryNodeReportDataResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryNodeReportDataResponseBody) GetReportDataModel() []*QueryNodeReportDataResponseBodyReportDataModel {
	return s.ReportDataModel
}

func (s *QueryNodeReportDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryNodeReportDataResponseBody) SetCode(v string) *QueryNodeReportDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryNodeReportDataResponseBody) SetCount(v int32) *QueryNodeReportDataResponseBody {
	s.Count = &v
	return s
}

func (s *QueryNodeReportDataResponseBody) SetMessage(v string) *QueryNodeReportDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryNodeReportDataResponseBody) SetPageNumber(v int32) *QueryNodeReportDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryNodeReportDataResponseBody) SetPageSize(v int32) *QueryNodeReportDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryNodeReportDataResponseBody) SetReportDataModel(v []*QueryNodeReportDataResponseBodyReportDataModel) *QueryNodeReportDataResponseBody {
	s.ReportDataModel = v
	return s
}

func (s *QueryNodeReportDataResponseBody) SetRequestId(v string) *QueryNodeReportDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNodeReportDataResponseBody) Validate() error {
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

type QueryNodeReportDataResponseBodyReportDataModel struct {
	// The total number of nodes for the record.
	//
	// example:
	//
	// 20
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The number of pay-as-you-go nodes in the record.
	//
	// example:
	//
	// 8
	PostpaidNodeCount *int64 `json:"PostpaidNodeCount,omitempty" xml:"PostpaidNodeCount,omitempty"`
	// The number of subscription nodes in the record.
	//
	// example:
	//
	// 12
	PrepaidNodeCount *int64 `json:"PrepaidNodeCount,omitempty" xml:"PrepaidNodeCount,omitempty"`
	// The statistical time. The 1d granularity returns the time in the yyyy-MM-dd format. The 1m, 30m, and 6h granularities return the time in the yyyy-MM-dd HH:mm format. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
}

func (s QueryNodeReportDataResponseBodyReportDataModel) String() string {
	return dara.Prettify(s)
}

func (s QueryNodeReportDataResponseBodyReportDataModel) GoString() string {
	return s.String()
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) GetPostpaidNodeCount() *int64 {
	return s.PostpaidNodeCount
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) GetPrepaidNodeCount() *int64 {
	return s.PrepaidNodeCount
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) GetStatDate() *string {
	return s.StatDate
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) SetNodeCount(v int64) *QueryNodeReportDataResponseBodyReportDataModel {
	s.NodeCount = &v
	return s
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) SetPostpaidNodeCount(v int64) *QueryNodeReportDataResponseBodyReportDataModel {
	s.PostpaidNodeCount = &v
	return s
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) SetPrepaidNodeCount(v int64) *QueryNodeReportDataResponseBodyReportDataModel {
	s.PrepaidNodeCount = &v
	return s
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) SetStatDate(v string) *QueryNodeReportDataResponseBodyReportDataModel {
	s.StatDate = &v
	return s
}

func (s *QueryNodeReportDataResponseBodyReportDataModel) Validate() error {
	return dara.Validate(s)
}
