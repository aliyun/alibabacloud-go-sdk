// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBaselineAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserBaselineAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserBaselineAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserBaselineAuthorizationResponseBody) *DescribeUserBaselineAuthorizationResponse
	GetBody() *DescribeUserBaselineAuthorizationResponseBody
}

type DescribeUserBaselineAuthorizationResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBaselineAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBaselineAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBaselineAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBaselineAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserBaselineAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserBaselineAuthorizationResponse) GetBody() *DescribeUserBaselineAuthorizationResponseBody {
	return s.Body
}

func (s *DescribeUserBaselineAuthorizationResponse) SetHeaders(v map[string]*string) *DescribeUserBaselineAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBaselineAuthorizationResponse) SetStatusCode(v int32) *DescribeUserBaselineAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationResponse) SetBody(v *DescribeUserBaselineAuthorizationResponseBody) *DescribeUserBaselineAuthorizationResponse {
	s.Body = v
	return s
}

func (s *DescribeUserBaselineAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
