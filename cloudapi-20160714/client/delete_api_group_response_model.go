// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApiGroupResponseBody) *DeleteApiGroupResponse
	GetBody() *DeleteApiGroupResponseBody
}

type DeleteApiGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApiGroupResponse) GetBody() *DeleteApiGroupResponseBody {
	return s.Body
}

func (s *DeleteApiGroupResponse) SetHeaders(v map[string]*string) *DeleteApiGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiGroupResponse) SetStatusCode(v int32) *DeleteApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiGroupResponse) SetBody(v *DeleteApiGroupResponseBody) *DeleteApiGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteApiGroupResponse) Validate() error {
	return dara.Validate(s)
}
