// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognizeRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRecognizeRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRecognizeRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRecognizeRuleResponseBody) *DeleteRecognizeRuleResponse
	GetBody() *DeleteRecognizeRuleResponseBody
}

type DeleteRecognizeRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecognizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRecognizeRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognizeRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecognizeRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRecognizeRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRecognizeRuleResponse) GetBody() *DeleteRecognizeRuleResponseBody {
	return s.Body
}

func (s *DeleteRecognizeRuleResponse) SetHeaders(v map[string]*string) *DeleteRecognizeRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecognizeRuleResponse) SetStatusCode(v int32) *DeleteRecognizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecognizeRuleResponse) SetBody(v *DeleteRecognizeRuleResponseBody) *DeleteRecognizeRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRecognizeRuleResponse) Validate() error {
	return dara.Validate(s)
}
