// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrepareUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrepareUploadResponse
	GetStatusCode() *int32
	SetBody(v *PrepareUploadResponseBody) *PrepareUploadResponse
	GetBody() *PrepareUploadResponseBody
}

type PrepareUploadResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrepareUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrepareUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s PrepareUploadResponse) GoString() string {
	return s.String()
}

func (s *PrepareUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrepareUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrepareUploadResponse) GetBody() *PrepareUploadResponseBody {
	return s.Body
}

func (s *PrepareUploadResponse) SetHeaders(v map[string]*string) *PrepareUploadResponse {
	s.Headers = v
	return s
}

func (s *PrepareUploadResponse) SetStatusCode(v int32) *PrepareUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *PrepareUploadResponse) SetBody(v *PrepareUploadResponseBody) *PrepareUploadResponse {
	s.Body = v
	return s
}

func (s *PrepareUploadResponse) Validate() error {
	return dara.Validate(s)
}
