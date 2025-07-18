// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCollectionDataMetadataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetDBInstanceId() *string
	SetFilter(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetFilter() *string
	SetIdsShrink(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetIdsShrink() *string
	SetMetadataShrink(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetMetadataShrink() *string
	SetNamespace(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UpdateCollectionDataMetadataShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *UpdateCollectionDataMetadataShrinkRequest
	GetWorkspaceId() *string
}

type UpdateCollectionDataMetadataShrinkRequest struct {
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
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
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
	MetadataShrink *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
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

func (s UpdateCollectionDataMetadataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectionDataMetadataShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCollectionDataMetadataShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetCollection(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.Collection = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetDBInstanceId(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetFilter(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.Filter = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetIdsShrink(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetMetadataShrink(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetNamespace(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetNamespacePassword(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetOwnerId(v int64) *UpdateCollectionDataMetadataShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetRegionId(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) SetWorkspaceId(v string) *UpdateCollectionDataMetadataShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateCollectionDataMetadataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
