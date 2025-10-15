// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryListByBizTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetHistoryListByBizTypeResponseBody
	GetCost() *int64
	SetData(v *GetHistoryListByBizTypeResponseBodyData) *GetHistoryListByBizTypeResponseBody
	GetData() *GetHistoryListByBizTypeResponseBodyData
	SetDataType(v string) *GetHistoryListByBizTypeResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetHistoryListByBizTypeResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetHistoryListByBizTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHistoryListByBizTypeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHistoryListByBizTypeResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetHistoryListByBizTypeResponseBody
	GetTime() *string
}

type GetHistoryListByBizTypeResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                   `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetHistoryListByBizTypeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 9DF9B3F3-9FFE-52CB-A8DC-F7BD5F842F0E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetHistoryListByBizTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetHistoryListByBizTypeResponseBody) GetData() *GetHistoryListByBizTypeResponseBodyData {
	return s.Data
}

func (s *GetHistoryListByBizTypeResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetHistoryListByBizTypeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetHistoryListByBizTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHistoryListByBizTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHistoryListByBizTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHistoryListByBizTypeResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetHistoryListByBizTypeResponseBody) SetCost(v int64) *GetHistoryListByBizTypeResponseBody {
	s.Cost = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetData(v *GetHistoryListByBizTypeResponseBodyData) *GetHistoryListByBizTypeResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetDataType(v string) *GetHistoryListByBizTypeResponseBody {
	s.DataType = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetErrCode(v string) *GetHistoryListByBizTypeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetMessage(v string) *GetHistoryListByBizTypeResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetRequestId(v string) *GetHistoryListByBizTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetSuccess(v bool) *GetHistoryListByBizTypeResponseBody {
	s.Success = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetTime(v string) *GetHistoryListByBizTypeResponseBody {
	s.Time = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHistoryListByBizTypeResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetHistoryListByBizTypeResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetHistoryListByBizTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetHistoryListByBizTypeResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetHistoryListByBizTypeResponseBodyData) GetRecords() []*GetHistoryListByBizTypeResponseBodyDataRecords {
	return s.Records
}

func (s *GetHistoryListByBizTypeResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *GetHistoryListByBizTypeResponseBodyData) GetTotalRecords() *int64 {
	return s.TotalRecords
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetCurrentPage(v int64) *GetHistoryListByBizTypeResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetPageSize(v int64) *GetHistoryListByBizTypeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetRecords(v []*GetHistoryListByBizTypeResponseBodyDataRecords) *GetHistoryListByBizTypeResponseBodyData {
	s.Records = v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetTotalPages(v int64) *GetHistoryListByBizTypeResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetTotalRecords(v int64) *GetHistoryListByBizTypeResponseBodyData {
	s.TotalRecords = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHistoryListByBizTypeResponseBodyDataRecords struct {
	// example:
	//
	// GysYBsxx
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// example:
	//
	// LibraryChat
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// null
	ExtraMessage interface{} `json:"extraMessage,omitempty" xml:"extraMessage,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 210
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LlmAnswer *string `json:"llmAnswer,omitempty" xml:"llmAnswer,omitempty"`
	LlmPrompt *string `json:"llmPrompt,omitempty" xml:"llmPrompt,omitempty"`
	// example:
	//
	// qwen-max
	LlmType *string `json:"llmType,omitempty" xml:"llmType,omitempty"`
	// example:
	//
	// null
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	UserQuery *string `json:"userQuery,omitempty" xml:"userQuery,omitempty"`
}

func (s GetHistoryListByBizTypeResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetBizId() *string {
	return s.BizId
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetBizType() *string {
	return s.BizType
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetExtraMessage() interface{} {
	return s.ExtraMessage
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetId() *int64 {
	return s.Id
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetLlmAnswer() *string {
	return s.LlmAnswer
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetLlmPrompt() *string {
	return s.LlmPrompt
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetLlmType() *string {
	return s.LlmType
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetSessionId() *string {
	return s.SessionId
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) GetUserQuery() *string {
	return s.UserQuery
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetBizId(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.BizId = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetBizType(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.BizType = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetExtraMessage(v interface{}) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.ExtraMessage = v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetGmtCreate(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetGmtModified(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetId(v int64) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetLlmAnswer(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.LlmAnswer = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetLlmPrompt(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.LlmPrompt = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetLlmType(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.LlmType = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetSessionId(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.SessionId = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetUserQuery(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.UserQuery = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
