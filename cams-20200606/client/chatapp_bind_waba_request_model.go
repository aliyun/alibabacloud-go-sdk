// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappBindWabaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ChatappBindWabaRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ChatappBindWabaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChatappBindWabaRequest
	GetResourceOwnerId() *int64
	SetWabaId(v string) *ChatappBindWabaRequest
	GetWabaId() *string
}

type ChatappBindWabaRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 33993***
	WabaId *string `json:"WabaId,omitempty" xml:"WabaId,omitempty"`
}

func (s ChatappBindWabaRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatappBindWabaRequest) GoString() string {
	return s.String()
}

func (s *ChatappBindWabaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatappBindWabaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChatappBindWabaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChatappBindWabaRequest) GetWabaId() *string {
	return s.WabaId
}

func (s *ChatappBindWabaRequest) SetOwnerId(v int64) *ChatappBindWabaRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatappBindWabaRequest) SetResourceOwnerAccount(v string) *ChatappBindWabaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChatappBindWabaRequest) SetResourceOwnerId(v int64) *ChatappBindWabaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChatappBindWabaRequest) SetWabaId(v string) *ChatappBindWabaRequest {
	s.WabaId = &v
	return s
}

func (s *ChatappBindWabaRequest) Validate() error {
	return dara.Validate(s)
}
