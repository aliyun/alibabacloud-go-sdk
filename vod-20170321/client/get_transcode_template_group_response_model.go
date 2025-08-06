// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTemplateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranscodeTemplateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranscodeTemplateGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetTranscodeTemplateGroupResponseBody) *GetTranscodeTemplateGroupResponse
	GetBody() *GetTranscodeTemplateGroupResponseBody
}

type GetTranscodeTemplateGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscodeTemplateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranscodeTemplateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranscodeTemplateGroupResponse) GetBody() *GetTranscodeTemplateGroupResponseBody {
	return s.Body
}

func (s *GetTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *GetTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *GetTranscodeTemplateGroupResponse) SetStatusCode(v int32) *GetTranscodeTemplateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponse) SetBody(v *GetTranscodeTemplateGroupResponseBody) *GetTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

func (s *GetTranscodeTemplateGroupResponse) Validate() error {
	return dara.Validate(s)
}
