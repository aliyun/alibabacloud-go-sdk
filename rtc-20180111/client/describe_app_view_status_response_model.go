// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppViewStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppViewStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppViewStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppViewStatusResponseBody) *DescribeAppViewStatusResponse
	GetBody() *DescribeAppViewStatusResponseBody
}

type DescribeAppViewStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppViewStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppViewStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppViewStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppViewStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppViewStatusResponse) GetBody() *DescribeAppViewStatusResponseBody {
	return s.Body
}

func (s *DescribeAppViewStatusResponse) SetHeaders(v map[string]*string) *DescribeAppViewStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppViewStatusResponse) SetStatusCode(v int32) *DescribeAppViewStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppViewStatusResponse) SetBody(v *DescribeAppViewStatusResponseBody) *DescribeAppViewStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAppViewStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
