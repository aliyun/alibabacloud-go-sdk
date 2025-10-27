// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHoneyPotAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHoneyPotAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHoneyPotAuthResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHoneyPotAuthResponseBody) *DescribeHoneyPotAuthResponse
	GetBody() *DescribeHoneyPotAuthResponseBody
}

type DescribeHoneyPotAuthResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHoneyPotAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHoneyPotAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHoneyPotAuthResponse) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHoneyPotAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHoneyPotAuthResponse) GetBody() *DescribeHoneyPotAuthResponseBody {
	return s.Body
}

func (s *DescribeHoneyPotAuthResponse) SetHeaders(v map[string]*string) *DescribeHoneyPotAuthResponse {
	s.Headers = v
	return s
}

func (s *DescribeHoneyPotAuthResponse) SetStatusCode(v int32) *DescribeHoneyPotAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHoneyPotAuthResponse) SetBody(v *DescribeHoneyPotAuthResponseBody) *DescribeHoneyPotAuthResponse {
	s.Body = v
	return s
}

func (s *DescribeHoneyPotAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
