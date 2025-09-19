// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridProxyClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *ModifyHybridProxyClusterRequest
	GetClusterName() *string
	SetRemark(v string) *ModifyHybridProxyClusterRequest
	GetRemark() *string
}

type ModifyHybridProxyClusterRequest struct {
	// The name of the proxy cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas-test-cnnf
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The remarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ModifyHybridProxyClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridProxyClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridProxyClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ModifyHybridProxyClusterRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyHybridProxyClusterRequest) SetClusterName(v string) *ModifyHybridProxyClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyHybridProxyClusterRequest) SetRemark(v string) *ModifyHybridProxyClusterRequest {
	s.Remark = &v
	return s
}

func (s *ModifyHybridProxyClusterRequest) Validate() error {
	return dara.Validate(s)
}
