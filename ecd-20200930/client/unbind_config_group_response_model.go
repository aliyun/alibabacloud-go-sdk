// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindConfigGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindConfigGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindConfigGroupResponse
	GetStatusCode() *int32
	SetBody(v *UnbindConfigGroupResponseBody) *UnbindConfigGroupResponse
	GetBody() *UnbindConfigGroupResponseBody
}

type UnbindConfigGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindConfigGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindConfigGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindConfigGroupResponse) GoString() string {
	return s.String()
}

func (s *UnbindConfigGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindConfigGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindConfigGroupResponse) GetBody() *UnbindConfigGroupResponseBody {
	return s.Body
}

func (s *UnbindConfigGroupResponse) SetHeaders(v map[string]*string) *UnbindConfigGroupResponse {
	s.Headers = v
	return s
}

func (s *UnbindConfigGroupResponse) SetStatusCode(v int32) *UnbindConfigGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindConfigGroupResponse) SetBody(v *UnbindConfigGroupResponseBody) *UnbindConfigGroupResponse {
	s.Body = v
	return s
}

func (s *UnbindConfigGroupResponse) Validate() error {
	return dara.Validate(s)
}
