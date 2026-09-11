// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAvailableVscsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArns(v []*ListKVCacheStoreAvailableVscsRequestArns) *ListKVCacheStoreAvailableVscsRequest
	GetArns() []*ListKVCacheStoreAvailableVscsRequestArns
	SetInstanceId(v string) *ListKVCacheStoreAvailableVscsRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ListKVCacheStoreAvailableVscsRequest
	GetInstanceType() *string
	SetKvcsId(v string) *ListKVCacheStoreAvailableVscsRequest
	GetKvcsId() *string
	SetRegionId(v string) *ListKVCacheStoreAvailableVscsRequest
	GetRegionId() *string
}

type ListKVCacheStoreAvailableVscsRequest struct {
	// The cross-account authorization role chain, used for cross-account VSC queries in ECS or EFLO scenarios.
	Arns []*ListKVCacheStoreAvailableVscsRequestArns `json:"Arns,omitempty" xml:"Arns,omitempty" type:"Repeated"`
	// The compute instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-55kl5wq6j6kvtl4xu5tgunddu
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The compute instance type. Valid values:
	//
	// - ECS
	//
	// - EFLO
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The KVCacheStore instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kvcs-cn-5y44vuqiz001
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The region ID, such as cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListKVCacheStoreAvailableVscsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableVscsRequest) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableVscsRequest) GetArns() []*ListKVCacheStoreAvailableVscsRequestArns {
	return s.Arns
}

func (s *ListKVCacheStoreAvailableVscsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListKVCacheStoreAvailableVscsRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListKVCacheStoreAvailableVscsRequest) GetKvcsId() *string {
	return s.KvcsId
}

func (s *ListKVCacheStoreAvailableVscsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKVCacheStoreAvailableVscsRequest) SetArns(v []*ListKVCacheStoreAvailableVscsRequestArns) *ListKVCacheStoreAvailableVscsRequest {
	s.Arns = v
	return s
}

func (s *ListKVCacheStoreAvailableVscsRequest) SetInstanceId(v string) *ListKVCacheStoreAvailableVscsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsRequest) SetInstanceType(v string) *ListKVCacheStoreAvailableVscsRequest {
	s.InstanceType = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsRequest) SetKvcsId(v string) *ListKVCacheStoreAvailableVscsRequest {
	s.KvcsId = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsRequest) SetRegionId(v string) *ListKVCacheStoreAvailableVscsRequest {
	s.RegionId = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsRequest) Validate() error {
	if s.Arns != nil {
		for _, item := range s.Arns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKVCacheStoreAvailableVscsRequestArns struct {
	// The target UID for role assumption.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1193522024229156
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::1028257687084022:role/zeus-locationservicerole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The role type. Valid values:
	//
	// - service
	//
	// - user
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListKVCacheStoreAvailableVscsRequestArns) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableVscsRequestArns) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableVscsRequestArns) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *ListKVCacheStoreAvailableVscsRequestArns) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListKVCacheStoreAvailableVscsRequestArns) GetRoleType() *string {
	return s.RoleType
}

func (s *ListKVCacheStoreAvailableVscsRequestArns) SetAssumeRoleFor(v string) *ListKVCacheStoreAvailableVscsRequestArns {
	s.AssumeRoleFor = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsRequestArns) SetRoleArn(v string) *ListKVCacheStoreAvailableVscsRequestArns {
	s.RoleArn = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsRequestArns) SetRoleType(v string) *ListKVCacheStoreAvailableVscsRequestArns {
	s.RoleType = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsRequestArns) Validate() error {
	return dara.Validate(s)
}
