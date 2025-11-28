// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadPrivacyKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadPrivacyKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadPrivacyKeyResponse
	GetStatusCode() *int32
	SetBody(v *DownloadPrivacyKeyResponseBody) *DownloadPrivacyKeyResponse
	GetBody() *DownloadPrivacyKeyResponseBody
}

type DownloadPrivacyKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadPrivacyKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadPrivacyKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadPrivacyKeyResponse) GoString() string {
	return s.String()
}

func (s *DownloadPrivacyKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadPrivacyKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadPrivacyKeyResponse) GetBody() *DownloadPrivacyKeyResponseBody {
	return s.Body
}

func (s *DownloadPrivacyKeyResponse) SetHeaders(v map[string]*string) *DownloadPrivacyKeyResponse {
	s.Headers = v
	return s
}

func (s *DownloadPrivacyKeyResponse) SetStatusCode(v int32) *DownloadPrivacyKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadPrivacyKeyResponse) SetBody(v *DownloadPrivacyKeyResponseBody) *DownloadPrivacyKeyResponse {
	s.Body = v
	return s
}

func (s *DownloadPrivacyKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
