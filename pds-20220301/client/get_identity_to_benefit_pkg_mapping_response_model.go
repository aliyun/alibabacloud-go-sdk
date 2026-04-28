// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityToBenefitPkgMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityToBenefitPkgMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityToBenefitPkgMappingResponse
	GetStatusCode() *int32
	SetBody(v *IdentityToBenefitPkgMapping) *GetIdentityToBenefitPkgMappingResponse
	GetBody() *IdentityToBenefitPkgMapping
}

type GetIdentityToBenefitPkgMappingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IdentityToBenefitPkgMapping `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityToBenefitPkgMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityToBenefitPkgMappingResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityToBenefitPkgMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityToBenefitPkgMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityToBenefitPkgMappingResponse) GetBody() *IdentityToBenefitPkgMapping {
	return s.Body
}

func (s *GetIdentityToBenefitPkgMappingResponse) SetHeaders(v map[string]*string) *GetIdentityToBenefitPkgMappingResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityToBenefitPkgMappingResponse) SetStatusCode(v int32) *GetIdentityToBenefitPkgMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityToBenefitPkgMappingResponse) SetBody(v *IdentityToBenefitPkgMapping) *GetIdentityToBenefitPkgMappingResponse {
	s.Body = v
	return s
}

func (s *GetIdentityToBenefitPkgMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
