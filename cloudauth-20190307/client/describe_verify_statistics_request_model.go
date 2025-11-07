// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgeGt(v string) *DescribeVerifyStatisticsRequest
	GetAgeGt() *string
	SetEndDate(v int64) *DescribeVerifyStatisticsRequest
	GetEndDate() *int64
	SetProductCode(v string) *DescribeVerifyStatisticsRequest
	GetProductCode() *string
	SetServiceCode(v string) *DescribeVerifyStatisticsRequest
	GetServiceCode() *string
	SetStartDate(v int64) *DescribeVerifyStatisticsRequest
	GetStartDate() *int64
}

type DescribeVerifyStatisticsRequest struct {
	// Whether the age is over 14 years old:
	//
	// - **T**: Over
	//
	// - **F**: Under
	//
	// example:
	//
	// T
	AgeGt *string `json:"AgeGt,omitempty" xml:"AgeGt,omitempty"`
	// End date of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760630399999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Product Code.
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
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Start date of the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1760025600000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyStatisticsRequest) GetAgeGt() *string {
	return s.AgeGt
}

func (s *DescribeVerifyStatisticsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeVerifyStatisticsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifyStatisticsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeVerifyStatisticsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeVerifyStatisticsRequest) SetAgeGt(v string) *DescribeVerifyStatisticsRequest {
	s.AgeGt = &v
	return s
}

func (s *DescribeVerifyStatisticsRequest) SetEndDate(v int64) *DescribeVerifyStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyStatisticsRequest) SetProductCode(v string) *DescribeVerifyStatisticsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifyStatisticsRequest) SetServiceCode(v string) *DescribeVerifyStatisticsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeVerifyStatisticsRequest) SetStartDate(v int64) *DescribeVerifyStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifyStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
