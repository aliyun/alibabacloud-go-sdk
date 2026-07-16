// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConditionIPBInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConditionIPBInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConditionIPBInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConditionIPBInfoResponseBody) *DescribeConditionIPBInfoResponse
	GetBody() *DescribeConditionIPBInfoResponseBody
}

type DescribeConditionIPBInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConditionIPBInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConditionIPBInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConditionIPBInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeConditionIPBInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConditionIPBInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConditionIPBInfoResponse) GetBody() *DescribeConditionIPBInfoResponseBody {
	return s.Body
}

func (s *DescribeConditionIPBInfoResponse) SetHeaders(v map[string]*string) *DescribeConditionIPBInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeConditionIPBInfoResponse) SetStatusCode(v int32) *DescribeConditionIPBInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConditionIPBInfoResponse) SetBody(v *DescribeConditionIPBInfoResponseBody) *DescribeConditionIPBInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeConditionIPBInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
