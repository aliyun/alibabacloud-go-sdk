// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUsersFromGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUsersFromGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUsersFromGroupResponseBody) *RemoveUsersFromGroupResponse
	GetBody() *RemoveUsersFromGroupResponseBody
}

type RemoveUsersFromGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUsersFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUsersFromGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUsersFromGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUsersFromGroupResponse) GetBody() *RemoveUsersFromGroupResponseBody {
	return s.Body
}

func (s *RemoveUsersFromGroupResponse) SetHeaders(v map[string]*string) *RemoveUsersFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersFromGroupResponse) SetStatusCode(v int32) *RemoveUsersFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUsersFromGroupResponse) SetBody(v *RemoveUsersFromGroupResponseBody) *RemoveUsersFromGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveUsersFromGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
