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
	// The service category. Valid values:
	//
	// 	- **standard**: Community Edition
	//
	// 	- **enterprise**: Enhanced Edition (Tair)
	//
	// This parameter is required.
	//
	// example:
	//
	// enterprise
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **redis**: ApsaraDB for Redis Community Edition instance or Tair DRAM-based instance
	//
	// 	- **tair_pena**: Tair persistent memory-optimized instance
	//
	// 	- **tair_pdb**: Tair ESSD/SSD-based instance
	//
	// This parameter is required.
	//
	// example:
	//
	// redis
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The compatible engine version. Valid values:
	//
	// 	- For ApsaraDB for Redis Community Edition instances, set the parameter to **5.0**, **6.0**, or **7.0**.
	//
	// 	- For Tair DRAM-based instances that are compatible with Redis 5.0, 6.0, or 7.0, set the parameter to **5.0**, **6.0**, or **7.0**.
	//
	// 	- For Tair persistent memory-optimized instances that are compatible with Redis 6.0, set the parameter to **1.0**.
	//
	// 	- For Tair ESSD-based instances that are compatible with Redis 6.0, set the parameter to **1.0**. For Tair SSD-based instances that are compatible with Redis 6.0, set the parameter to **2.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the parameter template. The description must be 0 to 200 characters in length.
	//
	// example:
	//
	// test
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
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
	// tw_test1
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// A JSON-formatted object that specifies the parameter-value pairs. Format: {"Parameter 1":"Value 1","Parameter 2":"Value 2"...}. The specified value overwrites the original content.
	//
	// >  The parameters that can be added for different editions are displayed in the console.
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
