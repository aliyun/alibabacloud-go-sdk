// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransferHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransferHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransferHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransferHistoryResponseBody) *DescribeTransferHistoryResponse
	GetBody() *DescribeTransferHistoryResponseBody
}

type DescribeTransferHistoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransferHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransferHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransferHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransferHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransferHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransferHistoryResponse) GetBody() *DescribeTransferHistoryResponseBody {
	return s.Body
}

func (s *DescribeTransferHistoryResponse) SetHeaders(v map[string]*string) *DescribeTransferHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransferHistoryResponse) SetStatusCode(v int32) *DescribeTransferHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransferHistoryResponse) SetBody(v *DescribeTransferHistoryResponseBody) *DescribeTransferHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeTransferHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
