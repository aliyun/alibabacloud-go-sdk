// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsVpcsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdsVpcsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdsVpcsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdsVpcsResponseBody) *DescribeRdsVpcsResponse
	GetBody() *DescribeRdsVpcsResponseBody
}

type DescribeRdsVpcsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdsVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdsVpcsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsVpcsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdsVpcsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdsVpcsResponse) GetBody() *DescribeRdsVpcsResponseBody {
	return s.Body
}

func (s *DescribeRdsVpcsResponse) SetHeaders(v map[string]*string) *DescribeRdsVpcsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdsVpcsResponse) SetStatusCode(v int32) *DescribeRdsVpcsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdsVpcsResponse) SetBody(v *DescribeRdsVpcsResponseBody) *DescribeRdsVpcsResponse {
	s.Body = v
	return s
}

func (s *DescribeRdsVpcsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
