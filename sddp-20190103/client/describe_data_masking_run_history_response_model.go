// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataMaskingRunHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataMaskingRunHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataMaskingRunHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataMaskingRunHistoryResponseBody) *DescribeDataMaskingRunHistoryResponse
	GetBody() *DescribeDataMaskingRunHistoryResponseBody
}

type DescribeDataMaskingRunHistoryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataMaskingRunHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataMaskingRunHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataMaskingRunHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataMaskingRunHistoryResponse) GetBody() *DescribeDataMaskingRunHistoryResponseBody {
	return s.Body
}

func (s *DescribeDataMaskingRunHistoryResponse) SetHeaders(v map[string]*string) *DescribeDataMaskingRunHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponse) SetStatusCode(v int32) *DescribeDataMaskingRunHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponse) SetBody(v *DescribeDataMaskingRunHistoryResponseBody) *DescribeDataMaskingRunHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
