// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVoiceCloneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVoiceCloneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVoiceCloneResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVoiceCloneResponseBody) *ModifyVoiceCloneResponse
	GetBody() *ModifyVoiceCloneResponseBody
}

type ModifyVoiceCloneResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVoiceCloneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVoiceCloneResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVoiceCloneResponse) GoString() string {
	return s.String()
}

func (s *ModifyVoiceCloneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVoiceCloneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVoiceCloneResponse) GetBody() *ModifyVoiceCloneResponseBody {
	return s.Body
}

func (s *ModifyVoiceCloneResponse) SetHeaders(v map[string]*string) *ModifyVoiceCloneResponse {
	s.Headers = v
	return s
}

func (s *ModifyVoiceCloneResponse) SetStatusCode(v int32) *ModifyVoiceCloneResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVoiceCloneResponse) SetBody(v *ModifyVoiceCloneResponseBody) *ModifyVoiceCloneResponse {
	s.Body = v
	return s
}

func (s *ModifyVoiceCloneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
