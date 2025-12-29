// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisPaiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisPaiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberAnalysisPaiResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberAnalysisPaiResponseBody) *DescribePhoneNumberAnalysisPaiResponse
	GetBody() *DescribePhoneNumberAnalysisPaiResponseBody
}

type DescribePhoneNumberAnalysisPaiResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAnalysisPaiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberAnalysisPaiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisPaiResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisPaiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberAnalysisPaiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberAnalysisPaiResponse) GetBody() *DescribePhoneNumberAnalysisPaiResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberAnalysisPaiResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisPaiResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisPaiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponse) SetBody(v *DescribePhoneNumberAnalysisPaiResponseBody) *DescribePhoneNumberAnalysisPaiResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
