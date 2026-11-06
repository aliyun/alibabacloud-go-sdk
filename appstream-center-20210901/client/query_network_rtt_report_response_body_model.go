// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNetworkRttReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryNetworkRttReportResponseBody
	GetCode() *string
	SetCount(v int32) *QueryNetworkRttReportResponseBody
	GetCount() *int32
	SetMessage(v string) *QueryNetworkRttReportResponseBody
	GetMessage() *string
	SetNetworkRttReportDataModel(v []*QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) *QueryNetworkRttReportResponseBody
	GetNetworkRttReportDataModel() []*QueryNetworkRttReportResponseBodyNetworkRttReportDataModel
	SetPageNumber(v int32) *QueryNetworkRttReportResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryNetworkRttReportResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryNetworkRttReportResponseBody
	GetRequestId() *string
}

type QueryNetworkRttReportResponseBody struct {
	// The business status code. `200` indicates that the request is processed successfully. Other values indicate that the processing failed.
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
	// The list of network RTT report data.
	NetworkRttReportDataModel []*QueryNetworkRttReportResponseBodyNetworkRttReportDataModel `json:"NetworkRttReportDataModel,omitempty" xml:"NetworkRttReportDataModel,omitempty" type:"Repeated"`
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
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryNetworkRttReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryNetworkRttReportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNetworkRttReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryNetworkRttReportResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QueryNetworkRttReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryNetworkRttReportResponseBody) GetNetworkRttReportDataModel() []*QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	return s.NetworkRttReportDataModel
}

func (s *QueryNetworkRttReportResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryNetworkRttReportResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryNetworkRttReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryNetworkRttReportResponseBody) SetCode(v string) *QueryNetworkRttReportResponseBody {
	s.Code = &v
	return s
}

func (s *QueryNetworkRttReportResponseBody) SetCount(v int32) *QueryNetworkRttReportResponseBody {
	s.Count = &v
	return s
}

func (s *QueryNetworkRttReportResponseBody) SetMessage(v string) *QueryNetworkRttReportResponseBody {
	s.Message = &v
	return s
}

func (s *QueryNetworkRttReportResponseBody) SetNetworkRttReportDataModel(v []*QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) *QueryNetworkRttReportResponseBody {
	s.NetworkRttReportDataModel = v
	return s
}

func (s *QueryNetworkRttReportResponseBody) SetPageNumber(v int32) *QueryNetworkRttReportResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryNetworkRttReportResponseBody) SetPageSize(v int32) *QueryNetworkRttReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryNetworkRttReportResponseBody) SetRequestId(v string) *QueryNetworkRttReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryNetworkRttReportResponseBody) Validate() error {
	if s.NetworkRttReportDataModel != nil {
		for _, item := range s.NetworkRttReportDataModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryNetworkRttReportResponseBodyNetworkRttReportDataModel struct {
	// The ID of the app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The number of users that meet the low-RTT criteria of the report. Unit: users. The current API operation does not return the specific threshold.
	//
	// example:
	//
	// 85
	LowRttUserCount *int64 `json:"LowRttUserCount,omitempty" xml:"LowRttUserCount,omitempty"`
	// The ratio of low-RTT users to the total number of users involved in the statistics. Valid values: 0 to 1. To obtain the percentage, multiply the value by 100.
	//
	// example:
	//
	// 0.85
	LowRttUserRate *float64 `json:"LowRttUserRate,omitempty" xml:"LowRttUserRate,omitempty"`
	// The ID of the business region to which the report data belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The average RTT of all users involved in the statistics. Unit: milliseconds.
	//
	// example:
	//
	// 35.8
	RttAvg *float64 `json:"RttAvg,omitempty" xml:"RttAvg,omitempty"`
	// The maximum value among the average RTTs of all users. Unit: milliseconds.
	//
	// example:
	//
	// 86.2
	RttMax *float64 `json:"RttMax,omitempty" xml:"RttMax,omitempty"`
	// The minimum value among the average RTTs of all users. Unit: milliseconds.
	//
	// example:
	//
	// 12.5
	RttMin *float64 `json:"RttMin,omitempty" xml:"RttMin,omitempty"`
	// The end time of the statistical window. Format: `yyyy-MM-dd HH:mm:ss`. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-01 10:05:00
	StatEndDate *string `json:"StatEndDate,omitempty" xml:"StatEndDate,omitempty"`
	// The start time of the statistical window. Format: `yyyy-MM-dd HH:mm:ss`. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-01 10:00:00
	StatStartDate *string `json:"StatStartDate,omitempty" xml:"StatStartDate,omitempty"`
	// The total number of users involved in the RTT calculation for the current statistical window. Unit: users.
	//
	// example:
	//
	// 100
	TotalUserCount *int64 `json:"TotalUserCount,omitempty" xml:"TotalUserCount,omitempty"`
}

func (s QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) String() string {
	return dara.Prettify(s)
}

func (s QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GoString() string {
	return s.String()
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetLowRttUserCount() *int64 {
	return s.LowRttUserCount
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetLowRttUserRate() *float64 {
	return s.LowRttUserRate
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetRttAvg() *float64 {
	return s.RttAvg
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetRttMax() *float64 {
	return s.RttMax
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetRttMin() *float64 {
	return s.RttMin
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetStatEndDate() *string {
	return s.StatEndDate
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetStatStartDate() *string {
	return s.StatStartDate
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) GetTotalUserCount() *int64 {
	return s.TotalUserCount
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetAppInstanceGroupId(v string) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetLowRttUserCount(v int64) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.LowRttUserCount = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetLowRttUserRate(v float64) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.LowRttUserRate = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetRegionId(v string) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.RegionId = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetRttAvg(v float64) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.RttAvg = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetRttMax(v float64) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.RttMax = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetRttMin(v float64) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.RttMin = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetStatEndDate(v string) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.StatEndDate = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetStatStartDate(v string) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.StatStartDate = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) SetTotalUserCount(v int64) *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel {
	s.TotalUserCount = &v
	return s
}

func (s *QueryNetworkRttReportResponseBodyNetworkRttReportDataModel) Validate() error {
	return dara.Validate(s)
}
