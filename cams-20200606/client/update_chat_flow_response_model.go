// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChatFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChatFlowResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChatFlowResponseBody) *UpdateChatFlowResponse
	GetBody() *UpdateChatFlowResponseBody
}

type UpdateChatFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChatFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChatFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatFlowResponse) GoString() string {
	return s.String()
}

func (s *UpdateChatFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChatFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChatFlowResponse) GetBody() *UpdateChatFlowResponseBody {
	return s.Body
}

func (s *UpdateChatFlowResponse) SetHeaders(v map[string]*string) *UpdateChatFlowResponse {
	s.Headers = v
	return s
}

func (s *UpdateChatFlowResponse) SetStatusCode(v int32) *UpdateChatFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChatFlowResponse) SetBody(v *UpdateChatFlowResponseBody) *UpdateChatFlowResponse {
	s.Body = v
	return s
}

func (s *UpdateChatFlowResponse) Validate() error {
	return dara.Validate(s)
}
