// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeDeviceInfoRequest
	GetBizType() *string
	SetCurrentPage(v int32) *DescribeDeviceInfoRequest
	GetCurrentPage() *int32
	SetDeviceId(v string) *DescribeDeviceInfoRequest
	GetDeviceId() *string
	SetExpiredEndDay(v string) *DescribeDeviceInfoRequest
	GetExpiredEndDay() *string
	SetExpiredStartDay(v string) *DescribeDeviceInfoRequest
	GetExpiredStartDay() *string
	SetPageSize(v int32) *DescribeDeviceInfoRequest
	GetPageSize() *int32
	SetUserDeviceId(v string) *DescribeDeviceInfoRequest
	GetUserDeviceId() *string
}

type DescribeDeviceInfoRequest struct {
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// wd.6ziUffspAeW5FVYbaqmexR-1qwNjM
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// 20200330
	ExpiredEndDay *string `json:"ExpiredEndDay,omitempty" xml:"ExpiredEndDay,omitempty"`
	// example:
	//
	// 20190401
	ExpiredStartDay *string `json:"ExpiredStartDay,omitempty" xml:"ExpiredStartDay,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s DescribeDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceInfoRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeDeviceInfoRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDeviceInfoRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DescribeDeviceInfoRequest) GetExpiredEndDay() *string {
	return s.ExpiredEndDay
}

func (s *DescribeDeviceInfoRequest) GetExpiredStartDay() *string {
	return s.ExpiredStartDay
}

func (s *DescribeDeviceInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeviceInfoRequest) GetUserDeviceId() *string {
	return s.UserDeviceId
}

func (s *DescribeDeviceInfoRequest) SetBizType(v string) *DescribeDeviceInfoRequest {
	s.BizType = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetCurrentPage(v int32) *DescribeDeviceInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetDeviceId(v string) *DescribeDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredEndDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredEndDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetExpiredStartDay(v string) *DescribeDeviceInfoRequest {
	s.ExpiredStartDay = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetPageSize(v int32) *DescribeDeviceInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceInfoRequest) SetUserDeviceId(v string) *DescribeDeviceInfoRequest {
	s.UserDeviceId = &v
	return s
}

func (s *DescribeDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
