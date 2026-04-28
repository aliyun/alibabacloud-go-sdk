// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityToBenefitPkgMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIdentityToBenefitPkgMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIdentityToBenefitPkgMappingResponse
	GetStatusCode() *int32
}

type UpdateIdentityToBenefitPkgMappingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateIdentityToBenefitPkgMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityToBenefitPkgMappingResponse) GoString() string {
	return s.String()
}

func (s *UpdateIdentityToBenefitPkgMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIdentityToBenefitPkgMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIdentityToBenefitPkgMappingResponse) SetHeaders(v map[string]*string) *UpdateIdentityToBenefitPkgMappingResponse {
	s.Headers = v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingResponse) SetStatusCode(v int32) *UpdateIdentityToBenefitPkgMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingResponse) Validate() error {
	return dara.Validate(s)
}
