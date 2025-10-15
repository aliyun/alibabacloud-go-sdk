// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationFederatedCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationFederatedCredential(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) *GetApplicationFederatedCredentialResponseBody
	GetApplicationFederatedCredential() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential
	SetRequestId(v string) *GetApplicationFederatedCredentialResponseBody
	GetRequestId() *string
}

type GetApplicationFederatedCredentialResponseBody struct {
	ApplicationFederatedCredential *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential `json:"ApplicationFederatedCredential,omitempty" xml:"ApplicationFederatedCredential,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationFederatedCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBody) GetApplicationFederatedCredential() *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	return s.ApplicationFederatedCredential
}

func (s *GetApplicationFederatedCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationFederatedCredentialResponseBody) SetApplicationFederatedCredential(v *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) *GetApplicationFederatedCredentialResponseBody {
	s.ApplicationFederatedCredential = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBody) SetRequestId(v string) *GetApplicationFederatedCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBody) Validate() error {
	if s.ApplicationFederatedCredential != nil {
		if err := s.ApplicationFederatedCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential struct {
	// 应用联邦凭证ID
	//
	// example:
	//
	// afc_aaaaa1111
	ApplicationFederatedCredentialId *string `json:"ApplicationFederatedCredentialId,omitempty" xml:"ApplicationFederatedCredentialId,omitempty"`
	// 应用联邦凭证名称
	//
	// example:
	//
	// test
	ApplicationFederatedCredentialName *string `json:"ApplicationFederatedCredentialName,omitempty" xml:"ApplicationFederatedCredentialName,omitempty"`
	// 应用联邦凭证类型
	//
	// example:
	//
	// oidc
	ApplicationFederatedCredentialType *string `json:"ApplicationFederatedCredentialType,omitempty" xml:"ApplicationFederatedCredentialType,omitempty"`
	// 应用ID
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 属性映射
	AttributeMappings []*GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings `json:"AttributeMappings,omitempty" xml:"AttributeMappings,omitempty" type:"Repeated"`
	// 创建时间
	//
	// example:
	//
	// 1758785994982
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用联邦凭证描述
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用联邦凭证提供者ID
	//
	// example:
	//
	// fcp_asda1dasdxxxx
	FederatedCredentialProviderId *string `json:"FederatedCredentialProviderId,omitempty" xml:"FederatedCredentialProviderId,omitempty"`
	// EAIM 实例ID
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 最近使用时间
	//
	// example:
	//
	// 1758785994982
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// 应用联邦凭证状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 1758785994982
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 验证条件
	//
	// example:
	//
	// Equals(cert.subject.CN, "test")
	VerificationCondition *string `json:"VerificationCondition,omitempty" xml:"VerificationCondition,omitempty"`
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetApplicationFederatedCredentialId() *string {
	return s.ApplicationFederatedCredentialId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetApplicationFederatedCredentialName() *string {
	return s.ApplicationFederatedCredentialName
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetApplicationFederatedCredentialType() *string {
	return s.ApplicationFederatedCredentialType
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetAttributeMappings() []*GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings {
	return s.AttributeMappings
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetFederatedCredentialProviderId() *string {
	return s.FederatedCredentialProviderId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) GetVerificationCondition() *string {
	return s.VerificationCondition
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetApplicationFederatedCredentialId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.ApplicationFederatedCredentialId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetApplicationFederatedCredentialName(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.ApplicationFederatedCredentialName = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetApplicationFederatedCredentialType(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.ApplicationFederatedCredentialType = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetApplicationId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetAttributeMappings(v []*GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.AttributeMappings = v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetCreateTime(v int64) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetDescription(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.Description = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetFederatedCredentialProviderId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.FederatedCredentialProviderId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetInstanceId(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetLastUsedTime(v int64) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.LastUsedTime = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetStatus(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.Status = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetUpdateTime(v int64) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.UpdateTime = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) SetVerificationCondition(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential {
	s.VerificationCondition = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredential) Validate() error {
	if s.AttributeMappings != nil {
		for _, item := range s.AttributeMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings struct {
	// 源值表达式
	//
	// example:
	//
	// Append(client.applicationFederatedCredentialId, ":", cert.subject.CN, ":", cert.serialNumber)
	SourceValueExpression *string `json:"SourceValueExpression,omitempty" xml:"SourceValueExpression,omitempty"`
	// 目标字段
	//
	// example:
	//
	// client.activeSubjectUrn
	TargetField *string `json:"TargetField,omitempty" xml:"TargetField,omitempty"`
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) GetSourceValueExpression() *string {
	return s.SourceValueExpression
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) GetTargetField() *string {
	return s.TargetField
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) SetSourceValueExpression(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings {
	s.SourceValueExpression = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) SetTargetField(v string) *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings {
	s.TargetField = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponseBodyApplicationFederatedCredentialAttributeMappings) Validate() error {
	return dara.Validate(s)
}
