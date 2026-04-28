// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetDownloadUrlResponseBody) *GetDownloadUrlResponse
	GetBody() *GetDownloadUrlResponseBody
}

type GetDownloadUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDownloadUrlResponse) GetBody() *GetDownloadUrlResponseBody {
	return s.Body
}

func (s *GetDownloadUrlResponse) SetHeaders(v map[string]*string) *GetDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDownloadUrlResponse) SetStatusCode(v int32) *GetDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDownloadUrlResponse) SetBody(v *GetDownloadUrlResponseBody) *GetDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
