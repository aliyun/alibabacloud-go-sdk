// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortCcAttackTopIPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortCcAttackTopIPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortCcAttackTopIPResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortCcAttackTopIPResponseBody) *DescribePortCcAttackTopIPResponse
	GetBody() *DescribePortCcAttackTopIPResponseBody
}

type DescribePortCcAttackTopIPResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortCcAttackTopIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortCcAttackTopIPResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortCcAttackTopIPResponse) GoString() string {
	return s.String()
}

func (s *DescribePortCcAttackTopIPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortCcAttackTopIPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortCcAttackTopIPResponse) GetBody() *DescribePortCcAttackTopIPResponseBody {
	return s.Body
}

func (s *DescribePortCcAttackTopIPResponse) SetHeaders(v map[string]*string) *DescribePortCcAttackTopIPResponse {
	s.Headers = v
	return s
}

func (s *DescribePortCcAttackTopIPResponse) SetStatusCode(v int32) *DescribePortCcAttackTopIPResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortCcAttackTopIPResponse) SetBody(v *DescribePortCcAttackTopIPResponseBody) *DescribePortCcAttackTopIPResponse {
	s.Body = v
	return s
}

func (s *DescribePortCcAttackTopIPResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
