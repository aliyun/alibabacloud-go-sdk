// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnrolledAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnrolledAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnrolledAccountsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnrolledAccountsResponseBody) *ListEnrolledAccountsResponse
	GetBody() *ListEnrolledAccountsResponseBody
}

type ListEnrolledAccountsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnrolledAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnrolledAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnrolledAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnrolledAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnrolledAccountsResponse) GetBody() *ListEnrolledAccountsResponseBody {
	return s.Body
}

func (s *ListEnrolledAccountsResponse) SetHeaders(v map[string]*string) *ListEnrolledAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListEnrolledAccountsResponse) SetStatusCode(v int32) *ListEnrolledAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnrolledAccountsResponse) SetBody(v *ListEnrolledAccountsResponseBody) *ListEnrolledAccountsResponse {
	s.Body = v
	return s
}

func (s *ListEnrolledAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
