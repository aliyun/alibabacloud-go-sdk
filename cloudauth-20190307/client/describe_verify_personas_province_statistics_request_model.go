// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasProvinceStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *DescribeVerifyPersonasProvinceStatisticsRequest
	GetProductCode() *string
	SetSceneId(v int64) *DescribeVerifyPersonasProvinceStatisticsRequest
	GetSceneId() *int64
	SetServiceCode(v string) *DescribeVerifyPersonasProvinceStatisticsRequest
	GetServiceCode() *string
	SetTimeRange(v string) *DescribeVerifyPersonasProvinceStatisticsRequest
	GetTimeRange() *string
}

type DescribeVerifyPersonasProvinceStatisticsRequest struct {
	// Cloud product code.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000004071
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Service type:
	//
	// - **antcloudauth**: Financial-grade real-person authentication.
	//
	// - **cloudauthst*	- (discontinued): Enhanced real-person authentication.
	//
	// This parameter is required.
	//
	// example:
	//
	// antcloudauth
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// Time range, search range is for the previous N days, TimeRange of 1 indicates the previous 1 day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s DescribeVerifyPersonasProvinceStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasProvinceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) SetProductCode(v string) *DescribeVerifyPersonasProvinceStatisticsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) SetSceneId(v int64) *DescribeVerifyPersonasProvinceStatisticsRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) SetServiceCode(v string) *DescribeVerifyPersonasProvinceStatisticsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) SetTimeRange(v string) *DescribeVerifyPersonasProvinceStatisticsRequest {
	s.TimeRange = &v
	return s
}

func (s *DescribeVerifyPersonasProvinceStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
