// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTranscodeTemplateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTranscodeTemplateGroupResponseBody
	GetRequestId() *string
	SetTranscodeTemplateGroupId(v string) *AddTranscodeTemplateGroupResponseBody
	GetTranscodeTemplateGroupId() *string
}

type AddTranscodeTemplateGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// 34e908aa4024af7821c31f93a2a****
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s AddTranscodeTemplateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddTranscodeTemplateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTranscodeTemplateGroupResponseBody) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *AddTranscodeTemplateGroupResponseBody) SetRequestId(v string) *AddTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTranscodeTemplateGroupResponseBody) SetTranscodeTemplateGroupId(v string) *AddTranscodeTemplateGroupResponseBody {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *AddTranscodeTemplateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
