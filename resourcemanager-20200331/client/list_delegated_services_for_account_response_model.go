// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDelegatedServicesForAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDelegatedServicesForAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDelegatedServicesForAccountResponse
	GetStatusCode() *int32
	SetBody(v *ListDelegatedServicesForAccountResponseBody) *ListDelegatedServicesForAccountResponse
	GetBody() *ListDelegatedServicesForAccountResponseBody
}

type ListDelegatedServicesForAccountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDelegatedServicesForAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDelegatedServicesForAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedServicesForAccountResponse) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDelegatedServicesForAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDelegatedServicesForAccountResponse) GetBody() *ListDelegatedServicesForAccountResponseBody {
	return s.Body
}

func (s *ListDelegatedServicesForAccountResponse) SetHeaders(v map[string]*string) *ListDelegatedServicesForAccountResponse {
	s.Headers = v
	return s
}

func (s *ListDelegatedServicesForAccountResponse) SetStatusCode(v int32) *ListDelegatedServicesForAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDelegatedServicesForAccountResponse) SetBody(v *ListDelegatedServicesForAccountResponseBody) *ListDelegatedServicesForAccountResponse {
	s.Body = v
	return s
}

func (s *ListDelegatedServicesForAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
