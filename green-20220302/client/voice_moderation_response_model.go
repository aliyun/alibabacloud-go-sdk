// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VoiceModerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VoiceModerationResponse
	GetStatusCode() *int32
	SetBody(v *VoiceModerationResponseBody) *VoiceModerationResponse
	GetBody() *VoiceModerationResponseBody
}

type VoiceModerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceModerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VoiceModerationResponse) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VoiceModerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VoiceModerationResponse) GetBody() *VoiceModerationResponseBody {
	return s.Body
}

func (s *VoiceModerationResponse) SetHeaders(v map[string]*string) *VoiceModerationResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationResponse) SetStatusCode(v int32) *VoiceModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationResponse) SetBody(v *VoiceModerationResponseBody) *VoiceModerationResponse {
	s.Body = v
	return s
}

func (s *VoiceModerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
