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
	// Specifies whether the age is greater than 14. Valid values:
	//
	// - **T**: greater than 14.
	//
	// - **F**: less than 14.
	//
	// example:
	//
	// T
	AgeGt *string `json:"AgeGt,omitempty" xml:"AgeGt,omitempty"`
	// The API code. Valid values:
	//
	// - **INIT_SERVICE**: server-side initialization failure.
	//
	// - **INIT_DEVICE**: client-side failure.
	//
	// - **VERIFY_DEVICE**: authentication not passed.
	//
	// This parameter is required.
	//
	// example:
	//
	// INIT_SERVICE
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// The device type. Valid values:
	//
	// - ios
	//
	// - android
	//
	// - websdk.
	//
	// example:
	//
	// ios
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The end time of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760630399999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The service type. Valid values:
	//
	// - **antcloudauth**: financial-grade ID Verification.
	//
	// - **cloudauthst*	- (discontinued): ID Verification Enhanced Edition.
	//
	// - **cloudauth*	- (discontinued): ID Verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloudauthst
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The start time of the query.
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
