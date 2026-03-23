// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *GetNetDeviceInfoRequest
	GetAppCode() *string
	SetAppName(v string) *GetNetDeviceInfoRequest
	GetAppName() *string
	SetCursor(v int64) *GetNetDeviceInfoRequest
	GetCursor() *int64
	SetHostName(v string) *GetNetDeviceInfoRequest
	GetHostName() *string
	SetId(v int64) *GetNetDeviceInfoRequest
	GetId() *int64
	SetIdc(v string) *GetNetDeviceInfoRequest
	GetIdc() *string
	SetLogicNetPod(v string) *GetNetDeviceInfoRequest
	GetLogicNetPod() *string
	SetManufacturer(v string) *GetNetDeviceInfoRequest
	GetManufacturer() *string
	SetMgnIp(v string) *GetNetDeviceInfoRequest
	GetMgnIp() *string
	SetModel(v string) *GetNetDeviceInfoRequest
	GetModel() *string
	SetNetPod(v string) *GetNetDeviceInfoRequest
	GetNetPod() *string
	SetPageSize(v int32) *GetNetDeviceInfoRequest
	GetPageSize() *int32
	SetPassword(v string) *GetNetDeviceInfoRequest
	GetPassword() *string
	SetRequestId(v string) *GetNetDeviceInfoRequest
	GetRequestId() *string
	SetRole(v string) *GetNetDeviceInfoRequest
	GetRole() *string
	SetServiceTag(v string) *GetNetDeviceInfoRequest
	GetServiceTag() *string
	SetUsername(v string) *GetNetDeviceInfoRequest
	GetUsername() *string
}

type GetNetDeviceInfoRequest struct {
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	ServiceTag   *string `json:"ServiceTag,omitempty" xml:"ServiceTag,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetNetDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetNetDeviceInfoRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *GetNetDeviceInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetNetDeviceInfoRequest) GetCursor() *int64 {
	return s.Cursor
}

func (s *GetNetDeviceInfoRequest) GetHostName() *string {
	return s.HostName
}

func (s *GetNetDeviceInfoRequest) GetId() *int64 {
	return s.Id
}

func (s *GetNetDeviceInfoRequest) GetIdc() *string {
	return s.Idc
}

func (s *GetNetDeviceInfoRequest) GetLogicNetPod() *string {
	return s.LogicNetPod
}

func (s *GetNetDeviceInfoRequest) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *GetNetDeviceInfoRequest) GetMgnIp() *string {
	return s.MgnIp
}

func (s *GetNetDeviceInfoRequest) GetModel() *string {
	return s.Model
}

func (s *GetNetDeviceInfoRequest) GetNetPod() *string {
	return s.NetPod
}

func (s *GetNetDeviceInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetNetDeviceInfoRequest) GetPassword() *string {
	return s.Password
}

func (s *GetNetDeviceInfoRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetDeviceInfoRequest) GetRole() *string {
	return s.Role
}

func (s *GetNetDeviceInfoRequest) GetServiceTag() *string {
	return s.ServiceTag
}

func (s *GetNetDeviceInfoRequest) GetUsername() *string {
	return s.Username
}

func (s *GetNetDeviceInfoRequest) SetAppCode(v string) *GetNetDeviceInfoRequest {
	s.AppCode = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetAppName(v string) *GetNetDeviceInfoRequest {
	s.AppName = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetCursor(v int64) *GetNetDeviceInfoRequest {
	s.Cursor = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetHostName(v string) *GetNetDeviceInfoRequest {
	s.HostName = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetId(v int64) *GetNetDeviceInfoRequest {
	s.Id = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetIdc(v string) *GetNetDeviceInfoRequest {
	s.Idc = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetLogicNetPod(v string) *GetNetDeviceInfoRequest {
	s.LogicNetPod = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetManufacturer(v string) *GetNetDeviceInfoRequest {
	s.Manufacturer = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetMgnIp(v string) *GetNetDeviceInfoRequest {
	s.MgnIp = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetModel(v string) *GetNetDeviceInfoRequest {
	s.Model = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetNetPod(v string) *GetNetDeviceInfoRequest {
	s.NetPod = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetPageSize(v int32) *GetNetDeviceInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetPassword(v string) *GetNetDeviceInfoRequest {
	s.Password = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetRequestId(v string) *GetNetDeviceInfoRequest {
	s.RequestId = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetRole(v string) *GetNetDeviceInfoRequest {
	s.Role = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetServiceTag(v string) *GetNetDeviceInfoRequest {
	s.ServiceTag = &v
	return s
}

func (s *GetNetDeviceInfoRequest) SetUsername(v string) *GetNetDeviceInfoRequest {
	s.Username = &v
	return s
}

func (s *GetNetDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
