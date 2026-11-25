// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterInfos(v []*DescribeHybridCloudClustersResponseBodyClusterInfos) *DescribeHybridCloudClustersResponseBody
	GetClusterInfos() []*DescribeHybridCloudClustersResponseBodyClusterInfos
	SetRequestId(v string) *DescribeHybridCloudClustersResponseBody
	GetRequestId() *string
}

type DescribeHybridCloudClustersResponseBody struct {
	// The list of clusters.
	ClusterInfos []*DescribeHybridCloudClustersResponseBodyClusterInfos `json:"ClusterInfos,omitempty" xml:"ClusterInfos,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-ER12-WE34-23PO-301469*****E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridCloudClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClustersResponseBody) GetClusterInfos() []*DescribeHybridCloudClustersResponseBodyClusterInfos {
	return s.ClusterInfos
}

func (s *DescribeHybridCloudClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudClustersResponseBody) SetClusterInfos(v []*DescribeHybridCloudClustersResponseBodyClusterInfos) *DescribeHybridCloudClustersResponseBody {
	s.ClusterInfos = v
	return s
}

func (s *DescribeHybridCloudClustersResponseBody) SetRequestId(v string) *DescribeHybridCloudClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBody) Validate() error {
	if s.ClusterInfos != nil {
		for _, item := range s.ClusterInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridCloudClustersResponseBodyClusterInfos struct {
	// The network access mode. Valid values:
	//
	// - **internet**: Internet access.
	//
	// - **vpc**: Express Connect private network access.
	//
	// example:
	//
	// internet
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The Express Connect access region. Valid values:
	//
	// - **cn-hangzhou**: Hangzhou.
	//
	// - **cn-beiijng**: Beijing.
	//
	// - **cn-shanghai**: Shanghai.
	//
	// example:
	//
	// cn-hangzhou
	AccessRegion *string `json:"AccessRegion,omitempty" xml:"AccessRegion,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// testcluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The resource ID of the hybrid cloud cluster.
	//
	// example:
	//
	// hdbc-cluster-t1****a
	ClusterResourceId *string `json:"ClusterResourceId,omitempty" xml:"ClusterResourceId,omitempty"`
	// The list of HTTP ports. Separate multiple ports with commas (,). Example: port1,port2,port3.
	//
	// example:
	//
	// 80,8080
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The list of HTTPS ports. Separate multiple ports with commas (,). Example: port1,port2,port3.
	//
	// example:
	//
	// 443,8443
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// 524**8
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of protection nodes that can be added to the cluster.
	//
	// example:
	//
	// 1
	ProtectionServerCount *int32 `json:"ProtectionServerCount,omitempty" xml:"ProtectionServerCount,omitempty"`
	// The proxy gateway mode. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// off
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The cluster type. Valid values:
	//
	// - **cname**: Reverse proxy cluster.
	//
	// - **service**: Service cluster.
	//
	// example:
	//
	// cname
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The rule configuration.
	//
	// example:
	//
	// {"enable":true,"param":{"breaker":{"duration":1,"failed":1,"recent_failed":1},"disable_protect":false,"max_request_body_len":1,"timeout":1}}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The manual bypass status. Valid values:
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// example:
	//
	// off
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The rule type. Valid values:
	//
	// - **bypass**: Skips security checks and directly allows requests.
	//
	// example:
	//
	// bypass
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeHybridCloudClustersResponseBodyClusterInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClustersResponseBodyClusterInfos) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetAccessMode() *string {
	return s.AccessMode
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetAccessRegion() *string {
	return s.AccessRegion
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetClusterResourceId() *string {
	return s.ClusterResourceId
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetHttpsPorts() *string {
	return s.HttpsPorts
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetId() *int64 {
	return s.Id
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetProtectionServerCount() *int32 {
	return s.ProtectionServerCount
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetRemark() *string {
	return s.Remark
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetAccessMode(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.AccessMode = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetAccessRegion(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.AccessRegion = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetClusterName(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.ClusterName = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetClusterResourceId(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.ClusterResourceId = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetHttpPorts(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.HttpPorts = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetHttpsPorts(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.HttpsPorts = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetId(v int64) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.Id = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetProtectionServerCount(v int32) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.ProtectionServerCount = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetProxyStatus(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.ProxyStatus = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetProxyType(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.ProxyType = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetRemark(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.Remark = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetRuleConfig(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.RuleConfig = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetRuleStatus(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.RuleStatus = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) SetRuleType(v string) *DescribeHybridCloudClustersResponseBodyClusterInfos {
	s.RuleType = &v
	return s
}

func (s *DescribeHybridCloudClustersResponseBodyClusterInfos) Validate() error {
	return dara.Validate(s)
}
