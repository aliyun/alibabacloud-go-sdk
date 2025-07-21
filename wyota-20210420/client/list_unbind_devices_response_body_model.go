// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnbindDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUnbindDevicesResponseBody
	GetCode() *string
	SetData(v *ListUnbindDevicesResponseBodyData) *ListUnbindDevicesResponseBody
	GetData() *ListUnbindDevicesResponseBodyData
	SetHttpStatusCode(v int32) *ListUnbindDevicesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUnbindDevicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUnbindDevicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUnbindDevicesResponseBody
	GetSuccess() *bool
}

type ListUnbindDevicesResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListUnbindDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUnbindDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUnbindDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnbindDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUnbindDevicesResponseBody) GetData() *ListUnbindDevicesResponseBodyData {
	return s.Data
}

func (s *ListUnbindDevicesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUnbindDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUnbindDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUnbindDevicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUnbindDevicesResponseBody) SetCode(v string) *ListUnbindDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUnbindDevicesResponseBody) SetData(v *ListUnbindDevicesResponseBodyData) *ListUnbindDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListUnbindDevicesResponseBody) SetHttpStatusCode(v int32) *ListUnbindDevicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUnbindDevicesResponseBody) SetMessage(v string) *ListUnbindDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUnbindDevicesResponseBody) SetRequestId(v string) *ListUnbindDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnbindDevicesResponseBody) SetSuccess(v bool) *ListUnbindDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUnbindDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUnbindDevicesResponseBodyData struct {
	Devices    []*ListUnbindDevicesResponseBodyDataDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUnbindDevicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUnbindDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUnbindDevicesResponseBodyData) GetDevices() []*ListUnbindDevicesResponseBodyDataDevices {
	return s.Devices
}

func (s *ListUnbindDevicesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUnbindDevicesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnbindDevicesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUnbindDevicesResponseBodyData) SetDevices(v []*ListUnbindDevicesResponseBodyDataDevices) *ListUnbindDevicesResponseBodyData {
	s.Devices = v
	return s
}

func (s *ListUnbindDevicesResponseBodyData) SetPageNumber(v int32) *ListUnbindDevicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyData) SetPageSize(v int32) *ListUnbindDevicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyData) SetTotalCount(v int32) *ListUnbindDevicesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUnbindDevicesResponseBodyDataDevices struct {
	Alias                         *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BoundTime                     *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	BuildId                       *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientType                    *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ConnectionStatus              *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	DeviceMqttConnectionStatus    *int32  `json:"DeviceMqttConnectionStatus,omitempty" xml:"DeviceMqttConnectionStatus,omitempty"`
	DeviceOs                      *string `json:"DeviceOs,omitempty" xml:"DeviceOs,omitempty"`
	DevicePlatform                *string `json:"DevicePlatform,omitempty" xml:"DevicePlatform,omitempty"`
	InManage                      *bool   `json:"InManage,omitempty" xml:"InManage,omitempty"`
	LastLoginTime                 *string `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	LastLoginUser                 *string `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	LoginUser                     *string `json:"LoginUser,omitempty" xml:"LoginUser,omitempty"`
	Model                         *string `json:"Model,omitempty" xml:"Model,omitempty"`
	PasswordFreeLoginUser         *string `json:"PasswordFreeLoginUser,omitempty" xml:"PasswordFreeLoginUser,omitempty"`
	PasswordFreeLoginUserNickName *string `json:"PasswordFreeLoginUserNickName,omitempty" xml:"PasswordFreeLoginUserNickName,omitempty"`
	PrivateIp                     *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	ProductName                   *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	PublicIp                      *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	SerialNo                      *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Uuid                          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListUnbindDevicesResponseBodyDataDevices) String() string {
	return dara.Prettify(s)
}

func (s ListUnbindDevicesResponseBodyDataDevices) GoString() string {
	return s.String()
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetAlias() *string {
	return s.Alias
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetBoundTime() *string {
	return s.BoundTime
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetBuildId() *string {
	return s.BuildId
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetClientType() *string {
	return s.ClientType
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetDeviceMqttConnectionStatus() *int32 {
	return s.DeviceMqttConnectionStatus
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetDeviceOs() *string {
	return s.DeviceOs
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetDevicePlatform() *string {
	return s.DevicePlatform
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetInManage() *bool {
	return s.InManage
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetLoginUser() *string {
	return s.LoginUser
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetModel() *string {
	return s.Model
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetPasswordFreeLoginUser() *string {
	return s.PasswordFreeLoginUser
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetPasswordFreeLoginUserNickName() *string {
	return s.PasswordFreeLoginUserNickName
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetProductName() *string {
	return s.ProductName
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListUnbindDevicesResponseBodyDataDevices) GetUuid() *string {
	return s.Uuid
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetAlias(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.Alias = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetBoundTime(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.BoundTime = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetBuildId(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.BuildId = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetClientType(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.ClientType = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetConnectionStatus(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.ConnectionStatus = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetDeviceMqttConnectionStatus(v int32) *ListUnbindDevicesResponseBodyDataDevices {
	s.DeviceMqttConnectionStatus = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetDeviceOs(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.DeviceOs = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetDevicePlatform(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.DevicePlatform = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetInManage(v bool) *ListUnbindDevicesResponseBodyDataDevices {
	s.InManage = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetLastLoginTime(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.LastLoginTime = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetLastLoginUser(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.LastLoginUser = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetLoginUser(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.LoginUser = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetModel(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.Model = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetPasswordFreeLoginUser(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.PasswordFreeLoginUser = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetPasswordFreeLoginUserNickName(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.PasswordFreeLoginUserNickName = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetPrivateIp(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.PrivateIp = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetProductName(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.ProductName = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetPublicIp(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.PublicIp = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetSerialNo(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.SerialNo = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) SetUuid(v string) *ListUnbindDevicesResponseBodyDataDevices {
	s.Uuid = &v
	return s
}

func (s *ListUnbindDevicesResponseBodyDataDevices) Validate() error {
	return dara.Validate(s)
}
