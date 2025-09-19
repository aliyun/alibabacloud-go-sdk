// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageEventOperationConditionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageEventOperationConditionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageEventOperationConditionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageEventOperationConditionResponseBody) *DescribeImageEventOperationConditionResponse
	GetBody() *DescribeImageEventOperationConditionResponseBody
}

type DescribeImageEventOperationConditionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageEventOperationConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageEventOperationConditionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageEventOperationConditionResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageEventOperationConditionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageEventOperationConditionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageEventOperationConditionResponse) GetBody() *DescribeImageEventOperationConditionResponseBody {
	return s.Body
}

func (s *DescribeImageEventOperationConditionResponse) SetHeaders(v map[string]*string) *DescribeImageEventOperationConditionResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageEventOperationConditionResponse) SetStatusCode(v int32) *DescribeImageEventOperationConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageEventOperationConditionResponse) SetBody(v *DescribeImageEventOperationConditionResponseBody) *DescribeImageEventOperationConditionResponse {
	s.Body = v
	return s
}

func (s *DescribeImageEventOperationConditionResponse) Validate() error {
	return dara.Validate(s)
}
