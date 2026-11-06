// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMauReportDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMauReportDataResponseBody
	GetCode() *string
	SetCount(v int32) *QueryMauReportDataResponseBody
	GetCount() *int32
	SetMauReportDataModel(v []*QueryMauReportDataResponseBodyMauReportDataModel) *QueryMauReportDataResponseBody
	GetMauReportDataModel() []*QueryMauReportDataResponseBodyMauReportDataModel
	SetMessage(v string) *QueryMauReportDataResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *QueryMauReportDataResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryMauReportDataResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryMauReportDataResponseBody
	GetRequestId() *string
}

type QueryMauReportDataResponseBody struct {
	// The business status code. A value of `200` indicates that the request is processed successfully. Other values indicate that the request fails to be processed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries that meet the query conditions.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of monthly active user report data.
	MauReportDataModel []*QueryMauReportDataResponseBodyMauReportDataModel `json:"MauReportDataModel,omitempty" xml:"MauReportDataModel,omitempty" type:"Repeated"`
	// The description of the request processing result. If the request is successful, `success` is usually returned.
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
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMauReportDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMauReportDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMauReportDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMauReportDataResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QueryMauReportDataResponseBody) GetMauReportDataModel() []*QueryMauReportDataResponseBodyMauReportDataModel {
	return s.MauReportDataModel
}

func (s *QueryMauReportDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMauReportDataResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryMauReportDataResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMauReportDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMauReportDataResponseBody) SetCode(v string) *QueryMauReportDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMauReportDataResponseBody) SetCount(v int32) *QueryMauReportDataResponseBody {
	s.Count = &v
	return s
}

func (s *QueryMauReportDataResponseBody) SetMauReportDataModel(v []*QueryMauReportDataResponseBodyMauReportDataModel) *QueryMauReportDataResponseBody {
	s.MauReportDataModel = v
	return s
}

func (s *QueryMauReportDataResponseBody) SetMessage(v string) *QueryMauReportDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMauReportDataResponseBody) SetPageNumber(v int32) *QueryMauReportDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryMauReportDataResponseBody) SetPageSize(v int32) *QueryMauReportDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMauReportDataResponseBody) SetRequestId(v string) *QueryMauReportDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMauReportDataResponseBody) Validate() error {
	if s.MauReportDataModel != nil {
		for _, item := range s.MauReportDataModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMauReportDataResponseBodyMauReportDataModel struct {
	// The app instance group ID.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The number of monthly active users in the statistical month. Unit: users.
	//
	// example:
	//
	// 286
	Mau *int64 `json:"Mau,omitempty" xml:"Mau,omitempty"`
	// The business region ID to which the report data belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The statistical cutoff date used for the month. Format: `yyyy-MM-dd`. For the last month, this is usually the day before the requested `EndDate`. The date is interpreted as a calendar date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-29
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
	// The statistical month. Format: `yyyy-MM`. The month is interpreted as a calendar month in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09
	StatMonth *string `json:"StatMonth,omitempty" xml:"StatMonth,omitempty"`
}

func (s QueryMauReportDataResponseBodyMauReportDataModel) String() string {
	return dara.Prettify(s)
}

func (s QueryMauReportDataResponseBodyMauReportDataModel) GoString() string {
	return s.String()
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) GetMau() *int64 {
	return s.Mau
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) GetStatDate() *string {
	return s.StatDate
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) GetStatMonth() *string {
	return s.StatMonth
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) SetAppInstanceGroupId(v string) *QueryMauReportDataResponseBodyMauReportDataModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) SetMau(v int64) *QueryMauReportDataResponseBodyMauReportDataModel {
	s.Mau = &v
	return s
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) SetRegionId(v string) *QueryMauReportDataResponseBodyMauReportDataModel {
	s.RegionId = &v
	return s
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) SetStatDate(v string) *QueryMauReportDataResponseBodyMauReportDataModel {
	s.StatDate = &v
	return s
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) SetStatMonth(v string) *QueryMauReportDataResponseBodyMauReportDataModel {
	s.StatMonth = &v
	return s
}

func (s *QueryMauReportDataResponseBodyMauReportDataModel) Validate() error {
	return dara.Validate(s)
}
