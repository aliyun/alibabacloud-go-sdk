// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZooKeeperSaslUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateZooKeeperSaslUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateZooKeeperSaslUserResponse
	GetStatusCode() *int32
	SetBody(v *UpdateZooKeeperSaslUserResponseBody) *UpdateZooKeeperSaslUserResponse
	GetBody() *UpdateZooKeeperSaslUserResponseBody
}

type UpdateZooKeeperSaslUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateZooKeeperSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateZooKeeperSaslUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateZooKeeperSaslUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateZooKeeperSaslUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateZooKeeperSaslUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateZooKeeperSaslUserResponse) GetBody() *UpdateZooKeeperSaslUserResponseBody {
	return s.Body
}

func (s *UpdateZooKeeperSaslUserResponse) SetHeaders(v map[string]*string) *UpdateZooKeeperSaslUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateZooKeeperSaslUserResponse) SetStatusCode(v int32) *UpdateZooKeeperSaslUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateZooKeeperSaslUserResponse) SetBody(v *UpdateZooKeeperSaslUserResponseBody) *UpdateZooKeeperSaslUserResponse {
	s.Body = v
	return s
}

func (s *UpdateZooKeeperSaslUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
