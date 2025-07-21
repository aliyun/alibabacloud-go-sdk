// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerChatFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerChatFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerChatFlowResponse
	GetStatusCode() *int32
	SetBody(v *TriggerChatFlowResponseBody) *TriggerChatFlowResponse
	GetBody() *TriggerChatFlowResponseBody
}

type TriggerChatFlowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerChatFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerChatFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerChatFlowResponse) GoString() string {
	return s.String()
}

func (s *TriggerChatFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerChatFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerChatFlowResponse) GetBody() *TriggerChatFlowResponseBody {
	return s.Body
}

func (s *TriggerChatFlowResponse) SetHeaders(v map[string]*string) *TriggerChatFlowResponse {
	s.Headers = v
	return s
}

func (s *TriggerChatFlowResponse) SetStatusCode(v int32) *TriggerChatFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerChatFlowResponse) SetBody(v *TriggerChatFlowResponseBody) *TriggerChatFlowResponse {
	s.Body = v
	return s
}

func (s *TriggerChatFlowResponse) Validate() error {
	return dara.Validate(s)
}
