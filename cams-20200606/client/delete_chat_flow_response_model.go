// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChatFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChatFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChatFlowResponseBody) *DeleteChatFlowResponse
	GetBody() *DeleteChatFlowResponseBody
}

type DeleteChatFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChatFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChatFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChatFlowResponse) GetBody() *DeleteChatFlowResponseBody {
	return s.Body
}

func (s *DeleteChatFlowResponse) SetHeaders(v map[string]*string) *DeleteChatFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatFlowResponse) SetStatusCode(v int32) *DeleteChatFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatFlowResponse) SetBody(v *DeleteChatFlowResponseBody) *DeleteChatFlowResponse {
	s.Body = v
	return s
}

func (s *DeleteChatFlowResponse) Validate() error {
	return dara.Validate(s)
}
