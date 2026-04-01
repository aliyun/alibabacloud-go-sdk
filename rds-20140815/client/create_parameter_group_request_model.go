// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *CreateParameterGroupRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateParameterGroupRequest
	GetEngineVersion() *string
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
	SetResourceGroupId(v string) *CreateParameterGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateParameterGroupRequest
	GetResourceOwnerId() *int64
}

type CreateParameterGroupRequest struct {
	// The database engine. Valid values:
	//
	// 	- **mysql**
	//
	// 	- **PostgreSQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance.
	//
	// 	- If the instance runs MySQL, the instance must run one of the following MySQL versions:
	//
	//     	- **5.6**
	//
	//     	- **5.7**
	//
	//     	- **8.0**
	//
	// 	- If the instance runs PostgreSQL, the instance must run one of the following PostgreSQL versions:
	//
	//     	- **10.0**
	//
	//     	- **11.0**
	//
	//     	- **12.0**
	//
	//     	- **13.0**
	//
	//     	- **14.0**
	//
	//     	- **15.0**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the parameter template. The value can be up to 200 characters in length.
	//
	// example:
	//
	// test
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The name of the parameter template.
	//
	// 	- The value must start with a letter and can contain letters, digits, periods (.), and underscores (_).
	//
	// 	- The value can be 8 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1234
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// A JSON string that consists of parameters and their values in the parameter template. Format: {"Parameter 1":"Value of Parameter 1","Parameter 2":"Value of Parameter 2"...}. For more information about the parameters that can be modified, see [Modify the parameters of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96063.html) or [Modify the parameters of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96751.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"back_log":"3000","wait_timeout":"86400"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the parameter template. You can call the DescribeRegions operation to query the most recent zone list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to obtain the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateParameterGroupRequest) GetEngineVersion() *string {
	return s.EngineVersion
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

func (s *CreateParameterGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateParameterGroupRequest) SetEngine(v string) *CreateParameterGroupRequest {
	s.Engine = &v
	return s
}

func (s *CreateParameterGroupRequest) SetEngineVersion(v string) *CreateParameterGroupRequest {
	s.EngineVersion = &v
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

func (s *CreateParameterGroupRequest) SetResourceGroupId(v string) *CreateParameterGroupRequest {
	s.ResourceGroupId = &v
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

func (s *CreateParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
