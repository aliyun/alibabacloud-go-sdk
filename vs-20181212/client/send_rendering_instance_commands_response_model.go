// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRenderingInstanceCommandsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendRenderingInstanceCommandsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendRenderingInstanceCommandsResponse
	GetStatusCode() *int32
	SetBody(v *SendRenderingInstanceCommandsResponseBody) *SendRenderingInstanceCommandsResponse
	GetBody() *SendRenderingInstanceCommandsResponseBody
}

type SendRenderingInstanceCommandsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendRenderingInstanceCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendRenderingInstanceCommandsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendRenderingInstanceCommandsResponse) GoString() string {
	return s.String()
}

func (s *SendRenderingInstanceCommandsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendRenderingInstanceCommandsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendRenderingInstanceCommandsResponse) GetBody() *SendRenderingInstanceCommandsResponseBody {
	return s.Body
}

func (s *SendRenderingInstanceCommandsResponse) SetHeaders(v map[string]*string) *SendRenderingInstanceCommandsResponse {
	s.Headers = v
	return s
}

func (s *SendRenderingInstanceCommandsResponse) SetStatusCode(v int32) *SendRenderingInstanceCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendRenderingInstanceCommandsResponse) SetBody(v *SendRenderingInstanceCommandsResponseBody) *SendRenderingInstanceCommandsResponse {
	s.Body = v
	return s
}

func (s *SendRenderingInstanceCommandsResponse) Validate() error {
	return dara.Validate(s)
}
