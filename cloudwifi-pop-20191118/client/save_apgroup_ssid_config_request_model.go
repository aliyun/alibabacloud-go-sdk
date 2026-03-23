// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupSsidConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcctPort(v int32) *SaveApgroupSsidConfigRequest
	GetAcctPort() *int32
	SetAcctSecret(v string) *SaveApgroupSsidConfigRequest
	GetAcctSecret() *string
	SetAcctServer(v string) *SaveApgroupSsidConfigRequest
	GetAcctServer() *string
	SetApgroupId(v string) *SaveApgroupSsidConfigRequest
	GetApgroupId() *string
	SetAppCode(v string) *SaveApgroupSsidConfigRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApgroupSsidConfigRequest
	GetAppName() *string
	SetAuthCache(v string) *SaveApgroupSsidConfigRequest
	GetAuthCache() *string
	SetAuthPort(v int32) *SaveApgroupSsidConfigRequest
	GetAuthPort() *int32
	SetAuthSecret(v string) *SaveApgroupSsidConfigRequest
	GetAuthSecret() *string
	SetAuthServer(v string) *SaveApgroupSsidConfigRequest
	GetAuthServer() *string
	SetBinding(v string) *SaveApgroupSsidConfigRequest
	GetBinding() *string
	SetCir(v int64) *SaveApgroupSsidConfigRequest
	GetCir() *int64
	SetDaeClient(v string) *SaveApgroupSsidConfigRequest
	GetDaeClient() *string
	SetDaePort(v int32) *SaveApgroupSsidConfigRequest
	GetDaePort() *int32
	SetDaeSecret(v string) *SaveApgroupSsidConfigRequest
	GetDaeSecret() *string
	SetDisabled(v string) *SaveApgroupSsidConfigRequest
	GetDisabled() *string
	SetDisassocLowAck(v string) *SaveApgroupSsidConfigRequest
	GetDisassocLowAck() *string
	SetDisassocWeakRssi(v int32) *SaveApgroupSsidConfigRequest
	GetDisassocWeakRssi() *int32
	SetDynamicVlan(v int32) *SaveApgroupSsidConfigRequest
	GetDynamicVlan() *int32
	SetEffect(v bool) *SaveApgroupSsidConfigRequest
	GetEffect() *bool
	SetEncKey(v string) *SaveApgroupSsidConfigRequest
	GetEncKey() *string
	SetEncryption(v string) *SaveApgroupSsidConfigRequest
	GetEncryption() *string
	SetHidden(v string) *SaveApgroupSsidConfigRequest
	GetHidden() *string
	SetId(v int64) *SaveApgroupSsidConfigRequest
	GetId() *int64
	SetIeee80211w(v string) *SaveApgroupSsidConfigRequest
	GetIeee80211w() *string
	SetIgnoreWeakProbe(v int32) *SaveApgroupSsidConfigRequest
	GetIgnoreWeakProbe() *int32
	SetIsolate(v string) *SaveApgroupSsidConfigRequest
	GetIsolate() *string
	SetLiteEffect(v bool) *SaveApgroupSsidConfigRequest
	GetLiteEffect() *bool
	SetMaxInactivity(v int32) *SaveApgroupSsidConfigRequest
	GetMaxInactivity() *int32
	SetMaxassoc(v string) *SaveApgroupSsidConfigRequest
	GetMaxassoc() *string
	SetMulticastForward(v int32) *SaveApgroupSsidConfigRequest
	GetMulticastForward() *int32
	SetNasid(v string) *SaveApgroupSsidConfigRequest
	GetNasid() *string
	SetNetwork(v int32) *SaveApgroupSsidConfigRequest
	GetNetwork() *int32
	SetNewSsid(v string) *SaveApgroupSsidConfigRequest
	GetNewSsid() *string
	SetOwnip(v string) *SaveApgroupSsidConfigRequest
	GetOwnip() *string
	SetSecondaryAcctPort(v int32) *SaveApgroupSsidConfigRequest
	GetSecondaryAcctPort() *int32
	SetSecondaryAcctSecret(v string) *SaveApgroupSsidConfigRequest
	GetSecondaryAcctSecret() *string
	SetSecondaryAcctServer(v string) *SaveApgroupSsidConfigRequest
	GetSecondaryAcctServer() *string
	SetSecondaryAuthPort(v int32) *SaveApgroupSsidConfigRequest
	GetSecondaryAuthPort() *int32
	SetSecondaryAuthSecret(v string) *SaveApgroupSsidConfigRequest
	GetSecondaryAuthSecret() *string
	SetSecondaryAuthServer(v string) *SaveApgroupSsidConfigRequest
	GetSecondaryAuthServer() *string
	SetShortPreamble(v string) *SaveApgroupSsidConfigRequest
	GetShortPreamble() *string
	SetSsid(v string) *SaveApgroupSsidConfigRequest
	GetSsid() *string
	SetSsidLb(v int32) *SaveApgroupSsidConfigRequest
	GetSsidLb() *int32
	SetType(v int32) *SaveApgroupSsidConfigRequest
	GetType() *int32
	SetVlanDhcp(v int32) *SaveApgroupSsidConfigRequest
	GetVlanDhcp() *int32
	SetWmm(v string) *SaveApgroupSsidConfigRequest
	GetWmm() *string
}

type SaveApgroupSsidConfigRequest struct {
	AcctPort   *int32  `json:"AcctPort,omitempty" xml:"AcctPort,omitempty"`
	AcctSecret *string `json:"AcctSecret,omitempty" xml:"AcctSecret,omitempty"`
	AcctServer *string `json:"AcctServer,omitempty" xml:"AcctServer,omitempty"`
	// This parameter is required.
	ApgroupId *string `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AuthCache  *string `json:"AuthCache,omitempty" xml:"AuthCache,omitempty"`
	AuthPort   *int32  `json:"AuthPort,omitempty" xml:"AuthPort,omitempty"`
	AuthSecret *string `json:"AuthSecret,omitempty" xml:"AuthSecret,omitempty"`
	AuthServer *string `json:"AuthServer,omitempty" xml:"AuthServer,omitempty"`
	// This parameter is required.
	Binding          *string `json:"Binding,omitempty" xml:"Binding,omitempty"`
	Cir              *int64  `json:"Cir,omitempty" xml:"Cir,omitempty"`
	DaeClient        *string `json:"DaeClient,omitempty" xml:"DaeClient,omitempty"`
	DaePort          *int32  `json:"DaePort,omitempty" xml:"DaePort,omitempty"`
	DaeSecret        *string `json:"DaeSecret,omitempty" xml:"DaeSecret,omitempty"`
	Disabled         *string `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DisassocLowAck   *string `json:"DisassocLowAck,omitempty" xml:"DisassocLowAck,omitempty"`
	DisassocWeakRssi *int32  `json:"DisassocWeakRssi,omitempty" xml:"DisassocWeakRssi,omitempty"`
	DynamicVlan      *int32  `json:"DynamicVlan,omitempty" xml:"DynamicVlan,omitempty"`
	Effect           *bool   `json:"Effect,omitempty" xml:"Effect,omitempty"`
	EncKey           *string `json:"EncKey,omitempty" xml:"EncKey,omitempty"`
	// This parameter is required.
	Encryption       *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	Hidden           *string `json:"Hidden,omitempty" xml:"Hidden,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Ieee80211w       *string `json:"Ieee80211w,omitempty" xml:"Ieee80211w,omitempty"`
	IgnoreWeakProbe  *int32  `json:"IgnoreWeakProbe,omitempty" xml:"IgnoreWeakProbe,omitempty"`
	Isolate          *string `json:"Isolate,omitempty" xml:"Isolate,omitempty"`
	LiteEffect       *bool   `json:"LiteEffect,omitempty" xml:"LiteEffect,omitempty"`
	MaxInactivity    *int32  `json:"MaxInactivity,omitempty" xml:"MaxInactivity,omitempty"`
	Maxassoc         *string `json:"Maxassoc,omitempty" xml:"Maxassoc,omitempty"`
	MulticastForward *int32  `json:"MulticastForward,omitempty" xml:"MulticastForward,omitempty"`
	Nasid            *string `json:"Nasid,omitempty" xml:"Nasid,omitempty"`
	// This parameter is required.
	Network *int32 `json:"Network,omitempty" xml:"Network,omitempty"`
	// This parameter is required.
	NewSsid             *string `json:"NewSsid,omitempty" xml:"NewSsid,omitempty"`
	Ownip               *string `json:"Ownip,omitempty" xml:"Ownip,omitempty"`
	SecondaryAcctPort   *int32  `json:"SecondaryAcctPort,omitempty" xml:"SecondaryAcctPort,omitempty"`
	SecondaryAcctSecret *string `json:"SecondaryAcctSecret,omitempty" xml:"SecondaryAcctSecret,omitempty"`
	SecondaryAcctServer *string `json:"SecondaryAcctServer,omitempty" xml:"SecondaryAcctServer,omitempty"`
	SecondaryAuthPort   *int32  `json:"SecondaryAuthPort,omitempty" xml:"SecondaryAuthPort,omitempty"`
	SecondaryAuthSecret *string `json:"SecondaryAuthSecret,omitempty" xml:"SecondaryAuthSecret,omitempty"`
	SecondaryAuthServer *string `json:"SecondaryAuthServer,omitempty" xml:"SecondaryAuthServer,omitempty"`
	ShortPreamble       *string `json:"ShortPreamble,omitempty" xml:"ShortPreamble,omitempty"`
	// This parameter is required.
	Ssid     *string `json:"Ssid,omitempty" xml:"Ssid,omitempty"`
	SsidLb   *int32  `json:"SsidLb,omitempty" xml:"SsidLb,omitempty"`
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	VlanDhcp *int32  `json:"VlanDhcp,omitempty" xml:"VlanDhcp,omitempty"`
	Wmm      *string `json:"Wmm,omitempty" xml:"Wmm,omitempty"`
}

func (s SaveApgroupSsidConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupSsidConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveApgroupSsidConfigRequest) GetAcctPort() *int32 {
	return s.AcctPort
}

func (s *SaveApgroupSsidConfigRequest) GetAcctSecret() *string {
	return s.AcctSecret
}

func (s *SaveApgroupSsidConfigRequest) GetAcctServer() *string {
	return s.AcctServer
}

func (s *SaveApgroupSsidConfigRequest) GetApgroupId() *string {
	return s.ApgroupId
}

func (s *SaveApgroupSsidConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApgroupSsidConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApgroupSsidConfigRequest) GetAuthCache() *string {
	return s.AuthCache
}

func (s *SaveApgroupSsidConfigRequest) GetAuthPort() *int32 {
	return s.AuthPort
}

func (s *SaveApgroupSsidConfigRequest) GetAuthSecret() *string {
	return s.AuthSecret
}

func (s *SaveApgroupSsidConfigRequest) GetAuthServer() *string {
	return s.AuthServer
}

func (s *SaveApgroupSsidConfigRequest) GetBinding() *string {
	return s.Binding
}

func (s *SaveApgroupSsidConfigRequest) GetCir() *int64 {
	return s.Cir
}

func (s *SaveApgroupSsidConfigRequest) GetDaeClient() *string {
	return s.DaeClient
}

func (s *SaveApgroupSsidConfigRequest) GetDaePort() *int32 {
	return s.DaePort
}

func (s *SaveApgroupSsidConfigRequest) GetDaeSecret() *string {
	return s.DaeSecret
}

func (s *SaveApgroupSsidConfigRequest) GetDisabled() *string {
	return s.Disabled
}

func (s *SaveApgroupSsidConfigRequest) GetDisassocLowAck() *string {
	return s.DisassocLowAck
}

func (s *SaveApgroupSsidConfigRequest) GetDisassocWeakRssi() *int32 {
	return s.DisassocWeakRssi
}

func (s *SaveApgroupSsidConfigRequest) GetDynamicVlan() *int32 {
	return s.DynamicVlan
}

func (s *SaveApgroupSsidConfigRequest) GetEffect() *bool {
	return s.Effect
}

func (s *SaveApgroupSsidConfigRequest) GetEncKey() *string {
	return s.EncKey
}

func (s *SaveApgroupSsidConfigRequest) GetEncryption() *string {
	return s.Encryption
}

func (s *SaveApgroupSsidConfigRequest) GetHidden() *string {
	return s.Hidden
}

func (s *SaveApgroupSsidConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *SaveApgroupSsidConfigRequest) GetIeee80211w() *string {
	return s.Ieee80211w
}

func (s *SaveApgroupSsidConfigRequest) GetIgnoreWeakProbe() *int32 {
	return s.IgnoreWeakProbe
}

func (s *SaveApgroupSsidConfigRequest) GetIsolate() *string {
	return s.Isolate
}

func (s *SaveApgroupSsidConfigRequest) GetLiteEffect() *bool {
	return s.LiteEffect
}

func (s *SaveApgroupSsidConfigRequest) GetMaxInactivity() *int32 {
	return s.MaxInactivity
}

func (s *SaveApgroupSsidConfigRequest) GetMaxassoc() *string {
	return s.Maxassoc
}

func (s *SaveApgroupSsidConfigRequest) GetMulticastForward() *int32 {
	return s.MulticastForward
}

func (s *SaveApgroupSsidConfigRequest) GetNasid() *string {
	return s.Nasid
}

func (s *SaveApgroupSsidConfigRequest) GetNetwork() *int32 {
	return s.Network
}

func (s *SaveApgroupSsidConfigRequest) GetNewSsid() *string {
	return s.NewSsid
}

func (s *SaveApgroupSsidConfigRequest) GetOwnip() *string {
	return s.Ownip
}

func (s *SaveApgroupSsidConfigRequest) GetSecondaryAcctPort() *int32 {
	return s.SecondaryAcctPort
}

func (s *SaveApgroupSsidConfigRequest) GetSecondaryAcctSecret() *string {
	return s.SecondaryAcctSecret
}

func (s *SaveApgroupSsidConfigRequest) GetSecondaryAcctServer() *string {
	return s.SecondaryAcctServer
}

func (s *SaveApgroupSsidConfigRequest) GetSecondaryAuthPort() *int32 {
	return s.SecondaryAuthPort
}

func (s *SaveApgroupSsidConfigRequest) GetSecondaryAuthSecret() *string {
	return s.SecondaryAuthSecret
}

func (s *SaveApgroupSsidConfigRequest) GetSecondaryAuthServer() *string {
	return s.SecondaryAuthServer
}

func (s *SaveApgroupSsidConfigRequest) GetShortPreamble() *string {
	return s.ShortPreamble
}

func (s *SaveApgroupSsidConfigRequest) GetSsid() *string {
	return s.Ssid
}

func (s *SaveApgroupSsidConfigRequest) GetSsidLb() *int32 {
	return s.SsidLb
}

func (s *SaveApgroupSsidConfigRequest) GetType() *int32 {
	return s.Type
}

func (s *SaveApgroupSsidConfigRequest) GetVlanDhcp() *int32 {
	return s.VlanDhcp
}

func (s *SaveApgroupSsidConfigRequest) GetWmm() *string {
	return s.Wmm
}

func (s *SaveApgroupSsidConfigRequest) SetAcctPort(v int32) *SaveApgroupSsidConfigRequest {
	s.AcctPort = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetAcctSecret(v string) *SaveApgroupSsidConfigRequest {
	s.AcctSecret = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetAcctServer(v string) *SaveApgroupSsidConfigRequest {
	s.AcctServer = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetApgroupId(v string) *SaveApgroupSsidConfigRequest {
	s.ApgroupId = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetAppCode(v string) *SaveApgroupSsidConfigRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetAppName(v string) *SaveApgroupSsidConfigRequest {
	s.AppName = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetAuthCache(v string) *SaveApgroupSsidConfigRequest {
	s.AuthCache = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetAuthPort(v int32) *SaveApgroupSsidConfigRequest {
	s.AuthPort = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetAuthSecret(v string) *SaveApgroupSsidConfigRequest {
	s.AuthSecret = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetAuthServer(v string) *SaveApgroupSsidConfigRequest {
	s.AuthServer = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetBinding(v string) *SaveApgroupSsidConfigRequest {
	s.Binding = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetCir(v int64) *SaveApgroupSsidConfigRequest {
	s.Cir = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetDaeClient(v string) *SaveApgroupSsidConfigRequest {
	s.DaeClient = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetDaePort(v int32) *SaveApgroupSsidConfigRequest {
	s.DaePort = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetDaeSecret(v string) *SaveApgroupSsidConfigRequest {
	s.DaeSecret = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetDisabled(v string) *SaveApgroupSsidConfigRequest {
	s.Disabled = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetDisassocLowAck(v string) *SaveApgroupSsidConfigRequest {
	s.DisassocLowAck = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetDisassocWeakRssi(v int32) *SaveApgroupSsidConfigRequest {
	s.DisassocWeakRssi = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetDynamicVlan(v int32) *SaveApgroupSsidConfigRequest {
	s.DynamicVlan = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetEffect(v bool) *SaveApgroupSsidConfigRequest {
	s.Effect = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetEncKey(v string) *SaveApgroupSsidConfigRequest {
	s.EncKey = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetEncryption(v string) *SaveApgroupSsidConfigRequest {
	s.Encryption = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetHidden(v string) *SaveApgroupSsidConfigRequest {
	s.Hidden = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetId(v int64) *SaveApgroupSsidConfigRequest {
	s.Id = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetIeee80211w(v string) *SaveApgroupSsidConfigRequest {
	s.Ieee80211w = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetIgnoreWeakProbe(v int32) *SaveApgroupSsidConfigRequest {
	s.IgnoreWeakProbe = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetIsolate(v string) *SaveApgroupSsidConfigRequest {
	s.Isolate = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetLiteEffect(v bool) *SaveApgroupSsidConfigRequest {
	s.LiteEffect = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetMaxInactivity(v int32) *SaveApgroupSsidConfigRequest {
	s.MaxInactivity = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetMaxassoc(v string) *SaveApgroupSsidConfigRequest {
	s.Maxassoc = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetMulticastForward(v int32) *SaveApgroupSsidConfigRequest {
	s.MulticastForward = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetNasid(v string) *SaveApgroupSsidConfigRequest {
	s.Nasid = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetNetwork(v int32) *SaveApgroupSsidConfigRequest {
	s.Network = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetNewSsid(v string) *SaveApgroupSsidConfigRequest {
	s.NewSsid = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetOwnip(v string) *SaveApgroupSsidConfigRequest {
	s.Ownip = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetSecondaryAcctPort(v int32) *SaveApgroupSsidConfigRequest {
	s.SecondaryAcctPort = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetSecondaryAcctSecret(v string) *SaveApgroupSsidConfigRequest {
	s.SecondaryAcctSecret = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetSecondaryAcctServer(v string) *SaveApgroupSsidConfigRequest {
	s.SecondaryAcctServer = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetSecondaryAuthPort(v int32) *SaveApgroupSsidConfigRequest {
	s.SecondaryAuthPort = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetSecondaryAuthSecret(v string) *SaveApgroupSsidConfigRequest {
	s.SecondaryAuthSecret = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetSecondaryAuthServer(v string) *SaveApgroupSsidConfigRequest {
	s.SecondaryAuthServer = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetShortPreamble(v string) *SaveApgroupSsidConfigRequest {
	s.ShortPreamble = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetSsid(v string) *SaveApgroupSsidConfigRequest {
	s.Ssid = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetSsidLb(v int32) *SaveApgroupSsidConfigRequest {
	s.SsidLb = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetType(v int32) *SaveApgroupSsidConfigRequest {
	s.Type = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetVlanDhcp(v int32) *SaveApgroupSsidConfigRequest {
	s.VlanDhcp = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) SetWmm(v string) *SaveApgroupSsidConfigRequest {
	s.Wmm = &v
	return s
}

func (s *SaveApgroupSsidConfigRequest) Validate() error {
	return dara.Validate(s)
}
