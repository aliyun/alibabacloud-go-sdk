// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDypnsSmsVerifyMessageQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CreateDypnsSmsVerifyMessageQueueRequest
	GetOwnerId() *int64
	SetQueryQueueTypes(v string) *CreateDypnsSmsVerifyMessageQueueRequest
	GetQueryQueueTypes() *string
	SetQueueType(v string) *CreateDypnsSmsVerifyMessageQueueRequest
	GetQueueType() *string
	SetRegion(v string) *CreateDypnsSmsVerifyMessageQueueRequest
	GetRegion() *string
	SetResourceOwnerAccount(v string) *CreateDypnsSmsVerifyMessageQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDypnsSmsVerifyMessageQueueRequest
	GetResourceOwnerId() *int64
}

type CreateDypnsSmsVerifyMessageQueueRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// -
	QueryQueueTypes *string `json:"QueryQueueTypes,omitempty" xml:"QueryQueueTypes,omitempty"`
	// example:
	//
	// dypnsSmsVerifyMnsReport
	QueueType *string `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDypnsSmsVerifyMessageQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDypnsSmsVerifyMessageQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) GetQueryQueueTypes() *string {
	return s.QueryQueueTypes
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) GetQueueType() *string {
	return s.QueueType
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) SetOwnerId(v int64) *CreateDypnsSmsVerifyMessageQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) SetQueryQueueTypes(v string) *CreateDypnsSmsVerifyMessageQueueRequest {
	s.QueryQueueTypes = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) SetQueueType(v string) *CreateDypnsSmsVerifyMessageQueueRequest {
	s.QueueType = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) SetRegion(v string) *CreateDypnsSmsVerifyMessageQueueRequest {
	s.Region = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) SetResourceOwnerAccount(v string) *CreateDypnsSmsVerifyMessageQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) SetResourceOwnerId(v int64) *CreateDypnsSmsVerifyMessageQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDypnsSmsVerifyMessageQueueRequest) Validate() error {
	return dara.Validate(s)
}
