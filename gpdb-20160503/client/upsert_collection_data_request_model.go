// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertCollectionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *UpsertCollectionDataRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UpsertCollectionDataRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *UpsertCollectionDataRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UpsertCollectionDataRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UpsertCollectionDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpsertCollectionDataRequest
	GetRegionId() *string
	SetRows(v []*UpsertCollectionDataRequestRows) *UpsertCollectionDataRequest
	GetRows() []*UpsertCollectionDataRequestRows
	SetWorkspaceId(v string) *UpsertCollectionDataRequest
	GetWorkspaceId() *string
}

type UpsertCollectionDataRequest struct {
	// The name of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
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
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vector data that you want to upload.
	Rows []*UpsertCollectionDataRequestRows `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. You must specify one of the WorkspaceId and DBInstanceId parameters. If you specify both parameters, the WorkspaceId parameter takes effect.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpsertCollectionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataRequest) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataRequest) GetCollection() *string {
	return s.Collection
}

func (s *UpsertCollectionDataRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpsertCollectionDataRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpsertCollectionDataRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UpsertCollectionDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpsertCollectionDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpsertCollectionDataRequest) GetRows() []*UpsertCollectionDataRequestRows {
	return s.Rows
}

func (s *UpsertCollectionDataRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpsertCollectionDataRequest) SetCollection(v string) *UpsertCollectionDataRequest {
	s.Collection = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetDBInstanceId(v string) *UpsertCollectionDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetNamespace(v string) *UpsertCollectionDataRequest {
	s.Namespace = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetNamespacePassword(v string) *UpsertCollectionDataRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetOwnerId(v int64) *UpsertCollectionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetRegionId(v string) *UpsertCollectionDataRequest {
	s.RegionId = &v
	return s
}

func (s *UpsertCollectionDataRequest) SetRows(v []*UpsertCollectionDataRequestRows) *UpsertCollectionDataRequest {
	s.Rows = v
	return s
}

func (s *UpsertCollectionDataRequest) SetWorkspaceId(v string) *UpsertCollectionDataRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpsertCollectionDataRequest) Validate() error {
	if s.Rows != nil {
		for _, item := range s.Rows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpsertCollectionDataRequestRows struct {
	// The unique ID of the vector data.
	//
	// >  If you leave this parameter empty, a unique ID is automatically generated by using uuidgen. If you specify a value that is the same as an existing ID, an updated vector data ID is used.
	//
	// example:
	//
	// doca-1234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metadata of the vector data. The key field of this parameter must be the same as the key field of Metadata that is specified when you call the [CreateCollection](https://help.aliyun.com/document_detail/2401497.html) operation.
	Metadata map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The sparse vector data.
	SparseVector *UpsertCollectionDataRequestRowsSparseVector `json:"SparseVector,omitempty" xml:"SparseVector,omitempty" type:"Struct"`
	// The vector data. The length of the vector data must be the same as the value of Dimension that is specified when you call the [CreateCollection](https://help.aliyun.com/document_detail/2401497.html) operation.
	//
	// This parameter is required.
	Vector []*float64 `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Repeated"`
}

func (s UpsertCollectionDataRequestRows) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataRequestRows) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataRequestRows) GetId() *string {
	return s.Id
}

func (s *UpsertCollectionDataRequestRows) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *UpsertCollectionDataRequestRows) GetSparseVector() *UpsertCollectionDataRequestRowsSparseVector {
	return s.SparseVector
}

func (s *UpsertCollectionDataRequestRows) GetVector() []*float64 {
	return s.Vector
}

func (s *UpsertCollectionDataRequestRows) SetId(v string) *UpsertCollectionDataRequestRows {
	s.Id = &v
	return s
}

func (s *UpsertCollectionDataRequestRows) SetMetadata(v map[string]*string) *UpsertCollectionDataRequestRows {
	s.Metadata = v
	return s
}

func (s *UpsertCollectionDataRequestRows) SetSparseVector(v *UpsertCollectionDataRequestRowsSparseVector) *UpsertCollectionDataRequestRows {
	s.SparseVector = v
	return s
}

func (s *UpsertCollectionDataRequestRows) SetVector(v []*float64) *UpsertCollectionDataRequestRows {
	s.Vector = v
	return s
}

func (s *UpsertCollectionDataRequestRows) Validate() error {
	if s.SparseVector != nil {
		if err := s.SparseVector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertCollectionDataRequestRowsSparseVector struct {
	// The indices.
	//
	// >  The size of indices cannot exceed 4000.
	Indices []*int64 `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Repeated"`
	// The sparse vector values.
	Values []*float64 `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s UpsertCollectionDataRequestRowsSparseVector) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataRequestRowsSparseVector) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataRequestRowsSparseVector) GetIndices() []*int64 {
	return s.Indices
}

func (s *UpsertCollectionDataRequestRowsSparseVector) GetValues() []*float64 {
	return s.Values
}

func (s *UpsertCollectionDataRequestRowsSparseVector) SetIndices(v []*int64) *UpsertCollectionDataRequestRowsSparseVector {
	s.Indices = v
	return s
}

func (s *UpsertCollectionDataRequestRowsSparseVector) SetValues(v []*float64) *UpsertCollectionDataRequestRowsSparseVector {
	s.Values = v
	return s
}

func (s *UpsertCollectionDataRequestRowsSparseVector) Validate() error {
	return dara.Validate(s)
}
