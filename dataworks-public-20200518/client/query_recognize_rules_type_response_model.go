// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognizeRulesTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRecognizeRulesTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRecognizeRulesTypeResponse
	GetStatusCode() *int32
	SetBody(v *QueryRecognizeRulesTypeResponseBody) *QueryRecognizeRulesTypeResponse
	GetBody() *QueryRecognizeRulesTypeResponseBody
}

type QueryRecognizeRulesTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecognizeRulesTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecognizeRulesTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognizeRulesTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryRecognizeRulesTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRecognizeRulesTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRecognizeRulesTypeResponse) GetBody() *QueryRecognizeRulesTypeResponseBody {
	return s.Body
}

func (s *QueryRecognizeRulesTypeResponse) SetHeaders(v map[string]*string) *QueryRecognizeRulesTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryRecognizeRulesTypeResponse) SetStatusCode(v int32) *QueryRecognizeRulesTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecognizeRulesTypeResponse) SetBody(v *QueryRecognizeRulesTypeResponseBody) *QueryRecognizeRulesTypeResponse {
	s.Body = v
	return s
}

func (s *QueryRecognizeRulesTypeResponse) Validate() error {
	return dara.Validate(s)
}
