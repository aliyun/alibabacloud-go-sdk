// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmbodiedAIPlatformsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeEmbodiedAIPlatformsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeEmbodiedAIPlatformsRequest
	GetEndTime() *string
	SetMaxResults(v int32) *DescribeEmbodiedAIPlatformsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeEmbodiedAIPlatformsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeEmbodiedAIPlatformsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEmbodiedAIPlatformsRequest
	GetPageSize() *int32
	SetPlatformName(v string) *DescribeEmbodiedAIPlatformsRequest
	GetPlatformName() *string
	SetRegionId(v string) *DescribeEmbodiedAIPlatformsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeEmbodiedAIPlatformsRequest
	GetStartTime() *string
	SetState(v string) *DescribeEmbodiedAIPlatformsRequest
	GetState() *string
}

type DescribeEmbodiedAIPlatformsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-12-01T03:05Z
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// platform1
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-12-01T02:05Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeEmbodiedAIPlatformsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmbodiedAIPlatformsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEmbodiedAIPlatformsRequest) GetState() *string {
	return s.State
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetDBClusterId(v string) *DescribeEmbodiedAIPlatformsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetEndTime(v string) *DescribeEmbodiedAIPlatformsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetMaxResults(v int32) *DescribeEmbodiedAIPlatformsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetNextToken(v string) *DescribeEmbodiedAIPlatformsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetPageNumber(v int32) *DescribeEmbodiedAIPlatformsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetPageSize(v int32) *DescribeEmbodiedAIPlatformsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetPlatformName(v string) *DescribeEmbodiedAIPlatformsRequest {
	s.PlatformName = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetRegionId(v string) *DescribeEmbodiedAIPlatformsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetStartTime(v string) *DescribeEmbodiedAIPlatformsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) SetState(v string) *DescribeEmbodiedAIPlatformsRequest {
	s.State = &v
	return s
}

func (s *DescribeEmbodiedAIPlatformsRequest) Validate() error {
	return dara.Validate(s)
}
