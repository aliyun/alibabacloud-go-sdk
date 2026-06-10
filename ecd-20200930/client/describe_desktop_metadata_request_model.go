// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTimeStart(v string) *DescribeDesktopMetadataRequest
	GetCreationTimeStart() *string
	SetDesktopIds(v []*string) *DescribeDesktopMetadataRequest
	GetDesktopIds() []*string
	SetEndUserId(v string) *DescribeDesktopMetadataRequest
	GetEndUserId() *string
	SetGroupId(v string) *DescribeDesktopMetadataRequest
	GetGroupId() *string
	SetHostName(v string) *DescribeDesktopMetadataRequest
	GetHostName() *string
	SetImageId(v string) *DescribeDesktopMetadataRequest
	GetImageId() *string
	SetIncludeDesktopGroup(v bool) *DescribeDesktopMetadataRequest
	GetIncludeDesktopGroup() *bool
	SetKeyword(v string) *DescribeDesktopMetadataRequest
	GetKeyword() *string
	SetMaxResults(v int32) *DescribeDesktopMetadataRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDesktopMetadataRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeDesktopMetadataRequest
	GetOfficeSiteId() *string
	SetOperationTimeStart(v string) *DescribeDesktopMetadataRequest
	GetOperationTimeStart() *string
	SetRegionId(v string) *DescribeDesktopMetadataRequest
	GetRegionId() *string
	SetSearchRegionId(v string) *DescribeDesktopMetadataRequest
	GetSearchRegionId() *string
}

type DescribeDesktopMetadataRequest struct {
	// The creation time of the cloud computer. The time must be in the `yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\"` format and in UTC.
	//
	// example:
	//
	// 2025-01-01T12:00:00Z
	CreationTimeStart *string `json:"CreationTimeStart,omitempty" xml:"CreationTimeStart,omitempty"`
	// A list of cloud computer IDs.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The ID of the end user.
	//
	// example:
	//
	// test-user
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the cloud computer share.
	//
	// example:
	//
	// dg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The hostname.
	//
	// example:
	//
	// ASW-2F-SRV-YXYZ-4.SHPTG
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Specifies whether to include cloud computers in cloud computer shares in the response.
	//
	// example:
	//
	// false
	IncludeDesktopGroup *bool `json:"IncludeDesktopGroup,omitempty" xml:"IncludeDesktopGroup,omitempty"`
	// > This parameter is not yet available.
	//
	// example:
	//
	// ecd
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of entries to return per page. Maximum: 100. Default: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token returned from the previous call to retrieve the next page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the office network.
	//
	// example:
	//
	// cn-hangzhou+dir-778418****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The start of the time range to query for operations. The time must be in the `yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\"` format and in UTC.
	//
	// example:
	//
	// 2025-01-01T12:00:00Z
	OperationTimeStart *string `json:"OperationTimeStart,omitempty" xml:"OperationTimeStart,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the region to search.
	//
	// example:
	//
	// cn-hangzhou
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
}

func (s DescribeDesktopMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopMetadataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopMetadataRequest) GetCreationTimeStart() *string {
	return s.CreationTimeStart
}

func (s *DescribeDesktopMetadataRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *DescribeDesktopMetadataRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopMetadataRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDesktopMetadataRequest) GetHostName() *string {
	return s.HostName
}

func (s *DescribeDesktopMetadataRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDesktopMetadataRequest) GetIncludeDesktopGroup() *bool {
	return s.IncludeDesktopGroup
}

func (s *DescribeDesktopMetadataRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDesktopMetadataRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopMetadataRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopMetadataRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopMetadataRequest) GetOperationTimeStart() *string {
	return s.OperationTimeStart
}

func (s *DescribeDesktopMetadataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopMetadataRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeDesktopMetadataRequest) SetCreationTimeStart(v string) *DescribeDesktopMetadataRequest {
	s.CreationTimeStart = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetDesktopIds(v []*string) *DescribeDesktopMetadataRequest {
	s.DesktopIds = v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetEndUserId(v string) *DescribeDesktopMetadataRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetGroupId(v string) *DescribeDesktopMetadataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetHostName(v string) *DescribeDesktopMetadataRequest {
	s.HostName = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetImageId(v string) *DescribeDesktopMetadataRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetIncludeDesktopGroup(v bool) *DescribeDesktopMetadataRequest {
	s.IncludeDesktopGroup = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetKeyword(v string) *DescribeDesktopMetadataRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetMaxResults(v int32) *DescribeDesktopMetadataRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetNextToken(v string) *DescribeDesktopMetadataRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetOfficeSiteId(v string) *DescribeDesktopMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetOperationTimeStart(v string) *DescribeDesktopMetadataRequest {
	s.OperationTimeStart = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetRegionId(v string) *DescribeDesktopMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) SetSearchRegionId(v string) *DescribeDesktopMetadataRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeDesktopMetadataRequest) Validate() error {
	return dara.Validate(s)
}
