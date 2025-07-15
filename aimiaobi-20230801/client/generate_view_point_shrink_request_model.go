// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateViewPointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GenerateViewPointShrinkRequest
	GetAgentKey() *string
	SetReferenceDataShrink(v string) *GenerateViewPointShrinkRequest
	GetReferenceDataShrink() *string
}

type GenerateViewPointShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey            *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
}

func (s GenerateViewPointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateViewPointShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateViewPointShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GenerateViewPointShrinkRequest) GetReferenceDataShrink() *string {
	return s.ReferenceDataShrink
}

func (s *GenerateViewPointShrinkRequest) SetAgentKey(v string) *GenerateViewPointShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateViewPointShrinkRequest) SetReferenceDataShrink(v string) *GenerateViewPointShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *GenerateViewPointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
