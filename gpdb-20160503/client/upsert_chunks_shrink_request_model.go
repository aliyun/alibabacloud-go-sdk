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
	// Specifies whether to allow data insertion when a filter is specified. This parameter takes effect only when the Filter parameter in TextChunks is specified.
	//
	// If AllowInsertWithFilter is set to true, an insert operation is performed when the filter does not match any data.
	//
	// If AllowInsertWithFilter is set to false, no operation is performed when the filter does not match any data.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AllowInsertWithFilter *bool `json:"AllowInsertWithFilter,omitempty" xml:"AllowInsertWithFilter,omitempty"`
	// The name of the document collection.
	//
	// > The document collection is created by calling the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. You can call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to view the created document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including the instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The file name.
	//
	// > If this parameter is not empty, the system determines whether to overwrite the data associated with this file name based on the value of the ShouldReplaceFile parameter. If this parameter is empty, the chunk data is directly appended and uploaded to the document collection.
	//
	// example:
	//
	// mydoc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The namespace. Default value: public.
	//
	// > You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to view the namespace list.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
	//
	// > This value is specified in the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to overwrite the data associated with the specified file name. This parameter takes effect only when the FileName parameter is specified.
	//
	// If ShouldReplaceFile is set to true, all data associated with the file name is deleted before new data is inserted.
	//
	// If ShouldReplaceFile is set to false, the data associated with the file name is not deleted. Instead, chunks are inserted or updated based on the specific TextChunks parameter.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	ShouldReplaceFile *bool `json:"ShouldReplaceFile,omitempty" xml:"ShouldReplaceFile,omitempty"`
	// The list of chunked documents.
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
