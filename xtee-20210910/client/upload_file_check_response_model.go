// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadFileCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadFileCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadFileCheckResponse
	GetStatusCode() *int32
	SetBody(v *UploadFileCheckResponseBody) *UploadFileCheckResponse
	GetBody() *UploadFileCheckResponseBody
}

type UploadFileCheckResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadFileCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadFileCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadFileCheckResponse) GoString() string {
	return s.String()
}

func (s *UploadFileCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadFileCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadFileCheckResponse) GetBody() *UploadFileCheckResponseBody {
	return s.Body
}

func (s *UploadFileCheckResponse) SetHeaders(v map[string]*string) *UploadFileCheckResponse {
	s.Headers = v
	return s
}

func (s *UploadFileCheckResponse) SetStatusCode(v int32) *UploadFileCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadFileCheckResponse) SetBody(v *UploadFileCheckResponseBody) *UploadFileCheckResponse {
	s.Body = v
	return s
}

func (s *UploadFileCheckResponse) Validate() error {
	return dara.Validate(s)
}
