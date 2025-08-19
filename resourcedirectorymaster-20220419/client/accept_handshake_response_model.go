// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptHandshakeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptHandshakeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptHandshakeResponse
	GetStatusCode() *int32
	SetBody(v *AcceptHandshakeResponseBody) *AcceptHandshakeResponse
	GetBody() *AcceptHandshakeResponseBody
}

type AcceptHandshakeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptHandshakeResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptHandshakeResponse) GoString() string {
	return s.String()
}

func (s *AcceptHandshakeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptHandshakeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptHandshakeResponse) GetBody() *AcceptHandshakeResponseBody {
	return s.Body
}

func (s *AcceptHandshakeResponse) SetHeaders(v map[string]*string) *AcceptHandshakeResponse {
	s.Headers = v
	return s
}

func (s *AcceptHandshakeResponse) SetStatusCode(v int32) *AcceptHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptHandshakeResponse) SetBody(v *AcceptHandshakeResponseBody) *AcceptHandshakeResponse {
	s.Body = v
	return s
}

func (s *AcceptHandshakeResponse) Validate() error {
	return dara.Validate(s)
}
