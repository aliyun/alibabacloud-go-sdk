// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageTerraformStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *ManageTerraformStateRequest
	GetAction() *string
	SetClientToken(v string) *ManageTerraformStateRequest
	GetClientToken() *string
	SetIdentifier(v string) *ManageTerraformStateRequest
	GetIdentifier() *string
	SetImportResourceId(v string) *ManageTerraformStateRequest
	GetImportResourceId() *string
	SetResourceIdentifier(v string) *ManageTerraformStateRequest
	GetResourceIdentifier() *string
	SetType(v string) *ManageTerraformStateRequest
	GetType() *string
}

type ManageTerraformStateRequest struct {
	// The modification action for the state file. Valid values:
	//
	// - Import
	//
	// - StateRemove
	//
	// This parameter is required.
	//
	// example:
	//
	// Import
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The idempotency token. Format: [0-9a-zA-Z-]{1,64}. A UUID is recommended.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The task identifier. For stack tasks, the format is <$stackId>:<$deploymentName>. For regular tasks, the format is <$TaskId>.
	//
	// This parameter is required.
	//
	// example:
	//
	// stack-as11xxxxxxxxx:developmentA
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// The actual ID of the Terraform resource.
	//
	// example:
	//
	// vsw-xxxxxxxx
	ImportResourceId *string `json:"importResourceId,omitempty" xml:"importResourceId,omitempty"`
	// The identifier of the resource in the Terraform template. For stack tasks, the format is <$componetName>:<$resourceName>. For regular tasks, the format is <$resourceName>.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc:alicloud_vswitch.vswitches[0]
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
	// The task type. Valid values:
	//
	// - Stack
	//
	// - Task
	//
	// This parameter is required.
	//
	// example:
	//
	// Stack
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ManageTerraformStateRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageTerraformStateRequest) GoString() string {
	return s.String()
}

func (s *ManageTerraformStateRequest) GetAction() *string {
	return s.Action
}

func (s *ManageTerraformStateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ManageTerraformStateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *ManageTerraformStateRequest) GetImportResourceId() *string {
	return s.ImportResourceId
}

func (s *ManageTerraformStateRequest) GetResourceIdentifier() *string {
	return s.ResourceIdentifier
}

func (s *ManageTerraformStateRequest) GetType() *string {
	return s.Type
}

func (s *ManageTerraformStateRequest) SetAction(v string) *ManageTerraformStateRequest {
	s.Action = &v
	return s
}

func (s *ManageTerraformStateRequest) SetClientToken(v string) *ManageTerraformStateRequest {
	s.ClientToken = &v
	return s
}

func (s *ManageTerraformStateRequest) SetIdentifier(v string) *ManageTerraformStateRequest {
	s.Identifier = &v
	return s
}

func (s *ManageTerraformStateRequest) SetImportResourceId(v string) *ManageTerraformStateRequest {
	s.ImportResourceId = &v
	return s
}

func (s *ManageTerraformStateRequest) SetResourceIdentifier(v string) *ManageTerraformStateRequest {
	s.ResourceIdentifier = &v
	return s
}

func (s *ManageTerraformStateRequest) SetType(v string) *ManageTerraformStateRequest {
	s.Type = &v
	return s
}

func (s *ManageTerraformStateRequest) Validate() error {
	return dara.Validate(s)
}
