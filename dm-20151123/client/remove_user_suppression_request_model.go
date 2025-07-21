// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserSuppressionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *RemoveUserSuppressionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RemoveUserSuppressionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveUserSuppressionRequest
	GetResourceOwnerId() *int64
	SetSuppressionIds(v string) *RemoveUserSuppressionRequest
	GetSuppressionIds() *string
}

type RemoveUserSuppressionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 59511
	SuppressionIds *string `json:"SuppressionIds,omitempty" xml:"SuppressionIds,omitempty"`
}

func (s RemoveUserSuppressionRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserSuppressionRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserSuppressionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveUserSuppressionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveUserSuppressionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveUserSuppressionRequest) GetSuppressionIds() *string {
	return s.SuppressionIds
}

func (s *RemoveUserSuppressionRequest) SetOwnerId(v int64) *RemoveUserSuppressionRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveUserSuppressionRequest) SetResourceOwnerAccount(v string) *RemoveUserSuppressionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveUserSuppressionRequest) SetResourceOwnerId(v int64) *RemoveUserSuppressionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveUserSuppressionRequest) SetSuppressionIds(v string) *RemoveUserSuppressionRequest {
	s.SuppressionIds = &v
	return s
}

func (s *RemoveUserSuppressionRequest) Validate() error {
	return dara.Validate(s)
}
