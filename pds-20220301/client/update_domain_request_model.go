// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDomainRequest
	GetDescription() *string
	SetDomainId(v string) *UpdateDomainRequest
	GetDomainId() *string
	SetDomainName(v string) *UpdateDomainRequest
	GetDomainName() *string
	SetInitDriveEnable(v bool) *UpdateDomainRequest
	GetInitDriveEnable() *bool
	SetInitDriveSize(v int64) *UpdateDomainRequest
	GetInitDriveSize() *int64
	SetPublishedAppAccessStrategy(v *AppAccessStrategy) *UpdateDomainRequest
	GetPublishedAppAccessStrategy() *AppAccessStrategy
	SetSizeQuota(v int64) *UpdateDomainRequest
	GetSizeQuota() *int64
	SetUserCountQuota(v int64) *UpdateDomainRequest
	GetUserCountQuota() *int64
}

type UpdateDomainRequest struct {
	// The description of the domain.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The name of the domain.
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	// Specifies whether to enable the default drive feature. A value of true specifies that all users are assigned a drive by default on the first logon. Default value: false.
	//
	// example:
	//
	// true
	InitDriveEnable *bool `json:"init_drive_enable,omitempty" xml:"init_drive_enable,omitempty"`
	// The size of the default drive. Unit: bytes. You must specify init_drive_size if you set init_drive_enable to true. Default value: 0. A value of 0 specifies that the size of the default drive is 0 bytes and you cannot upload files to the drive. To initialize the default drive, set init_drive_size to 0. A value of -1 specifies that the size is unlimited.
	//
	// example:
	//
	// 1073741824
	InitDriveSize *int64 `json:"init_drive_size,omitempty" xml:"init_drive_size,omitempty"`
	// The access policy of the application.
	PublishedAppAccessStrategy *AppAccessStrategy `json:"published_app_access_strategy,omitempty" xml:"published_app_access_strategy,omitempty"`
	// The total storage quota for all drives in the domain. A value of 0 specifies that the quota is unlimited.
	//
	// example:
	//
	// 1099511627776
	SizeQuota *int64 `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	// The maximum number of users that can be created in the domain.
	//
	// example:
	//
	// 50
	UserCountQuota *int64 `json:"user_count_quota,omitempty" xml:"user_count_quota,omitempty"`
}

func (s UpdateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDomainRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *UpdateDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDomainRequest) GetInitDriveEnable() *bool {
	return s.InitDriveEnable
}

func (s *UpdateDomainRequest) GetInitDriveSize() *int64 {
	return s.InitDriveSize
}

func (s *UpdateDomainRequest) GetPublishedAppAccessStrategy() *AppAccessStrategy {
	return s.PublishedAppAccessStrategy
}

func (s *UpdateDomainRequest) GetSizeQuota() *int64 {
	return s.SizeQuota
}

func (s *UpdateDomainRequest) GetUserCountQuota() *int64 {
	return s.UserCountQuota
}

func (s *UpdateDomainRequest) SetDescription(v string) *UpdateDomainRequest {
	s.Description = &v
	return s
}

func (s *UpdateDomainRequest) SetDomainId(v string) *UpdateDomainRequest {
	s.DomainId = &v
	return s
}

func (s *UpdateDomainRequest) SetDomainName(v string) *UpdateDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDomainRequest) SetInitDriveEnable(v bool) *UpdateDomainRequest {
	s.InitDriveEnable = &v
	return s
}

func (s *UpdateDomainRequest) SetInitDriveSize(v int64) *UpdateDomainRequest {
	s.InitDriveSize = &v
	return s
}

func (s *UpdateDomainRequest) SetPublishedAppAccessStrategy(v *AppAccessStrategy) *UpdateDomainRequest {
	s.PublishedAppAccessStrategy = v
	return s
}

func (s *UpdateDomainRequest) SetSizeQuota(v int64) *UpdateDomainRequest {
	s.SizeQuota = &v
	return s
}

func (s *UpdateDomainRequest) SetUserCountQuota(v int64) *UpdateDomainRequest {
	s.UserCountQuota = &v
	return s
}

func (s *UpdateDomainRequest) Validate() error {
	if s.PublishedAppAccessStrategy != nil {
		if err := s.PublishedAppAccessStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
