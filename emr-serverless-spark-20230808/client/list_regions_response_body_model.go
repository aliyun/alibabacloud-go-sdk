// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRegionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListRegionsResponseBody
	GetNextToken() *string
	SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody
	GetRegions() []*ListRegionsResponseBodyRegions
	SetRequestId(v string) *ListRegionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRegionsResponseBody
	GetTotalCount() *int32
}

type ListRegionsResponseBody struct {
	// The maximum number of records returned per request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The next token.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The requested region IDs.
	Regions []*ListRegionsResponseBodyRegions `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRegionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRegionsResponseBody) GetRegions() []*ListRegionsResponseBodyRegions {
	return s.Regions
}

func (s *ListRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRegionsResponseBody) SetMaxResults(v int32) *ListRegionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRegionsResponseBody) SetNextToken(v string) *ListRegionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetTotalCount(v int32) *ListRegionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRegionsResponseBody) Validate() error {
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

type ListRegionsResponseBodyRegions struct {
	// The English name.
	//
	// example:
	//
	// cn-shenzhen
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
	// The region name.
	//
	// example:
	//
	// Shenzhen
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) GetEnName() *string {
	return s.EnName
}

func (s *ListRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRegionsResponseBodyRegions) SetEnName(v string) *ListRegionsResponseBodyRegions {
	s.EnName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
