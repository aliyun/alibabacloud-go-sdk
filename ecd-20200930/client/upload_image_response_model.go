// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadImageResponse
	GetStatusCode() *int32
	SetBody(v *UploadImageResponseBody) *UploadImageResponse
	GetBody() *UploadImageResponseBody
}

type UploadImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadImageResponse) GoString() string {
	return s.String()
}

func (s *UploadImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadImageResponse) GetBody() *UploadImageResponseBody {
	return s.Body
}

func (s *UploadImageResponse) SetHeaders(v map[string]*string) *UploadImageResponse {
	s.Headers = v
	return s
}

func (s *UploadImageResponse) SetStatusCode(v int32) *UploadImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadImageResponse) SetBody(v *UploadImageResponseBody) *UploadImageResponse {
	s.Body = v
	return s
}

func (s *UploadImageResponse) Validate() error {
	return dara.Validate(s)
}
