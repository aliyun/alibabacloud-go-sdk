// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKmsKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKmsKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKmsKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListKmsKeysResponseBody) *ListKmsKeysResponse
	GetBody() *ListKmsKeysResponseBody
}

type ListKmsKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKmsKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKmsKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKmsKeysResponse) GoString() string {
	return s.String()
}

func (s *ListKmsKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKmsKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKmsKeysResponse) GetBody() *ListKmsKeysResponseBody {
	return s.Body
}

func (s *ListKmsKeysResponse) SetHeaders(v map[string]*string) *ListKmsKeysResponse {
	s.Headers = v
	return s
}

func (s *ListKmsKeysResponse) SetStatusCode(v int32) *ListKmsKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKmsKeysResponse) SetBody(v *ListKmsKeysResponseBody) *ListKmsKeysResponse {
	s.Body = v
	return s
}

func (s *ListKmsKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
