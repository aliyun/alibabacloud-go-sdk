// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClustersRequest
	GetClusterId() *string
	SetClusterStatus(v string) *ListClustersRequest
	GetClusterStatus() *string
	SetClusterType(v string) *ListClustersRequest
	GetClusterType() *string
	SetCurrent(v int64) *ListClustersRequest
	GetCurrent() *int64
	SetId(v string) *ListClustersRequest
	GetId() *string
	SetName(v string) *ListClustersRequest
	GetName() *string
	SetPageSize(v int64) *ListClustersRequest
	GetPageSize() *int64
}

type ListClustersRequest struct {
	// Filter by cluster ID
	//
	// > This cluster ID is not the ACK cluster ID, but the `id` field in the data returned by this API
	//
	// example:
	//
	// cb7d4cc26c8f845fb8a8255ffd394820e
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// - `Running`: Cluster management is Normal;
	//
	// - `Installing`: An install Job is in progress for the cluster;
	//
	// - `Uninstalling`: An uninstall Job is in progress for the cluster;
	//
	// - `Upgrading`: An Update Job is in progress for the cluster;
	//
	// - `Offline`: The cluster is offline and management is abnormal.
	//
	// example:
	//
	// Running
	ClusterStatus *string `json:"cluster_status,omitempty" xml:"cluster_status,omitempty"`
	// - `ACK`: ACK cluster
	//
	// - `CUSTOM`: Custom cluster (default clusters are classified as custom clusters)
	//
	// example:
	//
	// ACK
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// Current page number (starting from page 1)
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// This field is deprecated. Use the `cluster_id` field for filtering instead.
	//
	// example:
	//
	// cb7d4cc26c8f845fb8a8255ffd394820e
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Filter plugins by plugin name
	//
	// example:
	//
	// proxy-next-upstream
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClustersRequest) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *ListClustersRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClustersRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListClustersRequest) GetId() *string {
	return s.Id
}

func (s *ListClustersRequest) GetName() *string {
	return s.Name
}

func (s *ListClustersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListClustersRequest) SetClusterId(v string) *ListClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClustersRequest) SetClusterStatus(v string) *ListClustersRequest {
	s.ClusterStatus = &v
	return s
}

func (s *ListClustersRequest) SetClusterType(v string) *ListClustersRequest {
	s.ClusterType = &v
	return s
}

func (s *ListClustersRequest) SetCurrent(v int64) *ListClustersRequest {
	s.Current = &v
	return s
}

func (s *ListClustersRequest) SetId(v string) *ListClustersRequest {
	s.Id = &v
	return s
}

func (s *ListClustersRequest) SetName(v string) *ListClustersRequest {
	s.Name = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int64) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) Validate() error {
	return dara.Validate(s)
}
