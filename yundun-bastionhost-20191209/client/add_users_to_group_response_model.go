// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUsersToGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUsersToGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddUsersToGroupResponseBody) *AddUsersToGroupResponse
	GetBody() *AddUsersToGroupResponseBody
}

type AddUsersToGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUsersToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUsersToGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUsersToGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUsersToGroupResponse) GetBody() *AddUsersToGroupResponseBody {
	return s.Body
}

func (s *AddUsersToGroupResponse) SetHeaders(v map[string]*string) *AddUsersToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUsersToGroupResponse) SetStatusCode(v int32) *AddUsersToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUsersToGroupResponse) SetBody(v *AddUsersToGroupResponseBody) *AddUsersToGroupResponse {
	s.Body = v
	return s
}

func (s *AddUsersToGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
