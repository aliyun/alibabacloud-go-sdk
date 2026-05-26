// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLServiceProviderInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSAMLServiceProviderInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSAMLServiceProviderInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetSAMLServiceProviderInfoResponseBody) *GetSAMLServiceProviderInfoResponse
	GetBody() *GetSAMLServiceProviderInfoResponseBody
}

type GetSAMLServiceProviderInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSAMLServiceProviderInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSAMLServiceProviderInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLServiceProviderInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSAMLServiceProviderInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSAMLServiceProviderInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSAMLServiceProviderInfoResponse) GetBody() *GetSAMLServiceProviderInfoResponseBody {
	return s.Body
}

func (s *GetSAMLServiceProviderInfoResponse) SetHeaders(v map[string]*string) *GetSAMLServiceProviderInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSAMLServiceProviderInfoResponse) SetStatusCode(v int32) *GetSAMLServiceProviderInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSAMLServiceProviderInfoResponse) SetBody(v *GetSAMLServiceProviderInfoResponseBody) *GetSAMLServiceProviderInfoResponse {
	s.Body = v
	return s
}

func (s *GetSAMLServiceProviderInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
