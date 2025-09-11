// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendBatchSmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendBatchSmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendBatchSmsResponse
	GetStatusCode() *int32
	SetBody(v *SendBatchSmsResponseBody) *SendBatchSmsResponse
	GetBody() *SendBatchSmsResponseBody
}

type SendBatchSmsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendBatchSmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendBatchSmsResponse) String() string {
	return dara.Prettify(s)
}

func (s SendBatchSmsResponse) GoString() string {
	return s.String()
}

func (s *SendBatchSmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendBatchSmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendBatchSmsResponse) GetBody() *SendBatchSmsResponseBody {
	return s.Body
}

func (s *SendBatchSmsResponse) SetHeaders(v map[string]*string) *SendBatchSmsResponse {
	s.Headers = v
	return s
}

func (s *SendBatchSmsResponse) SetStatusCode(v int32) *SendBatchSmsResponse {
	s.StatusCode = &v
	return s
}

func (s *SendBatchSmsResponse) SetBody(v *SendBatchSmsResponseBody) *SendBatchSmsResponse {
	s.Body = v
	return s
}

func (s *SendBatchSmsResponse) Validate() error {
	return dara.Validate(s)
}
