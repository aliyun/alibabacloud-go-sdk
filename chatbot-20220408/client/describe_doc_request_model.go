// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DescribeDocRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *DescribeDocRequest
	GetKnowledgeId() *int64
	SetShowDetail(v bool) *DescribeDocRequest
	GetShowDetail() *bool
}

type DescribeDocRequest struct {
	// The key that identifies the workspace. If this parameter is omitted, the default workspace is used. You can find this key on the workspace management page of your root account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The knowledge ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001979424
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// Specifies whether to return detailed information for paragraphs. true: Detailed information is returned. false: Detailed information is not returned (default).
	//
	// example:
	//
	// false
	ShowDetail *bool `json:"ShowDetail,omitempty" xml:"ShowDetail,omitempty"`
}

func (s DescribeDocRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocRequest) GoString() string {
	return s.String()
}

func (s *DescribeDocRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeDocRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *DescribeDocRequest) GetShowDetail() *bool {
	return s.ShowDetail
}

func (s *DescribeDocRequest) SetAgentKey(v string) *DescribeDocRequest {
	s.AgentKey = &v
	return s
}

func (s *DescribeDocRequest) SetKnowledgeId(v int64) *DescribeDocRequest {
	s.KnowledgeId = &v
	return s
}

func (s *DescribeDocRequest) SetShowDetail(v bool) *DescribeDocRequest {
	s.ShowDetail = &v
	return s
}

func (s *DescribeDocRequest) Validate() error {
	return dara.Validate(s)
}
