// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostGroupAccountsFromUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachHostGroupAccountsFromUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachHostGroupAccountsFromUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *DetachHostGroupAccountsFromUserGroupResponseBody) *DetachHostGroupAccountsFromUserGroupResponse
	GetBody() *DetachHostGroupAccountsFromUserGroupResponseBody
}

type DetachHostGroupAccountsFromUserGroupResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostGroupAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) GetBody() *DetachHostGroupAccountsFromUserGroupResponseBody {
	return s.Body
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) SetHeaders(v map[string]*string) *DetachHostGroupAccountsFromUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) SetStatusCode(v int32) *DetachHostGroupAccountsFromUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) SetBody(v *DetachHostGroupAccountsFromUserGroupResponseBody) *DetachHostGroupAccountsFromUserGroupResponse {
	s.Body = v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
