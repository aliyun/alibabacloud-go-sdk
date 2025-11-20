// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDentry(v *CommitFileResponseBodyDentry) *CommitFileResponseBody
	GetDentry() *CommitFileResponseBodyDentry
	SetRequestId(v string) *CommitFileResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *CommitFileResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CommitFileResponseBody
	GetVendorType() *string
}

type CommitFileResponseBody struct {
	Dentry *CommitFileResponseBodyDentry `json:"dentry,omitempty" xml:"dentry,omitempty" type:"Struct"`
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

func (s CommitFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CommitFileResponseBody) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBody) GetDentry() *CommitFileResponseBodyDentry {
	return s.Dentry
}

func (s *CommitFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CommitFileResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CommitFileResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CommitFileResponseBody) SetDentry(v *CommitFileResponseBodyDentry) *CommitFileResponseBody {
	s.Dentry = v
	return s
}

func (s *CommitFileResponseBody) SetRequestId(v string) *CommitFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommitFileResponseBody) SetVendorRequestId(v string) *CommitFileResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CommitFileResponseBody) SetVendorType(v string) *CommitFileResponseBody {
	s.VendorType = &v
	return s
}

func (s *CommitFileResponseBody) Validate() error {
	if s.Dentry != nil {
		if err := s.Dentry.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CommitFileResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"AppProperties,omitempty" xml:"AppProperties,omitempty"`
	// example:
	//
	// DOCUMENT
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
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
	// parent_id
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// example:
	//
	// dentry_path
	Path       *string                                 `json:"Path,omitempty" xml:"Path,omitempty"`
	Properties *CommitFileResponseBodyDentryProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// space_id
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string                                `json:"StorageDriver,omitempty" xml:"StorageDriver,omitempty"`
	Thumbnail     *CommitFileResponseBodyDentryThumbnail `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty" type:"Struct"`
	// example:
	//
	// FILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// uuid
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CommitFileResponseBodyDentry) String() string {
	return dara.Prettify(s)
}

func (s CommitFileResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentry) GetAppProperties() map[string][]*DentryAppPropertiesValue {
	return s.AppProperties
}

func (s *CommitFileResponseBodyDentry) GetCategory() *string {
	return s.Category
}

func (s *CommitFileResponseBodyDentry) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CommitFileResponseBodyDentry) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CommitFileResponseBodyDentry) GetExtension() *string {
	return s.Extension
}

func (s *CommitFileResponseBodyDentry) GetId() *string {
	return s.Id
}

func (s *CommitFileResponseBodyDentry) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *CommitFileResponseBodyDentry) GetModifierId() *string {
	return s.ModifierId
}

func (s *CommitFileResponseBodyDentry) GetName() *string {
	return s.Name
}

func (s *CommitFileResponseBodyDentry) GetParentId() *string {
	return s.ParentId
}

func (s *CommitFileResponseBodyDentry) GetPartitionType() *string {
	return s.PartitionType
}

func (s *CommitFileResponseBodyDentry) GetPath() *string {
	return s.Path
}

func (s *CommitFileResponseBodyDentry) GetProperties() *CommitFileResponseBodyDentryProperties {
	return s.Properties
}

func (s *CommitFileResponseBodyDentry) GetSize() *int64 {
	return s.Size
}

func (s *CommitFileResponseBodyDentry) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CommitFileResponseBodyDentry) GetStatus() *string {
	return s.Status
}

func (s *CommitFileResponseBodyDentry) GetStorageDriver() *string {
	return s.StorageDriver
}

func (s *CommitFileResponseBodyDentry) GetThumbnail() *CommitFileResponseBodyDentryThumbnail {
	return s.Thumbnail
}

func (s *CommitFileResponseBodyDentry) GetType() *string {
	return s.Type
}

func (s *CommitFileResponseBodyDentry) GetUuid() *string {
	return s.Uuid
}

func (s *CommitFileResponseBodyDentry) GetVersion() *int64 {
	return s.Version
}

func (s *CommitFileResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *CommitFileResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCategory(v string) *CommitFileResponseBodyDentry {
	s.Category = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCreateTime(v string) *CommitFileResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetCreatorId(v string) *CommitFileResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetExtension(v string) *CommitFileResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetId(v string) *CommitFileResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetModifiedTime(v string) *CommitFileResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetModifierId(v string) *CommitFileResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetName(v string) *CommitFileResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetParentId(v string) *CommitFileResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetPartitionType(v string) *CommitFileResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetPath(v string) *CommitFileResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetProperties(v *CommitFileResponseBodyDentryProperties) *CommitFileResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetSize(v int64) *CommitFileResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetSpaceId(v string) *CommitFileResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetStatus(v string) *CommitFileResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetStorageDriver(v string) *CommitFileResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetThumbnail(v *CommitFileResponseBodyDentryThumbnail) *CommitFileResponseBodyDentry {
	s.Thumbnail = v
	return s
}

func (s *CommitFileResponseBodyDentry) SetType(v string) *CommitFileResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetUuid(v string) *CommitFileResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *CommitFileResponseBodyDentry) SetVersion(v int64) *CommitFileResponseBodyDentry {
	s.Version = &v
	return s
}

func (s *CommitFileResponseBodyDentry) Validate() error {
	if s.Properties != nil {
		if err := s.Properties.Validate(); err != nil {
			return err
		}
	}
	if s.Thumbnail != nil {
		if err := s.Thumbnail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CommitFileResponseBodyDentryProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s CommitFileResponseBodyDentryProperties) String() string {
	return dara.Prettify(s)
}

func (s CommitFileResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentryProperties) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CommitFileResponseBodyDentryProperties) SetReadOnly(v bool) *CommitFileResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

func (s *CommitFileResponseBodyDentryProperties) Validate() error {
	return dara.Validate(s)
}

type CommitFileResponseBodyDentryThumbnail struct {
	// example:
	//
	// 64
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 64
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CommitFileResponseBodyDentryThumbnail) String() string {
	return dara.Prettify(s)
}

func (s CommitFileResponseBodyDentryThumbnail) GoString() string {
	return s.String()
}

func (s *CommitFileResponseBodyDentryThumbnail) GetHeight() *int32 {
	return s.Height
}

func (s *CommitFileResponseBodyDentryThumbnail) GetUrl() *string {
	return s.Url
}

func (s *CommitFileResponseBodyDentryThumbnail) GetWidth() *int32 {
	return s.Width
}

func (s *CommitFileResponseBodyDentryThumbnail) SetHeight(v int32) *CommitFileResponseBodyDentryThumbnail {
	s.Height = &v
	return s
}

func (s *CommitFileResponseBodyDentryThumbnail) SetUrl(v string) *CommitFileResponseBodyDentryThumbnail {
	s.Url = &v
	return s
}

func (s *CommitFileResponseBodyDentryThumbnail) SetWidth(v int32) *CommitFileResponseBodyDentryThumbnail {
	s.Width = &v
	return s
}

func (s *CommitFileResponseBodyDentryThumbnail) Validate() error {
	return dara.Validate(s)
}
