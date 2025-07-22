// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityDistributionStatDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeQualityDistributionStatDataRequest
	GetAppId() *string
	SetEndDate(v int64) *DescribeQualityDistributionStatDataRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeQualityDistributionStatDataRequest
	GetStartDate() *int64
	SetStatDim(v string) *DescribeQualityDistributionStatDataRequest
	GetStatDim() *string
}

type DescribeQualityDistributionStatDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CHANNEL_ONLINE
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeQualityDistributionStatDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeQualityDistributionStatDataRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeQualityDistributionStatDataRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeQualityDistributionStatDataRequest) GetStatDim() *string {
	return s.StatDim
}

func (s *DescribeQualityDistributionStatDataRequest) SetAppId(v string) *DescribeQualityDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetStatDim(v string) *DescribeQualityDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) Validate() error {
	return dara.Validate(s)
}
