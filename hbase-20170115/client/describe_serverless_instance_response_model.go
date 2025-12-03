// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerlessInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServerlessInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServerlessInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServerlessInstanceResponseBody) *DescribeServerlessInstanceResponse
	GetBody() *DescribeServerlessInstanceResponseBody
}

type DescribeServerlessInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServerlessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServerlessInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerlessInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeServerlessInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServerlessInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServerlessInstanceResponse) GetBody() *DescribeServerlessInstanceResponseBody {
	return s.Body
}

func (s *DescribeServerlessInstanceResponse) SetHeaders(v map[string]*string) *DescribeServerlessInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeServerlessInstanceResponse) SetStatusCode(v int32) *DescribeServerlessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServerlessInstanceResponse) SetBody(v *DescribeServerlessInstanceResponseBody) *DescribeServerlessInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeServerlessInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
