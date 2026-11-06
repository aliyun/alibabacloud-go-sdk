// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDauReportDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryDauReportDataResponseBody
	GetCode() *string
	SetCount(v int32) *QueryDauReportDataResponseBody
	GetCount() *int32
	SetDauReportDataModel(v []*QueryDauReportDataResponseBodyDauReportDataModel) *QueryDauReportDataResponseBody
	GetDauReportDataModel() []*QueryDauReportDataResponseBodyDauReportDataModel
	SetMessage(v string) *QueryDauReportDataResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *QueryDauReportDataResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryDauReportDataResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryDauReportDataResponseBody
	GetRequestId() *string
}

type QueryDauReportDataResponseBody struct {
	// The business status code. A value of 200 indicates that the request is processed successfully. Other values indicate that the processing failed.
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
	// The list of daily active user report data.
	DauReportDataModel []*QueryDauReportDataResponseBodyDauReportDataModel `json:"DauReportDataModel,omitempty" xml:"DauReportDataModel,omitempty" type:"Repeated"`
	// The description of the request processing result. Typically, success is returned if the request is successful.
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
	// The number of records per page used in the current response.
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

func (s QueryDauReportDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDauReportDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDauReportDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDauReportDataResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QueryDauReportDataResponseBody) GetDauReportDataModel() []*QueryDauReportDataResponseBodyDauReportDataModel {
	return s.DauReportDataModel
}

func (s *QueryDauReportDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDauReportDataResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryDauReportDataResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDauReportDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDauReportDataResponseBody) SetCode(v string) *QueryDauReportDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDauReportDataResponseBody) SetCount(v int32) *QueryDauReportDataResponseBody {
	s.Count = &v
	return s
}

func (s *QueryDauReportDataResponseBody) SetDauReportDataModel(v []*QueryDauReportDataResponseBodyDauReportDataModel) *QueryDauReportDataResponseBody {
	s.DauReportDataModel = v
	return s
}

func (s *QueryDauReportDataResponseBody) SetMessage(v string) *QueryDauReportDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDauReportDataResponseBody) SetPageNumber(v int32) *QueryDauReportDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryDauReportDataResponseBody) SetPageSize(v int32) *QueryDauReportDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDauReportDataResponseBody) SetRequestId(v string) *QueryDauReportDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDauReportDataResponseBody) Validate() error {
	if s.DauReportDataModel != nil {
		for _, item := range s.DauReportDataModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDauReportDataResponseBodyDauReportDataModel struct {
	// The number of active users within the statistical epoch. Unit: users.
	//
	// example:
	//
	// 128
	Dau *int64 `json:"Dau,omitempty" xml:"Dau,omitempty"`
	// The statistics date. Format: yyyy-MM-dd. The string does not contain a time zone and is interpreted as a natural date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
	// The number of active users who generated sessions through streaming access within the statistical epoch. Unit: users.
	//
	// example:
	//
	// 96
	StreamingDau *int64 `json:"StreamingDau,omitempty" xml:"StreamingDau,omitempty"`
}

func (s QueryDauReportDataResponseBodyDauReportDataModel) String() string {
	return dara.Prettify(s)
}

func (s QueryDauReportDataResponseBodyDauReportDataModel) GoString() string {
	return s.String()
}

func (s *QueryDauReportDataResponseBodyDauReportDataModel) GetDau() *int64 {
	return s.Dau
}

func (s *QueryDauReportDataResponseBodyDauReportDataModel) GetStatDate() *string {
	return s.StatDate
}

func (s *QueryDauReportDataResponseBodyDauReportDataModel) GetStreamingDau() *int64 {
	return s.StreamingDau
}

func (s *QueryDauReportDataResponseBodyDauReportDataModel) SetDau(v int64) *QueryDauReportDataResponseBodyDauReportDataModel {
	s.Dau = &v
	return s
}

func (s *QueryDauReportDataResponseBodyDauReportDataModel) SetStatDate(v string) *QueryDauReportDataResponseBodyDauReportDataModel {
	s.StatDate = &v
	return s
}

func (s *QueryDauReportDataResponseBodyDauReportDataModel) SetStreamingDau(v int64) *QueryDauReportDataResponseBodyDauReportDataModel {
	s.StreamingDau = &v
	return s
}

func (s *QueryDauReportDataResponseBodyDauReportDataModel) Validate() error {
	return dara.Validate(s)
}
