// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DeleteNamespaceRequest
	GetNamespace() *string
	SetRegionId(v string) *DeleteNamespaceRequest
	GetRegionId() *string
}

type DeleteNamespaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteNamespaceRequest) SetRegionId(v string) *DeleteNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
