// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBType(v string) *CreateParameterGroupRequest
	GetDBType() *string
	SetDBVersion(v string) *CreateParameterGroupRequest
	GetDBVersion() *string
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
	SetResourceGroupId(v string) *CreateParameterGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateParameterGroupRequest
	GetResourceOwnerId() *int64
}

type CreateParameterGroupRequest struct {
	// The type of the database engine. Only **MySQL*	- is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine. Valid values:
	//
	// 	- **5.6**
	//
	// 	- **5.7**
	//
	// 	- **8.0**
	//
	// This parameter is required.
	//
	// example:
	//
	// 8.0
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the parameter template. It must be 0 to 199 characters in length.
	//
	// example:
	//
	// test_group
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The name of the parameter template. The name must meet the following requirements:
	//
	// 	- It can contain letters, digits, and underscores (_). It must start with a letter and cannot end with an underscore.**
	//
	// 	- It must be 8 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The JSON string that consists of parameters and values. The parameter values are strings. Example: `{"wait_timeout":"86400","innodb_old_blocks_time":"1000"}`.
	//
	// > You can call the [DescribeParameterTemplates](https://help.aliyun.com/document_detail/207428.html) operation to query the details of all parameters in the cluster of a specified engine version, such as the parameter name and valid values.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"wait_timeout":"86400","innodb_old_blocks_time":"1000"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-**********
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

func (s *CreateParameterGroupRequest) GetDBType() *string {
	return s.DBType
}

func (s *CreateParameterGroupRequest) GetDBVersion() *string {
	return s.DBVersion
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

func (s *CreateParameterGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateParameterGroupRequest) SetDBType(v string) *CreateParameterGroupRequest {
	s.DBType = &v
	return s
}

func (s *CreateParameterGroupRequest) SetDBVersion(v string) *CreateParameterGroupRequest {
	s.DBVersion = &v
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
