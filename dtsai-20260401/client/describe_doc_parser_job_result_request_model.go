// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DescribeDocParserJobResultRequest
	GetJobId() *string
	SetRegionId(v string) *DescribeDocParserJobResultRequest
	GetRegionId() *string
}

type DescribeDocParserJobResultRequest struct {
	// This parameter is required.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDocParserJobResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobResultRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDocParserJobResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDocParserJobResultRequest) SetJobId(v string) *DescribeDocParserJobResultRequest {
	s.JobId = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) SetRegionId(v string) *DescribeDocParserJobResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) Validate() error {
	return dara.Validate(s)
}
