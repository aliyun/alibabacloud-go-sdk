// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationCodeDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActivationCodeDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActivationCodeDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActivationCodeDetailsResponseBody) *DescribeActivationCodeDetailsResponse
	GetBody() *DescribeActivationCodeDetailsResponseBody
}

type DescribeActivationCodeDetailsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActivationCodeDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActivationCodeDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationCodeDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeActivationCodeDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActivationCodeDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActivationCodeDetailsResponse) GetBody() *DescribeActivationCodeDetailsResponseBody {
	return s.Body
}

func (s *DescribeActivationCodeDetailsResponse) SetHeaders(v map[string]*string) *DescribeActivationCodeDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeActivationCodeDetailsResponse) SetStatusCode(v int32) *DescribeActivationCodeDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActivationCodeDetailsResponse) SetBody(v *DescribeActivationCodeDetailsResponseBody) *DescribeActivationCodeDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeActivationCodeDetailsResponse) Validate() error {
	return dara.Validate(s)
}
