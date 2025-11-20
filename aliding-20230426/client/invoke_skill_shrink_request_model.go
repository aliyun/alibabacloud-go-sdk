// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParamsShrink(v string) *InvokeSkillShrinkRequest
	GetParamsShrink() *string
	SetSkillId(v string) *InvokeSkillShrinkRequest
	GetSkillId() *string
	SetStream(v bool) *InvokeSkillShrinkRequest
	GetStream() *bool
	SetSourceIdOfAssistantId(v string) *InvokeSkillShrinkRequest
	GetSourceIdOfAssistantId() *string
}

type InvokeSkillShrinkRequest struct {
	// example:
	//
	// {}
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a1d033dd-xxxx-49cf-b49b-2068081bb551
	SkillId               *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	Stream                *bool   `json:"Stream,omitempty" xml:"Stream,omitempty"`
	SourceIdOfAssistantId *string `json:"sourceIdOfAssistantId,omitempty" xml:"sourceIdOfAssistantId,omitempty"`
}

func (s InvokeSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvokeSkillShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *InvokeSkillShrinkRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *InvokeSkillShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *InvokeSkillShrinkRequest) GetSourceIdOfAssistantId() *string {
	return s.SourceIdOfAssistantId
}

func (s *InvokeSkillShrinkRequest) SetParamsShrink(v string) *InvokeSkillShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *InvokeSkillShrinkRequest) SetSkillId(v string) *InvokeSkillShrinkRequest {
	s.SkillId = &v
	return s
}

func (s *InvokeSkillShrinkRequest) SetStream(v bool) *InvokeSkillShrinkRequest {
	s.Stream = &v
	return s
}

func (s *InvokeSkillShrinkRequest) SetSourceIdOfAssistantId(v string) *InvokeSkillShrinkRequest {
	s.SourceIdOfAssistantId = &v
	return s
}

func (s *InvokeSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
