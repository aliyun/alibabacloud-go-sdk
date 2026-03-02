// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDependLibrarysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDependLibrarysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDependLibrarysResponse
	GetStatusCode() *int32
	SetBody(v *ListDependLibrarysResponseBody) *ListDependLibrarysResponse
	GetBody() *ListDependLibrarysResponseBody
}

type ListDependLibrarysResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDependLibrarysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDependLibrarysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDependLibrarysResponse) GoString() string {
	return s.String()
}

func (s *ListDependLibrarysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDependLibrarysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDependLibrarysResponse) GetBody() *ListDependLibrarysResponseBody {
	return s.Body
}

func (s *ListDependLibrarysResponse) SetHeaders(v map[string]*string) *ListDependLibrarysResponse {
	s.Headers = v
	return s
}

func (s *ListDependLibrarysResponse) SetStatusCode(v int32) *ListDependLibrarysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDependLibrarysResponse) SetBody(v *ListDependLibrarysResponseBody) *ListDependLibrarysResponse {
	s.Body = v
	return s
}

func (s *ListDependLibrarysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
