// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyInstancesStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolicyInstancesStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolicyInstancesStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolicyInstancesStatusResponseBody) *DescribePolicyInstancesStatusResponse
	GetBody() *DescribePolicyInstancesStatusResponseBody
}

type DescribePolicyInstancesStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyInstancesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyInstancesStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyInstancesStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyInstancesStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolicyInstancesStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolicyInstancesStatusResponse) GetBody() *DescribePolicyInstancesStatusResponseBody {
	return s.Body
}

func (s *DescribePolicyInstancesStatusResponse) SetHeaders(v map[string]*string) *DescribePolicyInstancesStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyInstancesStatusResponse) SetStatusCode(v int32) *DescribePolicyInstancesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyInstancesStatusResponse) SetBody(v *DescribePolicyInstancesStatusResponseBody) *DescribePolicyInstancesStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePolicyInstancesStatusResponse) Validate() error {
	return dara.Validate(s)
}
