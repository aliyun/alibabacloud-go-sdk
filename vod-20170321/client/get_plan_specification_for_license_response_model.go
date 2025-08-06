// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlanSpecificationForLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPlanSpecificationForLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPlanSpecificationForLicenseResponse
	GetStatusCode() *int32
	SetBody(v *GetPlanSpecificationForLicenseResponseBody) *GetPlanSpecificationForLicenseResponse
	GetBody() *GetPlanSpecificationForLicenseResponseBody
}

type GetPlanSpecificationForLicenseResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPlanSpecificationForLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPlanSpecificationForLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPlanSpecificationForLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetPlanSpecificationForLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPlanSpecificationForLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPlanSpecificationForLicenseResponse) GetBody() *GetPlanSpecificationForLicenseResponseBody {
	return s.Body
}

func (s *GetPlanSpecificationForLicenseResponse) SetHeaders(v map[string]*string) *GetPlanSpecificationForLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetPlanSpecificationForLicenseResponse) SetStatusCode(v int32) *GetPlanSpecificationForLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPlanSpecificationForLicenseResponse) SetBody(v *GetPlanSpecificationForLicenseResponseBody) *GetPlanSpecificationForLicenseResponse {
	s.Body = v
	return s
}

func (s *GetPlanSpecificationForLicenseResponse) Validate() error {
	return dara.Validate(s)
}
