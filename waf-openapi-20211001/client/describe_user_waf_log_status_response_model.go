// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserWafLogStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserWafLogStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserWafLogStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserWafLogStatusResponseBody) *DescribeUserWafLogStatusResponse
	GetBody() *DescribeUserWafLogStatusResponseBody
}

type DescribeUserWafLogStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserWafLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserWafLogStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserWafLogStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserWafLogStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserWafLogStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserWafLogStatusResponse) GetBody() *DescribeUserWafLogStatusResponseBody {
	return s.Body
}

func (s *DescribeUserWafLogStatusResponse) SetHeaders(v map[string]*string) *DescribeUserWafLogStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserWafLogStatusResponse) SetStatusCode(v int32) *DescribeUserWafLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserWafLogStatusResponse) SetBody(v *DescribeUserWafLogStatusResponseBody) *DescribeUserWafLogStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserWafLogStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
