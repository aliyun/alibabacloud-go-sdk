// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupConditionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupConditionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupConditionListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupConditionListResponseBody) *DescribeGroupConditionListResponse
	GetBody() *DescribeGroupConditionListResponseBody
}

type DescribeGroupConditionListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupConditionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupConditionListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupConditionListResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupConditionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupConditionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupConditionListResponse) GetBody() *DescribeGroupConditionListResponseBody {
	return s.Body
}

func (s *DescribeGroupConditionListResponse) SetHeaders(v map[string]*string) *DescribeGroupConditionListResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupConditionListResponse) SetStatusCode(v int32) *DescribeGroupConditionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupConditionListResponse) SetBody(v *DescribeGroupConditionListResponseBody) *DescribeGroupConditionListResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupConditionListResponse) Validate() error {
	return dara.Validate(s)
}
