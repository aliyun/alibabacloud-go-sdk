// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappEmbedSignUpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputToken(v string) *ChatappEmbedSignUpRequest
	GetInputToken() *string
	SetOwnerId(v int64) *ChatappEmbedSignUpRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ChatappEmbedSignUpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChatappEmbedSignUpRequest
	GetResourceOwnerId() *int64
}

type ChatappEmbedSignUpRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	InputToken           *string `json:"InputToken,omitempty" xml:"InputToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ChatappEmbedSignUpRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatappEmbedSignUpRequest) GoString() string {
	return s.String()
}

func (s *ChatappEmbedSignUpRequest) GetInputToken() *string {
	return s.InputToken
}

func (s *ChatappEmbedSignUpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatappEmbedSignUpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChatappEmbedSignUpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChatappEmbedSignUpRequest) SetInputToken(v string) *ChatappEmbedSignUpRequest {
	s.InputToken = &v
	return s
}

func (s *ChatappEmbedSignUpRequest) SetOwnerId(v int64) *ChatappEmbedSignUpRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatappEmbedSignUpRequest) SetResourceOwnerAccount(v string) *ChatappEmbedSignUpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChatappEmbedSignUpRequest) SetResourceOwnerId(v int64) *ChatappEmbedSignUpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChatappEmbedSignUpRequest) Validate() error {
	return dara.Validate(s)
}
