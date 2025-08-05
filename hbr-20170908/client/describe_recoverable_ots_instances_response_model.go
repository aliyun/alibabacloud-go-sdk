// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecoverableOtsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecoverableOtsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecoverableOtsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecoverableOtsInstancesResponseBody) *DescribeRecoverableOtsInstancesResponse
	GetBody() *DescribeRecoverableOtsInstancesResponseBody
}

type DescribeRecoverableOtsInstancesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecoverableOtsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecoverableOtsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecoverableOtsInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableOtsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecoverableOtsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecoverableOtsInstancesResponse) GetBody() *DescribeRecoverableOtsInstancesResponseBody {
	return s.Body
}

func (s *DescribeRecoverableOtsInstancesResponse) SetHeaders(v map[string]*string) *DescribeRecoverableOtsInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponse) SetStatusCode(v int32) *DescribeRecoverableOtsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponse) SetBody(v *DescribeRecoverableOtsInstancesResponseBody) *DescribeRecoverableOtsInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeRecoverableOtsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
