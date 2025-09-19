// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputingResourceProvider(v string) *CreateResourceGroupRequest
	GetComputingResourceProvider() *string
	SetDescription(v string) *CreateResourceGroupRequest
	GetDescription() *string
	SetName(v string) *CreateResourceGroupRequest
	GetName() *string
	SetResourceType(v string) *CreateResourceGroupRequest
	GetResourceType() *string
	SetTag(v []*CreateResourceGroupRequestTag) *CreateResourceGroupRequest
	GetTag() []*CreateResourceGroupRequestTag
	SetUserVpc(v *UserVpc) *CreateResourceGroupRequest
	GetUserVpc() *UserVpc
	SetVersion(v string) *CreateResourceGroupRequest
	GetVersion() *string
}

type CreateResourceGroupRequest struct {
	// example:
	//
	// Ecs
	ComputingResourceProvider *string `json:"ComputingResourceProvider,omitempty" xml:"ComputingResourceProvider,omitempty"`
	// example:
	//
	// test_api_report
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// testResourceGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Ecs
	ResourceType *string                          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*CreateResourceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UserVpc      *UserVpc                         `json:"UserVpc,omitempty" xml:"UserVpc,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) GetComputingResourceProvider() *string {
	return s.ComputingResourceProvider
}

func (s *CreateResourceGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateResourceGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateResourceGroupRequest) GetTag() []*CreateResourceGroupRequestTag {
	return s.Tag
}

func (s *CreateResourceGroupRequest) GetUserVpc() *UserVpc {
	return s.UserVpc
}

func (s *CreateResourceGroupRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateResourceGroupRequest) SetComputingResourceProvider(v string) *CreateResourceGroupRequest {
	s.ComputingResourceProvider = &v
	return s
}

func (s *CreateResourceGroupRequest) SetDescription(v string) *CreateResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceGroupRequest) SetName(v string) *CreateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupRequest) SetResourceType(v string) *CreateResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateResourceGroupRequest) SetTag(v []*CreateResourceGroupRequestTag) *CreateResourceGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateResourceGroupRequest) SetUserVpc(v *UserVpc) *CreateResourceGroupRequest {
	s.UserVpc = v
	return s
}

func (s *CreateResourceGroupRequest) SetVersion(v string) *CreateResourceGroupRequest {
	s.Version = &v
	return s
}

func (s *CreateResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateResourceGroupRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateResourceGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateResourceGroupRequestTag) SetKey(v string) *CreateResourceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateResourceGroupRequestTag) SetValue(v string) *CreateResourceGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateResourceGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
