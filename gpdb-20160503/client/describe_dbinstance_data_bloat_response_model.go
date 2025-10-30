// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataBloatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceDataBloatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceDataBloatResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceDataBloatResponseBody) *DescribeDBInstanceDataBloatResponse
	GetBody() *DescribeDBInstanceDataBloatResponseBody
}

type DescribeDBInstanceDataBloatResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceDataBloatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceDataBloatResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceDataBloatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceDataBloatResponse) GetBody() *DescribeDBInstanceDataBloatResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceDataBloatResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceDataBloatResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceDataBloatResponse) SetStatusCode(v int32) *DescribeDBInstanceDataBloatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponse) SetBody(v *DescribeDBInstanceDataBloatResponseBody) *DescribeDBInstanceDataBloatResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceDataBloatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
