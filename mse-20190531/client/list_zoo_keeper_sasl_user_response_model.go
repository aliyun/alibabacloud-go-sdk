// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZooKeeperSaslUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListZooKeeperSaslUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListZooKeeperSaslUserResponse
	GetStatusCode() *int32
	SetBody(v *ListZooKeeperSaslUserResponseBody) *ListZooKeeperSaslUserResponse
	GetBody() *ListZooKeeperSaslUserResponseBody
}

type ListZooKeeperSaslUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListZooKeeperSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListZooKeeperSaslUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListZooKeeperSaslUserResponse) GoString() string {
	return s.String()
}

func (s *ListZooKeeperSaslUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListZooKeeperSaslUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListZooKeeperSaslUserResponse) GetBody() *ListZooKeeperSaslUserResponseBody {
	return s.Body
}

func (s *ListZooKeeperSaslUserResponse) SetHeaders(v map[string]*string) *ListZooKeeperSaslUserResponse {
	s.Headers = v
	return s
}

func (s *ListZooKeeperSaslUserResponse) SetStatusCode(v int32) *ListZooKeeperSaslUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListZooKeeperSaslUserResponse) SetBody(v *ListZooKeeperSaslUserResponseBody) *ListZooKeeperSaslUserResponse {
	s.Body = v
	return s
}

func (s *ListZooKeeperSaslUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
