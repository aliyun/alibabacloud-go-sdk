// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDesktopOversoldUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddDesktopOversoldUserGroupResponseBodyData) *AddDesktopOversoldUserGroupResponseBody
	GetData() *AddDesktopOversoldUserGroupResponseBodyData
	SetRequestId(v string) *AddDesktopOversoldUserGroupResponseBody
	GetRequestId() *string
}

type AddDesktopOversoldUserGroupResponseBody struct {
	Data      *AddDesktopOversoldUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDesktopOversoldUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDesktopOversoldUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddDesktopOversoldUserGroupResponseBody) GetData() *AddDesktopOversoldUserGroupResponseBodyData {
	return s.Data
}

func (s *AddDesktopOversoldUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDesktopOversoldUserGroupResponseBody) SetData(v *AddDesktopOversoldUserGroupResponseBodyData) *AddDesktopOversoldUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *AddDesktopOversoldUserGroupResponseBody) SetRequestId(v string) *AddDesktopOversoldUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDesktopOversoldUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddDesktopOversoldUserGroupResponseBodyData struct {
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AddDesktopOversoldUserGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddDesktopOversoldUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDesktopOversoldUserGroupResponseBodyData) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddDesktopOversoldUserGroupResponseBodyData) SetUserGroupId(v string) *AddDesktopOversoldUserGroupResponseBodyData {
	s.UserGroupId = &v
	return s
}

func (s *AddDesktopOversoldUserGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
