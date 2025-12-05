// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTokenForMnsQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageType(v string) *QueryTokenForMnsQueueRequest
	GetMessageType() *string
	SetOwnerId(v int64) *QueryTokenForMnsQueueRequest
	GetOwnerId() *int64
	SetQueueName(v string) *QueryTokenForMnsQueueRequest
	GetQueueName() *string
	SetResourceOwnerAccount(v string) *QueryTokenForMnsQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryTokenForMnsQueueRequest
	GetResourceOwnerId() *int64
}

type QueryTokenForMnsQueueRequest struct {
	// This parameter is required.
	MessageType          *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QueueName            *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryTokenForMnsQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTokenForMnsQueueRequest) GoString() string {
	return s.String()
}

func (s *QueryTokenForMnsQueueRequest) GetMessageType() *string {
	return s.MessageType
}

func (s *QueryTokenForMnsQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryTokenForMnsQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *QueryTokenForMnsQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryTokenForMnsQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryTokenForMnsQueueRequest) SetMessageType(v string) *QueryTokenForMnsQueueRequest {
	s.MessageType = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) SetOwnerId(v int64) *QueryTokenForMnsQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) SetQueueName(v string) *QueryTokenForMnsQueueRequest {
	s.QueueName = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) SetResourceOwnerAccount(v string) *QueryTokenForMnsQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) SetResourceOwnerId(v int64) *QueryTokenForMnsQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryTokenForMnsQueueRequest) Validate() error {
	return dara.Validate(s)
}
