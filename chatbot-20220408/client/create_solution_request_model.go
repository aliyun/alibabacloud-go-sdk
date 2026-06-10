// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSolutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateSolutionRequest
	GetAgentKey() *string
	SetContent(v string) *CreateSolutionRequest
	GetContent() *string
	SetContentType(v int32) *CreateSolutionRequest
	GetContentType() *int32
	SetKnowledgeId(v int64) *CreateSolutionRequest
	GetKnowledgeId() *int64
	SetPerspectiveCodes(v []*string) *CreateSolutionRequest
	GetPerspectiveCodes() []*string
	SetTagIdList(v []*int64) *CreateSolutionRequest
	GetTagIdList() []*int64
}

type CreateSolutionRequest struct {
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
	TagIdList []*int64 `json:"TagIdList,omitempty" xml:"TagIdList,omitempty" type:"Repeated"`
}

func (s CreateSolutionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSolutionRequest) GoString() string {
	return s.String()
}

func (s *CreateSolutionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateSolutionRequest) GetContent() *string {
	return s.Content
}

func (s *CreateSolutionRequest) GetContentType() *int32 {
	return s.ContentType
}

func (s *CreateSolutionRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *CreateSolutionRequest) GetPerspectiveCodes() []*string {
	return s.PerspectiveCodes
}

func (s *CreateSolutionRequest) GetTagIdList() []*int64 {
	return s.TagIdList
}

func (s *CreateSolutionRequest) SetAgentKey(v string) *CreateSolutionRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateSolutionRequest) SetContent(v string) *CreateSolutionRequest {
	s.Content = &v
	return s
}

func (s *CreateSolutionRequest) SetContentType(v int32) *CreateSolutionRequest {
	s.ContentType = &v
	return s
}

func (s *CreateSolutionRequest) SetKnowledgeId(v int64) *CreateSolutionRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateSolutionRequest) SetPerspectiveCodes(v []*string) *CreateSolutionRequest {
	s.PerspectiveCodes = v
	return s
}

func (s *CreateSolutionRequest) SetTagIdList(v []*int64) *CreateSolutionRequest {
	s.TagIdList = v
	return s
}

func (s *CreateSolutionRequest) Validate() error {
	return dara.Validate(s)
}
