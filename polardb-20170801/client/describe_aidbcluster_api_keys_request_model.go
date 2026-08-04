// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAIDBClusterApiKeysRequest
	GetRegionId() *string
}

type DescribeAIDBClusterApiKeysRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAIDBClusterApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterApiKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterApiKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClusterApiKeysRequest) SetRegionId(v string) *DescribeAIDBClusterApiKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
