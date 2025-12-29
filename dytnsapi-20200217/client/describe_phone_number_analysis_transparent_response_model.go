// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisTransparentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisTransparentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePhoneNumberAnalysisTransparentResponse
	GetStatusCode() *int32
	SetBody(v *DescribePhoneNumberAnalysisTransparentResponseBody) *DescribePhoneNumberAnalysisTransparentResponse
	GetBody() *DescribePhoneNumberAnalysisTransparentResponseBody
}

type DescribePhoneNumberAnalysisTransparentResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePhoneNumberAnalysisTransparentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePhoneNumberAnalysisTransparentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisTransparentResponse) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) GetBody() *DescribePhoneNumberAnalysisTransparentResponseBody {
	return s.Body
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) SetHeaders(v map[string]*string) *DescribePhoneNumberAnalysisTransparentResponse {
	s.Headers = v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) SetStatusCode(v int32) *DescribePhoneNumberAnalysisTransparentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) SetBody(v *DescribePhoneNumberAnalysisTransparentResponseBody) *DescribePhoneNumberAnalysisTransparentResponse {
	s.Body = v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
