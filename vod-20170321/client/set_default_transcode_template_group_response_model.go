// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultTranscodeTemplateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultTranscodeTemplateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultTranscodeTemplateGroupResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultTranscodeTemplateGroupResponseBody) *SetDefaultTranscodeTemplateGroupResponse
	GetBody() *SetDefaultTranscodeTemplateGroupResponseBody
}

type SetDefaultTranscodeTemplateGroupResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultTranscodeTemplateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultTranscodeTemplateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultTranscodeTemplateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultTranscodeTemplateGroupResponse) GetBody() *SetDefaultTranscodeTemplateGroupResponseBody {
	return s.Body
}

func (s *SetDefaultTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *SetDefaultTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultTranscodeTemplateGroupResponse) SetStatusCode(v int32) *SetDefaultTranscodeTemplateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultTranscodeTemplateGroupResponse) SetBody(v *SetDefaultTranscodeTemplateGroupResponseBody) *SetDefaultTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

func (s *SetDefaultTranscodeTemplateGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
