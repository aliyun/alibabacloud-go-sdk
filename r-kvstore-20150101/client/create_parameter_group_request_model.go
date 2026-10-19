// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateParameterGroupRequest
	GetCategory() *string
	SetEngineType(v string) *CreateParameterGroupRequest
	GetEngineType() *string
	SetEngineVersion(v string) *CreateParameterGroupRequest
	GetEngineVersion() *string
	SetOwnerAccount(v string) *CreateParameterGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupDesc(v string) *CreateParameterGroupRequest
	GetParameterGroupDesc() *string
	SetParameterGroupName(v string) *CreateParameterGroupRequest
	GetParameterGroupName() *string
	SetParameters(v string) *CreateParameterGroupRequest
	GetParameters() *string
	SetRegionId(v string) *CreateParameterGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateParameterGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateParameterGroupRequest
	GetSecurityToken() *string
}

type CreateParameterGroupRequest struct {
	// The product category. Valid values:
	//
	// 	- **standard**: Redis Community Edition.
	//
	// 	- **enterprise**: Tair (Enterprise Edition).
	//
	// This parameter is required.
	//
	// example:
	//
	// enterprise
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **redis**: Redis Community Edition or Tair memory-optimized.
	//
	// 	- **tair_pena**: Tair persistent memory-optimized.
	//
	// 	- **tair_pdb**: Tair disk-optimized (ESSD/SSD).
	//
	// This parameter is required.
	//
	// example:
	//
	// redis
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The compatible version. Valid values:
	//
	// 	- Redis Community Edition supports **5.0**, **6.0**, and **7.0**.
	//
	// 	- Tair memory-optimized is compatible with Redis 5.0, 6.0, and 7.0 protocols. Set this parameter to **5.0**, **6.0**, or **7.0**.
	//
	// 	- Tair persistent memory-optimized is compatible with the Redis 6.0 protocol. Set this parameter to **1.0**.
	//
	// 	- Tair disk-optimized (ESSD/SSD) is compatible with the Redis 6.0 protocol. To create a parameter template for the ESSD type, set this parameter to **1.0**. To create a parameter template for the SSD type, set this parameter to **2.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the parameter template. The description can be 0 to 200 characters in length.
	//
	// example:
	//
	// test
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The name of the parameter template. The name must meet the following requirements:
	//
	// - The name can contain letters, digits, and underscores (_). It must start with a letter and cannot contain Chinese characters.
	//
	// - The name must be 8 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// tw_test1
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// A JSON string that consists of parameters and their values, in the following format: {"Parameter 1":"Value of Parameter 1","Parameter 2":"Value of Parameter 2"...}.
	//
	// > The parameters that can be added vary by version. Refer to the console for details.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"hz":"15","#no_loose_disabled-commands":"flushall"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupRequest) GetCategory() *string {
	return s.Category
}

func (s *CreateParameterGroupRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateParameterGroupRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateParameterGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateParameterGroupRequest) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *CreateParameterGroupRequest) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *CreateParameterGroupRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateParameterGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateParameterGroupRequest) SetCategory(v string) *CreateParameterGroupRequest {
	s.Category = &v
	return s
}

func (s *CreateParameterGroupRequest) SetEngineType(v string) *CreateParameterGroupRequest {
	s.EngineType = &v
	return s
}

func (s *CreateParameterGroupRequest) SetEngineVersion(v string) *CreateParameterGroupRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateParameterGroupRequest) SetOwnerAccount(v string) *CreateParameterGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateParameterGroupRequest) SetOwnerId(v int64) *CreateParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameterGroupDesc(v string) *CreateParameterGroupRequest {
	s.ParameterGroupDesc = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameterGroupName(v string) *CreateParameterGroupRequest {
	s.ParameterGroupName = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameters(v string) *CreateParameterGroupRequest {
	s.Parameters = &v
	return s
}

func (s *CreateParameterGroupRequest) SetRegionId(v string) *CreateParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetResourceOwnerAccount(v string) *CreateParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateParameterGroupRequest) SetResourceOwnerId(v int64) *CreateParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetSecurityToken(v string) *CreateParameterGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
