// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncBusinessAppHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *SyncBusinessAppHistoryRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *SyncBusinessAppHistoryRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *SyncBusinessAppHistoryRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *SyncBusinessAppHistoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SyncBusinessAppHistoryRequest
	GetResourceOwnerId() *int64
}

type SyncBusinessAppHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-x***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8613800**
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SyncBusinessAppHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncBusinessAppHistoryRequest) GoString() string {
	return s.String()
}

func (s *SyncBusinessAppHistoryRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SyncBusinessAppHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SyncBusinessAppHistoryRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *SyncBusinessAppHistoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SyncBusinessAppHistoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SyncBusinessAppHistoryRequest) SetCustSpaceId(v string) *SyncBusinessAppHistoryRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SyncBusinessAppHistoryRequest) SetOwnerId(v int64) *SyncBusinessAppHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *SyncBusinessAppHistoryRequest) SetPhoneNumber(v string) *SyncBusinessAppHistoryRequest {
	s.PhoneNumber = &v
	return s
}

func (s *SyncBusinessAppHistoryRequest) SetResourceOwnerAccount(v string) *SyncBusinessAppHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SyncBusinessAppHistoryRequest) SetResourceOwnerId(v int64) *SyncBusinessAppHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SyncBusinessAppHistoryRequest) Validate() error {
	return dara.Validate(s)
}
