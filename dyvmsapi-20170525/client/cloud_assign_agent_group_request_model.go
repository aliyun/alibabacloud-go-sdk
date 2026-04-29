// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAssignAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCnos(v string) *CloudAssignAgentGroupRequest
	GetCnos() *string
	SetEnterpriseId(v int64) *CloudAssignAgentGroupRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudAssignAgentGroupRequest
	GetGno() *string
	SetOwnerId(v int64) *CloudAssignAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudAssignAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudAssignAgentGroupRequest
	GetResourceOwnerId() *int64
}

type CloudAssignAgentGroupRequest struct {
	// 说明：多个座席编号之间以英文逗号( , )分隔，一次最多支持1000个cno；同一座席只能存在于一个外呼组，重复分配会自动移动到新的外呼组中
	//
	// example:
	//
	// 1000,1111
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 外呼组编号；说明：添加此gno对应的外呼组与cnos对应的座席的绑定从属关系; 同一外呼组最多可包含1000个座席
	//
	// example:
	//
	// WH123
	Gno                  *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudAssignAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudAssignAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloudAssignAgentGroupRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudAssignAgentGroupRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudAssignAgentGroupRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudAssignAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudAssignAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudAssignAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudAssignAgentGroupRequest) SetCnos(v string) *CloudAssignAgentGroupRequest {
	s.Cnos = &v
	return s
}

func (s *CloudAssignAgentGroupRequest) SetEnterpriseId(v int64) *CloudAssignAgentGroupRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudAssignAgentGroupRequest) SetGno(v string) *CloudAssignAgentGroupRequest {
	s.Gno = &v
	return s
}

func (s *CloudAssignAgentGroupRequest) SetOwnerId(v int64) *CloudAssignAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudAssignAgentGroupRequest) SetResourceOwnerAccount(v string) *CloudAssignAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudAssignAgentGroupRequest) SetResourceOwnerId(v int64) *CloudAssignAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudAssignAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
