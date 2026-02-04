// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *LlmSmartCallRequest
	GetBizId() *string
	SetBizType(v string) *LlmSmartCallRequest
	GetBizType() *string
	SetCallerNumber(v string) *LlmSmartCallRequest
	GetCallerNumber() *string
	SetProductCode(v string) *LlmSmartCallRequest
	GetProductCode() *string
	SetPromptParam(v string) *LlmSmartCallRequest
	GetPromptParam() *string
	SetScenesCode(v string) *LlmSmartCallRequest
	GetScenesCode() *string
	SetSkillType(v int32) *LlmSmartCallRequest
	GetSkillType() *int32
	SetStartWordParam(v string) *LlmSmartCallRequest
	GetStartWordParam() *string
	SetTenantCode(v string) *LlmSmartCallRequest
	GetTenantCode() *string
}

type LlmSmartCallRequest struct {
	// example:
	//
	// P20210208152920000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.companyreg_cloud
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 13518132662
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// example:
	//
	// xinxuan
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// {\\"prompt\\":\\"{\\\\\\"accountId\\\\\\":\\\\\\"03D31F61DD6512D23F5789A439CE4CA2\\\\\\",\\\\\\"conversationId\\\\\\":\\\\\\"K7nBeth9\\\\\\"}\\"}
	PromptParam *string `json:"PromptParam,omitempty" xml:"PromptParam,omitempty"`
	// example:
	//
	// robotcall_001
	ScenesCode *string `json:"ScenesCode,omitempty" xml:"ScenesCode,omitempty"`
	// example:
	//
	// 1
	SkillType      *int32  `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	StartWordParam *string `json:"StartWordParam,omitempty" xml:"StartWordParam,omitempty"`
	// example:
	//
	// msea
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s LlmSmartCallRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallRequest) GoString() string {
	return s.String()
}

func (s *LlmSmartCallRequest) GetBizId() *string {
	return s.BizId
}

func (s *LlmSmartCallRequest) GetBizType() *string {
	return s.BizType
}

func (s *LlmSmartCallRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *LlmSmartCallRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *LlmSmartCallRequest) GetPromptParam() *string {
	return s.PromptParam
}

func (s *LlmSmartCallRequest) GetScenesCode() *string {
	return s.ScenesCode
}

func (s *LlmSmartCallRequest) GetSkillType() *int32 {
	return s.SkillType
}

func (s *LlmSmartCallRequest) GetStartWordParam() *string {
	return s.StartWordParam
}

func (s *LlmSmartCallRequest) GetTenantCode() *string {
	return s.TenantCode
}

func (s *LlmSmartCallRequest) SetBizId(v string) *LlmSmartCallRequest {
	s.BizId = &v
	return s
}

func (s *LlmSmartCallRequest) SetBizType(v string) *LlmSmartCallRequest {
	s.BizType = &v
	return s
}

func (s *LlmSmartCallRequest) SetCallerNumber(v string) *LlmSmartCallRequest {
	s.CallerNumber = &v
	return s
}

func (s *LlmSmartCallRequest) SetProductCode(v string) *LlmSmartCallRequest {
	s.ProductCode = &v
	return s
}

func (s *LlmSmartCallRequest) SetPromptParam(v string) *LlmSmartCallRequest {
	s.PromptParam = &v
	return s
}

func (s *LlmSmartCallRequest) SetScenesCode(v string) *LlmSmartCallRequest {
	s.ScenesCode = &v
	return s
}

func (s *LlmSmartCallRequest) SetSkillType(v int32) *LlmSmartCallRequest {
	s.SkillType = &v
	return s
}

func (s *LlmSmartCallRequest) SetStartWordParam(v string) *LlmSmartCallRequest {
	s.StartWordParam = &v
	return s
}

func (s *LlmSmartCallRequest) SetTenantCode(v string) *LlmSmartCallRequest {
	s.TenantCode = &v
	return s
}

func (s *LlmSmartCallRequest) Validate() error {
	return dara.Validate(s)
}
