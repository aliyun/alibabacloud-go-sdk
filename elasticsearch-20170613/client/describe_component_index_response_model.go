// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentIndexResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentIndexResponseBody) *DescribeComponentIndexResponse
	GetBody() *DescribeComponentIndexResponseBody
}

type DescribeComponentIndexResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentIndexResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentIndexResponse) GetBody() *DescribeComponentIndexResponseBody {
	return s.Body
}

func (s *DescribeComponentIndexResponse) SetHeaders(v map[string]*string) *DescribeComponentIndexResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentIndexResponse) SetStatusCode(v int32) *DescribeComponentIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentIndexResponse) SetBody(v *DescribeComponentIndexResponseBody) *DescribeComponentIndexResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
