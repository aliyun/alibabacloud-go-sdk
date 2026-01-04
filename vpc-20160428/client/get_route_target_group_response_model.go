// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRouteTargetGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRouteTargetGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRouteTargetGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetRouteTargetGroupResponseBody) *GetRouteTargetGroupResponse
	GetBody() *GetRouteTargetGroupResponseBody
}

type GetRouteTargetGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRouteTargetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRouteTargetGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRouteTargetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetRouteTargetGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRouteTargetGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRouteTargetGroupResponse) GetBody() *GetRouteTargetGroupResponseBody {
	return s.Body
}

func (s *GetRouteTargetGroupResponse) SetHeaders(v map[string]*string) *GetRouteTargetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetRouteTargetGroupResponse) SetStatusCode(v int32) *GetRouteTargetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRouteTargetGroupResponse) SetBody(v *GetRouteTargetGroupResponseBody) *GetRouteTargetGroupResponse {
	s.Body = v
	return s
}

func (s *GetRouteTargetGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
