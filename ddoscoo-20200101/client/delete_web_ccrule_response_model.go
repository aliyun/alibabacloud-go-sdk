// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCCRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebCCRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebCCRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebCCRuleResponseBody) *DeleteWebCCRuleResponse
	GetBody() *DeleteWebCCRuleResponseBody
}

type DeleteWebCCRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebCCRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebCCRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebCCRuleResponse) GetBody() *DeleteWebCCRuleResponseBody {
	return s.Body
}

func (s *DeleteWebCCRuleResponse) SetHeaders(v map[string]*string) *DeleteWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebCCRuleResponse) SetStatusCode(v int32) *DeleteWebCCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebCCRuleResponse) SetBody(v *DeleteWebCCRuleResponseBody) *DeleteWebCCRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteWebCCRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
