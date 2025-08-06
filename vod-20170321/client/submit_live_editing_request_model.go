// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveEditingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SubmitLiveEditingRequest
	GetAppName() *string
	SetClips(v string) *SubmitLiveEditingRequest
	GetClips() *string
	SetCoverURL(v string) *SubmitLiveEditingRequest
	GetCoverURL() *string
	SetDescription(v string) *SubmitLiveEditingRequest
	GetDescription() *string
	SetDomainName(v string) *SubmitLiveEditingRequest
	GetDomainName() *string
	SetMediaMetadata(v string) *SubmitLiveEditingRequest
	GetMediaMetadata() *string
	SetOwnerId(v int64) *SubmitLiveEditingRequest
	GetOwnerId() *int64
	SetProduceConfig(v string) *SubmitLiveEditingRequest
	GetProduceConfig() *string
	SetResourceOwnerAccount(v string) *SubmitLiveEditingRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitLiveEditingRequest
	GetResourceOwnerId() *int64
	SetStreamName(v string) *SubmitLiveEditingRequest
	GetStreamName() *string
	SetTitle(v string) *SubmitLiveEditingRequest
	GetTitle() *string
	SetUserData(v string) *SubmitLiveEditingRequest
	GetUserData() *string
}

type SubmitLiveEditingRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Clips       *string `json:"Clips,omitempty" xml:"Clips,omitempty"`
	CoverURL    *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	MediaMetadata        *string `json:"MediaMetadata,omitempty" xml:"MediaMetadata,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProduceConfig        *string `json:"ProduceConfig,omitempty" xml:"ProduceConfig,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserData   *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitLiveEditingRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveEditingRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingRequest) GetAppName() *string {
	return s.AppName
}

func (s *SubmitLiveEditingRequest) GetClips() *string {
	return s.Clips
}

func (s *SubmitLiveEditingRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *SubmitLiveEditingRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitLiveEditingRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SubmitLiveEditingRequest) GetMediaMetadata() *string {
	return s.MediaMetadata
}

func (s *SubmitLiveEditingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitLiveEditingRequest) GetProduceConfig() *string {
	return s.ProduceConfig
}

func (s *SubmitLiveEditingRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitLiveEditingRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitLiveEditingRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *SubmitLiveEditingRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitLiveEditingRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitLiveEditingRequest) SetAppName(v string) *SubmitLiveEditingRequest {
	s.AppName = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetClips(v string) *SubmitLiveEditingRequest {
	s.Clips = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetCoverURL(v string) *SubmitLiveEditingRequest {
	s.CoverURL = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetDescription(v string) *SubmitLiveEditingRequest {
	s.Description = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetDomainName(v string) *SubmitLiveEditingRequest {
	s.DomainName = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetMediaMetadata(v string) *SubmitLiveEditingRequest {
	s.MediaMetadata = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetOwnerId(v int64) *SubmitLiveEditingRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetProduceConfig(v string) *SubmitLiveEditingRequest {
	s.ProduceConfig = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetResourceOwnerAccount(v string) *SubmitLiveEditingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetResourceOwnerId(v int64) *SubmitLiveEditingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetStreamName(v string) *SubmitLiveEditingRequest {
	s.StreamName = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetTitle(v string) *SubmitLiveEditingRequest {
	s.Title = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetUserData(v string) *SubmitLiveEditingRequest {
	s.UserData = &v
	return s
}

func (s *SubmitLiveEditingRequest) Validate() error {
	return dara.Validate(s)
}
