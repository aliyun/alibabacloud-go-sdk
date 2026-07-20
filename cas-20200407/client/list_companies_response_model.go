// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompaniesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCompaniesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCompaniesResponse
	GetStatusCode() *int32
	SetBody(v *ListCompaniesResponseBody) *ListCompaniesResponse
	GetBody() *ListCompaniesResponseBody
}

type ListCompaniesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCompaniesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCompaniesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCompaniesResponse) GoString() string {
	return s.String()
}

func (s *ListCompaniesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCompaniesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCompaniesResponse) GetBody() *ListCompaniesResponseBody {
	return s.Body
}

func (s *ListCompaniesResponse) SetHeaders(v map[string]*string) *ListCompaniesResponse {
	s.Headers = v
	return s
}

func (s *ListCompaniesResponse) SetStatusCode(v int32) *ListCompaniesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCompaniesResponse) SetBody(v *ListCompaniesResponseBody) *ListCompaniesResponse {
	s.Body = v
	return s
}

func (s *ListCompaniesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
