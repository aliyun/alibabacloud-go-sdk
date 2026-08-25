// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectorySAMLServiceProviderInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDirectorySAMLServiceProviderInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDirectorySAMLServiceProviderInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDirectorySAMLServiceProviderInfoResponseBody) *GetDirectorySAMLServiceProviderInfoResponse
	GetBody() *GetDirectorySAMLServiceProviderInfoResponseBody
}

type GetDirectorySAMLServiceProviderInfoResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectorySAMLServiceProviderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectorySAMLServiceProviderInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDirectorySAMLServiceProviderInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) GetBody() *GetDirectorySAMLServiceProviderInfoResponseBody {
	return s.Body
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) SetHeaders(v map[string]*string) *GetDirectorySAMLServiceProviderInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) SetStatusCode(v int32) *GetDirectorySAMLServiceProviderInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) SetBody(v *GetDirectorySAMLServiceProviderInfoResponseBody) *GetDirectorySAMLServiceProviderInfoResponse {
	s.Body = v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
