// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifyMode(v string) *ModifyParameterGroupRequest
	GetModifyMode() *string
	SetOwnerId(v int64) *ModifyParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupDesc(v string) *ModifyParameterGroupRequest
	GetParameterGroupDesc() *string
	SetParameterGroupId(v string) *ModifyParameterGroupRequest
	GetParameterGroupId() *string
	SetParameterGroupName(v string) *ModifyParameterGroupRequest
	GetParameterGroupName() *string
	SetParameters(v string) *ModifyParameterGroupRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyParameterGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyParameterGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyParameterGroupRequest
	GetResourceOwnerId() *int64
}

type ModifyParameterGroupRequest struct {
	// The modification mode of the parameter template. Valid values:
	//
	// 	- **Collectivity**: adds or modifies parameters. This is the default value.
	//
	// >The parameters specified in the **Parameters*	- parameter are added to the current parameter template, or the existing parameters in the current parameter template are modified. Other parameters in the current parameter template remain unaffected.
	//
	// 	- **Individual**: overwrites parameters.
	//
	// >The content in the current parameter template is replaced with the content specified in the **Parameters*	- parameter.
	//
	// example:
	//
	// Collectivity
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the parameter template. The description must be 0 to 200 characters in length.
	//
	// >If you do not specify this parameter, the original description of the parameter template is retained.
	//
	// example:
	//
	// test
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The ID of the parameter template. You can call the DescribeParameterGroups operation to query the ID of the parameter template.
	//
	// This parameter is required.
	//
	// example:
	//
	// rpg-13ppdh****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The name of the parameter template.
	//
	// 	- The name must start with a letter and can contain letters, digits, periods (.), and underscores (_).
	//
	// 	- The name must be 8 to 64 characters in length.
	//
	// >If you do not specify this parameter, the original name of the parameter template is retained.
	//
	// example:
	//
	// testgroup1
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The JSON string of parameters and values. Format: {"Parameter 1":"Value 1","Parameter 2":"Value 2"...}. For more information about the parameters that can be modified, see [Set the parameters of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96063.html) or [Set the parameters of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96751.html).
	//
	// > 	- If the **ModifyMode*	- parameter is set to **Individual**, the specified parameters overwrite the original parameters in the parameter template.
	//
	// > 	- If the **ModifyMode*	- parameter is set to **Collectivity**, the specified parameters are added to or modified in the original parameter template.
	//
	// > 	- If you do not specify this parameter, the original parameters are retained.
	//
	// example:
	//
	// {"back_log":"3000"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the region ID.
	//
	// >The region of a parameter template cannot be modified. You can call the CloneParameterGroup operation to copy the parameter template to another region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can call the DescribeDBInstanceAttribute operation to query the ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterGroupRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyParameterGroupRequest) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *ModifyParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyParameterGroupRequest) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *ModifyParameterGroupRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyParameterGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyParameterGroupRequest) SetModifyMode(v string) *ModifyParameterGroupRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetOwnerId(v int64) *ModifyParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetParameterGroupDesc(v string) *ModifyParameterGroupRequest {
	s.ParameterGroupDesc = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetParameterGroupId(v string) *ModifyParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetParameterGroupName(v string) *ModifyParameterGroupRequest {
	s.ParameterGroupName = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetParameters(v string) *ModifyParameterGroupRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetRegionId(v string) *ModifyParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetResourceGroupId(v string) *ModifyParameterGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetResourceOwnerAccount(v string) *ModifyParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetResourceOwnerId(v int64) *ModifyParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
