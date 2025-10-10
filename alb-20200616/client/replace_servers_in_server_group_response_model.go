// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceServersInServerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceServersInServerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceServersInServerGroupResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceServersInServerGroupResponseBody) *ReplaceServersInServerGroupResponse
	GetBody() *ReplaceServersInServerGroupResponseBody
}

type ReplaceServersInServerGroupResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceServersInServerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceServersInServerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceServersInServerGroupResponse) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceServersInServerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceServersInServerGroupResponse) GetBody() *ReplaceServersInServerGroupResponseBody {
	return s.Body
}

func (s *ReplaceServersInServerGroupResponse) SetHeaders(v map[string]*string) *ReplaceServersInServerGroupResponse {
	s.Headers = v
	return s
}

func (s *ReplaceServersInServerGroupResponse) SetStatusCode(v int32) *ReplaceServersInServerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceServersInServerGroupResponse) SetBody(v *ReplaceServersInServerGroupResponseBody) *ReplaceServersInServerGroupResponse {
	s.Body = v
	return s
}

func (s *ReplaceServersInServerGroupResponse) Validate() error {
	return dara.Validate(s)
}
