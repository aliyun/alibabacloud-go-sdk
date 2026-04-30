// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJVSInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJVSInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJVSInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJVSInstanceResponseBody) *DescribeJVSInstanceResponse
	GetBody() *DescribeJVSInstanceResponseBody
}

type DescribeJVSInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJVSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJVSInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJVSInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeJVSInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJVSInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJVSInstanceResponse) GetBody() *DescribeJVSInstanceResponseBody {
	return s.Body
}

func (s *DescribeJVSInstanceResponse) SetHeaders(v map[string]*string) *DescribeJVSInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeJVSInstanceResponse) SetStatusCode(v int32) *DescribeJVSInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJVSInstanceResponse) SetBody(v *DescribeJVSInstanceResponseBody) *DescribeJVSInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeJVSInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
