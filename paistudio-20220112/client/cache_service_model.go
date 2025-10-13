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
	SetCreatedBy(v string) *CacheService
	GetCreatedBy() *string
	SetGmtCreated(v string) *CacheService
	GetGmtCreated() *string
	SetQuotaId(v string) *CacheService
	GetQuotaId() *string
	SetStatus(v string) *CacheService
	GetStatus() *string
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
	CacheInfos              []*CacheInfo `json:"CacheInfos,omitempty" xml:"CacheInfos,omitempty" type:"Repeated"`
	CacheServiceId          *string      `json:"CacheServiceId,omitempty" xml:"CacheServiceId,omitempty"`
	CreatedBy               *string      `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	GmtCreated              *string      `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	QuotaId                 *string      `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	Status                  *string      `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportedClientQuotaIds []*string    `json:"SupportedClientQuotaIds,omitempty" xml:"SupportedClientQuotaIds,omitempty" type:"Repeated"`
	TenantId                *string      `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserId                  *string      `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserVpc                 *UserVpc     `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
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

func (s *CacheService) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CacheService) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CacheService) GetQuotaId() *string {
	return s.QuotaId
}

func (s *CacheService) GetStatus() *string {
	return s.Status
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

func (s *CacheService) SetCreatedBy(v string) *CacheService {
	s.CreatedBy = &v
	return s
}

func (s *CacheService) SetGmtCreated(v string) *CacheService {
	s.GmtCreated = &v
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
