// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LeaveMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LeaveMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *LeaveMessageGroupResponseBody) *LeaveMessageGroupResponse
	GetBody() *LeaveMessageGroupResponseBody
}

type LeaveMessageGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LeaveMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LeaveMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s LeaveMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *LeaveMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LeaveMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LeaveMessageGroupResponse) GetBody() *LeaveMessageGroupResponseBody {
	return s.Body
}

func (s *LeaveMessageGroupResponse) SetHeaders(v map[string]*string) *LeaveMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *LeaveMessageGroupResponse) SetStatusCode(v int32) *LeaveMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *LeaveMessageGroupResponse) SetBody(v *LeaveMessageGroupResponseBody) *LeaveMessageGroupResponse {
	s.Body = v
	return s
}

func (s *LeaveMessageGroupResponse) Validate() error {
	return dara.Validate(s)
}
