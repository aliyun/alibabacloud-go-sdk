// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribeDocParserJobResultRequest
	GetDtsJobId() *string
	SetRagInstanceId(v string) *DescribeDocParserJobResultRequest
	GetRagInstanceId() *string
	SetRegionId(v string) *DescribeDocParserJobResultRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDocParserJobResultRequest
	GetResourceGroupId() *string
}

type DescribeDocParserJobResultRequest struct {
	// example:
	//
	// dts-20250729-y0zz3t13h7d****
	DtsJobId      *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	RagInstanceId *string `json:"RagInstanceId,omitempty" xml:"RagInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDocParserJobResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobResultRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDocParserJobResultRequest) GetRagInstanceId() *string {
	return s.RagInstanceId
}

func (s *DescribeDocParserJobResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDocParserJobResultRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDocParserJobResultRequest) SetDtsJobId(v string) *DescribeDocParserJobResultRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) SetRagInstanceId(v string) *DescribeDocParserJobResultRequest {
	s.RagInstanceId = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) SetRegionId(v string) *DescribeDocParserJobResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) SetResourceGroupId(v string) *DescribeDocParserJobResultRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) Validate() error {
	return dara.Validate(s)
}
