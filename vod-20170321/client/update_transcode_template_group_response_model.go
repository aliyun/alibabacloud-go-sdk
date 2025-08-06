// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTranscodeTemplateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTranscodeTemplateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTranscodeTemplateGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTranscodeTemplateGroupResponseBody) *UpdateTranscodeTemplateGroupResponse
	GetBody() *UpdateTranscodeTemplateGroupResponseBody
}

type UpdateTranscodeTemplateGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTranscodeTemplateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateTranscodeTemplateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTranscodeTemplateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTranscodeTemplateGroupResponse) GetBody() *UpdateTranscodeTemplateGroupResponseBody {
	return s.Body
}

func (s *UpdateTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *UpdateTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateTranscodeTemplateGroupResponse) SetStatusCode(v int32) *UpdateTranscodeTemplateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupResponse) SetBody(v *UpdateTranscodeTemplateGroupResponseBody) *UpdateTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateTranscodeTemplateGroupResponse) Validate() error {
	return dara.Validate(s)
}
