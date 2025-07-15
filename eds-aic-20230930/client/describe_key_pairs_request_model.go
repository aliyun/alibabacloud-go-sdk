// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairIds(v []*string) *DescribeKeyPairsRequest
	GetKeyPairIds() []*string
	SetKeyPairName(v string) *DescribeKeyPairsRequest
	GetKeyPairName() *string
	SetMaxResults(v int32) *DescribeKeyPairsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeKeyPairsRequest
	GetNextToken() *string
}

type DescribeKeyPairsRequest struct {
	// The IDs of the ADB key pairs.
	KeyPairIds []*string `json:"KeyPairIds,omitempty" xml:"KeyPairIds,omitempty" type:"Repeated"`
	// The name of the ADB key pair.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAYRHtOLVQzCYj17y+OP7LZQBUVVbi0GTu8g5****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeKeyPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsRequest) GetKeyPairIds() []*string {
	return s.KeyPairIds
}

func (s *DescribeKeyPairsRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeKeyPairsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeKeyPairsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeKeyPairsRequest) SetKeyPairIds(v []*string) *DescribeKeyPairsRequest {
	s.KeyPairIds = v
	return s
}

func (s *DescribeKeyPairsRequest) SetKeyPairName(v string) *DescribeKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetMaxResults(v int32) *DescribeKeyPairsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetNextToken(v string) *DescribeKeyPairsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeKeyPairsRequest) Validate() error {
	return dara.Validate(s)
}
