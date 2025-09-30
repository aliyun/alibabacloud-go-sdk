// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithKnowledgeBaseStreamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ChatWithKnowledgeBaseStreamShrinkRequest
	GetDBInstanceId() *string
	SetIncludeKnowledgeBaseResults(v bool) *ChatWithKnowledgeBaseStreamShrinkRequest
	GetIncludeKnowledgeBaseResults() *bool
	SetKnowledgeParamsShrink(v string) *ChatWithKnowledgeBaseStreamShrinkRequest
	GetKnowledgeParamsShrink() *string
	SetModelParamsShrink(v string) *ChatWithKnowledgeBaseStreamShrinkRequest
	GetModelParamsShrink() *string
	SetOwnerId(v int64) *ChatWithKnowledgeBaseStreamShrinkRequest
	GetOwnerId() *int64
	SetPromptParams(v string) *ChatWithKnowledgeBaseStreamShrinkRequest
	GetPromptParams() *string
}

type ChatWithKnowledgeBaseStreamShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// false
	IncludeKnowledgeBaseResults *bool   `json:"IncludeKnowledgeBaseResults,omitempty" xml:"IncludeKnowledgeBaseResults,omitempty"`
	KnowledgeParamsShrink       *string `json:"KnowledgeParams,omitempty" xml:"KnowledgeParams,omitempty"`
	// This parameter is required.
	ModelParamsShrink *string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PromptParams      *string `json:"PromptParams,omitempty" xml:"PromptParams,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) GetIncludeKnowledgeBaseResults() *bool {
	return s.IncludeKnowledgeBaseResults
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) GetKnowledgeParamsShrink() *string {
	return s.KnowledgeParamsShrink
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) GetModelParamsShrink() *string {
	return s.ModelParamsShrink
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) GetPromptParams() *string {
	return s.PromptParams
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) SetDBInstanceId(v string) *ChatWithKnowledgeBaseStreamShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) SetIncludeKnowledgeBaseResults(v bool) *ChatWithKnowledgeBaseStreamShrinkRequest {
	s.IncludeKnowledgeBaseResults = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) SetKnowledgeParamsShrink(v string) *ChatWithKnowledgeBaseStreamShrinkRequest {
	s.KnowledgeParamsShrink = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) SetModelParamsShrink(v string) *ChatWithKnowledgeBaseStreamShrinkRequest {
	s.ModelParamsShrink = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) SetOwnerId(v int64) *ChatWithKnowledgeBaseStreamShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) SetPromptParams(v string) *ChatWithKnowledgeBaseStreamShrinkRequest {
	s.PromptParams = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
