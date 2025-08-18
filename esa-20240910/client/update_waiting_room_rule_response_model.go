// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWaitingRoomRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWaitingRoomRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWaitingRoomRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWaitingRoomRuleResponseBody) *UpdateWaitingRoomRuleResponse
	GetBody() *UpdateWaitingRoomRuleResponseBody
}

type UpdateWaitingRoomRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWaitingRoomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWaitingRoomRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWaitingRoomRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaitingRoomRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWaitingRoomRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWaitingRoomRuleResponse) GetBody() *UpdateWaitingRoomRuleResponseBody {
	return s.Body
}

func (s *UpdateWaitingRoomRuleResponse) SetHeaders(v map[string]*string) *UpdateWaitingRoomRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWaitingRoomRuleResponse) SetStatusCode(v int32) *UpdateWaitingRoomRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWaitingRoomRuleResponse) SetBody(v *UpdateWaitingRoomRuleResponseBody) *UpdateWaitingRoomRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateWaitingRoomRuleResponse) Validate() error {
	return dara.Validate(s)
}
