// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteWorkspaceResourceRequest
	GetGroupName() *string
	SetLabels(v string) *DeleteWorkspaceResourceRequest
	GetLabels() *string
	SetOption(v string) *DeleteWorkspaceResourceRequest
	GetOption() *string
	SetProductType(v string) *DeleteWorkspaceResourceRequest
	GetProductType() *string
	SetResourceIds(v string) *DeleteWorkspaceResourceRequest
	GetResourceIds() *string
	SetResourceType(v string) *DeleteWorkspaceResourceRequest
	GetResourceType() *string
}

type DeleteWorkspaceResourceRequest struct {
	// The name of the resource group. You can call [ListResources](https://help.aliyun.com/document_detail/449143.html) to obtain the name of the resource group.
	//
	// example:
	//
	// group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tags. Multiple tags are separated by commas (,).
	//
	// example:
	//
	// system.supported.eas=true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The operation to perform. Valid values:
	//
	// 	- DetachAndDelete: disassociates a resource from a workspace and deletes the resource in the workspace. This is the default value.
	//
	// 	- Detach: disassociates a resource group from a workspace.
	//
	// example:
	//
	// DetachAndDelete
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// **This field is no longer used and will be removed. Use the ResourceType field instead.
	//
	// example:
	//
	// DLC
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The resource IDs. Multiple resource IDs are separated by commas (,). The GroupName values for the specified resources must be the same. You cannot leave both GroupName and ResourceIds empty. You can specify both parameters.
	//
	// example:
	//
	// Resource-dks******jkf,Resource-adf******dss
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The resource type. Valid values:
	//
	// 	- ECS
	//
	// 	- Lingjun
	//
	// 	- ACS
	//
	// 	- FLINK
	//
	// 	- MaxCompute (This resource type is valid only if Option is set to Detach.)
	//
	// example:
	//
	// DLC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteWorkspaceResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResourceRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteWorkspaceResourceRequest) GetLabels() *string {
	return s.Labels
}

func (s *DeleteWorkspaceResourceRequest) GetOption() *string {
	return s.Option
}

func (s *DeleteWorkspaceResourceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DeleteWorkspaceResourceRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *DeleteWorkspaceResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteWorkspaceResourceRequest) SetGroupName(v string) *DeleteWorkspaceResourceRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetLabels(v string) *DeleteWorkspaceResourceRequest {
	s.Labels = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetOption(v string) *DeleteWorkspaceResourceRequest {
	s.Option = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetProductType(v string) *DeleteWorkspaceResourceRequest {
	s.ProductType = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetResourceIds(v string) *DeleteWorkspaceResourceRequest {
	s.ResourceIds = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetResourceType(v string) *DeleteWorkspaceResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) Validate() error {
	return dara.Validate(s)
}
