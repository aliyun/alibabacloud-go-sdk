// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStackGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStackGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStackGroupResponseBody) *UpdateStackGroupResponse
	GetBody() *UpdateStackGroupResponseBody
}

type UpdateStackGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStackGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStackGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStackGroupResponse) GetBody() *UpdateStackGroupResponseBody {
	return s.Body
}

func (s *UpdateStackGroupResponse) SetHeaders(v map[string]*string) *UpdateStackGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackGroupResponse) SetStatusCode(v int32) *UpdateStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStackGroupResponse) SetBody(v *UpdateStackGroupResponseBody) *UpdateStackGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateStackGroupResponse) Validate() error {
	return dara.Validate(s)
}
