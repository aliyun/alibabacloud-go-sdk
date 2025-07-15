// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApiGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApiGroupResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApiGroupResponseBody) *ModifyApiGroupResponse
	GetBody() *ModifyApiGroupResponseBody
}

type ModifyApiGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApiGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApiGroupResponse) GetBody() *ModifyApiGroupResponseBody {
	return s.Body
}

func (s *ModifyApiGroupResponse) SetHeaders(v map[string]*string) *ModifyApiGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiGroupResponse) SetStatusCode(v int32) *ModifyApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiGroupResponse) SetBody(v *ModifyApiGroupResponseBody) *ModifyApiGroupResponse {
	s.Body = v
	return s
}

func (s *ModifyApiGroupResponse) Validate() error {
	return dara.Validate(s)
}
