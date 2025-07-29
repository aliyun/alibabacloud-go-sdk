// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAmkClusterQuota(v int64) *DescribeUserQuotaResponseBody
	GetAmkClusterQuota() *int64
	SetAskClusterQuota(v int64) *DescribeUserQuotaResponseBody
	GetAskClusterQuota() *int64
	SetClusterNodepoolQuota(v int64) *DescribeUserQuotaResponseBody
	GetClusterNodepoolQuota() *int64
	SetClusterQuota(v int64) *DescribeUserQuotaResponseBody
	GetClusterQuota() *int64
	SetEdgeImprovedNodepoolQuota(v *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) *DescribeUserQuotaResponseBody
	GetEdgeImprovedNodepoolQuota() *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota
	SetNodeQuota(v int64) *DescribeUserQuotaResponseBody
	GetNodeQuota() *int64
	SetQuotas(v map[string]*QuotasValue) *DescribeUserQuotaResponseBody
	GetQuotas() map[string]*QuotasValue
}

type DescribeUserQuotaResponseBody struct {
	// The quota of Container Service for Kubernetes (ACK) managed clusters. Default value: 20. If the default quota limit is reached, submit an application in the [Quota Center console](https://quotas.console.aliyun.com/products/csk/quotas) to increase the quota.
	//
	// example:
	//
	// 20
	AmkClusterQuota *int64 `json:"amk_cluster_quota,omitempty" xml:"amk_cluster_quota,omitempty"`
	// The quota of ACK Serverless clusters. Default value: 20. If the default quota limit is reached, submit an application in the [Quota Center console](https://quotas.console.aliyun.com/products/csk/quotas) to increase the quota.
	//
	// example:
	//
	// 3
	AskClusterQuota *int64 `json:"ask_cluster_quota,omitempty" xml:"ask_cluster_quota,omitempty"`
	// The quota of node pools in an ACK cluster. Default value: 20. If the default quota limit is reached, submit an application in the [Quota Center console](https://quotas.console.aliyun.com/products/csk/quotas) to increase the quota.
	//
	// example:
	//
	// 10
	ClusterNodepoolQuota *int64 `json:"cluster_nodepool_quota,omitempty" xml:"cluster_nodepool_quota,omitempty"`
	// The quota of clusters that belong to an Alibaba Cloud account. Default value: 50. If the default quota limit is reached, submit an application in the [Quota Center console](https://quotas.console.aliyun.com/products/csk/quotas) to increase the quota.
	//
	// example:
	//
	// 50
	ClusterQuota *int64 `json:"cluster_quota,omitempty" xml:"cluster_quota,omitempty"`
	// This parameter is discontinued.
	//
	// The quotas of enhanced edge node pools.
	EdgeImprovedNodepoolQuota *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota `json:"edge_improved_nodepool_quota,omitempty" xml:"edge_improved_nodepool_quota,omitempty" type:"Struct"`
	// The quota of nodes in an ACK cluster. Default value: 100. If the default quota limit is reached, submit an application in the [Quota Center console](https://quotas.console.aliyun.com/products/csk/quotas) to increase the quota.
	//
	// example:
	//
	// 100
	NodeQuota *int64 `json:"node_quota,omitempty" xml:"node_quota,omitempty"`
	// Information about the new quota.
	Quotas map[string]*QuotasValue `json:"quotas,omitempty" xml:"quotas,omitempty"`
}

func (s DescribeUserQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponseBody) GetAmkClusterQuota() *int64 {
	return s.AmkClusterQuota
}

func (s *DescribeUserQuotaResponseBody) GetAskClusterQuota() *int64 {
	return s.AskClusterQuota
}

func (s *DescribeUserQuotaResponseBody) GetClusterNodepoolQuota() *int64 {
	return s.ClusterNodepoolQuota
}

func (s *DescribeUserQuotaResponseBody) GetClusterQuota() *int64 {
	return s.ClusterQuota
}

func (s *DescribeUserQuotaResponseBody) GetEdgeImprovedNodepoolQuota() *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota {
	return s.EdgeImprovedNodepoolQuota
}

func (s *DescribeUserQuotaResponseBody) GetNodeQuota() *int64 {
	return s.NodeQuota
}

func (s *DescribeUserQuotaResponseBody) GetQuotas() map[string]*QuotasValue {
	return s.Quotas
}

func (s *DescribeUserQuotaResponseBody) SetAmkClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.AmkClusterQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetAskClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.AskClusterQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetClusterNodepoolQuota(v int64) *DescribeUserQuotaResponseBody {
	s.ClusterNodepoolQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetClusterQuota(v int64) *DescribeUserQuotaResponseBody {
	s.ClusterQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetEdgeImprovedNodepoolQuota(v *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) *DescribeUserQuotaResponseBody {
	s.EdgeImprovedNodepoolQuota = v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetNodeQuota(v int64) *DescribeUserQuotaResponseBody {
	s.NodeQuota = &v
	return s
}

func (s *DescribeUserQuotaResponseBody) SetQuotas(v map[string]*QuotasValue) *DescribeUserQuotaResponseBody {
	s.Quotas = v
	return s
}

func (s *DescribeUserQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota struct {
	// This parameter is discontinued.
	//
	// The maximum bandwidth of each enhanced edge node pool. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
	// This parameter is discontinued.
	//
	// The maximum number of enhanced edge node pools that you can create within an Alibaba Cloud account.
	//
	// example:
	//
	// 3
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// This parameter is discontinued.
	//
	// The maximum subscription duration of an enhanced edge node pool. Unit: months.
	//
	// >  You are charged for enhanced edge node pools based on the pay-as-you-go billing method. Therefore, you can ignore this parameter.
	//
	// example:
	//
	// 3
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
}

func (s DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) GetCount() *int32 {
	return s.Count
}

func (s *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) SetBandwidth(v int32) *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota {
	s.Bandwidth = &v
	return s
}

func (s *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) SetCount(v int32) *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota {
	s.Count = &v
	return s
}

func (s *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) SetPeriod(v int32) *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota {
	s.Period = &v
	return s
}

func (s *DescribeUserQuotaResponseBodyEdgeImprovedNodepoolQuota) Validate() error {
	return dara.Validate(s)
}
