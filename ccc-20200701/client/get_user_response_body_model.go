// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserResponseBody
	GetCode() *string
	SetData(v *GetUserResponseBodyData) *GetUserResponseBody
	GetData() *GetUserResponseBodyData
	SetHttpStatusCode(v int32) *GetUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetUserResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
}

type GetUserResponseBody struct {
	// example:
	//
	// OK
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserResponseBody) GetData() *GetUserResponseBodyData {
	return s.Data
}

func (s *GetUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetHttpStatusCode(v int32) *GetUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetParams(v []*string) *GetUserResponseBody {
	s.Params = v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyData struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// example:
	//
	// 8033****
	DeviceExt *string `json:"DeviceExt,omitempty" xml:"DeviceExt,omitempty"`
	// example:
	//
	// ACC-YUNBS-1.0.10-****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// OFFLINE
	DeviceState *string `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	// example:
	//
	// 1001
	DisplayId   *string `json:"DisplayId,omitempty" xml:"DisplayId,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 8003****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// example:
	//
	// 1391234****
	Mobile   *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// Admin@ccc-test
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// Admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *GetUserResponseBodyData) GetDeviceExt() *string {
	return s.DeviceExt
}

func (s *GetUserResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetUserResponseBodyData) GetDeviceState() *string {
	return s.DeviceState
}

func (s *GetUserResponseBodyData) GetDisplayId() *string {
	return s.DisplayId
}

func (s *GetUserResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *GetUserResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserResponseBodyData) GetLoginName() *string {
	return s.LoginName
}

func (s *GetUserResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *GetUserResponseBodyData) GetNickname() *string {
	return s.Nickname
}

func (s *GetUserResponseBodyData) GetRoleId() *string {
	return s.RoleId
}

func (s *GetUserResponseBodyData) GetRoleName() *string {
	return s.RoleName
}

func (s *GetUserResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetUserResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *GetUserResponseBodyData) SetAvatarUrl(v string) *GetUserResponseBodyData {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserResponseBodyData) SetDeviceExt(v string) *GetUserResponseBodyData {
	s.DeviceExt = &v
	return s
}

func (s *GetUserResponseBodyData) SetDeviceId(v string) *GetUserResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetUserResponseBodyData) SetDeviceState(v string) *GetUserResponseBodyData {
	s.DeviceState = &v
	return s
}

func (s *GetUserResponseBodyData) SetDisplayId(v string) *GetUserResponseBodyData {
	s.DisplayId = &v
	return s
}

func (s *GetUserResponseBodyData) SetDisplayName(v string) *GetUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyData) SetEmail(v string) *GetUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyData) SetExtension(v string) *GetUserResponseBodyData {
	s.Extension = &v
	return s
}

func (s *GetUserResponseBodyData) SetInstanceId(v string) *GetUserResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUserResponseBodyData) SetLoginName(v string) *GetUserResponseBodyData {
	s.LoginName = &v
	return s
}

func (s *GetUserResponseBodyData) SetMobile(v string) *GetUserResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBodyData) SetNickname(v string) *GetUserResponseBodyData {
	s.Nickname = &v
	return s
}

func (s *GetUserResponseBodyData) SetRoleId(v string) *GetUserResponseBodyData {
	s.RoleId = &v
	return s
}

func (s *GetUserResponseBodyData) SetRoleName(v string) *GetUserResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *GetUserResponseBodyData) SetUserId(v string) *GetUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyData) SetWorkMode(v string) *GetUserResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *GetUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
