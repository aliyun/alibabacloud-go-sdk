// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSolutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateSolutionRequest
	GetAgentKey() *string
	SetContent(v string) *UpdateSolutionRequest
	GetContent() *string
	SetContentType(v int32) *UpdateSolutionRequest
	GetContentType() *int32
	SetPerspectiveCodes(v []*string) *UpdateSolutionRequest
	GetPerspectiveCodes() []*string
	SetSolutionId(v int64) *UpdateSolutionRequest
	GetSolutionId() *int64
	SetTagIdList(v []*int64) *UpdateSolutionRequest
	GetTagIdList() []*int64
}

type UpdateSolutionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 100001333260
	SolutionId *int64   `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	TagIdList  []*int64 `json:"TagIdList,omitempty" xml:"TagIdList,omitempty" type:"Repeated"`
}

func (s UpdateSolutionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSolutionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSolutionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateSolutionRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateSolutionRequest) GetContentType() *int32 {
	return s.ContentType
}

func (s *UpdateSolutionRequest) GetPerspectiveCodes() []*string {
	return s.PerspectiveCodes
}

func (s *UpdateSolutionRequest) GetSolutionId() *int64 {
	return s.SolutionId
}

func (s *UpdateSolutionRequest) GetTagIdList() []*int64 {
	return s.TagIdList
}

func (s *UpdateSolutionRequest) SetAgentKey(v string) *UpdateSolutionRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateSolutionRequest) SetContent(v string) *UpdateSolutionRequest {
	s.Content = &v
	return s
}

func (s *UpdateSolutionRequest) SetContentType(v int32) *UpdateSolutionRequest {
	s.ContentType = &v
	return s
}

func (s *UpdateSolutionRequest) SetPerspectiveCodes(v []*string) *UpdateSolutionRequest {
	s.PerspectiveCodes = v
	return s
}

func (s *UpdateSolutionRequest) SetSolutionId(v int64) *UpdateSolutionRequest {
	s.SolutionId = &v
	return s
}

func (s *UpdateSolutionRequest) SetTagIdList(v []*int64) *UpdateSolutionRequest {
	s.TagIdList = v
	return s
}

func (s *UpdateSolutionRequest) Validate() error {
	return dara.Validate(s)
}
