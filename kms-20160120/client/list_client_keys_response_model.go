// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListClientKeysResponseBody) *ListClientKeysResponse
	GetBody() *ListClientKeysResponseBody
}

type ListClientKeysResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientKeysResponse) GoString() string {
	return s.String()
}

func (s *ListClientKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientKeysResponse) GetBody() *ListClientKeysResponseBody {
	return s.Body
}

func (s *ListClientKeysResponse) SetHeaders(v map[string]*string) *ListClientKeysResponse {
	s.Headers = v
	return s
}

func (s *ListClientKeysResponse) SetStatusCode(v int32) *ListClientKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientKeysResponse) SetBody(v *ListClientKeysResponseBody) *ListClientKeysResponse {
	s.Body = v
	return s
}

func (s *ListClientKeysResponse) Validate() error {
	return dara.Validate(s)
}
