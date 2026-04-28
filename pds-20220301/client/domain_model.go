// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomain interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *Domain
	GetCreatedAt() *string
	SetDataHashName(v string) *Domain
	GetDataHashName() *string
	SetDescription(v string) *Domain
	GetDescription() *string
	SetDomainId(v string) *Domain
	GetDomainId() *string
	SetDomainName(v string) *Domain
	GetDomainName() *string
	SetInitDriveEnable(v bool) *Domain
	GetInitDriveEnable() *bool
	SetInitDriveSize(v int64) *Domain
	GetInitDriveSize() *int64
	SetParentDomainId(v string) *Domain
	GetParentDomainId() *string
	SetPublishedAppAccessStrategy(v *AppAccessStrategy) *Domain
	GetPublishedAppAccessStrategy() *AppAccessStrategy
	SetSharable(v bool) *Domain
	GetSharable() *bool
	SetSizeQuota(v int64) *Domain
	GetSizeQuota() *int64
	SetSizeQuotaUsed(v int64) *Domain
	GetSizeQuotaUsed() *int64
	SetStatus(v int64) *Domain
	GetStatus() *int64
	SetStoreRedundancyType(v string) *Domain
	GetStoreRedundancyType() *string
	SetUpdatedAt(v string) *Domain
	GetUpdatedAt() *string
	SetUsedSize(v int64) *Domain
	GetUsedSize() *int64
	SetUserCountQuota(v int64) *Domain
	GetUserCountQuota() *int64
}

type Domain struct {
	// The time when the domain was created. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The method used to calculate the hash value of the data.
	//
	// example:
	//
	// sha1
	DataHashName *string `json:"data_hash_name,omitempty" xml:"data_hash_name,omitempty"`
	// The description of the domain.
	//
	// example:
	//
	// my domain
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// domain ID
	//
	// example:
	//
	// bj2
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The name of the domain.
	//
	// example:
	//
	// pdsdomain
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	// Specifies whether to enable the default drive feature. Valid values: true and false. A value of true specifies that all users are assigned a drive by default on the first logon. Default value: false.
	//
	// example:
	//
	// true
	InitDriveEnable *bool `json:"init_drive_enable,omitempty" xml:"init_drive_enable,omitempty"`
	// The size of the default drive. Unit: bytes. This parameter is required if you set init_drive_enable to true. Default value: 0. A value of 0 indicates that the size of the default drive is 0 byte and you cannot upload files to the drive. To initialize the default drive, set init_drive_size to a positive number or -1. A value of -1 indicates that the size is unlimited.
	//
	// example:
	//
	// 1073741824
	InitDriveSize *int64 `json:"init_drive_size,omitempty" xml:"init_drive_size,omitempty"`
	// The ID of the parent domain. If the parent domain exists, the current domain is a child domain. Otherwise, the current domain is a common domain.
	//
	// example:
	//
	// bj1
	ParentDomainId *string `json:"parent_domain_id,omitempty" xml:"parent_domain_id,omitempty"`
	// The access policy of the application.
	PublishedAppAccessStrategy *AppAccessStrategy `json:"published_app_access_strategy,omitempty" xml:"published_app_access_strategy,omitempty"`
	// Specifies whether to enable sharing.
	Sharable *bool `json:"sharable,omitempty" xml:"sharable,omitempty"`
	// The total storage quota for all drives in the domain. A value of 0 indicates that the quota is unlimited.
	//
	// example:
	//
	// 1099511627776
	SizeQuota *int64 `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	// The used storage quota of all drives in the domain.
	//
	// example:
	//
	// 1099511627776
	SizeQuotaUsed *int64 `json:"size_quota_used,omitempty" xml:"size_quota_used,omitempty"`
	// The status of the domain. 1: The domain runs normally. 2: The domain is being created. 6: The domain has expired.
	//
	// example:
	//
	// 1
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// LRS
	StoreRedundancyType *string `json:"store_redundancy_type,omitempty" xml:"store_redundancy_type,omitempty"`
	// The time when the domain was last modified. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The usage of the logic space. Unit: bytes.
	//
	// example:
	//
	// 1099511627776
	UsedSize *int64 `json:"used_size,omitempty" xml:"used_size,omitempty"`
	// The maximum allowed number of users.
	//
	// example:
	//
	// 50
	UserCountQuota *int64 `json:"user_count_quota,omitempty" xml:"user_count_quota,omitempty"`
}

func (s Domain) String() string {
	return dara.Prettify(s)
}

func (s Domain) GoString() string {
	return s.String()
}

func (s *Domain) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Domain) GetDataHashName() *string {
	return s.DataHashName
}

func (s *Domain) GetDescription() *string {
	return s.Description
}

func (s *Domain) GetDomainId() *string {
	return s.DomainId
}

func (s *Domain) GetDomainName() *string {
	return s.DomainName
}

func (s *Domain) GetInitDriveEnable() *bool {
	return s.InitDriveEnable
}

func (s *Domain) GetInitDriveSize() *int64 {
	return s.InitDriveSize
}

func (s *Domain) GetParentDomainId() *string {
	return s.ParentDomainId
}

func (s *Domain) GetPublishedAppAccessStrategy() *AppAccessStrategy {
	return s.PublishedAppAccessStrategy
}

func (s *Domain) GetSharable() *bool {
	return s.Sharable
}

func (s *Domain) GetSizeQuota() *int64 {
	return s.SizeQuota
}

func (s *Domain) GetSizeQuotaUsed() *int64 {
	return s.SizeQuotaUsed
}

func (s *Domain) GetStatus() *int64 {
	return s.Status
}

func (s *Domain) GetStoreRedundancyType() *string {
	return s.StoreRedundancyType
}

func (s *Domain) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Domain) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *Domain) GetUserCountQuota() *int64 {
	return s.UserCountQuota
}

func (s *Domain) SetCreatedAt(v string) *Domain {
	s.CreatedAt = &v
	return s
}

func (s *Domain) SetDataHashName(v string) *Domain {
	s.DataHashName = &v
	return s
}

func (s *Domain) SetDescription(v string) *Domain {
	s.Description = &v
	return s
}

func (s *Domain) SetDomainId(v string) *Domain {
	s.DomainId = &v
	return s
}

func (s *Domain) SetDomainName(v string) *Domain {
	s.DomainName = &v
	return s
}

func (s *Domain) SetInitDriveEnable(v bool) *Domain {
	s.InitDriveEnable = &v
	return s
}

func (s *Domain) SetInitDriveSize(v int64) *Domain {
	s.InitDriveSize = &v
	return s
}

func (s *Domain) SetParentDomainId(v string) *Domain {
	s.ParentDomainId = &v
	return s
}

func (s *Domain) SetPublishedAppAccessStrategy(v *AppAccessStrategy) *Domain {
	s.PublishedAppAccessStrategy = v
	return s
}

func (s *Domain) SetSharable(v bool) *Domain {
	s.Sharable = &v
	return s
}

func (s *Domain) SetSizeQuota(v int64) *Domain {
	s.SizeQuota = &v
	return s
}

func (s *Domain) SetSizeQuotaUsed(v int64) *Domain {
	s.SizeQuotaUsed = &v
	return s
}

func (s *Domain) SetStatus(v int64) *Domain {
	s.Status = &v
	return s
}

func (s *Domain) SetStoreRedundancyType(v string) *Domain {
	s.StoreRedundancyType = &v
	return s
}

func (s *Domain) SetUpdatedAt(v string) *Domain {
	s.UpdatedAt = &v
	return s
}

func (s *Domain) SetUsedSize(v int64) *Domain {
	s.UsedSize = &v
	return s
}

func (s *Domain) SetUserCountQuota(v int64) *Domain {
	s.UserCountQuota = &v
	return s
}

func (s *Domain) Validate() error {
	if s.PublishedAppAccessStrategy != nil {
		if err := s.PublishedAppAccessStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
