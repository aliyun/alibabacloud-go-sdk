// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServerLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServerLockResponse
	GetStatusCode() *int32
	SetBody(v *ListServerLockResponseBody) *ListServerLockResponse
	GetBody() *ListServerLockResponseBody
}

type ListServerLockResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServerLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServerLockResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServerLockResponse) GoString() string {
	return s.String()
}

func (s *ListServerLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServerLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServerLockResponse) GetBody() *ListServerLockResponseBody {
	return s.Body
}

func (s *ListServerLockResponse) SetHeaders(v map[string]*string) *ListServerLockResponse {
	s.Headers = v
	return s
}

func (s *ListServerLockResponse) SetStatusCode(v int32) *ListServerLockResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerLockResponse) SetBody(v *ListServerLockResponseBody) *ListServerLockResponse {
	s.Body = v
	return s
}

func (s *ListServerLockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
