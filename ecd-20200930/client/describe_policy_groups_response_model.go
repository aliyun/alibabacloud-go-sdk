// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolicyGroupsResponseBody) *DescribePolicyGroupsResponse
	GetBody() *DescribePolicyGroupsResponseBody
}

type DescribePolicyGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyGroupsResponse) GetBody() *DescribePolicyGroupsResponseBody {
	return s.Body
}

func (s *DescribePolicyGroupsResponse) SetHeaders(v map[string]*string) *DescribePolicyGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyGroupsResponse) SetStatusCode(v int32) *DescribePolicyGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyGroupsResponse) SetBody(v *DescribePolicyGroupsResponseBody) *DescribePolicyGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyGroupsResponse) Validate() error {
	return dara.Validate(s)
}
