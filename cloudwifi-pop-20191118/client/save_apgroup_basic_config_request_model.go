// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupBasicConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *SaveApgroupBasicConfigRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApgroupBasicConfigRequest
	GetAppName() *string
	SetCountry(v string) *SaveApgroupBasicConfigRequest
	GetCountry() *string
	SetDai(v string) *SaveApgroupBasicConfigRequest
	GetDai() *string
	SetDescription(v string) *SaveApgroupBasicConfigRequest
	GetDescription() *string
	SetEchoInt(v int32) *SaveApgroupBasicConfigRequest
	GetEchoInt() *int32
	SetFailover(v int32) *SaveApgroupBasicConfigRequest
	GetFailover() *int32
	SetId(v int64) *SaveApgroupBasicConfigRequest
	GetId() *int64
	SetInsecureAllowed(v int32) *SaveApgroupBasicConfigRequest
	GetInsecureAllowed() *int32
	SetIsConfigStrongConsistency(v bool) *SaveApgroupBasicConfigRequest
	GetIsConfigStrongConsistency() *bool
	SetLanIp(v string) *SaveApgroupBasicConfigRequest
	GetLanIp() *string
	SetLanMask(v string) *SaveApgroupBasicConfigRequest
	GetLanMask() *string
	SetLanStatus(v int32) *SaveApgroupBasicConfigRequest
	GetLanStatus() *int32
	SetLogIp(v string) *SaveApgroupBasicConfigRequest
	GetLogIp() *string
	SetLogLevel(v int32) *SaveApgroupBasicConfigRequest
	GetLogLevel() *int32
	SetName(v string) *SaveApgroupBasicConfigRequest
	GetName() *string
	SetParentApgroupId(v int64) *SaveApgroupBasicConfigRequest
	GetParentApgroupId() *int64
	SetPasswd(v string) *SaveApgroupBasicConfigRequest
	GetPasswd() *string
	SetProtocol(v string) *SaveApgroupBasicConfigRequest
	GetProtocol() *string
	SetRoute(v string) *SaveApgroupBasicConfigRequest
	GetRoute() *string
	SetScan(v int32) *SaveApgroupBasicConfigRequest
	GetScan() *int32
	SetTokenServer(v string) *SaveApgroupBasicConfigRequest
	GetTokenServer() *string
	SetVpn(v string) *SaveApgroupBasicConfigRequest
	GetVpn() *string
	SetWorkMode(v int32) *SaveApgroupBasicConfigRequest
	GetWorkMode() *int32
}

type SaveApgroupBasicConfigRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Dai         *string `json:"Dai,omitempty" xml:"Dai,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	EchoInt  *int32 `json:"EchoInt,omitempty" xml:"EchoInt,omitempty"`
	Failover *int32 `json:"Failover,omitempty" xml:"Failover,omitempty"`
	// This parameter is required.
	Id              *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	InsecureAllowed *int32 `json:"InsecureAllowed,omitempty" xml:"InsecureAllowed,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsConfigStrongConsistency *bool   `json:"IsConfigStrongConsistency,omitempty" xml:"IsConfigStrongConsistency,omitempty"`
	LanIp                     *string `json:"LanIp,omitempty" xml:"LanIp,omitempty"`
	LanMask                   *string `json:"LanMask,omitempty" xml:"LanMask,omitempty"`
	LanStatus                 *int32  `json:"LanStatus,omitempty" xml:"LanStatus,omitempty"`
	LogIp                     *string `json:"LogIp,omitempty" xml:"LogIp,omitempty"`
	LogLevel                  *int32  `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// This parameter is required.
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentApgroupId *int64  `json:"ParentApgroupId,omitempty" xml:"ParentApgroupId,omitempty"`
	Passwd          *string `json:"Passwd,omitempty" xml:"Passwd,omitempty"`
	Protocol        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Route           *string `json:"Route,omitempty" xml:"Route,omitempty"`
	Scan            *int32  `json:"Scan,omitempty" xml:"Scan,omitempty"`
	TokenServer     *string `json:"TokenServer,omitempty" xml:"TokenServer,omitempty"`
	Vpn             *string `json:"Vpn,omitempty" xml:"Vpn,omitempty"`
	WorkMode        *int32  `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s SaveApgroupBasicConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupBasicConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveApgroupBasicConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApgroupBasicConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApgroupBasicConfigRequest) GetCountry() *string {
	return s.Country
}

func (s *SaveApgroupBasicConfigRequest) GetDai() *string {
	return s.Dai
}

func (s *SaveApgroupBasicConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *SaveApgroupBasicConfigRequest) GetEchoInt() *int32 {
	return s.EchoInt
}

func (s *SaveApgroupBasicConfigRequest) GetFailover() *int32 {
	return s.Failover
}

func (s *SaveApgroupBasicConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *SaveApgroupBasicConfigRequest) GetInsecureAllowed() *int32 {
	return s.InsecureAllowed
}

func (s *SaveApgroupBasicConfigRequest) GetIsConfigStrongConsistency() *bool {
	return s.IsConfigStrongConsistency
}

func (s *SaveApgroupBasicConfigRequest) GetLanIp() *string {
	return s.LanIp
}

func (s *SaveApgroupBasicConfigRequest) GetLanMask() *string {
	return s.LanMask
}

func (s *SaveApgroupBasicConfigRequest) GetLanStatus() *int32 {
	return s.LanStatus
}

func (s *SaveApgroupBasicConfigRequest) GetLogIp() *string {
	return s.LogIp
}

func (s *SaveApgroupBasicConfigRequest) GetLogLevel() *int32 {
	return s.LogLevel
}

func (s *SaveApgroupBasicConfigRequest) GetName() *string {
	return s.Name
}

func (s *SaveApgroupBasicConfigRequest) GetParentApgroupId() *int64 {
	return s.ParentApgroupId
}

func (s *SaveApgroupBasicConfigRequest) GetPasswd() *string {
	return s.Passwd
}

func (s *SaveApgroupBasicConfigRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *SaveApgroupBasicConfigRequest) GetRoute() *string {
	return s.Route
}

func (s *SaveApgroupBasicConfigRequest) GetScan() *int32 {
	return s.Scan
}

func (s *SaveApgroupBasicConfigRequest) GetTokenServer() *string {
	return s.TokenServer
}

func (s *SaveApgroupBasicConfigRequest) GetVpn() *string {
	return s.Vpn
}

func (s *SaveApgroupBasicConfigRequest) GetWorkMode() *int32 {
	return s.WorkMode
}

func (s *SaveApgroupBasicConfigRequest) SetAppCode(v string) *SaveApgroupBasicConfigRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetAppName(v string) *SaveApgroupBasicConfigRequest {
	s.AppName = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetCountry(v string) *SaveApgroupBasicConfigRequest {
	s.Country = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetDai(v string) *SaveApgroupBasicConfigRequest {
	s.Dai = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetDescription(v string) *SaveApgroupBasicConfigRequest {
	s.Description = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetEchoInt(v int32) *SaveApgroupBasicConfigRequest {
	s.EchoInt = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetFailover(v int32) *SaveApgroupBasicConfigRequest {
	s.Failover = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetId(v int64) *SaveApgroupBasicConfigRequest {
	s.Id = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetInsecureAllowed(v int32) *SaveApgroupBasicConfigRequest {
	s.InsecureAllowed = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetIsConfigStrongConsistency(v bool) *SaveApgroupBasicConfigRequest {
	s.IsConfigStrongConsistency = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetLanIp(v string) *SaveApgroupBasicConfigRequest {
	s.LanIp = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetLanMask(v string) *SaveApgroupBasicConfigRequest {
	s.LanMask = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetLanStatus(v int32) *SaveApgroupBasicConfigRequest {
	s.LanStatus = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetLogIp(v string) *SaveApgroupBasicConfigRequest {
	s.LogIp = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetLogLevel(v int32) *SaveApgroupBasicConfigRequest {
	s.LogLevel = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetName(v string) *SaveApgroupBasicConfigRequest {
	s.Name = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetParentApgroupId(v int64) *SaveApgroupBasicConfigRequest {
	s.ParentApgroupId = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetPasswd(v string) *SaveApgroupBasicConfigRequest {
	s.Passwd = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetProtocol(v string) *SaveApgroupBasicConfigRequest {
	s.Protocol = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetRoute(v string) *SaveApgroupBasicConfigRequest {
	s.Route = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetScan(v int32) *SaveApgroupBasicConfigRequest {
	s.Scan = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetTokenServer(v string) *SaveApgroupBasicConfigRequest {
	s.TokenServer = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetVpn(v string) *SaveApgroupBasicConfigRequest {
	s.Vpn = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) SetWorkMode(v int32) *SaveApgroupBasicConfigRequest {
	s.WorkMode = &v
	return s
}

func (s *SaveApgroupBasicConfigRequest) Validate() error {
	return dara.Validate(s)
}
