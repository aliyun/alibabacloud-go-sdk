// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateK8sResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrivateK8sInfos(v []*ListPrivateK8sResponseBodyPrivateK8sInfos) *ListPrivateK8sResponseBody
	GetPrivateK8sInfos() []*ListPrivateK8sResponseBodyPrivateK8sInfos
	SetRequestId(v string) *ListPrivateK8sResponseBody
	GetRequestId() *string
}

type ListPrivateK8sResponseBody struct {
	// The information about the self-managed Kubernetes clusters.
	PrivateK8sInfos []*ListPrivateK8sResponseBodyPrivateK8sInfos `json:"PrivateK8sInfos,omitempty" xml:"PrivateK8sInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 52870893-48A7-5A9E-9E05-6253E5B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivateK8sResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateK8sResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateK8sResponseBody) GetPrivateK8sInfos() []*ListPrivateK8sResponseBodyPrivateK8sInfos {
	return s.PrivateK8sInfos
}

func (s *ListPrivateK8sResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrivateK8sResponseBody) SetPrivateK8sInfos(v []*ListPrivateK8sResponseBodyPrivateK8sInfos) *ListPrivateK8sResponseBody {
	s.PrivateK8sInfos = v
	return s
}

func (s *ListPrivateK8sResponseBody) SetRequestId(v string) *ListPrivateK8sResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateK8sResponseBody) Validate() error {
	if s.PrivateK8sInfos != nil {
		for _, item := range s.PrivateK8sInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrivateK8sResponseBodyPrivateK8sInfos struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 12345
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The IP address of the API server.
	//
	// example:
	//
	// 192.168.XX.XX
	ApiServerIp *string `json:"ApiServerIp,omitempty" xml:"ApiServerIp,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// xxx
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The version of Kubernetes.
	//
	// example:
	//
	// 1.18
	K8sVersion *string `json:"K8sVersion,omitempty" xml:"K8sVersion,omitempty"`
	// The server configuration of Kubernetes.
	//
	// example:
	//
	// xxx
	KubeConfig *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	// The network type. Valid values:
	//
	// 	- **1**: Internet.
	//
	// 	- **2**: VPC.
	//
	// example:
	//
	// 1
	NetType *int64 `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2zet5l358k6z0gnz*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListPrivateK8sResponseBodyPrivateK8sInfos) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateK8sResponseBodyPrivateK8sInfos) GoString() string {
	return s.String()
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetApiServerIp() *string {
	return s.ApiServerIp
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetId() *int64 {
	return s.Id
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetK8sVersion() *string {
	return s.K8sVersion
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetKubeConfig() *string {
	return s.KubeConfig
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetNetType() *int64 {
	return s.NetType
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) GetVpcId() *string {
	return s.VpcId
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetAliUid(v int64) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.AliUid = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetApiServerIp(v string) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.ApiServerIp = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetClusterName(v string) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.ClusterName = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetId(v int64) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.Id = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetK8sVersion(v string) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.K8sVersion = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetKubeConfig(v string) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.KubeConfig = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetNetType(v int64) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.NetType = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetRegionId(v string) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.RegionId = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) SetVpcId(v string) *ListPrivateK8sResponseBodyPrivateK8sInfos {
	s.VpcId = &v
	return s
}

func (s *ListPrivateK8sResponseBodyPrivateK8sInfos) Validate() error {
	return dara.Validate(s)
}
