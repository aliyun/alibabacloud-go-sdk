// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocParserJobResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *DescribeDocParserJobResultRequest
	GetAgentName() *string
	SetJobId(v string) *DescribeDocParserJobResultRequest
	GetJobId() *string
	SetRegionId(v string) *DescribeDocParserJobResultRequest
	GetRegionId() *string
	SetResultType(v string) *DescribeDocParserJobResultRequest
	GetResultType() *string
}

type DescribeDocParserJobResultRequest struct {
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The document parsing task ID. You can obtain this ID by calling CreateDocParserJob.
	//
	// This parameter is required.
	//
	// example:
	//
	// job_abc123
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s DescribeDocParserJobResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocParserJobResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocParserJobResultRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *DescribeDocParserJobResultRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDocParserJobResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDocParserJobResultRequest) GetResultType() *string {
	return s.ResultType
}

func (s *DescribeDocParserJobResultRequest) SetAgentName(v string) *DescribeDocParserJobResultRequest {
	s.AgentName = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) SetJobId(v string) *DescribeDocParserJobResultRequest {
	s.JobId = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) SetRegionId(v string) *DescribeDocParserJobResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) SetResultType(v string) *DescribeDocParserJobResultRequest {
	s.ResultType = &v
	return s
}

func (s *DescribeDocParserJobResultRequest) Validate() error {
	return dara.Validate(s)
}
