// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyDeviceRiskStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *DescribeVerifyDeviceRiskStatisticsRequest
	GetEndDate() *int64
	SetProductCode(v string) *DescribeVerifyDeviceRiskStatisticsRequest
	GetProductCode() *string
	SetSceneId(v string) *DescribeVerifyDeviceRiskStatisticsRequest
	GetSceneId() *string
	SetServiceCode(v string) *DescribeVerifyDeviceRiskStatisticsRequest
	GetServiceCode() *string
	SetStartDate(v int64) *DescribeVerifyDeviceRiskStatisticsRequest
	GetStartDate() *int64
}

type DescribeVerifyDeviceRiskStatisticsRequest struct {
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1748624399999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
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
	// 100000xxxx
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
	// Start time of the query, in Unix timestamp format, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1746720000000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeVerifyDeviceRiskStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyDeviceRiskStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) SetEndDate(v int64) *DescribeVerifyDeviceRiskStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) SetProductCode(v string) *DescribeVerifyDeviceRiskStatisticsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) SetSceneId(v string) *DescribeVerifyDeviceRiskStatisticsRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) SetServiceCode(v string) *DescribeVerifyDeviceRiskStatisticsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) SetStartDate(v int64) *DescribeVerifyDeviceRiskStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeVerifyDeviceRiskStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
