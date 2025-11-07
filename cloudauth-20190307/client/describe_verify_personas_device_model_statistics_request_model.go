// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasDeviceModelStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *DescribeVerifyPersonasDeviceModelStatisticsRequest
	GetProductCode() *string
	SetSceneId(v int64) *DescribeVerifyPersonasDeviceModelStatisticsRequest
	GetSceneId() *int64
	SetServiceCode(v string) *DescribeVerifyPersonasDeviceModelStatisticsRequest
	GetServiceCode() *string
	SetTimeRange(v string) *DescribeVerifyPersonasDeviceModelStatisticsRequest
	GetTimeRange() *string
}

type DescribeVerifyPersonasDeviceModelStatisticsRequest struct {
	// Product Code.
	//
	// example:
	//
	// ID_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Scene ID.
	//
	// example:
	//
	// 1000015316
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
	// Time range for the query, indicating how many days back.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s DescribeVerifyPersonasDeviceModelStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasDeviceModelStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) SetProductCode(v string) *DescribeVerifyPersonasDeviceModelStatisticsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) SetSceneId(v int64) *DescribeVerifyPersonasDeviceModelStatisticsRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) SetServiceCode(v string) *DescribeVerifyPersonasDeviceModelStatisticsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) SetTimeRange(v string) *DescribeVerifyPersonasDeviceModelStatisticsRequest {
	s.TimeRange = &v
	return s
}

func (s *DescribeVerifyPersonasDeviceModelStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
