// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyPriorUsedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyPriorUsedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyPriorUsedResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolicyPriorUsedResponseBody) *DescribePolicyPriorUsedResponse
	GetBody() *DescribePolicyPriorUsedResponseBody
}

type DescribePolicyPriorUsedResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyPriorUsedResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyPriorUsedResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyPriorUsedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyPriorUsedResponse) GetBody() *DescribePolicyPriorUsedResponseBody {
	return s.Body
}

func (s *DescribePolicyPriorUsedResponse) SetHeaders(v map[string]*string) *DescribePolicyPriorUsedResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyPriorUsedResponse) SetStatusCode(v int32) *DescribePolicyPriorUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyPriorUsedResponse) SetBody(v *DescribePolicyPriorUsedResponseBody) *DescribePolicyPriorUsedResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyPriorUsedResponse) Validate() error {
	return dara.Validate(s)
}
