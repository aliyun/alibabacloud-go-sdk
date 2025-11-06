// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMaxYearOfServerLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckMaxYearOfServerLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckMaxYearOfServerLockResponse
	GetStatusCode() *int32
	SetBody(v *CheckMaxYearOfServerLockResponseBody) *CheckMaxYearOfServerLockResponse
	GetBody() *CheckMaxYearOfServerLockResponseBody
}

type CheckMaxYearOfServerLockResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMaxYearOfServerLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMaxYearOfServerLockResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckMaxYearOfServerLockResponse) GoString() string {
	return s.String()
}

func (s *CheckMaxYearOfServerLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckMaxYearOfServerLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckMaxYearOfServerLockResponse) GetBody() *CheckMaxYearOfServerLockResponseBody {
	return s.Body
}

func (s *CheckMaxYearOfServerLockResponse) SetHeaders(v map[string]*string) *CheckMaxYearOfServerLockResponse {
	s.Headers = v
	return s
}

func (s *CheckMaxYearOfServerLockResponse) SetStatusCode(v int32) *CheckMaxYearOfServerLockResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMaxYearOfServerLockResponse) SetBody(v *CheckMaxYearOfServerLockResponseBody) *CheckMaxYearOfServerLockResponse {
	s.Body = v
	return s
}

func (s *CheckMaxYearOfServerLockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
