// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApPortalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApConfigId(v int64) *SaveApPortalConfigRequest
	GetApConfigId() *int64
	SetAppAuthUrl(v string) *SaveApPortalConfigRequest
	GetAppAuthUrl() *string
	SetAppCode(v string) *SaveApPortalConfigRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApPortalConfigRequest
	GetAppName() *string
	SetAuthKey(v string) *SaveApPortalConfigRequest
	GetAuthKey() *string
	SetAuthSecret(v string) *SaveApPortalConfigRequest
	GetAuthSecret() *string
	SetCheckUrl(v string) *SaveApPortalConfigRequest
	GetCheckUrl() *string
	SetClientDownload(v int32) *SaveApPortalConfigRequest
	GetClientDownload() *int32
	SetClientUpload(v int32) *SaveApPortalConfigRequest
	GetClientUpload() *int32
	SetCountdown(v int32) *SaveApPortalConfigRequest
	GetCountdown() *int32
	SetNetwork(v int32) *SaveApPortalConfigRequest
	GetNetwork() *int32
	SetPortalTypes(v []*string) *SaveApPortalConfigRequest
	GetPortalTypes() []*string
	SetPortalUrl(v string) *SaveApPortalConfigRequest
	GetPortalUrl() *string
	SetTimeStamp(v int64) *SaveApPortalConfigRequest
	GetTimeStamp() *int64
	SetTotalDownload(v int32) *SaveApPortalConfigRequest
	GetTotalDownload() *int32
	SetTotalUpload(v int32) *SaveApPortalConfigRequest
	GetTotalUpload() *int32
	SetWebAuthUrl(v string) *SaveApPortalConfigRequest
	GetWebAuthUrl() *string
	SetWhitelist(v string) *SaveApPortalConfigRequest
	GetWhitelist() *string
}

type SaveApPortalConfigRequest struct {
	// This parameter is required.
	ApConfigId *int64  `json:"ApConfigId,omitempty" xml:"ApConfigId,omitempty"`
	AppAuthUrl *string `json:"AppAuthUrl,omitempty" xml:"AppAuthUrl,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName        *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AuthKey        *string   `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSecret     *string   `json:"AuthSecret,omitempty" xml:"AuthSecret,omitempty"`
	CheckUrl       *string   `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	ClientDownload *int32    `json:"ClientDownload,omitempty" xml:"ClientDownload,omitempty"`
	ClientUpload   *int32    `json:"ClientUpload,omitempty" xml:"ClientUpload,omitempty"`
	Countdown      *int32    `json:"Countdown,omitempty" xml:"Countdown,omitempty"`
	Network        *int32    `json:"Network,omitempty" xml:"Network,omitempty"`
	PortalTypes    []*string `json:"PortalTypes,omitempty" xml:"PortalTypes,omitempty" type:"Repeated"`
	PortalUrl      *string   `json:"PortalUrl,omitempty" xml:"PortalUrl,omitempty"`
	TimeStamp      *int64    `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TotalDownload  *int32    `json:"TotalDownload,omitempty" xml:"TotalDownload,omitempty"`
	TotalUpload    *int32    `json:"TotalUpload,omitempty" xml:"TotalUpload,omitempty"`
	WebAuthUrl     *string   `json:"WebAuthUrl,omitempty" xml:"WebAuthUrl,omitempty"`
	Whitelist      *string   `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s SaveApPortalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApPortalConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveApPortalConfigRequest) GetApConfigId() *int64 {
	return s.ApConfigId
}

func (s *SaveApPortalConfigRequest) GetAppAuthUrl() *string {
	return s.AppAuthUrl
}

func (s *SaveApPortalConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApPortalConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApPortalConfigRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *SaveApPortalConfigRequest) GetAuthSecret() *string {
	return s.AuthSecret
}

func (s *SaveApPortalConfigRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *SaveApPortalConfigRequest) GetClientDownload() *int32 {
	return s.ClientDownload
}

func (s *SaveApPortalConfigRequest) GetClientUpload() *int32 {
	return s.ClientUpload
}

func (s *SaveApPortalConfigRequest) GetCountdown() *int32 {
	return s.Countdown
}

func (s *SaveApPortalConfigRequest) GetNetwork() *int32 {
	return s.Network
}

func (s *SaveApPortalConfigRequest) GetPortalTypes() []*string {
	return s.PortalTypes
}

func (s *SaveApPortalConfigRequest) GetPortalUrl() *string {
	return s.PortalUrl
}

func (s *SaveApPortalConfigRequest) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *SaveApPortalConfigRequest) GetTotalDownload() *int32 {
	return s.TotalDownload
}

func (s *SaveApPortalConfigRequest) GetTotalUpload() *int32 {
	return s.TotalUpload
}

func (s *SaveApPortalConfigRequest) GetWebAuthUrl() *string {
	return s.WebAuthUrl
}

func (s *SaveApPortalConfigRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *SaveApPortalConfigRequest) SetApConfigId(v int64) *SaveApPortalConfigRequest {
	s.ApConfigId = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetAppAuthUrl(v string) *SaveApPortalConfigRequest {
	s.AppAuthUrl = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetAppCode(v string) *SaveApPortalConfigRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetAppName(v string) *SaveApPortalConfigRequest {
	s.AppName = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetAuthKey(v string) *SaveApPortalConfigRequest {
	s.AuthKey = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetAuthSecret(v string) *SaveApPortalConfigRequest {
	s.AuthSecret = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetCheckUrl(v string) *SaveApPortalConfigRequest {
	s.CheckUrl = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetClientDownload(v int32) *SaveApPortalConfigRequest {
	s.ClientDownload = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetClientUpload(v int32) *SaveApPortalConfigRequest {
	s.ClientUpload = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetCountdown(v int32) *SaveApPortalConfigRequest {
	s.Countdown = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetNetwork(v int32) *SaveApPortalConfigRequest {
	s.Network = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetPortalTypes(v []*string) *SaveApPortalConfigRequest {
	s.PortalTypes = v
	return s
}

func (s *SaveApPortalConfigRequest) SetPortalUrl(v string) *SaveApPortalConfigRequest {
	s.PortalUrl = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetTimeStamp(v int64) *SaveApPortalConfigRequest {
	s.TimeStamp = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetTotalDownload(v int32) *SaveApPortalConfigRequest {
	s.TotalDownload = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetTotalUpload(v int32) *SaveApPortalConfigRequest {
	s.TotalUpload = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetWebAuthUrl(v string) *SaveApPortalConfigRequest {
	s.WebAuthUrl = &v
	return s
}

func (s *SaveApPortalConfigRequest) SetWhitelist(v string) *SaveApPortalConfigRequest {
	s.Whitelist = &v
	return s
}

func (s *SaveApPortalConfigRequest) Validate() error {
	return dara.Validate(s)
}
