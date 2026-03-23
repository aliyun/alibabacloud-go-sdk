// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApBasicConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *SaveApBasicConfigRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApBasicConfigRequest
	GetAppName() *string
	SetCountry(v string) *SaveApBasicConfigRequest
	GetCountry() *string
	SetDai(v string) *SaveApBasicConfigRequest
	GetDai() *string
	SetDescription(v string) *SaveApBasicConfigRequest
	GetDescription() *string
	SetEchoInt(v int32) *SaveApBasicConfigRequest
	GetEchoInt() *int32
	SetFailover(v int32) *SaveApBasicConfigRequest
	GetFailover() *int32
	SetId(v int64) *SaveApBasicConfigRequest
	GetId() *int64
	SetInsecureAllowed(v int32) *SaveApBasicConfigRequest
	GetInsecureAllowed() *int32
	SetLanIp(v string) *SaveApBasicConfigRequest
	GetLanIp() *string
	SetLanMask(v string) *SaveApBasicConfigRequest
	GetLanMask() *string
	SetLanStatus(v int32) *SaveApBasicConfigRequest
	GetLanStatus() *int32
	SetLogIp(v string) *SaveApBasicConfigRequest
	GetLogIp() *string
	SetLogLevel(v int32) *SaveApBasicConfigRequest
	GetLogLevel() *int32
	SetName(v string) *SaveApBasicConfigRequest
	GetName() *string
	SetPasswd(v string) *SaveApBasicConfigRequest
	GetPasswd() *string
	SetProtocol(v string) *SaveApBasicConfigRequest
	GetProtocol() *string
	SetRoute(v string) *SaveApBasicConfigRequest
	GetRoute() *string
	SetScan(v int32) *SaveApBasicConfigRequest
	GetScan() *int32
	SetTokenServer(v string) *SaveApBasicConfigRequest
	GetTokenServer() *string
	SetVpn(v string) *SaveApBasicConfigRequest
	GetVpn() *string
	SetWorkMode(v int32) *SaveApBasicConfigRequest
	GetWorkMode() *int32
}

type SaveApBasicConfigRequest struct {
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
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InsecureAllowed *int32  `json:"InsecureAllowed,omitempty" xml:"InsecureAllowed,omitempty"`
	LanIp           *string `json:"LanIp,omitempty" xml:"LanIp,omitempty"`
	LanMask         *string `json:"LanMask,omitempty" xml:"LanMask,omitempty"`
	LanStatus       *int32  `json:"LanStatus,omitempty" xml:"LanStatus,omitempty"`
	LogIp           *string `json:"LogIp,omitempty" xml:"LogIp,omitempty"`
	LogLevel        *int32  `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	// This parameter is required.
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Passwd      *string `json:"Passwd,omitempty" xml:"Passwd,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Route       *string `json:"Route,omitempty" xml:"Route,omitempty"`
	Scan        *int32  `json:"Scan,omitempty" xml:"Scan,omitempty"`
	TokenServer *string `json:"TokenServer,omitempty" xml:"TokenServer,omitempty"`
	Vpn         *string `json:"Vpn,omitempty" xml:"Vpn,omitempty"`
	WorkMode    *int32  `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s SaveApBasicConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApBasicConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveApBasicConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApBasicConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApBasicConfigRequest) GetCountry() *string {
	return s.Country
}

func (s *SaveApBasicConfigRequest) GetDai() *string {
	return s.Dai
}

func (s *SaveApBasicConfigRequest) GetDescription() *string {
	return s.Description
}

func (s *SaveApBasicConfigRequest) GetEchoInt() *int32 {
	return s.EchoInt
}

func (s *SaveApBasicConfigRequest) GetFailover() *int32 {
	return s.Failover
}

func (s *SaveApBasicConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *SaveApBasicConfigRequest) GetInsecureAllowed() *int32 {
	return s.InsecureAllowed
}

func (s *SaveApBasicConfigRequest) GetLanIp() *string {
	return s.LanIp
}

func (s *SaveApBasicConfigRequest) GetLanMask() *string {
	return s.LanMask
}

func (s *SaveApBasicConfigRequest) GetLanStatus() *int32 {
	return s.LanStatus
}

func (s *SaveApBasicConfigRequest) GetLogIp() *string {
	return s.LogIp
}

func (s *SaveApBasicConfigRequest) GetLogLevel() *int32 {
	return s.LogLevel
}

func (s *SaveApBasicConfigRequest) GetName() *string {
	return s.Name
}

func (s *SaveApBasicConfigRequest) GetPasswd() *string {
	return s.Passwd
}

func (s *SaveApBasicConfigRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *SaveApBasicConfigRequest) GetRoute() *string {
	return s.Route
}

func (s *SaveApBasicConfigRequest) GetScan() *int32 {
	return s.Scan
}

func (s *SaveApBasicConfigRequest) GetTokenServer() *string {
	return s.TokenServer
}

func (s *SaveApBasicConfigRequest) GetVpn() *string {
	return s.Vpn
}

func (s *SaveApBasicConfigRequest) GetWorkMode() *int32 {
	return s.WorkMode
}

func (s *SaveApBasicConfigRequest) SetAppCode(v string) *SaveApBasicConfigRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetAppName(v string) *SaveApBasicConfigRequest {
	s.AppName = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetCountry(v string) *SaveApBasicConfigRequest {
	s.Country = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetDai(v string) *SaveApBasicConfigRequest {
	s.Dai = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetDescription(v string) *SaveApBasicConfigRequest {
	s.Description = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetEchoInt(v int32) *SaveApBasicConfigRequest {
	s.EchoInt = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetFailover(v int32) *SaveApBasicConfigRequest {
	s.Failover = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetId(v int64) *SaveApBasicConfigRequest {
	s.Id = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetInsecureAllowed(v int32) *SaveApBasicConfigRequest {
	s.InsecureAllowed = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetLanIp(v string) *SaveApBasicConfigRequest {
	s.LanIp = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetLanMask(v string) *SaveApBasicConfigRequest {
	s.LanMask = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetLanStatus(v int32) *SaveApBasicConfigRequest {
	s.LanStatus = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetLogIp(v string) *SaveApBasicConfigRequest {
	s.LogIp = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetLogLevel(v int32) *SaveApBasicConfigRequest {
	s.LogLevel = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetName(v string) *SaveApBasicConfigRequest {
	s.Name = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetPasswd(v string) *SaveApBasicConfigRequest {
	s.Passwd = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetProtocol(v string) *SaveApBasicConfigRequest {
	s.Protocol = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetRoute(v string) *SaveApBasicConfigRequest {
	s.Route = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetScan(v int32) *SaveApBasicConfigRequest {
	s.Scan = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetTokenServer(v string) *SaveApBasicConfigRequest {
	s.TokenServer = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetVpn(v string) *SaveApBasicConfigRequest {
	s.Vpn = &v
	return s
}

func (s *SaveApBasicConfigRequest) SetWorkMode(v int32) *SaveApBasicConfigRequest {
	s.WorkMode = &v
	return s
}

func (s *SaveApBasicConfigRequest) Validate() error {
	return dara.Validate(s)
}
