// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupInput interface {
	dara.Model
	String() string
	GoString() string
	SetNewResourceGroupId(v string) *ChangeResourceGroupInput
	GetNewResourceGroupId() *string
	SetResourceId(v string) *ChangeResourceGroupInput
	GetResourceId() *string
	SetResourceType(v string) *ChangeResourceGroupInput
	GetResourceType() *string
}

type ChangeResourceGroupInput struct {
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupInput) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupInput) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupInput) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupInput) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupInput) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupInput) SetNewResourceGroupId(v string) *ChangeResourceGroupInput {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupInput) SetResourceId(v string) *ChangeResourceGroupInput {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupInput) SetResourceType(v string) *ChangeResourceGroupInput {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceGroupInput) Validate() error {
	return dara.Validate(s)
}
