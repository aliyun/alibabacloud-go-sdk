// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePoliciesResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *DescribePoliciesResponse
	GetBody() map[string]interface{}
}

type DescribePoliciesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribePoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePoliciesResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *DescribePoliciesResponse) SetHeaders(v map[string]*string) *DescribePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribePoliciesResponse) SetStatusCode(v int32) *DescribePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePoliciesResponse) SetBody(v map[string]interface{}) *DescribePoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribePoliciesResponse) Validate() error {
	return dara.Validate(s)
}
