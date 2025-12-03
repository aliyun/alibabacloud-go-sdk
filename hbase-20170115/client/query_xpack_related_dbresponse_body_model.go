// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryXpackRelatedDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v *QueryXpackRelatedDBResponseBodyClusterList) *QueryXpackRelatedDBResponseBody
	GetClusterList() *QueryXpackRelatedDBResponseBodyClusterList
	SetPageNumber(v int32) *QueryXpackRelatedDBResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryXpackRelatedDBResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryXpackRelatedDBResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryXpackRelatedDBResponseBody
	GetTotalCount() *int32
}

type QueryXpackRelatedDBResponseBody struct {
	ClusterList *QueryXpackRelatedDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryXpackRelatedDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelatedDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBResponseBody) GetClusterList() *QueryXpackRelatedDBResponseBodyClusterList {
	return s.ClusterList
}

func (s *QueryXpackRelatedDBResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryXpackRelatedDBResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryXpackRelatedDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryXpackRelatedDBResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryXpackRelatedDBResponseBody) SetClusterList(v *QueryXpackRelatedDBResponseBodyClusterList) *QueryXpackRelatedDBResponseBody {
	s.ClusterList = v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) SetPageNumber(v int32) *QueryXpackRelatedDBResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) SetPageSize(v int32) *QueryXpackRelatedDBResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) SetRequestId(v string) *QueryXpackRelatedDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) SetTotalCount(v int32) *QueryXpackRelatedDBResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBody) Validate() error {
	if s.ClusterList != nil {
		if err := s.ClusterList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryXpackRelatedDBResponseBodyClusterList struct {
	Cluster []*QueryXpackRelatedDBResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QueryXpackRelatedDBResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelatedDBResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBResponseBodyClusterList) GetCluster() []*QueryXpackRelatedDBResponseBodyClusterListCluster {
	return s.Cluster
}

func (s *QueryXpackRelatedDBResponseBodyClusterList) SetCluster(v []*QueryXpackRelatedDBResponseBodyClusterListCluster) *QueryXpackRelatedDBResponseBodyClusterList {
	s.Cluster = v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterList) Validate() error {
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

type QueryXpackRelatedDBResponseBodyClusterListCluster struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DBType      *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion   *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	IsRelated   *bool   `json:"IsRelated,omitempty" xml:"IsRelated,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryXpackRelatedDBResponseBodyClusterListCluster) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelatedDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) GetDBType() *string {
	return s.DBType
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) GetDBVersion() *string {
	return s.DBVersion
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) GetIsRelated() *bool {
	return s.IsRelated
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) GetStatus() *string {
	return s.Status
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetClusterId(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetClusterName(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetDBType(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.DBType = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetDBVersion(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.DBVersion = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetIsRelated(v bool) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.IsRelated = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) SetStatus(v string) *QueryXpackRelatedDBResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

func (s *QueryXpackRelatedDBResponseBodyClusterListCluster) Validate() error {
	return dara.Validate(s)
}
