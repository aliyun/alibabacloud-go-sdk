// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApSsidConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcctPort(v int32) *SaveApSsidConfigRequest
	GetAcctPort() *int32
	SetAcctSecret(v string) *SaveApSsidConfigRequest
	GetAcctSecret() *string
	SetAcctServer(v string) *SaveApSsidConfigRequest
	GetAcctServer() *string
	SetAcctStatusServerWork(v int32) *SaveApSsidConfigRequest
	GetAcctStatusServerWork() *int32
	SetApAssetId(v int64) *SaveApSsidConfigRequest
	GetApAssetId() *int64
	SetAppCode(v string) *SaveApSsidConfigRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApSsidConfigRequest
	GetAppName() *string
	SetArpProxyEnable(v int32) *SaveApSsidConfigRequest
	GetArpProxyEnable() *int32
	SetAuthCache(v string) *SaveApSsidConfigRequest
	GetAuthCache() *string
	SetAuthPort(v int32) *SaveApSsidConfigRequest
	GetAuthPort() *int32
	SetAuthSecret(v string) *SaveApSsidConfigRequest
	GetAuthSecret() *string
	SetAuthServer(v string) *SaveApSsidConfigRequest
	GetAuthServer() *string
	SetAuthStatusServerWork(v int32) *SaveApSsidConfigRequest
	GetAuthStatusServerWork() *int32
	SetCir(v int64) *SaveApSsidConfigRequest
	GetCir() *int64
	SetCirStep(v int64) *SaveApSsidConfigRequest
	GetCirStep() *int64
	SetCirType(v int32) *SaveApSsidConfigRequest
	GetCirType() *int32
	SetCirUl(v int64) *SaveApSsidConfigRequest
	GetCirUl() *int64
	SetDaeClient(v string) *SaveApSsidConfigRequest
	GetDaeClient() *string
	SetDaePort(v int32) *SaveApSsidConfigRequest
	GetDaePort() *int32
	SetDaeSecret(v string) *SaveApSsidConfigRequest
	GetDaeSecret() *string
	SetDisabled(v string) *SaveApSsidConfigRequest
	GetDisabled() *string
	SetDisassocLowAck(v string) *SaveApSsidConfigRequest
	GetDisassocLowAck() *string
	SetDisassocWeakRssi(v int32) *SaveApSsidConfigRequest
	GetDisassocWeakRssi() *int32
	SetDynamicVlan(v int32) *SaveApSsidConfigRequest
	GetDynamicVlan() *int32
	SetEncKey(v string) *SaveApSsidConfigRequest
	GetEncKey() *string
	SetEncryption(v string) *SaveApSsidConfigRequest
	GetEncryption() *string
	SetFourthAuthPort(v int32) *SaveApSsidConfigRequest
	GetFourthAuthPort() *int32
	SetFourthAuthSecret(v string) *SaveApSsidConfigRequest
	GetFourthAuthSecret() *string
	SetFourthAuthServer(v string) *SaveApSsidConfigRequest
	GetFourthAuthServer() *string
	SetFtOverDs(v int32) *SaveApSsidConfigRequest
	GetFtOverDs() *int32
	SetHidden(v string) *SaveApSsidConfigRequest
	GetHidden() *string
	SetId(v int64) *SaveApSsidConfigRequest
	GetId() *int64
	SetIeee80211r(v int32) *SaveApSsidConfigRequest
	GetIeee80211r() *int32
	SetIeee80211w(v string) *SaveApSsidConfigRequest
	GetIeee80211w() *string
	SetIgnoreWeakProbe(v int32) *SaveApSsidConfigRequest
	GetIgnoreWeakProbe() *int32
	SetIsolate(v string) *SaveApSsidConfigRequest
	GetIsolate() *string
	SetLiteEffect(v bool) *SaveApSsidConfigRequest
	GetLiteEffect() *bool
	SetMac(v string) *SaveApSsidConfigRequest
	GetMac() *string
	SetMaxInactivity(v int32) *SaveApSsidConfigRequest
	GetMaxInactivity() *int32
	SetMaxassoc(v int32) *SaveApSsidConfigRequest
	GetMaxassoc() *int32
	SetMobilityDomain(v string) *SaveApSsidConfigRequest
	GetMobilityDomain() *string
	SetMulticastForward(v int32) *SaveApSsidConfigRequest
	GetMulticastForward() *int32
	SetNasid(v string) *SaveApSsidConfigRequest
	GetNasid() *string
	SetNdProxyWork(v int32) *SaveApSsidConfigRequest
	GetNdProxyWork() *int32
	SetNetwork(v int32) *SaveApSsidConfigRequest
	GetNetwork() *int32
	SetOwnip(v string) *SaveApSsidConfigRequest
	GetOwnip() *string
	SetRadioIndex(v string) *SaveApSsidConfigRequest
	GetRadioIndex() *string
	SetSecondaryAcctPort(v int32) *SaveApSsidConfigRequest
	GetSecondaryAcctPort() *int32
	SetSecondaryAcctSecret(v string) *SaveApSsidConfigRequest
	GetSecondaryAcctSecret() *string
	SetSecondaryAcctServer(v string) *SaveApSsidConfigRequest
	GetSecondaryAcctServer() *string
	SetSecondaryAuthPort(v int32) *SaveApSsidConfigRequest
	GetSecondaryAuthPort() *int32
	SetSecondaryAuthSecret(v string) *SaveApSsidConfigRequest
	GetSecondaryAuthSecret() *string
	SetSecondaryAuthServer(v string) *SaveApSsidConfigRequest
	GetSecondaryAuthServer() *string
	SetSendConfigToAp(v bool) *SaveApSsidConfigRequest
	GetSendConfigToAp() *bool
	SetShortPreamble(v string) *SaveApSsidConfigRequest
	GetShortPreamble() *string
	SetSsid(v string) *SaveApSsidConfigRequest
	GetSsid() *string
	SetSsidLb(v int32) *SaveApSsidConfigRequest
	GetSsidLb() *int32
	SetThirdAuthPort(v int32) *SaveApSsidConfigRequest
	GetThirdAuthPort() *int32
	SetThirdAuthSecret(v string) *SaveApSsidConfigRequest
	GetThirdAuthSecret() *string
	SetThirdAuthServer(v string) *SaveApSsidConfigRequest
	GetThirdAuthServer() *string
	SetType(v int32) *SaveApSsidConfigRequest
	GetType() *int32
	SetVlanDhcp(v int32) *SaveApSsidConfigRequest
	GetVlanDhcp() *int32
	SetWmm(v string) *SaveApSsidConfigRequest
	GetWmm() *string
}

type SaveApSsidConfigRequest struct {
	AcctPort             *int32  `json:"AcctPort,omitempty" xml:"AcctPort,omitempty"`
	AcctSecret           *string `json:"AcctSecret,omitempty" xml:"AcctSecret,omitempty"`
	AcctServer           *string `json:"AcctServer,omitempty" xml:"AcctServer,omitempty"`
	AcctStatusServerWork *int32  `json:"AcctStatusServerWork,omitempty" xml:"AcctStatusServerWork,omitempty"`
	ApAssetId            *int64  `json:"ApAssetId,omitempty" xml:"ApAssetId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ArpProxyEnable *int32  `json:"ArpProxyEnable,omitempty" xml:"ArpProxyEnable,omitempty"`
	// This parameter is required.
	AuthCache            *string `json:"AuthCache,omitempty" xml:"AuthCache,omitempty"`
	AuthPort             *int32  `json:"AuthPort,omitempty" xml:"AuthPort,omitempty"`
	AuthSecret           *string `json:"AuthSecret,omitempty" xml:"AuthSecret,omitempty"`
	AuthServer           *string `json:"AuthServer,omitempty" xml:"AuthServer,omitempty"`
	AuthStatusServerWork *int32  `json:"AuthStatusServerWork,omitempty" xml:"AuthStatusServerWork,omitempty"`
	Cir                  *int64  `json:"Cir,omitempty" xml:"Cir,omitempty"`
	CirStep              *int64  `json:"CirStep,omitempty" xml:"CirStep,omitempty"`
	CirType              *int32  `json:"CirType,omitempty" xml:"CirType,omitempty"`
	CirUl                *int64  `json:"CirUl,omitempty" xml:"CirUl,omitempty"`
	DaeClient            *string `json:"DaeClient,omitempty" xml:"DaeClient,omitempty"`
	DaePort              *int32  `json:"DaePort,omitempty" xml:"DaePort,omitempty"`
	DaeSecret            *string `json:"DaeSecret,omitempty" xml:"DaeSecret,omitempty"`
	// This parameter is required.
	Disabled *string `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is required.
	DisassocLowAck *string `json:"DisassocLowAck,omitempty" xml:"DisassocLowAck,omitempty"`
	// This parameter is required.
	DisassocWeakRssi *int32 `json:"DisassocWeakRssi,omitempty" xml:"DisassocWeakRssi,omitempty"`
	// This parameter is required.
	DynamicVlan *int32  `json:"DynamicVlan,omitempty" xml:"DynamicVlan,omitempty"`
	EncKey      *string `json:"EncKey,omitempty" xml:"EncKey,omitempty"`
	// This parameter is required.
	Encryption       *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	FourthAuthPort   *int32  `json:"FourthAuthPort,omitempty" xml:"FourthAuthPort,omitempty"`
	FourthAuthSecret *string `json:"FourthAuthSecret,omitempty" xml:"FourthAuthSecret,omitempty"`
	FourthAuthServer *string `json:"FourthAuthServer,omitempty" xml:"FourthAuthServer,omitempty"`
	FtOverDs         *int32  `json:"FtOverDs,omitempty" xml:"FtOverDs,omitempty"`
	// This parameter is required.
	Hidden     *string `json:"Hidden,omitempty" xml:"Hidden,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Ieee80211r *int32  `json:"Ieee80211r,omitempty" xml:"Ieee80211r,omitempty"`
	// This parameter is required.
	Ieee80211w      *string `json:"Ieee80211w,omitempty" xml:"Ieee80211w,omitempty"`
	IgnoreWeakProbe *int32  `json:"IgnoreWeakProbe,omitempty" xml:"IgnoreWeakProbe,omitempty"`
	// This parameter is required.
	Isolate    *string `json:"Isolate,omitempty" xml:"Isolate,omitempty"`
	LiteEffect *bool   `json:"LiteEffect,omitempty" xml:"LiteEffect,omitempty"`
	// This parameter is required.
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// This parameter is required.
	MaxInactivity *int32 `json:"MaxInactivity,omitempty" xml:"MaxInactivity,omitempty"`
	// This parameter is required.
	Maxassoc         *int32  `json:"Maxassoc,omitempty" xml:"Maxassoc,omitempty"`
	MobilityDomain   *string `json:"MobilityDomain,omitempty" xml:"MobilityDomain,omitempty"`
	MulticastForward *int32  `json:"MulticastForward,omitempty" xml:"MulticastForward,omitempty"`
	Nasid            *string `json:"Nasid,omitempty" xml:"Nasid,omitempty"`
	NdProxyWork      *int32  `json:"NdProxyWork,omitempty" xml:"NdProxyWork,omitempty"`
	// This parameter is required.
	Network *int32  `json:"Network,omitempty" xml:"Network,omitempty"`
	Ownip   *string `json:"Ownip,omitempty" xml:"Ownip,omitempty"`
	// This parameter is required.
	RadioIndex          *string `json:"RadioIndex,omitempty" xml:"RadioIndex,omitempty"`
	SecondaryAcctPort   *int32  `json:"SecondaryAcctPort,omitempty" xml:"SecondaryAcctPort,omitempty"`
	SecondaryAcctSecret *string `json:"SecondaryAcctSecret,omitempty" xml:"SecondaryAcctSecret,omitempty"`
	SecondaryAcctServer *string `json:"SecondaryAcctServer,omitempty" xml:"SecondaryAcctServer,omitempty"`
	SecondaryAuthPort   *int32  `json:"SecondaryAuthPort,omitempty" xml:"SecondaryAuthPort,omitempty"`
	SecondaryAuthSecret *string `json:"SecondaryAuthSecret,omitempty" xml:"SecondaryAuthSecret,omitempty"`
	SecondaryAuthServer *string `json:"SecondaryAuthServer,omitempty" xml:"SecondaryAuthServer,omitempty"`
	SendConfigToAp      *bool   `json:"SendConfigToAp,omitempty" xml:"SendConfigToAp,omitempty"`
	// This parameter is required.
	ShortPreamble *string `json:"ShortPreamble,omitempty" xml:"ShortPreamble,omitempty"`
	// This parameter is required.
	Ssid *string `json:"Ssid,omitempty" xml:"Ssid,omitempty"`
	// This parameter is required.
	SsidLb          *int32  `json:"SsidLb,omitempty" xml:"SsidLb,omitempty"`
	ThirdAuthPort   *int32  `json:"ThirdAuthPort,omitempty" xml:"ThirdAuthPort,omitempty"`
	ThirdAuthSecret *string `json:"ThirdAuthSecret,omitempty" xml:"ThirdAuthSecret,omitempty"`
	ThirdAuthServer *string `json:"ThirdAuthServer,omitempty" xml:"ThirdAuthServer,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	VlanDhcp *int32 `json:"VlanDhcp,omitempty" xml:"VlanDhcp,omitempty"`
	// This parameter is required.
	Wmm *string `json:"Wmm,omitempty" xml:"Wmm,omitempty"`
}

func (s SaveApSsidConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApSsidConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveApSsidConfigRequest) GetAcctPort() *int32 {
	return s.AcctPort
}

func (s *SaveApSsidConfigRequest) GetAcctSecret() *string {
	return s.AcctSecret
}

func (s *SaveApSsidConfigRequest) GetAcctServer() *string {
	return s.AcctServer
}

func (s *SaveApSsidConfigRequest) GetAcctStatusServerWork() *int32 {
	return s.AcctStatusServerWork
}

func (s *SaveApSsidConfigRequest) GetApAssetId() *int64 {
	return s.ApAssetId
}

func (s *SaveApSsidConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApSsidConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApSsidConfigRequest) GetArpProxyEnable() *int32 {
	return s.ArpProxyEnable
}

func (s *SaveApSsidConfigRequest) GetAuthCache() *string {
	return s.AuthCache
}

func (s *SaveApSsidConfigRequest) GetAuthPort() *int32 {
	return s.AuthPort
}

func (s *SaveApSsidConfigRequest) GetAuthSecret() *string {
	return s.AuthSecret
}

func (s *SaveApSsidConfigRequest) GetAuthServer() *string {
	return s.AuthServer
}

func (s *SaveApSsidConfigRequest) GetAuthStatusServerWork() *int32 {
	return s.AuthStatusServerWork
}

func (s *SaveApSsidConfigRequest) GetCir() *int64 {
	return s.Cir
}

func (s *SaveApSsidConfigRequest) GetCirStep() *int64 {
	return s.CirStep
}

func (s *SaveApSsidConfigRequest) GetCirType() *int32 {
	return s.CirType
}

func (s *SaveApSsidConfigRequest) GetCirUl() *int64 {
	return s.CirUl
}

func (s *SaveApSsidConfigRequest) GetDaeClient() *string {
	return s.DaeClient
}

func (s *SaveApSsidConfigRequest) GetDaePort() *int32 {
	return s.DaePort
}

func (s *SaveApSsidConfigRequest) GetDaeSecret() *string {
	return s.DaeSecret
}

func (s *SaveApSsidConfigRequest) GetDisabled() *string {
	return s.Disabled
}

func (s *SaveApSsidConfigRequest) GetDisassocLowAck() *string {
	return s.DisassocLowAck
}

func (s *SaveApSsidConfigRequest) GetDisassocWeakRssi() *int32 {
	return s.DisassocWeakRssi
}

func (s *SaveApSsidConfigRequest) GetDynamicVlan() *int32 {
	return s.DynamicVlan
}

func (s *SaveApSsidConfigRequest) GetEncKey() *string {
	return s.EncKey
}

func (s *SaveApSsidConfigRequest) GetEncryption() *string {
	return s.Encryption
}

func (s *SaveApSsidConfigRequest) GetFourthAuthPort() *int32 {
	return s.FourthAuthPort
}

func (s *SaveApSsidConfigRequest) GetFourthAuthSecret() *string {
	return s.FourthAuthSecret
}

func (s *SaveApSsidConfigRequest) GetFourthAuthServer() *string {
	return s.FourthAuthServer
}

func (s *SaveApSsidConfigRequest) GetFtOverDs() *int32 {
	return s.FtOverDs
}

func (s *SaveApSsidConfigRequest) GetHidden() *string {
	return s.Hidden
}

func (s *SaveApSsidConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *SaveApSsidConfigRequest) GetIeee80211r() *int32 {
	return s.Ieee80211r
}

func (s *SaveApSsidConfigRequest) GetIeee80211w() *string {
	return s.Ieee80211w
}

func (s *SaveApSsidConfigRequest) GetIgnoreWeakProbe() *int32 {
	return s.IgnoreWeakProbe
}

func (s *SaveApSsidConfigRequest) GetIsolate() *string {
	return s.Isolate
}

func (s *SaveApSsidConfigRequest) GetLiteEffect() *bool {
	return s.LiteEffect
}

func (s *SaveApSsidConfigRequest) GetMac() *string {
	return s.Mac
}

func (s *SaveApSsidConfigRequest) GetMaxInactivity() *int32 {
	return s.MaxInactivity
}

func (s *SaveApSsidConfigRequest) GetMaxassoc() *int32 {
	return s.Maxassoc
}

func (s *SaveApSsidConfigRequest) GetMobilityDomain() *string {
	return s.MobilityDomain
}

func (s *SaveApSsidConfigRequest) GetMulticastForward() *int32 {
	return s.MulticastForward
}

func (s *SaveApSsidConfigRequest) GetNasid() *string {
	return s.Nasid
}

func (s *SaveApSsidConfigRequest) GetNdProxyWork() *int32 {
	return s.NdProxyWork
}

func (s *SaveApSsidConfigRequest) GetNetwork() *int32 {
	return s.Network
}

func (s *SaveApSsidConfigRequest) GetOwnip() *string {
	return s.Ownip
}

func (s *SaveApSsidConfigRequest) GetRadioIndex() *string {
	return s.RadioIndex
}

func (s *SaveApSsidConfigRequest) GetSecondaryAcctPort() *int32 {
	return s.SecondaryAcctPort
}

func (s *SaveApSsidConfigRequest) GetSecondaryAcctSecret() *string {
	return s.SecondaryAcctSecret
}

func (s *SaveApSsidConfigRequest) GetSecondaryAcctServer() *string {
	return s.SecondaryAcctServer
}

func (s *SaveApSsidConfigRequest) GetSecondaryAuthPort() *int32 {
	return s.SecondaryAuthPort
}

func (s *SaveApSsidConfigRequest) GetSecondaryAuthSecret() *string {
	return s.SecondaryAuthSecret
}

func (s *SaveApSsidConfigRequest) GetSecondaryAuthServer() *string {
	return s.SecondaryAuthServer
}

func (s *SaveApSsidConfigRequest) GetSendConfigToAp() *bool {
	return s.SendConfigToAp
}

func (s *SaveApSsidConfigRequest) GetShortPreamble() *string {
	return s.ShortPreamble
}

func (s *SaveApSsidConfigRequest) GetSsid() *string {
	return s.Ssid
}

func (s *SaveApSsidConfigRequest) GetSsidLb() *int32 {
	return s.SsidLb
}

func (s *SaveApSsidConfigRequest) GetThirdAuthPort() *int32 {
	return s.ThirdAuthPort
}

func (s *SaveApSsidConfigRequest) GetThirdAuthSecret() *string {
	return s.ThirdAuthSecret
}

func (s *SaveApSsidConfigRequest) GetThirdAuthServer() *string {
	return s.ThirdAuthServer
}

func (s *SaveApSsidConfigRequest) GetType() *int32 {
	return s.Type
}

func (s *SaveApSsidConfigRequest) GetVlanDhcp() *int32 {
	return s.VlanDhcp
}

func (s *SaveApSsidConfigRequest) GetWmm() *string {
	return s.Wmm
}

func (s *SaveApSsidConfigRequest) SetAcctPort(v int32) *SaveApSsidConfigRequest {
	s.AcctPort = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAcctSecret(v string) *SaveApSsidConfigRequest {
	s.AcctSecret = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAcctServer(v string) *SaveApSsidConfigRequest {
	s.AcctServer = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAcctStatusServerWork(v int32) *SaveApSsidConfigRequest {
	s.AcctStatusServerWork = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetApAssetId(v int64) *SaveApSsidConfigRequest {
	s.ApAssetId = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAppCode(v string) *SaveApSsidConfigRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAppName(v string) *SaveApSsidConfigRequest {
	s.AppName = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetArpProxyEnable(v int32) *SaveApSsidConfigRequest {
	s.ArpProxyEnable = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAuthCache(v string) *SaveApSsidConfigRequest {
	s.AuthCache = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAuthPort(v int32) *SaveApSsidConfigRequest {
	s.AuthPort = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAuthSecret(v string) *SaveApSsidConfigRequest {
	s.AuthSecret = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAuthServer(v string) *SaveApSsidConfigRequest {
	s.AuthServer = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetAuthStatusServerWork(v int32) *SaveApSsidConfigRequest {
	s.AuthStatusServerWork = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetCir(v int64) *SaveApSsidConfigRequest {
	s.Cir = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetCirStep(v int64) *SaveApSsidConfigRequest {
	s.CirStep = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetCirType(v int32) *SaveApSsidConfigRequest {
	s.CirType = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetCirUl(v int64) *SaveApSsidConfigRequest {
	s.CirUl = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetDaeClient(v string) *SaveApSsidConfigRequest {
	s.DaeClient = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetDaePort(v int32) *SaveApSsidConfigRequest {
	s.DaePort = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetDaeSecret(v string) *SaveApSsidConfigRequest {
	s.DaeSecret = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetDisabled(v string) *SaveApSsidConfigRequest {
	s.Disabled = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetDisassocLowAck(v string) *SaveApSsidConfigRequest {
	s.DisassocLowAck = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetDisassocWeakRssi(v int32) *SaveApSsidConfigRequest {
	s.DisassocWeakRssi = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetDynamicVlan(v int32) *SaveApSsidConfigRequest {
	s.DynamicVlan = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetEncKey(v string) *SaveApSsidConfigRequest {
	s.EncKey = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetEncryption(v string) *SaveApSsidConfigRequest {
	s.Encryption = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetFourthAuthPort(v int32) *SaveApSsidConfigRequest {
	s.FourthAuthPort = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetFourthAuthSecret(v string) *SaveApSsidConfigRequest {
	s.FourthAuthSecret = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetFourthAuthServer(v string) *SaveApSsidConfigRequest {
	s.FourthAuthServer = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetFtOverDs(v int32) *SaveApSsidConfigRequest {
	s.FtOverDs = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetHidden(v string) *SaveApSsidConfigRequest {
	s.Hidden = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetId(v int64) *SaveApSsidConfigRequest {
	s.Id = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetIeee80211r(v int32) *SaveApSsidConfigRequest {
	s.Ieee80211r = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetIeee80211w(v string) *SaveApSsidConfigRequest {
	s.Ieee80211w = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetIgnoreWeakProbe(v int32) *SaveApSsidConfigRequest {
	s.IgnoreWeakProbe = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetIsolate(v string) *SaveApSsidConfigRequest {
	s.Isolate = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetLiteEffect(v bool) *SaveApSsidConfigRequest {
	s.LiteEffect = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetMac(v string) *SaveApSsidConfigRequest {
	s.Mac = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetMaxInactivity(v int32) *SaveApSsidConfigRequest {
	s.MaxInactivity = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetMaxassoc(v int32) *SaveApSsidConfigRequest {
	s.Maxassoc = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetMobilityDomain(v string) *SaveApSsidConfigRequest {
	s.MobilityDomain = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetMulticastForward(v int32) *SaveApSsidConfigRequest {
	s.MulticastForward = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetNasid(v string) *SaveApSsidConfigRequest {
	s.Nasid = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetNdProxyWork(v int32) *SaveApSsidConfigRequest {
	s.NdProxyWork = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetNetwork(v int32) *SaveApSsidConfigRequest {
	s.Network = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetOwnip(v string) *SaveApSsidConfigRequest {
	s.Ownip = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetRadioIndex(v string) *SaveApSsidConfigRequest {
	s.RadioIndex = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSecondaryAcctPort(v int32) *SaveApSsidConfigRequest {
	s.SecondaryAcctPort = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSecondaryAcctSecret(v string) *SaveApSsidConfigRequest {
	s.SecondaryAcctSecret = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSecondaryAcctServer(v string) *SaveApSsidConfigRequest {
	s.SecondaryAcctServer = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSecondaryAuthPort(v int32) *SaveApSsidConfigRequest {
	s.SecondaryAuthPort = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSecondaryAuthSecret(v string) *SaveApSsidConfigRequest {
	s.SecondaryAuthSecret = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSecondaryAuthServer(v string) *SaveApSsidConfigRequest {
	s.SecondaryAuthServer = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSendConfigToAp(v bool) *SaveApSsidConfigRequest {
	s.SendConfigToAp = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetShortPreamble(v string) *SaveApSsidConfigRequest {
	s.ShortPreamble = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSsid(v string) *SaveApSsidConfigRequest {
	s.Ssid = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetSsidLb(v int32) *SaveApSsidConfigRequest {
	s.SsidLb = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetThirdAuthPort(v int32) *SaveApSsidConfigRequest {
	s.ThirdAuthPort = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetThirdAuthSecret(v string) *SaveApSsidConfigRequest {
	s.ThirdAuthSecret = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetThirdAuthServer(v string) *SaveApSsidConfigRequest {
	s.ThirdAuthServer = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetType(v int32) *SaveApSsidConfigRequest {
	s.Type = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetVlanDhcp(v int32) *SaveApSsidConfigRequest {
	s.VlanDhcp = &v
	return s
}

func (s *SaveApSsidConfigRequest) SetWmm(v string) *SaveApSsidConfigRequest {
	s.Wmm = &v
	return s
}

func (s *SaveApSsidConfigRequest) Validate() error {
	return dara.Validate(s)
}
