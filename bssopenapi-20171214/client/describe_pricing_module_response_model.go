// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePricingModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePricingModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePricingModuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribePricingModuleResponseBody) *DescribePricingModuleResponse
	GetBody() *DescribePricingModuleResponseBody
}

type DescribePricingModuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePricingModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePricingModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleResponse) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePricingModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePricingModuleResponse) GetBody() *DescribePricingModuleResponseBody {
	return s.Body
}

func (s *DescribePricingModuleResponse) SetHeaders(v map[string]*string) *DescribePricingModuleResponse {
	s.Headers = v
	return s
}

func (s *DescribePricingModuleResponse) SetStatusCode(v int32) *DescribePricingModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePricingModuleResponse) SetBody(v *DescribePricingModuleResponseBody) *DescribePricingModuleResponse {
	s.Body = v
	return s
}

func (s *DescribePricingModuleResponse) Validate() error {
	return dara.Validate(s)
}
