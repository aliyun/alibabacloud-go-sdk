// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageToGlobeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMessageToGlobeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMessageToGlobeResponse
	GetStatusCode() *int32
	SetBody(v *SendMessageToGlobeResponseBody) *SendMessageToGlobeResponse
	GetBody() *SendMessageToGlobeResponseBody
}

type SendMessageToGlobeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageToGlobeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageToGlobeResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMessageToGlobeResponse) GoString() string {
	return s.String()
}

func (s *SendMessageToGlobeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMessageToGlobeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMessageToGlobeResponse) GetBody() *SendMessageToGlobeResponseBody {
	return s.Body
}

func (s *SendMessageToGlobeResponse) SetHeaders(v map[string]*string) *SendMessageToGlobeResponse {
	s.Headers = v
	return s
}

func (s *SendMessageToGlobeResponse) SetStatusCode(v int32) *SendMessageToGlobeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageToGlobeResponse) SetBody(v *SendMessageToGlobeResponseBody) *SendMessageToGlobeResponse {
	s.Body = v
	return s
}

func (s *SendMessageToGlobeResponse) Validate() error {
	return dara.Validate(s)
}
