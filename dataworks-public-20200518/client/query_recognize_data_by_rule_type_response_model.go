// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognizeDataByRuleTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecognizeDataByRuleTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecognizeDataByRuleTypeResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecognizeDataByRuleTypeResponseBody) *QueryRecognizeDataByRuleTypeResponse
	GetBody() *QueryRecognizeDataByRuleTypeResponseBody
}

type QueryRecognizeDataByRuleTypeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecognizeDataByRuleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecognizeDataByRuleTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognizeDataByRuleTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryRecognizeDataByRuleTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecognizeDataByRuleTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecognizeDataByRuleTypeResponse) GetBody() *QueryRecognizeDataByRuleTypeResponseBody {
	return s.Body
}

func (s *QueryRecognizeDataByRuleTypeResponse) SetHeaders(v map[string]*string) *QueryRecognizeDataByRuleTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponse) SetStatusCode(v int32) *QueryRecognizeDataByRuleTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponse) SetBody(v *QueryRecognizeDataByRuleTypeResponseBody) *QueryRecognizeDataByRuleTypeResponse {
	s.Body = v
	return s
}

func (s *QueryRecognizeDataByRuleTypeResponse) Validate() error {
	return dara.Validate(s)
}
