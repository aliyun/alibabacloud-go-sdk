// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadResponse
	GetStatusCode() *int32
	SetBody(v *UploadResponseBody) *UploadResponse
	GetBody() *UploadResponseBody
}

type UploadResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadResponse) GoString() string {
	return s.String()
}

func (s *UploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadResponse) GetBody() *UploadResponseBody {
	return s.Body
}

func (s *UploadResponse) SetHeaders(v map[string]*string) *UploadResponse {
	s.Headers = v
	return s
}

func (s *UploadResponse) SetStatusCode(v int32) *UploadResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadResponse) SetBody(v *UploadResponseBody) *UploadResponse {
	s.Body = v
	return s
}

func (s *UploadResponse) Validate() error {
	return dara.Validate(s)
}
