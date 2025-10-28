// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListK8sSecretsRequest
	GetClusterId() *string
	SetCondition(v string) *ListK8sSecretsRequest
	GetCondition() *string
	SetNamespace(v string) *ListK8sSecretsRequest
	GetNamespace() *string
	SetPageNo(v int32) *ListK8sSecretsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListK8sSecretsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListK8sSecretsRequest
	GetRegionId() *string
	SetShowRelatedApps(v bool) *ListK8sSecretsRequest
	GetShowRelatedApps() *bool
}

type ListK8sSecretsRequest struct {
	// The ID of the cluster.
	//
	// example:
	//
	// 7a953f9a-2946-4c7a-9d82-9939db******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The filter conditions. Set this parameter to a JSON string in the format of {"field":"Name", "pattern":"configmap-"}.
	//
	// example:
	//
	// {\\"field\\":\\"Name\\",\\"pattern\\":\\"product\\"}
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
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
	// The number of entries to return on each page. The value 0 indicates that all entries are returned on one page.
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
	// Specifies whether to return a list of applications that use a Secret.
	//
	// example:
	//
	// true
	ShowRelatedApps *bool `json:"ShowRelatedApps,omitempty" xml:"ShowRelatedApps,omitempty"`
}

func (s ListK8sSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sSecretsRequest) GetCondition() *string {
	return s.Condition
}

func (s *ListK8sSecretsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListK8sSecretsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListK8sSecretsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListK8sSecretsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListK8sSecretsRequest) GetShowRelatedApps() *bool {
	return s.ShowRelatedApps
}

func (s *ListK8sSecretsRequest) SetClusterId(v string) *ListK8sSecretsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListK8sSecretsRequest) SetCondition(v string) *ListK8sSecretsRequest {
	s.Condition = &v
	return s
}

func (s *ListK8sSecretsRequest) SetNamespace(v string) *ListK8sSecretsRequest {
	s.Namespace = &v
	return s
}

func (s *ListK8sSecretsRequest) SetPageNo(v int32) *ListK8sSecretsRequest {
	s.PageNo = &v
	return s
}

func (s *ListK8sSecretsRequest) SetPageSize(v int32) *ListK8sSecretsRequest {
	s.PageSize = &v
	return s
}

func (s *ListK8sSecretsRequest) SetRegionId(v string) *ListK8sSecretsRequest {
	s.RegionId = &v
	return s
}

func (s *ListK8sSecretsRequest) SetShowRelatedApps(v bool) *ListK8sSecretsRequest {
	s.ShowRelatedApps = &v
	return s
}

func (s *ListK8sSecretsRequest) Validate() error {
	return dara.Validate(s)
}
