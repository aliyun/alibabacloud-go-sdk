// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DescribeDocParserJobStatusRequest
	GetJobId() *string
	SetRegionId(v string) *DescribeDocParserJobStatusRequest
	GetRegionId() *string
}

type DescribeDocParserJobStatusRequest struct {
	// This parameter is required.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDocParserJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobStatusRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDocParserJobStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDocParserJobStatusRequest) SetJobId(v string) *DescribeDocParserJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *DescribeDocParserJobStatusRequest) SetRegionId(v string) *DescribeDocParserJobStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDocParserJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
