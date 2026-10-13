// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSolutionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateSolutionShrinkRequest
	GetAgentKey() *string
	SetContent(v string) *UpdateSolutionShrinkRequest
	GetContent() *string
	SetContentType(v int32) *UpdateSolutionShrinkRequest
	GetContentType() *int32
	SetPerspectiveCodes(v []*string) *UpdateSolutionShrinkRequest
	GetPerspectiveCodes() []*string
	SetSolutionId(v int64) *UpdateSolutionShrinkRequest
	GetSolutionId() *int64
	SetTagIdListShrink(v string) *UpdateSolutionShrinkRequest
	GetTagIdListShrink() *string
}

type UpdateSolutionShrinkRequest struct {
	// The workspace key. If this parameter is not specified, the default workspace is accessed. You can obtain the key from the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The answer content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Updated answer content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The answer type. Valid values:
	//
	// - 0: plain text
	//
	// - 1: rich text
	//
	// example:
	//
	// 1
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The list of perspective codes.
	//
	// This parameter is required.
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
	// The answer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100001333260
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	// The list of tag IDs.
	TagIdListShrink *string `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
}

func (s UpdateSolutionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSolutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSolutionShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateSolutionShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateSolutionShrinkRequest) GetContentType() *int32 {
	return s.ContentType
}

func (s *UpdateSolutionShrinkRequest) GetPerspectiveCodes() []*string {
	return s.PerspectiveCodes
}

func (s *UpdateSolutionShrinkRequest) GetSolutionId() *int64 {
	return s.SolutionId
}

func (s *UpdateSolutionShrinkRequest) GetTagIdListShrink() *string {
	return s.TagIdListShrink
}

func (s *UpdateSolutionShrinkRequest) SetAgentKey(v string) *UpdateSolutionShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateSolutionShrinkRequest) SetContent(v string) *UpdateSolutionShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateSolutionShrinkRequest) SetContentType(v int32) *UpdateSolutionShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *UpdateSolutionShrinkRequest) SetPerspectiveCodes(v []*string) *UpdateSolutionShrinkRequest {
	s.PerspectiveCodes = v
	return s
}

func (s *UpdateSolutionShrinkRequest) SetSolutionId(v int64) *UpdateSolutionShrinkRequest {
	s.SolutionId = &v
	return s
}

func (s *UpdateSolutionShrinkRequest) SetTagIdListShrink(v string) *UpdateSolutionShrinkRequest {
	s.TagIdListShrink = &v
	return s
}

func (s *UpdateSolutionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
