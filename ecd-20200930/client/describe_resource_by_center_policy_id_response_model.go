// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceByCenterPolicyIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceByCenterPolicyIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceByCenterPolicyIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceByCenterPolicyIdResponseBody) *DescribeResourceByCenterPolicyIdResponse
	GetBody() *DescribeResourceByCenterPolicyIdResponseBody
}

type DescribeResourceByCenterPolicyIdResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceByCenterPolicyIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceByCenterPolicyIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByCenterPolicyIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceByCenterPolicyIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceByCenterPolicyIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceByCenterPolicyIdResponse) GetBody() *DescribeResourceByCenterPolicyIdResponseBody {
	return s.Body
}

func (s *DescribeResourceByCenterPolicyIdResponse) SetHeaders(v map[string]*string) *DescribeResourceByCenterPolicyIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponse) SetStatusCode(v int32) *DescribeResourceByCenterPolicyIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponse) SetBody(v *DescribeResourceByCenterPolicyIdResponseBody) *DescribeResourceByCenterPolicyIdResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponse) Validate() error {
	return dara.Validate(s)
}
