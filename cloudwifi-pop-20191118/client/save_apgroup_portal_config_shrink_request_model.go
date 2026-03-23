// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupPortalConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApgroupId(v int64) *SaveApgroupPortalConfigShrinkRequest
	GetApgroupId() *int64
	SetAppAuthUrl(v string) *SaveApgroupPortalConfigShrinkRequest
	GetAppAuthUrl() *string
	SetAppCode(v string) *SaveApgroupPortalConfigShrinkRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApgroupPortalConfigShrinkRequest
	GetAppName() *string
	SetAuthKey(v string) *SaveApgroupPortalConfigShrinkRequest
	GetAuthKey() *string
	SetAuthSecret(v string) *SaveApgroupPortalConfigShrinkRequest
	GetAuthSecret() *string
	SetCheckUrl(v string) *SaveApgroupPortalConfigShrinkRequest
	GetCheckUrl() *string
	SetClientDownload(v int32) *SaveApgroupPortalConfigShrinkRequest
	GetClientDownload() *int32
	SetClientUpload(v int32) *SaveApgroupPortalConfigShrinkRequest
	GetClientUpload() *int32
	SetCountdown(v int32) *SaveApgroupPortalConfigShrinkRequest
	GetCountdown() *int32
	SetNetwork(v int32) *SaveApgroupPortalConfigShrinkRequest
	GetNetwork() *int32
	SetPortalTypesShrink(v string) *SaveApgroupPortalConfigShrinkRequest
	GetPortalTypesShrink() *string
	SetPortalUrl(v string) *SaveApgroupPortalConfigShrinkRequest
	GetPortalUrl() *string
	SetTimeStamp(v int64) *SaveApgroupPortalConfigShrinkRequest
	GetTimeStamp() *int64
	SetTotalDownload(v int32) *SaveApgroupPortalConfigShrinkRequest
	GetTotalDownload() *int32
	SetTotalUpload(v int32) *SaveApgroupPortalConfigShrinkRequest
	GetTotalUpload() *int32
	SetWebAuthUrl(v string) *SaveApgroupPortalConfigShrinkRequest
	GetWebAuthUrl() *string
	SetWhitelist(v string) *SaveApgroupPortalConfigShrinkRequest
	GetWhitelist() *string
}

type SaveApgroupPortalConfigShrinkRequest struct {
	// This parameter is required.
	ApgroupId  *int64  `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
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

func (s SaveApgroupPortalConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupPortalConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetApgroupId() *int64 {
	return s.ApgroupId
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetAppAuthUrl() *string {
	return s.AppAuthUrl
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetAuthSecret() *string {
	return s.AuthSecret
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetClientDownload() *int32 {
	return s.ClientDownload
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetClientUpload() *int32 {
	return s.ClientUpload
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetCountdown() *int32 {
	return s.Countdown
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetNetwork() *int32 {
	return s.Network
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetPortalTypesShrink() *string {
	return s.PortalTypesShrink
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetPortalUrl() *string {
	return s.PortalUrl
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetTotalDownload() *int32 {
	return s.TotalDownload
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetTotalUpload() *int32 {
	return s.TotalUpload
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetWebAuthUrl() *string {
	return s.WebAuthUrl
}

func (s *SaveApgroupPortalConfigShrinkRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetApgroupId(v int64) *SaveApgroupPortalConfigShrinkRequest {
	s.ApgroupId = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetAppAuthUrl(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.AppAuthUrl = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetAppCode(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetAppName(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.AppName = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetAuthKey(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.AuthKey = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetAuthSecret(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.AuthSecret = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetCheckUrl(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.CheckUrl = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetClientDownload(v int32) *SaveApgroupPortalConfigShrinkRequest {
	s.ClientDownload = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetClientUpload(v int32) *SaveApgroupPortalConfigShrinkRequest {
	s.ClientUpload = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetCountdown(v int32) *SaveApgroupPortalConfigShrinkRequest {
	s.Countdown = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetNetwork(v int32) *SaveApgroupPortalConfigShrinkRequest {
	s.Network = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetPortalTypesShrink(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.PortalTypesShrink = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetPortalUrl(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.PortalUrl = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetTimeStamp(v int64) *SaveApgroupPortalConfigShrinkRequest {
	s.TimeStamp = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetTotalDownload(v int32) *SaveApgroupPortalConfigShrinkRequest {
	s.TotalDownload = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetTotalUpload(v int32) *SaveApgroupPortalConfigShrinkRequest {
	s.TotalUpload = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetWebAuthUrl(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.WebAuthUrl = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) SetWhitelist(v string) *SaveApgroupPortalConfigShrinkRequest {
	s.Whitelist = &v
	return s
}

func (s *SaveApgroupPortalConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
