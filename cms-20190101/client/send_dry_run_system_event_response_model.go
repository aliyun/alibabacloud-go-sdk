// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDryRunSystemEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendDryRunSystemEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendDryRunSystemEventResponse
	GetStatusCode() *int32
	SetBody(v *SendDryRunSystemEventResponseBody) *SendDryRunSystemEventResponse
	GetBody() *SendDryRunSystemEventResponseBody
}

type SendDryRunSystemEventResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendDryRunSystemEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendDryRunSystemEventResponse) String() string {
	return dara.Prettify(s)
}

func (s SendDryRunSystemEventResponse) GoString() string {
	return s.String()
}

func (s *SendDryRunSystemEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendDryRunSystemEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendDryRunSystemEventResponse) GetBody() *SendDryRunSystemEventResponseBody {
	return s.Body
}

func (s *SendDryRunSystemEventResponse) SetHeaders(v map[string]*string) *SendDryRunSystemEventResponse {
	s.Headers = v
	return s
}

func (s *SendDryRunSystemEventResponse) SetStatusCode(v int32) *SendDryRunSystemEventResponse {
	s.StatusCode = &v
	return s
}

func (s *SendDryRunSystemEventResponse) SetBody(v *SendDryRunSystemEventResponseBody) *SendDryRunSystemEventResponse {
	s.Body = v
	return s
}

func (s *SendDryRunSystemEventResponse) Validate() error {
	return dara.Validate(s)
}
