// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroupName(v string) *ListGroupsRequest
	GetAppGroupName() *string
	SetNamespace(v string) *ListGroupsRequest
	GetNamespace() *string
	SetNamespaceSource(v string) *ListGroupsRequest
	GetNamespaceSource() *string
	SetRegionId(v string) *ListGroupsRequest
	GetRegionId() *string
}

type ListGroupsRequest struct {
	// The name of the application group.
	//
	// example:
	//
	// k8s-test
	AppGroupName *string `json:"AppGroupName,omitempty" xml:"AppGroupName,omitempty"`
	// The namespace. You can obtain the ID of the namespace on the **Namespace*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1a72ecb1-b4cc-400a-a71b-20cdec9b****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required only for specific third-party integrations.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) GetAppGroupName() *string {
	return s.AppGroupName
}

func (s *ListGroupsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListGroupsRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *ListGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGroupsRequest) SetAppGroupName(v string) *ListGroupsRequest {
	s.AppGroupName = &v
	return s
}

func (s *ListGroupsRequest) SetNamespace(v string) *ListGroupsRequest {
	s.Namespace = &v
	return s
}

func (s *ListGroupsRequest) SetNamespaceSource(v string) *ListGroupsRequest {
	s.NamespaceSource = &v
	return s
}

func (s *ListGroupsRequest) SetRegionId(v string) *ListGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListGroupsRequest) Validate() error {
	return dara.Validate(s)
}
