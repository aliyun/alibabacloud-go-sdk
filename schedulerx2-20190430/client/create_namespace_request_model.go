// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateNamespaceRequest
	GetDescription() *string
	SetName(v string) *CreateNamespaceRequest
	GetName() *string
	SetRegionId(v string) *CreateNamespaceRequest
	GetRegionId() *string
	SetUid(v string) *CreateNamespaceRequest
	GetUid() *string
}

type CreateNamespaceRequest struct {
	// The description of the namespace.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-env
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The unique identifier (UID) of the namespace. We recommend that you use the universally unique identifier (UUID) to generate the UID.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNamespaceRequest) GetUid() *string {
	return s.Uid
}

func (s *CreateNamespaceRequest) SetDescription(v string) *CreateNamespaceRequest {
	s.Description = &v
	return s
}

func (s *CreateNamespaceRequest) SetName(v string) *CreateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *CreateNamespaceRequest) SetRegionId(v string) *CreateNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNamespaceRequest) SetUid(v string) *CreateNamespaceRequest {
	s.Uid = &v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
