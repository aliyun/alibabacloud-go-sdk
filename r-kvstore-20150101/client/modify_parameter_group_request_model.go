// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ModifyParameterGroupRequest
	GetCategory() *string
	SetOwnerAccount(v string) *ModifyParameterGroupRequest
	GetOwnerAccount() *string
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
	SetResourceOwnerAccount(v string) *ModifyParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyParameterGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyParameterGroupRequest
	GetSecurityToken() *string
}

type ModifyParameterGroupRequest struct {
	// The service category. Valid values:
	//
	// 	- **standard**: Redis Open-Source Edition
	//
	// 	- **enterprise**: Tair (Enterprise Edition)
	//
	// This parameter is required.
	//
	// example:
	//
	// enterprise
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the parameter template. The description must be 0 to 200 characters in length.
	//
	// example:
	//
	// test
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sys-001****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The new name of the parameter template. The name must meet the following requirements:
	//
	// 	- The name can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
	//
	// 	- The name can be 8 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testGroupName
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// A JSON-formatted object that specifies the parameter-value pairs. Format: {"Parameter 1":"Value 1","Parameter 2":"Value 2"...}. The specified value overwrites the original content.
	//
	// >  The parameters that can be added for different editions are displayed in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"hz":"12"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterGroupRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyParameterGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
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

func (s *ModifyParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyParameterGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyParameterGroupRequest) SetCategory(v string) *ModifyParameterGroupRequest {
	s.Category = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetOwnerAccount(v string) *ModifyParameterGroupRequest {
	s.OwnerAccount = &v
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

func (s *ModifyParameterGroupRequest) SetResourceOwnerAccount(v string) *ModifyParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetResourceOwnerId(v int64) *ModifyParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetSecurityToken(v string) *ModifyParameterGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
