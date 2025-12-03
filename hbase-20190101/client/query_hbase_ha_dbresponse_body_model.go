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
	SetTotalCount(v int64) *QueryHBaseHaDBResponseBody
	GetTotalCount() *int64
}

type QueryHBaseHaDBResponseBody struct {
	ClusterList *QueryHBaseHaDBResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 963355AD-A3B1-4654-AFFC-B5186EB8F889
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *QueryHBaseHaDBResponseBody) GetTotalCount() *int64 {
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

func (s *QueryHBaseHaDBResponseBody) SetTotalCount(v int64) *QueryHBaseHaDBResponseBody {
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
	// example:
	//
	// hb-t4nn7dy1u1etbzmzm
	ActiveName *string `json:"ActiveName,omitempty" xml:"ActiveName,omitempty"`
	// bdsId
	//
	// example:
	//
	// bds-t4n3496whj23ia4k
	BdsName *string `json:"BdsName,omitempty" xml:"BdsName,omitempty"`
	// example:
	//
	// ha-v21tmnxjwh2yuy1il
	HaName        *string                                                    `json:"HaName,omitempty" xml:"HaName,omitempty"`
	HaSlbConnList *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList `json:"HaSlbConnList,omitempty" xml:"HaSlbConnList,omitempty" type:"Struct"`
	// example:
	//
	// hb-t4n0ye37832tx22vz
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

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) GetHaSlbConnList() *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList {
	return s.HaSlbConnList
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

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetHaSlbConnList(v *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.HaSlbConnList = v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) SetStandbyName(v string) *QueryHBaseHaDBResponseBodyClusterListCluster {
	s.StandbyName = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListCluster) Validate() error {
	if s.HaSlbConnList != nil {
		if err := s.HaSlbConnList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList struct {
	HaSlbConn []*QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn `json:"HaSlbConn,omitempty" xml:"HaSlbConn,omitempty" type:"Repeated"`
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) String() string {
	return dara.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) GetHaSlbConn() []*QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	return s.HaSlbConn
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) SetHaSlbConn(v []*QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList {
	s.HaSlbConn = v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnList) Validate() error {
	if s.HaSlbConn != nil {
		for _, item := range s.HaSlbConn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn struct {
	// example:
	//
	// Standby
	HbaseType *string `json:"HbaseType,omitempty" xml:"HbaseType,omitempty"`
	// example:
	//
	// ha-v21tmnxjwh2yuy1il-phoenix.bds.9b78df04-b.rds.aliyuncs.com:8765
	SlbConnAddr *string `json:"SlbConnAddr,omitempty" xml:"SlbConnAddr,omitempty"`
	// example:
	//
	// phoenix
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) String() string {
	return dara.Prettify(s)
}

func (s QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) GetHbaseType() *string {
	return s.HbaseType
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) GetSlbConnAddr() *string {
	return s.SlbConnAddr
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) GetSlbType() *string {
	return s.SlbType
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetHbaseType(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.HbaseType = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetSlbConnAddr(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.SlbConnAddr = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) SetSlbType(v string) *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn {
	s.SlbType = &v
	return s
}

func (s *QueryHBaseHaDBResponseBodyClusterListClusterHaSlbConnListHaSlbConn) Validate() error {
	return dara.Validate(s)
}
