// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHostGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHostGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHostGroupResponseBody) *ModifyHostGroupResponse
	GetBody() *ModifyHostGroupResponseBody
}

type ModifyHostGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHostGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHostGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHostGroupResponse) GetBody() *ModifyHostGroupResponseBody {
	return s.Body
}

func (s *ModifyHostGroupResponse) SetHeaders(v map[string]*string) *ModifyHostGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostGroupResponse) SetStatusCode(v int32) *ModifyHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostGroupResponse) SetBody(v *ModifyHostGroupResponseBody) *ModifyHostGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyHostGroupResponse) Validate() error {
	return dara.Validate(s)
}
