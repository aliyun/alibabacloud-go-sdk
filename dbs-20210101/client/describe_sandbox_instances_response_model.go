// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSandboxInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSandboxInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSandboxInstancesResponseBody) *DescribeSandboxInstancesResponse
	GetBody() *DescribeSandboxInstancesResponseBody
}

type DescribeSandboxInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSandboxInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSandboxInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSandboxInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSandboxInstancesResponse) GetBody() *DescribeSandboxInstancesResponseBody {
	return s.Body
}

func (s *DescribeSandboxInstancesResponse) SetHeaders(v map[string]*string) *DescribeSandboxInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxInstancesResponse) SetStatusCode(v int32) *DescribeSandboxInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxInstancesResponse) SetBody(v *DescribeSandboxInstancesResponseBody) *DescribeSandboxInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeSandboxInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
