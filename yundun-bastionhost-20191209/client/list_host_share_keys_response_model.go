// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostShareKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHostShareKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHostShareKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListHostShareKeysResponseBody) *ListHostShareKeysResponse
	GetBody() *ListHostShareKeysResponseBody
}

type ListHostShareKeysResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostShareKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHostShareKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHostShareKeysResponse) GoString() string {
	return s.String()
}

func (s *ListHostShareKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHostShareKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHostShareKeysResponse) GetBody() *ListHostShareKeysResponseBody {
	return s.Body
}

func (s *ListHostShareKeysResponse) SetHeaders(v map[string]*string) *ListHostShareKeysResponse {
	s.Headers = v
	return s
}

func (s *ListHostShareKeysResponse) SetStatusCode(v int32) *ListHostShareKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostShareKeysResponse) SetBody(v *ListHostShareKeysResponseBody) *ListHostShareKeysResponse {
	s.Body = v
	return s
}

func (s *ListHostShareKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
