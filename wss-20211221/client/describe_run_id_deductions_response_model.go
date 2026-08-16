// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRunIdDeductionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRunIdDeductionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRunIdDeductionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRunIdDeductionsResponseBody) *DescribeRunIdDeductionsResponse
	GetBody() *DescribeRunIdDeductionsResponseBody
}

type DescribeRunIdDeductionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRunIdDeductionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRunIdDeductionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunIdDeductionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRunIdDeductionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRunIdDeductionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRunIdDeductionsResponse) GetBody() *DescribeRunIdDeductionsResponseBody {
	return s.Body
}

func (s *DescribeRunIdDeductionsResponse) SetHeaders(v map[string]*string) *DescribeRunIdDeductionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRunIdDeductionsResponse) SetStatusCode(v int32) *DescribeRunIdDeductionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRunIdDeductionsResponse) SetBody(v *DescribeRunIdDeductionsResponseBody) *DescribeRunIdDeductionsResponse {
	s.Body = v
	return s
}

func (s *DescribeRunIdDeductionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
