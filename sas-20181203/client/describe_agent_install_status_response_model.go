// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentInstallStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgentInstallStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgentInstallStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgentInstallStatusResponseBody) *DescribeAgentInstallStatusResponse
	GetBody() *DescribeAgentInstallStatusResponseBody
}

type DescribeAgentInstallStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgentInstallStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgentInstallStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentInstallStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgentInstallStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgentInstallStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgentInstallStatusResponse) GetBody() *DescribeAgentInstallStatusResponseBody {
	return s.Body
}

func (s *DescribeAgentInstallStatusResponse) SetHeaders(v map[string]*string) *DescribeAgentInstallStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgentInstallStatusResponse) SetStatusCode(v int32) *DescribeAgentInstallStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgentInstallStatusResponse) SetBody(v *DescribeAgentInstallStatusResponseBody) *DescribeAgentInstallStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAgentInstallStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
