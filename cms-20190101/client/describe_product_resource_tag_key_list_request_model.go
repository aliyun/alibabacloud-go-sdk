// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductResourceTagKeyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeProductResourceTagKeyListRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeProductResourceTagKeyListRequest
	GetRegionId() *string
}

type DescribeProductResourceTagKeyListRequest struct {
	// The pagination token.
	//
	// example:
	//
	// dbc2826f237e****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeProductResourceTagKeyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResourceTagKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductResourceTagKeyListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProductResourceTagKeyListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeProductResourceTagKeyListRequest) SetNextToken(v string) *DescribeProductResourceTagKeyListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProductResourceTagKeyListRequest) SetRegionId(v string) *DescribeProductResourceTagKeyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProductResourceTagKeyListRequest) Validate() error {
	return dara.Validate(s)
}
