// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasOsStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *DescribeVerifyPersonasOsStatisticsRequest
	GetProductCode() *string
	SetSceneId(v int64) *DescribeVerifyPersonasOsStatisticsRequest
	GetSceneId() *int64
	SetServiceCode(v string) *DescribeVerifyPersonasOsStatisticsRequest
	GetServiceCode() *string
	SetTimeRange(v string) *DescribeVerifyPersonasOsStatisticsRequest
	GetTimeRange() *string
}

type DescribeVerifyPersonasOsStatisticsRequest struct {
	// Product code.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000002995
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
	// Time range for the query, indicating how many days ago.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s DescribeVerifyPersonasOsStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasOsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) SetProductCode(v string) *DescribeVerifyPersonasOsStatisticsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) SetSceneId(v int64) *DescribeVerifyPersonasOsStatisticsRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) SetServiceCode(v string) *DescribeVerifyPersonasOsStatisticsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) SetTimeRange(v string) *DescribeVerifyPersonasOsStatisticsRequest {
	s.TimeRange = &v
	return s
}

func (s *DescribeVerifyPersonasOsStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
