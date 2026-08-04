// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillingCostBreakdownRowDTO interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *BillingCostBreakdownRowDTO
	GetApiKeyId() *int64
	SetApiKeyName(v string) *BillingCostBreakdownRowDTO
	GetApiKeyName() *string
	SetBillingType(v string) *BillingCostBreakdownRowDTO
	GetBillingType() *string
	SetClientId(v int64) *BillingCostBreakdownRowDTO
	GetClientId() *int64
	SetClientName(v string) *BillingCostBreakdownRowDTO
	GetClientName() *string
	SetDimValues(v string) *BillingCostBreakdownRowDTO
	GetDimValues() *string
	SetMemberUserId(v int64) *BillingCostBreakdownRowDTO
	GetMemberUserId() *int64
	SetMemberUserName(v string) *BillingCostBreakdownRowDTO
	GetMemberUserName() *string
	SetModelCode(v string) *BillingCostBreakdownRowDTO
	GetModelCode() *string
	SetModelId(v int64) *BillingCostBreakdownRowDTO
	GetModelId() *int64
	SetModelName(v string) *BillingCostBreakdownRowDTO
	GetModelName() *string
	SetModelSymbol(v string) *BillingCostBreakdownRowDTO
	GetModelSymbol() *string
	SetModelType(v string) *BillingCostBreakdownRowDTO
	GetModelType() *string
	SetModelVersion(v int32) *BillingCostBreakdownRowDTO
	GetModelVersion() *int32
	SetPayableAmount(v float32) *BillingCostBreakdownRowDTO
	GetPayableAmount() *float32
	SetSummaryTime(v int64) *BillingCostBreakdownRowDTO
	GetSummaryTime() *int64
	SetTiers(v []*BillingBillTierDTO) *BillingCostBreakdownRowDTO
	GetTiers() []*BillingBillTierDTO
	SetValues(v string) *BillingCostBreakdownRowDTO
	GetValues() *string
}

type BillingCostBreakdownRowDTO struct {
	// example:
	//
	// 0
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// example:
	//
	// 默认密钥
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// Billing type; currently fixed as total_amount
	//
	// example:
	//
	// total_amount
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// Department ID; 0 indicates no associated department
	//
	// example:
	//
	// 0
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// Department name
	//
	// example:
	//
	// 研发部
	ClientName *string `json:"clientName,omitempty" xml:"clientName,omitempty"`
	// example:
	//
	// {"billing_version": "v1"}
	DimValues *string `json:"dimValues,omitempty" xml:"dimValues,omitempty"`
	// example:
	//
	// 30001
	MemberUserId *int64 `json:"memberUserId,omitempty" xml:"memberUserId,omitempty"`
	// example:
	//
	// 张三
	MemberUserName *string `json:"memberUserName,omitempty" xml:"memberUserName,omitempty"`
	// Model identifier
	//
	// example:
	//
	// qwen-plus
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// Model ID
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// Model name
	//
	// example:
	//
	// 通义千问-Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// qwen
	ModelSymbol *string `json:"modelSymbol,omitempty" xml:"modelSymbol,omitempty"`
	// Statistics category
	//
	// example:
	//
	// llm
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// example:
	//
	// 1
	ModelVersion *int32 `json:"modelVersion,omitempty" xml:"modelVersion,omitempty"`
	// Payable amount, rounded to 8 decimal places
	//
	// example:
	//
	// 0.00012800
	PayableAmount *float32 `json:"payableAmount,omitempty" xml:"payableAmount,omitempty"`
	// Statistics time point, Unix timestamp (seconds)
	//
	// example:
	//
	// 1700000000
	SummaryTime *int64 `json:"summaryTime,omitempty" xml:"summaryTime,omitempty"`
	// example:
	//
	// []
	Tiers []*BillingBillTierDTO `json:"tiers,omitempty" xml:"tiers,omitempty" type:"Repeated"`
	// example:
	//
	// {"input_tokens": 512000, "output_tokens": 256000}
	Values *string `json:"values,omitempty" xml:"values,omitempty"`
}

func (s BillingCostBreakdownRowDTO) String() string {
	return dara.Prettify(s)
}

func (s BillingCostBreakdownRowDTO) GoString() string {
	return s.String()
}

func (s *BillingCostBreakdownRowDTO) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *BillingCostBreakdownRowDTO) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *BillingCostBreakdownRowDTO) GetBillingType() *string {
	return s.BillingType
}

func (s *BillingCostBreakdownRowDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *BillingCostBreakdownRowDTO) GetClientName() *string {
	return s.ClientName
}

func (s *BillingCostBreakdownRowDTO) GetDimValues() *string {
	return s.DimValues
}

func (s *BillingCostBreakdownRowDTO) GetMemberUserId() *int64 {
	return s.MemberUserId
}

func (s *BillingCostBreakdownRowDTO) GetMemberUserName() *string {
	return s.MemberUserName
}

func (s *BillingCostBreakdownRowDTO) GetModelCode() *string {
	return s.ModelCode
}

func (s *BillingCostBreakdownRowDTO) GetModelId() *int64 {
	return s.ModelId
}

func (s *BillingCostBreakdownRowDTO) GetModelName() *string {
	return s.ModelName
}

func (s *BillingCostBreakdownRowDTO) GetModelSymbol() *string {
	return s.ModelSymbol
}

func (s *BillingCostBreakdownRowDTO) GetModelType() *string {
	return s.ModelType
}

func (s *BillingCostBreakdownRowDTO) GetModelVersion() *int32 {
	return s.ModelVersion
}

func (s *BillingCostBreakdownRowDTO) GetPayableAmount() *float32 {
	return s.PayableAmount
}

func (s *BillingCostBreakdownRowDTO) GetSummaryTime() *int64 {
	return s.SummaryTime
}

func (s *BillingCostBreakdownRowDTO) GetTiers() []*BillingBillTierDTO {
	return s.Tiers
}

func (s *BillingCostBreakdownRowDTO) GetValues() *string {
	return s.Values
}

func (s *BillingCostBreakdownRowDTO) SetApiKeyId(v int64) *BillingCostBreakdownRowDTO {
	s.ApiKeyId = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetApiKeyName(v string) *BillingCostBreakdownRowDTO {
	s.ApiKeyName = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetBillingType(v string) *BillingCostBreakdownRowDTO {
	s.BillingType = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetClientId(v int64) *BillingCostBreakdownRowDTO {
	s.ClientId = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetClientName(v string) *BillingCostBreakdownRowDTO {
	s.ClientName = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetDimValues(v string) *BillingCostBreakdownRowDTO {
	s.DimValues = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetMemberUserId(v int64) *BillingCostBreakdownRowDTO {
	s.MemberUserId = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetMemberUserName(v string) *BillingCostBreakdownRowDTO {
	s.MemberUserName = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetModelCode(v string) *BillingCostBreakdownRowDTO {
	s.ModelCode = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetModelId(v int64) *BillingCostBreakdownRowDTO {
	s.ModelId = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetModelName(v string) *BillingCostBreakdownRowDTO {
	s.ModelName = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetModelSymbol(v string) *BillingCostBreakdownRowDTO {
	s.ModelSymbol = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetModelType(v string) *BillingCostBreakdownRowDTO {
	s.ModelType = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetModelVersion(v int32) *BillingCostBreakdownRowDTO {
	s.ModelVersion = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetPayableAmount(v float32) *BillingCostBreakdownRowDTO {
	s.PayableAmount = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetSummaryTime(v int64) *BillingCostBreakdownRowDTO {
	s.SummaryTime = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetTiers(v []*BillingBillTierDTO) *BillingCostBreakdownRowDTO {
	s.Tiers = v
	return s
}

func (s *BillingCostBreakdownRowDTO) SetValues(v string) *BillingCostBreakdownRowDTO {
	s.Values = &v
	return s
}

func (s *BillingCostBreakdownRowDTO) Validate() error {
	if s.Tiers != nil {
		for _, item := range s.Tiers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
