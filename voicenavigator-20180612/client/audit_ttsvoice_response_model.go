// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditTTSVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuditTTSVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuditTTSVoiceResponse
	GetStatusCode() *int32
	SetBody(v *AuditTTSVoiceResponseBody) *AuditTTSVoiceResponse
	GetBody() *AuditTTSVoiceResponseBody
}

type AuditTTSVoiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditTTSVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditTTSVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s AuditTTSVoiceResponse) GoString() string {
	return s.String()
}

func (s *AuditTTSVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuditTTSVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuditTTSVoiceResponse) GetBody() *AuditTTSVoiceResponseBody {
	return s.Body
}

func (s *AuditTTSVoiceResponse) SetHeaders(v map[string]*string) *AuditTTSVoiceResponse {
	s.Headers = v
	return s
}

func (s *AuditTTSVoiceResponse) SetStatusCode(v int32) *AuditTTSVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditTTSVoiceResponse) SetBody(v *AuditTTSVoiceResponseBody) *AuditTTSVoiceResponse {
	s.Body = v
	return s
}

func (s *AuditTTSVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
