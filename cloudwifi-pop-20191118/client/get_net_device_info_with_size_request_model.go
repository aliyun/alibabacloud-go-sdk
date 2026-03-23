// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetDeviceInfoWithSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *GetNetDeviceInfoWithSizeRequest
	GetAppCode() *string
	SetAppName(v string) *GetNetDeviceInfoWithSizeRequest
	GetAppName() *string
	SetCursor(v int64) *GetNetDeviceInfoWithSizeRequest
	GetCursor() *int64
	SetHostName(v string) *GetNetDeviceInfoWithSizeRequest
	GetHostName() *string
	SetId(v int64) *GetNetDeviceInfoWithSizeRequest
	GetId() *int64
	SetIdc(v string) *GetNetDeviceInfoWithSizeRequest
	GetIdc() *string
	SetLogicNetPod(v string) *GetNetDeviceInfoWithSizeRequest
	GetLogicNetPod() *string
	SetManufacturer(v string) *GetNetDeviceInfoWithSizeRequest
	GetManufacturer() *string
	SetMgnIp(v string) *GetNetDeviceInfoWithSizeRequest
	GetMgnIp() *string
	SetModel(v string) *GetNetDeviceInfoWithSizeRequest
	GetModel() *string
	SetNetPod(v string) *GetNetDeviceInfoWithSizeRequest
	GetNetPod() *string
	SetPageSize(v int32) *GetNetDeviceInfoWithSizeRequest
	GetPageSize() *int32
	SetPassword(v string) *GetNetDeviceInfoWithSizeRequest
	GetPassword() *string
	SetRequestId(v string) *GetNetDeviceInfoWithSizeRequest
	GetRequestId() *string
	SetRole(v string) *GetNetDeviceInfoWithSizeRequest
	GetRole() *string
	SetServiceTag(v string) *GetNetDeviceInfoWithSizeRequest
	GetServiceTag() *string
	SetUsername(v string) *GetNetDeviceInfoWithSizeRequest
	GetUsername() *string
}

type GetNetDeviceInfoWithSizeRequest struct {
	AppCode      *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Cursor       *int64  `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	HostName     *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Idc          *string `json:"Idc,omitempty" xml:"Idc,omitempty"`
	LogicNetPod  *string `json:"LogicNetPod,omitempty" xml:"LogicNetPod,omitempty"`
	Manufacturer *string `json:"Manufacturer,omitempty" xml:"Manufacturer,omitempty"`
	MgnIp        *string `json:"MgnIp,omitempty" xml:"MgnIp,omitempty"`
	Model        *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NetPod       *string `json:"NetPod,omitempty" xml:"NetPod,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	ServiceTag *string `json:"ServiceTag,omitempty" xml:"ServiceTag,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetNetDeviceInfoWithSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetDeviceInfoWithSizeRequest) GoString() string {
	return s.String()
}

func (s *GetNetDeviceInfoWithSizeRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetNetDeviceInfoWithSizeRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetNetDeviceInfoWithSizeRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *GetNetDeviceInfoWithSizeRequest) GetHostName() *string {
	return s.HostName
}

func (s *GetNetDeviceInfoWithSizeRequest) GetId() *int64 {
	return s.Id
}

func (s *GetNetDeviceInfoWithSizeRequest) GetIdc() *string {
	return s.Idc
}

func (s *GetNetDeviceInfoWithSizeRequest) GetLogicNetPod() *string {
	return s.LogicNetPod
}

func (s *GetNetDeviceInfoWithSizeRequest) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *GetNetDeviceInfoWithSizeRequest) GetMgnIp() *string {
	return s.MgnIp
}

func (s *GetNetDeviceInfoWithSizeRequest) GetModel() *string {
	return s.Model
}

func (s *GetNetDeviceInfoWithSizeRequest) GetNetPod() *string {
	return s.NetPod
}

func (s *GetNetDeviceInfoWithSizeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetNetDeviceInfoWithSizeRequest) GetPassword() *string {
	return s.Password
}

func (s *GetNetDeviceInfoWithSizeRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetDeviceInfoWithSizeRequest) GetRole() *string {
	return s.Role
}

func (s *GetNetDeviceInfoWithSizeRequest) GetServiceTag() *string {
	return s.ServiceTag
}

func (s *GetNetDeviceInfoWithSizeRequest) GetUsername() *string {
	return s.Username
}

func (s *GetNetDeviceInfoWithSizeRequest) SetAppCode(v string) *GetNetDeviceInfoWithSizeRequest {
	s.AppCode = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetAppName(v string) *GetNetDeviceInfoWithSizeRequest {
	s.AppName = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetCursor(v int64) *GetNetDeviceInfoWithSizeRequest {
	s.Cursor = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetHostName(v string) *GetNetDeviceInfoWithSizeRequest {
	s.HostName = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetId(v int64) *GetNetDeviceInfoWithSizeRequest {
	s.Id = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetIdc(v string) *GetNetDeviceInfoWithSizeRequest {
	s.Idc = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetLogicNetPod(v string) *GetNetDeviceInfoWithSizeRequest {
	s.LogicNetPod = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetManufacturer(v string) *GetNetDeviceInfoWithSizeRequest {
	s.Manufacturer = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetMgnIp(v string) *GetNetDeviceInfoWithSizeRequest {
	s.MgnIp = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetModel(v string) *GetNetDeviceInfoWithSizeRequest {
	s.Model = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetNetPod(v string) *GetNetDeviceInfoWithSizeRequest {
	s.NetPod = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetPageSize(v int32) *GetNetDeviceInfoWithSizeRequest {
	s.PageSize = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetPassword(v string) *GetNetDeviceInfoWithSizeRequest {
	s.Password = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetRequestId(v string) *GetNetDeviceInfoWithSizeRequest {
	s.RequestId = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetRole(v string) *GetNetDeviceInfoWithSizeRequest {
	s.Role = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetServiceTag(v string) *GetNetDeviceInfoWithSizeRequest {
	s.ServiceTag = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) SetUsername(v string) *GetNetDeviceInfoWithSizeRequest {
	s.Username = &v
	return s
}

func (s *GetNetDeviceInfoWithSizeRequest) Validate() error {
	return dara.Validate(s)
}
