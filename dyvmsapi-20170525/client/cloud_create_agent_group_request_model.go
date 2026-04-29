// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAgentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CloudCreateAgentGroupRequest
	GetComment() *string
	SetEnterpriseId(v int64) *CloudCreateAgentGroupRequest
	GetEnterpriseId() *int64
	SetGno(v string) *CloudCreateAgentGroupRequest
	GetGno() *string
	SetGroupName(v string) *CloudCreateAgentGroupRequest
	GetGroupName() *string
	SetOwnerId(v int64) *CloudCreateAgentGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudCreateAgentGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudCreateAgentGroupRequest
	GetResourceOwnerId() *int64
}

type CloudCreateAgentGroupRequest struct {
	// 说明：描述长度不能大于100,需进行UTF-8格式的URLEncode编码
	//
	// example:
	//
	// comment3
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 说明：同一企业外呼组编号编号不能重复，外呼组编号长度为2-20，以字母开头,须同时且只能包含字母和数字 同一企业最多可创建1000个外呼组
	//
	// This parameter is required.
	//
	// example:
	//
	// WH12
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 说明：外呼组名称不能为空，长度不能大于50,需进行UTF-8格式的URLEncode编码
	//
	// This parameter is required.
	//
	// example:
	//
	// gpname1
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudCreateAgentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentGroupRequest) GetComment() *string {
	return s.Comment
}

func (s *CloudCreateAgentGroupRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateAgentGroupRequest) GetGno() *string {
	return s.Gno
}

func (s *CloudCreateAgentGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CloudCreateAgentGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudCreateAgentGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudCreateAgentGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudCreateAgentGroupRequest) SetComment(v string) *CloudCreateAgentGroupRequest {
	s.Comment = &v
	return s
}

func (s *CloudCreateAgentGroupRequest) SetEnterpriseId(v int64) *CloudCreateAgentGroupRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateAgentGroupRequest) SetGno(v string) *CloudCreateAgentGroupRequest {
	s.Gno = &v
	return s
}

func (s *CloudCreateAgentGroupRequest) SetGroupName(v string) *CloudCreateAgentGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CloudCreateAgentGroupRequest) SetOwnerId(v int64) *CloudCreateAgentGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudCreateAgentGroupRequest) SetResourceOwnerAccount(v string) *CloudCreateAgentGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudCreateAgentGroupRequest) SetResourceOwnerId(v int64) *CloudCreateAgentGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudCreateAgentGroupRequest) Validate() error {
	return dara.Validate(s)
}
