// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsAIWorkspaceDataEnabled(v bool) *GetResourceGroupShrinkRequest
	GetIsAIWorkspaceDataEnabled() *bool
	SetTagShrink(v string) *GetResourceGroupShrinkRequest
	GetTagShrink() *string
}

type GetResourceGroupShrinkRequest struct {
	// example:
	//
	// true
	IsAIWorkspaceDataEnabled *bool   `json:"IsAIWorkspaceDataEnabled,omitempty" xml:"IsAIWorkspaceDataEnabled,omitempty"`
	TagShrink                *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetResourceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupShrinkRequest) GetIsAIWorkspaceDataEnabled() *bool {
	return s.IsAIWorkspaceDataEnabled
}

func (s *GetResourceGroupShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *GetResourceGroupShrinkRequest) SetIsAIWorkspaceDataEnabled(v bool) *GetResourceGroupShrinkRequest {
	s.IsAIWorkspaceDataEnabled = &v
	return s
}

func (s *GetResourceGroupShrinkRequest) SetTagShrink(v string) *GetResourceGroupShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *GetResourceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
