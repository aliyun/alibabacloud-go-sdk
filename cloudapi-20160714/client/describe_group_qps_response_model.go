// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupQpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupQpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupQpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupQpsResponseBody) *DescribeGroupQpsResponse
	GetBody() *DescribeGroupQpsResponseBody
}

type DescribeGroupQpsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupQpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupQpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupQpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupQpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupQpsResponse) GetBody() *DescribeGroupQpsResponseBody {
	return s.Body
}

func (s *DescribeGroupQpsResponse) SetHeaders(v map[string]*string) *DescribeGroupQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupQpsResponse) SetStatusCode(v int32) *DescribeGroupQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupQpsResponse) SetBody(v *DescribeGroupQpsResponseBody) *DescribeGroupQpsResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupQpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
