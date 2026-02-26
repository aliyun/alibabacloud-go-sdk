// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppMaterialFile interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *AppMaterialFile
	GetBizId() *string
	SetContentType(v string) *AppMaterialFile
	GetContentType() *string
	SetCreateTime(v string) *AppMaterialFile
	GetCreateTime() *string
	SetDeletedTime(v string) *AppMaterialFile
	GetDeletedTime() *string
	SetDirectoryId(v string) *AppMaterialFile
	GetDirectoryId() *string
	SetFileId(v string) *AppMaterialFile
	GetFileId() *string
	SetFileUrl(v string) *AppMaterialFile
	GetFileUrl() *string
	SetHeight(v int32) *AppMaterialFile
	GetHeight() *int32
	SetName(v string) *AppMaterialFile
	GetName() *string
	SetStatus(v string) *AppMaterialFile
	GetStatus() *string
	SetStorageSize(v string) *AppMaterialFile
	GetStorageSize() *string
	SetSuffix(v string) *AppMaterialFile
	GetSuffix() *string
	SetType(v string) *AppMaterialFile
	GetType() *string
	SetWidth(v int32) *AppMaterialFile
	GetWidth() *int32
}

type AppMaterialFile struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// image/png
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 2025-11-03 02:05:01
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2025-11-03 02:05:01
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// example:
	//
	// ddddwww
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// d75e1e9693cc460da218419d67735567
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// https://demo-moshicloud-test.oss-cn-hangzhou-internal.aliyuncs.com/document/testcase/aliyun/api_case/pdf_angle90_page1.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 1280
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 23M
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// example:
	//
	// png
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
	// example:
	//
	// IMAGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AppMaterialFile) String() string {
	return dara.Prettify(s)
}

func (s AppMaterialFile) GoString() string {
	return s.String()
}

func (s *AppMaterialFile) GetBizId() *string {
	return s.BizId
}

func (s *AppMaterialFile) GetContentType() *string {
	return s.ContentType
}

func (s *AppMaterialFile) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AppMaterialFile) GetDeletedTime() *string {
	return s.DeletedTime
}

func (s *AppMaterialFile) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *AppMaterialFile) GetFileId() *string {
	return s.FileId
}

func (s *AppMaterialFile) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AppMaterialFile) GetHeight() *int32 {
	return s.Height
}

func (s *AppMaterialFile) GetName() *string {
	return s.Name
}

func (s *AppMaterialFile) GetStatus() *string {
	return s.Status
}

func (s *AppMaterialFile) GetStorageSize() *string {
	return s.StorageSize
}

func (s *AppMaterialFile) GetSuffix() *string {
	return s.Suffix
}

func (s *AppMaterialFile) GetType() *string {
	return s.Type
}

func (s *AppMaterialFile) GetWidth() *int32 {
	return s.Width
}

func (s *AppMaterialFile) SetBizId(v string) *AppMaterialFile {
	s.BizId = &v
	return s
}

func (s *AppMaterialFile) SetContentType(v string) *AppMaterialFile {
	s.ContentType = &v
	return s
}

func (s *AppMaterialFile) SetCreateTime(v string) *AppMaterialFile {
	s.CreateTime = &v
	return s
}

func (s *AppMaterialFile) SetDeletedTime(v string) *AppMaterialFile {
	s.DeletedTime = &v
	return s
}

func (s *AppMaterialFile) SetDirectoryId(v string) *AppMaterialFile {
	s.DirectoryId = &v
	return s
}

func (s *AppMaterialFile) SetFileId(v string) *AppMaterialFile {
	s.FileId = &v
	return s
}

func (s *AppMaterialFile) SetFileUrl(v string) *AppMaterialFile {
	s.FileUrl = &v
	return s
}

func (s *AppMaterialFile) SetHeight(v int32) *AppMaterialFile {
	s.Height = &v
	return s
}

func (s *AppMaterialFile) SetName(v string) *AppMaterialFile {
	s.Name = &v
	return s
}

func (s *AppMaterialFile) SetStatus(v string) *AppMaterialFile {
	s.Status = &v
	return s
}

func (s *AppMaterialFile) SetStorageSize(v string) *AppMaterialFile {
	s.StorageSize = &v
	return s
}

func (s *AppMaterialFile) SetSuffix(v string) *AppMaterialFile {
	s.Suffix = &v
	return s
}

func (s *AppMaterialFile) SetType(v string) *AppMaterialFile {
	s.Type = &v
	return s
}

func (s *AppMaterialFile) SetWidth(v int32) *AppMaterialFile {
	s.Width = &v
	return s
}

func (s *AppMaterialFile) Validate() error {
	return dara.Validate(s)
}
