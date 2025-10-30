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
