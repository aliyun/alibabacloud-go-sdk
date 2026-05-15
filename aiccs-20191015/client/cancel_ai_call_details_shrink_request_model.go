// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAiCallDetailsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *CancelAiCallDetailsShrinkRequest
	GetBatchId() *string
	SetDetailIdListShrink(v string) *CancelAiCallDetailsShrinkRequest
	GetDetailIdListShrink() *string
	SetEncryptionType(v int64) *CancelAiCallDetailsShrinkRequest
	GetEncryptionType() *int64
	SetOwnerId(v int64) *CancelAiCallDetailsShrinkRequest
	GetOwnerId() *int64
	SetPhoneNumbersShrink(v string) *CancelAiCallDetailsShrinkRequest
	GetPhoneNumbersShrink() *string
	SetResourceOwnerAccount(v string) *CancelAiCallDetailsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelAiCallDetailsShrinkRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *CancelAiCallDetailsShrinkRequest
	GetTaskId() *string
}

type CancelAiCallDetailsShrinkRequest struct {
	// example:
	//
	// 4253331213*****
	BatchId            *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	DetailIdListShrink *string `json:"DetailIdList,omitempty" xml:"DetailIdList,omitempty"`
	// example:
	//
	// 36
	EncryptionType       *int64  `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumbersShrink   *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1223123123****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelAiCallDetailsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAiCallDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CancelAiCallDetailsShrinkRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *CancelAiCallDetailsShrinkRequest) GetDetailIdListShrink() *string {
	return s.DetailIdListShrink
}

func (s *CancelAiCallDetailsShrinkRequest) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *CancelAiCallDetailsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelAiCallDetailsShrinkRequest) GetPhoneNumbersShrink() *string {
	return s.PhoneNumbersShrink
}

func (s *CancelAiCallDetailsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelAiCallDetailsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelAiCallDetailsShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelAiCallDetailsShrinkRequest) SetBatchId(v string) *CancelAiCallDetailsShrinkRequest {
	s.BatchId = &v
	return s
}

func (s *CancelAiCallDetailsShrinkRequest) SetDetailIdListShrink(v string) *CancelAiCallDetailsShrinkRequest {
	s.DetailIdListShrink = &v
	return s
}

func (s *CancelAiCallDetailsShrinkRequest) SetEncryptionType(v int64) *CancelAiCallDetailsShrinkRequest {
	s.EncryptionType = &v
	return s
}

func (s *CancelAiCallDetailsShrinkRequest) SetOwnerId(v int64) *CancelAiCallDetailsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelAiCallDetailsShrinkRequest) SetPhoneNumbersShrink(v string) *CancelAiCallDetailsShrinkRequest {
	s.PhoneNumbersShrink = &v
	return s
}

func (s *CancelAiCallDetailsShrinkRequest) SetResourceOwnerAccount(v string) *CancelAiCallDetailsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelAiCallDetailsShrinkRequest) SetResourceOwnerId(v int64) *CancelAiCallDetailsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelAiCallDetailsShrinkRequest) SetTaskId(v string) *CancelAiCallDetailsShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *CancelAiCallDetailsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
