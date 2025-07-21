// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVersionDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVersionDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVersionDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetVersionDownloadUrlResponseBody) *GetVersionDownloadUrlResponse
	GetBody() *GetVersionDownloadUrlResponseBody
}

type GetVersionDownloadUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVersionDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVersionDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVersionDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetVersionDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVersionDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVersionDownloadUrlResponse) GetBody() *GetVersionDownloadUrlResponseBody {
	return s.Body
}

func (s *GetVersionDownloadUrlResponse) SetHeaders(v map[string]*string) *GetVersionDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetVersionDownloadUrlResponse) SetStatusCode(v int32) *GetVersionDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVersionDownloadUrlResponse) SetBody(v *GetVersionDownloadUrlResponseBody) *GetVersionDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *GetVersionDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}
