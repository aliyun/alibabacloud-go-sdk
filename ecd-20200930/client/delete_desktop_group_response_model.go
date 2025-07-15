// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDesktopGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDesktopGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDesktopGroupResponseBody) *DeleteDesktopGroupResponse
	GetBody() *DeleteDesktopGroupResponseBody
}

type DeleteDesktopGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDesktopGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDesktopGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDesktopGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDesktopGroupResponse) GetBody() *DeleteDesktopGroupResponseBody {
	return s.Body
}

func (s *DeleteDesktopGroupResponse) SetHeaders(v map[string]*string) *DeleteDesktopGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDesktopGroupResponse) SetStatusCode(v int32) *DeleteDesktopGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDesktopGroupResponse) SetBody(v *DeleteDesktopGroupResponseBody) *DeleteDesktopGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDesktopGroupResponse) Validate() error {
	return dara.Validate(s)
}
