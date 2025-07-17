// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScalingGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScalingGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateScalingGroupResponseBody) *CreateScalingGroupResponse
	GetBody() *CreateScalingGroupResponseBody
}

type CreateScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScalingGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScalingGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScalingGroupResponse) GetBody() *CreateScalingGroupResponseBody {
	return s.Body
}

func (s *CreateScalingGroupResponse) SetHeaders(v map[string]*string) *CreateScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingGroupResponse) SetStatusCode(v int32) *CreateScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingGroupResponse) SetBody(v *CreateScalingGroupResponseBody) *CreateScalingGroupResponse {
	s.Body = v
	return s
}

func (s *CreateScalingGroupResponse) Validate() error {
	return dara.Validate(s)
}
