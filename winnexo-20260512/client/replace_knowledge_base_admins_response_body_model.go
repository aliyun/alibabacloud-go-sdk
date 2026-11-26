// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceKnowledgeBaseAdminsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReplaceKnowledgeBaseAdminsResponseBody
	GetCode() *string
	SetCreator(v *ReplaceKnowledgeBaseAdminsResponseBodyCreator) *ReplaceKnowledgeBaseAdminsResponseBody
	GetCreator() *ReplaceKnowledgeBaseAdminsResponseBodyCreator
	SetDirectoryId(v string) *ReplaceKnowledgeBaseAdminsResponseBody
	GetDirectoryId() *string
	SetEffectiveApprovers(v []*ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) *ReplaceKnowledgeBaseAdminsResponseBody
	GetEffectiveApprovers() []*ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers
	SetMessage(v string) *ReplaceKnowledgeBaseAdminsResponseBody
	GetMessage() *string
	SetPrincipals(v []*ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) *ReplaceKnowledgeBaseAdminsResponseBody
	GetPrincipals() []*ReplaceKnowledgeBaseAdminsResponseBodyPrincipals
	SetRequestId(v string) *ReplaceKnowledgeBaseAdminsResponseBody
	GetRequestId() *string
}

type ReplaceKnowledgeBaseAdminsResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 知识库创建者
	Creator *ReplaceKnowledgeBaseAdminsResponseBodyCreator `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	// 已配置的知识库根目录业务 ID
	//
	// example:
	//
	// kb_001
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 用户组展开和去重后的有效审批人
	EffectiveApprovers []*ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers `json:"effectiveApprovers,omitempty" xml:"effectiveApprovers,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 配置完成后的管理员主体列表，包含创建者
	Principals []*ReplaceKnowledgeBaseAdminsResponseBodyPrincipals `json:"principals,omitempty" xml:"principals,omitempty" type:"Repeated"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ReplaceKnowledgeBaseAdminsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseAdminsResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) GetCreator() *ReplaceKnowledgeBaseAdminsResponseBodyCreator {
	return s.Creator
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) GetEffectiveApprovers() []*ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers {
	return s.EffectiveApprovers
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) GetPrincipals() []*ReplaceKnowledgeBaseAdminsResponseBodyPrincipals {
	return s.Principals
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) SetCode(v string) *ReplaceKnowledgeBaseAdminsResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) SetCreator(v *ReplaceKnowledgeBaseAdminsResponseBodyCreator) *ReplaceKnowledgeBaseAdminsResponseBody {
	s.Creator = v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) SetDirectoryId(v string) *ReplaceKnowledgeBaseAdminsResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) SetEffectiveApprovers(v []*ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) *ReplaceKnowledgeBaseAdminsResponseBody {
	s.EffectiveApprovers = v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) SetMessage(v string) *ReplaceKnowledgeBaseAdminsResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) SetPrincipals(v []*ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) *ReplaceKnowledgeBaseAdminsResponseBody {
	s.Principals = v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) SetRequestId(v string) *ReplaceKnowledgeBaseAdminsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBody) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.EffectiveApprovers != nil {
		for _, item := range s.EffectiveApprovers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Principals != nil {
		for _, item := range s.Principals {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceKnowledgeBaseAdminsResponseBodyCreator struct {
	// 是否为不可移除的创建者
	//
	// example:
	//
	// true
	Immutable *bool `json:"immutable,omitempty" xml:"immutable,omitempty"`
	// 用户 ID 或用户组 ID
	//
	// example:
	//
	// 10001
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// 主体显示名称，无法解析时为空
	//
	// example:
	//
	// 张三
	PrincipalName *string `json:"principalName,omitempty" xml:"principalName,omitempty"`
	// 主体类型：USER 或 USER_GROUP
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"principalType,omitempty" xml:"principalType,omitempty"`
	// 权限来源：CREATOR 或 DIRECT
	//
	// example:
	//
	// CREATOR
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ReplaceKnowledgeBaseAdminsResponseBodyCreator) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseAdminsResponseBodyCreator) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) GetImmutable() *bool {
	return s.Immutable
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) GetSource() *string {
	return s.Source
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) SetImmutable(v bool) *ReplaceKnowledgeBaseAdminsResponseBodyCreator {
	s.Immutable = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) SetPrincipalId(v string) *ReplaceKnowledgeBaseAdminsResponseBodyCreator {
	s.PrincipalId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) SetPrincipalName(v string) *ReplaceKnowledgeBaseAdminsResponseBodyCreator {
	s.PrincipalName = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) SetPrincipalType(v string) *ReplaceKnowledgeBaseAdminsResponseBodyCreator {
	s.PrincipalType = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) SetSource(v string) *ReplaceKnowledgeBaseAdminsResponseBodyCreator {
	s.Source = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyCreator) Validate() error {
	return dara.Validate(s)
}

type ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers struct {
	// 产生该审批权限的主体 ID 列表
	//
	// example:
	//
	// ["10001","group-1"]
	SourcePrincipalIds []*string `json:"sourcePrincipalIds,omitempty" xml:"sourcePrincipalIds,omitempty" type:"Repeated"`
	// 有效审批用户 ID
	//
	// example:
	//
	// 10001
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// 租户内显示名称，无法解析时为空
	//
	// example:
	//
	// 张三
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) GetSourcePrincipalIds() []*string {
	return s.SourcePrincipalIds
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) GetUserId() *int64 {
	return s.UserId
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) GetUserName() *string {
	return s.UserName
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) SetSourcePrincipalIds(v []*string) *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers {
	s.SourcePrincipalIds = v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) SetUserId(v int64) *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers {
	s.UserId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) SetUserName(v string) *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers {
	s.UserName = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyEffectiveApprovers) Validate() error {
	return dara.Validate(s)
}

type ReplaceKnowledgeBaseAdminsResponseBodyPrincipals struct {
	// 是否为不可移除的创建者
	//
	// example:
	//
	// true
	Immutable *bool `json:"immutable,omitempty" xml:"immutable,omitempty"`
	// 用户 ID 或用户组 ID
	//
	// example:
	//
	// 10001
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// 主体显示名称，无法解析时为空
	//
	// example:
	//
	// 张三
	PrincipalName *string `json:"principalName,omitempty" xml:"principalName,omitempty"`
	// 主体类型：USER 或 USER_GROUP
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"principalType,omitempty" xml:"principalType,omitempty"`
	// 权限来源：CREATOR 或 DIRECT
	//
	// example:
	//
	// CREATOR
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) String() string {
	return dara.Prettify(s)
}

func (s ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) GoString() string {
	return s.String()
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) GetImmutable() *bool {
	return s.Immutable
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) GetSource() *string {
	return s.Source
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) SetImmutable(v bool) *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals {
	s.Immutable = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) SetPrincipalId(v string) *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals {
	s.PrincipalId = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) SetPrincipalName(v string) *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals {
	s.PrincipalName = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) SetPrincipalType(v string) *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals {
	s.PrincipalType = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) SetSource(v string) *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals {
	s.Source = &v
	return s
}

func (s *ReplaceKnowledgeBaseAdminsResponseBodyPrincipals) Validate() error {
	return dara.Validate(s)
}
