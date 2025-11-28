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
	// Based on the Filter input specified under TextChunks, this parameter controls whether data insertion is allowed when a Filter is provided.
	//
	// If AllowInsertWithFilter = true, the insert operation is performed when the filter does not match any data.
	//
	// If AllowInsertWithFilter = false, no action is performed if the filter does not match any data.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AllowInsertWithFilter *bool `json:"AllowInsertWithFilter,omitempty" xml:"AllowInsertWithFilter,omitempty"`
	// The name of the document collection.
	//
	// > You can call the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation to create a document collection and call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to query a list of document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The file name of the document.
	//
	// > When a non-empty filename is specified, the system will decide whether to overwrite the data associated with that filename based on the value of the ShouldReplaceFile parameter. If you leave this parameter empty, the data of chunks is appended to the document collection.
	//
	// example:
	//
	// mydoc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// > You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// > The value of this parameter is specified when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to overwrite the data associated with the file name specified by the FileName parameter.
	//
	// If you set ShouldReplaceFile to true, the system deletes all data associated with the file name and then inserts new data.
	//
	// If you set ShouldReplaceFile to false, the system does not delete the data associated with the file name, but inserts or updates the data of chunks based on the TextChunks parameter.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	ShouldReplaceFile *bool `json:"ShouldReplaceFile,omitempty" xml:"ShouldReplaceFile,omitempty"`
	// List of document chunks after splitting.
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
