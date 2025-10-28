// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportK8sClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ImportK8sClusterRequest
	GetClusterId() *string
	SetEnableAsm(v bool) *ImportK8sClusterRequest
	GetEnableAsm() *bool
	SetMode(v int32) *ImportK8sClusterRequest
	GetMode() *int32
	SetNamespaceId(v string) *ImportK8sClusterRequest
	GetNamespaceId() *string
}

type ImportK8sClusterRequest struct {
	// The ID of the ACK cluster or serverless Kubernetes cluster. You can obtain the cluster ID by calling the GetK8sCluster operation. For more information, see [GetK8sCluster](https://help.aliyun.com/document_detail/181437.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 9c28bbb9-****-44b3-b953-54ef8a2d0be2
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to enable the integration with Alibaba Cloud Service Mesh (ASM). Valid values:
	//
	// 	- true: Enables the integration with ASM.
	//
	// 	- false: Disables the integration with ASM.
	//
	// example:
	//
	// true
	EnableAsm *bool `json:"EnableAsm,omitempty" xml:"EnableAsm,omitempty"`
	// You can ignore this parameter.
	//
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the namespace. It is in the format of `Region ID:Identifier of the microservices namespace`. Example: `cn-hangzhou:doc`.
	//
	// example:
	//
	// cn-beijing:doc
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ImportK8sClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportK8sClusterRequest) GoString() string {
	return s.String()
}

func (s *ImportK8sClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ImportK8sClusterRequest) GetEnableAsm() *bool {
	return s.EnableAsm
}

func (s *ImportK8sClusterRequest) GetMode() *int32 {
	return s.Mode
}

func (s *ImportK8sClusterRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ImportK8sClusterRequest) SetClusterId(v string) *ImportK8sClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportK8sClusterRequest) SetEnableAsm(v bool) *ImportK8sClusterRequest {
	s.EnableAsm = &v
	return s
}

func (s *ImportK8sClusterRequest) SetMode(v int32) *ImportK8sClusterRequest {
	s.Mode = &v
	return s
}

func (s *ImportK8sClusterRequest) SetNamespaceId(v string) *ImportK8sClusterRequest {
	s.NamespaceId = &v
	return s
}

func (s *ImportK8sClusterRequest) Validate() error {
	return dara.Validate(s)
}
