// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeAnnualResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeAnnualResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeAnnualResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberOperatorAttributeAnnualResponseBody) *DescribePhoneNumberOperatorAttributeAnnualResponse
	GetBody() *DescribePhoneNumberOperatorAttributeAnnualResponseBody
}

type DescribePhoneNumberOperatorAttributeAnnualResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberOperatorAttributeAnnualResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeAnnualResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) GetBody() *DescribePhoneNumberOperatorAttributeAnnualResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberOperatorAttributeAnnualResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) SetStatusCode(v int32) *DescribePhoneNumberOperatorAttributeAnnualResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) SetBody(v *DescribePhoneNumberOperatorAttributeAnnualResponseBody) *DescribePhoneNumberOperatorAttributeAnnualResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeAnnualResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
