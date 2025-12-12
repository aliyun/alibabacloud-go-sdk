// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNodeInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeInfos(v []*DescribeDBClusterNodeInfosResponseBodyNodeInfos) *DescribeDBClusterNodeInfosResponseBody
	GetNodeInfos() []*DescribeDBClusterNodeInfosResponseBodyNodeInfos
	SetPageNumber(v int32) *DescribeDBClusterNodeInfosResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClusterNodeInfosResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDBClusterNodeInfosResponseBody
	GetRequestId() *string
	SetTotalNodeCount(v int32) *DescribeDBClusterNodeInfosResponseBody
	GetTotalNodeCount() *int32
	SetZkNodeInfos(v []*DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) *DescribeDBClusterNodeInfosResponseBody
	GetZkNodeInfos() []*DescribeDBClusterNodeInfosResponseBodyZkNodeInfos
}

type DescribeDBClusterNodeInfosResponseBody struct {
	NodeInfos []*DescribeDBClusterNodeInfosResponseBodyNodeInfos `json:"NodeInfos,omitempty" xml:"NodeInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 746CD303-0B82-5E8D-886D-93A9FAF3A876
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNodeCount *int32                                               `json:"TotalNodeCount,omitempty" xml:"TotalNodeCount,omitempty"`
	ZkNodeInfos    []*DescribeDBClusterNodeInfosResponseBodyZkNodeInfos `json:"ZkNodeInfos,omitempty" xml:"ZkNodeInfos,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNodeInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeInfosResponseBody) GetNodeInfos() []*DescribeDBClusterNodeInfosResponseBodyNodeInfos {
	return s.NodeInfos
}

func (s *DescribeDBClusterNodeInfosResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClusterNodeInfosResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClusterNodeInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterNodeInfosResponseBody) GetTotalNodeCount() *int32 {
	return s.TotalNodeCount
}

func (s *DescribeDBClusterNodeInfosResponseBody) GetZkNodeInfos() []*DescribeDBClusterNodeInfosResponseBodyZkNodeInfos {
	return s.ZkNodeInfos
}

func (s *DescribeDBClusterNodeInfosResponseBody) SetNodeInfos(v []*DescribeDBClusterNodeInfosResponseBodyNodeInfos) *DescribeDBClusterNodeInfosResponseBody {
	s.NodeInfos = v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBody) SetPageNumber(v int32) *DescribeDBClusterNodeInfosResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBody) SetPageSize(v int32) *DescribeDBClusterNodeInfosResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBody) SetRequestId(v string) *DescribeDBClusterNodeInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBody) SetTotalNodeCount(v int32) *DescribeDBClusterNodeInfosResponseBody {
	s.TotalNodeCount = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBody) SetZkNodeInfos(v []*DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) *DescribeDBClusterNodeInfosResponseBody {
	s.ZkNodeInfos = v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBody) Validate() error {
	if s.NodeInfos != nil {
		for _, item := range s.NodeInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ZkNodeInfos != nil {
		for _, item := range s.ZkNodeInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterNodeInfosResponseBodyNodeInfos struct {
	// example:
	//
	// true
	FailoverTesting *bool `json:"FailoverTesting,omitempty" xml:"FailoverTesting,omitempty"`
	// example:
	//
	// 172.168.0.1
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// example:
	//
	// ck-bp108z124a8****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// 1
	ReplicaId *string `json:"ReplicaId,omitempty" xml:"ReplicaId,omitempty"`
	// example:
	//
	// 2
	ShardId *string `json:"ShardId,omitempty" xml:"ShardId,omitempty"`
}

func (s DescribeDBClusterNodeInfosResponseBodyNodeInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeInfosResponseBodyNodeInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) GetFailoverTesting() *bool {
	return s.FailoverTesting
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) GetNodeIp() *string {
	return s.NodeIp
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) GetReplicaId() *string {
	return s.ReplicaId
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) GetShardId() *string {
	return s.ShardId
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) SetFailoverTesting(v bool) *DescribeDBClusterNodeInfosResponseBodyNodeInfos {
	s.FailoverTesting = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) SetNodeIp(v string) *DescribeDBClusterNodeInfosResponseBodyNodeInfos {
	s.NodeIp = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) SetNodeName(v string) *DescribeDBClusterNodeInfosResponseBodyNodeInfos {
	s.NodeName = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) SetReplicaId(v string) *DescribeDBClusterNodeInfosResponseBodyNodeInfos {
	s.ReplicaId = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) SetShardId(v string) *DescribeDBClusterNodeInfosResponseBodyNodeInfos {
	s.ShardId = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBodyNodeInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterNodeInfosResponseBodyZkNodeInfos struct {
	// example:
	//
	// true
	FailoverTesting *bool `json:"FailoverTesting,omitempty" xml:"FailoverTesting,omitempty"`
	// example:
	//
	// ck-bp108z124a8****
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// 1
	ReplicaId *string `json:"ReplicaId,omitempty" xml:"ReplicaId,omitempty"`
}

func (s DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) GetFailoverTesting() *bool {
	return s.FailoverTesting
}

func (s *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) GetReplicaId() *string {
	return s.ReplicaId
}

func (s *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) SetFailoverTesting(v bool) *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos {
	s.FailoverTesting = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) SetNodeName(v string) *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos {
	s.NodeName = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) SetReplicaId(v string) *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos {
	s.ReplicaId = &v
	return s
}

func (s *DescribeDBClusterNodeInfosResponseBodyZkNodeInfos) Validate() error {
	return dara.Validate(s)
}
