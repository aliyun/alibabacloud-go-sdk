// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatAclPageStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatAclPageStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatAclPageStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatAclPageStatusResponseBody) *DescribeNatAclPageStatusResponse
	GetBody() *DescribeNatAclPageStatusResponseBody
}

type DescribeNatAclPageStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatAclPageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatAclPageStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatAclPageStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatAclPageStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatAclPageStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatAclPageStatusResponse) GetBody() *DescribeNatAclPageStatusResponseBody {
	return s.Body
}

func (s *DescribeNatAclPageStatusResponse) SetHeaders(v map[string]*string) *DescribeNatAclPageStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatAclPageStatusResponse) SetStatusCode(v int32) *DescribeNatAclPageStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatAclPageStatusResponse) SetBody(v *DescribeNatAclPageStatusResponseBody) *DescribeNatAclPageStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeNatAclPageStatusResponse) Validate() error {
	return dara.Validate(s)
}
