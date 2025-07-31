// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterModificationHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParameterModificationHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParameterModificationHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParameterModificationHistoryResponseBody) *DescribeParameterModificationHistoryResponse
	GetBody() *DescribeParameterModificationHistoryResponseBody
}

type DescribeParameterModificationHistoryResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParameterModificationHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParameterModificationHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterModificationHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParameterModificationHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParameterModificationHistoryResponse) GetBody() *DescribeParameterModificationHistoryResponseBody {
	return s.Body
}

func (s *DescribeParameterModificationHistoryResponse) SetHeaders(v map[string]*string) *DescribeParameterModificationHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterModificationHistoryResponse) SetStatusCode(v int32) *DescribeParameterModificationHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParameterModificationHistoryResponse) SetBody(v *DescribeParameterModificationHistoryResponseBody) *DescribeParameterModificationHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeParameterModificationHistoryResponse) Validate() error {
	return dara.Validate(s)
}
