// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaslUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSaslUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSaslUserResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSaslUserResponseBody) *DeleteSaslUserResponse
	GetBody() *DeleteSaslUserResponseBody
}

type DeleteSaslUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSaslUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaslUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteSaslUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSaslUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSaslUserResponse) GetBody() *DeleteSaslUserResponseBody {
	return s.Body
}

func (s *DeleteSaslUserResponse) SetHeaders(v map[string]*string) *DeleteSaslUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteSaslUserResponse) SetStatusCode(v int32) *DeleteSaslUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSaslUserResponse) SetBody(v *DeleteSaslUserResponseBody) *DeleteSaslUserResponse {
	s.Body = v
	return s
}

func (s *DeleteSaslUserResponse) Validate() error {
	return dara.Validate(s)
}
