// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserGroupResponseBody
	GetRequestId() *string
}

type ModifyUserGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserGroupResponseBody) SetRequestId(v string) *ModifyUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
