// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityAreaDistributionStatDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeQualityAreaDistributionStatDataRequest
	GetAppId() *string
	SetEndDate(v int64) *DescribeQualityAreaDistributionStatDataRequest
	GetEndDate() *int64
	SetParentArea(v string) *DescribeQualityAreaDistributionStatDataRequest
	GetParentArea() *string
	SetStartDate(v int64) *DescribeQualityAreaDistributionStatDataRequest
	GetStartDate() *int64
}

type DescribeQualityAreaDistributionStatDataRequest struct {
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
	// example:
	//
	// 中国
	ParentArea *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeQualityAreaDistributionStatDataRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeQualityAreaDistributionStatDataRequest) GetParentArea() *string {
	return s.ParentArea
}

func (s *DescribeQualityAreaDistributionStatDataRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetAppId(v string) *DescribeQualityAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityAreaDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeQualityAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityAreaDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) Validate() error {
	return dara.Validate(s)
}
