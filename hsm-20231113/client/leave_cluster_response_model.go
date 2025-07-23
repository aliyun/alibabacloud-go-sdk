// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLeaveClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LeaveClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LeaveClusterResponse
	GetStatusCode() *int32
	SetBody(v *LeaveClusterResponseBody) *LeaveClusterResponse
	GetBody() *LeaveClusterResponseBody
}

type LeaveClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LeaveClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LeaveClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s LeaveClusterResponse) GoString() string {
	return s.String()
}

func (s *LeaveClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LeaveClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LeaveClusterResponse) GetBody() *LeaveClusterResponseBody {
	return s.Body
}

func (s *LeaveClusterResponse) SetHeaders(v map[string]*string) *LeaveClusterResponse {
	s.Headers = v
	return s
}

func (s *LeaveClusterResponse) SetStatusCode(v int32) *LeaveClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *LeaveClusterResponse) SetBody(v *LeaveClusterResponseBody) *LeaveClusterResponse {
	s.Body = v
	return s
}

func (s *LeaveClusterResponse) Validate() error {
	return dara.Validate(s)
}
