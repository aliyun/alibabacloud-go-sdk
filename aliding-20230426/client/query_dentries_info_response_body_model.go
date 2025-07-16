// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentriesInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDentry(v *QueryDentriesInfoResponseBodyDentry) *QueryDentriesInfoResponseBody
	GetDentry() *QueryDentriesInfoResponseBodyDentry
	SetRequestId(v string) *QueryDentriesInfoResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *QueryDentriesInfoResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *QueryDentriesInfoResponseBody
	GetVendorType() *string
}

type QueryDentriesInfoResponseBody struct {
	Dentry *QueryDentriesInfoResponseBodyDentry `json:"Dentry,omitempty" xml:"Dentry,omitempty" type:"Struct"`
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

func (s QueryDentriesInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoResponseBody) GetDentry() *QueryDentriesInfoResponseBodyDentry {
	return s.Dentry
}

func (s *QueryDentriesInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDentriesInfoResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *QueryDentriesInfoResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *QueryDentriesInfoResponseBody) SetDentry(v *QueryDentriesInfoResponseBodyDentry) *QueryDentriesInfoResponseBody {
	s.Dentry = v
	return s
}

func (s *QueryDentriesInfoResponseBody) SetRequestId(v string) *QueryDentriesInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDentriesInfoResponseBody) SetVendorRequestId(v string) *QueryDentriesInfoResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *QueryDentriesInfoResponseBody) SetVendorType(v string) *QueryDentriesInfoResponseBody {
	s.VendorType = &v
	return s
}

func (s *QueryDentriesInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDentriesInfoResponseBodyDentry struct {
	AppProperties map[string][]*DentryAppPropertiesValue `json:"AppProperties,omitempty" xml:"AppProperties,omitempty"`
	// example:
	//
	// 2025-03-26T02:19:35Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// qt8bGiSa7WnHKeRPtMuoiSJwiE
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 140901622636
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// qt8bGiSa7WnHKeRPtMuoiSJwiE
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string                                        `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	Path          *string                                        `json:"Path,omitempty" xml:"Path,omitempty"`
	Properties    *QueryDentriesInfoResponseBodyDentryProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 22443475065
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string                                       `json:"StorageDriver,omitempty" xml:"StorageDriver,omitempty"`
	Thumbnail     *QueryDentriesInfoResponseBodyDentryThumbnail `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty" type:"Struct"`
	// example:
	//
	// FILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1716258459684
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryDentriesInfoResponseBodyDentry) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoResponseBodyDentry) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoResponseBodyDentry) GetAppProperties() map[string][]*DentryAppPropertiesValue {
	return s.AppProperties
}

func (s *QueryDentriesInfoResponseBodyDentry) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryDentriesInfoResponseBodyDentry) GetCreatorId() *string {
	return s.CreatorId
}

func (s *QueryDentriesInfoResponseBodyDentry) GetExtension() *string {
	return s.Extension
}

func (s *QueryDentriesInfoResponseBodyDentry) GetId() *string {
	return s.Id
}

func (s *QueryDentriesInfoResponseBodyDentry) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *QueryDentriesInfoResponseBodyDentry) GetModifierId() *string {
	return s.ModifierId
}

func (s *QueryDentriesInfoResponseBodyDentry) GetName() *string {
	return s.Name
}

func (s *QueryDentriesInfoResponseBodyDentry) GetParentId() *string {
	return s.ParentId
}

func (s *QueryDentriesInfoResponseBodyDentry) GetPartitionType() *string {
	return s.PartitionType
}

func (s *QueryDentriesInfoResponseBodyDentry) GetPath() *string {
	return s.Path
}

func (s *QueryDentriesInfoResponseBodyDentry) GetProperties() *QueryDentriesInfoResponseBodyDentryProperties {
	return s.Properties
}

func (s *QueryDentriesInfoResponseBodyDentry) GetSize() *int64 {
	return s.Size
}

func (s *QueryDentriesInfoResponseBodyDentry) GetSpaceId() *string {
	return s.SpaceId
}

func (s *QueryDentriesInfoResponseBodyDentry) GetStatus() *string {
	return s.Status
}

func (s *QueryDentriesInfoResponseBodyDentry) GetStorageDriver() *string {
	return s.StorageDriver
}

func (s *QueryDentriesInfoResponseBodyDentry) GetThumbnail() *QueryDentriesInfoResponseBodyDentryThumbnail {
	return s.Thumbnail
}

func (s *QueryDentriesInfoResponseBodyDentry) GetType() *string {
	return s.Type
}

func (s *QueryDentriesInfoResponseBodyDentry) GetUuid() *string {
	return s.Uuid
}

func (s *QueryDentriesInfoResponseBodyDentry) GetVersion() *int64 {
	return s.Version
}

func (s *QueryDentriesInfoResponseBodyDentry) SetAppProperties(v map[string][]*DentryAppPropertiesValue) *QueryDentriesInfoResponseBodyDentry {
	s.AppProperties = v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetCreateTime(v string) *QueryDentriesInfoResponseBodyDentry {
	s.CreateTime = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetCreatorId(v string) *QueryDentriesInfoResponseBodyDentry {
	s.CreatorId = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetExtension(v string) *QueryDentriesInfoResponseBodyDentry {
	s.Extension = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetId(v string) *QueryDentriesInfoResponseBodyDentry {
	s.Id = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetModifiedTime(v string) *QueryDentriesInfoResponseBodyDentry {
	s.ModifiedTime = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetModifierId(v string) *QueryDentriesInfoResponseBodyDentry {
	s.ModifierId = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetName(v string) *QueryDentriesInfoResponseBodyDentry {
	s.Name = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetParentId(v string) *QueryDentriesInfoResponseBodyDentry {
	s.ParentId = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetPartitionType(v string) *QueryDentriesInfoResponseBodyDentry {
	s.PartitionType = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetPath(v string) *QueryDentriesInfoResponseBodyDentry {
	s.Path = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetProperties(v *QueryDentriesInfoResponseBodyDentryProperties) *QueryDentriesInfoResponseBodyDentry {
	s.Properties = v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetSize(v int64) *QueryDentriesInfoResponseBodyDentry {
	s.Size = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetSpaceId(v string) *QueryDentriesInfoResponseBodyDentry {
	s.SpaceId = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetStatus(v string) *QueryDentriesInfoResponseBodyDentry {
	s.Status = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetStorageDriver(v string) *QueryDentriesInfoResponseBodyDentry {
	s.StorageDriver = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetThumbnail(v *QueryDentriesInfoResponseBodyDentryThumbnail) *QueryDentriesInfoResponseBodyDentry {
	s.Thumbnail = v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetType(v string) *QueryDentriesInfoResponseBodyDentry {
	s.Type = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetUuid(v string) *QueryDentriesInfoResponseBodyDentry {
	s.Uuid = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) SetVersion(v int64) *QueryDentriesInfoResponseBodyDentry {
	s.Version = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentry) Validate() error {
	return dara.Validate(s)
}

type QueryDentriesInfoResponseBodyDentryProperties struct {
	// example:
	//
	// True
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s QueryDentriesInfoResponseBodyDentryProperties) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoResponseBodyDentryProperties) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoResponseBodyDentryProperties) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *QueryDentriesInfoResponseBodyDentryProperties) SetReadOnly(v bool) *QueryDentriesInfoResponseBodyDentryProperties {
	s.ReadOnly = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentryProperties) Validate() error {
	return dara.Validate(s)
}

type QueryDentriesInfoResponseBodyDentryThumbnail struct {
	// example:
	//
	// 720
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryDentriesInfoResponseBodyDentryThumbnail) String() string {
	return dara.Prettify(s)
}

func (s QueryDentriesInfoResponseBodyDentryThumbnail) GoString() string {
	return s.String()
}

func (s *QueryDentriesInfoResponseBodyDentryThumbnail) GetHeight() *int32 {
	return s.Height
}

func (s *QueryDentriesInfoResponseBodyDentryThumbnail) GetUrl() *string {
	return s.Url
}

func (s *QueryDentriesInfoResponseBodyDentryThumbnail) GetWidth() *int32 {
	return s.Width
}

func (s *QueryDentriesInfoResponseBodyDentryThumbnail) SetHeight(v int32) *QueryDentriesInfoResponseBodyDentryThumbnail {
	s.Height = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentryThumbnail) SetUrl(v string) *QueryDentriesInfoResponseBodyDentryThumbnail {
	s.Url = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentryThumbnail) SetWidth(v int32) *QueryDentriesInfoResponseBodyDentryThumbnail {
	s.Width = &v
	return s
}

func (s *QueryDentriesInfoResponseBodyDentryThumbnail) Validate() error {
	return dara.Validate(s)
}
