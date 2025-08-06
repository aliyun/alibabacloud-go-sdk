// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultTranscodeTemplateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTranscodeTemplateGroupId(v string) *SetDefaultTranscodeTemplateGroupRequest
	GetTranscodeTemplateGroupId() *string
}

type SetDefaultTranscodeTemplateGroupRequest struct {
	// The ID of the transcoding template group.
	//
	// This parameter is required.
	//
	// example:
	//
	// d58079958be8d*****b699ab7ab6e1bf
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s SetDefaultTranscodeTemplateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultTranscodeTemplateGroupRequest) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *SetDefaultTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *SetDefaultTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *SetDefaultTranscodeTemplateGroupRequest) Validate() error {
	return dara.Validate(s)
}
