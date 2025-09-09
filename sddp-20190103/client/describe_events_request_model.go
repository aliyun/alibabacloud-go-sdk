// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeEventsRequest
	GetCurrentPage() *int32
	SetDealUserId(v string) *DescribeEventsRequest
	GetDealUserId() *string
	SetEndTime(v string) *DescribeEventsRequest
	GetEndTime() *string
	SetId(v int64) *DescribeEventsRequest
	GetId() *int64
	SetInstanceName(v string) *DescribeEventsRequest
	GetInstanceName() *string
	SetLang(v string) *DescribeEventsRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeEventsRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeEventsRequest
	GetProductCode() *string
	SetStartTime(v string) *DescribeEventsRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeEventsRequest
	GetStatus() *string
	SetSubTypeCode(v string) *DescribeEventsRequest
	GetSubTypeCode() *string
	SetTargetProductCode(v string) *DescribeEventsRequest
	GetTargetProductCode() *string
	SetTypeCode(v string) *DescribeEventsRequest
	GetTypeCode() *string
	SetUserId(v int64) *DescribeEventsRequest
	GetUserId() *int64
	SetUserName(v string) *DescribeEventsRequest
	GetUserName() *string
	SetWarnLevel(v int32) *DescribeEventsRequest
	GetWarnLevel() *int32
}

type DescribeEventsRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the account that handles the anomalous event.
	//
	// example:
	//
	// yundun-***
	DealUserId *string `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	// The end of the time range during which the anomalous events are detected. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1698700000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The unique ID of the anomalous event.
	//
	// example:
	//
	// 789026
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data asset.
	//
	// example:
	//
	// rm-uf6yzvbc2tg90iuxk.l****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service to which the table belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// OSS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The beginning of the time range during which the anomalous events are detected. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1657900000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The handling status of the anomalous event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: confirmed
	//
	// 	- 2: marked as false positive
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the anomalous event subtype.
	//
	// > You can call the **DescribeEventTypes*	- operation to query the name of the anomalous event subtype.
	//
	// example:
	//
	// Anomalous volume of downloaded data
	SubTypeCode *string `json:"SubTypeCode,omitempty" xml:"SubTypeCode,omitempty"`
	// The name of the destination service in an anomalous data flow. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**
	//
	// example:
	//
	// RDS
	TargetProductCode *string `json:"TargetProductCode,omitempty" xml:"TargetProductCode,omitempty"`
	// The name of the anomalous event type. Valid values:
	//
	// 	- 01: anomalous permission usage
	//
	// 	- 02: anomalous data flow
	//
	// 	- 03: anomalous data operation
	//
	// example:
	//
	// 02
	TypeCode *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	// The ID of the account that triggered the anomalous event.
	//
	// example:
	//
	// 1978132506596***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the RAM user.
	//
	// example:
	//
	// name
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The risk level of the alert that is triggered. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 1
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEventsRequest) GetDealUserId() *string {
	return s.DealUserId
}

func (s *DescribeEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEventsRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEventsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEventsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeEventsRequest) GetSubTypeCode() *string {
	return s.SubTypeCode
}

func (s *DescribeEventsRequest) GetTargetProductCode() *string {
	return s.TargetProductCode
}

func (s *DescribeEventsRequest) GetTypeCode() *string {
	return s.TypeCode
}

func (s *DescribeEventsRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeEventsRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeEventsRequest) GetWarnLevel() *int32 {
	return s.WarnLevel
}

func (s *DescribeEventsRequest) SetCurrentPage(v int32) *DescribeEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventsRequest) SetDealUserId(v string) *DescribeEventsRequest {
	s.DealUserId = &v
	return s
}

func (s *DescribeEventsRequest) SetEndTime(v string) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetId(v int64) *DescribeEventsRequest {
	s.Id = &v
	return s
}

func (s *DescribeEventsRequest) SetInstanceName(v string) *DescribeEventsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeEventsRequest) SetLang(v string) *DescribeEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int32) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetProductCode(v string) *DescribeEventsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v string) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsRequest) SetStatus(v string) *DescribeEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeEventsRequest) SetSubTypeCode(v string) *DescribeEventsRequest {
	s.SubTypeCode = &v
	return s
}

func (s *DescribeEventsRequest) SetTargetProductCode(v string) *DescribeEventsRequest {
	s.TargetProductCode = &v
	return s
}

func (s *DescribeEventsRequest) SetTypeCode(v string) *DescribeEventsRequest {
	s.TypeCode = &v
	return s
}

func (s *DescribeEventsRequest) SetUserId(v int64) *DescribeEventsRequest {
	s.UserId = &v
	return s
}

func (s *DescribeEventsRequest) SetUserName(v string) *DescribeEventsRequest {
	s.UserName = &v
	return s
}

func (s *DescribeEventsRequest) SetWarnLevel(v int32) *DescribeEventsRequest {
	s.WarnLevel = &v
	return s
}

func (s *DescribeEventsRequest) Validate() error {
	return dara.Validate(s)
}
