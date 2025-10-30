// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHitRuleTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHitRuleTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHitRuleTrendResponseBody) *DescribeHitRuleTrendResponse
	GetBody() *DescribeHitRuleTrendResponseBody
}

type DescribeHitRuleTrendResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHitRuleTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHitRuleTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHitRuleTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHitRuleTrendResponse) GetBody() *DescribeHitRuleTrendResponseBody {
	return s.Body
}

func (s *DescribeHitRuleTrendResponse) SetHeaders(v map[string]*string) *DescribeHitRuleTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeHitRuleTrendResponse) SetStatusCode(v int32) *DescribeHitRuleTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHitRuleTrendResponse) SetBody(v *DescribeHitRuleTrendResponseBody) *DescribeHitRuleTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeHitRuleTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
