// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadCdsFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadCdsFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadCdsFileResponse
	GetStatusCode() *int32
	SetBody(v *DownloadCdsFileResponseBody) *DownloadCdsFileResponse
	GetBody() *DownloadCdsFileResponseBody
}

type DownloadCdsFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DownloadCdsFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DownloadCdsFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadCdsFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadCdsFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadCdsFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadCdsFileResponse) GetBody() *DownloadCdsFileResponseBody {
	return s.Body
}

func (s *DownloadCdsFileResponse) SetHeaders(v map[string]*string) *DownloadCdsFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadCdsFileResponse) SetStatusCode(v int32) *DownloadCdsFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadCdsFileResponse) SetBody(v *DownloadCdsFileResponseBody) *DownloadCdsFileResponse {
	s.Body = v
	return s
}

func (s *DownloadCdsFileResponse) Validate() error {
	return dara.Validate(s)
}
