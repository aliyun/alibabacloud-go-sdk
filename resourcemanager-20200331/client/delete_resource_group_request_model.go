// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *DeleteResourceGroupRequest
	GetResourceGroupId() *string
}

type DeleteResourceGroupRequest struct {
	// The ID of the resource group.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteResourceGroupRequest) SetResourceGroupId(v string) *DeleteResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
