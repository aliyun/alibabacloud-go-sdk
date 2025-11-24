// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSwimLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSwimLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSwimLaneGroupResponseBody) *UpdateSwimLaneGroupResponse
	GetBody() *UpdateSwimLaneGroupResponseBody
}

type UpdateSwimLaneGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSwimLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSwimLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSwimLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSwimLaneGroupResponse) GetBody() *UpdateSwimLaneGroupResponseBody {
	return s.Body
}

func (s *UpdateSwimLaneGroupResponse) SetHeaders(v map[string]*string) *UpdateSwimLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateSwimLaneGroupResponse) SetStatusCode(v int32) *UpdateSwimLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSwimLaneGroupResponse) SetBody(v *UpdateSwimLaneGroupResponseBody) *UpdateSwimLaneGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateSwimLaneGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
