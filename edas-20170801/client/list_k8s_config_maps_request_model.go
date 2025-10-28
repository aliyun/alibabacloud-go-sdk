// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sConfigMapsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListK8sConfigMapsRequest
	GetClusterId() *string
	SetCondition(v map[string]interface{}) *ListK8sConfigMapsRequest
	GetCondition() map[string]interface{}
	SetNamespace(v string) *ListK8sConfigMapsRequest
	GetNamespace() *string
	SetPageNo(v int32) *ListK8sConfigMapsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListK8sConfigMapsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListK8sConfigMapsRequest
	GetRegionId() *string
	SetShowRelatedApps(v bool) *ListK8sConfigMapsRequest
	GetShowRelatedApps() *bool
}

type ListK8sConfigMapsRequest struct {
	// The ID of the cluster.
	//
	// example:
	//
	// c0830281-366c-41b6-80fb-542e76******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The filter conditions. Set this parameter to a JSON string in the format of {"field":"Name", "pattern":"configmap-"}.
	//
	// example:
	//
	// {"field":"Name", "pattern":"configmap-"}
	Condition map[string]interface{} `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of the page to return. Pages start from Page 0.
	//
	// example:
	//
	// 0
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to return a list of applications that use a ConfigMap. Valid values: true and false.
	//
	// example:
	//
	// true
	ShowRelatedApps *bool `json:"ShowRelatedApps,omitempty" xml:"ShowRelatedApps,omitempty"`
}

func (s ListK8sConfigMapsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListK8sConfigMapsRequest) GoString() string {
	return s.String()
}

func (s *ListK8sConfigMapsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sConfigMapsRequest) GetCondition() map[string]interface{} {
	return s.Condition
}

func (s *ListK8sConfigMapsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListK8sConfigMapsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListK8sConfigMapsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListK8sConfigMapsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListK8sConfigMapsRequest) GetShowRelatedApps() *bool {
	return s.ShowRelatedApps
}

func (s *ListK8sConfigMapsRequest) SetClusterId(v string) *ListK8sConfigMapsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListK8sConfigMapsRequest) SetCondition(v map[string]interface{}) *ListK8sConfigMapsRequest {
	s.Condition = v
	return s
}

func (s *ListK8sConfigMapsRequest) SetNamespace(v string) *ListK8sConfigMapsRequest {
	s.Namespace = &v
	return s
}

func (s *ListK8sConfigMapsRequest) SetPageNo(v int32) *ListK8sConfigMapsRequest {
	s.PageNo = &v
	return s
}

func (s *ListK8sConfigMapsRequest) SetPageSize(v int32) *ListK8sConfigMapsRequest {
	s.PageSize = &v
	return s
}

func (s *ListK8sConfigMapsRequest) SetRegionId(v string) *ListK8sConfigMapsRequest {
	s.RegionId = &v
	return s
}

func (s *ListK8sConfigMapsRequest) SetShowRelatedApps(v bool) *ListK8sConfigMapsRequest {
	s.ShowRelatedApps = &v
	return s
}

func (s *ListK8sConfigMapsRequest) Validate() error {
	return dara.Validate(s)
}
