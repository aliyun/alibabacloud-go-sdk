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
	// The task ID (DtsJobId) returned when the document parsing task was created.
	//
	// example:
	//
	// dts-20250729-y0zz3t13h7d****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The instance ID of the DTS RagFlow instance.
	//
	// example:
	//
	// dtsgbe12he619iy801
	RagInstanceId *string `json:"RagInstanceId,omitempty" xml:"RagInstanceId,omitempty"`
	// The region ID of the task. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
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
