// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableScalingGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableScalingGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableScalingGroupResponse
	GetStatusCode() *int32
	SetBody(v *DisableScalingGroupResponseBody) *DisableScalingGroupResponse
	GetBody() *DisableScalingGroupResponseBody
}

type DisableScalingGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableScalingGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableScalingGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableScalingGroupResponse) GetBody() *DisableScalingGroupResponseBody {
	return s.Body
}

func (s *DisableScalingGroupResponse) SetHeaders(v map[string]*string) *DisableScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *DisableScalingGroupResponse) SetStatusCode(v int32) *DisableScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableScalingGroupResponse) SetBody(v *DisableScalingGroupResponseBody) *DisableScalingGroupResponse {
	s.Body = v
	return s
}

func (s *DisableScalingGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
