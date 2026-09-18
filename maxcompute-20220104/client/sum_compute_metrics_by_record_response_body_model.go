// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumComputeMetricsByRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SumComputeMetricsByRecordResponseBodyData) *SumComputeMetricsByRecordResponseBody
	GetData() []*SumComputeMetricsByRecordResponseBodyData
	SetHttpCode(v int32) *SumComputeMetricsByRecordResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SumComputeMetricsByRecordResponseBody
	GetRequestId() *string
}

type SumComputeMetricsByRecordResponseBody struct {
	// The response data.
	Data []*SumComputeMetricsByRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// - 1xx: Informational. The server has received the request and is processing it.
	//
	// - 2xx: Success. The server successfully received, understood, and accepted the request.
	//
	// - 3xx: Redirection. The client must take further action to complete the request.
	//
	// - 4xx: Client-side error. The request contains invalid syntax or parameters and cannot be fulfilled.
	//
	// - 5xx: Server-side error. The server failed to fulfill a valid request.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0a06dfe517540143853845404e83af
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SumComputeMetricsByRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByRecordResponseBody) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByRecordResponseBody) GetData() []*SumComputeMetricsByRecordResponseBodyData {
	return s.Data
}

func (s *SumComputeMetricsByRecordResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SumComputeMetricsByRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SumComputeMetricsByRecordResponseBody) SetData(v []*SumComputeMetricsByRecordResponseBodyData) *SumComputeMetricsByRecordResponseBody {
	s.Data = v
	return s
}

func (s *SumComputeMetricsByRecordResponseBody) SetHttpCode(v int32) *SumComputeMetricsByRecordResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SumComputeMetricsByRecordResponseBody) SetRequestId(v string) *SumComputeMetricsByRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *SumComputeMetricsByRecordResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumComputeMetricsByRecordResponseBodyData struct {
	// A list of daily usage records.
	DailyComputeRecords []*SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords `json:"dailyComputeRecords,omitempty" xml:"dailyComputeRecords,omitempty" type:"Repeated"`
	// The usage type. For example: ComputationSql
	//
	// example:
	//
	// ComputationSql
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SumComputeMetricsByRecordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByRecordResponseBodyData) GetDailyComputeRecords() []*SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords {
	return s.DailyComputeRecords
}

func (s *SumComputeMetricsByRecordResponseBodyData) GetType() *string {
	return s.Type
}

func (s *SumComputeMetricsByRecordResponseBodyData) SetDailyComputeRecords(v []*SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) *SumComputeMetricsByRecordResponseBodyData {
	s.DailyComputeRecords = v
	return s
}

func (s *SumComputeMetricsByRecordResponseBodyData) SetType(v string) *SumComputeMetricsByRecordResponseBodyData {
	s.Type = &v
	return s
}

func (s *SumComputeMetricsByRecordResponseBodyData) Validate() error {
	if s.DailyComputeRecords != nil {
		for _, item := range s.DailyComputeRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords struct {
	// The statistics date. The format is yyyyMMdd.
	//
	// example:
	//
	// 20260411
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// This day\\"s usage as a percentage of the total usage for the specified period. The value does not include a percent sign (%).
	//
	// example:
	//
	// 50
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage,omitempty"`
	// The record count.
	//
	// example:
	//
	// 1200
	Record *string `json:"record,omitempty" xml:"record,omitempty"`
}

func (s SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) GetDateTime() *string {
	return s.DateTime
}

func (s *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) GetPercentage() *float64 {
	return s.Percentage
}

func (s *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) GetRecord() *string {
	return s.Record
}

func (s *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) SetDateTime(v string) *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords {
	s.DateTime = &v
	return s
}

func (s *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) SetPercentage(v float64) *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords {
	s.Percentage = &v
	return s
}

func (s *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) SetRecord(v string) *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords {
	s.Record = &v
	return s
}

func (s *SumComputeMetricsByRecordResponseBodyDataDailyComputeRecords) Validate() error {
	return dara.Validate(s)
}
