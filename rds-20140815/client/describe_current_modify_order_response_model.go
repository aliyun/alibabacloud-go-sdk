// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCurrentModifyOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCurrentModifyOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCurrentModifyOrderResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCurrentModifyOrderResponseBody) *DescribeCurrentModifyOrderResponse
	GetBody() *DescribeCurrentModifyOrderResponseBody
}

type DescribeCurrentModifyOrderResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCurrentModifyOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCurrentModifyOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCurrentModifyOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeCurrentModifyOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCurrentModifyOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCurrentModifyOrderResponse) GetBody() *DescribeCurrentModifyOrderResponseBody {
	return s.Body
}

func (s *DescribeCurrentModifyOrderResponse) SetHeaders(v map[string]*string) *DescribeCurrentModifyOrderResponse {
	s.Headers = v
	return s
}

func (s *DescribeCurrentModifyOrderResponse) SetStatusCode(v int32) *DescribeCurrentModifyOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCurrentModifyOrderResponse) SetBody(v *DescribeCurrentModifyOrderResponseBody) *DescribeCurrentModifyOrderResponse {
	s.Body = v
	return s
}

func (s *DescribeCurrentModifyOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
