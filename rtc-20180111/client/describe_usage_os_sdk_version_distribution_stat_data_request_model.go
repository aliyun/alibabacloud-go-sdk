// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageOsSdkVersionDistributionStatDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeUsageOsSdkVersionDistributionStatDataRequest
	GetAppId() *string
	SetEndDate(v int64) *DescribeUsageOsSdkVersionDistributionStatDataRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeUsageOsSdkVersionDistributionStatDataRequest
	GetStartDate() *int64
}

type DescribeUsageOsSdkVersionDistributionStatDataRequest struct {
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

func (s DescribeUsageOsSdkVersionDistributionStatDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetAppId(v string) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetEndDate(v int64) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetStartDate(v int64) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) Validate() error {
	return dara.Validate(s)
}
