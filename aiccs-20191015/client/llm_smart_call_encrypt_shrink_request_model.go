// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallEncryptShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *LlmSmartCallEncryptShrinkRequest
	GetApplicationCode() *string
	SetCallerNumber(v string) *LlmSmartCallEncryptShrinkRequest
	GetCallerNumber() *string
	SetEncryptCalledNumber(v string) *LlmSmartCallEncryptShrinkRequest
	GetEncryptCalledNumber() *string
	SetOutId(v string) *LlmSmartCallEncryptShrinkRequest
	GetOutId() *string
	SetOwnerId(v int64) *LlmSmartCallEncryptShrinkRequest
	GetOwnerId() *int64
	SetPromptParamShrink(v string) *LlmSmartCallEncryptShrinkRequest
	GetPromptParamShrink() *string
	SetResourceOwnerAccount(v string) *LlmSmartCallEncryptShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *LlmSmartCallEncryptShrinkRequest
	GetResourceOwnerId() *int64
	SetStartWordParamShrink(v string) *LlmSmartCallEncryptShrinkRequest
	GetStartWordParamShrink() *string
}

type LlmSmartCallEncryptShrinkRequest struct {
	// The application code for the large language model.
	//
	// This parameter is required.
	//
	// example:
	//
	// AD******45
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// The caller number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 132******65
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// The encrypted called number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 140*********243
	EncryptCalledNumber *string `json:"EncryptCalledNumber,omitempty" xml:"EncryptCalledNumber,omitempty"`
	// A user-defined ID for the outbound call.
	//
	// example:
	//
	// dsa*******sad
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Parameters for the large language model\\"s prompt.
	//
	// example:
	//
	// {"prompt":"推荐一部电影"}
	PromptParamShrink    *string `json:"PromptParam,omitempty" xml:"PromptParam,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Parameters for the large language model\\"s opening statement.
	//
	// example:
	//
	// {"name":"小明","address":"浙江省杭州市"}
	StartWordParamShrink *string `json:"StartWordParam,omitempty" xml:"StartWordParam,omitempty"`
}

func (s LlmSmartCallEncryptShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallEncryptShrinkRequest) GoString() string {
	return s.String()
}

func (s *LlmSmartCallEncryptShrinkRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *LlmSmartCallEncryptShrinkRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *LlmSmartCallEncryptShrinkRequest) GetEncryptCalledNumber() *string {
	return s.EncryptCalledNumber
}

func (s *LlmSmartCallEncryptShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *LlmSmartCallEncryptShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LlmSmartCallEncryptShrinkRequest) GetPromptParamShrink() *string {
	return s.PromptParamShrink
}

func (s *LlmSmartCallEncryptShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *LlmSmartCallEncryptShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *LlmSmartCallEncryptShrinkRequest) GetStartWordParamShrink() *string {
	return s.StartWordParamShrink
}

func (s *LlmSmartCallEncryptShrinkRequest) SetApplicationCode(v string) *LlmSmartCallEncryptShrinkRequest {
	s.ApplicationCode = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) SetCallerNumber(v string) *LlmSmartCallEncryptShrinkRequest {
	s.CallerNumber = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) SetEncryptCalledNumber(v string) *LlmSmartCallEncryptShrinkRequest {
	s.EncryptCalledNumber = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) SetOutId(v string) *LlmSmartCallEncryptShrinkRequest {
	s.OutId = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) SetOwnerId(v int64) *LlmSmartCallEncryptShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) SetPromptParamShrink(v string) *LlmSmartCallEncryptShrinkRequest {
	s.PromptParamShrink = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) SetResourceOwnerAccount(v string) *LlmSmartCallEncryptShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) SetResourceOwnerId(v int64) *LlmSmartCallEncryptShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) SetStartWordParamShrink(v string) *LlmSmartCallEncryptShrinkRequest {
	s.StartWordParamShrink = &v
	return s
}

func (s *LlmSmartCallEncryptShrinkRequest) Validate() error {
	return dara.Validate(s)
}
