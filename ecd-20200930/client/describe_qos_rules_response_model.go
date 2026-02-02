// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQosRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQosRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQosRulesResponseBody) *DescribeQosRulesResponse
	GetBody() *DescribeQosRulesResponseBody
}

type DescribeQosRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQosRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQosRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeQosRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQosRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQosRulesResponse) GetBody() *DescribeQosRulesResponseBody {
	return s.Body
}

func (s *DescribeQosRulesResponse) SetHeaders(v map[string]*string) *DescribeQosRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeQosRulesResponse) SetStatusCode(v int32) *DescribeQosRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQosRulesResponse) SetBody(v *DescribeQosRulesResponseBody) *DescribeQosRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeQosRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
