// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQualityOsSdkVersionDistributionStatDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeQualityOsSdkVersionDistributionStatDataRequest
	GetAppId() *string
	SetEndDate(v int64) *DescribeQualityOsSdkVersionDistributionStatDataRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeQualityOsSdkVersionDistributionStatDataRequest
	GetStartDate() *int64
}

type DescribeQualityOsSdkVersionDistributionStatDataRequest struct {
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
}

func (s DescribeQualityOsSdkVersionDistributionStatDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetAppId(v string) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) Validate() error {
	return dara.Validate(s)
}
