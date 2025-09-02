// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognizeRuleDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecognizeRuleDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecognizeRuleDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecognizeRuleDetailResponseBody) *QueryRecognizeRuleDetailResponse
	GetBody() *QueryRecognizeRuleDetailResponseBody
}

type QueryRecognizeRuleDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecognizeRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecognizeRuleDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognizeRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRecognizeRuleDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecognizeRuleDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecognizeRuleDetailResponse) GetBody() *QueryRecognizeRuleDetailResponseBody {
	return s.Body
}

func (s *QueryRecognizeRuleDetailResponse) SetHeaders(v map[string]*string) *QueryRecognizeRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRecognizeRuleDetailResponse) SetStatusCode(v int32) *QueryRecognizeRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecognizeRuleDetailResponse) SetBody(v *QueryRecognizeRuleDetailResponseBody) *QueryRecognizeRuleDetailResponse {
	s.Body = v
	return s
}

func (s *QueryRecognizeRuleDetailResponse) Validate() error {
	return dara.Validate(s)
}
