// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBoundDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBoundDevicesResponseBody
	GetCode() *string
	SetData(v *ListBoundDevicesResponseBodyData) *ListBoundDevicesResponseBody
	GetData() *ListBoundDevicesResponseBodyData
	SetHttpStatusCode(v int32) *ListBoundDevicesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBoundDevicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBoundDevicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBoundDevicesResponseBody
	GetSuccess() *bool
}

type ListBoundDevicesResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListBoundDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBoundDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBoundDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBoundDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBoundDevicesResponseBody) GetData() *ListBoundDevicesResponseBodyData {
	return s.Data
}

func (s *ListBoundDevicesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBoundDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBoundDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBoundDevicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBoundDevicesResponseBody) SetCode(v string) *ListBoundDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListBoundDevicesResponseBody) SetData(v *ListBoundDevicesResponseBodyData) *ListBoundDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListBoundDevicesResponseBody) SetHttpStatusCode(v int32) *ListBoundDevicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBoundDevicesResponseBody) SetMessage(v string) *ListBoundDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListBoundDevicesResponseBody) SetRequestId(v string) *ListBoundDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBoundDevicesResponseBody) SetSuccess(v bool) *ListBoundDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListBoundDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBoundDevicesResponseBodyData struct {
	Devices    []*ListBoundDevicesResponseBodyDataDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBoundDevicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBoundDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBoundDevicesResponseBodyData) GetDevices() []*ListBoundDevicesResponseBodyDataDevices {
	return s.Devices
}

func (s *ListBoundDevicesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBoundDevicesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBoundDevicesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBoundDevicesResponseBodyData) SetDevices(v []*ListBoundDevicesResponseBodyDataDevices) *ListBoundDevicesResponseBodyData {
	s.Devices = v
	return s
}

func (s *ListBoundDevicesResponseBodyData) SetPageNumber(v int32) *ListBoundDevicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBoundDevicesResponseBodyData) SetPageSize(v int32) *ListBoundDevicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBoundDevicesResponseBodyData) SetTotalCount(v int32) *ListBoundDevicesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListBoundDevicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListBoundDevicesResponseBodyDataDevices struct {
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

func (s ListBoundDevicesResponseBodyDataDevices) String() string {
	return dara.Prettify(s)
}

func (s ListBoundDevicesResponseBodyDataDevices) GoString() string {
	return s.String()
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetAlias() *string {
	return s.Alias
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetBoundTime() *string {
	return s.BoundTime
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetBuildId() *string {
	return s.BuildId
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetClientType() *string {
	return s.ClientType
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetDeviceMqttConnectionStatus() *int32 {
	return s.DeviceMqttConnectionStatus
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetDeviceOs() *string {
	return s.DeviceOs
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetDevicePlatform() *string {
	return s.DevicePlatform
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetInManage() *bool {
	return s.InManage
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetLastLoginTime() *string {
	return s.LastLoginTime
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetLoginUser() *string {
	return s.LoginUser
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetModel() *string {
	return s.Model
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetPasswordFreeLoginUser() *string {
	return s.PasswordFreeLoginUser
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetPasswordFreeLoginUserNickName() *string {
	return s.PasswordFreeLoginUserNickName
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetProductName() *string {
	return s.ProductName
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListBoundDevicesResponseBodyDataDevices) GetUuid() *string {
	return s.Uuid
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetAlias(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.Alias = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetBoundTime(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.BoundTime = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetBuildId(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.BuildId = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetClientType(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.ClientType = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetConnectionStatus(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.ConnectionStatus = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetDeviceMqttConnectionStatus(v int32) *ListBoundDevicesResponseBodyDataDevices {
	s.DeviceMqttConnectionStatus = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetDeviceOs(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.DeviceOs = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetDevicePlatform(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.DevicePlatform = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetInManage(v bool) *ListBoundDevicesResponseBodyDataDevices {
	s.InManage = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetLastLoginTime(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.LastLoginTime = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetLastLoginUser(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.LastLoginUser = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetLoginUser(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.LoginUser = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetModel(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.Model = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetPasswordFreeLoginUser(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.PasswordFreeLoginUser = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetPasswordFreeLoginUserNickName(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.PasswordFreeLoginUserNickName = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetPrivateIp(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.PrivateIp = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetProductName(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.ProductName = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetPublicIp(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.PublicIp = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetSerialNo(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.SerialNo = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) SetUuid(v string) *ListBoundDevicesResponseBodyDataDevices {
	s.Uuid = &v
	return s
}

func (s *ListBoundDevicesResponseBodyDataDevices) Validate() error {
	return dara.Validate(s)
}
