// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseAnswerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseAnswerRequest
	GetKnowledgeBaseId() *string
	SetQueryId(v string) *DescribeKnowledgeBaseAnswerRequest
	GetQueryId() *string
	SetRegionId(v string) *DescribeKnowledgeBaseAnswerRequest
	GetRegionId() *string
}

type DescribeKnowledgeBaseAnswerRequest struct {
	// The unique ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The unique ID of the Q&A task.
	//
	// This parameter is required.
	//
	// example:
	//
	// R3BGbnBqcXN******.2a5a23c9-******-179970533d30
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeKnowledgeBaseAnswerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseAnswerRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseAnswerRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKnowledgeBaseAnswerRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribeKnowledgeBaseAnswerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKnowledgeBaseAnswerRequest) SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseAnswerRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerRequest) SetQueryId(v string) *DescribeKnowledgeBaseAnswerRequest {
	s.QueryId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerRequest) SetRegionId(v string) *DescribeKnowledgeBaseAnswerRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerRequest) Validate() error {
	return dara.Validate(s)
}
