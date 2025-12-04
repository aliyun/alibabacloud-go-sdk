// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *ChangeResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceRegionId(v string) *ChangeResourceGroupRequest
	GetResourceRegionId() *string
}

type ChangeResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1i43033051****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceRegionId(v string) *ChangeResourceGroupRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
