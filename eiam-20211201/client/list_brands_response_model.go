// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrandsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBrandsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBrandsResponse
	GetStatusCode() *int32
	SetBody(v *ListBrandsResponseBody) *ListBrandsResponse
	GetBody() *ListBrandsResponseBody
}

type ListBrandsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBrandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBrandsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBrandsResponse) GoString() string {
	return s.String()
}

func (s *ListBrandsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBrandsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBrandsResponse) GetBody() *ListBrandsResponseBody {
	return s.Body
}

func (s *ListBrandsResponse) SetHeaders(v map[string]*string) *ListBrandsResponse {
	s.Headers = v
	return s
}

func (s *ListBrandsResponse) SetStatusCode(v int32) *ListBrandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBrandsResponse) SetBody(v *ListBrandsResponseBody) *ListBrandsResponse {
	s.Body = v
	return s
}

func (s *ListBrandsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
