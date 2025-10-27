// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNeedAsyncQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNeedAsyncQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNeedAsyncQueryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNeedAsyncQueryResponseBody) *DescribeNeedAsyncQueryResponse
	GetBody() *DescribeNeedAsyncQueryResponseBody
}

type DescribeNeedAsyncQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNeedAsyncQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNeedAsyncQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNeedAsyncQueryResponse) GoString() string {
	return s.String()
}

func (s *DescribeNeedAsyncQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNeedAsyncQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNeedAsyncQueryResponse) GetBody() *DescribeNeedAsyncQueryResponseBody {
	return s.Body
}

func (s *DescribeNeedAsyncQueryResponse) SetHeaders(v map[string]*string) *DescribeNeedAsyncQueryResponse {
	s.Headers = v
	return s
}

func (s *DescribeNeedAsyncQueryResponse) SetStatusCode(v int32) *DescribeNeedAsyncQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNeedAsyncQueryResponse) SetBody(v *DescribeNeedAsyncQueryResponseBody) *DescribeNeedAsyncQueryResponse {
	s.Body = v
	return s
}

func (s *DescribeNeedAsyncQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
