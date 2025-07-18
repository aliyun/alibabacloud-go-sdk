// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowInsertWithFilter(v bool) *UpsertChunksRequest
	GetAllowInsertWithFilter() *bool
	SetCollection(v string) *UpsertChunksRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UpsertChunksRequest
	GetDBInstanceId() *string
	SetFileName(v string) *UpsertChunksRequest
	GetFileName() *string
	SetNamespace(v string) *UpsertChunksRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UpsertChunksRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UpsertChunksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpsertChunksRequest
	GetRegionId() *string
	SetShouldReplaceFile(v bool) *UpsertChunksRequest
	GetShouldReplaceFile() *bool
	SetTextChunks(v []*UpsertChunksRequestTextChunks) *UpsertChunksRequest
	GetTextChunks() []*UpsertChunksRequestTextChunks
}

type UpsertChunksRequest struct {
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
	TextChunks []*UpsertChunksRequestTextChunks `json:"TextChunks,omitempty" xml:"TextChunks,omitempty" type:"Repeated"`
}

func (s UpsertChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertChunksRequest) GoString() string {
	return s.String()
}

func (s *UpsertChunksRequest) GetAllowInsertWithFilter() *bool {
	return s.AllowInsertWithFilter
}

func (s *UpsertChunksRequest) GetCollection() *string {
	return s.Collection
}

func (s *UpsertChunksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpsertChunksRequest) GetFileName() *string {
	return s.FileName
}

func (s *UpsertChunksRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpsertChunksRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UpsertChunksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpsertChunksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpsertChunksRequest) GetShouldReplaceFile() *bool {
	return s.ShouldReplaceFile
}

func (s *UpsertChunksRequest) GetTextChunks() []*UpsertChunksRequestTextChunks {
	return s.TextChunks
}

func (s *UpsertChunksRequest) SetAllowInsertWithFilter(v bool) *UpsertChunksRequest {
	s.AllowInsertWithFilter = &v
	return s
}

func (s *UpsertChunksRequest) SetCollection(v string) *UpsertChunksRequest {
	s.Collection = &v
	return s
}

func (s *UpsertChunksRequest) SetDBInstanceId(v string) *UpsertChunksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpsertChunksRequest) SetFileName(v string) *UpsertChunksRequest {
	s.FileName = &v
	return s
}

func (s *UpsertChunksRequest) SetNamespace(v string) *UpsertChunksRequest {
	s.Namespace = &v
	return s
}

func (s *UpsertChunksRequest) SetNamespacePassword(v string) *UpsertChunksRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpsertChunksRequest) SetOwnerId(v int64) *UpsertChunksRequest {
	s.OwnerId = &v
	return s
}

func (s *UpsertChunksRequest) SetRegionId(v string) *UpsertChunksRequest {
	s.RegionId = &v
	return s
}

func (s *UpsertChunksRequest) SetShouldReplaceFile(v bool) *UpsertChunksRequest {
	s.ShouldReplaceFile = &v
	return s
}

func (s *UpsertChunksRequest) SetTextChunks(v []*UpsertChunksRequestTextChunks) *UpsertChunksRequest {
	s.TextChunks = v
	return s
}

func (s *UpsertChunksRequest) Validate() error {
	return dara.Validate(s)
}

type UpsertChunksRequestTextChunks struct {
	// Document content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cloud-native data warehouse AnalyticDB PostgreSQL Edition provides a simple, fast, and cost-effective PB-level cloud data warehouse solution.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Filter  *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Metadata.
	//
	// example:
	//
	// {"title":"test"}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
}

func (s UpsertChunksRequestTextChunks) String() string {
	return dara.Prettify(s)
}

func (s UpsertChunksRequestTextChunks) GoString() string {
	return s.String()
}

func (s *UpsertChunksRequestTextChunks) GetContent() *string {
	return s.Content
}

func (s *UpsertChunksRequestTextChunks) GetFilter() *string {
	return s.Filter
}

func (s *UpsertChunksRequestTextChunks) GetId() *string {
	return s.Id
}

func (s *UpsertChunksRequestTextChunks) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *UpsertChunksRequestTextChunks) SetContent(v string) *UpsertChunksRequestTextChunks {
	s.Content = &v
	return s
}

func (s *UpsertChunksRequestTextChunks) SetFilter(v string) *UpsertChunksRequestTextChunks {
	s.Filter = &v
	return s
}

func (s *UpsertChunksRequestTextChunks) SetId(v string) *UpsertChunksRequestTextChunks {
	s.Id = &v
	return s
}

func (s *UpsertChunksRequestTextChunks) SetMetadata(v map[string]interface{}) *UpsertChunksRequestTextChunks {
	s.Metadata = v
	return s
}

func (s *UpsertChunksRequestTextChunks) Validate() error {
	return dara.Validate(s)
}
