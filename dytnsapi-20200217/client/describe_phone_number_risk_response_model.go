// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberRiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberRiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberRiskResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberRiskResponseBody) *DescribePhoneNumberRiskResponse
	GetBody() *DescribePhoneNumberRiskResponseBody
}

type DescribePhoneNumberRiskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberRiskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberRiskResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberRiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberRiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberRiskResponse) GetBody() *DescribePhoneNumberRiskResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberRiskResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberRiskResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberRiskResponse) SetStatusCode(v int32) *DescribePhoneNumberRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberRiskResponse) SetBody(v *DescribePhoneNumberRiskResponseBody) *DescribePhoneNumberRiskResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberRiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
