// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendConsoleMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendConsoleMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendConsoleMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendConsoleMessageResponseBody) *SendConsoleMessageResponse
	GetBody() *SendConsoleMessageResponseBody
}

type SendConsoleMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendConsoleMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendConsoleMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendConsoleMessageResponse) GoString() string {
	return s.String()
}

func (s *SendConsoleMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendConsoleMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendConsoleMessageResponse) GetBody() *SendConsoleMessageResponseBody {
	return s.Body
}

func (s *SendConsoleMessageResponse) SetHeaders(v map[string]*string) *SendConsoleMessageResponse {
	s.Headers = v
	return s
}

func (s *SendConsoleMessageResponse) SetStatusCode(v int32) *SendConsoleMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendConsoleMessageResponse) SetBody(v *SendConsoleMessageResponseBody) *SendConsoleMessageResponse {
	s.Body = v
	return s
}

func (s *SendConsoleMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
