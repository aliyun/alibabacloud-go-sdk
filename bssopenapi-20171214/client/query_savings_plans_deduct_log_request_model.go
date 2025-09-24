// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansDeductLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *QuerySavingsPlansDeductLogRequest
	GetEndTime() *string
	SetInstanceId(v string) *QuerySavingsPlansDeductLogRequest
	GetInstanceId() *string
	SetInstanceType(v string) *QuerySavingsPlansDeductLogRequest
	GetInstanceType() *string
	SetLocale(v string) *QuerySavingsPlansDeductLogRequest
	GetLocale() *string
	SetPageNum(v int32) *QuerySavingsPlansDeductLogRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QuerySavingsPlansDeductLogRequest
	GetPageSize() *int32
	SetStartTime(v string) *QuerySavingsPlansDeductLogRequest
	GetStartTime() *string
}

type QuerySavingsPlansDeductLogRequest struct {
	// The end of the time range to query. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2022-01-05 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// spn-XXXXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance ID based on which the data is queried. Valid values:
	//
	// 	- spn: queries data based on the ID of the savings plan instance.
	//
	// 	- product: queries data based on the ID of the cloud service instance.
	//
	// example:
	//
	// spn
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The language of the return data. Valid values:
	//
	// 	- ZH: Chinese
	//
	// 	- EN: English
	//
	// example:
	//
	// ZH
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2022-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QuerySavingsPlansDeductLogRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDeductLogRequest) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySavingsPlansDeductLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QuerySavingsPlansDeductLogRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *QuerySavingsPlansDeductLogRequest) GetLocale() *string {
	return s.Locale
}

func (s *QuerySavingsPlansDeductLogRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySavingsPlansDeductLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySavingsPlansDeductLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySavingsPlansDeductLogRequest) SetEndTime(v string) *QuerySavingsPlansDeductLogRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetInstanceId(v string) *QuerySavingsPlansDeductLogRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetInstanceType(v string) *QuerySavingsPlansDeductLogRequest {
	s.InstanceType = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetLocale(v string) *QuerySavingsPlansDeductLogRequest {
	s.Locale = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetPageNum(v int32) *QuerySavingsPlansDeductLogRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetPageSize(v int32) *QuerySavingsPlansDeductLogRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetStartTime(v string) *QuerySavingsPlansDeductLogRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) Validate() error {
	return dara.Validate(s)
}
