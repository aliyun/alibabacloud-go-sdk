// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGroupResponseBody) *UpdateGroupResponse
	GetBody() *UpdateGroupResponseBody
}

type UpdateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGroupResponse) GetBody() *UpdateGroupResponseBody {
	return s.Body
}

func (s *UpdateGroupResponse) SetHeaders(v map[string]*string) *UpdateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupResponse) SetStatusCode(v int32) *UpdateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupResponse) SetBody(v *UpdateGroupResponseBody) *UpdateGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateGroupResponse) Validate() error {
	return dara.Validate(s)
}
