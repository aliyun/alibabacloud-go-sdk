// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseAttributeRequest
	GetKnowledgeBaseId() *string
	SetRegionId(v string) *DescribeKnowledgeBaseAttributeRequest
	GetRegionId() *string
}

type DescribeKnowledgeBaseAttributeRequest struct {
	// The unique identifier of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeKnowledgeBaseAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseAttributeRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKnowledgeBaseAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKnowledgeBaseAttributeRequest) SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseAttributeRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeRequest) SetRegionId(v string) *DescribeKnowledgeBaseAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKnowledgeBaseAttributeRequest) Validate() error {
	return dara.Validate(s)
}
