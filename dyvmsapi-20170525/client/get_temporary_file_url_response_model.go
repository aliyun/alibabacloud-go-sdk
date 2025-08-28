// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemporaryFileUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTemporaryFileUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTemporaryFileUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetTemporaryFileUrlResponseBody) *GetTemporaryFileUrlResponse
	GetBody() *GetTemporaryFileUrlResponseBody
}

type GetTemporaryFileUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemporaryFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemporaryFileUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTemporaryFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTemporaryFileUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTemporaryFileUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTemporaryFileUrlResponse) GetBody() *GetTemporaryFileUrlResponseBody {
	return s.Body
}

func (s *GetTemporaryFileUrlResponse) SetHeaders(v map[string]*string) *GetTemporaryFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTemporaryFileUrlResponse) SetStatusCode(v int32) *GetTemporaryFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemporaryFileUrlResponse) SetBody(v *GetTemporaryFileUrlResponseBody) *GetTemporaryFileUrlResponse {
	s.Body = v
	return s
}

func (s *GetTemporaryFileUrlResponse) Validate() error {
	return dara.Validate(s)
}
