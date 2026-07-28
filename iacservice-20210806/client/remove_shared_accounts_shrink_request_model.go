// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSharedAccountsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsShrink(v string) *RemoveSharedAccountsShrinkRequest
	GetAccountIdsShrink() *string
	SetResourceId(v string) *RemoveSharedAccountsShrinkRequest
	GetResourceId() *string
	SetResourceType(v string) *RemoveSharedAccountsShrinkRequest
	GetResourceType() *string
}

type RemoveSharedAccountsShrinkRequest struct {
	// The list of Alibaba Cloud account IDs.
	//
	// This parameter is required.
	AccountIdsShrink *string `json:"accountIds,omitempty" xml:"accountIds,omitempty"`
	// The ID of the resource to unshare.
	//
	//  - If the type is Namespace, set this parameter to the workspace name.
	//
	// - If the type is RegistryModule, set this parameter to \\<namespaceName>/\\<ModuleName>.
	//
	// This parameter is required.
	//
	// example:
	//
	// terraform-alicloud-modules/mongodb
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - RegistryModule: Registry template.
	//
	// - Namespace: workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// RegistryModule
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s RemoveSharedAccountsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSharedAccountsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveSharedAccountsShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *RemoveSharedAccountsShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *RemoveSharedAccountsShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *RemoveSharedAccountsShrinkRequest) SetAccountIdsShrink(v string) *RemoveSharedAccountsShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *RemoveSharedAccountsShrinkRequest) SetResourceId(v string) *RemoveSharedAccountsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *RemoveSharedAccountsShrinkRequest) SetResourceType(v string) *RemoveSharedAccountsShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *RemoveSharedAccountsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
