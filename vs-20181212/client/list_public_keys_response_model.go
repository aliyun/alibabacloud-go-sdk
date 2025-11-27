// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublicKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublicKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListPublicKeysResponseBody) *ListPublicKeysResponse
	GetBody() *ListPublicKeysResponseBody
}

type ListPublicKeysResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublicKeysResponse) GoString() string {
	return s.String()
}

func (s *ListPublicKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublicKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublicKeysResponse) GetBody() *ListPublicKeysResponseBody {
	return s.Body
}

func (s *ListPublicKeysResponse) SetHeaders(v map[string]*string) *ListPublicKeysResponse {
	s.Headers = v
	return s
}

func (s *ListPublicKeysResponse) SetStatusCode(v int32) *ListPublicKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicKeysResponse) SetBody(v *ListPublicKeysResponseBody) *ListPublicKeysResponse {
	s.Body = v
	return s
}

func (s *ListPublicKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
