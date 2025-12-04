// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVoiceFileAuditInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVoiceFileAuditInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVoiceFileAuditInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryVoiceFileAuditInfoResponseBody) *QueryVoiceFileAuditInfoResponse
	GetBody() *QueryVoiceFileAuditInfoResponseBody
}

type QueryVoiceFileAuditInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVoiceFileAuditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVoiceFileAuditInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVoiceFileAuditInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryVoiceFileAuditInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVoiceFileAuditInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVoiceFileAuditInfoResponse) GetBody() *QueryVoiceFileAuditInfoResponseBody {
	return s.Body
}

func (s *QueryVoiceFileAuditInfoResponse) SetHeaders(v map[string]*string) *QueryVoiceFileAuditInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryVoiceFileAuditInfoResponse) SetStatusCode(v int32) *QueryVoiceFileAuditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponse) SetBody(v *QueryVoiceFileAuditInfoResponseBody) *QueryVoiceFileAuditInfoResponse {
	s.Body = v
	return s
}

func (s *QueryVoiceFileAuditInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
