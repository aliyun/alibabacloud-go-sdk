// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductDimensionGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDimensionGroups(v []*ListProductDimensionGroupsResponseBodyDimensionGroups) *ListProductDimensionGroupsResponseBody
	GetDimensionGroups() []*ListProductDimensionGroupsResponseBodyDimensionGroups
	SetMaxResults(v int32) *ListProductDimensionGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductDimensionGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListProductDimensionGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductDimensionGroupsResponseBody
	GetTotalCount() *int32
}

type ListProductDimensionGroupsResponseBody struct {
	// The dimension groups.
	DimensionGroups []*ListProductDimensionGroupsResponseBodyDimensionGroups `json:"DimensionGroups,omitempty" xml:"DimensionGroups,omitempty" type:"Repeated"`
	// The maximum number of records that are returned for the query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position at which the query ends. An empty value indicates that all data is returned.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 057D210F-F2FC-5329-A536-26C16628BB09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records that are returned for the query.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductDimensionGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductDimensionGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductDimensionGroupsResponseBody) GetDimensionGroups() []*ListProductDimensionGroupsResponseBodyDimensionGroups {
	return s.DimensionGroups
}

func (s *ListProductDimensionGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductDimensionGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductDimensionGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductDimensionGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductDimensionGroupsResponseBody) SetDimensionGroups(v []*ListProductDimensionGroupsResponseBodyDimensionGroups) *ListProductDimensionGroupsResponseBody {
	s.DimensionGroups = v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) SetMaxResults(v int32) *ListProductDimensionGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) SetNextToken(v string) *ListProductDimensionGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) SetRequestId(v string) *ListProductDimensionGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) SetTotalCount(v int32) *ListProductDimensionGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProductDimensionGroupsResponseBodyDimensionGroups struct {
	// The key of the dimension group.
	DimensionKeys []*string `json:"DimensionKeys,omitempty" xml:"DimensionKeys,omitempty" type:"Repeated"`
	// The code of the dimension group.
	//
	// example:
	//
	// oss_wf1ngqmd7q
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the dimension group.
	//
	// example:
	//
	// OSS_Group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The service code.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductDimensionGroupsResponseBodyDimensionGroups) String() string {
	return dara.Prettify(s)
}

func (s ListProductDimensionGroupsResponseBodyDimensionGroups) GoString() string {
	return s.String()
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) GetDimensionKeys() []*string {
	return s.DimensionKeys
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) GetGroupCode() *string {
	return s.GroupCode
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) SetDimensionKeys(v []*string) *ListProductDimensionGroupsResponseBodyDimensionGroups {
	s.DimensionKeys = v
	return s
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) SetGroupCode(v string) *ListProductDimensionGroupsResponseBodyDimensionGroups {
	s.GroupCode = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) SetGroupName(v string) *ListProductDimensionGroupsResponseBodyDimensionGroups {
	s.GroupName = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) SetProductCode(v string) *ListProductDimensionGroupsResponseBodyDimensionGroups {
	s.ProductCode = &v
	return s
}

func (s *ListProductDimensionGroupsResponseBodyDimensionGroups) Validate() error {
	return dara.Validate(s)
}
