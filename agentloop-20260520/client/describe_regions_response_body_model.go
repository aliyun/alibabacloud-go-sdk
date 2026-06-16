// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeRegionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRegionsResponseBody
	GetNextToken() *string
	SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() []*DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	MaxResults *int32                                `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Regions    []*DescribeRegionsResponseBodyRegions `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRegionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRegionsResponseBody) GetRegions() []*DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetMaxResults(v int32) *DescribeRegionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetNextToken(v string) *DescribeRegionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegions struct {
	InternetEndpoint *string `json:"internetEndpoint,omitempty" xml:"internetEndpoint,omitempty"`
	LocalName        *string `json:"localName,omitempty" xml:"localName,omitempty"`
	RegionId         *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	VpcEndpoint      *string `json:"vpcEndpoint,omitempty" xml:"vpcEndpoint,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetInternetEndpoint() *string {
	return s.InternetEndpoint
}

func (s *DescribeRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegions) GetVpcEndpoint() *string {
	return s.VpcEndpoint
}

func (s *DescribeRegionsResponseBodyRegions) SetInternetEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetVpcEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.VpcEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
