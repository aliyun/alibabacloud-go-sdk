// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDentriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDentries(v []*ListDentriesResponseBodyDentries) *ListDentriesResponseBody
	GetDentries() []*ListDentriesResponseBodyDentries
	SetNextToken(v string) *ListDentriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDentriesResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *ListDentriesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListDentriesResponseBody
	GetVendorType() *string
}

type ListDentriesResponseBody struct {
	Dentries []*ListDentriesResponseBodyDentries `json:"dentries,omitempty" xml:"dentries,omitempty" type:"Repeated"`
	// example:
	//
	// next_token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

func (s ListDentriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBody) GetDentries() []*ListDentriesResponseBodyDentries {
	return s.Dentries
}

func (s *ListDentriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDentriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDentriesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListDentriesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListDentriesResponseBody) SetDentries(v []*ListDentriesResponseBodyDentries) *ListDentriesResponseBody {
	s.Dentries = v
	return s
}

func (s *ListDentriesResponseBody) SetNextToken(v string) *ListDentriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDentriesResponseBody) SetRequestId(v string) *ListDentriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDentriesResponseBody) SetVendorRequestId(v string) *ListDentriesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListDentriesResponseBody) SetVendorType(v string) *ListDentriesResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListDentriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDentriesResponseBodyDentries struct {
	AppProperties map[string][]*DentriesAppPropertiesValue `json:"AppProperties,omitempty" xml:"AppProperties,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// cHtUxxxxx
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// txt
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// 657xxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-01-01T10:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// cHtUxxxxx
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// 测试文件夹
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// PUBLIC_OSS_PARTITION
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// example:
	//
	// ./test.txt
	Path       *string                                     `json:"Path,omitempty" xml:"Path,omitempty"`
	Properties *ListDentriesResponseBodyDentriesProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	// example:
	//
	// 512
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 854xxxxx
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DINGTALK
	StorageDriver *string                                    `json:"StorageDriver,omitempty" xml:"StorageDriver,omitempty"`
	Thumbnail     *ListDentriesResponseBodyDentriesThumbnail `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty" type:"Struct"`
	// example:
	//
	// FILE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 123xxxxx
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDentriesResponseBodyDentries) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesResponseBodyDentries) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentries) GetAppProperties() map[string][]*DentriesAppPropertiesValue {
	return s.AppProperties
}

func (s *ListDentriesResponseBodyDentries) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDentriesResponseBodyDentries) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDentriesResponseBodyDentries) GetExtension() *string {
	return s.Extension
}

func (s *ListDentriesResponseBodyDentries) GetId() *string {
	return s.Id
}

func (s *ListDentriesResponseBodyDentries) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListDentriesResponseBodyDentries) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListDentriesResponseBodyDentries) GetName() *string {
	return s.Name
}

func (s *ListDentriesResponseBodyDentries) GetParentId() *string {
	return s.ParentId
}

func (s *ListDentriesResponseBodyDentries) GetPartitionType() *string {
	return s.PartitionType
}

func (s *ListDentriesResponseBodyDentries) GetPath() *string {
	return s.Path
}

func (s *ListDentriesResponseBodyDentries) GetProperties() *ListDentriesResponseBodyDentriesProperties {
	return s.Properties
}

func (s *ListDentriesResponseBodyDentries) GetSize() *int64 {
	return s.Size
}

func (s *ListDentriesResponseBodyDentries) GetSpaceId() *string {
	return s.SpaceId
}

func (s *ListDentriesResponseBodyDentries) GetStatus() *string {
	return s.Status
}

func (s *ListDentriesResponseBodyDentries) GetStorageDriver() *string {
	return s.StorageDriver
}

func (s *ListDentriesResponseBodyDentries) GetThumbnail() *ListDentriesResponseBodyDentriesThumbnail {
	return s.Thumbnail
}

func (s *ListDentriesResponseBodyDentries) GetType() *string {
	return s.Type
}

func (s *ListDentriesResponseBodyDentries) GetUuid() *string {
	return s.Uuid
}

func (s *ListDentriesResponseBodyDentries) GetVersion() *int64 {
	return s.Version
}

func (s *ListDentriesResponseBodyDentries) SetAppProperties(v map[string][]*DentriesAppPropertiesValue) *ListDentriesResponseBodyDentries {
	s.AppProperties = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetCreateTime(v string) *ListDentriesResponseBodyDentries {
	s.CreateTime = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetCreatorId(v string) *ListDentriesResponseBodyDentries {
	s.CreatorId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetExtension(v string) *ListDentriesResponseBodyDentries {
	s.Extension = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetId(v string) *ListDentriesResponseBodyDentries {
	s.Id = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetModifiedTime(v string) *ListDentriesResponseBodyDentries {
	s.ModifiedTime = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetModifierId(v string) *ListDentriesResponseBodyDentries {
	s.ModifierId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetName(v string) *ListDentriesResponseBodyDentries {
	s.Name = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetParentId(v string) *ListDentriesResponseBodyDentries {
	s.ParentId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetPartitionType(v string) *ListDentriesResponseBodyDentries {
	s.PartitionType = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetPath(v string) *ListDentriesResponseBodyDentries {
	s.Path = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetProperties(v *ListDentriesResponseBodyDentriesProperties) *ListDentriesResponseBodyDentries {
	s.Properties = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetSize(v int64) *ListDentriesResponseBodyDentries {
	s.Size = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetSpaceId(v string) *ListDentriesResponseBodyDentries {
	s.SpaceId = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetStatus(v string) *ListDentriesResponseBodyDentries {
	s.Status = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetStorageDriver(v string) *ListDentriesResponseBodyDentries {
	s.StorageDriver = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetThumbnail(v *ListDentriesResponseBodyDentriesThumbnail) *ListDentriesResponseBodyDentries {
	s.Thumbnail = v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetType(v string) *ListDentriesResponseBodyDentries {
	s.Type = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetUuid(v string) *ListDentriesResponseBodyDentries {
	s.Uuid = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) SetVersion(v int64) *ListDentriesResponseBodyDentries {
	s.Version = &v
	return s
}

func (s *ListDentriesResponseBodyDentries) Validate() error {
	return dara.Validate(s)
}

type ListDentriesResponseBodyDentriesProperties struct {
	// example:
	//
	// true
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s ListDentriesResponseBodyDentriesProperties) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesResponseBodyDentriesProperties) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentriesProperties) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListDentriesResponseBodyDentriesProperties) SetReadOnly(v bool) *ListDentriesResponseBodyDentriesProperties {
	s.ReadOnly = &v
	return s
}

func (s *ListDentriesResponseBodyDentriesProperties) Validate() error {
	return dara.Validate(s)
}

type ListDentriesResponseBodyDentriesThumbnail struct {
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

func (s ListDentriesResponseBodyDentriesThumbnail) String() string {
	return dara.Prettify(s)
}

func (s ListDentriesResponseBodyDentriesThumbnail) GoString() string {
	return s.String()
}

func (s *ListDentriesResponseBodyDentriesThumbnail) GetHeight() *int32 {
	return s.Height
}

func (s *ListDentriesResponseBodyDentriesThumbnail) GetUrl() *string {
	return s.Url
}

func (s *ListDentriesResponseBodyDentriesThumbnail) GetWidth() *int32 {
	return s.Width
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetHeight(v int32) *ListDentriesResponseBodyDentriesThumbnail {
	s.Height = &v
	return s
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetUrl(v string) *ListDentriesResponseBodyDentriesThumbnail {
	s.Url = &v
	return s
}

func (s *ListDentriesResponseBodyDentriesThumbnail) SetWidth(v int32) *ListDentriesResponseBodyDentriesThumbnail {
	s.Width = &v
	return s
}

func (s *ListDentriesResponseBodyDentriesThumbnail) Validate() error {
	return dara.Validate(s)
}
