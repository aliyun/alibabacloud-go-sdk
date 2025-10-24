// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChargeResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChargeResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChargeResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChargeResultResponseBody) *DescribeChargeResultResponse
	GetBody() *DescribeChargeResultResponseBody
}

type DescribeChargeResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChargeResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChargeResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeChargeResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChargeResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChargeResultResponse) GetBody() *DescribeChargeResultResponseBody {
	return s.Body
}

func (s *DescribeChargeResultResponse) SetHeaders(v map[string]*string) *DescribeChargeResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeChargeResultResponse) SetStatusCode(v int32) *DescribeChargeResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChargeResultResponse) SetBody(v *DescribeChargeResultResponseBody) *DescribeChargeResultResponse {
	s.Body = v
	return s
}

func (s *DescribeChargeResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
