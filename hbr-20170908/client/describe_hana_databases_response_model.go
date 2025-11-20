// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHanaDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHanaDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHanaDatabasesResponseBody) *DescribeHanaDatabasesResponse
	GetBody() *DescribeHanaDatabasesResponseBody
}

type DescribeHanaDatabasesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHanaDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHanaDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeHanaDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHanaDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHanaDatabasesResponse) GetBody() *DescribeHanaDatabasesResponseBody {
	return s.Body
}

func (s *DescribeHanaDatabasesResponse) SetHeaders(v map[string]*string) *DescribeHanaDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeHanaDatabasesResponse) SetStatusCode(v int32) *DescribeHanaDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHanaDatabasesResponse) SetBody(v *DescribeHanaDatabasesResponseBody) *DescribeHanaDatabasesResponse {
	s.Body = v
	return s
}

func (s *DescribeHanaDatabasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
