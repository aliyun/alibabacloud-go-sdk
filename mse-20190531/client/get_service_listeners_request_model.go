// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListenersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetServiceListenersRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *GetServiceListenersRequest
	GetClusterId() *string
	SetClusterName(v string) *GetServiceListenersRequest
	GetClusterName() *string
	SetGroupName(v string) *GetServiceListenersRequest
	GetGroupName() *string
	SetHasIpCount(v string) *GetServiceListenersRequest
	GetHasIpCount() *string
	SetInstanceId(v string) *GetServiceListenersRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *GetServiceListenersRequest
	GetNamespaceId() *string
	SetPageNum(v int32) *GetServiceListenersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *GetServiceListenersRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetServiceListenersRequest
	GetRegionId() *string
	SetRequestPars(v string) *GetServiceListenersRequest
	GetRequestPars() *string
	SetServiceName(v string) *GetServiceListenersRequest
	GetServiceName() *string
}

type GetServiceListenersRequest struct {
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
	// The ID of the MSE instance to which the service belongs.
	//
	// >  You must specify InstanceId or ClusterId.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster to which the service belongs.
	//
	// > The cluster is a concept for Nacos services and is not equivalent to a Microservices Engine (MSE) instance.
	//
	// example:
	//
	// DEFAULT
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The group to which the service belongs.
	//
	// example:
	//
	// WEB_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// none
	HasIpCount *string `json:"HasIpCount,omitempty" xml:"HasIpCount,omitempty"`
	// The ID of the instance to which the service belongs.
	//
	// > You must specify InstanceId or ClusterId.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32odtt20j
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace to which the service belongs.
	//
	// example:
	//
	// ddaf8f12-****-b1c1-86e7c72e266b
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The name of the service whose listeners you want to query.
	//
	// example:
	//
	// zeekr-orderboss
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetServiceListenersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListenersRequest) GoString() string {
	return s.String()
}

func (s *GetServiceListenersRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetServiceListenersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetServiceListenersRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetServiceListenersRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetServiceListenersRequest) GetHasIpCount() *string {
	return s.HasIpCount
}

func (s *GetServiceListenersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetServiceListenersRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetServiceListenersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetServiceListenersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetServiceListenersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceListenersRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *GetServiceListenersRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceListenersRequest) SetAcceptLanguage(v string) *GetServiceListenersRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetServiceListenersRequest) SetClusterId(v string) *GetServiceListenersRequest {
	s.ClusterId = &v
	return s
}

func (s *GetServiceListenersRequest) SetClusterName(v string) *GetServiceListenersRequest {
	s.ClusterName = &v
	return s
}

func (s *GetServiceListenersRequest) SetGroupName(v string) *GetServiceListenersRequest {
	s.GroupName = &v
	return s
}

func (s *GetServiceListenersRequest) SetHasIpCount(v string) *GetServiceListenersRequest {
	s.HasIpCount = &v
	return s
}

func (s *GetServiceListenersRequest) SetInstanceId(v string) *GetServiceListenersRequest {
	s.InstanceId = &v
	return s
}

func (s *GetServiceListenersRequest) SetNamespaceId(v string) *GetServiceListenersRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetServiceListenersRequest) SetPageNum(v int32) *GetServiceListenersRequest {
	s.PageNum = &v
	return s
}

func (s *GetServiceListenersRequest) SetPageSize(v int32) *GetServiceListenersRequest {
	s.PageSize = &v
	return s
}

func (s *GetServiceListenersRequest) SetRegionId(v string) *GetServiceListenersRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceListenersRequest) SetRequestPars(v string) *GetServiceListenersRequest {
	s.RequestPars = &v
	return s
}

func (s *GetServiceListenersRequest) SetServiceName(v string) *GetServiceListenersRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceListenersRequest) Validate() error {
	return dara.Validate(s)
}
