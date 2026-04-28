// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DownloadFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DownloadFileResponse
	GetStatusCode() *int32
}

type DownloadFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DownloadFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DownloadFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DownloadFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DownloadFileResponse) SetHeaders(v map[string]*string) *DownloadFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadFileResponse) SetStatusCode(v int32) *DownloadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DownloadFileResponse) Validate() error {
	return dara.Validate(s)
}
