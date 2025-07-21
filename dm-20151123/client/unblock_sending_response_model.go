// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnblockSendingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnblockSendingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnblockSendingResponse
	GetStatusCode() *int32
	SetBody(v *UnblockSendingResponseBody) *UnblockSendingResponse
	GetBody() *UnblockSendingResponseBody
}

type UnblockSendingResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnblockSendingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnblockSendingResponse) String() string {
	return dara.Prettify(s)
}

func (s UnblockSendingResponse) GoString() string {
	return s.String()
}

func (s *UnblockSendingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnblockSendingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnblockSendingResponse) GetBody() *UnblockSendingResponseBody {
	return s.Body
}

func (s *UnblockSendingResponse) SetHeaders(v map[string]*string) *UnblockSendingResponse {
	s.Headers = v
	return s
}

func (s *UnblockSendingResponse) SetStatusCode(v int32) *UnblockSendingResponse {
	s.StatusCode = &v
	return s
}

func (s *UnblockSendingResponse) SetBody(v *UnblockSendingResponseBody) *UnblockSendingResponse {
	s.Body = v
	return s
}

func (s *UnblockSendingResponse) Validate() error {
	return dara.Validate(s)
}
