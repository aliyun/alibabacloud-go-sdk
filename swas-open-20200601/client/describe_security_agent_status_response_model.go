// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityAgentStatusResponseBody) *DescribeSecurityAgentStatusResponse
	GetBody() *DescribeSecurityAgentStatusResponseBody
}

type DescribeSecurityAgentStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityAgentStatusResponse) GetBody() *DescribeSecurityAgentStatusResponseBody {
	return s.Body
}

func (s *DescribeSecurityAgentStatusResponse) SetHeaders(v map[string]*string) *DescribeSecurityAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityAgentStatusResponse) SetStatusCode(v int32) *DescribeSecurityAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityAgentStatusResponse) SetBody(v *DescribeSecurityAgentStatusResponseBody) *DescribeSecurityAgentStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
