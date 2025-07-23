// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventRuleResponseBody) *DeleteEventRuleResponse
	GetBody() *DeleteEventRuleResponseBody
}

type DeleteEventRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventRuleResponse) GetBody() *DeleteEventRuleResponseBody {
	return s.Body
}

func (s *DeleteEventRuleResponse) SetHeaders(v map[string]*string) *DeleteEventRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRuleResponse) SetStatusCode(v int32) *DeleteEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventRuleResponse) SetBody(v *DeleteEventRuleResponseBody) *DeleteEventRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteEventRuleResponse) Validate() error {
	return dara.Validate(s)
}
