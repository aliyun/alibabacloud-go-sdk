// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebPathResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebPathResponseBody) *DescribeWebPathResponse
	GetBody() *DescribeWebPathResponseBody
}

type DescribeWebPathResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebPathResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPathResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebPathResponse) GetBody() *DescribeWebPathResponseBody {
	return s.Body
}

func (s *DescribeWebPathResponse) SetHeaders(v map[string]*string) *DescribeWebPathResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebPathResponse) SetStatusCode(v int32) *DescribeWebPathResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebPathResponse) SetBody(v *DescribeWebPathResponseBody) *DescribeWebPathResponse {
	s.Body = v
	return s
}

func (s *DescribeWebPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
