// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListUserKeysResponseBody) *ListUserKeysResponse
	GetBody() *ListUserKeysResponseBody
}

type ListUserKeysResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserKeysResponse) GoString() string {
	return s.String()
}

func (s *ListUserKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserKeysResponse) GetBody() *ListUserKeysResponseBody {
	return s.Body
}

func (s *ListUserKeysResponse) SetHeaders(v map[string]*string) *ListUserKeysResponse {
	s.Headers = v
	return s
}

func (s *ListUserKeysResponse) SetStatusCode(v int32) *ListUserKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserKeysResponse) SetBody(v *ListUserKeysResponseBody) *ListUserKeysResponse {
	s.Body = v
	return s
}

func (s *ListUserKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
