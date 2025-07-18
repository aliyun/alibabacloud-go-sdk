// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCollectionDataMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *UpdateCollectionDataMetadataRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UpdateCollectionDataMetadataRequest
	GetDBInstanceId() *string
	SetFilter(v string) *UpdateCollectionDataMetadataRequest
	GetFilter() *string
	SetIds(v []*string) *UpdateCollectionDataMetadataRequest
	GetIds() []*string
	SetMetadata(v map[string]interface{}) *UpdateCollectionDataMetadataRequest
	GetMetadata() map[string]interface{}
	SetNamespace(v string) *UpdateCollectionDataMetadataRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UpdateCollectionDataMetadataRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UpdateCollectionDataMetadataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateCollectionDataMetadataRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *UpdateCollectionDataMetadataRequest
	GetWorkspaceId() *string
}

type UpdateCollectionDataMetadataRequest struct {
	// Collection name.
	//
	// > You can use the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) API to view the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB for PostgreSQL instances in the target region, including the instance ID.
	//
	// example:
	//
	// gp-j788ghhjjxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Filter condition for the data to be updated, in SQL WHERE format. This field cannot be empty at the same time as the Ids field.
	//
	// example:
	//
	// business_value = \\"chat_file_1\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// ID list of the data to be updated, i.e., the Row.Id specified when uploading the data. This field cannot be empty at the same time as the Filter field.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// Data to be updated, in a JSON string of MAP format. The key is the field name, and the value is the new data value.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//       "title": "new title",
	//
	//       "content": "new content"
	//
	// }
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Namespace.
	//
	// > You can use the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) API to view the list.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Password corresponding to the namespace.
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// ID of the Workspace composed of multiple database instances. This parameter and the DBInstanceId parameter cannot both be empty. When both are specified, this parameter takes precedence.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateCollectionDataMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectionDataMetadataRequest) GoString() string {
	return s.String()
}

func (s *UpdateCollectionDataMetadataRequest) GetCollection() *string {
	return s.Collection
}

func (s *UpdateCollectionDataMetadataRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpdateCollectionDataMetadataRequest) GetFilter() *string {
	return s.Filter
}

func (s *UpdateCollectionDataMetadataRequest) GetIds() []*string {
	return s.Ids
}

func (s *UpdateCollectionDataMetadataRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *UpdateCollectionDataMetadataRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateCollectionDataMetadataRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UpdateCollectionDataMetadataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCollectionDataMetadataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCollectionDataMetadataRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateCollectionDataMetadataRequest) SetCollection(v string) *UpdateCollectionDataMetadataRequest {
	s.Collection = &v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetDBInstanceId(v string) *UpdateCollectionDataMetadataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetFilter(v string) *UpdateCollectionDataMetadataRequest {
	s.Filter = &v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetIds(v []*string) *UpdateCollectionDataMetadataRequest {
	s.Ids = v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetMetadata(v map[string]interface{}) *UpdateCollectionDataMetadataRequest {
	s.Metadata = v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetNamespace(v string) *UpdateCollectionDataMetadataRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetNamespacePassword(v string) *UpdateCollectionDataMetadataRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetOwnerId(v int64) *UpdateCollectionDataMetadataRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetRegionId(v string) *UpdateCollectionDataMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) SetWorkspaceId(v string) *UpdateCollectionDataMetadataRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateCollectionDataMetadataRequest) Validate() error {
	return dara.Validate(s)
}
