// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScriptVoiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScriptVoiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScriptVoiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScriptVoiceConfigResponseBody) *ModifyScriptVoiceConfigResponse
	GetBody() *ModifyScriptVoiceConfigResponseBody
}

type ModifyScriptVoiceConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScriptVoiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScriptVoiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScriptVoiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyScriptVoiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScriptVoiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScriptVoiceConfigResponse) GetBody() *ModifyScriptVoiceConfigResponseBody {
	return s.Body
}

func (s *ModifyScriptVoiceConfigResponse) SetHeaders(v map[string]*string) *ModifyScriptVoiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyScriptVoiceConfigResponse) SetStatusCode(v int32) *ModifyScriptVoiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScriptVoiceConfigResponse) SetBody(v *ModifyScriptVoiceConfigResponseBody) *ModifyScriptVoiceConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyScriptVoiceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
