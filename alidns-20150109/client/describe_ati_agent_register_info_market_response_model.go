// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiAgentRegisterInfoMarketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAtiAgentRegisterInfoMarketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAtiAgentRegisterInfoMarketResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAtiAgentRegisterInfoMarketResponseBody) *DescribeAtiAgentRegisterInfoMarketResponse
	GetBody() *DescribeAtiAgentRegisterInfoMarketResponseBody
}

type DescribeAtiAgentRegisterInfoMarketResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAtiAgentRegisterInfoMarketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAtiAgentRegisterInfoMarketResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoMarketResponse) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoMarketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAtiAgentRegisterInfoMarketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAtiAgentRegisterInfoMarketResponse) GetBody() *DescribeAtiAgentRegisterInfoMarketResponseBody {
	return s.Body
}

func (s *DescribeAtiAgentRegisterInfoMarketResponse) SetHeaders(v map[string]*string) *DescribeAtiAgentRegisterInfoMarketResponse {
	s.Headers = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponse) SetStatusCode(v int32) *DescribeAtiAgentRegisterInfoMarketResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponse) SetBody(v *DescribeAtiAgentRegisterInfoMarketResponseBody) *DescribeAtiAgentRegisterInfoMarketResponse {
	s.Body = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
