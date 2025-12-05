// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserGroupResponseBody) *ModifyUserGroupResponse
	GetBody() *ModifyUserGroupResponseBody
}

type ModifyUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserGroupResponse) GetBody() *ModifyUserGroupResponseBody {
	return s.Body
}

func (s *ModifyUserGroupResponse) SetHeaders(v map[string]*string) *ModifyUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserGroupResponse) SetStatusCode(v int32) *ModifyUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserGroupResponse) SetBody(v *ModifyUserGroupResponseBody) *ModifyUserGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
