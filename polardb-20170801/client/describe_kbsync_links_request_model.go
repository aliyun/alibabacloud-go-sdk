// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKBSyncLinksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImPlatform(v string) *DescribeKBSyncLinksRequest
	GetImPlatform() *string
	SetKnowledgeBaseId(v string) *DescribeKBSyncLinksRequest
	GetKnowledgeBaseId() *string
	SetPageNumber(v int32) *DescribeKBSyncLinksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKBSyncLinksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeKBSyncLinksRequest
	GetRegionId() *string
}

type DescribeKBSyncLinksRequest struct {
	// The source channel of the synchronization link.
	//
	// example:
	//
	// FEISHU
	ImPlatform *string `json:"ImPlatform,omitempty" xml:"ImPlatform,omitempty"`
	// The unique ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The page number of the query results. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of synchronization links returned per page. Valid values: 10, 20, 30, 50, 100, 200, and 500. Default value: 30.
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

func (s DescribeKBSyncLinksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKBSyncLinksRequest) GoString() string {
	return s.String()
}

func (s *DescribeKBSyncLinksRequest) GetImPlatform() *string {
	return s.ImPlatform
}

func (s *DescribeKBSyncLinksRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKBSyncLinksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKBSyncLinksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKBSyncLinksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKBSyncLinksRequest) SetImPlatform(v string) *DescribeKBSyncLinksRequest {
	s.ImPlatform = &v
	return s
}

func (s *DescribeKBSyncLinksRequest) SetKnowledgeBaseId(v string) *DescribeKBSyncLinksRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKBSyncLinksRequest) SetPageNumber(v int32) *DescribeKBSyncLinksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKBSyncLinksRequest) SetPageSize(v int32) *DescribeKBSyncLinksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeKBSyncLinksRequest) SetRegionId(v string) *DescribeKBSyncLinksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKBSyncLinksRequest) Validate() error {
	return dara.Validate(s)
}
