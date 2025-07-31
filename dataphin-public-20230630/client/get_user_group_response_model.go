// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetUserGroupResponseBody) *GetUserGroupResponse
	GetBody() *GetUserGroupResponseBody
}

type GetUserGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserGroupResponse) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserGroupResponse) GetBody() *GetUserGroupResponseBody {
	return s.Body
}

func (s *GetUserGroupResponse) SetHeaders(v map[string]*string) *GetUserGroupResponse {
	s.Headers = v
	return s
}

func (s *GetUserGroupResponse) SetStatusCode(v int32) *GetUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGroupResponse) SetBody(v *GetUserGroupResponseBody) *GetUserGroupResponse {
	s.Body = v
	return s
}

func (s *GetUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
