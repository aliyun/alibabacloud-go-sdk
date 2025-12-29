// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisAIResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisAIResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberAnalysisAIResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberAnalysisAIResponseBody) *DescribePhoneNumberAnalysisAIResponse
	GetBody() *DescribePhoneNumberAnalysisAIResponseBody
}

type DescribePhoneNumberAnalysisAIResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAnalysisAIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberAnalysisAIResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberAnalysisAIResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberAnalysisAIResponse) GetBody() *DescribePhoneNumberAnalysisAIResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberAnalysisAIResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisAIResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisAIResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponse) SetBody(v *DescribePhoneNumberAnalysisAIResponseBody) *DescribePhoneNumberAnalysisAIResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberAnalysisAIResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
