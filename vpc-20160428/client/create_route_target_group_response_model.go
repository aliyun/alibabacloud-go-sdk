// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteTargetGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRouteTargetGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRouteTargetGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateRouteTargetGroupResponseBody) *CreateRouteTargetGroupResponse
	GetBody() *CreateRouteTargetGroupResponseBody
}

type CreateRouteTargetGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteTargetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteTargetGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTargetGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteTargetGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRouteTargetGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRouteTargetGroupResponse) GetBody() *CreateRouteTargetGroupResponseBody {
	return s.Body
}

func (s *CreateRouteTargetGroupResponse) SetHeaders(v map[string]*string) *CreateRouteTargetGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteTargetGroupResponse) SetStatusCode(v int32) *CreateRouteTargetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteTargetGroupResponse) SetBody(v *CreateRouteTargetGroupResponseBody) *CreateRouteTargetGroupResponse {
	s.Body = v
	return s
}

func (s *CreateRouteTargetGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
