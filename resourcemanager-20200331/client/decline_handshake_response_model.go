// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeclineHandshakeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeclineHandshakeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeclineHandshakeResponse
	GetStatusCode() *int32
	SetBody(v *DeclineHandshakeResponseBody) *DeclineHandshakeResponse
	GetBody() *DeclineHandshakeResponseBody
}

type DeclineHandshakeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeclineHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeclineHandshakeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeclineHandshakeResponse) GoString() string {
	return s.String()
}

func (s *DeclineHandshakeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeclineHandshakeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeclineHandshakeResponse) GetBody() *DeclineHandshakeResponseBody {
	return s.Body
}

func (s *DeclineHandshakeResponse) SetHeaders(v map[string]*string) *DeclineHandshakeResponse {
	s.Headers = v
	return s
}

func (s *DeclineHandshakeResponse) SetStatusCode(v int32) *DeclineHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeclineHandshakeResponse) SetBody(v *DeclineHandshakeResponseBody) *DeclineHandshakeResponse {
	s.Body = v
	return s
}

func (s *DeclineHandshakeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
