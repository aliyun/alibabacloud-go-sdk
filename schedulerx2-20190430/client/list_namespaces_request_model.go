// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *ListNamespacesRequest
	GetNamespace() *string
	SetNamespaceName(v string) *ListNamespacesRequest
	GetNamespaceName() *string
	SetRegionId(v string) *ListNamespacesRequest
	GetRegionId() *string
}

type ListNamespacesRequest struct {
	// The namespace ID.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// schedulerx-dev
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListNamespacesRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListNamespacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNamespacesRequest) SetNamespace(v string) *ListNamespacesRequest {
	s.Namespace = &v
	return s
}

func (s *ListNamespacesRequest) SetNamespaceName(v string) *ListNamespacesRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListNamespacesRequest) SetRegionId(v string) *ListNamespacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
