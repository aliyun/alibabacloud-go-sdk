// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCredentialProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCredentialProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListCredentialProvidersResponseBody) *ListCredentialProvidersResponse
	GetBody() *ListCredentialProvidersResponseBody
}

type ListCredentialProvidersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCredentialProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCredentialProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListCredentialProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCredentialProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCredentialProvidersResponse) GetBody() *ListCredentialProvidersResponseBody {
	return s.Body
}

func (s *ListCredentialProvidersResponse) SetHeaders(v map[string]*string) *ListCredentialProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListCredentialProvidersResponse) SetStatusCode(v int32) *ListCredentialProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCredentialProvidersResponse) SetBody(v *ListCredentialProvidersResponseBody) *ListCredentialProvidersResponse {
	s.Body = v
	return s
}

func (s *ListCredentialProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
