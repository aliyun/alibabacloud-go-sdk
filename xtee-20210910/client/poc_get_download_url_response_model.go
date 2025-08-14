// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocGetDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PocGetDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PocGetDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *PocGetDownloadUrlResponseBody) *PocGetDownloadUrlResponse
	GetBody() *PocGetDownloadUrlResponseBody
}

type PocGetDownloadUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PocGetDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PocGetDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s PocGetDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *PocGetDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PocGetDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PocGetDownloadUrlResponse) GetBody() *PocGetDownloadUrlResponseBody {
	return s.Body
}

func (s *PocGetDownloadUrlResponse) SetHeaders(v map[string]*string) *PocGetDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *PocGetDownloadUrlResponse) SetStatusCode(v int32) *PocGetDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *PocGetDownloadUrlResponse) SetBody(v *PocGetDownloadUrlResponseBody) *PocGetDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *PocGetDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}
