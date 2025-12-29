// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberOperatorAttributeResponseBody) *DescribePhoneNumberOperatorAttributeResponse
	GetBody() *DescribePhoneNumberOperatorAttributeResponseBody
}

type DescribePhoneNumberOperatorAttributeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberOperatorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberOperatorAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberOperatorAttributeResponse) GetBody() *DescribePhoneNumberOperatorAttributeResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponse) SetBody(v *DescribePhoneNumberOperatorAttributeResponseBody) *DescribePhoneNumberOperatorAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
