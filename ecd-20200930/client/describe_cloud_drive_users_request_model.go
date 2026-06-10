// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDriveUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DescribeCloudDriveUsersRequest
	GetCdsId() *string
	SetEndUserId(v string) *DescribeCloudDriveUsersRequest
	GetEndUserId() *string
	SetMaxResults(v int32) *DescribeCloudDriveUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCloudDriveUsersRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeCloudDriveUsersRequest
	GetRegionId() *string
}

type DescribeCloudDriveUsersRequest struct {
	// The ID of the enterprise network drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-066224****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The end user ID.
	//
	// If specified, the query returns only the personal drive of that user. If left empty, the query returns the personal drives of all users.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The number of entries per page.
	//
	// - Maximum: 100.
	//
	// - Default: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results. This is the `NextToken` value from the previous response. Leave this parameter empty for the first request.
	//
	// example:
	//
	// aGN4YzAxQGNuLWhhbmd6aG91LjExNzU5NTMyNjgzMTQ1****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. To obtain a list of supported regions, call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudDriveUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveUsersRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DescribeCloudDriveUsersRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeCloudDriveUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCloudDriveUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudDriveUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudDriveUsersRequest) SetCdsId(v string) *DescribeCloudDriveUsersRequest {
	s.CdsId = &v
	return s
}

func (s *DescribeCloudDriveUsersRequest) SetEndUserId(v string) *DescribeCloudDriveUsersRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeCloudDriveUsersRequest) SetMaxResults(v int32) *DescribeCloudDriveUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudDriveUsersRequest) SetNextToken(v string) *DescribeCloudDriveUsersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudDriveUsersRequest) SetRegionId(v string) *DescribeCloudDriveUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudDriveUsersRequest) Validate() error {
	return dara.Validate(s)
}
