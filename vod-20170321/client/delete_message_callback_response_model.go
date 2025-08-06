// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMessageCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMessageCallbackResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMessageCallbackResponseBody) *DeleteMessageCallbackResponse
	GetBody() *DeleteMessageCallbackResponseBody
}

type DeleteMessageCallbackResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessageCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessageCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageCallbackResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMessageCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessageCallbackResponse) GetBody() *DeleteMessageCallbackResponseBody {
	return s.Body
}

func (s *DeleteMessageCallbackResponse) SetHeaders(v map[string]*string) *DeleteMessageCallbackResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageCallbackResponse) SetStatusCode(v int32) *DeleteMessageCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageCallbackResponse) SetBody(v *DeleteMessageCallbackResponseBody) *DeleteMessageCallbackResponse {
	s.Body = v
	return s
}

func (s *DeleteMessageCallbackResponse) Validate() error {
	return dara.Validate(s)
}
