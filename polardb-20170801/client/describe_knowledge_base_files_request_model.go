// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileIds(v string) *DescribeKnowledgeBaseFilesRequest
	GetFileIds() *string
	SetKeyword(v string) *DescribeKnowledgeBaseFilesRequest
	GetKeyword() *string
	SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseFilesRequest
	GetKnowledgeBaseId() *string
	SetLinkId(v string) *DescribeKnowledgeBaseFilesRequest
	GetLinkId() *string
	SetPageNumber(v int32) *DescribeKnowledgeBaseFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKnowledgeBaseFilesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeKnowledgeBaseFilesRequest
	GetRegionId() *string
	SetSourceType(v string) *DescribeKnowledgeBaseFilesRequest
	GetSourceType() *string
}

type DescribeKnowledgeBaseFilesRequest struct {
	// The list of file IDs, separated by commas (,).
	//
	// example:
	//
	// doc_a,doc_b
	FileIds *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// The keyword used to filter file names.
	//
	// example:
	//
	// Financial report
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The unique ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The unique ID of the synchronization link.
	//
	// example:
	//
	// pkbl-2ze123456789abc
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source type.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeKnowledgeBaseFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesRequest) GetFileIds() *string {
	return s.FileIds
}

func (s *DescribeKnowledgeBaseFilesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeKnowledgeBaseFilesRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKnowledgeBaseFilesRequest) GetLinkId() *string {
	return s.LinkId
}

func (s *DescribeKnowledgeBaseFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKnowledgeBaseFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKnowledgeBaseFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKnowledgeBaseFilesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeKnowledgeBaseFilesRequest) SetFileIds(v string) *DescribeKnowledgeBaseFilesRequest {
	s.FileIds = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesRequest) SetKeyword(v string) *DescribeKnowledgeBaseFilesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesRequest) SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseFilesRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesRequest) SetLinkId(v string) *DescribeKnowledgeBaseFilesRequest {
	s.LinkId = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesRequest) SetPageNumber(v int32) *DescribeKnowledgeBaseFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesRequest) SetPageSize(v int32) *DescribeKnowledgeBaseFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesRequest) SetRegionId(v string) *DescribeKnowledgeBaseFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesRequest) SetSourceType(v string) *DescribeKnowledgeBaseFilesRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesRequest) Validate() error {
	return dara.Validate(s)
}
