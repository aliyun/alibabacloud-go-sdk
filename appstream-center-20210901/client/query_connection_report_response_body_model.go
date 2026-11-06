// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConnectionReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryConnectionReportResponseBody
	GetCode() *string
	SetConnectionReportDataModel(v []*QueryConnectionReportResponseBodyConnectionReportDataModel) *QueryConnectionReportResponseBody
	GetConnectionReportDataModel() []*QueryConnectionReportResponseBodyConnectionReportDataModel
	SetCount(v int32) *QueryConnectionReportResponseBody
	GetCount() *int32
	SetMessage(v string) *QueryConnectionReportResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *QueryConnectionReportResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryConnectionReportResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryConnectionReportResponseBody
	GetRequestId() *string
}

type QueryConnectionReportResponseBody struct {
	// The business status code. `200` indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of connection report data.
	ConnectionReportDataModel []*QueryConnectionReportResponseBodyConnectionReportDataModel `json:"ConnectionReportDataModel,omitempty" xml:"ConnectionReportDataModel,omitempty" type:"Repeated"`
	// The total number of records that meet the query conditions.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The description of the request processing result. `success` is usually returned when the request is successful.
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
	// The unique identifier of the current request, which can be used to troubleshoot the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryConnectionReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConnectionReportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConnectionReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryConnectionReportResponseBody) GetConnectionReportDataModel() []*QueryConnectionReportResponseBodyConnectionReportDataModel {
	return s.ConnectionReportDataModel
}

func (s *QueryConnectionReportResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QueryConnectionReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryConnectionReportResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryConnectionReportResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryConnectionReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConnectionReportResponseBody) SetCode(v string) *QueryConnectionReportResponseBody {
	s.Code = &v
	return s
}

func (s *QueryConnectionReportResponseBody) SetConnectionReportDataModel(v []*QueryConnectionReportResponseBodyConnectionReportDataModel) *QueryConnectionReportResponseBody {
	s.ConnectionReportDataModel = v
	return s
}

func (s *QueryConnectionReportResponseBody) SetCount(v int32) *QueryConnectionReportResponseBody {
	s.Count = &v
	return s
}

func (s *QueryConnectionReportResponseBody) SetMessage(v string) *QueryConnectionReportResponseBody {
	s.Message = &v
	return s
}

func (s *QueryConnectionReportResponseBody) SetPageNumber(v int32) *QueryConnectionReportResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryConnectionReportResponseBody) SetPageSize(v int32) *QueryConnectionReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryConnectionReportResponseBody) SetRequestId(v string) *QueryConnectionReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConnectionReportResponseBody) Validate() error {
	if s.ConnectionReportDataModel != nil {
		for _, item := range s.ConnectionReportDataModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryConnectionReportResponseBodyConnectionReportDataModel struct {
	// The ID of the app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The average duration of successful connections. In the current version, the return value is in milliseconds, so parse it as milliseconds. The `Seconds` in the field name is retained for compatibility and does not indicate the unit. This value may not be returned if there are no successful connections.
	//
	// example:
	//
	// 1250
	ConnectionAvgDurationSeconds *float64 `json:"ConnectionAvgDurationSeconds,omitempty" xml:"ConnectionAvgDurationSeconds,omitempty"`
	// The number of connection failures within the statistical window.
	//
	// example:
	//
	// 2
	ConnectionFailureCount *int64 `json:"ConnectionFailureCount,omitempty" xml:"ConnectionFailureCount,omitempty"`
	// The number of successful connections within the statistical window.
	//
	// example:
	//
	// 98
	ConnectionSuccessCount *int64 `json:"ConnectionSuccessCount,omitempty" xml:"ConnectionSuccessCount,omitempty"`
	// The connection success rate, which is calculated by dividing the number of successful connections by the sum of successful and failed connections. The value ranges from 0 to 1. Multiply the value by 100 to get the percentage. This value may not be returned if there are no connection records.
	//
	// example:
	//
	// 0.98
	ConnectionSuccessRate *float64 `json:"ConnectionSuccessRate,omitempty" xml:"ConnectionSuccessRate,omitempty"`
	// The ID of the business region to which the report data belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The end time of the statistical window. Format: `yyyy-MM-dd HH:mm:ss`. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01 10:05:00
	StatEndDate *string `json:"StatEndDate,omitempty" xml:"StatEndDate,omitempty"`
	// The start time of the statistical window. Format: `yyyy-MM-dd HH:mm:ss`. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01 10:00:00
	StatStartDate *string `json:"StatStartDate,omitempty" xml:"StatStartDate,omitempty"`
}

func (s QueryConnectionReportResponseBodyConnectionReportDataModel) String() string {
	return dara.Prettify(s)
}

func (s QueryConnectionReportResponseBodyConnectionReportDataModel) GoString() string {
	return s.String()
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) GetConnectionAvgDurationSeconds() *float64 {
	return s.ConnectionAvgDurationSeconds
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) GetConnectionFailureCount() *int64 {
	return s.ConnectionFailureCount
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) GetConnectionSuccessCount() *int64 {
	return s.ConnectionSuccessCount
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) GetConnectionSuccessRate() *float64 {
	return s.ConnectionSuccessRate
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) GetStatEndDate() *string {
	return s.StatEndDate
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) GetStatStartDate() *string {
	return s.StatStartDate
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) SetAppInstanceGroupId(v string) *QueryConnectionReportResponseBodyConnectionReportDataModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) SetConnectionAvgDurationSeconds(v float64) *QueryConnectionReportResponseBodyConnectionReportDataModel {
	s.ConnectionAvgDurationSeconds = &v
	return s
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) SetConnectionFailureCount(v int64) *QueryConnectionReportResponseBodyConnectionReportDataModel {
	s.ConnectionFailureCount = &v
	return s
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) SetConnectionSuccessCount(v int64) *QueryConnectionReportResponseBodyConnectionReportDataModel {
	s.ConnectionSuccessCount = &v
	return s
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) SetConnectionSuccessRate(v float64) *QueryConnectionReportResponseBodyConnectionReportDataModel {
	s.ConnectionSuccessRate = &v
	return s
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) SetRegionId(v string) *QueryConnectionReportResponseBodyConnectionReportDataModel {
	s.RegionId = &v
	return s
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) SetStatEndDate(v string) *QueryConnectionReportResponseBodyConnectionReportDataModel {
	s.StatEndDate = &v
	return s
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) SetStatStartDate(v string) *QueryConnectionReportResponseBodyConnectionReportDataModel {
	s.StatStartDate = &v
	return s
}

func (s *QueryConnectionReportResponseBodyConnectionReportDataModel) Validate() error {
	return dara.Validate(s)
}
