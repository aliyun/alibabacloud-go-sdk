// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApgroupPortalConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApgroupId(v int64) *SaveApgroupPortalConfigRequest
	GetApgroupId() *int64
	SetAppAuthUrl(v string) *SaveApgroupPortalConfigRequest
	GetAppAuthUrl() *string
	SetAppCode(v string) *SaveApgroupPortalConfigRequest
	GetAppCode() *string
	SetAppName(v string) *SaveApgroupPortalConfigRequest
	GetAppName() *string
	SetAuthKey(v string) *SaveApgroupPortalConfigRequest
	GetAuthKey() *string
	SetAuthSecret(v string) *SaveApgroupPortalConfigRequest
	GetAuthSecret() *string
	SetCheckUrl(v string) *SaveApgroupPortalConfigRequest
	GetCheckUrl() *string
	SetClientDownload(v int32) *SaveApgroupPortalConfigRequest
	GetClientDownload() *int32
	SetClientUpload(v int32) *SaveApgroupPortalConfigRequest
	GetClientUpload() *int32
	SetCountdown(v int32) *SaveApgroupPortalConfigRequest
	GetCountdown() *int32
	SetNetwork(v int32) *SaveApgroupPortalConfigRequest
	GetNetwork() *int32
	SetPortalTypes(v []*string) *SaveApgroupPortalConfigRequest
	GetPortalTypes() []*string
	SetPortalUrl(v string) *SaveApgroupPortalConfigRequest
	GetPortalUrl() *string
	SetTimeStamp(v int64) *SaveApgroupPortalConfigRequest
	GetTimeStamp() *int64
	SetTotalDownload(v int32) *SaveApgroupPortalConfigRequest
	GetTotalDownload() *int32
	SetTotalUpload(v int32) *SaveApgroupPortalConfigRequest
	GetTotalUpload() *int32
	SetWebAuthUrl(v string) *SaveApgroupPortalConfigRequest
	GetWebAuthUrl() *string
	SetWhitelist(v string) *SaveApgroupPortalConfigRequest
	GetWhitelist() *string
}

type SaveApgroupPortalConfigRequest struct {
	// This parameter is required.
	ApgroupId  *int64  `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
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

func (s SaveApgroupPortalConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveApgroupPortalConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveApgroupPortalConfigRequest) GetApgroupId() *int64 {
	return s.ApgroupId
}

func (s *SaveApgroupPortalConfigRequest) GetAppAuthUrl() *string {
	return s.AppAuthUrl
}

func (s *SaveApgroupPortalConfigRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SaveApgroupPortalConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveApgroupPortalConfigRequest) GetAuthKey() *string {
	return s.AuthKey
}

func (s *SaveApgroupPortalConfigRequest) GetAuthSecret() *string {
	return s.AuthSecret
}

func (s *SaveApgroupPortalConfigRequest) GetCheckUrl() *string {
	return s.CheckUrl
}

func (s *SaveApgroupPortalConfigRequest) GetClientDownload() *int32 {
	return s.ClientDownload
}

func (s *SaveApgroupPortalConfigRequest) GetClientUpload() *int32 {
	return s.ClientUpload
}

func (s *SaveApgroupPortalConfigRequest) GetCountdown() *int32 {
	return s.Countdown
}

func (s *SaveApgroupPortalConfigRequest) GetNetwork() *int32 {
	return s.Network
}

func (s *SaveApgroupPortalConfigRequest) GetPortalTypes() []*string {
	return s.PortalTypes
}

func (s *SaveApgroupPortalConfigRequest) GetPortalUrl() *string {
	return s.PortalUrl
}

func (s *SaveApgroupPortalConfigRequest) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *SaveApgroupPortalConfigRequest) GetTotalDownload() *int32 {
	return s.TotalDownload
}

func (s *SaveApgroupPortalConfigRequest) GetTotalUpload() *int32 {
	return s.TotalUpload
}

func (s *SaveApgroupPortalConfigRequest) GetWebAuthUrl() *string {
	return s.WebAuthUrl
}

func (s *SaveApgroupPortalConfigRequest) GetWhitelist() *string {
	return s.Whitelist
}

func (s *SaveApgroupPortalConfigRequest) SetApgroupId(v int64) *SaveApgroupPortalConfigRequest {
	s.ApgroupId = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetAppAuthUrl(v string) *SaveApgroupPortalConfigRequest {
	s.AppAuthUrl = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetAppCode(v string) *SaveApgroupPortalConfigRequest {
	s.AppCode = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetAppName(v string) *SaveApgroupPortalConfigRequest {
	s.AppName = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetAuthKey(v string) *SaveApgroupPortalConfigRequest {
	s.AuthKey = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetAuthSecret(v string) *SaveApgroupPortalConfigRequest {
	s.AuthSecret = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetCheckUrl(v string) *SaveApgroupPortalConfigRequest {
	s.CheckUrl = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetClientDownload(v int32) *SaveApgroupPortalConfigRequest {
	s.ClientDownload = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetClientUpload(v int32) *SaveApgroupPortalConfigRequest {
	s.ClientUpload = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetCountdown(v int32) *SaveApgroupPortalConfigRequest {
	s.Countdown = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetNetwork(v int32) *SaveApgroupPortalConfigRequest {
	s.Network = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetPortalTypes(v []*string) *SaveApgroupPortalConfigRequest {
	s.PortalTypes = v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetPortalUrl(v string) *SaveApgroupPortalConfigRequest {
	s.PortalUrl = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetTimeStamp(v int64) *SaveApgroupPortalConfigRequest {
	s.TimeStamp = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetTotalDownload(v int32) *SaveApgroupPortalConfigRequest {
	s.TotalDownload = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetTotalUpload(v int32) *SaveApgroupPortalConfigRequest {
	s.TotalUpload = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetWebAuthUrl(v string) *SaveApgroupPortalConfigRequest {
	s.WebAuthUrl = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) SetWhitelist(v string) *SaveApgroupPortalConfigRequest {
	s.Whitelist = &v
	return s
}

func (s *SaveApgroupPortalConfigRequest) Validate() error {
	return dara.Validate(s)
}
