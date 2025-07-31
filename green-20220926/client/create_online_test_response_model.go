// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOnlineTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOnlineTestResponse
	GetStatusCode() *int32
	SetBody(v *CreateOnlineTestResponseBody) *CreateOnlineTestResponse
	GetBody() *CreateOnlineTestResponseBody
}

type CreateOnlineTestResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOnlineTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOnlineTestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineTestResponse) GoString() string {
	return s.String()
}

func (s *CreateOnlineTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOnlineTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOnlineTestResponse) GetBody() *CreateOnlineTestResponseBody {
	return s.Body
}

func (s *CreateOnlineTestResponse) SetHeaders(v map[string]*string) *CreateOnlineTestResponse {
	s.Headers = v
	return s
}

func (s *CreateOnlineTestResponse) SetStatusCode(v int32) *CreateOnlineTestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOnlineTestResponse) SetBody(v *CreateOnlineTestResponseBody) *CreateOnlineTestResponse {
	s.Body = v
	return s
}

func (s *CreateOnlineTestResponse) Validate() error {
	return dara.Validate(s)
}
