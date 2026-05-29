// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableTtsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListAvailableTtsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAvailableTtsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAvailableTtsRequest
	GetResourceOwnerId() *int64
	SetTtsVoiceCode(v string) *ListAvailableTtsRequest
	GetTtsVoiceCode() *string
	SetVoiceType(v string) *ListAvailableTtsRequest
	GetVoiceType() *string
}

type ListAvailableTtsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 复刻音色编码
	//
	// example:
	//
	// V123456789
	TtsVoiceCode *string `json:"TtsVoiceCode,omitempty" xml:"TtsVoiceCode,omitempty"`
	// example:
	//
	// 示例值示例值
	VoiceType *string `json:"VoiceType,omitempty" xml:"VoiceType,omitempty"`
}

func (s ListAvailableTtsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableTtsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableTtsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAvailableTtsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAvailableTtsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAvailableTtsRequest) GetTtsVoiceCode() *string {
	return s.TtsVoiceCode
}

func (s *ListAvailableTtsRequest) GetVoiceType() *string {
	return s.VoiceType
}

func (s *ListAvailableTtsRequest) SetOwnerId(v int64) *ListAvailableTtsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAvailableTtsRequest) SetResourceOwnerAccount(v string) *ListAvailableTtsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAvailableTtsRequest) SetResourceOwnerId(v int64) *ListAvailableTtsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAvailableTtsRequest) SetTtsVoiceCode(v string) *ListAvailableTtsRequest {
	s.TtsVoiceCode = &v
	return s
}

func (s *ListAvailableTtsRequest) SetVoiceType(v string) *ListAvailableTtsRequest {
	s.VoiceType = &v
	return s
}

func (s *ListAvailableTtsRequest) Validate() error {
	return dara.Validate(s)
}
