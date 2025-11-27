// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsResourceSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeRdsResourceSettingsRequest
	GetOwnerId() *int64
	SetResourceNiche(v string) *DescribeRdsResourceSettingsRequest
	GetResourceNiche() *string
	SetResourceOwnerAccount(v string) *DescribeRdsResourceSettingsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRdsResourceSettingsRequest
	GetResourceOwnerId() *int64
}

type DescribeRdsResourceSettingsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The location of the notification.
	//
	// 	- noticeBar: notification bar
	//
	// 	- popUp: popup
	//
	// This parameter is required.
	//
	// example:
	//
	// noticeBar
	ResourceNiche        *string `json:"ResourceNiche,omitempty" xml:"ResourceNiche,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRdsResourceSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsResourceSettingsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsResourceSettingsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRdsResourceSettingsRequest) GetResourceNiche() *string {
	return s.ResourceNiche
}

func (s *DescribeRdsResourceSettingsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRdsResourceSettingsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRdsResourceSettingsRequest) SetOwnerId(v int64) *DescribeRdsResourceSettingsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsResourceSettingsRequest) SetResourceNiche(v string) *DescribeRdsResourceSettingsRequest {
	s.ResourceNiche = &v
	return s
}

func (s *DescribeRdsResourceSettingsRequest) SetResourceOwnerAccount(v string) *DescribeRdsResourceSettingsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsResourceSettingsRequest) SetResourceOwnerId(v int64) *DescribeRdsResourceSettingsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsResourceSettingsRequest) Validate() error {
	return dara.Validate(s)
}
