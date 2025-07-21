// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolicyResponseBody) *DescribePolicyResponse
	GetBody() *DescribePolicyResponseBody
}

type DescribePolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyResponse) GetBody() *DescribePolicyResponseBody {
	return s.Body
}

func (s *DescribePolicyResponse) SetHeaders(v map[string]*string) *DescribePolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyResponse) SetStatusCode(v int32) *DescribePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyResponse) SetBody(v *DescribePolicyResponseBody) *DescribePolicyResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyResponse) Validate() error {
	return dara.Validate(s)
}
