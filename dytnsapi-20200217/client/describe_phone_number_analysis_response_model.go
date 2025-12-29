// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberAnalysisResponseBody) *DescribePhoneNumberAnalysisResponse
	GetBody() *DescribePhoneNumberAnalysisResponseBody
}

type DescribePhoneNumberAnalysisResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberAnalysisResponse) GetBody() *DescribePhoneNumberAnalysisResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberAnalysisResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisResponse) SetBody(v *DescribePhoneNumberAnalysisResponseBody) *DescribePhoneNumberAnalysisResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
