// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCreateIndexJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CancelCreateIndexJobRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CancelCreateIndexJobRequest
	GetDBInstanceId() *string
	SetJobId(v string) *CancelCreateIndexJobRequest
	GetJobId() *string
	SetNamespace(v string) *CancelCreateIndexJobRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *CancelCreateIndexJobRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *CancelCreateIndexJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelCreateIndexJobRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CancelCreateIndexJobRequest
	GetWorkspaceId() *string
}

type CancelCreateIndexJobRequest struct {
	// The collection name.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to view the collection list.
	//
	// This parameter is required.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in the specified region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the index creation task, which is returned by the CreateIndex operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace.
	//
	// > You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace, or call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to view the namespace list.
	//
	// This parameter is required.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
	//
	// > This value is specified when you call the CreateNamespace operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workspace that consists of multiple database instances. This parameter and the DBInstanceId parameter cannot be both empty. If both parameters are specified, this parameter takes precedence.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CancelCreateIndexJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCreateIndexJobRequest) GoString() string {
	return s.String()
}

func (s *CancelCreateIndexJobRequest) GetCollection() *string {
	return s.Collection
}

func (s *CancelCreateIndexJobRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CancelCreateIndexJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelCreateIndexJobRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CancelCreateIndexJobRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *CancelCreateIndexJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelCreateIndexJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelCreateIndexJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CancelCreateIndexJobRequest) SetCollection(v string) *CancelCreateIndexJobRequest {
	s.Collection = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetDBInstanceId(v string) *CancelCreateIndexJobRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetJobId(v string) *CancelCreateIndexJobRequest {
	s.JobId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetNamespace(v string) *CancelCreateIndexJobRequest {
	s.Namespace = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetNamespacePassword(v string) *CancelCreateIndexJobRequest {
	s.NamespacePassword = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetOwnerId(v int64) *CancelCreateIndexJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetRegionId(v string) *CancelCreateIndexJobRequest {
	s.RegionId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) SetWorkspaceId(v string) *CancelCreateIndexJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CancelCreateIndexJobRequest) Validate() error {
	return dara.Validate(s)
}
