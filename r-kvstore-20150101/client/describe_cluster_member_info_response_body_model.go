// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterMemberInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterChildren(v []*DescribeClusterMemberInfoResponseBodyClusterChildren) *DescribeClusterMemberInfoResponseBody
	GetClusterChildren() []*DescribeClusterMemberInfoResponseBodyClusterChildren
	SetRequestId(v string) *DescribeClusterMemberInfoResponseBody
	GetRequestId() *string
}

type DescribeClusterMemberInfoResponseBody struct {
	// Details about data nodes in the cluster instance.
	ClusterChildren []*DescribeClusterMemberInfoResponseBodyClusterChildren `json:"ClusterChildren,omitempty" xml:"ClusterChildren,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2D9F3768-EDA9-4811-943E-42C8006E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterMemberInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterMemberInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterMemberInfoResponseBody) GetClusterChildren() []*DescribeClusterMemberInfoResponseBodyClusterChildren {
	return s.ClusterChildren
}

func (s *DescribeClusterMemberInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterMemberInfoResponseBody) SetClusterChildren(v []*DescribeClusterMemberInfoResponseBodyClusterChildren) *DescribeClusterMemberInfoResponseBody {
	s.ClusterChildren = v
	return s
}

func (s *DescribeClusterMemberInfoResponseBody) SetRequestId(v string) *DescribeClusterMemberInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterMemberInfoResponseBodyClusterChildren struct {
	// The maximum bandwidth of the node. Unit: MB/s.
	//
	// > This parameter is returned only if the return value of **Service*	- is **redis**, which indicates a data node.
	//
	// example:
	//
	// 96
	BandWidth *int64 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The retention period of binlogs.
	//
	// example:
	//
	// 7
	BinlogRetentionDays *int32 `json:"BinlogRetentionDays,omitempty" xml:"BinlogRetentionDays,omitempty"`
	// The type of workload. The return value is **ALIYUN**.
	//
	// example:
	//
	// ALIYUN
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The maximum memory capacity per data node. Unit: MB.
	//
	// > This parameter is returned only if the return value of **Service*	- is **redis**, which indicates a data node.
	//
	// example:
	//
	// 1024
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The specifications of the data node. For more information, see [Community Edition instances that use cloud disks](https://help.aliyun.com/document_detail/164477.html).
	//
	// example:
	//
	// redis.shard.small.ce
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The maximum number of connections supported by the data node.
	//
	// example:
	//
	// 20000
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The current bandwidth of the data node, which is the sum of the default bandwidth and any extra bandwidth that is purchased. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	CurrentBandWidth *int64 `json:"CurrentBandWidth,omitempty" xml:"CurrentBandWidth,omitempty"`
	// The storage capacity of the [enhanced SSD (ESSD)](https://help.aliyun.com/document_detail/122389.html) that is used by the data node. Unit: MB.
	//
	// > The ESSD is used only to store system operating data, such as Persistent Memory (PMEM). It is not used as a medium to write and read data.
	//
	// example:
	//
	// 4096
	DiskSizeMB *int32 `json:"DiskSizeMB,omitempty" xml:"DiskSizeMB,omitempty"`
	// The ID of the replica set in the node.
	//
	// example:
	//
	// 501791111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data node.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of replica nodes.
	//
	// example:
	//
	// 0
	ReplicaSize *int32 `json:"ReplicaSize,omitempty" xml:"ReplicaSize,omitempty"`
	// The name of the resource group to which the node belongs.
	//
	// example:
	//
	// GLOBAL_ZHANGJIAKOU_A
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The node type. Valid values:
	//
	// 	- **redis**: data node
	//
	// 	- **redis_cs**: config server
	//
	// example:
	//
	// redis
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The major version of the node.
	//
	// example:
	//
	// 5.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 2****_176498472570****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeClusterMemberInfoResponseBodyClusterChildren) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterMemberInfoResponseBodyClusterChildren) GoString() string {
	return s.String()
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetBandWidth() *int64 {
	return s.BandWidth
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetBinlogRetentionDays() *int32 {
	return s.BinlogRetentionDays
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetBizType() *string {
	return s.BizType
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetConnections() *int64 {
	return s.Connections
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetCurrentBandWidth() *int64 {
	return s.CurrentBandWidth
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetDiskSizeMB() *int32 {
	return s.DiskSizeMB
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetId() *int64 {
	return s.Id
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetName() *string {
	return s.Name
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetReplicaSize() *int32 {
	return s.ReplicaSize
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetService() *string {
	return s.Service
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) GetUserId() *string {
	return s.UserId
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetBandWidth(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.BandWidth = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetBinlogRetentionDays(v int32) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.BinlogRetentionDays = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetBizType(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.BizType = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetCapacity(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Capacity = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetClassCode(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.ClassCode = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetConnections(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Connections = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetCurrentBandWidth(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.CurrentBandWidth = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetDiskSizeMB(v int32) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.DiskSizeMB = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetId(v int64) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Id = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetName(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Name = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetReplicaSize(v int32) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.ReplicaSize = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetResourceGroupName(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetService(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.Service = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetServiceVersion(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.ServiceVersion = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) SetUserId(v string) *DescribeClusterMemberInfoResponseBodyClusterChildren {
	s.UserId = &v
	return s
}

func (s *DescribeClusterMemberInfoResponseBodyClusterChildren) Validate() error {
	return dara.Validate(s)
}
