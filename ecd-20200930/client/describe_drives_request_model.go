// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrivesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainIds(v []*string) *DescribeDrivesRequest
	GetDomainIds() []*string
	SetMaxResults(v int32) *DescribeDrivesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDrivesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDrivesRequest
	GetRegionId() *string
	SetResourceType(v string) *DescribeDrivesRequest
	GetResourceType() *string
	SetUserId(v string) *DescribeDrivesRequest
	GetUserId() *string
}

type DescribeDrivesRequest struct {
	DomainIds []*string `json:"DomainIds,omitempty" xml:"DomainIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAA****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// NAS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeDrivesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrivesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrivesRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *DescribeDrivesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDrivesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDrivesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrivesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDrivesRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeDrivesRequest) SetDomainIds(v []*string) *DescribeDrivesRequest {
	s.DomainIds = v
	return s
}

func (s *DescribeDrivesRequest) SetMaxResults(v int32) *DescribeDrivesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDrivesRequest) SetNextToken(v string) *DescribeDrivesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDrivesRequest) SetRegionId(v string) *DescribeDrivesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrivesRequest) SetResourceType(v string) *DescribeDrivesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDrivesRequest) SetUserId(v string) *DescribeDrivesRequest {
	s.UserId = &v
	return s
}

func (s *DescribeDrivesRequest) Validate() error {
	return dara.Validate(s)
}
