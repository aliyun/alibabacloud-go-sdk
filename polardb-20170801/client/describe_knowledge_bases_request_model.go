// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeKnowledgeBasesRequest
	GetKeyword() *string
	SetKnowledgeSpaceId(v string) *DescribeKnowledgeBasesRequest
	GetKnowledgeSpaceId() *string
	SetPageNumber(v int32) *DescribeKnowledgeBasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKnowledgeBasesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeKnowledgeBasesRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeKnowledgeBasesRequest
	GetStatus() *string
}

type DescribeKnowledgeBasesRequest struct {
	// The keyword for searching knowledge bases.
	//
	// example:
	//
	// testkb
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The unique identifier of the knowledge space.
	//
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the knowledge base.
	//
	// example:
	//
	// Activation
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBasesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeKnowledgeBasesRequest) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *DescribeKnowledgeBasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKnowledgeBasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKnowledgeBasesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKnowledgeBasesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeKnowledgeBasesRequest) SetKeyword(v string) *DescribeKnowledgeBasesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeKnowledgeBasesRequest) SetKnowledgeSpaceId(v string) *DescribeKnowledgeBasesRequest {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *DescribeKnowledgeBasesRequest) SetPageNumber(v int32) *DescribeKnowledgeBasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKnowledgeBasesRequest) SetPageSize(v int32) *DescribeKnowledgeBasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeKnowledgeBasesRequest) SetRegionId(v string) *DescribeKnowledgeBasesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKnowledgeBasesRequest) SetStatus(v string) *DescribeKnowledgeBasesRequest {
	s.Status = &v
	return s
}

func (s *DescribeKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
