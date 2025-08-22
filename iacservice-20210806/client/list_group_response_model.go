// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupResponseBody) *ListGroupResponse
	GetBody() *ListGroupResponseBody
}

type ListGroupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupResponse) GoString() string {
	return s.String()
}

func (s *ListGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupResponse) GetBody() *ListGroupResponseBody {
	return s.Body
}

func (s *ListGroupResponse) SetHeaders(v map[string]*string) *ListGroupResponse {
	s.Headers = v
	return s
}

func (s *ListGroupResponse) SetStatusCode(v int32) *ListGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupResponse) SetBody(v *ListGroupResponseBody) *ListGroupResponse {
	s.Body = v
	return s
}

func (s *ListGroupResponse) Validate() error {
	return dara.Validate(s)
}
