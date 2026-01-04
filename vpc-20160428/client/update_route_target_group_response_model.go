// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRouteTargetGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRouteTargetGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRouteTargetGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRouteTargetGroupResponseBody) *UpdateRouteTargetGroupResponse
	GetBody() *UpdateRouteTargetGroupResponseBody
}

type UpdateRouteTargetGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRouteTargetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRouteTargetGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRouteTargetGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateRouteTargetGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRouteTargetGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRouteTargetGroupResponse) GetBody() *UpdateRouteTargetGroupResponseBody {
	return s.Body
}

func (s *UpdateRouteTargetGroupResponse) SetHeaders(v map[string]*string) *UpdateRouteTargetGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateRouteTargetGroupResponse) SetStatusCode(v int32) *UpdateRouteTargetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRouteTargetGroupResponse) SetBody(v *UpdateRouteTargetGroupResponseBody) *UpdateRouteTargetGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateRouteTargetGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
