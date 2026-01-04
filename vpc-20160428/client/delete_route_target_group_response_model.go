// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteTargetGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteTargetGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteTargetGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteTargetGroupResponseBody) *DeleteRouteTargetGroupResponse
	GetBody() *DeleteRouteTargetGroupResponseBody
}

type DeleteRouteTargetGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteTargetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteTargetGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteTargetGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteTargetGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteTargetGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteTargetGroupResponse) GetBody() *DeleteRouteTargetGroupResponseBody {
	return s.Body
}

func (s *DeleteRouteTargetGroupResponse) SetHeaders(v map[string]*string) *DeleteRouteTargetGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteTargetGroupResponse) SetStatusCode(v int32) *DeleteRouteTargetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteTargetGroupResponse) SetBody(v *DeleteRouteTargetGroupResponseBody) *DeleteRouteTargetGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteTargetGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
