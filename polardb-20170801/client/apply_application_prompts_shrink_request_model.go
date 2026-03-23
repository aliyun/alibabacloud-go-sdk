// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyApplicationPromptsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ApplyApplicationPromptsShrinkRequest
	GetApplicationId() *string
	SetDisabledPromptIdsShrink(v string) *ApplyApplicationPromptsShrinkRequest
	GetDisabledPromptIdsShrink() *string
	SetEnabledPromptIdsShrink(v string) *ApplyApplicationPromptsShrinkRequest
	GetEnabledPromptIdsShrink() *string
}

type ApplyApplicationPromptsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// papt-************
	DisabledPromptIdsShrink *string `json:"DisabledPromptIds,omitempty" xml:"DisabledPromptIds,omitempty"`
	// example:
	//
	// papt-************
	EnabledPromptIdsShrink *string `json:"EnabledPromptIds,omitempty" xml:"EnabledPromptIds,omitempty"`
}

func (s ApplyApplicationPromptsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyApplicationPromptsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyApplicationPromptsShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ApplyApplicationPromptsShrinkRequest) GetDisabledPromptIdsShrink() *string {
	return s.DisabledPromptIdsShrink
}

func (s *ApplyApplicationPromptsShrinkRequest) GetEnabledPromptIdsShrink() *string {
	return s.EnabledPromptIdsShrink
}

func (s *ApplyApplicationPromptsShrinkRequest) SetApplicationId(v string) *ApplyApplicationPromptsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ApplyApplicationPromptsShrinkRequest) SetDisabledPromptIdsShrink(v string) *ApplyApplicationPromptsShrinkRequest {
	s.DisabledPromptIdsShrink = &v
	return s
}

func (s *ApplyApplicationPromptsShrinkRequest) SetEnabledPromptIdsShrink(v string) *ApplyApplicationPromptsShrinkRequest {
	s.EnabledPromptIdsShrink = &v
	return s
}

func (s *ApplyApplicationPromptsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
