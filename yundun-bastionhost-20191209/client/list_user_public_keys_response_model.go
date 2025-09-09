// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPublicKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserPublicKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserPublicKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListUserPublicKeysResponseBody) *ListUserPublicKeysResponse
	GetBody() *ListUserPublicKeysResponseBody
}

type ListUserPublicKeysResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPublicKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPublicKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserPublicKeysResponse) GoString() string {
	return s.String()
}

func (s *ListUserPublicKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserPublicKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserPublicKeysResponse) GetBody() *ListUserPublicKeysResponseBody {
	return s.Body
}

func (s *ListUserPublicKeysResponse) SetHeaders(v map[string]*string) *ListUserPublicKeysResponse {
	s.Headers = v
	return s
}

func (s *ListUserPublicKeysResponse) SetStatusCode(v int32) *ListUserPublicKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPublicKeysResponse) SetBody(v *ListUserPublicKeysResponseBody) *ListUserPublicKeysResponse {
	s.Body = v
	return s
}

func (s *ListUserPublicKeysResponse) Validate() error {
	return dara.Validate(s)
}
