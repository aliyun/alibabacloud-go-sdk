// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyPersonasSexStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductCode(v string) *DescribeVerifyPersonasSexStatisticsRequest
	GetProductCode() *string
	SetSceneId(v int64) *DescribeVerifyPersonasSexStatisticsRequest
	GetSceneId() *int64
	SetServiceCode(v string) *DescribeVerifyPersonasSexStatisticsRequest
	GetServiceCode() *string
	SetTimeRange(v string) *DescribeVerifyPersonasSexStatisticsRequest
	GetTimeRange() *string
}

type DescribeVerifyPersonasSexStatisticsRequest struct {
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
	// 1000002996
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
	// Time range, the search range is for the previous N days, TimeRange of 1 indicates the previous 1 day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s DescribeVerifyPersonasSexStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyPersonasSexStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) SetProductCode(v string) *DescribeVerifyPersonasSexStatisticsRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) SetSceneId(v int64) *DescribeVerifyPersonasSexStatisticsRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) SetServiceCode(v string) *DescribeVerifyPersonasSexStatisticsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) SetTimeRange(v string) *DescribeVerifyPersonasSexStatisticsRequest {
	s.TimeRange = &v
	return s
}

func (s *DescribeVerifyPersonasSexStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
