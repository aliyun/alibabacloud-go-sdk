// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseFileShardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v string) *DescribeKnowledgeBaseFileShardsRequest
	GetFileId() *string
	SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseFileShardsRequest
	GetKnowledgeBaseId() *string
	SetPageNumber(v int32) *DescribeKnowledgeBaseFileShardsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKnowledgeBaseFileShardsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeKnowledgeBaseFileShardsRequest
	GetRegionId() *string
}

type DescribeKnowledgeBaseFileShardsRequest struct {
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The unique ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
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
}

func (s DescribeKnowledgeBaseFileShardsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFileShardsRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFileShardsRequest) GetFileId() *string {
	return s.FileId
}

func (s *DescribeKnowledgeBaseFileShardsRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKnowledgeBaseFileShardsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKnowledgeBaseFileShardsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKnowledgeBaseFileShardsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKnowledgeBaseFileShardsRequest) SetFileId(v string) *DescribeKnowledgeBaseFileShardsRequest {
	s.FileId = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsRequest) SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseFileShardsRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsRequest) SetPageNumber(v int32) *DescribeKnowledgeBaseFileShardsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsRequest) SetPageSize(v int32) *DescribeKnowledgeBaseFileShardsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsRequest) SetRegionId(v string) *DescribeKnowledgeBaseFileShardsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKnowledgeBaseFileShardsRequest) Validate() error {
	return dara.Validate(s)
}
