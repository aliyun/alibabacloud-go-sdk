// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody
	GetClusters() []*ListClustersResponseBodyClusters
	SetCurrentPage(v int32) *ListClustersResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListClustersResponseBody
	GetTotal() *int32
}

type ListClustersResponseBody struct {
	// The clusters.
	Clusters []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of clusters.
	//
	// example:
	//
	// 114
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetClusters() []*ListClustersResponseBodyClusters {
	return s.Clusters
}

func (s *ListClustersResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetCurrentPage(v int32) *ListClustersResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetTotal(v int32) *ListClustersResponseBody {
	s.Total = &v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClustersResponseBodyClusters struct {
	// The ID of the cluster.
	//
	// example:
	//
	// cluster-w3G9vOJI2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- NEW: The cluster is not initialized.
	//
	// 	- INITIALIZED: The cluster is initialized.
	//
	// 	- DELETED: The cluster is deleted.
	//
	// 	- SYNCHRONIZING: The cluster is being synchronized.
	//
	// 	- TO_DELETE: The cluster is pending deletion.
	//
	// example:
	//
	// INITIALIZED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClustersResponseBodyClusters) GetStatus() *string {
	return s.Status
}

func (s *ListClustersResponseBodyClusters) SetClusterId(v string) *ListClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetStatus(v string) *ListClustersResponseBodyClusters {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}
