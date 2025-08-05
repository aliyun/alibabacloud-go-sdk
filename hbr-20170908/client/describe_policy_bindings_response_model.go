// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyBindingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolicyBindingsResponseBody) *DescribePolicyBindingsResponse
	GetBody() *DescribePolicyBindingsResponseBody
}

type DescribePolicyBindingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyBindingsResponse) GetBody() *DescribePolicyBindingsResponseBody {
	return s.Body
}

func (s *DescribePolicyBindingsResponse) SetHeaders(v map[string]*string) *DescribePolicyBindingsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyBindingsResponse) SetStatusCode(v int32) *DescribePolicyBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyBindingsResponse) SetBody(v *DescribePolicyBindingsResponseBody) *DescribePolicyBindingsResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyBindingsResponse) Validate() error {
	return dara.Validate(s)
}
