// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCheckChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SetCheckChannelRequest
	GetAppId() *string
	SetChannel(v string) *SetCheckChannelRequest
	GetChannel() *string
	SetLegalSwitch(v string) *SetCheckChannelRequest
	GetLegalSwitch() *string
	SetOwnerId(v int64) *SetCheckChannelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SetCheckChannelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetCheckChannelRequest
	GetResourceOwnerId() *int64
}

type SetCheckChannelRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Channel              *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	LegalSwitch          *string `json:"LegalSwitch,omitempty" xml:"LegalSwitch,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetCheckChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCheckChannelRequest) GoString() string {
	return s.String()
}

func (s *SetCheckChannelRequest) GetAppId() *string {
	return s.AppId
}

func (s *SetCheckChannelRequest) GetChannel() *string {
	return s.Channel
}

func (s *SetCheckChannelRequest) GetLegalSwitch() *string {
	return s.LegalSwitch
}

func (s *SetCheckChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCheckChannelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetCheckChannelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetCheckChannelRequest) SetAppId(v string) *SetCheckChannelRequest {
	s.AppId = &v
	return s
}

func (s *SetCheckChannelRequest) SetChannel(v string) *SetCheckChannelRequest {
	s.Channel = &v
	return s
}

func (s *SetCheckChannelRequest) SetLegalSwitch(v string) *SetCheckChannelRequest {
	s.LegalSwitch = &v
	return s
}

func (s *SetCheckChannelRequest) SetOwnerId(v int64) *SetCheckChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCheckChannelRequest) SetResourceOwnerAccount(v string) *SetCheckChannelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetCheckChannelRequest) SetResourceOwnerId(v int64) *SetCheckChannelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetCheckChannelRequest) Validate() error {
	return dara.Validate(s)
}
