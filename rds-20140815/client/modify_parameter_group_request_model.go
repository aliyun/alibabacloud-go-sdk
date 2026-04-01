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
	// 	- **Collectivity*	- (default): adds new parameters or modifies parameters in the original parameter template.
	//
	// >  If you set the ModifyMode parameter to Collectivity, the system adds the value of the **Parameters*	- parameter to the original parameter template or modifies the corresponding parameters in the original parameter template. Other parameters in the original parameter template are not affected.
	//
	// 	- **Individual**: overwrites original parameters.
	//
	// >  If you set the ModifyMode parameter to Individual, the system uses the value of the **Parameters*	- parameter to overwrite the parameter settings in the original parameter template.
	//
	// example:
	//
	// Collectivity
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new description of the parameter template. The description can be up to 200 characters in length.
	//
	// > If you do not specify this parameter, the original description of the parameter template is retained.
	//
	// example:
	//
	// test
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The parameter template ID. You can call the DescribeParameterGroups operation to query the parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rpg-13ppdh****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The parameter template name.
	//
	// 	- The name can contain letters, digits, periods (.), and underscores (_). It must start with a letter.
	//
	// 	- It can be 8 to 64 characters in length.
	//
	// > If you do not specify this parameter, the original name of the parameter template is retained.
	//
	// example:
	//
	// testgroup1
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// A JSON string that consists of parameters and their values in the parameter template. Format: {"Parameter 1":"Value of Parameter 1","Parameter 2":"Value of Parameter 2"...}. For more information about the parameters that can be modified, see [Modify the parameters of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96063.html) or [Modify the parameters of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96751.html).
	//
	// > 	- If **ModifyMode*	- is set to **Individual*	- and this parameter is specified, the new parameters overwrite the parameters in the original parameter template.
	//
	// > 	- If you set **ModifyMode*	- to **Collectivity*	- and specify this parameter, the new parameters are added to the original parameter template, or the parameters in the original parameter template are modified.
	//
	// > 	- If you do not specify this parameter, the parameters in the original parameter template remain unchanged.
	//
	// example:
	//
	// {"back_log":"3000"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// >  The region of a parameter template cannot be changed. You can call the CloneParameterGroup operation to replicate a parameter template to a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
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
