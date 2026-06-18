// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatRecordDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListChatRecordDetailRequest
	GetClientToken() *string
	SetCloseTimeEnd(v int64) *ListChatRecordDetailRequest
	GetCloseTimeEnd() *int64
	SetCloseTimeStart(v int64) *ListChatRecordDetailRequest
	GetCloseTimeStart() *int64
	SetCurrentPage(v int32) *ListChatRecordDetailRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListChatRecordDetailRequest
	GetInstanceId() *string
	SetPageSize(v int32) *ListChatRecordDetailRequest
	GetPageSize() *int32
}

type ListChatRecordDetailRequest struct {
	// Unique ID for the customer request. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Right boundary of the time range for session end time. Format: UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1614582000000
	CloseTimeEnd *int64 `json:"CloseTimeEnd,omitempty" xml:"CloseTimeEnd,omitempty"`
	// Left boundary of the time range for session end time. Format: UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1614578400000
	CloseTimeStart *int64 `json:"CloseTimeStart,omitempty" xml:"CloseTimeStart,omitempty"`
	// Current page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Artificial Intelligence Cloud Call Service (AICCS) instance ID. You can obtain it from the Artificial Intelligence Cloud Call Service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page size. Default value: **500**.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListChatRecordDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatRecordDetailRequest) GoString() string {
	return s.String()
}

func (s *ListChatRecordDetailRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListChatRecordDetailRequest) GetCloseTimeEnd() *int64 {
	return s.CloseTimeEnd
}

func (s *ListChatRecordDetailRequest) GetCloseTimeStart() *int64 {
	return s.CloseTimeStart
}

func (s *ListChatRecordDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListChatRecordDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListChatRecordDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatRecordDetailRequest) SetClientToken(v string) *ListChatRecordDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetCloseTimeEnd(v int64) *ListChatRecordDetailRequest {
	s.CloseTimeEnd = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetCloseTimeStart(v int64) *ListChatRecordDetailRequest {
	s.CloseTimeStart = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetCurrentPage(v int32) *ListChatRecordDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetInstanceId(v string) *ListChatRecordDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChatRecordDetailRequest) SetPageSize(v int32) *ListChatRecordDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatRecordDetailRequest) Validate() error {
	return dara.Validate(s)
}
