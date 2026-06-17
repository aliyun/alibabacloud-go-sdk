// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithKnowledgeBaseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ChatWithKnowledgeBaseShrinkRequest
	GetDBInstanceId() *string
	SetIncludeKnowledgeBaseResults(v bool) *ChatWithKnowledgeBaseShrinkRequest
	GetIncludeKnowledgeBaseResults() *bool
	SetKnowledgeParamsShrink(v string) *ChatWithKnowledgeBaseShrinkRequest
	GetKnowledgeParamsShrink() *string
	SetModelParamsShrink(v string) *ChatWithKnowledgeBaseShrinkRequest
	GetModelParamsShrink() *string
	SetOwnerId(v int64) *ChatWithKnowledgeBaseShrinkRequest
	GetOwnerId() *int64
	SetPromptParams(v string) *ChatWithKnowledgeBaseShrinkRequest
	GetPromptParams() *string
	SetRegionId(v string) *ChatWithKnowledgeBaseShrinkRequest
	GetRegionId() *string
}

type ChatWithKnowledgeBaseShrinkRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to view the details of all instances in a target region, including their instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Whether to include the raw retrieval results from the knowledge base in the response. Default: `false`.
	//
	// example:
	//
	// false
	IncludeKnowledgeBaseResults *bool `json:"IncludeKnowledgeBaseResults,omitempty" xml:"IncludeKnowledgeBaseResults,omitempty"`
	// Parameters for knowledge retrieval. If omitted, the operation performs a standard chat without retrieving from a knowledge base.
	KnowledgeParamsShrink *string `json:"KnowledgeParams,omitempty" xml:"KnowledgeParams,omitempty"`
	// The parameters for calling the large language model (LLM).
	//
	// This parameter is required.
	ModelParamsShrink *string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A custom system prompt template. If specified, it overrides the default prompt. The template must include the {{ text_chunks }}, {{ user_system_prompt }}, {{ graph_entities }}, and {{ graph_relations }} placeholders.
	//
	// example:
	//
	// "参考以下知识回答问题:{{ text_chunks }}"
	PromptParams *string `json:"PromptParams,omitempty" xml:"PromptParams,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ChatWithKnowledgeBaseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ChatWithKnowledgeBaseShrinkRequest) GetIncludeKnowledgeBaseResults() *bool {
	return s.IncludeKnowledgeBaseResults
}

func (s *ChatWithKnowledgeBaseShrinkRequest) GetKnowledgeParamsShrink() *string {
	return s.KnowledgeParamsShrink
}

func (s *ChatWithKnowledgeBaseShrinkRequest) GetModelParamsShrink() *string {
	return s.ModelParamsShrink
}

func (s *ChatWithKnowledgeBaseShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatWithKnowledgeBaseShrinkRequest) GetPromptParams() *string {
	return s.PromptParams
}

func (s *ChatWithKnowledgeBaseShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChatWithKnowledgeBaseShrinkRequest) SetDBInstanceId(v string) *ChatWithKnowledgeBaseShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ChatWithKnowledgeBaseShrinkRequest) SetIncludeKnowledgeBaseResults(v bool) *ChatWithKnowledgeBaseShrinkRequest {
	s.IncludeKnowledgeBaseResults = &v
	return s
}

func (s *ChatWithKnowledgeBaseShrinkRequest) SetKnowledgeParamsShrink(v string) *ChatWithKnowledgeBaseShrinkRequest {
	s.KnowledgeParamsShrink = &v
	return s
}

func (s *ChatWithKnowledgeBaseShrinkRequest) SetModelParamsShrink(v string) *ChatWithKnowledgeBaseShrinkRequest {
	s.ModelParamsShrink = &v
	return s
}

func (s *ChatWithKnowledgeBaseShrinkRequest) SetOwnerId(v int64) *ChatWithKnowledgeBaseShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatWithKnowledgeBaseShrinkRequest) SetPromptParams(v string) *ChatWithKnowledgeBaseShrinkRequest {
	s.PromptParams = &v
	return s
}

func (s *ChatWithKnowledgeBaseShrinkRequest) SetRegionId(v string) *ChatWithKnowledgeBaseShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ChatWithKnowledgeBaseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
