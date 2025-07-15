// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBgpGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBgpGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBgpGroupResponseBody) *DeleteBgpGroupResponse
	GetBody() *DeleteBgpGroupResponseBody
}

type DeleteBgpGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBgpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBgpGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteBgpGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBgpGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBgpGroupResponse) GetBody() *DeleteBgpGroupResponseBody {
	return s.Body
}

func (s *DeleteBgpGroupResponse) SetHeaders(v map[string]*string) *DeleteBgpGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteBgpGroupResponse) SetStatusCode(v int32) *DeleteBgpGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBgpGroupResponse) SetBody(v *DeleteBgpGroupResponseBody) *DeleteBgpGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteBgpGroupResponse) Validate() error {
	return dara.Validate(s)
}
