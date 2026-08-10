// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScansResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScansResponseBody) *DescribeScansResponse
	GetBody() *DescribeScansResponseBody
}

type DescribeScansResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScansResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScansResponse) GoString() string {
	return s.String()
}

func (s *DescribeScansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScansResponse) GetBody() *DescribeScansResponseBody {
	return s.Body
}

func (s *DescribeScansResponse) SetHeaders(v map[string]*string) *DescribeScansResponse {
	s.Headers = v
	return s
}

func (s *DescribeScansResponse) SetStatusCode(v int32) *DescribeScansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScansResponse) SetBody(v *DescribeScansResponseBody) *DescribeScansResponse {
	s.Body = v
	return s
}

func (s *DescribeScansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
