// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostGroupAccountsToUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachHostGroupAccountsToUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachHostGroupAccountsToUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *AttachHostGroupAccountsToUserGroupResponseBody) *AttachHostGroupAccountsToUserGroupResponse
	GetBody() *AttachHostGroupAccountsToUserGroupResponseBody
}

type AttachHostGroupAccountsToUserGroupResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostGroupAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachHostGroupAccountsToUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachHostGroupAccountsToUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachHostGroupAccountsToUserGroupResponse) GetBody() *AttachHostGroupAccountsToUserGroupResponseBody {
	return s.Body
}

func (s *AttachHostGroupAccountsToUserGroupResponse) SetHeaders(v map[string]*string) *AttachHostGroupAccountsToUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponse) SetStatusCode(v int32) *AttachHostGroupAccountsToUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponse) SetBody(v *AttachHostGroupAccountsToUserGroupResponseBody) *AttachHostGroupAccountsToUserGroupResponse {
	s.Body = v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponse) Validate() error {
	return dara.Validate(s)
}
