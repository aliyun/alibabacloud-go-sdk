// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTranscodeTemplateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTranscodeTemplateGroupResponseBody
	GetRequestId() *string
	SetTranscodeTemplateGroupId(v string) *UpdateTranscodeTemplateGroupResponseBody
	GetTranscodeTemplateGroupId() *string
}

type UpdateTranscodeTemplateGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// 34e908aa4024a*****f7821c31f93a2a
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s UpdateTranscodeTemplateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTranscodeTemplateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTranscodeTemplateGroupResponseBody) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *UpdateTranscodeTemplateGroupResponseBody) SetRequestId(v string) *UpdateTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupResponseBody) SetTranscodeTemplateGroupId(v string) *UpdateTranscodeTemplateGroupResponseBody {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
