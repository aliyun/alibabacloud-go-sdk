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
	// The name of the resource group. To obtain the name of the resource group, see [ListResources](https://help.aliyun.com/document_detail/449143.html).
	//
	// example:
	//
	// group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The list of tags, separated by commas.
	//
	// example:
	//
	// system.supported.eas=true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The deletion behavior. Valid values:
	//
	// - DetachAndDelete: Unbinds the resources from the workspace and deletes the resources. This is the default value.
	//
	// - Detach: Unbinds the resources from the workspace.
	//
	// example:
	//
	// DetachAndDelete
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// **[Deprecated]*	- This field is deprecated and will be retired. Use the ResourceType field.
	//
	// example:
	//
	// DLC
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The IDs of the resources, separated by commas. The GroupName of these resource IDs must be the same. GroupName and ResourceIds cannot be empty at the same time, but both can be specified.
	//
	// example:
	//
	// Resource-dks******jkf,Resource-adf******dss
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The list of resource types. Valid values:
	//
	// - ECS: general computing resources
	//
	// - Lingjun: Lingjun AI computing resources
	//
	// - ACS: ACS computing resources
	//
	// - FLINK: Flink resources
	//
	// - MaxCompute: MaxCompute resources. Only the Detach option is supported. The DetachAndDelete option is not supported.
	//
	// - SelfManagedAckPro: ACK Pro unified management cluster resources
	//
	// - SelfManagedAckLingjun: ACK Lingjun unified management cluster resources
	//
	// - SelfManagedASI: ASI unified management cluster resources (third-party cloud)
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
