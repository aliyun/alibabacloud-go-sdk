// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataAgentMemoryResponseBody
	GetCode() *string
	SetData(v *ListDataAgentMemoryResponseBodyData) *ListDataAgentMemoryResponseBody
	GetData() *ListDataAgentMemoryResponseBodyData
	SetErrorCode(v string) *ListDataAgentMemoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentMemoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDataAgentMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataAgentMemoryResponseBody
	GetSuccess() *bool
	SetTimestamp(v string) *ListDataAgentMemoryResponseBody
	GetTimestamp() *string
}

type ListDataAgentMemoryResponseBody struct {
	// The status code. A value of Success indicates success.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response struct.
	Data *ListDataAgentMemoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The operation timestamp.
	//
	// example:
	//
	// 1768270172
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListDataAgentMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataAgentMemoryResponseBody) GetData() *ListDataAgentMemoryResponseBodyData {
	return s.Data
}

func (s *ListDataAgentMemoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentMemoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataAgentMemoryResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListDataAgentMemoryResponseBody) SetCode(v string) *ListDataAgentMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataAgentMemoryResponseBody) SetData(v *ListDataAgentMemoryResponseBodyData) *ListDataAgentMemoryResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentMemoryResponseBody) SetErrorCode(v string) *ListDataAgentMemoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentMemoryResponseBody) SetErrorMessage(v string) *ListDataAgentMemoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentMemoryResponseBody) SetRequestId(v string) *ListDataAgentMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentMemoryResponseBody) SetSuccess(v bool) *ListDataAgentMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentMemoryResponseBody) SetTimestamp(v string) *ListDataAgentMemoryResponseBody {
	s.Timestamp = &v
	return s
}

func (s *ListDataAgentMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataAgentMemoryResponseBodyData struct {
	// The response struct.
	Data []*ListDataAgentMemoryResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDataAgentMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentMemoryResponseBodyData) GetData() []*ListDataAgentMemoryResponseBodyDataData {
	return s.Data
}

func (s *ListDataAgentMemoryResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListDataAgentMemoryResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataAgentMemoryResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListDataAgentMemoryResponseBodyData) SetData(v []*ListDataAgentMemoryResponseBodyDataData) *ListDataAgentMemoryResponseBodyData {
	s.Data = v
	return s
}

func (s *ListDataAgentMemoryResponseBodyData) SetPageNum(v int64) *ListDataAgentMemoryResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyData) SetPageSize(v int64) *ListDataAgentMemoryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyData) SetTotal(v int64) *ListDataAgentMemoryResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataAgentMemoryResponseBodyDataData struct {
	// The memory content.
	//
	// example:
	//
	// Diamond pricing analysis requires examining the skewness and outliers of the distribution of each feature.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The memory source ID.
	//
	// example:
	//
	// w3xa1********x6y8zm
	FromId *string `json:"FromId,omitempty" xml:"FromId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-07-29T07:11:23Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-07-29T07:11:23Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The memory hit level (hotness).
	//
	// example:
	//
	// 1
	HintLevel *int64 `json:"HintLevel,omitempty" xml:"HintLevel,omitempty"`
	// The memory source.
	//
	// example:
	//
	// session
	MemFrom *string `json:"MemFrom,omitempty" xml:"MemFrom,omitempty"`
	// The memory status.
	//
	// example:
	//
	// memorized
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The memory UUID.
	//
	// example:
	//
	// 8zm3w********g3yxa1
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListDataAgentMemoryResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentMemoryResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListDataAgentMemoryResponseBodyDataData) GetContent() *string {
	return s.Content
}

func (s *ListDataAgentMemoryResponseBodyDataData) GetFromId() *string {
	return s.FromId
}

func (s *ListDataAgentMemoryResponseBodyDataData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListDataAgentMemoryResponseBodyDataData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataAgentMemoryResponseBodyDataData) GetHintLevel() *int64 {
	return s.HintLevel
}

func (s *ListDataAgentMemoryResponseBodyDataData) GetMemFrom() *string {
	return s.MemFrom
}

func (s *ListDataAgentMemoryResponseBodyDataData) GetStatus() *string {
	return s.Status
}

func (s *ListDataAgentMemoryResponseBodyDataData) GetUuid() *string {
	return s.Uuid
}

func (s *ListDataAgentMemoryResponseBodyDataData) SetContent(v string) *ListDataAgentMemoryResponseBodyDataData {
	s.Content = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyDataData) SetFromId(v string) *ListDataAgentMemoryResponseBodyDataData {
	s.FromId = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyDataData) SetGmtCreated(v string) *ListDataAgentMemoryResponseBodyDataData {
	s.GmtCreated = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyDataData) SetGmtModified(v string) *ListDataAgentMemoryResponseBodyDataData {
	s.GmtModified = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyDataData) SetHintLevel(v int64) *ListDataAgentMemoryResponseBodyDataData {
	s.HintLevel = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyDataData) SetMemFrom(v string) *ListDataAgentMemoryResponseBodyDataData {
	s.MemFrom = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyDataData) SetStatus(v string) *ListDataAgentMemoryResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyDataData) SetUuid(v string) *ListDataAgentMemoryResponseBodyDataData {
	s.Uuid = &v
	return s
}

func (s *ListDataAgentMemoryResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}
