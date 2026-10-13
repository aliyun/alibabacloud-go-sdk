// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogviewDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogviewDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogviewDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetLogviewDownloadUrlResponseBody) *GetLogviewDownloadUrlResponse
	GetBody() *GetLogviewDownloadUrlResponseBody
}

type GetLogviewDownloadUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogviewDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogviewDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogviewDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetLogviewDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogviewDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogviewDownloadUrlResponse) GetBody() *GetLogviewDownloadUrlResponseBody {
	return s.Body
}

func (s *GetLogviewDownloadUrlResponse) SetHeaders(v map[string]*string) *GetLogviewDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetLogviewDownloadUrlResponse) SetStatusCode(v int32) *GetLogviewDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogviewDownloadUrlResponse) SetBody(v *GetLogviewDownloadUrlResponseBody) *GetLogviewDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetLogviewDownloadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
