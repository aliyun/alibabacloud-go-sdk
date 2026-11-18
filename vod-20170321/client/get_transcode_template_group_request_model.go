// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTemplateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTranscodeTemplateGroupId(v string) *GetTranscodeTemplateGroupRequest
	GetTranscodeTemplateGroupId() *string
}

type GetTranscodeTemplateGroupRequest struct {
	// The ID of the transcoding template group.
	//
	// This parameter is required.
	//
	// example:
	//
	// a591f697c7167*****6ae1502142d0
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s GetTranscodeTemplateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupRequest) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *GetTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *GetTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *GetTranscodeTemplateGroupRequest) Validate() error {
	return dara.Validate(s)
}
