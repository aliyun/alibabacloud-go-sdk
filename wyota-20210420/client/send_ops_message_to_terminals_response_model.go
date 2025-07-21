// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendOpsMessageToTerminalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendOpsMessageToTerminalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendOpsMessageToTerminalsResponse
	GetStatusCode() *int32
	SetBody(v *SendOpsMessageToTerminalsResponseBody) *SendOpsMessageToTerminalsResponse
	GetBody() *SendOpsMessageToTerminalsResponseBody
}

type SendOpsMessageToTerminalsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendOpsMessageToTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendOpsMessageToTerminalsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendOpsMessageToTerminalsResponse) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendOpsMessageToTerminalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendOpsMessageToTerminalsResponse) GetBody() *SendOpsMessageToTerminalsResponseBody {
	return s.Body
}

func (s *SendOpsMessageToTerminalsResponse) SetHeaders(v map[string]*string) *SendOpsMessageToTerminalsResponse {
	s.Headers = v
	return s
}

func (s *SendOpsMessageToTerminalsResponse) SetStatusCode(v int32) *SendOpsMessageToTerminalsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendOpsMessageToTerminalsResponse) SetBody(v *SendOpsMessageToTerminalsResponseBody) *SendOpsMessageToTerminalsResponse {
	s.Body = v
	return s
}

func (s *SendOpsMessageToTerminalsResponse) Validate() error {
	return dara.Validate(s)
}
