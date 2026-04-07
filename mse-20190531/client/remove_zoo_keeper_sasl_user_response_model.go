// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveZooKeeperSaslUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveZooKeeperSaslUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveZooKeeperSaslUserResponse
	GetStatusCode() *int32
	SetBody(v *RemoveZooKeeperSaslUserResponseBody) *RemoveZooKeeperSaslUserResponse
	GetBody() *RemoveZooKeeperSaslUserResponseBody
}

type RemoveZooKeeperSaslUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveZooKeeperSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveZooKeeperSaslUserResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveZooKeeperSaslUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveZooKeeperSaslUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveZooKeeperSaslUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveZooKeeperSaslUserResponse) GetBody() *RemoveZooKeeperSaslUserResponseBody {
	return s.Body
}

func (s *RemoveZooKeeperSaslUserResponse) SetHeaders(v map[string]*string) *RemoveZooKeeperSaslUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveZooKeeperSaslUserResponse) SetStatusCode(v int32) *RemoveZooKeeperSaslUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveZooKeeperSaslUserResponse) SetBody(v *RemoveZooKeeperSaslUserResponseBody) *RemoveZooKeeperSaslUserResponse {
	s.Body = v
	return s
}

func (s *RemoveZooKeeperSaslUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
