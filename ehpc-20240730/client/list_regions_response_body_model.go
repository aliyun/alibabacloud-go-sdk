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
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0axxxx
	NextToken *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Regions   []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
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
