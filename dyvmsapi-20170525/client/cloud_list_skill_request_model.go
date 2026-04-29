// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CloudListSkillRequest
	GetComment() *string
	SetEnterpriseId(v int64) *CloudListSkillRequest
	GetEnterpriseId() *int64
	SetLimit(v string) *CloudListSkillRequest
	GetLimit() *string
	SetName(v string) *CloudListSkillRequest
	GetName() *string
	SetOffset(v string) *CloudListSkillRequest
	GetOffset() *string
	SetOrder(v int64) *CloudListSkillRequest
	GetOrder() *int64
	SetOwnerId(v int64) *CloudListSkillRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudListSkillRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudListSkillRequest
	GetResourceOwnerId() *int64
}

type CloudListSkillRequest struct {
	// 描述
	//
	// example:
	//
	// comment1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 条数；正整数，默认值是10，最大值是500
	//
	// example:
	//
	// 10
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 技能名称
	//
	// example:
	//
	// name1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 从第几条开始；正整数，默认值是0
	//
	// example:
	//
	// 0
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// 排序方式,按照创建时间排序；0正序 1倒序
	//
	// example:
	//
	// 0
	Order                *int64  `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudListSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudListSkillRequest) GoString() string {
	return s.String()
}

func (s *CloudListSkillRequest) GetComment() *string {
	return s.Comment
}

func (s *CloudListSkillRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListSkillRequest) GetLimit() *string {
	return s.Limit
}

func (s *CloudListSkillRequest) GetName() *string {
	return s.Name
}

func (s *CloudListSkillRequest) GetOffset() *string {
	return s.Offset
}

func (s *CloudListSkillRequest) GetOrder() *int64 {
	return s.Order
}

func (s *CloudListSkillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudListSkillRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudListSkillRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudListSkillRequest) SetComment(v string) *CloudListSkillRequest {
	s.Comment = &v
	return s
}

func (s *CloudListSkillRequest) SetEnterpriseId(v int64) *CloudListSkillRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListSkillRequest) SetLimit(v string) *CloudListSkillRequest {
	s.Limit = &v
	return s
}

func (s *CloudListSkillRequest) SetName(v string) *CloudListSkillRequest {
	s.Name = &v
	return s
}

func (s *CloudListSkillRequest) SetOffset(v string) *CloudListSkillRequest {
	s.Offset = &v
	return s
}

func (s *CloudListSkillRequest) SetOrder(v int64) *CloudListSkillRequest {
	s.Order = &v
	return s
}

func (s *CloudListSkillRequest) SetOwnerId(v int64) *CloudListSkillRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudListSkillRequest) SetResourceOwnerAccount(v string) *CloudListSkillRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudListSkillRequest) SetResourceOwnerId(v int64) *CloudListSkillRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudListSkillRequest) Validate() error {
	return dara.Validate(s)
}
