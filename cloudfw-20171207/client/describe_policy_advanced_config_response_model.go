// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyAdvancedConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyAdvancedConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyAdvancedConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolicyAdvancedConfigResponseBody) *DescribePolicyAdvancedConfigResponse
	GetBody() *DescribePolicyAdvancedConfigResponseBody
}

type DescribePolicyAdvancedConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyAdvancedConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyAdvancedConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyAdvancedConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyAdvancedConfigResponse) GetBody() *DescribePolicyAdvancedConfigResponseBody {
	return s.Body
}

func (s *DescribePolicyAdvancedConfigResponse) SetHeaders(v map[string]*string) *DescribePolicyAdvancedConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyAdvancedConfigResponse) SetStatusCode(v int32) *DescribePolicyAdvancedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyAdvancedConfigResponse) SetBody(v *DescribePolicyAdvancedConfigResponseBody) *DescribePolicyAdvancedConfigResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyAdvancedConfigResponse) Validate() error {
	return dara.Validate(s)
}
