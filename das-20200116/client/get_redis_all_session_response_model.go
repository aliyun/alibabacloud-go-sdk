// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedisAllSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRedisAllSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRedisAllSessionResponse
	GetStatusCode() *int32
	SetBody(v *GetRedisAllSessionResponseBody) *GetRedisAllSessionResponse
	GetBody() *GetRedisAllSessionResponseBody
}

type GetRedisAllSessionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRedisAllSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRedisAllSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRedisAllSessionResponse) GoString() string {
	return s.String()
}

func (s *GetRedisAllSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRedisAllSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRedisAllSessionResponse) GetBody() *GetRedisAllSessionResponseBody {
	return s.Body
}

func (s *GetRedisAllSessionResponse) SetHeaders(v map[string]*string) *GetRedisAllSessionResponse {
	s.Headers = v
	return s
}

func (s *GetRedisAllSessionResponse) SetStatusCode(v int32) *GetRedisAllSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRedisAllSessionResponse) SetBody(v *GetRedisAllSessionResponseBody) *GetRedisAllSessionResponse {
	s.Body = v
	return s
}

func (s *GetRedisAllSessionResponse) Validate() error {
	return dara.Validate(s)
}
