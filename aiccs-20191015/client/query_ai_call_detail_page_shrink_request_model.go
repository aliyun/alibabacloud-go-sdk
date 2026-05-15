// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallDetailPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *QueryAiCallDetailPageShrinkRequest
	GetBatchId() *string
	SetCallResult(v string) *QueryAiCallDetailPageShrinkRequest
	GetCallResult() *string
	SetCalledNumber(v string) *QueryAiCallDetailPageShrinkRequest
	GetCalledNumber() *string
	SetDetailIdsShrink(v string) *QueryAiCallDetailPageShrinkRequest
	GetDetailIdsShrink() *string
	SetEncryptionType(v int64) *QueryAiCallDetailPageShrinkRequest
	GetEncryptionType() *int64
	SetEndCallingTime(v int64) *QueryAiCallDetailPageShrinkRequest
	GetEndCallingTime() *int64
	SetEndImportedTime(v int64) *QueryAiCallDetailPageShrinkRequest
	GetEndImportedTime() *int64
	SetMajorIntent(v string) *QueryAiCallDetailPageShrinkRequest
	GetMajorIntent() *string
	SetMaxConversationDuration(v int64) *QueryAiCallDetailPageShrinkRequest
	GetMaxConversationDuration() *int64
	SetMinConversationDuration(v int64) *QueryAiCallDetailPageShrinkRequest
	GetMinConversationDuration() *int64
	SetOutId(v string) *QueryAiCallDetailPageShrinkRequest
	GetOutId() *string
	SetOwnerId(v int64) *QueryAiCallDetailPageShrinkRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryAiCallDetailPageShrinkRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryAiCallDetailPageShrinkRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *QueryAiCallDetailPageShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAiCallDetailPageShrinkRequest
	GetResourceOwnerId() *int64
	SetStartCallingTime(v int64) *QueryAiCallDetailPageShrinkRequest
	GetStartCallingTime() *int64
	SetStartImportedTime(v int64) *QueryAiCallDetailPageShrinkRequest
	GetStartImportedTime() *int64
	SetStatus(v int64) *QueryAiCallDetailPageShrinkRequest
	GetStatus() *int64
	SetTaskId(v string) *QueryAiCallDetailPageShrinkRequest
	GetTaskId() *string
}

type QueryAiCallDetailPageShrinkRequest struct {
	// example:
	//
	// 1212131231****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// ANSWERED
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// example:
	//
	// 053714454****
	CalledNumber    *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	DetailIdsShrink *string `json:"DetailIds,omitempty" xml:"DetailIds,omitempty"`
	// example:
	//
	// 73
	EncryptionType *int64 `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// example:
	//
	// 1748948749000
	EndCallingTime *int64 `json:"EndCallingTime,omitempty" xml:"EndCallingTime,omitempty"`
	// example:
	//
	// 1748948749000
	EndImportedTime *int64 `json:"EndImportedTime,omitempty" xml:"EndImportedTime,omitempty"`
	// example:
	//
	// A
	MajorIntent *string `json:"MajorIntent,omitempty" xml:"MajorIntent,omitempty"`
	// example:
	//
	// 20
	MaxConversationDuration *int64 `json:"MaxConversationDuration,omitempty" xml:"MaxConversationDuration,omitempty"`
	// example:
	//
	// 0
	MinConversationDuration *int64 `json:"MinConversationDuration,omitempty" xml:"MinConversationDuration,omitempty"`
	// example:
	//
	// 示例值示例值
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1748948749000
	StartCallingTime *int64 `json:"StartCallingTime,omitempty" xml:"StartCallingTime,omitempty"`
	// example:
	//
	// 1748948749000
	StartImportedTime *int64 `json:"StartImportedTime,omitempty" xml:"StartImportedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1212131231****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryAiCallDetailPageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallDetailPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAiCallDetailPageShrinkRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *QueryAiCallDetailPageShrinkRequest) GetCallResult() *string {
	return s.CallResult
}

func (s *QueryAiCallDetailPageShrinkRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *QueryAiCallDetailPageShrinkRequest) GetDetailIdsShrink() *string {
	return s.DetailIdsShrink
}

func (s *QueryAiCallDetailPageShrinkRequest) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *QueryAiCallDetailPageShrinkRequest) GetEndCallingTime() *int64 {
	return s.EndCallingTime
}

func (s *QueryAiCallDetailPageShrinkRequest) GetEndImportedTime() *int64 {
	return s.EndImportedTime
}

func (s *QueryAiCallDetailPageShrinkRequest) GetMajorIntent() *string {
	return s.MajorIntent
}

func (s *QueryAiCallDetailPageShrinkRequest) GetMaxConversationDuration() *int64 {
	return s.MaxConversationDuration
}

func (s *QueryAiCallDetailPageShrinkRequest) GetMinConversationDuration() *int64 {
	return s.MinConversationDuration
}

func (s *QueryAiCallDetailPageShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *QueryAiCallDetailPageShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAiCallDetailPageShrinkRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryAiCallDetailPageShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryAiCallDetailPageShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAiCallDetailPageShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAiCallDetailPageShrinkRequest) GetStartCallingTime() *int64 {
	return s.StartCallingTime
}

func (s *QueryAiCallDetailPageShrinkRequest) GetStartImportedTime() *int64 {
	return s.StartImportedTime
}

func (s *QueryAiCallDetailPageShrinkRequest) GetStatus() *int64 {
	return s.Status
}

func (s *QueryAiCallDetailPageShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAiCallDetailPageShrinkRequest) SetBatchId(v string) *QueryAiCallDetailPageShrinkRequest {
	s.BatchId = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetCallResult(v string) *QueryAiCallDetailPageShrinkRequest {
	s.CallResult = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetCalledNumber(v string) *QueryAiCallDetailPageShrinkRequest {
	s.CalledNumber = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetDetailIdsShrink(v string) *QueryAiCallDetailPageShrinkRequest {
	s.DetailIdsShrink = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetEncryptionType(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.EncryptionType = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetEndCallingTime(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.EndCallingTime = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetEndImportedTime(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.EndImportedTime = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetMajorIntent(v string) *QueryAiCallDetailPageShrinkRequest {
	s.MajorIntent = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetMaxConversationDuration(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.MaxConversationDuration = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetMinConversationDuration(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.MinConversationDuration = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetOutId(v string) *QueryAiCallDetailPageShrinkRequest {
	s.OutId = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetOwnerId(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetPageNo(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetPageSize(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetResourceOwnerAccount(v string) *QueryAiCallDetailPageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetResourceOwnerId(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetStartCallingTime(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.StartCallingTime = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetStartImportedTime(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.StartImportedTime = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetStatus(v int64) *QueryAiCallDetailPageShrinkRequest {
	s.Status = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) SetTaskId(v string) *QueryAiCallDetailPageShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *QueryAiCallDetailPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
