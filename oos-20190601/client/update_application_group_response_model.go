// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationGroupResponseBody) *UpdateApplicationGroupResponse
	GetBody() *UpdateApplicationGroupResponseBody
}

type UpdateApplicationGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationGroupResponse) GetBody() *UpdateApplicationGroupResponseBody {
	return s.Body
}

func (s *UpdateApplicationGroupResponse) SetHeaders(v map[string]*string) *UpdateApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationGroupResponse) SetStatusCode(v int32) *UpdateApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationGroupResponse) SetBody(v *UpdateApplicationGroupResponseBody) *UpdateApplicationGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationGroupResponse) Validate() error {
	return dara.Validate(s)
}
