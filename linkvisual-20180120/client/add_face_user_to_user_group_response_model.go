// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserToUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceUserToUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceUserToUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceUserToUserGroupResponseBody) *AddFaceUserToUserGroupResponse
	GetBody() *AddFaceUserToUserGroupResponseBody
}

type AddFaceUserToUserGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserToUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserToUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserToUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceUserToUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceUserToUserGroupResponse) GetBody() *AddFaceUserToUserGroupResponseBody {
	return s.Body
}

func (s *AddFaceUserToUserGroupResponse) SetHeaders(v map[string]*string) *AddFaceUserToUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserToUserGroupResponse) SetStatusCode(v int32) *AddFaceUserToUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserToUserGroupResponse) SetBody(v *AddFaceUserToUserGroupResponseBody) *AddFaceUserToUserGroupResponse {
	s.Body = v
	return s
}

func (s *AddFaceUserToUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
