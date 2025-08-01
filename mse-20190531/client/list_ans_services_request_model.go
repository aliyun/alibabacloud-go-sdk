// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListAnsServicesRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *ListAnsServicesRequest
	GetClusterId() *string
	SetClusterName(v string) *ListAnsServicesRequest
	GetClusterName() *string
	SetGroupName(v string) *ListAnsServicesRequest
	GetGroupName() *string
	SetHasIpCount(v string) *ListAnsServicesRequest
	GetHasIpCount() *string
	SetInstanceId(v string) *ListAnsServicesRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *ListAnsServicesRequest
	GetNamespaceId() *string
	SetPageNum(v int32) *ListAnsServicesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAnsServicesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListAnsServicesRequest
	GetRegionId() *string
	SetRequestPars(v string) *ListAnsServicesRequest
	GetRequestPars() *string
	SetServiceName(v string) *ListAnsServicesRequest
	GetServiceName() *string
}

type ListAnsServicesRequest struct {
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
	// The ID of the cluster.
	//
	// > This operation contains both the InstanceId and ClusterId parameters. You must specify one of them.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 查询服务下某个集群的实例列表是所需要的参数
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The name of the contact group.
	//
	// example:
	//
	// name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether to query the number of instances that are used for the service.
	//
	// example:
	//
	// true
	HasIpCount *string `json:"HasIpCount,omitempty" xml:"HasIpCount,omitempty"`
	// The ID of the instance.
	//
	// > This operation contains both the InstanceId and ClusterId parameters. You must specify one of them.
	//
	// example:
	//
	// mse-cn-st21v5****
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
	// The number of entries returned per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// name
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListAnsServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnsServicesRequest) GoString() string {
	return s.String()
}

func (s *ListAnsServicesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListAnsServicesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAnsServicesRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListAnsServicesRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAnsServicesRequest) GetHasIpCount() *string {
	return s.HasIpCount
}

func (s *ListAnsServicesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnsServicesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAnsServicesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAnsServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAnsServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAnsServicesRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListAnsServicesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAnsServicesRequest) SetAcceptLanguage(v string) *ListAnsServicesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListAnsServicesRequest) SetClusterId(v string) *ListAnsServicesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAnsServicesRequest) SetClusterName(v string) *ListAnsServicesRequest {
	s.ClusterName = &v
	return s
}

func (s *ListAnsServicesRequest) SetGroupName(v string) *ListAnsServicesRequest {
	s.GroupName = &v
	return s
}

func (s *ListAnsServicesRequest) SetHasIpCount(v string) *ListAnsServicesRequest {
	s.HasIpCount = &v
	return s
}

func (s *ListAnsServicesRequest) SetInstanceId(v string) *ListAnsServicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAnsServicesRequest) SetNamespaceId(v string) *ListAnsServicesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAnsServicesRequest) SetPageNum(v int32) *ListAnsServicesRequest {
	s.PageNum = &v
	return s
}

func (s *ListAnsServicesRequest) SetPageSize(v int32) *ListAnsServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnsServicesRequest) SetRegionId(v string) *ListAnsServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAnsServicesRequest) SetRequestPars(v string) *ListAnsServicesRequest {
	s.RequestPars = &v
	return s
}

func (s *ListAnsServicesRequest) SetServiceName(v string) *ListAnsServicesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListAnsServicesRequest) Validate() error {
	return dara.Validate(s)
}
