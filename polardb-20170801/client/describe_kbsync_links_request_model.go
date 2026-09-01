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
	SetRegionId(v string) *DescribeKBSyncLinksRequest
	GetRegionId() *string
}

type DescribeKBSyncLinksRequest struct {
	// example:
	//
	// FEISHU
	ImPlatform *string `json:"ImPlatform,omitempty" xml:"ImPlatform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
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

func (s *DescribeKBSyncLinksRequest) SetRegionId(v string) *DescribeKBSyncLinksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKBSyncLinksRequest) Validate() error {
	return dara.Validate(s)
}
