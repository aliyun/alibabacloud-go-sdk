// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKBSyncLinkAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *DescribeKBSyncLinkAttributeRequest
	GetKnowledgeBaseId() *string
	SetLinkId(v string) *DescribeKBSyncLinkAttributeRequest
	GetLinkId() *string
	SetRegionId(v string) *DescribeKBSyncLinkAttributeRequest
	GetRegionId() *string
}

type DescribeKBSyncLinkAttributeRequest struct {
	// The unique identifier of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-123456
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The unique identifier of the synchronization link.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkbl-123456
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeKBSyncLinkAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKBSyncLinkAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeKBSyncLinkAttributeRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKBSyncLinkAttributeRequest) GetLinkId() *string {
	return s.LinkId
}

func (s *DescribeKBSyncLinkAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKBSyncLinkAttributeRequest) SetKnowledgeBaseId(v string) *DescribeKBSyncLinkAttributeRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeRequest) SetLinkId(v string) *DescribeKBSyncLinkAttributeRequest {
	s.LinkId = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeRequest) SetRegionId(v string) *DescribeKBSyncLinkAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKBSyncLinkAttributeRequest) Validate() error {
	return dara.Validate(s)
}
