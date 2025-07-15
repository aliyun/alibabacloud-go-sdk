// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLiveMessageUsersInGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckLiveMessageUsersInGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckLiveMessageUsersInGroupResponse
	GetStatusCode() *int32
	SetBody(v *CheckLiveMessageUsersInGroupResponseBody) *CheckLiveMessageUsersInGroupResponse
	GetBody() *CheckLiveMessageUsersInGroupResponseBody
}

type CheckLiveMessageUsersInGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckLiveMessageUsersInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckLiveMessageUsersInGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersInGroupResponse) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersInGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckLiveMessageUsersInGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckLiveMessageUsersInGroupResponse) GetBody() *CheckLiveMessageUsersInGroupResponseBody {
	return s.Body
}

func (s *CheckLiveMessageUsersInGroupResponse) SetHeaders(v map[string]*string) *CheckLiveMessageUsersInGroupResponse {
	s.Headers = v
	return s
}

func (s *CheckLiveMessageUsersInGroupResponse) SetStatusCode(v int32) *CheckLiveMessageUsersInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupResponse) SetBody(v *CheckLiveMessageUsersInGroupResponseBody) *CheckLiveMessageUsersInGroupResponse {
	s.Body = v
	return s
}

func (s *CheckLiveMessageUsersInGroupResponse) Validate() error {
	return dara.Validate(s)
}
