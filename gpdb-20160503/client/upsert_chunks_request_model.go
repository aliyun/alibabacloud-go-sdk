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
	if s.TextChunks != nil {
		for _, item := range s.TextChunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpsertChunksRequestTextChunks struct {
	// The document content.
	//
	// This parameter is required.
	//
	// example:
	//
	// AnalyticDB for PostgreSQL provides a simple, fast, and cost-effective PB-level cloud data warehousing solution.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The filter condition for the data to be updated, in the format of an SQL WHERE clause.
	//
	// example:
	//
	// title = \\"seagull\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The unique ID of the vector data.
	//
	// example:
	//
	// 273e3fc7-8f56-4167-a1bb-d35d2f3b9043
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metadata.
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
