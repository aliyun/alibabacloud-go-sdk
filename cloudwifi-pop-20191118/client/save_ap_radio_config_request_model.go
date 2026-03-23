// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApRadioConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *SaveApRadioConfigRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApRadioConfigRequest
	GetAppName() *string
	SetBcastRate(v int32) *SaveApRadioConfigRequest
	GetBcastRate() *int32
	SetBeaconInt(v int32) *SaveApRadioConfigRequest
	GetBeaconInt() *int32
	SetChannel(v string) *SaveApRadioConfigRequest
	GetChannel() *string
	SetDisabled(v string) *SaveApRadioConfigRequest
	GetDisabled() *string
	SetFrag(v int32) *SaveApRadioConfigRequest
	GetFrag() *int32
	SetHtmode(v string) *SaveApRadioConfigRequest
	GetHtmode() *string
	SetHwmode(v string) *SaveApRadioConfigRequest
	GetHwmode() *string
	SetId(v int64) *SaveApRadioConfigRequest
	GetId() *int64
	SetMcastRate(v int32) *SaveApRadioConfigRequest
	GetMcastRate() *int32
	SetMgmtRate(v int32) *SaveApRadioConfigRequest
	GetMgmtRate() *int32
	SetMinrate(v int32) *SaveApRadioConfigRequest
	GetMinrate() *int32
	SetNoscan(v string) *SaveApRadioConfigRequest
	GetNoscan() *string
	SetProbereq(v string) *SaveApRadioConfigRequest
	GetProbereq() *string
	SetRequireMode(v string) *SaveApRadioConfigRequest
	GetRequireMode() *string
	SetRts(v int32) *SaveApRadioConfigRequest
	GetRts() *int32
	SetShortgi(v string) *SaveApRadioConfigRequest
	GetShortgi() *string
	SetTxpower(v string) *SaveApRadioConfigRequest
	GetTxpower() *string
	SetUapsd(v int32) *SaveApRadioConfigRequest
	GetUapsd() *int32
}

type SaveApRadioConfigRequest struct {
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	BcastRate *int32 `json:"BcastRate,omitempty" xml:"BcastRate,omitempty"`
	// This parameter is required.
	BeaconInt *int32 `json:"BeaconInt,omitempty" xml:"BeaconInt,omitempty"`
	// This parameter is required.
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// This parameter is required.
	Disabled *string `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is required.
	Frag *int32 `json:"Frag,omitempty" xml:"Frag,omitempty"`
	// This parameter is required.
	Htmode *string `json:"Htmode,omitempty" xml:"Htmode,omitempty"`
	// This parameter is required.
	Hwmode *string `json:"Hwmode,omitempty" xml:"Hwmode,omitempty"`
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	McastRate *int32 `json:"McastRate,omitempty" xml:"McastRate,omitempty"`
	// This parameter is required.
	MgmtRate *int32 `json:"MgmtRate,omitempty" xml:"MgmtRate,omitempty"`
	// This parameter is required.
	Minrate *int32 `json:"Minrate,omitempty" xml:"Minrate,omitempty"`
	// This parameter is required.
	Noscan *string `json:"Noscan,omitempty" xml:"Noscan,omitempty"`
	// This parameter is required.
	Probereq    *string `json:"Probereq,omitempty" xml:"Probereq,omitempty"`
	RequireMode *string `json:"RequireMode,omitempty" xml:"RequireMode,omitempty"`
	// This parameter is required.
	Rts *int32 `json:"Rts,omitempty" xml:"Rts,omitempty"`
	// This parameter is required.
	Shortgi *string `json:"Shortgi,omitempty" xml:"Shortgi,omitempty"`
	// This parameter is required.
	Txpower *string `json:"Txpower,omitempty" xml:"Txpower,omitempty"`
	// This parameter is required.
	Uapsd *int32 `json:"Uapsd,omitempty" xml:"Uapsd,omitempty"`
}

func (s SaveApRadioConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApRadioConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveApRadioConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApRadioConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApRadioConfigRequest) GetBcastRate() *int32 {
	return s.BcastRate
}

func (s *SaveApRadioConfigRequest) GetBeaconInt() *int32 {
	return s.BeaconInt
}

func (s *SaveApRadioConfigRequest) GetChannel() *string {
	return s.Channel
}

func (s *SaveApRadioConfigRequest) GetDisabled() *string {
	return s.Disabled
}

func (s *SaveApRadioConfigRequest) GetFrag() *int32 {
	return s.Frag
}

func (s *SaveApRadioConfigRequest) GetHtmode() *string {
	return s.Htmode
}

func (s *SaveApRadioConfigRequest) GetHwmode() *string {
	return s.Hwmode
}

func (s *SaveApRadioConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *SaveApRadioConfigRequest) GetMcastRate() *int32 {
	return s.McastRate
}

func (s *SaveApRadioConfigRequest) GetMgmtRate() *int32 {
	return s.MgmtRate
}

func (s *SaveApRadioConfigRequest) GetMinrate() *int32 {
	return s.Minrate
}

func (s *SaveApRadioConfigRequest) GetNoscan() *string {
	return s.Noscan
}

func (s *SaveApRadioConfigRequest) GetProbereq() *string {
	return s.Probereq
}

func (s *SaveApRadioConfigRequest) GetRequireMode() *string {
	return s.RequireMode
}

func (s *SaveApRadioConfigRequest) GetRts() *int32 {
	return s.Rts
}

func (s *SaveApRadioConfigRequest) GetShortgi() *string {
	return s.Shortgi
}

func (s *SaveApRadioConfigRequest) GetTxpower() *string {
	return s.Txpower
}

func (s *SaveApRadioConfigRequest) GetUapsd() *int32 {
	return s.Uapsd
}

func (s *SaveApRadioConfigRequest) SetAppCode(v string) *SaveApRadioConfigRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetAppName(v string) *SaveApRadioConfigRequest {
	s.AppName = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetBcastRate(v int32) *SaveApRadioConfigRequest {
	s.BcastRate = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetBeaconInt(v int32) *SaveApRadioConfigRequest {
	s.BeaconInt = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetChannel(v string) *SaveApRadioConfigRequest {
	s.Channel = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetDisabled(v string) *SaveApRadioConfigRequest {
	s.Disabled = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetFrag(v int32) *SaveApRadioConfigRequest {
	s.Frag = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetHtmode(v string) *SaveApRadioConfigRequest {
	s.Htmode = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetHwmode(v string) *SaveApRadioConfigRequest {
	s.Hwmode = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetId(v int64) *SaveApRadioConfigRequest {
	s.Id = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetMcastRate(v int32) *SaveApRadioConfigRequest {
	s.McastRate = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetMgmtRate(v int32) *SaveApRadioConfigRequest {
	s.MgmtRate = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetMinrate(v int32) *SaveApRadioConfigRequest {
	s.Minrate = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetNoscan(v string) *SaveApRadioConfigRequest {
	s.Noscan = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetProbereq(v string) *SaveApRadioConfigRequest {
	s.Probereq = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetRequireMode(v string) *SaveApRadioConfigRequest {
	s.RequireMode = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetRts(v int32) *SaveApRadioConfigRequest {
	s.Rts = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetShortgi(v string) *SaveApRadioConfigRequest {
	s.Shortgi = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetTxpower(v string) *SaveApRadioConfigRequest {
	s.Txpower = &v
	return s
}

func (s *SaveApRadioConfigRequest) SetUapsd(v int32) *SaveApRadioConfigRequest {
	s.Uapsd = &v
	return s
}

func (s *SaveApRadioConfigRequest) Validate() error {
	return dara.Validate(s)
}
