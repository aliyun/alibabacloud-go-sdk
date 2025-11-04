// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScalingGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScalingGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScalingGroupResponseBody) *DeleteScalingGroupResponse
	GetBody() *DeleteScalingGroupResponseBody
}

type DeleteScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScalingGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScalingGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScalingGroupResponse) GetBody() *DeleteScalingGroupResponseBody {
	return s.Body
}

func (s *DeleteScalingGroupResponse) SetHeaders(v map[string]*string) *DeleteScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingGroupResponse) SetStatusCode(v int32) *DeleteScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingGroupResponse) SetBody(v *DeleteScalingGroupResponseBody) *DeleteScalingGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteScalingGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
