// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillInstanceAllSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KillInstanceAllSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KillInstanceAllSessionResponse
	GetStatusCode() *int32
	SetBody(v *KillInstanceAllSessionResponseBody) *KillInstanceAllSessionResponse
	GetBody() *KillInstanceAllSessionResponseBody
}

type KillInstanceAllSessionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillInstanceAllSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillInstanceAllSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s KillInstanceAllSessionResponse) GoString() string {
	return s.String()
}

func (s *KillInstanceAllSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KillInstanceAllSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KillInstanceAllSessionResponse) GetBody() *KillInstanceAllSessionResponseBody {
	return s.Body
}

func (s *KillInstanceAllSessionResponse) SetHeaders(v map[string]*string) *KillInstanceAllSessionResponse {
	s.Headers = v
	return s
}

func (s *KillInstanceAllSessionResponse) SetStatusCode(v int32) *KillInstanceAllSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *KillInstanceAllSessionResponse) SetBody(v *KillInstanceAllSessionResponseBody) *KillInstanceAllSessionResponse {
	s.Body = v
	return s
}

func (s *KillInstanceAllSessionResponse) Validate() error {
	return dara.Validate(s)
}
