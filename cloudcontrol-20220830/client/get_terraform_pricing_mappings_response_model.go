// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerraformPricingMappingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTerraformPricingMappingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTerraformPricingMappingsResponse
	GetStatusCode() *int32
	SetBody(v *GetTerraformPricingMappingsResponseBody) *GetTerraformPricingMappingsResponse
	GetBody() *GetTerraformPricingMappingsResponseBody
}

type GetTerraformPricingMappingsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTerraformPricingMappingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTerraformPricingMappingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformPricingMappingsResponse) GoString() string {
	return s.String()
}

func (s *GetTerraformPricingMappingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTerraformPricingMappingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTerraformPricingMappingsResponse) GetBody() *GetTerraformPricingMappingsResponseBody {
	return s.Body
}

func (s *GetTerraformPricingMappingsResponse) SetHeaders(v map[string]*string) *GetTerraformPricingMappingsResponse {
	s.Headers = v
	return s
}

func (s *GetTerraformPricingMappingsResponse) SetStatusCode(v int32) *GetTerraformPricingMappingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTerraformPricingMappingsResponse) SetBody(v *GetTerraformPricingMappingsResponseBody) *GetTerraformPricingMappingsResponse {
	s.Body = v
	return s
}

func (s *GetTerraformPricingMappingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
