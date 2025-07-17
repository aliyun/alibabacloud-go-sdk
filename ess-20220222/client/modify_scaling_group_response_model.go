// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScalingGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScalingGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScalingGroupResponseBody) *ModifyScalingGroupResponse
	GetBody() *ModifyScalingGroupResponseBody
}

type ModifyScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScalingGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScalingGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScalingGroupResponse) GetBody() *ModifyScalingGroupResponseBody {
	return s.Body
}

func (s *ModifyScalingGroupResponse) SetHeaders(v map[string]*string) *ModifyScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingGroupResponse) SetStatusCode(v int32) *ModifyScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScalingGroupResponse) SetBody(v *ModifyScalingGroupResponseBody) *ModifyScalingGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyScalingGroupResponse) Validate() error {
	return dara.Validate(s)
}
