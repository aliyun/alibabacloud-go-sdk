// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAPIKeyCredentialProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAPIKeyCredentialProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAPIKeyCredentialProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListAPIKeyCredentialProvidersResponseBody) *ListAPIKeyCredentialProvidersResponse
	GetBody() *ListAPIKeyCredentialProvidersResponseBody
}

type ListAPIKeyCredentialProvidersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAPIKeyCredentialProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAPIKeyCredentialProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAPIKeyCredentialProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListAPIKeyCredentialProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAPIKeyCredentialProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAPIKeyCredentialProvidersResponse) GetBody() *ListAPIKeyCredentialProvidersResponseBody {
	return s.Body
}

func (s *ListAPIKeyCredentialProvidersResponse) SetHeaders(v map[string]*string) *ListAPIKeyCredentialProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponse) SetStatusCode(v int32) *ListAPIKeyCredentialProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponse) SetBody(v *ListAPIKeyCredentialProvidersResponseBody) *ListAPIKeyCredentialProvidersResponse {
	s.Body = v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
