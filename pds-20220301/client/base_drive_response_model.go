// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseDriveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetActionList(v []*string) *BaseDriveResponse
	GetActionList() []*string
	SetCategory(v string) *BaseDriveResponse
	GetCategory() *string
	SetCreatedAt(v string) *BaseDriveResponse
	GetCreatedAt() *string
	SetCreator(v string) *BaseDriveResponse
	GetCreator() *string
	SetDeltaEnabled(v bool) *BaseDriveResponse
	GetDeltaEnabled() *bool
	SetDescription(v string) *BaseDriveResponse
	GetDescription() *string
	SetDomainId(v string) *BaseDriveResponse
	GetDomainId() *string
	SetDriveId(v string) *BaseDriveResponse
	GetDriveId() *string
	SetDriveName(v string) *BaseDriveResponse
	GetDriveName() *string
	SetDriveType(v string) *BaseDriveResponse
	GetDriveType() *string
	SetEncryptDataAccess(v bool) *BaseDriveResponse
	GetEncryptDataAccess() *bool
	SetEncryptMode(v string) *BaseDriveResponse
	GetEncryptMode() *string
	SetIsHandover(v bool) *BaseDriveResponse
	GetIsHandover() *bool
	SetOwner(v string) *BaseDriveResponse
	GetOwner() *string
	SetOwnerType(v string) *BaseDriveResponse
	GetOwnerType() *string
	SetPathStatus(v string) *BaseDriveResponse
	GetPathStatus() *string
	SetPermission(v map[string]*IDPermission) *BaseDriveResponse
	GetPermission() map[string]*IDPermission
	SetRelativePath(v string) *BaseDriveResponse
	GetRelativePath() *string
	SetStatus(v string) *BaseDriveResponse
	GetStatus() *string
	SetStoreId(v string) *BaseDriveResponse
	GetStoreId() *string
	SetTotalSize(v int64) *BaseDriveResponse
	GetTotalSize() *int64
	SetUpdatedAt(v string) *BaseDriveResponse
	GetUpdatedAt() *string
	SetUsedSize(v int64) *BaseDriveResponse
	GetUsedSize() *int64
}

type BaseDriveResponse struct {
	ActionList []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	Category   *string   `json:"category,omitempty" xml:"category,omitempty"`
	CreatedAt  *string   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// ccp
	Creator      *string `json:"creator,omitempty" xml:"creator,omitempty"`
	DeltaEnabled *bool   `json:"delta_enabled,omitempty" xml:"delta_enabled,omitempty"`
	// example:
	//
	// ccp team drive
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// hz999
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// example:
	//
	// 123
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// ccpdrive
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// example:
	//
	// normal
	DriveType         *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	EncryptDataAccess *bool   `json:"encrypt_data_access,omitempty" xml:"encrypt_data_access,omitempty"`
	EncryptMode       *string `json:"encrypt_mode,omitempty" xml:"encrypt_mode,omitempty"`
	IsHandover        *bool   `json:"is_handover,omitempty" xml:"is_handover,omitempty"`
	// example:
	//
	// ccp
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// user
	OwnerType    *string                  `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	PathStatus   *string                  `json:"path_status,omitempty" xml:"path_status,omitempty"`
	Permission   map[string]*IDPermission `json:"permission,omitempty" xml:"permission,omitempty"`
	RelativePath *string                  `json:"relative_path,omitempty" xml:"relative_path,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 123
	StoreId *string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// example:
	//
	// 102400
	TotalSize *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// example:
	//
	// 1024
	UsedSize *int64 `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s BaseDriveResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseDriveResponse) GoString() string {
	return s.String()
}

func (s *BaseDriveResponse) GetActionList() []*string {
	return s.ActionList
}

func (s *BaseDriveResponse) GetCategory() *string {
	return s.Category
}

func (s *BaseDriveResponse) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *BaseDriveResponse) GetCreator() *string {
	return s.Creator
}

func (s *BaseDriveResponse) GetDeltaEnabled() *bool {
	return s.DeltaEnabled
}

func (s *BaseDriveResponse) GetDescription() *string {
	return s.Description
}

func (s *BaseDriveResponse) GetDomainId() *string {
	return s.DomainId
}

func (s *BaseDriveResponse) GetDriveId() *string {
	return s.DriveId
}

func (s *BaseDriveResponse) GetDriveName() *string {
	return s.DriveName
}

func (s *BaseDriveResponse) GetDriveType() *string {
	return s.DriveType
}

func (s *BaseDriveResponse) GetEncryptDataAccess() *bool {
	return s.EncryptDataAccess
}

func (s *BaseDriveResponse) GetEncryptMode() *string {
	return s.EncryptMode
}

func (s *BaseDriveResponse) GetIsHandover() *bool {
	return s.IsHandover
}

func (s *BaseDriveResponse) GetOwner() *string {
	return s.Owner
}

func (s *BaseDriveResponse) GetOwnerType() *string {
	return s.OwnerType
}

func (s *BaseDriveResponse) GetPathStatus() *string {
	return s.PathStatus
}

func (s *BaseDriveResponse) GetPermission() map[string]*IDPermission {
	return s.Permission
}

func (s *BaseDriveResponse) GetRelativePath() *string {
	return s.RelativePath
}

func (s *BaseDriveResponse) GetStatus() *string {
	return s.Status
}

func (s *BaseDriveResponse) GetStoreId() *string {
	return s.StoreId
}

func (s *BaseDriveResponse) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *BaseDriveResponse) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *BaseDriveResponse) GetUsedSize() *int64 {
	return s.UsedSize
}

func (s *BaseDriveResponse) SetActionList(v []*string) *BaseDriveResponse {
	s.ActionList = v
	return s
}

func (s *BaseDriveResponse) SetCategory(v string) *BaseDriveResponse {
	s.Category = &v
	return s
}

func (s *BaseDriveResponse) SetCreatedAt(v string) *BaseDriveResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseDriveResponse) SetCreator(v string) *BaseDriveResponse {
	s.Creator = &v
	return s
}

func (s *BaseDriveResponse) SetDeltaEnabled(v bool) *BaseDriveResponse {
	s.DeltaEnabled = &v
	return s
}

func (s *BaseDriveResponse) SetDescription(v string) *BaseDriveResponse {
	s.Description = &v
	return s
}

func (s *BaseDriveResponse) SetDomainId(v string) *BaseDriveResponse {
	s.DomainId = &v
	return s
}

func (s *BaseDriveResponse) SetDriveId(v string) *BaseDriveResponse {
	s.DriveId = &v
	return s
}

func (s *BaseDriveResponse) SetDriveName(v string) *BaseDriveResponse {
	s.DriveName = &v
	return s
}

func (s *BaseDriveResponse) SetDriveType(v string) *BaseDriveResponse {
	s.DriveType = &v
	return s
}

func (s *BaseDriveResponse) SetEncryptDataAccess(v bool) *BaseDriveResponse {
	s.EncryptDataAccess = &v
	return s
}

func (s *BaseDriveResponse) SetEncryptMode(v string) *BaseDriveResponse {
	s.EncryptMode = &v
	return s
}

func (s *BaseDriveResponse) SetIsHandover(v bool) *BaseDriveResponse {
	s.IsHandover = &v
	return s
}

func (s *BaseDriveResponse) SetOwner(v string) *BaseDriveResponse {
	s.Owner = &v
	return s
}

func (s *BaseDriveResponse) SetOwnerType(v string) *BaseDriveResponse {
	s.OwnerType = &v
	return s
}

func (s *BaseDriveResponse) SetPathStatus(v string) *BaseDriveResponse {
	s.PathStatus = &v
	return s
}

func (s *BaseDriveResponse) SetPermission(v map[string]*IDPermission) *BaseDriveResponse {
	s.Permission = v
	return s
}

func (s *BaseDriveResponse) SetRelativePath(v string) *BaseDriveResponse {
	s.RelativePath = &v
	return s
}

func (s *BaseDriveResponse) SetStatus(v string) *BaseDriveResponse {
	s.Status = &v
	return s
}

func (s *BaseDriveResponse) SetStoreId(v string) *BaseDriveResponse {
	s.StoreId = &v
	return s
}

func (s *BaseDriveResponse) SetTotalSize(v int64) *BaseDriveResponse {
	s.TotalSize = &v
	return s
}

func (s *BaseDriveResponse) SetUpdatedAt(v string) *BaseDriveResponse {
	s.UpdatedAt = &v
	return s
}

func (s *BaseDriveResponse) SetUsedSize(v int64) *BaseDriveResponse {
	s.UsedSize = &v
	return s
}

func (s *BaseDriveResponse) Validate() error {
	return dara.Validate(s)
}
