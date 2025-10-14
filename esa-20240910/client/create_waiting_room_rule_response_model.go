// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWaitingRoomRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWaitingRoomRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWaitingRoomRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateWaitingRoomRuleResponseBody) *CreateWaitingRoomRuleResponse
	GetBody() *CreateWaitingRoomRuleResponseBody
}

type CreateWaitingRoomRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWaitingRoomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWaitingRoomRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWaitingRoomRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateWaitingRoomRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWaitingRoomRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWaitingRoomRuleResponse) GetBody() *CreateWaitingRoomRuleResponseBody {
	return s.Body
}

func (s *CreateWaitingRoomRuleResponse) SetHeaders(v map[string]*string) *CreateWaitingRoomRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateWaitingRoomRuleResponse) SetStatusCode(v int32) *CreateWaitingRoomRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWaitingRoomRuleResponse) SetBody(v *CreateWaitingRoomRuleResponseBody) *CreateWaitingRoomRuleResponse {
	s.Body = v
	return s
}

func (s *CreateWaitingRoomRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
