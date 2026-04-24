// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelOperatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelOperatorResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelOperatorResponseBody) *DescribeModelOperatorResponse
	GetBody() *DescribeModelOperatorResponseBody
}

type DescribeModelOperatorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelOperatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelOperatorResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelOperatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelOperatorResponse) GetBody() *DescribeModelOperatorResponseBody {
	return s.Body
}

func (s *DescribeModelOperatorResponse) SetHeaders(v map[string]*string) *DescribeModelOperatorResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelOperatorResponse) SetStatusCode(v int32) *DescribeModelOperatorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelOperatorResponse) SetBody(v *DescribeModelOperatorResponseBody) *DescribeModelOperatorResponse {
	s.Body = v
	return s
}

func (s *DescribeModelOperatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
