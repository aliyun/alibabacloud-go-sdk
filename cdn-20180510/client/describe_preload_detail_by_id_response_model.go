// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreloadDetailByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePreloadDetailByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePreloadDetailByIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribePreloadDetailByIdResponseBody) *DescribePreloadDetailByIdResponse
	GetBody() *DescribePreloadDetailByIdResponseBody
}

type DescribePreloadDetailByIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePreloadDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePreloadDetailByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePreloadDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribePreloadDetailByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePreloadDetailByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePreloadDetailByIdResponse) GetBody() *DescribePreloadDetailByIdResponseBody {
	return s.Body
}

func (s *DescribePreloadDetailByIdResponse) SetHeaders(v map[string]*string) *DescribePreloadDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribePreloadDetailByIdResponse) SetStatusCode(v int32) *DescribePreloadDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreloadDetailByIdResponse) SetBody(v *DescribePreloadDetailByIdResponseBody) *DescribePreloadDetailByIdResponse {
	s.Body = v
	return s
}

func (s *DescribePreloadDetailByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
