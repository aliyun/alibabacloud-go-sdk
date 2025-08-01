// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSwimmingLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSwimmingLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSwimmingLaneGroupResponseBody) *DeleteSwimmingLaneGroupResponse
	GetBody() *DeleteSwimmingLaneGroupResponseBody
}

type DeleteSwimmingLaneGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSwimmingLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSwimmingLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSwimmingLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSwimmingLaneGroupResponse) GetBody() *DeleteSwimmingLaneGroupResponseBody {
	return s.Body
}

func (s *DeleteSwimmingLaneGroupResponse) SetHeaders(v map[string]*string) *DeleteSwimmingLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSwimmingLaneGroupResponse) SetStatusCode(v int32) *DeleteSwimmingLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponse) SetBody(v *DeleteSwimmingLaneGroupResponseBody) *DeleteSwimmingLaneGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteSwimmingLaneGroupResponse) Validate() error {
	return dara.Validate(s)
}
