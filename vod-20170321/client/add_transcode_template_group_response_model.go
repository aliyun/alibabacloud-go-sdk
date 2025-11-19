// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTranscodeTemplateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTranscodeTemplateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTranscodeTemplateGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddTranscodeTemplateGroupResponseBody) *AddTranscodeTemplateGroupResponse
	GetBody() *AddTranscodeTemplateGroupResponseBody
}

type AddTranscodeTemplateGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTranscodeTemplateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *AddTranscodeTemplateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTranscodeTemplateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTranscodeTemplateGroupResponse) GetBody() *AddTranscodeTemplateGroupResponseBody {
	return s.Body
}

func (s *AddTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *AddTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *AddTranscodeTemplateGroupResponse) SetStatusCode(v int32) *AddTranscodeTemplateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTranscodeTemplateGroupResponse) SetBody(v *AddTranscodeTemplateGroupResponseBody) *AddTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

func (s *AddTranscodeTemplateGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
