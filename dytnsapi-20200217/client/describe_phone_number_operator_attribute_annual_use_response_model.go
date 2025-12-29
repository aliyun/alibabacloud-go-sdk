// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeAnnualUseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeAnnualUseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeAnnualUseResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) *DescribePhoneNumberOperatorAttributeAnnualUseResponse
	GetBody() *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody
}

type DescribePhoneNumberOperatorAttributeAnnualUseResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualUseResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) GetBody() *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeAnnualUseResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeAnnualUseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) SetBody(v *DescribePhoneNumberOperatorAttributeAnnualUseResponseBody) *DescribePhoneNumberOperatorAttributeAnnualUseResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualUseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
