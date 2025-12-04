// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyDBConfigRequest
	GetConfig() *string
	SetDBClusterId(v string) *ModifyDBConfigRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBConfigRequest
	GetResourceOwnerId() *int64
}

type ModifyDBConfigRequest struct {
	// The dictionary configuration.
	//
	// example:
	//
	// {"name":"test"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1r59y779o04****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyDBConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBConfigRequest) SetConfig(v string) *ModifyDBConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyDBConfigRequest) SetDBClusterId(v string) *ModifyDBConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetOwnerAccount(v string) *ModifyDBConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBConfigRequest) SetOwnerId(v int64) *ModifyDBConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetRegionId(v string) *ModifyDBConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBConfigRequest) SetResourceOwnerId(v int64) *ModifyDBConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBConfigRequest) Validate() error {
	return dara.Validate(s)
}
