// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaAuditAudioResultDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaAuditAudioResultDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaAuditAudioResultDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaAuditAudioResultDetailResponseBody) *GetMediaAuditAudioResultDetailResponse
	GetBody() *GetMediaAuditAudioResultDetailResponseBody
}

type GetMediaAuditAudioResultDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaAuditAudioResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaAuditAudioResultDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaAuditAudioResultDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaAuditAudioResultDetailResponse) GetBody() *GetMediaAuditAudioResultDetailResponseBody {
	return s.Body
}

func (s *GetMediaAuditAudioResultDetailResponse) SetHeaders(v map[string]*string) *GetMediaAuditAudioResultDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponse) SetStatusCode(v int32) *GetMediaAuditAudioResultDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponse) SetBody(v *GetMediaAuditAudioResultDetailResponseBody) *GetMediaAuditAudioResultDetailResponse {
	s.Body = v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
