// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetCheckChannelRequest
	GetAppId() *string
	SetOwnerId(v int64) *GetCheckChannelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetCheckChannelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCheckChannelRequest
	GetResourceOwnerId() *int64
}

type GetCheckChannelRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCheckChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckChannelRequest) GoString() string {
	return s.String()
}

func (s *GetCheckChannelRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetCheckChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCheckChannelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCheckChannelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCheckChannelRequest) SetAppId(v string) *GetCheckChannelRequest {
	s.AppId = &v
	return s
}

func (s *GetCheckChannelRequest) SetOwnerId(v int64) *GetCheckChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCheckChannelRequest) SetResourceOwnerAccount(v string) *GetCheckChannelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCheckChannelRequest) SetResourceOwnerId(v int64) *GetCheckChannelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCheckChannelRequest) Validate() error {
	return dara.Validate(s)
}
