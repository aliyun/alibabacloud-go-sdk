// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseCode(v string) *ListEnterpriseCodeRequest
	GetEnterpriseCode() *string
	SetIsDefault(v bool) *ListEnterpriseCodeRequest
	GetIsDefault() *bool
	SetMaxResults(v int32) *ListEnterpriseCodeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEnterpriseCodeRequest
	GetNextToken() *string
	SetRegionId(v string) *ListEnterpriseCodeRequest
	GetRegionId() *string
}

type ListEnterpriseCodeRequest struct {
	// The enterprise code that you want to query.
	//
	// example:
	//
	// 12P**
	EnterpriseCode *string `json:"EnterpriseCode,omitempty" xml:"EnterpriseCode,omitempty"`
	// Specifies whether to query only default enterprise codes. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false*	- (default): no
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The number of entries returned per page.
	//
	// Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for returning the next page when the data is returned in more than one page.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnterpriseCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseCodeRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseCodeRequest) GetEnterpriseCode() *string {
	return s.EnterpriseCode
}

func (s *ListEnterpriseCodeRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListEnterpriseCodeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEnterpriseCodeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEnterpriseCodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnterpriseCodeRequest) SetEnterpriseCode(v string) *ListEnterpriseCodeRequest {
	s.EnterpriseCode = &v
	return s
}

func (s *ListEnterpriseCodeRequest) SetIsDefault(v bool) *ListEnterpriseCodeRequest {
	s.IsDefault = &v
	return s
}

func (s *ListEnterpriseCodeRequest) SetMaxResults(v int32) *ListEnterpriseCodeRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEnterpriseCodeRequest) SetNextToken(v string) *ListEnterpriseCodeRequest {
	s.NextToken = &v
	return s
}

func (s *ListEnterpriseCodeRequest) SetRegionId(v string) *ListEnterpriseCodeRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnterpriseCodeRequest) Validate() error {
	return dara.Validate(s)
}
