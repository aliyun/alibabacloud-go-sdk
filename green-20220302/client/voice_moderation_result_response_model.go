// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VoiceModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VoiceModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *VoiceModerationResultResponseBody) *VoiceModerationResultResponse
	GetBody() *VoiceModerationResultResponseBody
}

type VoiceModerationResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VoiceModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationResultResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VoiceModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VoiceModerationResultResponse) GetBody() *VoiceModerationResultResponseBody {
	return s.Body
}

func (s *VoiceModerationResultResponse) SetHeaders(v map[string]*string) *VoiceModerationResultResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationResultResponse) SetStatusCode(v int32) *VoiceModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationResultResponse) SetBody(v *VoiceModerationResultResponseBody) *VoiceModerationResultResponse {
	s.Body = v
	return s
}

func (s *VoiceModerationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
