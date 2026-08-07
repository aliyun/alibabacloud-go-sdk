// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelSpaceName(v string) *DescribeAIDBClusterApiKeysRequest
	GetModelSpaceName() *string
	SetRegionId(v string) *DescribeAIDBClusterApiKeysRequest
	GetRegionId() *string
}

type DescribeAIDBClusterApiKeysRequest struct {
	// example:
	//
	// pms-xxx
	ModelSpaceName *string `json:"ModelSpaceName,omitempty" xml:"ModelSpaceName,omitempty"`
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

func (s *DescribeAIDBClusterApiKeysRequest) GetModelSpaceName() *string {
	return s.ModelSpaceName
}

func (s *DescribeAIDBClusterApiKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClusterApiKeysRequest) SetModelSpaceName(v string) *DescribeAIDBClusterApiKeysRequest {
	s.ModelSpaceName = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysRequest) SetRegionId(v string) *DescribeAIDBClusterApiKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClusterApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
