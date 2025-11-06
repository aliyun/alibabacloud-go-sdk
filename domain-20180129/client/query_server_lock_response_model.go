// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryServerLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryServerLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryServerLockResponse
	GetStatusCode() *int32
	SetBody(v *QueryServerLockResponseBody) *QueryServerLockResponse
	GetBody() *QueryServerLockResponseBody
}

type QueryServerLockResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryServerLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryServerLockResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryServerLockResponse) GoString() string {
	return s.String()
}

func (s *QueryServerLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryServerLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryServerLockResponse) GetBody() *QueryServerLockResponseBody {
	return s.Body
}

func (s *QueryServerLockResponse) SetHeaders(v map[string]*string) *QueryServerLockResponse {
	s.Headers = v
	return s
}

func (s *QueryServerLockResponse) SetStatusCode(v int32) *QueryServerLockResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryServerLockResponse) SetBody(v *QueryServerLockResponseBody) *QueryServerLockResponse {
	s.Body = v
	return s
}

func (s *QueryServerLockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
