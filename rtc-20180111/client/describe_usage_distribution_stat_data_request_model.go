// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageDistributionStatDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeUsageDistributionStatDataRequest
	GetAppId() *string
	SetEndDate(v int64) *DescribeUsageDistributionStatDataRequest
	GetEndDate() *int64
	SetStartDate(v int64) *DescribeUsageDistributionStatDataRequest
	GetStartDate() *int64
	SetStatDim(v string) *DescribeUsageDistributionStatDataRequest
	GetStatDim() *string
}

type DescribeUsageDistributionStatDataRequest struct {
	// The app ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The statistical dimension. Valid values:
	//
	// - **CHANNEL_ONLINE**: aggregates data by the number of online users in a channel.
	//
	// - **NETWORK**: aggregates data by network type.
	//
	// - **OS**: aggregates data by operating system type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CHANNEL_ONLINE
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeUsageDistributionStatDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeUsageDistributionStatDataRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeUsageDistributionStatDataRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeUsageDistributionStatDataRequest) GetStatDim() *string {
	return s.StatDim
}

func (s *DescribeUsageDistributionStatDataRequest) SetAppId(v string) *DescribeUsageDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetEndDate(v int64) *DescribeUsageDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetStartDate(v int64) *DescribeUsageDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetStatDim(v string) *DescribeUsageDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) Validate() error {
	return dara.Validate(s)
}
