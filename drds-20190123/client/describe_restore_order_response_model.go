// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreOrderResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreOrderResponseBody) *DescribeRestoreOrderResponse
	GetBody() *DescribeRestoreOrderResponseBody
}

type DescribeRestoreOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreOrderResponse) GetBody() *DescribeRestoreOrderResponseBody {
	return s.Body
}

func (s *DescribeRestoreOrderResponse) SetHeaders(v map[string]*string) *DescribeRestoreOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreOrderResponse) SetStatusCode(v int32) *DescribeRestoreOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreOrderResponse) SetBody(v *DescribeRestoreOrderResponseBody) *DescribeRestoreOrderResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreOrderResponse) Validate() error {
	return dara.Validate(s)
}
