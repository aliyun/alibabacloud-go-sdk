// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *ListDevicesRequest
	GetAlias() *string
	SetBuildId(v string) *ListDevicesRequest
	GetBuildId() *string
	SetClientType(v int32) *ListDevicesRequest
	GetClientType() *int32
	SetDeviceGroupId(v string) *ListDevicesRequest
	GetDeviceGroupId() *string
	SetDeviceIpV4(v string) *ListDevicesRequest
	GetDeviceIpV4() *string
	SetDeviceName(v string) *ListDevicesRequest
	GetDeviceName() *string
	SetDeviceOS(v string) *ListDevicesRequest
	GetDeviceOS() *string
	SetDevicePlatform(v string) *ListDevicesRequest
	GetDevicePlatform() *string
	SetEndUserId(v string) *ListDevicesRequest
	GetEndUserId() *string
	SetLabelContent(v string) *ListDevicesRequest
	GetLabelContent() *string
	SetLabelId(v string) *ListDevicesRequest
	GetLabelId() *string
	SetLastLoginUser(v string) *ListDevicesRequest
	GetLastLoginUser() *string
	SetLocationInfo(v string) *ListDevicesRequest
	GetLocationInfo() *string
	SetModel(v string) *ListDevicesRequest
	GetModel() *string
	SetPageNumber(v int32) *ListDevicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDevicesRequest
	GetPageSize() *int32
	SetSerialNo(v string) *ListDevicesRequest
	GetSerialNo() *string
	SetUserType(v string) *ListDevicesRequest
	GetUserType() *string
	SetUuid(v string) *ListDevicesRequest
	GetUuid() *string
}

type ListDevicesRequest struct {
	Alias          *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	BuildId        *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	ClientType     *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	DeviceGroupId  *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	DeviceIpV4     *string `json:"DeviceIpV4,omitempty" xml:"DeviceIpV4,omitempty"`
	DeviceName     *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceOS       *string `json:"DeviceOS,omitempty" xml:"DeviceOS,omitempty"`
	DevicePlatform *string `json:"DevicePlatform,omitempty" xml:"DevicePlatform,omitempty"`
	EndUserId      *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LabelContent   *string `json:"LabelContent,omitempty" xml:"LabelContent,omitempty"`
	LabelId        *string `json:"LabelId,omitempty" xml:"LabelId,omitempty"`
	LastLoginUser  *string `json:"LastLoginUser,omitempty" xml:"LastLoginUser,omitempty"`
	LocationInfo   *string `json:"LocationInfo,omitempty" xml:"LocationInfo,omitempty"`
	Model          *string `json:"Model,omitempty" xml:"Model,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SerialNo       *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	UserType       *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	Uuid           *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) GetAlias() *string {
	return s.Alias
}

func (s *ListDevicesRequest) GetBuildId() *string {
	return s.BuildId
}

func (s *ListDevicesRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *ListDevicesRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *ListDevicesRequest) GetDeviceIpV4() *string {
	return s.DeviceIpV4
}

func (s *ListDevicesRequest) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ListDevicesRequest) GetDeviceOS() *string {
	return s.DeviceOS
}

func (s *ListDevicesRequest) GetDevicePlatform() *string {
	return s.DevicePlatform
}

func (s *ListDevicesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *ListDevicesRequest) GetLabelContent() *string {
	return s.LabelContent
}

func (s *ListDevicesRequest) GetLabelId() *string {
	return s.LabelId
}

func (s *ListDevicesRequest) GetLastLoginUser() *string {
	return s.LastLoginUser
}

func (s *ListDevicesRequest) GetLocationInfo() *string {
	return s.LocationInfo
}

func (s *ListDevicesRequest) GetModel() *string {
	return s.Model
}

func (s *ListDevicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDevicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDevicesRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListDevicesRequest) GetUserType() *string {
	return s.UserType
}

func (s *ListDevicesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ListDevicesRequest) SetAlias(v string) *ListDevicesRequest {
	s.Alias = &v
	return s
}

func (s *ListDevicesRequest) SetBuildId(v string) *ListDevicesRequest {
	s.BuildId = &v
	return s
}

func (s *ListDevicesRequest) SetClientType(v int32) *ListDevicesRequest {
	s.ClientType = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceGroupId(v string) *ListDevicesRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceIpV4(v string) *ListDevicesRequest {
	s.DeviceIpV4 = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceName(v string) *ListDevicesRequest {
	s.DeviceName = &v
	return s
}

func (s *ListDevicesRequest) SetDeviceOS(v string) *ListDevicesRequest {
	s.DeviceOS = &v
	return s
}

func (s *ListDevicesRequest) SetDevicePlatform(v string) *ListDevicesRequest {
	s.DevicePlatform = &v
	return s
}

func (s *ListDevicesRequest) SetEndUserId(v string) *ListDevicesRequest {
	s.EndUserId = &v
	return s
}

func (s *ListDevicesRequest) SetLabelContent(v string) *ListDevicesRequest {
	s.LabelContent = &v
	return s
}

func (s *ListDevicesRequest) SetLabelId(v string) *ListDevicesRequest {
	s.LabelId = &v
	return s
}

func (s *ListDevicesRequest) SetLastLoginUser(v string) *ListDevicesRequest {
	s.LastLoginUser = &v
	return s
}

func (s *ListDevicesRequest) SetLocationInfo(v string) *ListDevicesRequest {
	s.LocationInfo = &v
	return s
}

func (s *ListDevicesRequest) SetModel(v string) *ListDevicesRequest {
	s.Model = &v
	return s
}

func (s *ListDevicesRequest) SetPageNumber(v int32) *ListDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDevicesRequest) SetPageSize(v int32) *ListDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevicesRequest) SetSerialNo(v string) *ListDevicesRequest {
	s.SerialNo = &v
	return s
}

func (s *ListDevicesRequest) SetUserType(v string) *ListDevicesRequest {
	s.UserType = &v
	return s
}

func (s *ListDevicesRequest) SetUuid(v string) *ListDevicesRequest {
	s.Uuid = &v
	return s
}

func (s *ListDevicesRequest) Validate() error {
	return dara.Validate(s)
}
