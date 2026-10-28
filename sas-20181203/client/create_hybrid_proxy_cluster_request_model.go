// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridProxyClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *CreateHybridProxyClusterRequest
	GetClusterName() *string
	SetIp(v string) *CreateHybridProxyClusterRequest
	GetIp() *string
	SetRemark(v string) *CreateHybridProxyClusterRequest
	GetRemark() *string
}

type CreateHybridProxyClusterRequest struct {
	// The name of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// Chester-Test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The endpoint of the cluster.
	//
	// >  You can specify an IP address or a domain name
	//
	// example:
	//
	// 192.168.xx.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// remark test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateHybridProxyClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridProxyClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridProxyClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateHybridProxyClusterRequest) GetIp() *string {
	return s.Ip
}

func (s *CreateHybridProxyClusterRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateHybridProxyClusterRequest) SetClusterName(v string) *CreateHybridProxyClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateHybridProxyClusterRequest) SetIp(v string) *CreateHybridProxyClusterRequest {
	s.Ip = &v
	return s
}

func (s *CreateHybridProxyClusterRequest) SetRemark(v string) *CreateHybridProxyClusterRequest {
	s.Remark = &v
	return s
}

func (s *CreateHybridProxyClusterRequest) Validate() error {
	return dara.Validate(s)
}
