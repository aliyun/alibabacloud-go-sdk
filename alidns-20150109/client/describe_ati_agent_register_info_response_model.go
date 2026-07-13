// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiAgentRegisterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAtiAgentRegisterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAtiAgentRegisterInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAtiAgentRegisterInfoResponseBody) *DescribeAtiAgentRegisterInfoResponse
	GetBody() *DescribeAtiAgentRegisterInfoResponseBody
}

type DescribeAtiAgentRegisterInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAtiAgentRegisterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAtiAgentRegisterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAtiAgentRegisterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAtiAgentRegisterInfoResponse) GetBody() *DescribeAtiAgentRegisterInfoResponseBody {
	return s.Body
}

func (s *DescribeAtiAgentRegisterInfoResponse) SetHeaders(v map[string]*string) *DescribeAtiAgentRegisterInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoResponse) SetStatusCode(v int32) *DescribeAtiAgentRegisterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoResponse) SetBody(v *DescribeAtiAgentRegisterInfoResponseBody) *DescribeAtiAgentRegisterInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeAtiAgentRegisterInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
