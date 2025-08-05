// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultIPSConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefaultIPSConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefaultIPSConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefaultIPSConfigResponseBody) *DescribeDefaultIPSConfigResponse
	GetBody() *DescribeDefaultIPSConfigResponseBody
}

type DescribeDefaultIPSConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefaultIPSConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefaultIPSConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefaultIPSConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefaultIPSConfigResponse) GetBody() *DescribeDefaultIPSConfigResponseBody {
	return s.Body
}

func (s *DescribeDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *DescribeDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefaultIPSConfigResponse) SetStatusCode(v int32) *DescribeDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponse) SetBody(v *DescribeDefaultIPSConfigResponseBody) *DescribeDefaultIPSConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDefaultIPSConfigResponse) Validate() error {
	return dara.Validate(s)
}
