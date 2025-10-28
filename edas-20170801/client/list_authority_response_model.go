// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorityResponseBody) *ListAuthorityResponse
	GetBody() *ListAuthorityResponseBody
}

type ListAuthorityResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorityResponse) GetBody() *ListAuthorityResponseBody {
	return s.Body
}

func (s *ListAuthorityResponse) SetHeaders(v map[string]*string) *ListAuthorityResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorityResponse) SetStatusCode(v int32) *ListAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorityResponse) SetBody(v *ListAuthorityResponseBody) *ListAuthorityResponse {
	s.Body = v
	return s
}

func (s *ListAuthorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
