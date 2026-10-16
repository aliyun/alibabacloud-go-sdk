// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeTags(v bool) *GetResourceGroupRequest
	GetIncludeTags() *bool
	SetResourceGroupId(v string) *GetResourceGroupRequest
	GetResourceGroupId() *string
}

type GetResourceGroupRequest struct {
	// Specifies whether to return tag information. Valid values:
	//
	// - false (default): does not return tag information.
	//
	// - true: returns tag information.
	//
	// example:
	//
	// true
	IncludeTags *bool `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The resource group ID.
	//
	// You can call [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) to query resource group IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequest) GetIncludeTags() *bool {
	return s.IncludeTags
}

func (s *GetResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetResourceGroupRequest) SetIncludeTags(v bool) *GetResourceGroupRequest {
	s.IncludeTags = &v
	return s
}

func (s *GetResourceGroupRequest) SetResourceGroupId(v string) *GetResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
