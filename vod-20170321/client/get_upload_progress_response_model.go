// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadProgressResponseBody) *GetUploadProgressResponse
	GetBody() *GetUploadProgressResponseBody
}

type GetUploadProgressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadProgressResponse) GoString() string {
	return s.String()
}

func (s *GetUploadProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadProgressResponse) GetBody() *GetUploadProgressResponseBody {
	return s.Body
}

func (s *GetUploadProgressResponse) SetHeaders(v map[string]*string) *GetUploadProgressResponse {
	s.Headers = v
	return s
}

func (s *GetUploadProgressResponse) SetStatusCode(v int32) *GetUploadProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadProgressResponse) SetBody(v *GetUploadProgressResponseBody) *GetUploadProgressResponse {
	s.Body = v
	return s
}

func (s *GetUploadProgressResponse) Validate() error {
	return dara.Validate(s)
}
