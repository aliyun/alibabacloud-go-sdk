// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInterAuthStatisticsHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInterAuthStatisticsHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInterAuthStatisticsHistoryResponseBody) *DescribeInterAuthStatisticsHistoryResponse
	GetBody() *DescribeInterAuthStatisticsHistoryResponseBody
}

type DescribeInterAuthStatisticsHistoryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInterAuthStatisticsHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInterAuthStatisticsHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInterAuthStatisticsHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInterAuthStatisticsHistoryResponse) GetBody() *DescribeInterAuthStatisticsHistoryResponseBody {
	return s.Body
}

func (s *DescribeInterAuthStatisticsHistoryResponse) SetHeaders(v map[string]*string) *DescribeInterAuthStatisticsHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponse) SetStatusCode(v int32) *DescribeInterAuthStatisticsHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponse) SetBody(v *DescribeInterAuthStatisticsHistoryResponseBody) *DescribeInterAuthStatisticsHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeInterAuthStatisticsHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
