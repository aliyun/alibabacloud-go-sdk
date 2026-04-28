// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *BaseDomainResponse
	GetCreatedAt() *string
	SetDescription(v string) *BaseDomainResponse
	GetDescription() *string
	SetDomainId(v string) *BaseDomainResponse
	GetDomainId() *string
	SetDomainName(v string) *BaseDomainResponse
	GetDomainName() *string
	SetInitDriveEnable(v bool) *BaseDomainResponse
	GetInitDriveEnable() *bool
	SetInitDriveSize(v int64) *BaseDomainResponse
	GetInitDriveSize() *int64
	SetParentDomainId(v string) *BaseDomainResponse
	GetParentDomainId() *string
	SetPublishedAppAccessStrategy(v *AppAccessStrategy) *BaseDomainResponse
	GetPublishedAppAccessStrategy() *AppAccessStrategy
	SetShareLinkEnabled(v bool) *BaseDomainResponse
	GetShareLinkEnabled() *bool
	SetSizeQuota(v int64) *BaseDomainResponse
	GetSizeQuota() *int64
	SetSizeQuotaUsed(v int64) *BaseDomainResponse
	GetSizeQuotaUsed() *int64
	SetStatus(v int64) *BaseDomainResponse
	GetStatus() *int64
	SetUpdatedAt(v string) *BaseDomainResponse
	GetUpdatedAt() *string
	SetUsedSize(v int64) *BaseDomainResponse
	GetUsedSize() *int64
}

type BaseDomainResponse struct {
	CreatedAt                  *string            `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description                *string            `json:"description,omitempty" xml:"description,omitempty"`
	DomainId                   *string            `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DomainName                 *string            `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	InitDriveEnable            *bool              `json:"init_drive_enable,omitempty" xml:"init_drive_enable,omitempty"`
	InitDriveSize              *int64             `json:"init_drive_size,omitempty" xml:"init_drive_size,omitempty"`
	ParentDomainId             *string            `json:"parent_domain_id,omitempty" xml:"parent_domain_id,omitempty"`
	PublishedAppAccessStrategy *AppAccessStrategy `json:"published_app_access_strategy,omitempty" xml:"published_app_access_strategy,omitempty"`
	ShareLinkEnabled           *bool              `json:"share_link_enabled,omitempty" xml:"share_link_enabled,omitempty"`
	SizeQuota                  *int64             `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	SizeQuotaUsed              *int64             `json:"size_quota_used,omitempty" xml:"size_quota_used,omitempty"`
	Status                     *int64             `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt                  *string            `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UsedSize                   *int64             `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s BaseDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseDomainResponse) GoString() string {
	return s.String()
}

func (s *BaseDomainResponse) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *BaseDomainResponse) GetDescription() *string {
	return s.Description
}

func (s *BaseDomainResponse) GetDomainId() *string {
	return s.DomainId
}

func (s *BaseDomainResponse) GetDomainName() *string {
	return s.DomainName
}

func (s *BaseDomainResponse) GetInitDriveEnable() *bool {
	return s.InitDriveEnable
}

func (s *BaseDomainResponse) GetInitDriveSize() *int64 {
	return s.InitDriveSize
}

func (s *BaseDomainResponse) GetParentDomainId() *string {
	return s.ParentDomainId
}

func (s *BaseDomainResponse) GetPublishedAppAccessStrategy() *AppAccessStrategy {
	return s.PublishedAppAccessStrategy
}

func (s *BaseDomainResponse) GetShareLinkEnabled() *bool {
	return s.ShareLinkEnabled
}

func (s *BaseDomainResponse) GetSizeQuota() *int64 {
	return s.SizeQuota
}

func (s *BaseDomainResponse) GetSizeQuotaUsed() *int64 {
	return s.SizeQuotaUsed
}

func (s *BaseDomainResponse) GetStatus() *int64 {
	return s.Status
}

func (s *BaseDomainResponse) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *BaseDomainResponse) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *BaseDomainResponse) SetCreatedAt(v string) *BaseDomainResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseDomainResponse) SetDescription(v string) *BaseDomainResponse {
	s.Description = &v
	return s
}

func (s *BaseDomainResponse) SetDomainId(v string) *BaseDomainResponse {
	s.DomainId = &v
	return s
}

func (s *BaseDomainResponse) SetDomainName(v string) *BaseDomainResponse {
	s.DomainName = &v
	return s
}

func (s *BaseDomainResponse) SetInitDriveEnable(v bool) *BaseDomainResponse {
	s.InitDriveEnable = &v
	return s
}

func (s *BaseDomainResponse) SetInitDriveSize(v int64) *BaseDomainResponse {
	s.InitDriveSize = &v
	return s
}

func (s *BaseDomainResponse) SetParentDomainId(v string) *BaseDomainResponse {
	s.ParentDomainId = &v
	return s
}

func (s *BaseDomainResponse) SetPublishedAppAccessStrategy(v *AppAccessStrategy) *BaseDomainResponse {
	s.PublishedAppAccessStrategy = v
	return s
}

func (s *BaseDomainResponse) SetShareLinkEnabled(v bool) *BaseDomainResponse {
	s.ShareLinkEnabled = &v
	return s
}

func (s *BaseDomainResponse) SetSizeQuota(v int64) *BaseDomainResponse {
	s.SizeQuota = &v
	return s
}

func (s *BaseDomainResponse) SetSizeQuotaUsed(v int64) *BaseDomainResponse {
	s.SizeQuotaUsed = &v
	return s
}

func (s *BaseDomainResponse) SetStatus(v int64) *BaseDomainResponse {
	s.Status = &v
	return s
}

func (s *BaseDomainResponse) SetUpdatedAt(v string) *BaseDomainResponse {
	s.UpdatedAt = &v
	return s
}

func (s *BaseDomainResponse) SetUsedSize(v int64) *BaseDomainResponse {
	s.UsedSize = &v
	return s
}

func (s *BaseDomainResponse) Validate() error {
	if s.PublishedAppAccessStrategy != nil {
		if err := s.PublishedAppAccessStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
