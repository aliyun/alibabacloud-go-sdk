// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDownloadInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileDownloadInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileDownloadInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetFileDownloadInfoResponseBody) *GetFileDownloadInfoResponse
	GetBody() *GetFileDownloadInfoResponseBody
}

type GetFileDownloadInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileDownloadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileDownloadInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileDownloadInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileDownloadInfoResponse) GetBody() *GetFileDownloadInfoResponseBody {
	return s.Body
}

func (s *GetFileDownloadInfoResponse) SetHeaders(v map[string]*string) *GetFileDownloadInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFileDownloadInfoResponse) SetStatusCode(v int32) *GetFileDownloadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileDownloadInfoResponse) SetBody(v *GetFileDownloadInfoResponseBody) *GetFileDownloadInfoResponse {
	s.Body = v
	return s
}

func (s *GetFileDownloadInfoResponse) Validate() error {
	return dara.Validate(s)
}
