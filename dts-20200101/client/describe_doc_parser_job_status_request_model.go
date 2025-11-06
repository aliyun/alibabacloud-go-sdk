// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribeDocParserJobStatusRequest
	GetDtsJobId() *string
	SetRagInstanceId(v string) *DescribeDocParserJobStatusRequest
	GetRagInstanceId() *string
	SetRegionId(v string) *DescribeDocParserJobStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDocParserJobStatusRequest
	GetResourceGroupId() *string
}

type DescribeDocParserJobStatusRequest struct {
	// example:
	//
	// dts-20250729-l3m1213ye7l****
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

func (s DescribeDocParserJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobStatusRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDocParserJobStatusRequest) GetRagInstanceId() *string {
	return s.RagInstanceId
}

func (s *DescribeDocParserJobStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDocParserJobStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDocParserJobStatusRequest) SetDtsJobId(v string) *DescribeDocParserJobStatusRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDocParserJobStatusRequest) SetRagInstanceId(v string) *DescribeDocParserJobStatusRequest {
	s.RagInstanceId = &v
	return s
}

func (s *DescribeDocParserJobStatusRequest) SetRegionId(v string) *DescribeDocParserJobStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDocParserJobStatusRequest) SetResourceGroupId(v string) *DescribeDocParserJobStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDocParserJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
