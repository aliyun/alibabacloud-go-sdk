// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTempFileUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTempFileUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateTempFileUploadUrlResponseBody) *CreateTempFileUploadUrlResponse
	GetBody() *CreateTempFileUploadUrlResponseBody
}

type CreateTempFileUploadUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTempFileUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTempFileUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateTempFileUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTempFileUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTempFileUploadUrlResponse) GetBody() *CreateTempFileUploadUrlResponseBody {
	return s.Body
}

func (s *CreateTempFileUploadUrlResponse) SetHeaders(v map[string]*string) *CreateTempFileUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateTempFileUploadUrlResponse) SetStatusCode(v int32) *CreateTempFileUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTempFileUploadUrlResponse) SetBody(v *CreateTempFileUploadUrlResponseBody) *CreateTempFileUploadUrlResponse {
	s.Body = v
	return s
}

func (s *CreateTempFileUploadUrlResponse) Validate() error {
	return dara.Validate(s)
}
