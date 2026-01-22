// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessKeysResponseBody) *ListAccessKeysResponse
	GetBody() *ListAccessKeysResponseBody
}

type ListAccessKeysResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessKeysResponse) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessKeysResponse) GetBody() *ListAccessKeysResponseBody {
	return s.Body
}

func (s *ListAccessKeysResponse) SetHeaders(v map[string]*string) *ListAccessKeysResponse {
	s.Headers = v
	return s
}

func (s *ListAccessKeysResponse) SetStatusCode(v int32) *ListAccessKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessKeysResponse) SetBody(v *ListAccessKeysResponseBody) *ListAccessKeysResponse {
	s.Body = v
	return s
}

func (s *ListAccessKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
