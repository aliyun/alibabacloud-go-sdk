// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberAttributeResponseBody) *DescribePhoneNumberAttributeResponse
	GetBody() *DescribePhoneNumberAttributeResponseBody
}

type DescribePhoneNumberAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberAttributeResponse) GetBody() *DescribePhoneNumberAttributeResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberAttributeResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAttributeResponse) SetStatusCode(v int32) *DescribePhoneNumberAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAttributeResponse) SetBody(v *DescribePhoneNumberAttributeResponseBody) *DescribePhoneNumberAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
