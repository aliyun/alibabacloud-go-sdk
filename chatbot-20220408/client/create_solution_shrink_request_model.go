// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSolutionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateSolutionShrinkRequest
	GetAgentKey() *string
	SetContent(v string) *CreateSolutionShrinkRequest
	GetContent() *string
	SetContentType(v int32) *CreateSolutionShrinkRequest
	GetContentType() *int32
	SetKnowledgeId(v int64) *CreateSolutionShrinkRequest
	GetKnowledgeId() *int64
	SetPerspectiveCodes(v []*string) *CreateSolutionShrinkRequest
	GetPerspectiveCodes() []*string
	SetTagIdListShrink(v string) *CreateSolutionShrinkRequest
	GetTagIdListShrink() *string
}

type CreateSolutionShrinkRequest struct {
	// The key of the business space. If this parameter is omitted, the system uses the default business space. You can obtain the key from the Business Management page of your primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The content of the knowledge answer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 答案内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The answer type. Valid values: `0` (plain text) and `1` (rich text).
	//
	// example:
	//
	// 1
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The ID of the knowledge.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// A list of perspective codes.
	//
	// This parameter is required.
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
	// A list of tag IDs.
	TagIdListShrink *string `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
}

func (s CreateSolutionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSolutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSolutionShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateSolutionShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *CreateSolutionShrinkRequest) GetContentType() *int32 {
	return s.ContentType
}

func (s *CreateSolutionShrinkRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *CreateSolutionShrinkRequest) GetPerspectiveCodes() []*string {
	return s.PerspectiveCodes
}

func (s *CreateSolutionShrinkRequest) GetTagIdListShrink() *string {
	return s.TagIdListShrink
}

func (s *CreateSolutionShrinkRequest) SetAgentKey(v string) *CreateSolutionShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateSolutionShrinkRequest) SetContent(v string) *CreateSolutionShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateSolutionShrinkRequest) SetContentType(v int32) *CreateSolutionShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *CreateSolutionShrinkRequest) SetKnowledgeId(v int64) *CreateSolutionShrinkRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateSolutionShrinkRequest) SetPerspectiveCodes(v []*string) *CreateSolutionShrinkRequest {
	s.PerspectiveCodes = v
	return s
}

func (s *CreateSolutionShrinkRequest) SetTagIdListShrink(v string) *CreateSolutionShrinkRequest {
	s.TagIdListShrink = &v
	return s
}

func (s *CreateSolutionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
