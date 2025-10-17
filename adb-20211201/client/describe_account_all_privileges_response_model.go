// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAllPrivilegesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountAllPrivilegesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountAllPrivilegesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountAllPrivilegesResponseBody) *DescribeAccountAllPrivilegesResponse
	GetBody() *DescribeAccountAllPrivilegesResponseBody
}

type DescribeAccountAllPrivilegesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountAllPrivilegesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountAllPrivilegesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountAllPrivilegesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountAllPrivilegesResponse) GetBody() *DescribeAccountAllPrivilegesResponseBody {
	return s.Body
}

func (s *DescribeAccountAllPrivilegesResponse) SetHeaders(v map[string]*string) *DescribeAccountAllPrivilegesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponse) SetStatusCode(v int32) *DescribeAccountAllPrivilegesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponse) SetBody(v *DescribeAccountAllPrivilegesResponseBody) *DescribeAccountAllPrivilegesResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
