// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataAgentSessionResponseBody) *DescribeDataAgentSessionResponse
	GetBody() *DescribeDataAgentSessionResponseBody
}

type DescribeDataAgentSessionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataAgentSessionResponse) GetBody() *DescribeDataAgentSessionResponseBody {
	return s.Body
}

func (s *DescribeDataAgentSessionResponse) SetHeaders(v map[string]*string) *DescribeDataAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataAgentSessionResponse) SetStatusCode(v int32) *DescribeDataAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataAgentSessionResponse) SetBody(v *DescribeDataAgentSessionResponseBody) *DescribeDataAgentSessionResponse {
	s.Body = v
	return s
}

func (s *DescribeDataAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
