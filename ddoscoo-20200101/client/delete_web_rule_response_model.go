// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebRuleResponseBody) *DeleteWebRuleResponse
	GetBody() *DeleteWebRuleResponseBody
}

type DeleteWebRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebRuleResponse) GetBody() *DeleteWebRuleResponseBody {
	return s.Body
}

func (s *DeleteWebRuleResponse) SetHeaders(v map[string]*string) *DeleteWebRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebRuleResponse) SetStatusCode(v int32) *DeleteWebRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebRuleResponse) SetBody(v *DeleteWebRuleResponseBody) *DeleteWebRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteWebRuleResponse) Validate() error {
	return dara.Validate(s)
}
