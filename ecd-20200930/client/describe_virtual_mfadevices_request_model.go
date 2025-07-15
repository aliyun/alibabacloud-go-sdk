// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualMFADevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v []*string) *DescribeVirtualMFADevicesRequest
	GetEndUserId() []*string
	SetMaxResults(v int32) *DescribeVirtualMFADevicesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVirtualMFADevicesRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeVirtualMFADevicesRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeVirtualMFADevicesRequest
	GetRegionId() *string
}

type DescribeVirtualMFADevicesRequest struct {
	// The names of the AD users.
	//
	// example:
	//
	// testuser
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 500. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou+dir-269345****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeVirtualMFADevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualMFADevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *DescribeVirtualMFADevicesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVirtualMFADevicesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVirtualMFADevicesRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeVirtualMFADevicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVirtualMFADevicesRequest) SetEndUserId(v []*string) *DescribeVirtualMFADevicesRequest {
	s.EndUserId = v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetMaxResults(v int32) *DescribeVirtualMFADevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetNextToken(v string) *DescribeVirtualMFADevicesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetOfficeSiteId(v string) *DescribeVirtualMFADevicesRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) SetRegionId(v string) *DescribeVirtualMFADevicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualMFADevicesRequest) Validate() error {
	return dara.Validate(s)
}
