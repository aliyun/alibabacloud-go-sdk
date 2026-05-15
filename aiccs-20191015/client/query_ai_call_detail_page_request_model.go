// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallDetailPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *QueryAiCallDetailPageRequest
	GetBatchId() *string
	SetCallResult(v string) *QueryAiCallDetailPageRequest
	GetCallResult() *string
	SetCalledNumber(v string) *QueryAiCallDetailPageRequest
	GetCalledNumber() *string
	SetDetailIds(v []*int64) *QueryAiCallDetailPageRequest
	GetDetailIds() []*int64
	SetEncryptionType(v int64) *QueryAiCallDetailPageRequest
	GetEncryptionType() *int64
	SetEndCallingTime(v int64) *QueryAiCallDetailPageRequest
	GetEndCallingTime() *int64
	SetEndImportedTime(v int64) *QueryAiCallDetailPageRequest
	GetEndImportedTime() *int64
	SetMajorIntent(v string) *QueryAiCallDetailPageRequest
	GetMajorIntent() *string
	SetMaxConversationDuration(v int64) *QueryAiCallDetailPageRequest
	GetMaxConversationDuration() *int64
	SetMinConversationDuration(v int64) *QueryAiCallDetailPageRequest
	GetMinConversationDuration() *int64
	SetOutId(v string) *QueryAiCallDetailPageRequest
	GetOutId() *string
	SetOwnerId(v int64) *QueryAiCallDetailPageRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryAiCallDetailPageRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryAiCallDetailPageRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *QueryAiCallDetailPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAiCallDetailPageRequest
	GetResourceOwnerId() *int64
	SetStartCallingTime(v int64) *QueryAiCallDetailPageRequest
	GetStartCallingTime() *int64
	SetStartImportedTime(v int64) *QueryAiCallDetailPageRequest
	GetStartImportedTime() *int64
	SetStatus(v int64) *QueryAiCallDetailPageRequest
	GetStatus() *int64
	SetTaskId(v string) *QueryAiCallDetailPageRequest
	GetTaskId() *string
}

type QueryAiCallDetailPageRequest struct {
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
	CalledNumber *string  `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	DetailIds    []*int64 `json:"DetailIds,omitempty" xml:"DetailIds,omitempty" type:"Repeated"`
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

func (s QueryAiCallDetailPageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallDetailPageRequest) GoString() string {
	return s.String()
}

func (s *QueryAiCallDetailPageRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *QueryAiCallDetailPageRequest) GetCallResult() *string {
	return s.CallResult
}

func (s *QueryAiCallDetailPageRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *QueryAiCallDetailPageRequest) GetDetailIds() []*int64 {
	return s.DetailIds
}

func (s *QueryAiCallDetailPageRequest) GetEncryptionType() *int64 {
	return s.EncryptionType
}

func (s *QueryAiCallDetailPageRequest) GetEndCallingTime() *int64 {
	return s.EndCallingTime
}

func (s *QueryAiCallDetailPageRequest) GetEndImportedTime() *int64 {
	return s.EndImportedTime
}

func (s *QueryAiCallDetailPageRequest) GetMajorIntent() *string {
	return s.MajorIntent
}

func (s *QueryAiCallDetailPageRequest) GetMaxConversationDuration() *int64 {
	return s.MaxConversationDuration
}

func (s *QueryAiCallDetailPageRequest) GetMinConversationDuration() *int64 {
	return s.MinConversationDuration
}

func (s *QueryAiCallDetailPageRequest) GetOutId() *string {
	return s.OutId
}

func (s *QueryAiCallDetailPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAiCallDetailPageRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryAiCallDetailPageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryAiCallDetailPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAiCallDetailPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAiCallDetailPageRequest) GetStartCallingTime() *int64 {
	return s.StartCallingTime
}

func (s *QueryAiCallDetailPageRequest) GetStartImportedTime() *int64 {
	return s.StartImportedTime
}

func (s *QueryAiCallDetailPageRequest) GetStatus() *int64 {
	return s.Status
}

func (s *QueryAiCallDetailPageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAiCallDetailPageRequest) SetBatchId(v string) *QueryAiCallDetailPageRequest {
	s.BatchId = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetCallResult(v string) *QueryAiCallDetailPageRequest {
	s.CallResult = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetCalledNumber(v string) *QueryAiCallDetailPageRequest {
	s.CalledNumber = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetDetailIds(v []*int64) *QueryAiCallDetailPageRequest {
	s.DetailIds = v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetEncryptionType(v int64) *QueryAiCallDetailPageRequest {
	s.EncryptionType = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetEndCallingTime(v int64) *QueryAiCallDetailPageRequest {
	s.EndCallingTime = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetEndImportedTime(v int64) *QueryAiCallDetailPageRequest {
	s.EndImportedTime = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetMajorIntent(v string) *QueryAiCallDetailPageRequest {
	s.MajorIntent = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetMaxConversationDuration(v int64) *QueryAiCallDetailPageRequest {
	s.MaxConversationDuration = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetMinConversationDuration(v int64) *QueryAiCallDetailPageRequest {
	s.MinConversationDuration = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetOutId(v string) *QueryAiCallDetailPageRequest {
	s.OutId = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetOwnerId(v int64) *QueryAiCallDetailPageRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetPageNo(v int64) *QueryAiCallDetailPageRequest {
	s.PageNo = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetPageSize(v int64) *QueryAiCallDetailPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetResourceOwnerAccount(v string) *QueryAiCallDetailPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetResourceOwnerId(v int64) *QueryAiCallDetailPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetStartCallingTime(v int64) *QueryAiCallDetailPageRequest {
	s.StartCallingTime = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetStartImportedTime(v int64) *QueryAiCallDetailPageRequest {
	s.StartImportedTime = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetStatus(v int64) *QueryAiCallDetailPageRequest {
	s.Status = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) SetTaskId(v string) *QueryAiCallDetailPageRequest {
	s.TaskId = &v
	return s
}

func (s *QueryAiCallDetailPageRequest) Validate() error {
	return dara.Validate(s)
}
