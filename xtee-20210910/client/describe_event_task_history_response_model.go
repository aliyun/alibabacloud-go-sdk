// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTaskHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventTaskHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventTaskHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventTaskHistoryResponseBody) *DescribeEventTaskHistoryResponse
	GetBody() *DescribeEventTaskHistoryResponseBody
}

type DescribeEventTaskHistoryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventTaskHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventTaskHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTaskHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventTaskHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventTaskHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventTaskHistoryResponse) GetBody() *DescribeEventTaskHistoryResponseBody {
	return s.Body
}

func (s *DescribeEventTaskHistoryResponse) SetHeaders(v map[string]*string) *DescribeEventTaskHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventTaskHistoryResponse) SetStatusCode(v int32) *DescribeEventTaskHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventTaskHistoryResponse) SetBody(v *DescribeEventTaskHistoryResponseBody) *DescribeEventTaskHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeEventTaskHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
