// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListAnsInstancesRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *ListAnsInstancesRequest
	GetClusterId() *string
	SetClusterName(v string) *ListAnsInstancesRequest
	GetClusterName() *string
	SetGroupName(v string) *ListAnsInstancesRequest
	GetGroupName() *string
	SetInstanceId(v string) *ListAnsInstancesRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *ListAnsInstancesRequest
	GetNamespaceId() *string
	SetPageNum(v int32) *ListAnsInstancesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAnsInstancesRequest
	GetPageSize() *int32
	SetRequestPars(v string) *ListAnsInstancesRequest
	GetRequestPars() *string
	SetServiceName(v string) *ListAnsInstancesRequest
	GetServiceName() *string
}

type ListAnsInstancesRequest struct {
	// The language of the returned information. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Deprecated
	//
	// The cluster ID.	Notice: If this operation includes both InstanceId and ClusterId, you must specify at least one of them.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster alias.
	//
	// example:
	//
	// DEFAULT
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The contact group name.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The instance ID.	Notice: If this operation includes both InstanceId and ClusterId, you must specify at least one of them.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl327w****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// 12233****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of instances to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The extended request parameters in JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListAnsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListAnsInstancesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListAnsInstancesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAnsInstancesRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListAnsInstancesRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListAnsInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAnsInstancesRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAnsInstancesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAnsInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAnsInstancesRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *ListAnsInstancesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAnsInstancesRequest) SetAcceptLanguage(v string) *ListAnsInstancesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListAnsInstancesRequest) SetClusterId(v string) *ListAnsInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAnsInstancesRequest) SetClusterName(v string) *ListAnsInstancesRequest {
	s.ClusterName = &v
	return s
}

func (s *ListAnsInstancesRequest) SetGroupName(v string) *ListAnsInstancesRequest {
	s.GroupName = &v
	return s
}

func (s *ListAnsInstancesRequest) SetInstanceId(v string) *ListAnsInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAnsInstancesRequest) SetNamespaceId(v string) *ListAnsInstancesRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListAnsInstancesRequest) SetPageNum(v int32) *ListAnsInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *ListAnsInstancesRequest) SetPageSize(v int32) *ListAnsInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAnsInstancesRequest) SetRequestPars(v string) *ListAnsInstancesRequest {
	s.RequestPars = &v
	return s
}

func (s *ListAnsInstancesRequest) SetServiceName(v string) *ListAnsInstancesRequest {
	s.ServiceName = &v
	return s
}

func (s *ListAnsInstancesRequest) Validate() error {
	return dara.Validate(s)
}
