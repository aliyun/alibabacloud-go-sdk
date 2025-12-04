// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConfigHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConfigHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConfigHistoryResponseBody) *DescribeConfigHistoryResponse
	GetBody() *DescribeConfigHistoryResponseBody
}

type DescribeConfigHistoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfigHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConfigHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConfigHistoryResponse) GetBody() *DescribeConfigHistoryResponseBody {
	return s.Body
}

func (s *DescribeConfigHistoryResponse) SetHeaders(v map[string]*string) *DescribeConfigHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigHistoryResponse) SetStatusCode(v int32) *DescribeConfigHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigHistoryResponse) SetBody(v *DescribeConfigHistoryResponseBody) *DescribeConfigHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeConfigHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
