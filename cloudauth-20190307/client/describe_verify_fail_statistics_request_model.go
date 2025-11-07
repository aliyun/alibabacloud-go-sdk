// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyFailStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgeGt(v string) *DescribeVerifyFailStatisticsRequest
	GetAgeGt() *string
	SetApi(v string) *DescribeVerifyFailStatisticsRequest
	GetApi() *string
	SetDeviceType(v string) *DescribeVerifyFailStatisticsRequest
	GetDeviceType() *string
	SetEndDate(v int64) *DescribeVerifyFailStatisticsRequest
	GetEndDate() *int64
	SetProductCode(v string) *DescribeVerifyFailStatisticsRequest
	GetProductCode() *string
	SetServiceCode(v string) *DescribeVerifyFailStatisticsRequest
	GetServiceCode() *string
	SetStartDate(v int64) *DescribeVerifyFailStatisticsRequest
	GetStartDate() *int64
}

type DescribeVerifyFailStatisticsRequest struct {
	// Age greater than 14 years old:
	//
	// - **T**: Greater than
	//
	// - **F**: Less than
	//
	// example:
	//
	// T
	AgeGt *string `json:"AgeGt,omitempty" xml:"AgeGt,omitempty"`
	// API code:
	//
	// - **INIT_SERVICE**: Server-side initialization failure
	//
	// - **INIT_DEVICE**: Client-side failure
	//
	// - **VERIFY_DEVICE**: Authentication failed
	//
	// This parameter is required.
	//
	// example:
	//
	// INIT_SERVICE
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// Device type.
	//
	// - ios
	//
	// - android
	//
	// - websdk
	//
	// example:
	//
	// ios
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// End time of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760630399999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Service type:
	//
	// - **antcloudauth**: Financial-grade real-person authentication.
	//
	// - **cloudauthst*	- (discontinued): Enhanced real-person authentication.
	//
	// - **cloudauth*	- (discontinued): Real-person authentication.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudauthst
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Start time of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760025600000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyFailStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyFailStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyFailStatisticsRequest) GetAgeGt() *string {
	return s.AgeGt
}

func (s *DescribeVerifyFailStatisticsRequest) GetApi() *string {
	return s.Api
}

func (s *DescribeVerifyFailStatisticsRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeVerifyFailStatisticsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeVerifyFailStatisticsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifyFailStatisticsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeVerifyFailStatisticsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeVerifyFailStatisticsRequest) SetAgeGt(v string) *DescribeVerifyFailStatisticsRequest {
	s.AgeGt = &v
	return s
}

func (s *DescribeVerifyFailStatisticsRequest) SetApi(v string) *DescribeVerifyFailStatisticsRequest {
	s.Api = &v
	return s
}

func (s *DescribeVerifyFailStatisticsRequest) SetDeviceType(v string) *DescribeVerifyFailStatisticsRequest {
	s.DeviceType = &v
	return s
}

func (s *DescribeVerifyFailStatisticsRequest) SetEndDate(v int64) *DescribeVerifyFailStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyFailStatisticsRequest) SetProductCode(v string) *DescribeVerifyFailStatisticsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifyFailStatisticsRequest) SetServiceCode(v string) *DescribeVerifyFailStatisticsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeVerifyFailStatisticsRequest) SetStartDate(v int64) *DescribeVerifyFailStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifyFailStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
