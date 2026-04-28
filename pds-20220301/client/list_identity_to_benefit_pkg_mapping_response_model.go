// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityToBenefitPkgMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdentityToBenefitPkgMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdentityToBenefitPkgMappingResponse
	GetStatusCode() *int32
	SetBody(v *ListIdentityToBenefitPkgMappingResponseBody) *ListIdentityToBenefitPkgMappingResponse
	GetBody() *ListIdentityToBenefitPkgMappingResponseBody
}

type ListIdentityToBenefitPkgMappingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdentityToBenefitPkgMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentityToBenefitPkgMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityToBenefitPkgMappingResponse) GoString() string {
	return s.String()
}

func (s *ListIdentityToBenefitPkgMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdentityToBenefitPkgMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdentityToBenefitPkgMappingResponse) GetBody() *ListIdentityToBenefitPkgMappingResponseBody {
	return s.Body
}

func (s *ListIdentityToBenefitPkgMappingResponse) SetHeaders(v map[string]*string) *ListIdentityToBenefitPkgMappingResponse {
	s.Headers = v
	return s
}

func (s *ListIdentityToBenefitPkgMappingResponse) SetStatusCode(v int32) *ListIdentityToBenefitPkgMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentityToBenefitPkgMappingResponse) SetBody(v *ListIdentityToBenefitPkgMappingResponseBody) *ListIdentityToBenefitPkgMappingResponse {
	s.Body = v
	return s
}

func (s *ListIdentityToBenefitPkgMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
