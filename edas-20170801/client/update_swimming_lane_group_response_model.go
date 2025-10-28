// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSwimmingLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSwimmingLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSwimmingLaneGroupResponseBody) *UpdateSwimmingLaneGroupResponse
	GetBody() *UpdateSwimmingLaneGroupResponseBody
}

type UpdateSwimmingLaneGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSwimmingLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSwimmingLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSwimmingLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSwimmingLaneGroupResponse) GetBody() *UpdateSwimmingLaneGroupResponseBody {
	return s.Body
}

func (s *UpdateSwimmingLaneGroupResponse) SetHeaders(v map[string]*string) *UpdateSwimmingLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateSwimmingLaneGroupResponse) SetStatusCode(v int32) *UpdateSwimmingLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSwimmingLaneGroupResponse) SetBody(v *UpdateSwimmingLaneGroupResponseBody) *UpdateSwimmingLaneGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateSwimmingLaneGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
