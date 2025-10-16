// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserProfilePathRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserProfilePathRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserProfilePathRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserProfilePathRulesResponseBody) *DescribeUserProfilePathRulesResponse
	GetBody() *DescribeUserProfilePathRulesResponseBody
}

type DescribeUserProfilePathRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserProfilePathRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserProfilePathRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserProfilePathRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserProfilePathRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserProfilePathRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserProfilePathRulesResponse) GetBody() *DescribeUserProfilePathRulesResponseBody {
	return s.Body
}

func (s *DescribeUserProfilePathRulesResponse) SetHeaders(v map[string]*string) *DescribeUserProfilePathRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserProfilePathRulesResponse) SetStatusCode(v int32) *DescribeUserProfilePathRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserProfilePathRulesResponse) SetBody(v *DescribeUserProfilePathRulesResponseBody) *DescribeUserProfilePathRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeUserProfilePathRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
