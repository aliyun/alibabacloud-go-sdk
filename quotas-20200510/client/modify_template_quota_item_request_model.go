// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateQuotaItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesireValue(v float32) *ModifyTemplateQuotaItemRequest
	GetDesireValue() *float32
	SetDimensions(v []*ModifyTemplateQuotaItemRequestDimensions) *ModifyTemplateQuotaItemRequest
	GetDimensions() []*ModifyTemplateQuotaItemRequestDimensions
	SetEffectiveTime(v string) *ModifyTemplateQuotaItemRequest
	GetEffectiveTime() *string
	SetEnvLanguage(v string) *ModifyTemplateQuotaItemRequest
	GetEnvLanguage() *string
	SetExpireTime(v string) *ModifyTemplateQuotaItemRequest
	GetExpireTime() *string
	SetId(v string) *ModifyTemplateQuotaItemRequest
	GetId() *string
	SetNoticeType(v int64) *ModifyTemplateQuotaItemRequest
	GetNoticeType() *int64
	SetProductCode(v string) *ModifyTemplateQuotaItemRequest
	GetProductCode() *string
	SetQuotaActionCode(v string) *ModifyTemplateQuotaItemRequest
	GetQuotaActionCode() *string
	SetQuotaCategory(v string) *ModifyTemplateQuotaItemRequest
	GetQuotaCategory() *string
}

type ModifyTemplateQuotaItemRequest struct {
	// example:
	//
	// 804
	DesireValue *float32                                    `json:"DesireValue,omitempty" xml:"DesireValue,omitempty"`
	Dimensions  []*ModifyTemplateQuotaItemRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-01-19T09:25:56Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// example:
	//
	// zh
	EnvLanguage *string `json:"EnvLanguage,omitempty" xml:"EnvLanguage,omitempty"`
	// example:
	//
	// 2021-01-20T09:25:56Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	NoticeType *int64 `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
}

func (s ModifyTemplateQuotaItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateQuotaItemRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemRequest) GetDesireValue() *float32 {
	return s.DesireValue
}

func (s *ModifyTemplateQuotaItemRequest) GetDimensions() []*ModifyTemplateQuotaItemRequestDimensions {
	return s.Dimensions
}

func (s *ModifyTemplateQuotaItemRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyTemplateQuotaItemRequest) GetEnvLanguage() *string {
	return s.EnvLanguage
}

func (s *ModifyTemplateQuotaItemRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModifyTemplateQuotaItemRequest) GetId() *string {
	return s.Id
}

func (s *ModifyTemplateQuotaItemRequest) GetNoticeType() *int64 {
	return s.NoticeType
}

func (s *ModifyTemplateQuotaItemRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ModifyTemplateQuotaItemRequest) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ModifyTemplateQuotaItemRequest) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ModifyTemplateQuotaItemRequest) SetDesireValue(v float32) *ModifyTemplateQuotaItemRequest {
	s.DesireValue = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetDimensions(v []*ModifyTemplateQuotaItemRequestDimensions) *ModifyTemplateQuotaItemRequest {
	s.Dimensions = v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetEffectiveTime(v string) *ModifyTemplateQuotaItemRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetEnvLanguage(v string) *ModifyTemplateQuotaItemRequest {
	s.EnvLanguage = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetExpireTime(v string) *ModifyTemplateQuotaItemRequest {
	s.ExpireTime = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetId(v string) *ModifyTemplateQuotaItemRequest {
	s.Id = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetNoticeType(v int64) *ModifyTemplateQuotaItemRequest {
	s.NoticeType = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetProductCode(v string) *ModifyTemplateQuotaItemRequest {
	s.ProductCode = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetQuotaActionCode(v string) *ModifyTemplateQuotaItemRequest {
	s.QuotaActionCode = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) SetQuotaCategory(v string) *ModifyTemplateQuotaItemRequest {
	s.QuotaCategory = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyTemplateQuotaItemRequestDimensions struct {
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyTemplateQuotaItemRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateQuotaItemRequestDimensions) GoString() string {
	return s.String()
}

func (s *ModifyTemplateQuotaItemRequestDimensions) GetKey() *string {
	return s.Key
}

func (s *ModifyTemplateQuotaItemRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *ModifyTemplateQuotaItemRequestDimensions) SetKey(v string) *ModifyTemplateQuotaItemRequestDimensions {
	s.Key = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequestDimensions) SetValue(v string) *ModifyTemplateQuotaItemRequestDimensions {
	s.Value = &v
	return s
}

func (s *ModifyTemplateQuotaItemRequestDimensions) Validate() error {
	return dara.Validate(s)
}
