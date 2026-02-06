// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptVoiceConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScriptVoiceConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScriptVoiceConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListScriptVoiceConfigsResponseBody) *ListScriptVoiceConfigsResponse
	GetBody() *ListScriptVoiceConfigsResponseBody
}

type ListScriptVoiceConfigsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScriptVoiceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScriptVoiceConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVoiceConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListScriptVoiceConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScriptVoiceConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScriptVoiceConfigsResponse) GetBody() *ListScriptVoiceConfigsResponseBody {
	return s.Body
}

func (s *ListScriptVoiceConfigsResponse) SetHeaders(v map[string]*string) *ListScriptVoiceConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListScriptVoiceConfigsResponse) SetStatusCode(v int32) *ListScriptVoiceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScriptVoiceConfigsResponse) SetBody(v *ListScriptVoiceConfigsResponseBody) *ListScriptVoiceConfigsResponse {
	s.Body = v
	return s
}

func (s *ListScriptVoiceConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
