// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventCenterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventCenterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventCenterRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventCenterRuleResponseBody) *DeleteEventCenterRuleResponse
	GetBody() *DeleteEventCenterRuleResponseBody
}

type DeleteEventCenterRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventCenterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventCenterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventCenterRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventCenterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventCenterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventCenterRuleResponse) GetBody() *DeleteEventCenterRuleResponseBody {
	return s.Body
}

func (s *DeleteEventCenterRuleResponse) SetHeaders(v map[string]*string) *DeleteEventCenterRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventCenterRuleResponse) SetStatusCode(v int32) *DeleteEventCenterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventCenterRuleResponse) SetBody(v *DeleteEventCenterRuleResponseBody) *DeleteEventCenterRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteEventCenterRuleResponse) Validate() error {
	return dara.Validate(s)
}
