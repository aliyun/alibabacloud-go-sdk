// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmSmartCallEncryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *LlmSmartCallEncryptRequest
	GetApplicationCode() *string
	SetCallerNumber(v string) *LlmSmartCallEncryptRequest
	GetCallerNumber() *string
	SetEncryptCalledNumber(v string) *LlmSmartCallEncryptRequest
	GetEncryptCalledNumber() *string
	SetOutId(v string) *LlmSmartCallEncryptRequest
	GetOutId() *string
	SetOwnerId(v int64) *LlmSmartCallEncryptRequest
	GetOwnerId() *int64
	SetPromptParam(v map[string]interface{}) *LlmSmartCallEncryptRequest
	GetPromptParam() map[string]interface{}
	SetResourceOwnerAccount(v string) *LlmSmartCallEncryptRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *LlmSmartCallEncryptRequest
	GetResourceOwnerId() *int64
	SetStartWordParam(v map[string]interface{}) *LlmSmartCallEncryptRequest
	GetStartWordParam() map[string]interface{}
}

type LlmSmartCallEncryptRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ADDFA32145
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ADDFA32145
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 140432432432243
	EncryptCalledNumber *string `json:"EncryptCalledNumber,omitempty" xml:"EncryptCalledNumber,omitempty"`
	// example:
	//
	// dsadsaasfdsad
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// {}
	PromptParam          map[string]interface{} `json:"PromptParam,omitempty" xml:"PromptParam,omitempty"`
	ResourceOwnerAccount *string                `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartWordParam       map[string]interface{} `json:"StartWordParam,omitempty" xml:"StartWordParam,omitempty"`
}

func (s LlmSmartCallEncryptRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmSmartCallEncryptRequest) GoString() string {
	return s.String()
}

func (s *LlmSmartCallEncryptRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *LlmSmartCallEncryptRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *LlmSmartCallEncryptRequest) GetEncryptCalledNumber() *string {
	return s.EncryptCalledNumber
}

func (s *LlmSmartCallEncryptRequest) GetOutId() *string {
	return s.OutId
}

func (s *LlmSmartCallEncryptRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LlmSmartCallEncryptRequest) GetPromptParam() map[string]interface{} {
	return s.PromptParam
}

func (s *LlmSmartCallEncryptRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *LlmSmartCallEncryptRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *LlmSmartCallEncryptRequest) GetStartWordParam() map[string]interface{} {
	return s.StartWordParam
}

func (s *LlmSmartCallEncryptRequest) SetApplicationCode(v string) *LlmSmartCallEncryptRequest {
	s.ApplicationCode = &v
	return s
}

func (s *LlmSmartCallEncryptRequest) SetCallerNumber(v string) *LlmSmartCallEncryptRequest {
	s.CallerNumber = &v
	return s
}

func (s *LlmSmartCallEncryptRequest) SetEncryptCalledNumber(v string) *LlmSmartCallEncryptRequest {
	s.EncryptCalledNumber = &v
	return s
}

func (s *LlmSmartCallEncryptRequest) SetOutId(v string) *LlmSmartCallEncryptRequest {
	s.OutId = &v
	return s
}

func (s *LlmSmartCallEncryptRequest) SetOwnerId(v int64) *LlmSmartCallEncryptRequest {
	s.OwnerId = &v
	return s
}

func (s *LlmSmartCallEncryptRequest) SetPromptParam(v map[string]interface{}) *LlmSmartCallEncryptRequest {
	s.PromptParam = v
	return s
}

func (s *LlmSmartCallEncryptRequest) SetResourceOwnerAccount(v string) *LlmSmartCallEncryptRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LlmSmartCallEncryptRequest) SetResourceOwnerId(v int64) *LlmSmartCallEncryptRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LlmSmartCallEncryptRequest) SetStartWordParam(v map[string]interface{}) *LlmSmartCallEncryptRequest {
	s.StartWordParam = v
	return s
}

func (s *LlmSmartCallEncryptRequest) Validate() error {
	return dara.Validate(s)
}
