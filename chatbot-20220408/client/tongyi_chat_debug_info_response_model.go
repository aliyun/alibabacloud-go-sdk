// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTongyiChatDebugInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TongyiChatDebugInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TongyiChatDebugInfoResponse
	GetStatusCode() *int32
	SetBody(v *TongyiChatDebugInfoResponseBody) *TongyiChatDebugInfoResponse
	GetBody() *TongyiChatDebugInfoResponseBody
}

type TongyiChatDebugInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TongyiChatDebugInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TongyiChatDebugInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoResponse) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TongyiChatDebugInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TongyiChatDebugInfoResponse) GetBody() *TongyiChatDebugInfoResponseBody {
	return s.Body
}

func (s *TongyiChatDebugInfoResponse) SetHeaders(v map[string]*string) *TongyiChatDebugInfoResponse {
	s.Headers = v
	return s
}

func (s *TongyiChatDebugInfoResponse) SetStatusCode(v int32) *TongyiChatDebugInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *TongyiChatDebugInfoResponse) SetBody(v *TongyiChatDebugInfoResponseBody) *TongyiChatDebugInfoResponse {
	s.Body = v
	return s
}

func (s *TongyiChatDebugInfoResponse) Validate() error {
	return dara.Validate(s)
}
