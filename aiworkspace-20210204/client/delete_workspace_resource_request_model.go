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
	// The resource group name. To get the resource group name, see [ListResources](https://help.aliyun.com/document_detail/449143.html).
	//
	// example:
	//
	// group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// A comma-separated list of labels.
	//
	// example:
	//
	// system.supported.eas=true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The deletion behavior. Valid values:
	//
	// - `DetachAndDelete` (default): Detaches the resource from the workspace and deletes the resource.
	//
	// - `Detach`: Detaches the resource from the workspace.
	//
	// example:
	//
	// DetachAndDelete
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// **This parameter is deprecated and will be removed. Use the `ResourceType` parameter instead.**
	//
	// example:
	//
	// DLC
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// A comma-separated list of resource IDs. All specified resources must belong to the same `GroupName`. You must specify a value for at least one of the `GroupName` or `ResourceIds` parameters.
	//
	// example:
	//
	// Resource-dks******jkf,Resource-adf******dss
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The resource type. Valid values:
	//
	// - `ECS`: general-purpose computing resources
	//
	// - `Lingjun`: Lingjun intelligent computing resources
	//
	// - `ACS`: ACS computing resources
	//
	// - `Flink`: Flink resources.
	//
	// - `MaxCompute`: MaxCompute resources. For this resource type, the `Option` parameter can only be set to `Detach`.
	//
	// - `SelfManagedAckPro`: AckPro unified management cluster resources
	//
	// - `SelfManagedAckLingjun`: AckLinjun unified management cluster resources
	//
	// - `SelfManagedASI`: ASI unified management cluster resources (third-party cloud)
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
