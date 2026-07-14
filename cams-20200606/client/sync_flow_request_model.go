// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *SyncFlowRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *SyncFlowRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SyncFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SyncFlowRequest
	GetResourceOwnerId() *int64
}

type SyncFlowRequest struct {
	// 通道实例
	//
	// This parameter is required.
	//
	// example:
	//
	// cams-kei****
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SyncFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncFlowRequest) GoString() string {
	return s.String()
}

func (s *SyncFlowRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SyncFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SyncFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SyncFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SyncFlowRequest) SetCustSpaceId(v string) *SyncFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SyncFlowRequest) SetOwnerId(v int64) *SyncFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *SyncFlowRequest) SetResourceOwnerAccount(v string) *SyncFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SyncFlowRequest) SetResourceOwnerId(v int64) *SyncFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SyncFlowRequest) Validate() error {
	return dara.Validate(s)
}
