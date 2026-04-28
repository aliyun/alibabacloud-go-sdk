// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityToBenefitPkgMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIdentityToBenefitPkgMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIdentityToBenefitPkgMappingResponse
	GetStatusCode() *int32
}

type CreateIdentityToBenefitPkgMappingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateIdentityToBenefitPkgMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityToBenefitPkgMappingResponse) GoString() string {
	return s.String()
}

func (s *CreateIdentityToBenefitPkgMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIdentityToBenefitPkgMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIdentityToBenefitPkgMappingResponse) SetHeaders(v map[string]*string) *CreateIdentityToBenefitPkgMappingResponse {
	s.Headers = v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingResponse) SetStatusCode(v int32) *CreateIdentityToBenefitPkgMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingResponse) Validate() error {
	return dara.Validate(s)
}
