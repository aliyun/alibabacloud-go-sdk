// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *KillProcessRequest
	GetDBClusterId() *string
	SetInitialQueryId(v string) *KillProcessRequest
	GetInitialQueryId() *string
	SetOwnerAccount(v string) *KillProcessRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *KillProcessRequest
	GetOwnerId() *int64
	SetRegionId(v string) *KillProcessRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *KillProcessRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *KillProcessRequest
	GetResourceOwnerId() *int64
}

type KillProcessRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The query statement or query statements that you want to stop executing. If you want to stop executing multiple query statements, separate the statements with commas (,).
	//
	// >  If you do not set this parameter, all query statements are stopped by default.
	//
	// example:
	//
	// SELECT 	- FROM `test01` ;
	InitialQueryId *string `json:"InitialQueryId,omitempty" xml:"InitialQueryId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s KillProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s KillProcessRequest) GoString() string {
	return s.String()
}

func (s *KillProcessRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *KillProcessRequest) GetInitialQueryId() *string {
	return s.InitialQueryId
}

func (s *KillProcessRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *KillProcessRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *KillProcessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KillProcessRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *KillProcessRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *KillProcessRequest) SetDBClusterId(v string) *KillProcessRequest {
	s.DBClusterId = &v
	return s
}

func (s *KillProcessRequest) SetInitialQueryId(v string) *KillProcessRequest {
	s.InitialQueryId = &v
	return s
}

func (s *KillProcessRequest) SetOwnerAccount(v string) *KillProcessRequest {
	s.OwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetOwnerId(v int64) *KillProcessRequest {
	s.OwnerId = &v
	return s
}

func (s *KillProcessRequest) SetRegionId(v string) *KillProcessRequest {
	s.RegionId = &v
	return s
}

func (s *KillProcessRequest) SetResourceOwnerAccount(v string) *KillProcessRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *KillProcessRequest) SetResourceOwnerId(v int64) *KillProcessRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *KillProcessRequest) Validate() error {
	return dara.Validate(s)
}
