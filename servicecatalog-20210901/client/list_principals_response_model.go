// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrincipalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrincipalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrincipalsResponse
	GetStatusCode() *int32
	SetBody(v *ListPrincipalsResponseBody) *ListPrincipalsResponse
	GetBody() *ListPrincipalsResponseBody
}

type ListPrincipalsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrincipalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrincipalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrincipalsResponse) GoString() string {
	return s.String()
}

func (s *ListPrincipalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrincipalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrincipalsResponse) GetBody() *ListPrincipalsResponseBody {
	return s.Body
}

func (s *ListPrincipalsResponse) SetHeaders(v map[string]*string) *ListPrincipalsResponse {
	s.Headers = v
	return s
}

func (s *ListPrincipalsResponse) SetStatusCode(v int32) *ListPrincipalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrincipalsResponse) SetBody(v *ListPrincipalsResponseBody) *ListPrincipalsResponse {
	s.Body = v
	return s
}

func (s *ListPrincipalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
