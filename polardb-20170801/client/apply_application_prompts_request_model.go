// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyApplicationPromptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ApplyApplicationPromptsRequest
	GetApplicationId() *string
	SetDisabledPromptIds(v []*string) *ApplyApplicationPromptsRequest
	GetDisabledPromptIds() []*string
	SetEnabledPromptIds(v []*string) *ApplyApplicationPromptsRequest
	GetEnabledPromptIds() []*string
}

type ApplyApplicationPromptsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// papt-************
	DisabledPromptIds []*string `json:"DisabledPromptIds,omitempty" xml:"DisabledPromptIds,omitempty" type:"Repeated"`
	// example:
	//
	// papt-************
	EnabledPromptIds []*string `json:"EnabledPromptIds,omitempty" xml:"EnabledPromptIds,omitempty" type:"Repeated"`
}

func (s ApplyApplicationPromptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyApplicationPromptsRequest) GoString() string {
	return s.String()
}

func (s *ApplyApplicationPromptsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ApplyApplicationPromptsRequest) GetDisabledPromptIds() []*string {
	return s.DisabledPromptIds
}

func (s *ApplyApplicationPromptsRequest) GetEnabledPromptIds() []*string {
	return s.EnabledPromptIds
}

func (s *ApplyApplicationPromptsRequest) SetApplicationId(v string) *ApplyApplicationPromptsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ApplyApplicationPromptsRequest) SetDisabledPromptIds(v []*string) *ApplyApplicationPromptsRequest {
	s.DisabledPromptIds = v
	return s
}

func (s *ApplyApplicationPromptsRequest) SetEnabledPromptIds(v []*string) *ApplyApplicationPromptsRequest {
	s.EnabledPromptIds = v
	return s
}

func (s *ApplyApplicationPromptsRequest) Validate() error {
	return dara.Validate(s)
}
