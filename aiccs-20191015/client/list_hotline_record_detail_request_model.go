// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotlineRecordDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListHotlineRecordDetailRequest
	GetClientToken() *string
	SetCloseTimeEnd(v int64) *ListHotlineRecordDetailRequest
	GetCloseTimeEnd() *int64
	SetCloseTimeStart(v int64) *ListHotlineRecordDetailRequest
	GetCloseTimeStart() *int64
	SetCurrentPage(v int32) *ListHotlineRecordDetailRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListHotlineRecordDetailRequest
	GetInstanceId() *string
	SetPageSize(v int32) *ListHotlineRecordDetailRequest
	GetPageSize() *int32
}

type ListHotlineRecordDetailRequest struct {
	// A unique ID for the customer request. Used for idempotency validation and can be generated using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The start time when the hotline call ends. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1614582000000
	CloseTimeEnd *int64 `json:"CloseTimeEnd,omitempty" xml:"CloseTimeEnd,omitempty"`
	// The end time when the hotline call ends. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1614578400000
	CloseTimeStart *int64 `json:"CloseTimeStart,omitempty" xml:"CloseTimeStart,omitempty"`
	// Current page number. Default value: **1**.
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
	// Page size. Default value: **100**.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotlineRecordDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotlineRecordDetailRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordDetailRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListHotlineRecordDetailRequest) GetCloseTimeEnd() *int64 {
	return s.CloseTimeEnd
}

func (s *ListHotlineRecordDetailRequest) GetCloseTimeStart() *int64 {
	return s.CloseTimeStart
}

func (s *ListHotlineRecordDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHotlineRecordDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHotlineRecordDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotlineRecordDetailRequest) SetClientToken(v string) *ListHotlineRecordDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetCloseTimeEnd(v int64) *ListHotlineRecordDetailRequest {
	s.CloseTimeEnd = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetCloseTimeStart(v int64) *ListHotlineRecordDetailRequest {
	s.CloseTimeStart = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetCurrentPage(v int32) *ListHotlineRecordDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetInstanceId(v string) *ListHotlineRecordDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) SetPageSize(v int32) *ListHotlineRecordDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListHotlineRecordDetailRequest) Validate() error {
	return dara.Validate(s)
}
