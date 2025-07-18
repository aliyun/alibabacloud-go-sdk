// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertChunksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowInsertWithFilter(v bool) *UpsertChunksShrinkRequest
	GetAllowInsertWithFilter() *bool
	SetCollection(v string) *UpsertChunksShrinkRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UpsertChunksShrinkRequest
	GetDBInstanceId() *string
	SetFileName(v string) *UpsertChunksShrinkRequest
	GetFileName() *string
	SetNamespace(v string) *UpsertChunksShrinkRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UpsertChunksShrinkRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UpsertChunksShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpsertChunksShrinkRequest
	GetRegionId() *string
	SetShouldReplaceFile(v bool) *UpsertChunksShrinkRequest
	GetShouldReplaceFile() *bool
	SetTextChunksShrink(v string) *UpsertChunksShrinkRequest
	GetTextChunksShrink() *string
}

type UpsertChunksShrinkRequest struct {
	AllowInsertWithFilter *bool `json:"AllowInsertWithFilter,omitempty" xml:"AllowInsertWithFilter,omitempty"`
	// Document collection name.
	//
	// > Created by the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) API. You can use the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) API to view the already created document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// File name.
	//
	// > If a file name is specified and not empty, it will overwrite the data for this file name; if empty, the chunks data will be appended directly to the document collection.
	//
	// example:
	//
	// mydoc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Namespace, default is public.
	//
	// > You can create it using the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) API and view the list using the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) API.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Password corresponding to the namespace.
	//
	// > This value is specified by the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region ID where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShouldReplaceFile *bool   `json:"ShouldReplaceFile,omitempty" xml:"ShouldReplaceFile,omitempty"`
	// List of split documents.
	TextChunksShrink *string `json:"TextChunks,omitempty" xml:"TextChunks,omitempty"`
}

func (s UpsertChunksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertChunksShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpsertChunksShrinkRequest) GetAllowInsertWithFilter() *bool {
	return s.AllowInsertWithFilter
}

func (s *UpsertChunksShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *UpsertChunksShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpsertChunksShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *UpsertChunksShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpsertChunksShrinkRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UpsertChunksShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpsertChunksShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpsertChunksShrinkRequest) GetShouldReplaceFile() *bool {
	return s.ShouldReplaceFile
}

func (s *UpsertChunksShrinkRequest) GetTextChunksShrink() *string {
	return s.TextChunksShrink
}

func (s *UpsertChunksShrinkRequest) SetAllowInsertWithFilter(v bool) *UpsertChunksShrinkRequest {
	s.AllowInsertWithFilter = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetCollection(v string) *UpsertChunksShrinkRequest {
	s.Collection = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetDBInstanceId(v string) *UpsertChunksShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetFileName(v string) *UpsertChunksShrinkRequest {
	s.FileName = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetNamespace(v string) *UpsertChunksShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetNamespacePassword(v string) *UpsertChunksShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetOwnerId(v int64) *UpsertChunksShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetRegionId(v string) *UpsertChunksShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetShouldReplaceFile(v bool) *UpsertChunksShrinkRequest {
	s.ShouldReplaceFile = &v
	return s
}

func (s *UpsertChunksShrinkRequest) SetTextChunksShrink(v string) *UpsertChunksShrinkRequest {
	s.TextChunksShrink = &v
	return s
}

func (s *UpsertChunksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
