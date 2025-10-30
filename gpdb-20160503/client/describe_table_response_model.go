// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTableResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTableResponseBody) *DescribeTableResponse
	GetBody() *DescribeTableResponseBody
}

type DescribeTableResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableResponse) GoString() string {
	return s.String()
}

func (s *DescribeTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTableResponse) GetBody() *DescribeTableResponseBody {
	return s.Body
}

func (s *DescribeTableResponse) SetHeaders(v map[string]*string) *DescribeTableResponse {
	s.Headers = v
	return s
}

func (s *DescribeTableResponse) SetStatusCode(v int32) *DescribeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTableResponse) SetBody(v *DescribeTableResponseBody) *DescribeTableResponse {
	s.Body = v
	return s
}

func (s *DescribeTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
