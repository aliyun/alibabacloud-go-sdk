// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitMultipartFileUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitMultipartFileUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitMultipartFileUploadResponse
	GetStatusCode() *int32
	SetBody(v *InitMultipartFileUploadResponseBody) *InitMultipartFileUploadResponse
	GetBody() *InitMultipartFileUploadResponseBody
}

type InitMultipartFileUploadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitMultipartFileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitMultipartFileUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s InitMultipartFileUploadResponse) GoString() string {
	return s.String()
}

func (s *InitMultipartFileUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitMultipartFileUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitMultipartFileUploadResponse) GetBody() *InitMultipartFileUploadResponseBody {
	return s.Body
}

func (s *InitMultipartFileUploadResponse) SetHeaders(v map[string]*string) *InitMultipartFileUploadResponse {
	s.Headers = v
	return s
}

func (s *InitMultipartFileUploadResponse) SetStatusCode(v int32) *InitMultipartFileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *InitMultipartFileUploadResponse) SetBody(v *InitMultipartFileUploadResponseBody) *InitMultipartFileUploadResponse {
	s.Body = v
	return s
}

func (s *InitMultipartFileUploadResponse) Validate() error {
	return dara.Validate(s)
}
