// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestUploadTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RequestUploadTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RequestUploadTokenResponse
	GetStatusCode() *int32
	SetBody(v *RequestUploadTokenResponseBody) *RequestUploadTokenResponse
	GetBody() *RequestUploadTokenResponseBody
}

type RequestUploadTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestUploadTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestUploadTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s RequestUploadTokenResponse) GoString() string {
	return s.String()
}

func (s *RequestUploadTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RequestUploadTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RequestUploadTokenResponse) GetBody() *RequestUploadTokenResponseBody {
	return s.Body
}

func (s *RequestUploadTokenResponse) SetHeaders(v map[string]*string) *RequestUploadTokenResponse {
	s.Headers = v
	return s
}

func (s *RequestUploadTokenResponse) SetStatusCode(v int32) *RequestUploadTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestUploadTokenResponse) SetBody(v *RequestUploadTokenResponseBody) *RequestUploadTokenResponse {
	s.Body = v
	return s
}

func (s *RequestUploadTokenResponse) Validate() error {
	return dara.Validate(s)
}
