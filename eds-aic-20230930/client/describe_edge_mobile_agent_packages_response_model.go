// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMobileAgentPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeMobileAgentPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeMobileAgentPackagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeMobileAgentPackagesResponseBody) *DescribeEdgeMobileAgentPackagesResponse
	GetBody() *DescribeEdgeMobileAgentPackagesResponseBody
}

type DescribeEdgeMobileAgentPackagesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeMobileAgentPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeMobileAgentPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMobileAgentPackagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMobileAgentPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeMobileAgentPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeMobileAgentPackagesResponse) GetBody() *DescribeEdgeMobileAgentPackagesResponseBody {
	return s.Body
}

func (s *DescribeEdgeMobileAgentPackagesResponse) SetHeaders(v map[string]*string) *DescribeEdgeMobileAgentPackagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponse) SetStatusCode(v int32) *DescribeEdgeMobileAgentPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponse) SetBody(v *DescribeEdgeMobileAgentPackagesResponseBody) *DescribeEdgeMobileAgentPackagesResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeMobileAgentPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
