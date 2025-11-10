// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDelegatedAdministratorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDelegatedAdministratorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDelegatedAdministratorsResponse
	GetStatusCode() *int32
	SetBody(v *ListDelegatedAdministratorsResponseBody) *ListDelegatedAdministratorsResponse
	GetBody() *ListDelegatedAdministratorsResponseBody
}

type ListDelegatedAdministratorsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDelegatedAdministratorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDelegatedAdministratorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedAdministratorsResponse) GoString() string {
	return s.String()
}

func (s *ListDelegatedAdministratorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDelegatedAdministratorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDelegatedAdministratorsResponse) GetBody() *ListDelegatedAdministratorsResponseBody {
	return s.Body
}

func (s *ListDelegatedAdministratorsResponse) SetHeaders(v map[string]*string) *ListDelegatedAdministratorsResponse {
	s.Headers = v
	return s
}

func (s *ListDelegatedAdministratorsResponse) SetStatusCode(v int32) *ListDelegatedAdministratorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDelegatedAdministratorsResponse) SetBody(v *ListDelegatedAdministratorsResponseBody) *ListDelegatedAdministratorsResponse {
	s.Body = v
	return s
}

func (s *ListDelegatedAdministratorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
