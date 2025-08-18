// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaitingRoomRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWaitingRoomRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWaitingRoomRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWaitingRoomRuleResponseBody) *DeleteWaitingRoomRuleResponse
	GetBody() *DeleteWaitingRoomRuleResponseBody
}

type DeleteWaitingRoomRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWaitingRoomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWaitingRoomRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaitingRoomRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWaitingRoomRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWaitingRoomRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWaitingRoomRuleResponse) GetBody() *DeleteWaitingRoomRuleResponseBody {
	return s.Body
}

func (s *DeleteWaitingRoomRuleResponse) SetHeaders(v map[string]*string) *DeleteWaitingRoomRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWaitingRoomRuleResponse) SetStatusCode(v int32) *DeleteWaitingRoomRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWaitingRoomRuleResponse) SetBody(v *DeleteWaitingRoomRuleResponseBody) *DeleteWaitingRoomRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteWaitingRoomRuleResponse) Validate() error {
	return dara.Validate(s)
}
