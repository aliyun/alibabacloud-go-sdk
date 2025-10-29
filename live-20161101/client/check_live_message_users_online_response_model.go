// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLiveMessageUsersOnlineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckLiveMessageUsersOnlineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckLiveMessageUsersOnlineResponse
	GetStatusCode() *int32
	SetBody(v *CheckLiveMessageUsersOnlineResponseBody) *CheckLiveMessageUsersOnlineResponse
	GetBody() *CheckLiveMessageUsersOnlineResponseBody
}

type CheckLiveMessageUsersOnlineResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckLiveMessageUsersOnlineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckLiveMessageUsersOnlineResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersOnlineResponse) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersOnlineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckLiveMessageUsersOnlineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckLiveMessageUsersOnlineResponse) GetBody() *CheckLiveMessageUsersOnlineResponseBody {
	return s.Body
}

func (s *CheckLiveMessageUsersOnlineResponse) SetHeaders(v map[string]*string) *CheckLiveMessageUsersOnlineResponse {
	s.Headers = v
	return s
}

func (s *CheckLiveMessageUsersOnlineResponse) SetStatusCode(v int32) *CheckLiveMessageUsersOnlineResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineResponse) SetBody(v *CheckLiveMessageUsersOnlineResponseBody) *CheckLiveMessageUsersOnlineResponse {
	s.Body = v
	return s
}

func (s *CheckLiveMessageUsersOnlineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
