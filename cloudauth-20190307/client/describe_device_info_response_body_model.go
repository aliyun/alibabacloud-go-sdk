// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDeviceInfoResponseBody
	GetCurrentPage() *int32
	SetDeviceInfoList(v *DescribeDeviceInfoResponseBodyDeviceInfoList) *DescribeDeviceInfoResponseBody
	GetDeviceInfoList() *DescribeDeviceInfoResponseBodyDeviceInfoList
	SetPageSize(v int32) *DescribeDeviceInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDeviceInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDeviceInfoResponseBody
	GetTotalCount() *int32
}

type DescribeDeviceInfoResponseBody struct {
	// The current page number being queried.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Array of device information.
	DeviceInfoList *DescribeDeviceInfoResponseBodyDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Struct"`
	// Number of items per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of this request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDeviceInfoResponseBody) GetDeviceInfoList() *DescribeDeviceInfoResponseBodyDeviceInfoList {
	return s.DeviceInfoList
}

func (s *DescribeDeviceInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDeviceInfoResponseBody) SetCurrentPage(v int32) *DescribeDeviceInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetDeviceInfoList(v *DescribeDeviceInfoResponseBodyDeviceInfoList) *DescribeDeviceInfoResponseBody {
	s.DeviceInfoList = v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetPageSize(v int32) *DescribeDeviceInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetRequestId(v string) *DescribeDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) SetTotalCount(v int32) *DescribeDeviceInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeviceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceInfoResponseBodyDeviceInfoList struct {
	DeviceInfo []*DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Repeated"`
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoList) GetDeviceInfo() []*DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	return s.DeviceInfo
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoList) SetDeviceInfo(v []*DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) *DescribeDeviceInfoResponseBodyDeviceInfoList {
	s.DeviceInfo = v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo struct {
	// Authorization start date.
	//
	// example:
	//
	// 20180101
	BeginDay *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty"`
	// Corresponds to the BizType in the request.
	//
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Corresponds to the DeviceId in the request.
	//
	// example:
	//
	// wd.6ziUffspAeW5FVYbaqmexR-1qwNjM
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// Authorization expiration date.
	//
	// example:
	//
	// 20180101
	ExpiredDay *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	// Corresponds to the UserDeviceId in the request.
	//
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) GetBeginDay() *string {
	return s.BeginDay
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) GetBizType() *string {
	return s.BizType
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) GetExpiredDay() *string {
	return s.ExpiredDay
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) GetUserDeviceId() *string {
	return s.UserDeviceId
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetBeginDay(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.BeginDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetBizType(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetDeviceId(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetExpiredDay(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.ExpiredDay = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) SetUserDeviceId(v string) *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo {
	s.UserDeviceId = &v
	return s
}

func (s *DescribeDeviceInfoResponseBodyDeviceInfoListDeviceInfo) Validate() error {
	return dara.Validate(s)
}
