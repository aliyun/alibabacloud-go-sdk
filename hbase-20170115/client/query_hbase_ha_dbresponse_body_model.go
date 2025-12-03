// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHBaseHaDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v *QueryHBaseHaDBResponseBodyClusterList) *QueryHBaseHaDBResponseBody
	GetClusterList() *QueryHBaseHaDBResponseBodyClusterList
	SetPageNumber(v int32) *QueryHBaseHaDBResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryHBaseHaDBResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryHBaseHaDBResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryHBaseHaDBResponseBody
	GetTotalCount() *int32
}

type QueryHBaseHaDBResponseBody struct {
	ClusterList *QueryHBaseHaDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	PageNumber  *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryHBaseHaDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHBaseHaDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBody) GetClusterList() *QueryHBaseHaDBResponseBodyClusterList {
	return s.ClusterList
}

func (s *QueryHBaseHaDBResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryHBaseHaDBResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryHBaseHaDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHBaseHaDBResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryHBaseHaDBResponseBody) SetClusterList(v *QueryHBaseHaDBResponseBodyClusterList) *QueryHBaseHaDBResponseBody {
	s.ClusterList = v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetPageNumber(v int32) *QueryHBaseHaDBResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetPageSize(v int32) *QueryHBaseHaDBResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetRequestId(v string) *QueryHBaseHaDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) SetTotalCount(v int32) *QueryHBaseHaDBResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHBaseHaDBResponseBody) Validate() error {
	if s.ClusterList != nil {
		if err := s.ClusterList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHBaseHaDBResponseBodyClusterList struct {
	Cluster []*QueryHBaseHaDBResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QueryHBaseHaDBResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterList) GetCluster() []*QueryHBaseHaDBResponseBodyClusterListCluster {
	return s.Cluster
}

func (s *QueryHBaseHaDBResponseBodyClusterList) SetCluster(v []*QueryHBaseHaDBResponseBodyClusterListCluster) *QueryHBaseHaDBResponseBodyClusterList {
	s.Cluster = v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterList) Validate() error {
	if s.Cluster != nil {
		for _, item := range s.Cluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHBaseHaDBResponseBodyClusterListCluster struct {
	ActiveName  *string `json:"ActiveName,omitempty" xml:"ActiveName,omitempty"`
	BdsName     *string `json:"BdsName,omitempty" xml:"BdsName,omitempty"`
	HaName      *string `json:"HaName,omitempty" xml:"HaName,omitempty"`
	StandbyName *string `json:"StandbyName,omitempty" xml:"StandbyName,omitempty"`
}

func (s QueryHBaseHaDBResponseBodyClusterListCluster) String() string {
	return dara.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) GetActiveName() *string {
	return s.ActiveName
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) GetBdsName() *string {
	return s.BdsName
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) GetHaName() *string {
	return s.HaName
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) GetStandbyName() *string {
	return s.StandbyName
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetActiveName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.ActiveName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetBdsName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.BdsName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetHaName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.HaName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetStandbyName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.StandbyName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) Validate() error {
	return dara.Validate(s)
}
