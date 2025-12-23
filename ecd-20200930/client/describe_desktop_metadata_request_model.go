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
	// example:
	//
	// 2025-01-01T12:00:00Z
	CreationTimeStart *string   `json:"CreationTimeStart,omitempty" xml:"CreationTimeStart,omitempty"`
	DesktopIds        []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// example:
	//
	// dg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// ASW-2F-SRV-YXYZ-4.SHPTG
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// false
	IncludeDesktopGroup *bool `json:"IncludeDesktopGroup,omitempty" xml:"IncludeDesktopGroup,omitempty"`
	// example:
	//
	// ecd
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-778418****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// 2025-01-01T12:00:00Z
	OperationTimeStart *string `json:"OperationTimeStart,omitempty" xml:"OperationTimeStart,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
