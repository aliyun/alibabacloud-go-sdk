// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryXpackRelateDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v *QueryXpackRelateDBResponseBodyClusterList) *QueryXpackRelateDBResponseBody
	GetClusterList() *QueryXpackRelateDBResponseBodyClusterList
	SetRequestId(v string) *QueryXpackRelateDBResponseBody
	GetRequestId() *string
}

type QueryXpackRelateDBResponseBody struct {
	ClusterList *QueryXpackRelateDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	// example:
	//
	// 288E9010-36DD-499C-B4DA-61E4362DA4CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryXpackRelateDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelateDBResponseBody) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBody) GetClusterList() *QueryXpackRelateDBResponseBodyClusterList {
	return s.ClusterList
}

func (s *QueryXpackRelateDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryXpackRelateDBResponseBody) SetClusterList(v *QueryXpackRelateDBResponseBodyClusterList) *QueryXpackRelateDBResponseBody {
	s.ClusterList = v
	return s
}

func (s *QueryXpackRelateDBResponseBody) SetRequestId(v string) *QueryXpackRelateDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryXpackRelateDBResponseBody) Validate() error {
	if s.ClusterList != nil {
		if err := s.ClusterList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryXpackRelateDBResponseBodyClusterList struct {
	Cluster []*QueryXpackRelateDBResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s QueryXpackRelateDBResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelateDBResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBodyClusterList) GetCluster() []*QueryXpackRelateDBResponseBodyClusterListCluster {
	return s.Cluster
}

func (s *QueryXpackRelateDBResponseBodyClusterList) SetCluster(v []*QueryXpackRelateDBResponseBodyClusterListCluster) *QueryXpackRelateDBResponseBodyClusterList {
	s.Cluster = v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterList) Validate() error {
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

type QueryXpackRelateDBResponseBodyClusterListCluster struct {
	// example:
	//
	// hb-bp16o0pd52e3y582s
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// hbase_test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// hbase
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 2.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// false
	IsRelated *bool `json:"IsRelated,omitempty" xml:"IsRelated,omitempty"`
	// example:
	//
	// ..
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryXpackRelateDBResponseBodyClusterListCluster) String() string {
	return dara.Prettify(s)
}

func (s QueryXpackRelateDBResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) GetDBType() *string {
	return s.DBType
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) GetDBVersion() *string {
	return s.DBVersion
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) GetIsRelated() *bool {
	return s.IsRelated
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) GetStatus() *string {
	return s.Status
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetClusterId(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetClusterName(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetDBType(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.DBType = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetDBVersion(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.DBVersion = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetIsRelated(v bool) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.IsRelated = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetLockMode(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.LockMode = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) SetStatus(v string) *QueryXpackRelateDBResponseBodyClusterListCluster {
	s.Status = &v
	return s
}

func (s *QueryXpackRelateDBResponseBodyClusterListCluster) Validate() error {
	return dara.Validate(s)
}
