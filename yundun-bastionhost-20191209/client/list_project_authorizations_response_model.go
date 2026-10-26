// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectAuthorizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectAuthorizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectAuthorizationsResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectAuthorizationsResponseBody) *ListProjectAuthorizationsResponse
	GetBody() *ListProjectAuthorizationsResponseBody
}

type ListProjectAuthorizationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectAuthorizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectAuthorizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectAuthorizationsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectAuthorizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectAuthorizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectAuthorizationsResponse) GetBody() *ListProjectAuthorizationsResponseBody {
	return s.Body
}

func (s *ListProjectAuthorizationsResponse) SetHeaders(v map[string]*string) *ListProjectAuthorizationsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectAuthorizationsResponse) SetStatusCode(v int32) *ListProjectAuthorizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectAuthorizationsResponse) SetBody(v *ListProjectAuthorizationsResponseBody) *ListProjectAuthorizationsResponse {
	s.Body = v
	return s
}

func (s *ListProjectAuthorizationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
