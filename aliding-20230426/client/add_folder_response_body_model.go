// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFolderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDentry(v *AddFolderResponseBodyDentry) *AddFolderResponseBody
	GetDentry() *AddFolderResponseBodyDentry
	SetRequestId(v string) *AddFolderResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *AddFolderResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *AddFolderResponseBody
	GetVendorType() *string
}

type AddFolderResponseBody struct {
	Dentry *AddFolderResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s AddFolderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFolderResponseBody) GoString() string {
	return s.String()
}

func (s *AddFolderResponseBody) GetDentry() *AddFolderResponseBodyDentry {
	return s.Dentry
}

func (s *AddFolderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFolderResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *AddFolderResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *AddFolderResponseBody) SetDentry(v *AddFolderResponseBodyDentry) *AddFolderResponseBody {
	s.Dentry = v
	return s
}

func (s *AddFolderResponseBody) SetRequestId(v string) *AddFolderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFolderResponseBody) SetVendorRequestId(v string) *AddFolderResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *AddFolderResponseBody) SetVendorType(v string) *AddFolderResponseBody {
	s.VendorType = &v
	return s
}

func (s *AddFolderResponseBody) Validate() error {
	if s.Dentry != nil {
		if err := s.Dentry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFolderResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"AppProperties,omitempty" xml:"AppProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// creator_id
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// dentry_id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// modifier_id
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// dentry_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 163201723391
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// example:
	//
	// ./test.txt
	Path       *string                                `json:"Path,omitempty" xml:"Path,omitempty"`
	Properties *AddFolderResponseBodyDentryProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// example:
	//
	// 6020771
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// Ao01nSzzBxZQ68JW
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string `json:"StorageDriver,omitempty" xml:"StorageDriver,omitempty"`
	// example:
	//
	// FOLDER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 80a7201602b34450a7a97d8d4e255421
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AddFolderResponseBodyDentry) String() string {
	return dara.Prettify(s)
}

func (s AddFolderResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *AddFolderResponseBodyDentry) GetAppProperties() map[string][]*DentryAppPropertiesValue {
	return s.AppProperties
}

func (s *AddFolderResponseBodyDentry) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AddFolderResponseBodyDentry) GetCreatorId() *string {
	return s.CreatorId
}

func (s *AddFolderResponseBodyDentry) GetExtension() *string {
	return s.Extension
}

func (s *AddFolderResponseBodyDentry) GetId() *string {
	return s.Id
}

func (s *AddFolderResponseBodyDentry) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *AddFolderResponseBodyDentry) GetModifierId() *string {
	return s.ModifierId
}

func (s *AddFolderResponseBodyDentry) GetName() *string {
	return s.Name
}

func (s *AddFolderResponseBodyDentry) GetParentId() *string {
	return s.ParentId
}

func (s *AddFolderResponseBodyDentry) GetPartitionType() *string {
	return s.PartitionType
}

func (s *AddFolderResponseBodyDentry) GetPath() *string {
	return s.Path
}

func (s *AddFolderResponseBodyDentry) GetProperties() *AddFolderResponseBodyDentryProperties {
	return s.Properties
}

func (s *AddFolderResponseBodyDentry) GetSize() *int64 {
	return s.Size
}

func (s *AddFolderResponseBodyDentry) GetSpaceId() *string {
	return s.SpaceId
}

func (s *AddFolderResponseBodyDentry) GetStatus() *string {
	return s.Status
}

func (s *AddFolderResponseBodyDentry) GetStorageDriver() *string {
	return s.StorageDriver
}

func (s *AddFolderResponseBodyDentry) GetType() *string {
	return s.Type
}

func (s *AddFolderResponseBodyDentry) GetUuid() *string {
	return s.Uuid
}

func (s *AddFolderResponseBodyDentry) GetVersion() *int64 {
	return s.Version
}

func (s *AddFolderResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *AddFolderResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *AddFolderResponseBodyDentry) SetCreateTime(v string) *AddFolderResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetCreatorId(v string) *AddFolderResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetExtension(v string) *AddFolderResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetId(v string) *AddFolderResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetModifiedTime(v string) *AddFolderResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetModifierId(v string) *AddFolderResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetName(v string) *AddFolderResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetParentId(v string) *AddFolderResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetPartitionType(v string) *AddFolderResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetPath(v string) *AddFolderResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetProperties(v *AddFolderResponseBodyDentryProperties) *AddFolderResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *AddFolderResponseBodyDentry) SetSize(v int64) *AddFolderResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetSpaceId(v string) *AddFolderResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetStatus(v string) *AddFolderResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetStorageDriver(v string) *AddFolderResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetType(v string) *AddFolderResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetUuid(v string) *AddFolderResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *AddFolderResponseBodyDentry) SetVersion(v int64) *AddFolderResponseBodyDentry {
	s.Version = &v
	return s
}

func (s *AddFolderResponseBodyDentry) Validate() error {
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFolderResponseBodyDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s AddFolderResponseBodyDentryProperties) String() string {
	return dara.Prettify(s)
}

func (s AddFolderResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *AddFolderResponseBodyDentryProperties) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *AddFolderResponseBodyDentryProperties) SetReadOnly(v bool) *AddFolderResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

func (s *AddFolderResponseBodyDentryProperties) Validate() error {
	return dara.Validate(s)
}
