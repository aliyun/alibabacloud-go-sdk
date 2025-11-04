// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScalingGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyScalingGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyScalingGroupResponse
	GetStatusCode() *int32
	SetBody(v *ApplyScalingGroupResponseBody) *ApplyScalingGroupResponse
	GetBody() *ApplyScalingGroupResponseBody
}

type ApplyScalingGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyScalingGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *ApplyScalingGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyScalingGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyScalingGroupResponse) GetBody() *ApplyScalingGroupResponseBody {
	return s.Body
}

func (s *ApplyScalingGroupResponse) SetHeaders(v map[string]*string) *ApplyScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *ApplyScalingGroupResponse) SetStatusCode(v int32) *ApplyScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyScalingGroupResponse) SetBody(v *ApplyScalingGroupResponseBody) *ApplyScalingGroupResponse {
	s.Body = v
	return s
}

func (s *ApplyScalingGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
