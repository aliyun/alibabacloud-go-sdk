// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTempDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTempDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTempDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetTempDownloadUrlResponseBody) *GetTempDownloadUrlResponse
	GetBody() *GetTempDownloadUrlResponseBody
}

type GetTempDownloadUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTempDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTempDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTempDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTempDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTempDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTempDownloadUrlResponse) GetBody() *GetTempDownloadUrlResponseBody {
	return s.Body
}

func (s *GetTempDownloadUrlResponse) SetHeaders(v map[string]*string) *GetTempDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTempDownloadUrlResponse) SetStatusCode(v int32) *GetTempDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTempDownloadUrlResponse) SetBody(v *GetTempDownloadUrlResponseBody) *GetTempDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetTempDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
