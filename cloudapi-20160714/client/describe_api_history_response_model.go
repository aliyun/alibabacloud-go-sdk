// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiHistoryResponseBody) *DescribeApiHistoryResponse
	GetBody() *DescribeApiHistoryResponseBody
}

type DescribeApiHistoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiHistoryResponse) GetBody() *DescribeApiHistoryResponseBody {
	return s.Body
}

func (s *DescribeApiHistoryResponse) SetHeaders(v map[string]*string) *DescribeApiHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiHistoryResponse) SetStatusCode(v int32) *DescribeApiHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiHistoryResponse) SetBody(v *DescribeApiHistoryResponseBody) *DescribeApiHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeApiHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
