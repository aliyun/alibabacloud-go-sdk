// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsServiceClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListAnsServiceClustersRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *ListAnsServiceClustersRequest
	GetClusterId() *string
	SetClusterName(v string) *ListAnsServiceClustersRequest
	GetClusterName() *string
	SetGroupName(v string) *ListAnsServiceClustersRequest
	GetGroupName() *string
	SetInstanceId(v string) *ListAnsServiceClustersRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *ListAnsServiceClustersRequest
	GetNamespaceId() *string
	SetPageNum(v int32) *ListAnsServiceClustersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAnsServiceClustersRequest
	GetPageSize() *int32
	SetServiceName(v string) *ListAnsServiceClustersRequest
	GetServiceName() *string
}

type ListAnsServiceClustersRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Deprecated
	//
	// The ID of the MSE cluster.
	//
	// >  The MSE cluster is different from the cluster of the Nacos service.
	//
	// example:
	//
	// mse-09k1q110q01
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The alias of the cluster.
	//
	// example:
	//
	// mse-7413****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32a6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// 12233****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// nacos.test.3
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListAnsServiceClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServiceClustersRequest) GoString() string {
	return s.String()
}

func (s *ListAnsServiceClustersRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListAnsServiceClustersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAnsServiceClustersRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListAnsServiceClustersRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAnsServiceClustersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnsServiceClustersRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAnsServiceClustersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAnsServiceClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAnsServiceClustersRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAnsServiceClustersRequest) SetAcceptLanguage(v string) *ListAnsServiceClustersRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetClusterId(v string) *ListAnsServiceClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetClusterName(v string) *ListAnsServiceClustersRequest {
	s.ClusterName = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetGroupName(v string) *ListAnsServiceClustersRequest {
	s.GroupName = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetInstanceId(v string) *ListAnsServiceClustersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetNamespaceId(v string) *ListAnsServiceClustersRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetPageNum(v int32) *ListAnsServiceClustersRequest {
	s.PageNum = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetPageSize(v int32) *ListAnsServiceClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnsServiceClustersRequest) SetServiceName(v string) *ListAnsServiceClustersRequest {
	s.ServiceName = &v
	return s
}

func (s *ListAnsServiceClustersRequest) Validate() error {
	return dara.Validate(s)
}
