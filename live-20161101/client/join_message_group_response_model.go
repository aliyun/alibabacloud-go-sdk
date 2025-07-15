// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *JoinMessageGroupResponseBody) *JoinMessageGroupResponse
	GetBody() *JoinMessageGroupResponseBody
}

type JoinMessageGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinMessageGroupResponse) GetBody() *JoinMessageGroupResponseBody {
	return s.Body
}

func (s *JoinMessageGroupResponse) SetHeaders(v map[string]*string) *JoinMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *JoinMessageGroupResponse) SetStatusCode(v int32) *JoinMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinMessageGroupResponse) SetBody(v *JoinMessageGroupResponseBody) *JoinMessageGroupResponse {
	s.Body = v
	return s
}

func (s *JoinMessageGroupResponse) Validate() error {
	return dara.Validate(s)
}
