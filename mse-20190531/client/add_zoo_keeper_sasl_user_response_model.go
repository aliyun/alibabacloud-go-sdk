// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZooKeeperSaslUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddZooKeeperSaslUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddZooKeeperSaslUserResponse
	GetStatusCode() *int32
	SetBody(v *AddZooKeeperSaslUserResponseBody) *AddZooKeeperSaslUserResponse
	GetBody() *AddZooKeeperSaslUserResponseBody
}

type AddZooKeeperSaslUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddZooKeeperSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddZooKeeperSaslUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddZooKeeperSaslUserResponse) GoString() string {
	return s.String()
}

func (s *AddZooKeeperSaslUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddZooKeeperSaslUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddZooKeeperSaslUserResponse) GetBody() *AddZooKeeperSaslUserResponseBody {
	return s.Body
}

func (s *AddZooKeeperSaslUserResponse) SetHeaders(v map[string]*string) *AddZooKeeperSaslUserResponse {
	s.Headers = v
	return s
}

func (s *AddZooKeeperSaslUserResponse) SetStatusCode(v int32) *AddZooKeeperSaslUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddZooKeeperSaslUserResponse) SetBody(v *AddZooKeeperSaslUserResponseBody) *AddZooKeeperSaslUserResponse {
	s.Body = v
	return s
}

func (s *AddZooKeeperSaslUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
