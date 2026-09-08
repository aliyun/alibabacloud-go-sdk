// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheService interface {
	dara.Model
	String() string
	GoString() string
	SetCacheInfos(v []*CacheInfo) *CacheService
	GetCacheInfos() []*CacheInfo
	SetCacheServiceId(v string) *CacheService
	GetCacheServiceId() *string
	SetClusterId(v string) *CacheService
	GetClusterId() *string
	SetCreatedBy(v string) *CacheService
	GetCreatedBy() *string
	SetGmtCreated(v string) *CacheService
	GetGmtCreated() *string
	SetIsSharded(v bool) *CacheService
	GetIsSharded() *bool
	SetNetworkType(v string) *CacheService
	GetNetworkType() *string
	SetQuotaId(v string) *CacheService
	GetQuotaId() *string
	SetStatus(v string) *CacheService
	GetStatus() *string
	SetSupportRDMA(v string) *CacheService
	GetSupportRDMA() *string
	SetSupportedClientQuotaIds(v []*string) *CacheService
	GetSupportedClientQuotaIds() []*string
	SetTenantId(v string) *CacheService
	GetTenantId() *string
	SetUserId(v string) *CacheService
	GetUserId() *string
	SetUserVpc(v *UserVpc) *CacheService
	GetUserVpc() *UserVpc
}

type CacheService struct {
	// The list of data source cache information connected to the cache service. Each element corresponds to a data source and its access port.
	CacheInfos []*CacheInfo `json:"CacheInfos,omitempty" xml:"CacheInfos,omitempty" type:"Repeated"`
	// The cache service ID. This is the unique identifier of the cache service.
	//
	// example:
	//
	// cachea1b2c3d4e5f
	CacheServiceId *string `json:"CacheServiceId,omitempty" xml:"CacheServiceId,omitempty"`
	// The ID of the cluster where the cache service resides.
	//
	// example:
	//
	// c1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the resource quota that created the cache service.
	//
	// example:
	//
	// quota1a2b3c4d5e6
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the cache service was created, in UTC in ISO 8601 format.
	//
	// example:
	//
	// 2026-08-10T03:17:31Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// Indicates whether the service discovery of the cache service uses shard mode.
	IsSharded *bool `json:"IsSharded,omitempty" xml:"IsSharded,omitempty"`
	// The type of RDMA network interface controller used by the cache service. This parameter is returned only when SupportRDMA is set to true. Valid values:
	//
	// - eic: EIC network interface controller.
	//
	// - mlx: Mellanox network interface controller.
	//
	// This parameter is empty when RDMA is not enabled.
	//
	// example:
	//
	// eic
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the resource quota to which the cache service belongs.
	//
	// example:
	//
	// quota1a2b3c4d5e6
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The current status of the cache service.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the cache service supports access through RDMA networks.
	SupportRDMA *string `json:"SupportRDMA,omitempty" xml:"SupportRDMA,omitempty"`
	// The list of client quota IDs that are allowed to access the cache service.
	SupportedClientQuotaIds []*string `json:"SupportedClientQuotaIds,omitempty" xml:"SupportedClientQuotaIds,omitempty" type:"Repeated"`
	// The tenant ID to which the cache service belongs, which is the Alibaba Cloud account ID.
	//
	// example:
	//
	// 1234567890123456
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The ID of the user who created the cache service.
	//
	// example:
	//
	// 123456789012345678
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The VPC network information of the cache service, including the VPC, vSwitch, and security group configurations.
	UserVpc *UserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
}

func (s CacheService) String() string {
	return dara.Prettify(s)
}

func (s CacheService) GoString() string {
	return s.String()
}

func (s *CacheService) GetCacheInfos() []*CacheInfo {
	return s.CacheInfos
}

func (s *CacheService) GetCacheServiceId() *string {
	return s.CacheServiceId
}

func (s *CacheService) GetClusterId() *string {
	return s.ClusterId
}

func (s *CacheService) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CacheService) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CacheService) GetIsSharded() *bool {
	return s.IsSharded
}

func (s *CacheService) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CacheService) GetQuotaId() *string {
	return s.QuotaId
}

func (s *CacheService) GetStatus() *string {
	return s.Status
}

func (s *CacheService) GetSupportRDMA() *string {
	return s.SupportRDMA
}

func (s *CacheService) GetSupportedClientQuotaIds() []*string {
	return s.SupportedClientQuotaIds
}

func (s *CacheService) GetTenantId() *string {
	return s.TenantId
}

func (s *CacheService) GetUserId() *string {
	return s.UserId
}

func (s *CacheService) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *CacheService) SetCacheInfos(v []*CacheInfo) *CacheService {
	s.CacheInfos = v
	return s
}

func (s *CacheService) SetCacheServiceId(v string) *CacheService {
	s.CacheServiceId = &v
	return s
}

func (s *CacheService) SetClusterId(v string) *CacheService {
	s.ClusterId = &v
	return s
}

func (s *CacheService) SetCreatedBy(v string) *CacheService {
	s.CreatedBy = &v
	return s
}

func (s *CacheService) SetGmtCreated(v string) *CacheService {
	s.GmtCreated = &v
	return s
}

func (s *CacheService) SetIsSharded(v bool) *CacheService {
	s.IsSharded = &v
	return s
}

func (s *CacheService) SetNetworkType(v string) *CacheService {
	s.NetworkType = &v
	return s
}

func (s *CacheService) SetQuotaId(v string) *CacheService {
	s.QuotaId = &v
	return s
}

func (s *CacheService) SetStatus(v string) *CacheService {
	s.Status = &v
	return s
}

func (s *CacheService) SetSupportRDMA(v string) *CacheService {
	s.SupportRDMA = &v
	return s
}

func (s *CacheService) SetSupportedClientQuotaIds(v []*string) *CacheService {
	s.SupportedClientQuotaIds = v
	return s
}

func (s *CacheService) SetTenantId(v string) *CacheService {
	s.TenantId = &v
	return s
}

func (s *CacheService) SetUserId(v string) *CacheService {
	s.UserId = &v
	return s
}

func (s *CacheService) SetUserVpc(v *UserVpc) *CacheService {
	s.UserVpc = v
	return s
}

func (s *CacheService) Validate() error {
	if s.CacheInfos != nil {
		for _, item := range s.CacheInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}
