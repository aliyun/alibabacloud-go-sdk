// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGroupResponseBody) *ModifyGroupResponse
	GetBody() *ModifyGroupResponseBody
}

type ModifyGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGroupResponse) GetBody() *ModifyGroupResponseBody {
	return s.Body
}

func (s *ModifyGroupResponse) SetHeaders(v map[string]*string) *ModifyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyGroupResponse) SetStatusCode(v int32) *ModifyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGroupResponse) SetBody(v *ModifyGroupResponseBody) *ModifyGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyGroupResponse) Validate() error {
	return dara.Validate(s)
}
