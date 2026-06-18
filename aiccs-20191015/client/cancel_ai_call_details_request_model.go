// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAiCallDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *CancelAiCallDetailsRequest
	GetBatchId() *string
	SetDetailIdList(v []*string) *CancelAiCallDetailsRequest
	GetDetailIdList() []*string
	SetEncryptionType(v int64) *CancelAiCallDetailsRequest
	GetEncryptionType() *int64
	SetOwnerId(v int64) *CancelAiCallDetailsRequest
	GetOwnerId() *int64
	SetPhoneNumbers(v []*string) *CancelAiCallDetailsRequest
	GetPhoneNumbers() []*string
	SetResourceOwnerAccount(v string) *CancelAiCallDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelAiCallDetailsRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *CancelAiCallDetailsRequest
	GetTaskId() *string
}

type CancelAiCallDetailsRequest struct {
	// The batch ID. This ID is returned when you import callee data. You can find the task batch ID on the **Call Task Management**>**Details**>**Execution Record*	- page, or use the import batch ID that is returned by the [ImportTaskNumberDatas](https://help.aliyun.com/document_detail/2926815.html) API operation. If DetailIdList is specified, this parameter is invalid.
	//
	// example:
	//
	// 4253331213*****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// A list of detail IDs. If you specify this parameter, `BatchId` and `PhoneNumbers` are ignored.
	DetailIdList []*string `json:"DetailIdList,omitempty" xml:"DetailIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 36
	EncryptionType *int64 `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	OwnerId        *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A list of phone numbers. This parameter takes effect only when `BatchId` is also specified.
	PhoneNumbers         []*string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task ID. You can view the task ID on the **Call Task Management*	- page or obtain it by calling the [QueryAiCallTaskPage](https://help.aliyun.com/document_detail/2926799.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelAiCallDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAiCallDetailsRequest) GoString() string {
	return s.String()
}

func (s *CancelAiCallDetailsRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *CancelAiCallDetailsRequest) GetDetailIdList() []*string {
	return s.DetailIdList
}

func (s *CancelAiCallDetailsRequest) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *CancelAiCallDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelAiCallDetailsRequest) GetPhoneNumbers() []*string {
	return s.PhoneNumbers
}

func (s *CancelAiCallDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelAiCallDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelAiCallDetailsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelAiCallDetailsRequest) SetBatchId(v string) *CancelAiCallDetailsRequest {
	s.BatchId = &v
	return s
}

func (s *CancelAiCallDetailsRequest) SetDetailIdList(v []*string) *CancelAiCallDetailsRequest {
	s.DetailIdList = v
	return s
}

func (s *CancelAiCallDetailsRequest) SetEncryptionType(v int64) *CancelAiCallDetailsRequest {
	s.EncryptionType = &v
	return s
}

func (s *CancelAiCallDetailsRequest) SetOwnerId(v int64) *CancelAiCallDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelAiCallDetailsRequest) SetPhoneNumbers(v []*string) *CancelAiCallDetailsRequest {
	s.PhoneNumbers = v
	return s
}

func (s *CancelAiCallDetailsRequest) SetResourceOwnerAccount(v string) *CancelAiCallDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelAiCallDetailsRequest) SetResourceOwnerId(v int64) *CancelAiCallDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelAiCallDetailsRequest) SetTaskId(v string) *CancelAiCallDetailsRequest {
	s.TaskId = &v
	return s
}

func (s *CancelAiCallDetailsRequest) Validate() error {
	return dara.Validate(s)
}
