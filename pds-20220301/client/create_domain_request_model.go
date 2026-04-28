// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDomainRequest
	GetDescription() *string
	SetDomainName(v string) *CreateDomainRequest
	GetDomainName() *string
	SetInitDriveEnable(v bool) *CreateDomainRequest
	GetInitDriveEnable() *bool
	SetInitDriveSize(v int64) *CreateDomainRequest
	GetInitDriveSize() *int64
	SetParentDomainId(v string) *CreateDomainRequest
	GetParentDomainId() *string
	SetSizeQuota(v int64) *CreateDomainRequest
	GetSizeQuota() *int64
	SetStoreRedundancyType(v string) *CreateDomainRequest
	GetStoreRedundancyType() *string
	SetUserCountQuota(v int64) *CreateDomainRequest
	GetUserCountQuota() *int64
}

type CreateDomainRequest struct {
	// The description of the domain.
	//
	// example:
	//
	// cloud drive dev
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud drive
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	// Specifies whether to enable the default drive feature. A value of true specifies that all users are assigned a drive by default on the first logon. Default value: false.
	//
	// example:
	//
	// true
	InitDriveEnable *bool `json:"init_drive_enable,omitempty" xml:"init_drive_enable,omitempty"`
	// This parameter is required when the init_drive_enable is set to true. The size of the default drive. Unit: bytes. The default is 0, meaning the created drive size is 0, and files cannot be uploaded. If you need to initialize the drive, set this value. A value of -1 indicates that the size is unlimited.
	//
	// example:
	//
	// 1073741824
	InitDriveSize *int64 `json:"init_drive_size,omitempty" xml:"init_drive_size,omitempty"`
	// The ID of the parent domain. If you want to create a child domain, specify parent_domain_id. In most cases, you do not need to create a child domain. If you want to perform secondary operations based on Drive and Photo Service, contact the customer service.
	//
	// example:
	//
	// bj1
	ParentDomainId *string `json:"parent_domain_id,omitempty" xml:"parent_domain_id,omitempty"`
	// The total storage quota for all drives in the domain. A value of 0 indicates that the quota is unlimited.
	//
	// example:
	//
	// 1099511627776
	SizeQuota *int64 `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	// Specifies the storage redundancy type. Valid values:
	//
	// 	- LRS: locally redundant storage
	//
	// 	- ZRS: zone-redundant storage
	//
	// example:
	//
	// LRS
	StoreRedundancyType *string `json:"store_redundancy_type,omitempty" xml:"store_redundancy_type,omitempty"`
	// The largest number of users that can be created in the domain. A value of 0 specifies that the number is unlimited.
	//
	// example:
	//
	// 50
	UserCountQuota *int64 `json:"user_count_quota,omitempty" xml:"user_count_quota,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateDomainRequest) GetInitDriveEnable() *bool {
	return s.InitDriveEnable
}

func (s *CreateDomainRequest) GetInitDriveSize() *int64 {
	return s.InitDriveSize
}

func (s *CreateDomainRequest) GetParentDomainId() *string {
	return s.ParentDomainId
}

func (s *CreateDomainRequest) GetSizeQuota() *int64 {
	return s.SizeQuota
}

func (s *CreateDomainRequest) GetStoreRedundancyType() *string {
	return s.StoreRedundancyType
}

func (s *CreateDomainRequest) GetUserCountQuota() *int64 {
	return s.UserCountQuota
}

func (s *CreateDomainRequest) SetDescription(v string) *CreateDomainRequest {
	s.Description = &v
	return s
}

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDomainRequest) SetInitDriveEnable(v bool) *CreateDomainRequest {
	s.InitDriveEnable = &v
	return s
}

func (s *CreateDomainRequest) SetInitDriveSize(v int64) *CreateDomainRequest {
	s.InitDriveSize = &v
	return s
}

func (s *CreateDomainRequest) SetParentDomainId(v string) *CreateDomainRequest {
	s.ParentDomainId = &v
	return s
}

func (s *CreateDomainRequest) SetSizeQuota(v int64) *CreateDomainRequest {
	s.SizeQuota = &v
	return s
}

func (s *CreateDomainRequest) SetStoreRedundancyType(v string) *CreateDomainRequest {
	s.StoreRedundancyType = &v
	return s
}

func (s *CreateDomainRequest) SetUserCountQuota(v int64) *CreateDomainRequest {
	s.UserCountQuota = &v
	return s
}

func (s *CreateDomainRequest) Validate() error {
	return dara.Validate(s)
}
