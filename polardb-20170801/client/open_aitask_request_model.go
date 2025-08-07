// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAITaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *OpenAITaskRequest
	GetDBClusterId() *string
	SetNodeType(v string) *OpenAITaskRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *OpenAITaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenAITaskRequest
	GetOwnerId() *int64
	SetPassword(v string) *OpenAITaskRequest
	GetPassword() *string
	SetRegionId(v string) *OpenAITaskRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *OpenAITaskRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *OpenAITaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenAITaskRequest
	GetResourceOwnerId() *int64
	SetUsername(v string) *OpenAITaskRequest
	GetUsername() *string
}

type OpenAITaskRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node type. Valid values:
	//
	// 	- **DLNode**: This node is an AI node.
	//
	// 	- **SearchNode**: This node is a node for which the PolarDB for AI feature is enabled.
	//
	// example:
	//
	// DLNode
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password used to access the database for which you want to enable the PolarDB for AI feature.
	//
	// example:
	//
	// testPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the region.
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
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The username used to access the database for which you want to enable the PolarDB for AI feature.
	//
	// example:
	//
	// testAccountName
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s OpenAITaskRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenAITaskRequest) GoString() string {
	return s.String()
}

func (s *OpenAITaskRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *OpenAITaskRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *OpenAITaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenAITaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenAITaskRequest) GetPassword() *string {
	return s.Password
}

func (s *OpenAITaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenAITaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *OpenAITaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenAITaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenAITaskRequest) GetUsername() *string {
	return s.Username
}

func (s *OpenAITaskRequest) SetDBClusterId(v string) *OpenAITaskRequest {
	s.DBClusterId = &v
	return s
}

func (s *OpenAITaskRequest) SetNodeType(v string) *OpenAITaskRequest {
	s.NodeType = &v
	return s
}

func (s *OpenAITaskRequest) SetOwnerAccount(v string) *OpenAITaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenAITaskRequest) SetOwnerId(v int64) *OpenAITaskRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenAITaskRequest) SetPassword(v string) *OpenAITaskRequest {
	s.Password = &v
	return s
}

func (s *OpenAITaskRequest) SetRegionId(v string) *OpenAITaskRequest {
	s.RegionId = &v
	return s
}

func (s *OpenAITaskRequest) SetResourceGroupId(v string) *OpenAITaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *OpenAITaskRequest) SetResourceOwnerAccount(v string) *OpenAITaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenAITaskRequest) SetResourceOwnerId(v int64) *OpenAITaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenAITaskRequest) SetUsername(v string) *OpenAITaskRequest {
	s.Username = &v
	return s
}

func (s *OpenAITaskRequest) Validate() error {
	return dara.Validate(s)
}
