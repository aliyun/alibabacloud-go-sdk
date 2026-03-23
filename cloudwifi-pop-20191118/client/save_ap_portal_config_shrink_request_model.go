// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApPortalConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApConfigId(v int64) *SaveApPortalConfigShrinkRequest
	GetApConfigId() *int64
	SetAppAuthUrl(v string) *SaveApPortalConfigShrinkRequest
	GetAppAuthUrl() *string
	SetAppCode(v string) *SaveApPortalConfigShrinkRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApPortalConfigShrinkRequest
	GetAppName() *string
	SetAuthKey(v string) *SaveApPortalConfigShrinkRequest
	GetAuthKey() *string
	SetAuthSecret(v string) *SaveApPortalConfigShrinkRequest
	GetAuthSecret() *string
	SetCheckUrl(v string) *SaveApPortalConfigShrinkRequest
	GetCheckUrl() *string
	SetClientDownload(v int32) *SaveApPortalConfigShrinkRequest
	GetClientDownload() *int32
	SetClientUpload(v int32) *SaveApPortalConfigShrinkRequest
	GetClientUpload() *int32
	SetCountdown(v int32) *SaveApPortalConfigShrinkRequest
	GetCountdown() *int32
	SetNetwork(v int32) *SaveApPortalConfigShrinkRequest
	GetNetwork() *int32
	SetPortalTypesShrink(v string) *SaveApPortalConfigShrinkRequest
	GetPortalTypesShrink() *string
	SetPortalUrl(v string) *SaveApPortalConfigShrinkRequest
	GetPortalUrl() *string
	SetTimeStamp(v int64) *SaveApPortalConfigShrinkRequest
	GetTimeStamp() *int64
	SetTotalDownload(v int32) *SaveApPortalConfigShrinkRequest
	GetTotalDownload() *int32
	SetTotalUpload(v int32) *SaveApPortalConfigShrinkRequest
	GetTotalUpload() *int32
	SetWebAuthUrl(v string) *SaveApPortalConfigShrinkRequest
	GetWebAuthUrl() *string
	SetWhitelist(v string) *SaveApPortalConfigShrinkRequest
	GetWhitelist() *string
}

type SaveApPortalConfigShrinkRequest struct {
	// This parameter is required.
	ApConfigId *int64  `json:"ApConfigId,omitempty" xml:"ApConfigId,omitempty"`
	AppAuthUrl *string `json:"AppAuthUrl,omitempty" xml:"AppAuthUrl,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName           *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AuthKey           *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSecret        *string `json:"AuthSecret,omitempty" xml:"AuthSecret,omitempty"`
	CheckUrl          *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	ClientDownload    *int32  `json:"ClientDownload,omitempty" xml:"ClientDownload,omitempty"`
	ClientUpload      *int32  `json:"ClientUpload,omitempty" xml:"ClientUpload,omitempty"`
	Countdown         *int32  `json:"Countdown,omitempty" xml:"Countdown,omitempty"`
	Network           *int32  `json:"Network,omitempty" xml:"Network,omitempty"`
	PortalTypesShrink *string `json:"PortalTypes,omitempty" xml:"PortalTypes,omitempty"`
	PortalUrl         *string `json:"PortalUrl,omitempty" xml:"PortalUrl,omitempty"`
	TimeStamp         *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TotalDownload     *int32  `json:"TotalDownload,omitempty" xml:"TotalDownload,omitempty"`
	TotalUpload       *int32  `json:"TotalUpload,omitempty" xml:"TotalUpload,omitempty"`
	WebAuthUrl        *string `json:"WebAuthUrl,omitempty" xml:"WebAuthUrl,omitempty"`
	Whitelist         *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s SaveApPortalConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApPortalConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveApPortalConfigShrinkRequest) GetApConfigId() *int64 {
	return s.ApConfigId
}

func (s *SaveApPortalConfigShrinkRequest) GetAppAuthUrl() *string {
	return s.AppAuthUrl
}

func (s *SaveApPortalConfigShrinkRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApPortalConfigShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApPortalConfigShrinkRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *SaveApPortalConfigShrinkRequest) GetAuthSecret() *string {
	return s.AuthSecret
}

func (s *SaveApPortalConfigShrinkRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *SaveApPortalConfigShrinkRequest) GetClientDownload() *int32 {
	return s.ClientDownload
}

func (s *SaveApPortalConfigShrinkRequest) GetClientUpload() *int32 {
	return s.ClientUpload
}

func (s *SaveApPortalConfigShrinkRequest) GetCountdown() *int32 {
	return s.Countdown
}

func (s *SaveApPortalConfigShrinkRequest) GetNetwork() *int32 {
	return s.Network
}

func (s *SaveApPortalConfigShrinkRequest) GetPortalTypesShrink() *string {
	return s.PortalTypesShrink
}

func (s *SaveApPortalConfigShrinkRequest) GetPortalUrl() *string {
	return s.PortalUrl
}

func (s *SaveApPortalConfigShrinkRequest) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *SaveApPortalConfigShrinkRequest) GetTotalDownload() *int32 {
	return s.TotalDownload
}

func (s *SaveApPortalConfigShrinkRequest) GetTotalUpload() *int32 {
	return s.TotalUpload
}

func (s *SaveApPortalConfigShrinkRequest) GetWebAuthUrl() *string {
	return s.WebAuthUrl
}

func (s *SaveApPortalConfigShrinkRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *SaveApPortalConfigShrinkRequest) SetApConfigId(v int64) *SaveApPortalConfigShrinkRequest {
	s.ApConfigId = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetAppAuthUrl(v string) *SaveApPortalConfigShrinkRequest {
	s.AppAuthUrl = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetAppCode(v string) *SaveApPortalConfigShrinkRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetAppName(v string) *SaveApPortalConfigShrinkRequest {
	s.AppName = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetAuthKey(v string) *SaveApPortalConfigShrinkRequest {
	s.AuthKey = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetAuthSecret(v string) *SaveApPortalConfigShrinkRequest {
	s.AuthSecret = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetCheckUrl(v string) *SaveApPortalConfigShrinkRequest {
	s.CheckUrl = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetClientDownload(v int32) *SaveApPortalConfigShrinkRequest {
	s.ClientDownload = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetClientUpload(v int32) *SaveApPortalConfigShrinkRequest {
	s.ClientUpload = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetCountdown(v int32) *SaveApPortalConfigShrinkRequest {
	s.Countdown = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetNetwork(v int32) *SaveApPortalConfigShrinkRequest {
	s.Network = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetPortalTypesShrink(v string) *SaveApPortalConfigShrinkRequest {
	s.PortalTypesShrink = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetPortalUrl(v string) *SaveApPortalConfigShrinkRequest {
	s.PortalUrl = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetTimeStamp(v int64) *SaveApPortalConfigShrinkRequest {
	s.TimeStamp = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetTotalDownload(v int32) *SaveApPortalConfigShrinkRequest {
	s.TotalDownload = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetTotalUpload(v int32) *SaveApPortalConfigShrinkRequest {
	s.TotalUpload = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetWebAuthUrl(v string) *SaveApPortalConfigShrinkRequest {
	s.WebAuthUrl = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) SetWhitelist(v string) *SaveApPortalConfigShrinkRequest {
	s.Whitelist = &v
	return s
}

func (s *SaveApPortalConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
