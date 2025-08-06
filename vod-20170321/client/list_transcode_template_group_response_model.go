// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeTemplateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTranscodeTemplateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTranscodeTemplateGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListTranscodeTemplateGroupResponseBody) *ListTranscodeTemplateGroupResponse
	GetBody() *ListTranscodeTemplateGroupResponseBody
}

type ListTranscodeTemplateGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTranscodeTemplateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *ListTranscodeTemplateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTranscodeTemplateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTranscodeTemplateGroupResponse) GetBody() *ListTranscodeTemplateGroupResponseBody {
	return s.Body
}

func (s *ListTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *ListTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *ListTranscodeTemplateGroupResponse) SetStatusCode(v int32) *ListTranscodeTemplateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponse) SetBody(v *ListTranscodeTemplateGroupResponseBody) *ListTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

func (s *ListTranscodeTemplateGroupResponse) Validate() error {
	return dara.Validate(s)
}
