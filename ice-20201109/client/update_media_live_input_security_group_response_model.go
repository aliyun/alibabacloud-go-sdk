// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveInputSecurityGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaLiveInputSecurityGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaLiveInputSecurityGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaLiveInputSecurityGroupResponseBody) *UpdateMediaLiveInputSecurityGroupResponse
	GetBody() *UpdateMediaLiveInputSecurityGroupResponseBody
}

type UpdateMediaLiveInputSecurityGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaLiveInputSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaLiveInputSecurityGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputSecurityGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaLiveInputSecurityGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaLiveInputSecurityGroupResponse) GetBody() *UpdateMediaLiveInputSecurityGroupResponseBody {
	return s.Body
}

func (s *UpdateMediaLiveInputSecurityGroupResponse) SetHeaders(v map[string]*string) *UpdateMediaLiveInputSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupResponse) SetStatusCode(v int32) *UpdateMediaLiveInputSecurityGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupResponse) SetBody(v *UpdateMediaLiveInputSecurityGroupResponseBody) *UpdateMediaLiveInputSecurityGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaLiveInputSecurityGroupResponse) Validate() error {
	return dara.Validate(s)
}
