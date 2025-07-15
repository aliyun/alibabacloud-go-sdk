// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDesktopOversoldUserGroupResponseBodyData) *ModifyDesktopOversoldUserGroupResponseBody
	GetData() *ModifyDesktopOversoldUserGroupResponseBodyData
	SetRequestId(v string) *ModifyDesktopOversoldUserGroupResponseBody
	GetRequestId() *string
}

type ModifyDesktopOversoldUserGroupResponseBody struct {
	Data      *ModifyDesktopOversoldUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDesktopOversoldUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldUserGroupResponseBody) GetData() *ModifyDesktopOversoldUserGroupResponseBodyData {
	return s.Data
}

func (s *ModifyDesktopOversoldUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDesktopOversoldUserGroupResponseBody) SetData(v *ModifyDesktopOversoldUserGroupResponseBodyData) *ModifyDesktopOversoldUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDesktopOversoldUserGroupResponseBody) SetRequestId(v string) *ModifyDesktopOversoldUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDesktopOversoldUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyDesktopOversoldUserGroupResponseBodyData struct {
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ModifyDesktopOversoldUserGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldUserGroupResponseBodyData) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ModifyDesktopOversoldUserGroupResponseBodyData) SetUserGroupId(v string) *ModifyDesktopOversoldUserGroupResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldUserGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
