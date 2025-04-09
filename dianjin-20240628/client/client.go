// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type CreateAnnualDocSummaryTaskRequest struct {
	// This parameter is required.
	AnaYears []*int32 `json:"anaYears,omitempty" xml:"anaYears,omitempty" type:"Repeated"`
	// This parameter is required.
	DocInfos []*CreateAnnualDocSummaryTaskRequestDocInfos `json:"docInfos,omitempty" xml:"docInfos,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableTable *bool   `json:"enableTable,omitempty" xml:"enableTable,omitempty"`
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s CreateAnnualDocSummaryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAnnualDocSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAnnualDocSummaryTaskRequest) SetAnaYears(v []*int32) *CreateAnnualDocSummaryTaskRequest {
	s.AnaYears = v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) SetDocInfos(v []*CreateAnnualDocSummaryTaskRequestDocInfos) *CreateAnnualDocSummaryTaskRequest {
	s.DocInfos = v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) SetEnableTable(v bool) *CreateAnnualDocSummaryTaskRequest {
	s.EnableTable = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) SetInstruction(v string) *CreateAnnualDocSummaryTaskRequest {
	s.Instruction = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequest) SetModelId(v string) *CreateAnnualDocSummaryTaskRequest {
	s.ModelId = &v
	return s
}

type CreateAnnualDocSummaryTaskRequestDocInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 198386463432
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023
	DocYear *int32 `json:"docYear,omitempty" xml:"docYear,omitempty"`
	// example:
	//
	// 2
	EndPage *int32 `json:"endPage,omitempty" xml:"endPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rdxrmo6amk
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// 1
	StartPage *int32 `json:"startPage,omitempty" xml:"startPage,omitempty"`
}

func (s CreateAnnualDocSummaryTaskRequestDocInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateAnnualDocSummaryTaskRequestDocInfos) GoString() string {
	return s.String()
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetDocId(v string) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.DocId = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetDocYear(v int32) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.DocYear = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetEndPage(v int32) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.EndPage = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetLibraryId(v string) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.LibraryId = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskRequestDocInfos) SetStartPage(v int32) *CreateAnnualDocSummaryTaskRequestDocInfos {
	s.StartPage = &v
	return s
}

type CreateAnnualDocSummaryTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 3284627354
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// 32FFC91D-0A9F-585A-B84F-8A54C5187035
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

func (s CreateAnnualDocSummaryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAnnualDocSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetCost(v int64) *CreateAnnualDocSummaryTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetData(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetDataType(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetErrCode(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetMessage(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetRequestId(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetSuccess(v bool) *CreateAnnualDocSummaryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponseBody) SetTime(v string) *CreateAnnualDocSummaryTaskResponseBody {
	s.Time = &v
	return s
}

type CreateAnnualDocSummaryTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAnnualDocSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAnnualDocSummaryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAnnualDocSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAnnualDocSummaryTaskResponse) SetHeaders(v map[string]*string) *CreateAnnualDocSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponse) SetStatusCode(v int32) *CreateAnnualDocSummaryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAnnualDocSummaryTaskResponse) SetBody(v *CreateAnnualDocSummaryTaskResponseBody) *CreateAnnualDocSummaryTaskResponse {
	s.Body = v
	return s
}

type CreateDialogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// taobao
	Channel       *string `json:"channel,omitempty" xml:"channel,omitempty"`
	EnableLibrary *bool   `json:"enableLibrary,omitempty" xml:"enableLibrary,omitempty"`
	// example:
	//
	// null
	MetaData map[string]interface{} `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// live_broadcast_qa
	PlayCode      *string   `json:"playCode,omitempty" xml:"playCode,omitempty"`
	QaLibraryList []*string `json:"qaLibraryList,omitempty" xml:"qaLibraryList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ebf83826-dc1c-46f8-9759-0fb6da4c8xxx
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SelfDirected *bool   `json:"selfDirected,omitempty" xml:"selfDirected,omitempty"`
}

func (s CreateDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogRequest) GoString() string {
	return s.String()
}

func (s *CreateDialogRequest) SetChannel(v string) *CreateDialogRequest {
	s.Channel = &v
	return s
}

func (s *CreateDialogRequest) SetEnableLibrary(v bool) *CreateDialogRequest {
	s.EnableLibrary = &v
	return s
}

func (s *CreateDialogRequest) SetMetaData(v map[string]interface{}) *CreateDialogRequest {
	s.MetaData = v
	return s
}

func (s *CreateDialogRequest) SetPlayCode(v string) *CreateDialogRequest {
	s.PlayCode = &v
	return s
}

func (s *CreateDialogRequest) SetQaLibraryList(v []*string) *CreateDialogRequest {
	s.QaLibraryList = v
	return s
}

func (s *CreateDialogRequest) SetRequestId(v string) *CreateDialogRequest {
	s.RequestId = &v
	return s
}

func (s *CreateDialogRequest) SetSelfDirected(v bool) *CreateDialogRequest {
	s.SelfDirected = &v
	return s
}

type CreateDialogResponseBody struct {
	// example:
	//
	// null
	Cost *int64                        `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *CreateDialogResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 003D019A-1BB3-53EC-A0D2-CE76DA5D73B1
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

func (s CreateDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDialogResponseBody) SetCost(v int64) *CreateDialogResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateDialogResponseBody) SetData(v *CreateDialogResponseBodyData) *CreateDialogResponseBody {
	s.Data = v
	return s
}

func (s *CreateDialogResponseBody) SetDataType(v string) *CreateDialogResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateDialogResponseBody) SetErrCode(v string) *CreateDialogResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDialogResponseBody) SetMessage(v string) *CreateDialogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDialogResponseBody) SetRequestId(v string) *CreateDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDialogResponseBody) SetSuccess(v bool) *CreateDialogResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDialogResponseBody) SetTime(v string) *CreateDialogResponseBody {
	s.Time = &v
	return s
}

type CreateDialogResponseBodyData struct {
	OpeningRemarks *string `json:"openingRemarks,omitempty" xml:"openingRemarks,omitempty"`
	// example:
	//
	// 1728545917713234
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CreateDialogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDialogResponseBodyData) SetOpeningRemarks(v string) *CreateDialogResponseBodyData {
	s.OpeningRemarks = &v
	return s
}

func (s *CreateDialogResponseBodyData) SetSessionId(v string) *CreateDialogResponseBodyData {
	s.SessionId = &v
	return s
}

type CreateDialogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogResponse) GoString() string {
	return s.String()
}

func (s *CreateDialogResponse) SetHeaders(v map[string]*string) *CreateDialogResponse {
	s.Headers = v
	return s
}

func (s *CreateDialogResponse) SetStatusCode(v int32) *CreateDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDialogResponse) SetBody(v *CreateDialogResponseBody) *CreateDialogResponse {
	s.Body = v
	return s
}

type CreateDialogAnalysisTaskRequest struct {
	AnalysisNodes []*string `json:"analysisNodes,omitempty" xml:"analysisNodes,omitempty" type:"Repeated"`
	// This parameter is required.
	ConversationList []*CreateDialogAnalysisTaskRequestConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	// "labels": "XXX",
	//
	// "summaryConstraints": "XXX",
	//
	// "sopInfo": "XXX"
	//
	// }
	MetaData map[string]interface{} `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// common
	PlayCode *string `json:"playCode,omitempty" xml:"playCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDialogAnalysisTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogAnalysisTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskRequest) SetAnalysisNodes(v []*string) *CreateDialogAnalysisTaskRequest {
	s.AnalysisNodes = v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) SetConversationList(v []*CreateDialogAnalysisTaskRequestConversationList) *CreateDialogAnalysisTaskRequest {
	s.ConversationList = v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) SetMetaData(v map[string]interface{}) *CreateDialogAnalysisTaskRequest {
	s.MetaData = v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) SetPlayCode(v string) *CreateDialogAnalysisTaskRequest {
	s.PlayCode = &v
	return s
}

func (s *CreateDialogAnalysisTaskRequest) SetRequestId(v string) *CreateDialogAnalysisTaskRequest {
	s.RequestId = &v
	return s
}

type CreateDialogAnalysisTaskRequestConversationList struct {
	// This parameter is required.
	DialogueList []*CreateDialogAnalysisTaskRequestConversationListDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
}

func (s CreateDialogAnalysisTaskRequestConversationList) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogAnalysisTaskRequestConversationList) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskRequestConversationList) SetDialogueList(v []*CreateDialogAnalysisTaskRequestConversationListDialogueList) *CreateDialogAnalysisTaskRequestConversationList {
	s.DialogueList = v
	return s
}

type CreateDialogAnalysisTaskRequestConversationListDialogueList struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s CreateDialogAnalysisTaskRequestConversationListDialogueList) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogAnalysisTaskRequestConversationListDialogueList) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskRequestConversationListDialogueList) SetContent(v string) *CreateDialogAnalysisTaskRequestConversationListDialogueList {
	s.Content = &v
	return s
}

func (s *CreateDialogAnalysisTaskRequestConversationListDialogueList) SetRole(v string) *CreateDialogAnalysisTaskRequestConversationListDialogueList {
	s.Role = &v
	return s
}

type CreateDialogAnalysisTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64    `json:"cost,omitempty" xml:"cost,omitempty"`
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateDialogAnalysisTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskResponseBody) SetCost(v int64) *CreateDialogAnalysisTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetData(v []*string) *CreateDialogAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetDataType(v string) *CreateDialogAnalysisTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetErrCode(v string) *CreateDialogAnalysisTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetMessage(v string) *CreateDialogAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetRequestId(v string) *CreateDialogAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetSuccess(v bool) *CreateDialogAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponseBody) SetTime(v string) *CreateDialogAnalysisTaskResponseBody {
	s.Time = &v
	return s
}

type CreateDialogAnalysisTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDialogAnalysisTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDialogAnalysisTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDialogAnalysisTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDialogAnalysisTaskResponse) SetHeaders(v map[string]*string) *CreateDialogAnalysisTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDialogAnalysisTaskResponse) SetStatusCode(v int32) *CreateDialogAnalysisTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDialogAnalysisTaskResponse) SetBody(v *CreateDialogAnalysisTaskResponseBody) *CreateDialogAnalysisTaskResponse {
	s.Body = v
	return s
}

type CreateDocsSummaryTaskRequest struct {
	// This parameter is required.
	DocInfos []*CreateDocsSummaryTaskRequestDocInfos `json:"docInfos,omitempty" xml:"docInfos,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableTable *bool   `json:"enableTable,omitempty" xml:"enableTable,omitempty"`
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s CreateDocsSummaryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDocsSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDocsSummaryTaskRequest) SetDocInfos(v []*CreateDocsSummaryTaskRequestDocInfos) *CreateDocsSummaryTaskRequest {
	s.DocInfos = v
	return s
}

func (s *CreateDocsSummaryTaskRequest) SetEnableTable(v bool) *CreateDocsSummaryTaskRequest {
	s.EnableTable = &v
	return s
}

func (s *CreateDocsSummaryTaskRequest) SetInstruction(v string) *CreateDocsSummaryTaskRequest {
	s.Instruction = &v
	return s
}

func (s *CreateDocsSummaryTaskRequest) SetModelId(v string) *CreateDocsSummaryTaskRequest {
	s.ModelId = &v
	return s
}

type CreateDocsSummaryTaskRequestDocInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 198386463432
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// 2
	EndPage *int32 `json:"endPage,omitempty" xml:"endPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rdxrmo6amk
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// 1
	StartPage *int32 `json:"startPage,omitempty" xml:"startPage,omitempty"`
}

func (s CreateDocsSummaryTaskRequestDocInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateDocsSummaryTaskRequestDocInfos) GoString() string {
	return s.String()
}

func (s *CreateDocsSummaryTaskRequestDocInfos) SetDocId(v string) *CreateDocsSummaryTaskRequestDocInfos {
	s.DocId = &v
	return s
}

func (s *CreateDocsSummaryTaskRequestDocInfos) SetEndPage(v int32) *CreateDocsSummaryTaskRequestDocInfos {
	s.EndPage = &v
	return s
}

func (s *CreateDocsSummaryTaskRequestDocInfos) SetLibraryId(v string) *CreateDocsSummaryTaskRequestDocInfos {
	s.LibraryId = &v
	return s
}

func (s *CreateDocsSummaryTaskRequestDocInfos) SetStartPage(v int32) *CreateDocsSummaryTaskRequestDocInfos {
	s.StartPage = &v
	return s
}

type CreateDocsSummaryTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 765675376
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// 32FFC91D-0A9F-585A-B84F-8A54C5187035
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

func (s CreateDocsSummaryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDocsSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocsSummaryTaskResponseBody) SetCost(v int64) *CreateDocsSummaryTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetData(v string) *CreateDocsSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetDataType(v string) *CreateDocsSummaryTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetErrCode(v string) *CreateDocsSummaryTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetMessage(v string) *CreateDocsSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetRequestId(v string) *CreateDocsSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetSuccess(v bool) *CreateDocsSummaryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDocsSummaryTaskResponseBody) SetTime(v string) *CreateDocsSummaryTaskResponseBody {
	s.Time = &v
	return s
}

type CreateDocsSummaryTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocsSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocsSummaryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDocsSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDocsSummaryTaskResponse) SetHeaders(v map[string]*string) *CreateDocsSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDocsSummaryTaskResponse) SetStatusCode(v int32) *CreateDocsSummaryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocsSummaryTaskResponse) SetBody(v *CreateDocsSummaryTaskResponseBody) *CreateDocsSummaryTaskResponse {
	s.Body = v
	return s
}

type CreateFinReportSummaryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableTable *bool `json:"enableTable,omitempty" xml:"enableTable,omitempty"`
	// example:
	//
	// 10
	EndPage     *int32  `json:"endPage,omitempty" xml:"endPage,omitempty"`
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 1
	StartPage *int32 `json:"startPage,omitempty" xml:"startPage,omitempty"`
	// example:
	//
	// custom
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateFinReportSummaryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFinReportSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskRequest) SetDocId(v string) *CreateFinReportSummaryTaskRequest {
	s.DocId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetEnableTable(v bool) *CreateFinReportSummaryTaskRequest {
	s.EnableTable = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetEndPage(v int32) *CreateFinReportSummaryTaskRequest {
	s.EndPage = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetInstruction(v string) *CreateFinReportSummaryTaskRequest {
	s.Instruction = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetLibraryId(v string) *CreateFinReportSummaryTaskRequest {
	s.LibraryId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetModelId(v string) *CreateFinReportSummaryTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetStartPage(v int32) *CreateFinReportSummaryTaskRequest {
	s.StartPage = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetTaskType(v string) *CreateFinReportSummaryTaskRequest {
	s.TaskType = &v
	return s
}

type CreateFinReportSummaryTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 3284627354
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateFinReportSummaryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFinReportSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskResponseBody) SetCost(v int64) *CreateFinReportSummaryTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetData(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetDataType(v string) *CreateFinReportSummaryTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetErrCode(v string) *CreateFinReportSummaryTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetMessage(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetRequestId(v string) *CreateFinReportSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetSuccess(v bool) *CreateFinReportSummaryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetTime(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Time = &v
	return s
}

type CreateFinReportSummaryTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFinReportSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFinReportSummaryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFinReportSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskResponse) SetHeaders(v map[string]*string) *CreateFinReportSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFinReportSummaryTaskResponse) SetStatusCode(v int32) *CreateFinReportSummaryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponse) SetBody(v *CreateFinReportSummaryTaskResponseBody) *CreateFinReportSummaryTaskResponse {
	s.Body = v
	return s
}

type CreateLibraryRequest struct {
	// This parameter is required.
	Description  *string                           `json:"description,omitempty" xml:"description,omitempty"`
	IndexSetting *CreateLibraryRequestIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	// This parameter is required.
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s CreateLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequest) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequest) SetDescription(v string) *CreateLibraryRequest {
	s.Description = &v
	return s
}

func (s *CreateLibraryRequest) SetIndexSetting(v *CreateLibraryRequestIndexSetting) *CreateLibraryRequest {
	s.IndexSetting = v
	return s
}

func (s *CreateLibraryRequest) SetLibraryName(v string) *CreateLibraryRequest {
	s.LibraryName = &v
	return s
}

type CreateLibraryRequestIndexSetting struct {
	ChunkStrategy      *CreateLibraryRequestIndexSettingChunkStrategy      `json:"chunkStrategy,omitempty" xml:"chunkStrategy,omitempty" type:"Struct"`
	ModelConfig        *CreateLibraryRequestIndexSettingModelConfig        `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	PromptRoleStyle    *string                                             `json:"promptRoleStyle,omitempty" xml:"promptRoleStyle,omitempty"`
	QueryEnhancer      *CreateLibraryRequestIndexSettingQueryEnhancer      `json:"queryEnhancer,omitempty" xml:"queryEnhancer,omitempty" type:"Struct"`
	RecallStrategy     *CreateLibraryRequestIndexSettingRecallStrategy     `json:"recallStrategy,omitempty" xml:"recallStrategy,omitempty" type:"Struct"`
	TextIndexSetting   *CreateLibraryRequestIndexSettingTextIndexSetting   `json:"textIndexSetting,omitempty" xml:"textIndexSetting,omitempty" type:"Struct"`
	VectorIndexSetting *CreateLibraryRequestIndexSettingVectorIndexSetting `json:"vectorIndexSetting,omitempty" xml:"vectorIndexSetting,omitempty" type:"Struct"`
}

func (s CreateLibraryRequestIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSetting) SetChunkStrategy(v *CreateLibraryRequestIndexSettingChunkStrategy) *CreateLibraryRequestIndexSetting {
	s.ChunkStrategy = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetModelConfig(v *CreateLibraryRequestIndexSettingModelConfig) *CreateLibraryRequestIndexSetting {
	s.ModelConfig = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetPromptRoleStyle(v string) *CreateLibraryRequestIndexSetting {
	s.PromptRoleStyle = &v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetQueryEnhancer(v *CreateLibraryRequestIndexSettingQueryEnhancer) *CreateLibraryRequestIndexSetting {
	s.QueryEnhancer = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetRecallStrategy(v *CreateLibraryRequestIndexSettingRecallStrategy) *CreateLibraryRequestIndexSetting {
	s.RecallStrategy = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetTextIndexSetting(v *CreateLibraryRequestIndexSettingTextIndexSetting) *CreateLibraryRequestIndexSetting {
	s.TextIndexSetting = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetVectorIndexSetting(v *CreateLibraryRequestIndexSettingVectorIndexSetting) *CreateLibraryRequestIndexSetting {
	s.VectorIndexSetting = v
	return s
}

type CreateLibraryRequestIndexSettingChunkStrategy struct {
	// example:
	//
	// true
	DocTreeSplit *bool `json:"docTreeSplit,omitempty" xml:"docTreeSplit,omitempty"`
	// example:
	//
	// 300
	DocTreeSplitSize *int32 `json:"docTreeSplitSize,omitempty" xml:"docTreeSplitSize,omitempty"`
	// example:
	//
	// true
	EnhanceGraph *bool `json:"enhanceGraph,omitempty" xml:"enhanceGraph,omitempty"`
	// example:
	//
	// true
	EnhanceTable *bool `json:"enhanceTable,omitempty" xml:"enhanceTable,omitempty"`
	// example:
	//
	// 20
	Overlap *int32 `json:"overlap,omitempty" xml:"overlap,omitempty"`
	// example:
	//
	// true
	SentenceSplit *bool `json:"sentenceSplit,omitempty" xml:"sentenceSplit,omitempty"`
	// example:
	//
	// 300
	SentenceSplitSize *int32 `json:"sentenceSplitSize,omitempty" xml:"sentenceSplitSize,omitempty"`
	// example:
	//
	// 300
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// true
	Split *bool `json:"split,omitempty" xml:"split,omitempty"`
}

func (s CreateLibraryRequestIndexSettingChunkStrategy) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetDocTreeSplit(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.DocTreeSplit = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetDocTreeSplitSize(v int32) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.DocTreeSplitSize = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetEnhanceGraph(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.EnhanceGraph = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetEnhanceTable(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.EnhanceTable = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetOverlap(v int32) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.Overlap = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetSentenceSplit(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.SentenceSplit = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetSentenceSplitSize(v int32) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.SentenceSplitSize = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetSize(v int32) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.Size = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetSplit(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.Split = &v
	return s
}

type CreateLibraryRequestIndexSettingModelConfig struct {
	// example:
	//
	// 0.8
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// example:
	//
	// 0.8
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s CreateLibraryRequestIndexSettingModelConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingModelConfig) SetTemperature(v float64) *CreateLibraryRequestIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingModelConfig) SetTopP(v float64) *CreateLibraryRequestIndexSettingModelConfig {
	s.TopP = &v
	return s
}

type CreateLibraryRequestIndexSettingQueryEnhancer struct {
	// example:
	//
	// true
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// true
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// true
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// true
	EnableQueryRewrite *bool `json:"enableQueryRewrite,omitempty" xml:"enableQueryRewrite,omitempty"`
	// example:
	//
	// true
	EnableSession *bool `json:"enableSession,omitempty" xml:"enableSession,omitempty"`
	// example:
	//
	// xxxx
	LocalKnowledgeId *string `json:"localKnowledgeId,omitempty" xml:"localKnowledgeId,omitempty"`
	// example:
	//
	// true
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s CreateLibraryRequestIndexSettingQueryEnhancer) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableFollowUp(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableFollowUp = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableMultiQuery(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableMultiQuery = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableOpenQa(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableOpenQa = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableQueryRewrite(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableQueryRewrite = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableSession(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableSession = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetLocalKnowledgeId(v string) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.LocalKnowledgeId = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetWithDocumentReference(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.WithDocumentReference = &v
	return s
}

type CreateLibraryRequestIndexSettingRecallStrategy struct {
	// example:
	//
	// model
	DocumentRankType *string `json:"documentRankType,omitempty" xml:"documentRankType,omitempty"`
	// example:
	//
	// 20
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s CreateLibraryRequestIndexSettingRecallStrategy) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) SetDocumentRankType(v string) *CreateLibraryRequestIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) SetLimit(v int32) *CreateLibraryRequestIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

type CreateLibraryRequestIndexSettingTextIndexSetting struct {
	// example:
	//
	// ElasticSearch
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// Standard
	IndexAnalyzer *string `json:"indexAnalyzer,omitempty" xml:"indexAnalyzer,omitempty"`
	// example:
	//
	// 0.5
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// Standard
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" xml:"searchAnalyzer,omitempty"`
	// example:
	//
	// 50
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s CreateLibraryRequestIndexSettingTextIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetCategory(v string) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.Category = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetEnable(v bool) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.Enable = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetIndexAnalyzer(v string) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.IndexAnalyzer = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetRankThreshold(v float64) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetSearchAnalyzer(v string) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.SearchAnalyzer = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetTopK(v int32) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.TopK = &v
	return s
}

type CreateLibraryRequestIndexSettingVectorIndexSetting struct {
	// example:
	//
	// ADB
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// DashScope
	EmbeddingType *string `json:"embeddingType,omitempty" xml:"embeddingType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 0.5
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// 50
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s CreateLibraryRequestIndexSettingVectorIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetCategory(v string) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.Category = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetEmbeddingType(v string) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.EmbeddingType = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetEnable(v bool) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.Enable = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetRankThreshold(v float64) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetTopK(v int32) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.TopK = &v
	return s
}

type CreateLibraryResponseBody struct {
	// example:
	//
	// 300
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// a1b2c3
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// null
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLibraryResponseBody) SetCost(v int64) *CreateLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateLibraryResponseBody) SetData(v string) *CreateLibraryResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLibraryResponseBody) SetDataType(v string) *CreateLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateLibraryResponseBody) SetErrCode(v string) *CreateLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateLibraryResponseBody) SetMessage(v string) *CreateLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLibraryResponseBody) SetRequestId(v string) *CreateLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLibraryResponseBody) SetSuccess(v bool) *CreateLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLibraryResponseBody) SetTime(v string) *CreateLibraryResponseBody {
	s.Time = &v
	return s
}

type CreateLibraryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryResponse) GoString() string {
	return s.String()
}

func (s *CreateLibraryResponse) SetHeaders(v map[string]*string) *CreateLibraryResponse {
	s.Headers = v
	return s
}

func (s *CreateLibraryResponse) SetStatusCode(v int32) *CreateLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLibraryResponse) SetBody(v *CreateLibraryResponseBody) *CreateLibraryResponse {
	s.Body = v
	return s
}

type CreatePdfTranslateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 873648346573245
	DocId     *string `json:"docId,omitempty" xml:"docId,omitempty"`
	Knowledge *string `json:"knowledge,omitempty" xml:"knowledge,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 中文
	TranslateTo *string `json:"translateTo,omitempty" xml:"translateTo,omitempty"`
}

func (s CreatePdfTranslateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePdfTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePdfTranslateTaskRequest) SetDocId(v string) *CreatePdfTranslateTaskRequest {
	s.DocId = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) SetKnowledge(v string) *CreatePdfTranslateTaskRequest {
	s.Knowledge = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) SetLibraryId(v string) *CreatePdfTranslateTaskRequest {
	s.LibraryId = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) SetModelId(v string) *CreatePdfTranslateTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) SetTranslateTo(v string) *CreatePdfTranslateTaskRequest {
	s.TranslateTo = &v
	return s
}

type CreatePdfTranslateTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 3284627354
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreatePdfTranslateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePdfTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePdfTranslateTaskResponseBody) SetCost(v int64) *CreatePdfTranslateTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetData(v string) *CreatePdfTranslateTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetDataType(v string) *CreatePdfTranslateTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetErrCode(v string) *CreatePdfTranslateTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetMessage(v string) *CreatePdfTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetRequestId(v string) *CreatePdfTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetSuccess(v bool) *CreatePdfTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePdfTranslateTaskResponseBody) SetTime(v string) *CreatePdfTranslateTaskResponseBody {
	s.Time = &v
	return s
}

type CreatePdfTranslateTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePdfTranslateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdfTranslateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePdfTranslateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreatePdfTranslateTaskResponse) SetHeaders(v map[string]*string) *CreatePdfTranslateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreatePdfTranslateTaskResponse) SetStatusCode(v int32) *CreatePdfTranslateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdfTranslateTaskResponse) SetBody(v *CreatePdfTranslateTaskResponseBody) *CreatePdfTranslateTaskResponse {
	s.Body = v
	return s
}

type CreatePredefinedDocumentRequest struct {
	Chunks []*CreatePredefinedDocumentRequestChunks `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	// example:
	//
	// a1b2c3
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// {"a": "1"}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// 测试文档
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreatePredefinedDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePredefinedDocumentRequest) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentRequest) SetChunks(v []*CreatePredefinedDocumentRequestChunks) *CreatePredefinedDocumentRequest {
	s.Chunks = v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetLibraryId(v string) *CreatePredefinedDocumentRequest {
	s.LibraryId = &v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetMetadata(v map[string]interface{}) *CreatePredefinedDocumentRequest {
	s.Metadata = v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetTitle(v string) *CreatePredefinedDocumentRequest {
	s.Title = &v
	return s
}

type CreatePredefinedDocumentRequestChunks struct {
	// example:
	//
	// {"a": "1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// 1
	ChunkOrder *int32 `json:"chunkOrder,omitempty" xml:"chunkOrder,omitempty"`
	// example:
	//
	// 这是一段测试文本
	ChunkText *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
}

func (s CreatePredefinedDocumentRequestChunks) String() string {
	return tea.Prettify(s)
}

func (s CreatePredefinedDocumentRequestChunks) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkMeta(v map[string]interface{}) *CreatePredefinedDocumentRequestChunks {
	s.ChunkMeta = v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkOrder(v int32) *CreatePredefinedDocumentRequestChunks {
	s.ChunkOrder = &v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkText(v string) *CreatePredefinedDocumentRequestChunks {
	s.ChunkText = &v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkType(v string) *CreatePredefinedDocumentRequestChunks {
	s.ChunkType = &v
	return s
}

type CreatePredefinedDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 1782981430906818562
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// 0a06dfe617018288881568684e2937
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreatePredefinedDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePredefinedDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentResponseBody) SetCost(v int64) *CreatePredefinedDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetData(v string) *CreatePredefinedDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetDataType(v string) *CreatePredefinedDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetErrCode(v string) *CreatePredefinedDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetMessage(v string) *CreatePredefinedDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetRequestId(v string) *CreatePredefinedDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetSuccess(v bool) *CreatePredefinedDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetTime(v string) *CreatePredefinedDocumentResponseBody {
	s.Time = &v
	return s
}

type CreatePredefinedDocumentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePredefinedDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePredefinedDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePredefinedDocumentResponse) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentResponse) SetHeaders(v map[string]*string) *CreatePredefinedDocumentResponse {
	s.Headers = v
	return s
}

func (s *CreatePredefinedDocumentResponse) SetStatusCode(v int32) *CreatePredefinedDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePredefinedDocumentResponse) SetBody(v *CreatePredefinedDocumentResponseBody) *CreatePredefinedDocumentResponse {
	s.Body = v
	return s
}

type CreateQualityCheckTaskRequest struct {
	// This parameter is required.
	ConversationList *CreateQualityCheckTaskRequestConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-27 11:23:20
	GmtService   *string            `json:"gmtService,omitempty" xml:"gmtService,omitempty"`
	MetaData     map[string]*string `json:"metaData,omitempty" xml:"metaData,omitempty"`
	QualityGroup []*string          `json:"qualityGroup,omitempty" xml:"qualityGroup,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateQualityCheckTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskRequest) SetConversationList(v *CreateQualityCheckTaskRequestConversationList) *CreateQualityCheckTaskRequest {
	s.ConversationList = v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetGmtService(v string) *CreateQualityCheckTaskRequest {
	s.GmtService = &v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetMetaData(v map[string]*string) *CreateQualityCheckTaskRequest {
	s.MetaData = v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetQualityGroup(v []*string) *CreateQualityCheckTaskRequest {
	s.QualityGroup = v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetRequestId(v string) *CreateQualityCheckTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateQualityCheckTaskRequest) SetType(v string) *CreateQualityCheckTaskRequest {
	s.Type = &v
	return s
}

type CreateQualityCheckTaskRequestConversationList struct {
	// example:
	//
	// 1
	CallType *string `json:"callType,omitempty" xml:"callType,omitempty"`
	// example:
	//
	// 1
	CustomerId   *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	CustomerName *string `json:"customerName,omitempty" xml:"customerName,omitempty"`
	// example:
	//
	// xxx
	CustomerServiceId   *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	CustomerServiceName *string `json:"customerServiceName,omitempty" xml:"customerServiceName,omitempty"`
	// This parameter is required.
	DialogueList []*CreateQualityCheckTaskRequestConversationListDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtService *string `json:"gmtService,omitempty" xml:"gmtService,omitempty"`
}

func (s CreateQualityCheckTaskRequestConversationList) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckTaskRequestConversationList) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCallType(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CallType = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCustomerId(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CustomerId = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCustomerName(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CustomerName = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCustomerServiceId(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CustomerServiceId = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetCustomerServiceName(v string) *CreateQualityCheckTaskRequestConversationList {
	s.CustomerServiceName = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetDialogueList(v []*CreateQualityCheckTaskRequestConversationListDialogueList) *CreateQualityCheckTaskRequestConversationList {
	s.DialogueList = v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationList) SetGmtService(v string) *CreateQualityCheckTaskRequestConversationList {
	s.GmtService = &v
	return s
}

type CreateQualityCheckTaskRequestConversationListDialogueList struct {
	// example:
	//
	// 0
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// example:
	//
	// 2024-05-23 14:57:50
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2348234
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// 23874627346
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// 0
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateQualityCheckTaskRequestConversationListDialogueList) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckTaskRequestConversationListDialogueList) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetBegin(v int32) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.Begin = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetBeginTime(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.BeginTime = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetContent(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.Content = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetCustomerId(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.CustomerId = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetCustomerServiceId(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.CustomerServiceId = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetCustomerServiceType(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.CustomerServiceType = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetEnd(v int32) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.End = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetRole(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.Role = &v
	return s
}

func (s *CreateQualityCheckTaskRequestConversationListDialogueList) SetType(v string) *CreateQualityCheckTaskRequestConversationListDialogueList {
	s.Type = &v
	return s
}

type CreateQualityCheckTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                  `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *CreateQualityCheckTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateQualityCheckTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskResponseBody) SetCost(v int64) *CreateQualityCheckTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetData(v *CreateQualityCheckTaskResponseBodyData) *CreateQualityCheckTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetDataType(v string) *CreateQualityCheckTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetErrCode(v string) *CreateQualityCheckTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetMessage(v string) *CreateQualityCheckTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetRequestId(v string) *CreateQualityCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetSuccess(v bool) *CreateQualityCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateQualityCheckTaskResponseBody) SetTime(v string) *CreateQualityCheckTaskResponseBody {
	s.Time = &v
	return s
}

type CreateQualityCheckTaskResponseBodyData struct {
	// taskId
	//
	// example:
	//
	// 172373500521
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateQualityCheckTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskResponseBodyData) SetTaskId(v string) *CreateQualityCheckTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type CreateQualityCheckTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQualityCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQualityCheckTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQualityCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateQualityCheckTaskResponse) SetHeaders(v map[string]*string) *CreateQualityCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateQualityCheckTaskResponse) SetStatusCode(v int32) *CreateQualityCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQualityCheckTaskResponse) SetBody(v *CreateQualityCheckTaskResponseBody) *CreateQualityCheckTaskResponse {
	s.Body = v
	return s
}

type DeleteDocumentRequest struct {
	// This parameter is required.
	DocIds []*string `json:"docIds,omitempty" xml:"docIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s DeleteDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentRequest) SetDocIds(v []*string) *DeleteDocumentRequest {
	s.DocIds = v
	return s
}

func (s *DeleteDocumentRequest) SetLibraryId(v string) *DeleteDocumentRequest {
	s.LibraryId = &v
	return s
}

type DeleteDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
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
	// 67C7021A-D268-553D-8C15-A087B9604028
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s DeleteDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponseBody) SetCost(v int64) *DeleteDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetData(v bool) *DeleteDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetDataType(v string) *DeleteDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetErrCode(v string) *DeleteDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetMessage(v string) *DeleteDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetRequestId(v string) *DeleteDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetSuccess(v bool) *DeleteDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetTime(v string) *DeleteDocumentResponseBody {
	s.Time = &v
	return s
}

type DeleteDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponse) SetHeaders(v map[string]*string) *DeleteDocumentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocumentResponse) SetStatusCode(v int32) *DeleteDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocumentResponse) SetBody(v *DeleteDocumentResponseBody) *DeleteDocumentResponse {
	s.Body = v
	return s
}

type DeleteLibraryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// skdfefxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s DeleteLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibraryRequest) GoString() string {
	return s.String()
}

func (s *DeleteLibraryRequest) SetLibraryId(v string) *DeleteLibraryRequest {
	s.LibraryId = &v
	return s
}

type DeleteLibraryResponseBody struct {
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
	// 30F6AD44-F078-540D-B5A5-1E519C8E9E6D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLibraryResponseBody) SetErrCode(v string) *DeleteLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetMessage(v string) *DeleteLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetRequestId(v string) *DeleteLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetSuccess(v bool) *DeleteLibraryResponseBody {
	s.Success = &v
	return s
}

type DeleteLibraryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibraryResponse) GoString() string {
	return s.String()
}

func (s *DeleteLibraryResponse) SetHeaders(v map[string]*string) *DeleteLibraryResponse {
	s.Headers = v
	return s
}

func (s *DeleteLibraryResponse) SetStatusCode(v int32) *DeleteLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLibraryResponse) SetBody(v *DeleteLibraryResponseBody) *DeleteLibraryResponse {
	s.Body = v
	return s
}

type EvictTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s EvictTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s EvictTaskRequest) GoString() string {
	return s.String()
}

func (s *EvictTaskRequest) SetTaskId(v string) *EvictTaskRequest {
	s.TaskId = &v
	return s
}

type EvictTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 17071319
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 44BD277A-87F9-5310-8D63-3E6645F1DA85
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EvictTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvictTaskResponseBody) GoString() string {
	return s.String()
}

func (s *EvictTaskResponseBody) SetCost(v int64) *EvictTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *EvictTaskResponseBody) SetData(v string) *EvictTaskResponseBody {
	s.Data = &v
	return s
}

func (s *EvictTaskResponseBody) SetDataType(v string) *EvictTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *EvictTaskResponseBody) SetErrCode(v string) *EvictTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *EvictTaskResponseBody) SetMessage(v string) *EvictTaskResponseBody {
	s.Message = &v
	return s
}

func (s *EvictTaskResponseBody) SetRequestId(v string) *EvictTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *EvictTaskResponseBody) SetSuccess(v bool) *EvictTaskResponseBody {
	s.Success = &v
	return s
}

func (s *EvictTaskResponseBody) SetTime(v string) *EvictTaskResponseBody {
	s.Time = &v
	return s
}

type EvictTaskResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EvictTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvictTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s EvictTaskResponse) GoString() string {
	return s.String()
}

func (s *EvictTaskResponse) SetHeaders(v map[string]*string) *EvictTaskResponse {
	s.Headers = v
	return s
}

func (s *EvictTaskResponse) SetStatusCode(v int32) *EvictTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *EvictTaskResponse) SetBody(v *EvictTaskResponseBody) *EvictTaskResponse {
	s.Body = v
	return s
}

type GenDocQaResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 182364872346
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdgdsfg
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenDocQaResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GenDocQaResultRequest) GoString() string {
	return s.String()
}

func (s *GenDocQaResultRequest) SetDocId(v string) *GenDocQaResultRequest {
	s.DocId = &v
	return s
}

func (s *GenDocQaResultRequest) SetLibraryId(v string) *GenDocQaResultRequest {
	s.LibraryId = &v
	return s
}

func (s *GenDocQaResultRequest) SetRequestId(v string) *GenDocQaResultRequest {
	s.RequestId = &v
	return s
}

type GenDocQaResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GenDocQaResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 44BD277A-87F9-5310-8D63-3E6645F1DA85
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GenDocQaResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenDocQaResultResponseBody) GoString() string {
	return s.String()
}

func (s *GenDocQaResultResponseBody) SetCost(v int64) *GenDocQaResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetData(v *GenDocQaResultResponseBodyData) *GenDocQaResultResponseBody {
	s.Data = v
	return s
}

func (s *GenDocQaResultResponseBody) SetDataType(v string) *GenDocQaResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetErrCode(v string) *GenDocQaResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetMessage(v string) *GenDocQaResultResponseBody {
	s.Message = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetRequestId(v string) *GenDocQaResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetSuccess(v bool) *GenDocQaResultResponseBody {
	s.Success = &v
	return s
}

func (s *GenDocQaResultResponseBody) SetTime(v string) *GenDocQaResultResponseBody {
	s.Time = &v
	return s
}

type GenDocQaResultResponseBodyData struct {
	// example:
	//
	// PROCESSING
	CurrentStatus *string `json:"currentStatus,omitempty" xml:"currentStatus,omitempty"`
	// example:
	//
	// 873648346573245
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// 7wxwrjpabj
	LibraryId      *string                                         `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	ParseQaResults []*GenDocQaResultResponseBodyDataParseQaResults `json:"parseQaResults,omitempty" xml:"parseQaResults,omitempty" type:"Repeated"`
}

func (s GenDocQaResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenDocQaResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenDocQaResultResponseBodyData) SetCurrentStatus(v string) *GenDocQaResultResponseBodyData {
	s.CurrentStatus = &v
	return s
}

func (s *GenDocQaResultResponseBodyData) SetDocId(v string) *GenDocQaResultResponseBodyData {
	s.DocId = &v
	return s
}

func (s *GenDocQaResultResponseBodyData) SetLibraryId(v string) *GenDocQaResultResponseBodyData {
	s.LibraryId = &v
	return s
}

func (s *GenDocQaResultResponseBodyData) SetParseQaResults(v []*GenDocQaResultResponseBodyDataParseQaResults) *GenDocQaResultResponseBodyData {
	s.ParseQaResults = v
	return s
}

type GenDocQaResultResponseBodyDataParseQaResults struct {
	Answer   *string `json:"answer,omitempty" xml:"answer,omitempty"`
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
}

func (s GenDocQaResultResponseBodyDataParseQaResults) String() string {
	return tea.Prettify(s)
}

func (s GenDocQaResultResponseBodyDataParseQaResults) GoString() string {
	return s.String()
}

func (s *GenDocQaResultResponseBodyDataParseQaResults) SetAnswer(v string) *GenDocQaResultResponseBodyDataParseQaResults {
	s.Answer = &v
	return s
}

func (s *GenDocQaResultResponseBodyDataParseQaResults) SetQuestion(v string) *GenDocQaResultResponseBodyDataParseQaResults {
	s.Question = &v
	return s
}

type GenDocQaResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenDocQaResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenDocQaResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GenDocQaResultResponse) GoString() string {
	return s.String()
}

func (s *GenDocQaResultResponse) SetHeaders(v map[string]*string) *GenDocQaResultResponse {
	s.Headers = v
	return s
}

func (s *GenDocQaResultResponse) SetStatusCode(v int32) *GenDocQaResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GenDocQaResultResponse) SetBody(v *GenDocQaResultResponseBody) *GenDocQaResultResponse {
	s.Body = v
	return s
}

type GetAppConfigResponseBody struct {
	// example:
	//
	// null
	Cost *int64                        `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetAppConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// None
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetAppConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponseBody) SetCost(v int64) *GetAppConfigResponseBody {
	s.Cost = &v
	return s
}

func (s *GetAppConfigResponseBody) SetData(v *GetAppConfigResponseBodyData) *GetAppConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAppConfigResponseBody) SetDataType(v string) *GetAppConfigResponseBody {
	s.DataType = &v
	return s
}

func (s *GetAppConfigResponseBody) SetErrCode(v string) *GetAppConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAppConfigResponseBody) SetMessage(v string) *GetAppConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppConfigResponseBody) SetRequestId(v string) *GetAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppConfigResponseBody) SetSuccess(v bool) *GetAppConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppConfigResponseBody) SetTime(v string) *GetAppConfigResponseBody {
	s.Time = &v
	return s
}

type GetAppConfigResponseBodyData struct {
	EmbeddingTypeList         []map[string]*string `json:"embeddingTypeList,omitempty" xml:"embeddingTypeList,omitempty" type:"Repeated"`
	FrontendConfig            map[string]*bool     `json:"frontendConfig,omitempty" xml:"frontendConfig,omitempty"`
	LibraryDocumentStatusList []map[string]*string `json:"libraryDocumentStatusList,omitempty" xml:"libraryDocumentStatusList,omitempty" type:"Repeated"`
	LlmHelperTypeList         []map[string]*string `json:"llmHelperTypeList,omitempty" xml:"llmHelperTypeList,omitempty" type:"Repeated"`
	TextIndexCategoryList     []*string            `json:"textIndexCategoryList,omitempty" xml:"textIndexCategoryList,omitempty" type:"Repeated"`
	VectorIndexCategoryList   []*string            `json:"vectorIndexCategoryList,omitempty" xml:"vectorIndexCategoryList,omitempty" type:"Repeated"`
}

func (s GetAppConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponseBodyData) SetEmbeddingTypeList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.EmbeddingTypeList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetFrontendConfig(v map[string]*bool) *GetAppConfigResponseBodyData {
	s.FrontendConfig = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetLibraryDocumentStatusList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.LibraryDocumentStatusList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetLlmHelperTypeList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.LlmHelperTypeList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetTextIndexCategoryList(v []*string) *GetAppConfigResponseBodyData {
	s.TextIndexCategoryList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetVectorIndexCategoryList(v []*string) *GetAppConfigResponseBodyData {
	s.VectorIndexCategoryList = v
	return s
}

type GetAppConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponse) SetHeaders(v map[string]*string) *GetAppConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAppConfigResponse) SetStatusCode(v int32) *GetAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppConfigResponse) SetBody(v *GetAppConfigResponseBody) *GetAppConfigResponse {
	s.Body = v
	return s
}

type GetChatQuestionRespRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1869307330227937280
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetChatQuestionRespRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChatQuestionRespRequest) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespRequest) SetBatchId(v string) *GetChatQuestionRespRequest {
	s.BatchId = &v
	return s
}

func (s *GetChatQuestionRespRequest) SetSessionId(v string) *GetChatQuestionRespRequest {
	s.SessionId = &v
	return s
}

type GetChatQuestionRespResponseBody struct {
	// example:
	//
	// null
	Cost *int64                               `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetChatQuestionRespResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 44BD277A-87F9-5310-8D63-3E6645F1DA85
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

func (s GetChatQuestionRespResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChatQuestionRespResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespResponseBody) SetCost(v int64) *GetChatQuestionRespResponseBody {
	s.Cost = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetData(v *GetChatQuestionRespResponseBodyData) *GetChatQuestionRespResponseBody {
	s.Data = v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetDataType(v string) *GetChatQuestionRespResponseBody {
	s.DataType = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetErrCode(v string) *GetChatQuestionRespResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetMessage(v string) *GetChatQuestionRespResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetRequestId(v string) *GetChatQuestionRespResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetSuccess(v bool) *GetChatQuestionRespResponseBody {
	s.Success = &v
	return s
}

func (s *GetChatQuestionRespResponseBody) SetTime(v string) *GetChatQuestionRespResponseBody {
	s.Time = &v
	return s
}

type GetChatQuestionRespResponseBodyData struct {
	// example:
	//
	// PROCESSING
	CurrentState *string                                            `json:"currentState,omitempty" xml:"currentState,omitempty"`
	QuestionList []*GetChatQuestionRespResponseBodyDataQuestionList `json:"questionList,omitempty" xml:"questionList,omitempty" type:"Repeated"`
}

func (s GetChatQuestionRespResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChatQuestionRespResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespResponseBodyData) SetCurrentState(v string) *GetChatQuestionRespResponseBodyData {
	s.CurrentState = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyData) SetQuestionList(v []*GetChatQuestionRespResponseBodyDataQuestionList) *GetChatQuestionRespResponseBodyData {
	s.QuestionList = v
	return s
}

type GetChatQuestionRespResponseBodyDataQuestionList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-11-17 10:05:00
	GmtCreate  *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	OriContent *string `json:"oriContent,omitempty" xml:"oriContent,omitempty"`
	Reply      *string `json:"reply,omitempty" xml:"reply,omitempty"`
	// example:
	//
	// 1732846760323001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// PRODUCT_QA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 39847834568436
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetChatQuestionRespResponseBodyDataQuestionList) String() string {
	return tea.Prettify(s)
}

func (s GetChatQuestionRespResponseBodyDataQuestionList) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetContent(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.Content = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetGmtCreate(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.GmtCreate = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetOriContent(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.OriContent = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetReply(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.Reply = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetSessionId(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.SessionId = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetType(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.Type = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetUserId(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.UserId = &v
	return s
}

func (s *GetChatQuestionRespResponseBodyDataQuestionList) SetUserName(v string) *GetChatQuestionRespResponseBodyDataQuestionList {
	s.UserName = &v
	return s
}

type GetChatQuestionRespResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatQuestionRespResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatQuestionRespResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChatQuestionRespResponse) GoString() string {
	return s.String()
}

func (s *GetChatQuestionRespResponse) SetHeaders(v map[string]*string) *GetChatQuestionRespResponse {
	s.Headers = v
	return s
}

func (s *GetChatQuestionRespResponse) SetStatusCode(v int32) *GetChatQuestionRespResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatQuestionRespResponse) SetBody(v *GetChatQuestionRespResponseBody) *GetChatQuestionRespResponse {
	s.Body = v
	return s
}

type GetDialogAnalysisResultRequest struct {
	// example:
	//
	// true
	Asc *bool `json:"asc,omitempty" xml:"asc,omitempty"`
	// example:
	//
	// 2024-09-23 09:20:02
	EndTime    *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	SessionIds []*string `json:"sessionIds,omitempty" xml:"sessionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-14 09:11:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// true
	UseUrl *bool `json:"useUrl,omitempty" xml:"useUrl,omitempty"`
}

func (s GetDialogAnalysisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDialogAnalysisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultRequest) SetAsc(v bool) *GetDialogAnalysisResultRequest {
	s.Asc = &v
	return s
}

func (s *GetDialogAnalysisResultRequest) SetEndTime(v string) *GetDialogAnalysisResultRequest {
	s.EndTime = &v
	return s
}

func (s *GetDialogAnalysisResultRequest) SetSessionIds(v []*string) *GetDialogAnalysisResultRequest {
	s.SessionIds = v
	return s
}

func (s *GetDialogAnalysisResultRequest) SetStartTime(v string) *GetDialogAnalysisResultRequest {
	s.StartTime = &v
	return s
}

func (s *GetDialogAnalysisResultRequest) SetUseUrl(v bool) *GetDialogAnalysisResultRequest {
	s.UseUrl = &v
	return s
}

type GetDialogAnalysisResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                   `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDialogAnalysisResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 88A006F0-B565-53BA-B38A-DBDF9D0B2935
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDialogAnalysisResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBody) SetCost(v int64) *GetDialogAnalysisResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetData(v *GetDialogAnalysisResultResponseBodyData) *GetDialogAnalysisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetDataType(v string) *GetDialogAnalysisResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetErrCode(v string) *GetDialogAnalysisResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetMessage(v string) *GetDialogAnalysisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetRequestId(v string) *GetDialogAnalysisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetSuccess(v bool) *GetDialogAnalysisResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBody) SetTime(v string) *GetDialogAnalysisResultResponseBody {
	s.Time = &v
	return s
}

type GetDialogAnalysisResultResponseBodyData struct {
	DialogAnalysisRespList []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList `json:"dialogAnalysisRespList,omitempty" xml:"dialogAnalysisRespList,omitempty" type:"Repeated"`
}

func (s GetDialogAnalysisResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBodyData) SetDialogAnalysisRespList(v []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) *GetDialogAnalysisResultResponseBodyData {
	s.DialogAnalysisRespList = v
	return s
}

type GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList struct {
	AnalysisResp *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp `json:"analysisResp,omitempty" xml:"analysisResp,omitempty" type:"Struct"`
	// example:
	//
	// 2024-04-24 11:54:34
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// https://xxx.oss-cn-beijing.aliyuncs.com/dialog-analysis/2024-12-30/2/1826661605606129665
	OssUrl *string `json:"ossUrl,omitempty" xml:"ossUrl,omitempty"`
	// example:
	//
	// 183764873624
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) String() string {
	return tea.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetAnalysisResp(v *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.AnalysisResp = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetGmtCreate(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.GmtCreate = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetOssUrl(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.OssUrl = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetSessionId(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.SessionId = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList) SetStatus(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespList {
	s.Status = &v
	return s
}

type GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp struct {
	DialogExecPlan        *string                                                                                  `json:"dialogExecPlan,omitempty" xml:"dialogExecPlan,omitempty"`
	DialogLabels          []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels `json:"dialogLabels,omitempty" xml:"dialogLabels,omitempty" type:"Repeated"`
	DialogOpenAnalysis    map[string]interface{}                                                                   `json:"dialogOpenAnalysis,omitempty" xml:"dialogOpenAnalysis,omitempty"`
	DialogProcessAnalysis map[string]interface{}                                                                   `json:"dialogProcessAnalysis,omitempty" xml:"dialogProcessAnalysis,omitempty"`
	DialogSop             *string                                                                                  `json:"dialogSop,omitempty" xml:"dialogSop,omitempty"`
	DialogSummary         *string                                                                                  `json:"dialogSummary,omitempty" xml:"dialogSummary,omitempty"`
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) String() string {
	return tea.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogExecPlan(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogExecPlan = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogLabels(v []*GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogLabels = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogOpenAnalysis(v map[string]interface{}) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogOpenAnalysis = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogProcessAnalysis(v map[string]interface{}) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogProcessAnalysis = v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogSop(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogSop = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp) SetDialogSummary(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisResp {
	s.DialogSummary = &v
	return s
}

type GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) String() string {
	return tea.Prettify(s)
}

func (s GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) SetName(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels {
	s.Name = &v
	return s
}

func (s *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels) SetValue(v string) *GetDialogAnalysisResultResponseBodyDataDialogAnalysisRespListAnalysisRespDialogLabels {
	s.Value = &v
	return s
}

type GetDialogAnalysisResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDialogAnalysisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDialogAnalysisResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDialogAnalysisResultResponse) GoString() string {
	return s.String()
}

func (s *GetDialogAnalysisResultResponse) SetHeaders(v map[string]*string) *GetDialogAnalysisResultResponse {
	s.Headers = v
	return s
}

func (s *GetDialogAnalysisResultResponse) SetStatusCode(v int32) *GetDialogAnalysisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDialogAnalysisResultResponse) SetBody(v *GetDialogAnalysisResultResponseBody) *GetDialogAnalysisResultResponse {
	s.Body = v
	return s
}

type GetDialogDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1906623923815534xxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetDialogDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDialogDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDialogDetailRequest) SetSessionId(v string) *GetDialogDetailRequest {
	s.SessionId = &v
	return s
}

type GetDialogDetailResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDialogDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDialogDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDialogDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDialogDetailResponseBody) SetCost(v int64) *GetDialogDetailResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetData(v *GetDialogDetailResponseBodyData) *GetDialogDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetDialogDetailResponseBody) SetDataType(v string) *GetDialogDetailResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetErrCode(v string) *GetDialogDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetMessage(v string) *GetDialogDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetRequestId(v string) *GetDialogDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetSuccess(v bool) *GetDialogDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetTime(v string) *GetDialogDetailResponseBody {
	s.Time = &v
	return s
}

type GetDialogDetailResponseBodyData struct {
	DialogueList []*GetDialogDetailResponseBodyDataDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 10
	TotalDialogTurns *int32 `json:"totalDialogTurns,omitempty" xml:"totalDialogTurns,omitempty"`
	// example:
	//
	// 5
	ValidDialogTurns *int32 `json:"validDialogTurns,omitempty" xml:"validDialogTurns,omitempty"`
}

func (s GetDialogDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDialogDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDialogDetailResponseBodyData) SetDialogueList(v []*GetDialogDetailResponseBodyDataDialogueList) *GetDialogDetailResponseBodyData {
	s.DialogueList = v
	return s
}

func (s *GetDialogDetailResponseBodyData) SetGmtCreate(v string) *GetDialogDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetDialogDetailResponseBodyData) SetStatus(v string) *GetDialogDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDialogDetailResponseBodyData) SetTotalDialogTurns(v int32) *GetDialogDetailResponseBodyData {
	s.TotalDialogTurns = &v
	return s
}

func (s *GetDialogDetailResponseBodyData) SetValidDialogTurns(v int32) *GetDialogDetailResponseBodyData {
	s.ValidDialogTurns = &v
	return s
}

type GetDialogDetailResponseBodyDataDialogueList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 123761283
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// BOT
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// true
	HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
	// example:
	//
	// 1742869659849
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 193874634xxx
	IntentCode *string `json:"intentCode,omitempty" xml:"intentCode,omitempty"`
	IntentName *string `json:"intentName,omitempty" xml:"intentName,omitempty"`
	// example:
	//
	// 0
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDialogDetailResponseBodyDataDialogueList) String() string {
	return tea.Prettify(s)
}

func (s GetDialogDetailResponseBodyDataDialogueList) GoString() string {
	return s.String()
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetContent(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.Content = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetCustomerId(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.CustomerId = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetCustomerServiceId(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.CustomerServiceId = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetCustomerServiceType(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.CustomerServiceType = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetHangUpDialog(v bool) *GetDialogDetailResponseBodyDataDialogueList {
	s.HangUpDialog = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetId(v int64) *GetDialogDetailResponseBodyDataDialogueList {
	s.Id = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetIntentCode(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.IntentCode = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetIntentName(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.IntentName = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetRole(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.Role = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetType(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.Type = &v
	return s
}

type GetDialogDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDialogDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDialogDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDialogDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDialogDetailResponse) SetHeaders(v map[string]*string) *GetDialogDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDialogDetailResponse) SetStatusCode(v int32) *GetDialogDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDialogDetailResponse) SetBody(v *GetDialogDetailResponseBody) *GetDialogDetailResponse {
	s.Body = v
	return s
}

type GetDocumentChunkListRequest struct {
	ChunkIdList []*string `json:"chunkIdList,omitempty" xml:"chunkIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 182364872346
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dsjgfdjgfxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
}

func (s GetDocumentChunkListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListRequest) SetChunkIdList(v []*string) *GetDocumentChunkListRequest {
	s.ChunkIdList = v
	return s
}

func (s *GetDocumentChunkListRequest) SetDocId(v string) *GetDocumentChunkListRequest {
	s.DocId = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetLibraryId(v string) *GetDocumentChunkListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetOrder(v string) *GetDocumentChunkListRequest {
	s.Order = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetOrderBy(v string) *GetDocumentChunkListRequest {
	s.OrderBy = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetPage(v int32) *GetDocumentChunkListRequest {
	s.Page = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetPageSize(v int32) *GetDocumentChunkListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetSearchQuery(v string) *GetDocumentChunkListRequest {
	s.SearchQuery = &v
	return s
}

type GetDocumentChunkListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDocumentChunkListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 2B8F6DC9-6FAF-576F-9095-CCD90FB2BDDF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDocumentChunkListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBody) SetCost(v int64) *GetDocumentChunkListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetData(v *GetDocumentChunkListResponseBodyData) *GetDocumentChunkListResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetDataType(v string) *GetDocumentChunkListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetErrCode(v string) *GetDocumentChunkListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetMessage(v string) *GetDocumentChunkListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetRequestId(v string) *GetDocumentChunkListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetSuccess(v bool) *GetDocumentChunkListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetTime(v string) *GetDocumentChunkListResponseBody {
	s.Time = &v
	return s
}

type GetDocumentChunkListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetDocumentChunkListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetDocumentChunkListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyData) SetCurrentPage(v int64) *GetDocumentChunkListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetPageSize(v int64) *GetDocumentChunkListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetRecords(v []*GetDocumentChunkListResponseBodyDataRecords) *GetDocumentChunkListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetTotalPages(v int64) *GetDocumentChunkListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetTotalRecords(v int64) *GetDocumentChunkListResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetDocumentChunkListResponseBodyDataRecords struct {
	// example:
	//
	// 28377468263482764
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// oss-xxxx-hangzhou.com/test.pdf
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 8947387648356
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// jhsdvne
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 947538465
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*GetDocumentChunkListResponseBodyDataRecordsPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 9848346548365
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetDocumentChunkListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkMeta(v map[string]interface{}) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkMeta = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkOssUrl(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkOssUrl = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkText(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkText = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkType(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkType = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetDocId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetFileType(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetLibraryId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetLibraryName(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.LibraryName = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetNextChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.NextChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetPos(v []*GetDocumentChunkListResponseBodyDataRecordsPos) *GetDocumentChunkListResponseBodyDataRecords {
	s.Pos = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetPreChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.PreChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetScore(v float32) *GetDocumentChunkListResponseBodyDataRecords {
	s.Score = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetTitle(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.Title = &v
	return s
}

type GetDocumentChunkListResponseBodyDataRecordsPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s GetDocumentChunkListResponseBodyDataRecordsPos) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyDataRecordsPos) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetAxisArray(v []*float64) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.AxisArray = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetPage(v int32) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.Page = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetTextHighlightArea(v []*int32) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.TextHighlightArea = v
	return s
}

type GetDocumentChunkListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentChunkListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentChunkListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponse) SetHeaders(v map[string]*string) *GetDocumentChunkListResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentChunkListResponse) SetStatusCode(v int32) *GetDocumentChunkListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentChunkListResponse) SetBody(v *GetDocumentChunkListResponseBody) *GetDocumentChunkListResponse {
	s.Body = v
	return s
}

type GetDocumentListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetDocumentListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentListRequest) SetLibraryId(v string) *GetDocumentListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentListRequest) SetPage(v int32) *GetDocumentListRequest {
	s.Page = &v
	return s
}

func (s *GetDocumentListRequest) SetPageSize(v int32) *GetDocumentListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentListRequest) SetStatus(v string) *GetDocumentListRequest {
	s.Status = &v
	return s
}

type GetDocumentListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDocumentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDocumentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBody) SetCost(v int64) *GetDocumentListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentListResponseBody) SetData(v *GetDocumentListResponseBodyData) *GetDocumentListResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentListResponseBody) SetDataType(v string) *GetDocumentListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentListResponseBody) SetErrCode(v string) *GetDocumentListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentListResponseBody) SetMessage(v string) *GetDocumentListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentListResponseBody) SetRequestId(v string) *GetDocumentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentListResponseBody) SetSuccess(v bool) *GetDocumentListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentListResponseBody) SetTime(v string) *GetDocumentListResponseBody {
	s.Time = &v
	return s
}

type GetDocumentListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetDocumentListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetDocumentListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBodyData) SetCurrentPage(v int64) *GetDocumentListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetPageSize(v int64) *GetDocumentListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetRecords(v []*GetDocumentListResponseBodyDataRecords) *GetDocumentListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetDocumentListResponseBodyData) SetTotalPages(v int64) *GetDocumentListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetTotalRecords(v int64) *GetDocumentListResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetDocumentListResponseBodyDataRecords struct {
	// example:
	//
	// 8326748346
	DocId        *string                `json:"docId,omitempty" xml:"docId,omitempty"`
	DocumentMeta map[string]interface{} `json:"documentMeta,omitempty" xml:"documentMeta,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
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
	// skjdhshbv
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// WaitRefresh
	StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// null
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDocumentListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBodyDataRecords) SetDocId(v string) *GetDocumentListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetDocumentMeta(v map[string]interface{}) *GetDocumentListResponseBodyDataRecords {
	s.DocumentMeta = v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetFileType(v string) *GetDocumentListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetGmtCreate(v string) *GetDocumentListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetGmtModified(v string) *GetDocumentListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetLibraryId(v string) *GetDocumentListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetStatusCode(v string) *GetDocumentListResponseBodyDataRecords {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetTitle(v string) *GetDocumentListResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetUrl(v string) *GetDocumentListResponseBodyDataRecords {
	s.Url = &v
	return s
}

type GetDocumentListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponse) SetHeaders(v map[string]*string) *GetDocumentListResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentListResponse) SetStatusCode(v int32) *GetDocumentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentListResponse) SetBody(v *GetDocumentListResponseBody) *GetDocumentListResponse {
	s.Body = v
	return s
}

type GetDocumentUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12681367362
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s GetDocumentUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlRequest) SetDocumentId(v string) *GetDocumentUrlRequest {
	s.DocumentId = &v
	return s
}

type GetDocumentUrlResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// https://path_to_file
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// 66249B43-8C2B-5EE7-AE78-B382306621C6
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

func (s GetDocumentUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlResponseBody) SetCost(v int64) *GetDocumentUrlResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetData(v string) *GetDocumentUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetDataType(v string) *GetDocumentUrlResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetErrCode(v string) *GetDocumentUrlResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetMessage(v string) *GetDocumentUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetRequestId(v string) *GetDocumentUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetSuccess(v bool) *GetDocumentUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetTime(v string) *GetDocumentUrlResponseBody {
	s.Time = &v
	return s
}

type GetDocumentUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlResponse) SetHeaders(v map[string]*string) *GetDocumentUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentUrlResponse) SetStatusCode(v int32) *GetDocumentUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentUrlResponse) SetBody(v *GetDocumentUrlResponseBody) *GetDocumentUrlResponse {
	s.Body = v
	return s
}

type GetFilterDocumentListRequest struct {
	And       []*GetFilterDocumentListRequestAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	DocIdList []*string                          `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string                           `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	Or        []*GetFilterDocumentListRequestOr `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
}

func (s GetFilterDocumentListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListRequest) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequest) SetAnd(v []*GetFilterDocumentListRequestAnd) *GetFilterDocumentListRequest {
	s.And = v
	return s
}

func (s *GetFilterDocumentListRequest) SetDocIdList(v []*string) *GetFilterDocumentListRequest {
	s.DocIdList = v
	return s
}

func (s *GetFilterDocumentListRequest) SetLibraryId(v string) *GetFilterDocumentListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetOr(v []*GetFilterDocumentListRequestOr) *GetFilterDocumentListRequest {
	s.Or = v
	return s
}

func (s *GetFilterDocumentListRequest) SetPage(v int32) *GetFilterDocumentListRequest {
	s.Page = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetPageSize(v int32) *GetFilterDocumentListRequest {
	s.PageSize = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetStatus(v []*string) *GetFilterDocumentListRequest {
	s.Status = v
	return s
}

type GetFilterDocumentListRequestAnd struct {
	// example:
	//
	// 1
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// company
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// alibaba
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFilterDocumentListRequestAnd) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListRequestAnd) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequestAnd) SetBoost(v float32) *GetFilterDocumentListRequestAnd {
	s.Boost = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetKey(v string) *GetFilterDocumentListRequestAnd {
	s.Key = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetOperator(v string) *GetFilterDocumentListRequestAnd {
	s.Operator = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetValue(v string) *GetFilterDocumentListRequestAnd {
	s.Value = &v
	return s
}

type GetFilterDocumentListRequestOr struct {
	// example:
	//
	// 1
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// company
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// contains
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// alibaba
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFilterDocumentListRequestOr) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListRequestOr) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequestOr) SetBoost(v float32) *GetFilterDocumentListRequestOr {
	s.Boost = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetKey(v string) *GetFilterDocumentListRequestOr {
	s.Key = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetOperator(v string) *GetFilterDocumentListRequestOr {
	s.Operator = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetValue(v string) *GetFilterDocumentListRequestOr {
	s.Value = &v
	return s
}

type GetFilterDocumentListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                 `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetFilterDocumentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 7ADF010C-FD89-569D-A079-2D4D5247E943
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

func (s GetFilterDocumentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBody) SetCost(v int64) *GetFilterDocumentListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetData(v *GetFilterDocumentListResponseBodyData) *GetFilterDocumentListResponseBody {
	s.Data = v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetDataType(v string) *GetFilterDocumentListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetErrCode(v string) *GetFilterDocumentListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetMessage(v string) *GetFilterDocumentListResponseBody {
	s.Message = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetRequestId(v string) *GetFilterDocumentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetSuccess(v bool) *GetFilterDocumentListResponseBody {
	s.Success = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetTime(v string) *GetFilterDocumentListResponseBody {
	s.Time = &v
	return s
}

type GetFilterDocumentListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                          `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetFilterDocumentListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetFilterDocumentListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBodyData) SetCurrentPage(v int64) *GetFilterDocumentListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetPageSize(v int64) *GetFilterDocumentListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetRecords(v []*GetFilterDocumentListResponseBodyDataRecords) *GetFilterDocumentListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetTotalPages(v int64) *GetFilterDocumentListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetTotalRecords(v int64) *GetFilterDocumentListResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetFilterDocumentListResponseBodyDataRecords struct {
	// example:
	//
	// 29368126816
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// {"a": "1"}
	DocumentMeta map[string]interface{} `json:"documentMeta,omitempty" xml:"documentMeta,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
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
	// sdfgsjdfg
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// WaitRefresh
	StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// null
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetFilterDocumentListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetDocId(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetDocumentMeta(v map[string]interface{}) *GetFilterDocumentListResponseBodyDataRecords {
	s.DocumentMeta = v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetFileType(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetGmtCreate(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetGmtModified(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetLibraryId(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetStatusCode(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.StatusCode = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetTitle(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetUrl(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.Url = &v
	return s
}

type GetFilterDocumentListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFilterDocumentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFilterDocumentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListResponse) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponse) SetHeaders(v map[string]*string) *GetFilterDocumentListResponse {
	s.Headers = v
	return s
}

func (s *GetFilterDocumentListResponse) SetStatusCode(v int32) *GetFilterDocumentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFilterDocumentListResponse) SetBody(v *GetFilterDocumentListResponseBody) *GetFilterDocumentListResponse {
	s.Body = v
	return s
}

type GetHistoryListByBizTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GysYBsxx
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// LibraryChat
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetHistoryListByBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeRequest) SetBizId(v string) *GetHistoryListByBizTypeRequest {
	s.BizId = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetBizType(v string) *GetHistoryListByBizTypeRequest {
	s.BizType = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetPage(v int32) *GetHistoryListByBizTypeRequest {
	s.Page = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetPageSize(v int32) *GetHistoryListByBizTypeRequest {
	s.PageSize = &v
	return s
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
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBodyData) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBodyDataRecords) GoString() string {
	return s.String()
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

type GetHistoryListByBizTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoryListByBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoryListByBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponse) SetHeaders(v map[string]*string) *GetHistoryListByBizTypeResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryListByBizTypeResponse) SetStatusCode(v int32) *GetHistoryListByBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryListByBizTypeResponse) SetBody(v *GetHistoryListByBizTypeResponseBody) *GetHistoryListByBizTypeResponse {
	s.Body = v
	return s
}

type GetLibraryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s GetLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryRequest) GoString() string {
	return s.String()
}

func (s *GetLibraryRequest) SetLibraryId(v string) *GetLibraryRequest {
	s.LibraryId = &v
	return s
}

type GetLibraryResponseBody struct {
	// example:
	//
	// null
	Cost *int64                      `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetLibraryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 44BD277A-87F9-5310-8D63-3E6645F1DA85
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

func (s GetLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBody) SetCost(v int64) *GetLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *GetLibraryResponseBody) SetData(v *GetLibraryResponseBodyData) *GetLibraryResponseBody {
	s.Data = v
	return s
}

func (s *GetLibraryResponseBody) SetDataType(v string) *GetLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *GetLibraryResponseBody) SetErrCode(v string) *GetLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetLibraryResponseBody) SetMessage(v string) *GetLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *GetLibraryResponseBody) SetRequestId(v string) *GetLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLibraryResponseBody) SetSuccess(v bool) *GetLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *GetLibraryResponseBody) SetTime(v string) *GetLibraryResponseBody {
	s.Time = &v
	return s
}

type GetLibraryResponseBodyData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 10
	DocumentCount *int64 `json:"documentCount,omitempty" xml:"documentCount,omitempty"`
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
	// 19386728376
	Id           *string                                 `json:"id,omitempty" xml:"id,omitempty"`
	IndexSetting *GetLibraryResponseBodyDataIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	LibraryName  *string                                 `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s GetLibraryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyData) SetDescription(v string) *GetLibraryResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetDocumentCount(v int64) *GetLibraryResponseBodyData {
	s.DocumentCount = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetGmtCreate(v string) *GetLibraryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetGmtModified(v string) *GetLibraryResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetId(v string) *GetLibraryResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetIndexSetting(v *GetLibraryResponseBodyDataIndexSetting) *GetLibraryResponseBodyData {
	s.IndexSetting = v
	return s
}

func (s *GetLibraryResponseBodyData) SetLibraryName(v string) *GetLibraryResponseBodyData {
	s.LibraryName = &v
	return s
}

type GetLibraryResponseBodyDataIndexSetting struct {
	ChunkStrategy      *GetLibraryResponseBodyDataIndexSettingChunkStrategy      `json:"chunkStrategy,omitempty" xml:"chunkStrategy,omitempty" type:"Struct"`
	ModelConfig        *GetLibraryResponseBodyDataIndexSettingModelConfig        `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	PromptRoleStyle    *string                                                   `json:"promptRoleStyle,omitempty" xml:"promptRoleStyle,omitempty"`
	QueryEnhancer      *GetLibraryResponseBodyDataIndexSettingQueryEnhancer      `json:"queryEnhancer,omitempty" xml:"queryEnhancer,omitempty" type:"Struct"`
	RecallStrategy     *GetLibraryResponseBodyDataIndexSettingRecallStrategy     `json:"recallStrategy,omitempty" xml:"recallStrategy,omitempty" type:"Struct"`
	TextIndexSetting   *GetLibraryResponseBodyDataIndexSettingTextIndexSetting   `json:"textIndexSetting,omitempty" xml:"textIndexSetting,omitempty" type:"Struct"`
	VectorIndexSetting *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting `json:"vectorIndexSetting,omitempty" xml:"vectorIndexSetting,omitempty" type:"Struct"`
}

func (s GetLibraryResponseBodyDataIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetChunkStrategy(v *GetLibraryResponseBodyDataIndexSettingChunkStrategy) *GetLibraryResponseBodyDataIndexSetting {
	s.ChunkStrategy = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetModelConfig(v *GetLibraryResponseBodyDataIndexSettingModelConfig) *GetLibraryResponseBodyDataIndexSetting {
	s.ModelConfig = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetPromptRoleStyle(v string) *GetLibraryResponseBodyDataIndexSetting {
	s.PromptRoleStyle = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetQueryEnhancer(v *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) *GetLibraryResponseBodyDataIndexSetting {
	s.QueryEnhancer = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetRecallStrategy(v *GetLibraryResponseBodyDataIndexSettingRecallStrategy) *GetLibraryResponseBodyDataIndexSetting {
	s.RecallStrategy = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetTextIndexSetting(v *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) *GetLibraryResponseBodyDataIndexSetting {
	s.TextIndexSetting = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetVectorIndexSetting(v *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) *GetLibraryResponseBodyDataIndexSetting {
	s.VectorIndexSetting = v
	return s
}

type GetLibraryResponseBodyDataIndexSettingChunkStrategy struct {
	// example:
	//
	// true
	DocTreeSplit *bool `json:"docTreeSplit,omitempty" xml:"docTreeSplit,omitempty"`
	// example:
	//
	// 160
	DocTreeSplitSize *int32 `json:"docTreeSplitSize,omitempty" xml:"docTreeSplitSize,omitempty"`
	// example:
	//
	// true
	EnhanceGraph *bool `json:"enhanceGraph,omitempty" xml:"enhanceGraph,omitempty"`
	// example:
	//
	// true
	EnhanceTable *bool `json:"enhanceTable,omitempty" xml:"enhanceTable,omitempty"`
	// example:
	//
	// 40
	Overlap *int32 `json:"overlap,omitempty" xml:"overlap,omitempty"`
	// example:
	//
	// true
	SentenceSplit *bool `json:"sentenceSplit,omitempty" xml:"sentenceSplit,omitempty"`
	// example:
	//
	// 160
	SentenceSplitSize *int32 `json:"sentenceSplitSize,omitempty" xml:"sentenceSplitSize,omitempty"`
	// example:
	//
	// 256
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// true
	Split *bool `json:"split,omitempty" xml:"split,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingChunkStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetDocTreeSplit(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.DocTreeSplit = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetDocTreeSplitSize(v int32) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.DocTreeSplitSize = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetEnhanceGraph(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.EnhanceGraph = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetEnhanceTable(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.EnhanceTable = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetOverlap(v int32) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.Overlap = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetSentenceSplit(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.SentenceSplit = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetSentenceSplitSize(v int32) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.SentenceSplitSize = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetSize(v int32) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.Size = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetSplit(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.Split = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingModelConfig struct {
	// example:
	//
	// 0.8
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// topP
	//
	// example:
	//
	// 0.8
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingModelConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) SetTemperature(v float64) *GetLibraryResponseBodyDataIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) SetTopP(v float64) *GetLibraryResponseBodyDataIndexSettingModelConfig {
	s.TopP = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingQueryEnhancer struct {
	// example:
	//
	// true
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// true
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// true
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// true
	EnableQueryRewrite *bool `json:"enableQueryRewrite,omitempty" xml:"enableQueryRewrite,omitempty"`
	// example:
	//
	// true
	EnableSession *bool `json:"enableSession,omitempty" xml:"enableSession,omitempty"`
	// example:
	//
	// 2836482634
	LocalKnowledgeId *string `json:"localKnowledgeId,omitempty" xml:"localKnowledgeId,omitempty"`
	// example:
	//
	// true
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingQueryEnhancer) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableFollowUp(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableFollowUp = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableMultiQuery(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableMultiQuery = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableOpenQa(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableOpenQa = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableQueryRewrite(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableQueryRewrite = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableSession(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableSession = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetLocalKnowledgeId(v string) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.LocalKnowledgeId = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetWithDocumentReference(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.WithDocumentReference = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingRecallStrategy struct {
	// example:
	//
	// model
	DocumentRankType *string `json:"documentRankType,omitempty" xml:"documentRankType,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingRecallStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) SetDocumentRankType(v string) *GetLibraryResponseBodyDataIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) SetLimit(v int32) *GetLibraryResponseBodyDataIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingTextIndexSetting struct {
	// example:
	//
	// ElasticSearch
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// IkMaxWord
	IndexAnalyzer *string `json:"indexAnalyzer,omitempty" xml:"indexAnalyzer,omitempty"`
	// example:
	//
	// null
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// Standard
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" xml:"searchAnalyzer,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingTextIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetCategory(v string) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.Category = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetEnable(v bool) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.Enable = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetIndexAnalyzer(v string) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.IndexAnalyzer = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetRankThreshold(v float64) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetSearchAnalyzer(v string) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.SearchAnalyzer = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetTopK(v int32) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.TopK = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingVectorIndexSetting struct {
	// example:
	//
	// ADB
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// DashScope
	EmbeddingType *string `json:"embeddingType,omitempty" xml:"embeddingType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// null
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetCategory(v string) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.Category = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetEmbeddingType(v string) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.EmbeddingType = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetEnable(v bool) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.Enable = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetRankThreshold(v float64) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetTopK(v int32) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.TopK = &v
	return s
}

type GetLibraryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryResponse) SetHeaders(v map[string]*string) *GetLibraryResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryResponse) SetStatusCode(v int32) *GetLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryResponse) SetBody(v *GetLibraryResponseBody) *GetLibraryResponse {
	s.Body = v
	return s
}

type GetLibraryListRequest struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query    *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetLibraryListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListRequest) GoString() string {
	return s.String()
}

func (s *GetLibraryListRequest) SetPage(v int32) *GetLibraryListRequest {
	s.Page = &v
	return s
}

func (s *GetLibraryListRequest) SetPageSize(v int32) *GetLibraryListRequest {
	s.PageSize = &v
	return s
}

func (s *GetLibraryListRequest) SetQuery(v string) *GetLibraryListRequest {
	s.Query = &v
	return s
}

type GetLibraryListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetLibraryListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 0a06dfe817156528535968405edce3
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

func (s GetLibraryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBody) SetCost(v int64) *GetLibraryListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetLibraryListResponseBody) SetData(v *GetLibraryListResponseBodyData) *GetLibraryListResponseBody {
	s.Data = v
	return s
}

func (s *GetLibraryListResponseBody) SetDataType(v string) *GetLibraryListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetLibraryListResponseBody) SetErrCode(v string) *GetLibraryListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetLibraryListResponseBody) SetMessage(v string) *GetLibraryListResponseBody {
	s.Message = &v
	return s
}

func (s *GetLibraryListResponseBody) SetRequestId(v string) *GetLibraryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLibraryListResponseBody) SetSuccess(v bool) *GetLibraryListResponseBody {
	s.Success = &v
	return s
}

func (s *GetLibraryListResponseBody) SetTime(v string) *GetLibraryListResponseBody {
	s.Time = &v
	return s
}

type GetLibraryListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetLibraryListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetLibraryListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyData) SetCurrentPage(v int64) *GetLibraryListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetLibraryListResponseBodyData) SetPageSize(v int64) *GetLibraryListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetLibraryListResponseBodyData) SetRecords(v []*GetLibraryListResponseBodyDataRecords) *GetLibraryListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetLibraryListResponseBodyData) SetTotalPages(v int64) *GetLibraryListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetLibraryListResponseBodyData) SetTotalRecords(v int64) *GetLibraryListResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetLibraryListResponseBodyDataRecords struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 10
	DocumentCount *int64 `json:"documentCount,omitempty" xml:"documentCount,omitempty"`
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
	// 24vs4aa42jv1rg7
	Id           *string                                            `json:"id,omitempty" xml:"id,omitempty"`
	IndexSetting *GetLibraryListResponseBodyDataRecordsIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	LibraryName  *string                                            `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecords) SetDescription(v string) *GetLibraryListResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetDocumentCount(v int64) *GetLibraryListResponseBodyDataRecords {
	s.DocumentCount = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetGmtCreate(v string) *GetLibraryListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetGmtModified(v string) *GetLibraryListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetId(v string) *GetLibraryListResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetIndexSetting(v *GetLibraryListResponseBodyDataRecordsIndexSetting) *GetLibraryListResponseBodyDataRecords {
	s.IndexSetting = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetLibraryName(v string) *GetLibraryListResponseBodyDataRecords {
	s.LibraryName = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSetting struct {
	ChunkStrategy      *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy      `json:"chunkStrategy,omitempty" xml:"chunkStrategy,omitempty" type:"Struct"`
	ModelConfig        *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig        `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	PromptRoleStyle    *string                                                              `json:"promptRoleStyle,omitempty" xml:"promptRoleStyle,omitempty"`
	QueryEnhancer      *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer      `json:"queryEnhancer,omitempty" xml:"queryEnhancer,omitempty" type:"Struct"`
	RecallStrategy     *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy     `json:"recallStrategy,omitempty" xml:"recallStrategy,omitempty" type:"Struct"`
	TextIndexSetting   *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting   `json:"textIndexSetting,omitempty" xml:"textIndexSetting,omitempty" type:"Struct"`
	VectorIndexSetting *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting `json:"vectorIndexSetting,omitempty" xml:"vectorIndexSetting,omitempty" type:"Struct"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetChunkStrategy(v *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.ChunkStrategy = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetModelConfig(v *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.ModelConfig = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetPromptRoleStyle(v string) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.PromptRoleStyle = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetQueryEnhancer(v *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.QueryEnhancer = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetRecallStrategy(v *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.RecallStrategy = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetTextIndexSetting(v *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.TextIndexSetting = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetVectorIndexSetting(v *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.VectorIndexSetting = v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy struct {
	// example:
	//
	// true
	DocTreeSplit *bool `json:"docTreeSplit,omitempty" xml:"docTreeSplit,omitempty"`
	// example:
	//
	// 160
	DocTreeSplitSize *int32 `json:"docTreeSplitSize,omitempty" xml:"docTreeSplitSize,omitempty"`
	// example:
	//
	// true
	EnhanceGraph *bool `json:"enhanceGraph,omitempty" xml:"enhanceGraph,omitempty"`
	// example:
	//
	// true
	EnhanceTable *bool `json:"enhanceTable,omitempty" xml:"enhanceTable,omitempty"`
	// example:
	//
	// 40
	Overlap *int32 `json:"overlap,omitempty" xml:"overlap,omitempty"`
	// example:
	//
	// true
	SentenceSplit *bool `json:"sentenceSplit,omitempty" xml:"sentenceSplit,omitempty"`
	// example:
	//
	// 160
	SentenceSplitSize *int32 `json:"sentenceSplitSize,omitempty" xml:"sentenceSplitSize,omitempty"`
	// example:
	//
	// 256
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// true
	Split *bool `json:"split,omitempty" xml:"split,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetDocTreeSplit(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.DocTreeSplit = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetDocTreeSplitSize(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.DocTreeSplitSize = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetEnhanceGraph(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.EnhanceGraph = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetEnhanceTable(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.EnhanceTable = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetOverlap(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.Overlap = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetSentenceSplit(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.SentenceSplit = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetSentenceSplitSize(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.SentenceSplitSize = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetSize(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.Size = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetSplit(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.Split = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig struct {
	// example:
	//
	// 0.8
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// topP
	//
	// example:
	//
	// 0.8
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) SetTemperature(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) SetTopP(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig {
	s.TopP = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer struct {
	// example:
	//
	// true
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// true
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// true
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// true
	EnableQueryRewrite *bool `json:"enableQueryRewrite,omitempty" xml:"enableQueryRewrite,omitempty"`
	// example:
	//
	// true
	EnableSession *bool `json:"enableSession,omitempty" xml:"enableSession,omitempty"`
	// example:
	//
	// sdbcjsbc
	LocalKnowledgeId *string `json:"localKnowledgeId,omitempty" xml:"localKnowledgeId,omitempty"`
	// example:
	//
	// true
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableFollowUp(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableFollowUp = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableMultiQuery(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableMultiQuery = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableOpenQa(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableOpenQa = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableQueryRewrite(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableQueryRewrite = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableSession(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableSession = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetLocalKnowledgeId(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.LocalKnowledgeId = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetWithDocumentReference(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.WithDocumentReference = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy struct {
	// example:
	//
	// model
	DocumentRankType *string `json:"documentRankType,omitempty" xml:"documentRankType,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) SetDocumentRankType(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) SetLimit(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting struct {
	// example:
	//
	// ElasticSearch
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// Standard
	IndexAnalyzer *string `json:"indexAnalyzer,omitempty" xml:"indexAnalyzer,omitempty"`
	// example:
	//
	// null
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// Standard
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" xml:"searchAnalyzer,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetCategory(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.Category = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetEnable(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.Enable = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetIndexAnalyzer(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.IndexAnalyzer = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetRankThreshold(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetSearchAnalyzer(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.SearchAnalyzer = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetTopK(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.TopK = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting struct {
	// example:
	//
	// ADB
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// DashScope
	EmbeddingType *string `json:"embeddingType,omitempty" xml:"embeddingType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// null
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetCategory(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.Category = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetEmbeddingType(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.EmbeddingType = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetEnable(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.Enable = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetRankThreshold(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetTopK(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.TopK = &v
	return s
}

type GetLibraryListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibraryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponse) SetHeaders(v map[string]*string) *GetLibraryListResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryListResponse) SetStatusCode(v int32) *GetLibraryListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryListResponse) SetBody(v *GetLibraryListResponseBody) *GetLibraryListResponse {
	s.Body = v
	return s
}

type GetParseResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 873648346573245
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdgdsfg
	LibraryId    *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	UseUrlResult *bool   `json:"useUrlResult,omitempty" xml:"useUrlResult,omitempty"`
}

func (s GetParseResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParseResultRequest) GoString() string {
	return s.String()
}

func (s *GetParseResultRequest) SetDocId(v string) *GetParseResultRequest {
	s.DocId = &v
	return s
}

func (s *GetParseResultRequest) SetLibraryId(v string) *GetParseResultRequest {
	s.LibraryId = &v
	return s
}

func (s *GetParseResultRequest) SetUseUrlResult(v bool) *GetParseResultRequest {
	s.UseUrlResult = &v
	return s
}

type GetParseResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetParseResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 0abb793617204049360065953ec6dd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetParseResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParseResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetParseResultResponseBody) SetCost(v int64) *GetParseResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetParseResultResponseBody) SetData(v *GetParseResultResponseBodyData) *GetParseResultResponseBody {
	s.Data = v
	return s
}

func (s *GetParseResultResponseBody) SetDataType(v string) *GetParseResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetParseResultResponseBody) SetErrCode(v string) *GetParseResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetParseResultResponseBody) SetMessage(v string) *GetParseResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetParseResultResponseBody) SetRequestId(v string) *GetParseResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParseResultResponseBody) SetSuccess(v bool) *GetParseResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetParseResultResponseBody) SetTime(v string) *GetParseResultResponseBody {
	s.Time = &v
	return s
}

type GetParseResultResponseBodyData struct {
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// null
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
	// example:
	//
	// b0a202e2-5031-4589-a6d7-39185f0d8d01
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {
	//
	//           "Status": "Success",
	//
	//           "Data": {},
	//
	//           "Message": null,
	//
	//           "TaskId": "docmind-20240601-123abc"
	//
	//         }
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	ResultUrl *string                `json:"resultUrl,omitempty" xml:"resultUrl,omitempty"`
	// example:
	//
	// WaitRefresh
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetParseResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetParseResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetParseResultResponseBodyData) SetFileType(v string) *GetParseResultResponseBodyData {
	s.FileType = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetProviderType(v string) *GetParseResultResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetRequestId(v string) *GetParseResultResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetResult(v map[string]interface{}) *GetParseResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetParseResultResponseBodyData) SetResultUrl(v string) *GetParseResultResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetStatus(v string) *GetParseResultResponseBodyData {
	s.Status = &v
	return s
}

type GetParseResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParseResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParseResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParseResultResponse) GoString() string {
	return s.String()
}

func (s *GetParseResultResponse) SetHeaders(v map[string]*string) *GetParseResultResponse {
	s.Headers = v
	return s
}

func (s *GetParseResultResponse) SetStatusCode(v int32) *GetParseResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParseResultResponse) SetBody(v *GetParseResultResponseBody) *GetParseResultResponse {
	s.Body = v
	return s
}

type GetQualityCheckTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetQualityCheckTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultRequest) SetTaskId(v string) *GetQualityCheckTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetQualityCheckTaskResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                     `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetQualityCheckTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 67C7021A-D268-553D-8C15-A087B9604028
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

func (s GetQualityCheckTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBody) SetCost(v int64) *GetQualityCheckTaskResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetData(v *GetQualityCheckTaskResultResponseBodyData) *GetQualityCheckTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetDataType(v string) *GetQualityCheckTaskResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetErrCode(v string) *GetQualityCheckTaskResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetMessage(v string) *GetQualityCheckTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetRequestId(v string) *GetQualityCheckTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetSuccess(v bool) *GetQualityCheckTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetTime(v string) *GetQualityCheckTaskResultResponseBody {
	s.Time = &v
	return s
}

type GetQualityCheckTaskResultResponseBodyData struct {
	ConversationList *GetQualityCheckTaskResultResponseBodyDataConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Struct"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtStart         *string                                                      `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	QualityCheckList []*GetQualityCheckTaskResultResponseBodyDataQualityCheckList `json:"qualityCheckList,omitempty" xml:"qualityCheckList,omitempty" type:"Repeated"`
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1703557101831
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetConversationList(v *GetQualityCheckTaskResultResponseBodyDataConversationList) *GetQualityCheckTaskResultResponseBodyData {
	s.ConversationList = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetGmtCreate(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetGmtEnd(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.GmtEnd = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetGmtStart(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.GmtStart = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetQualityCheckList(v []*GetQualityCheckTaskResultResponseBodyDataQualityCheckList) *GetQualityCheckTaskResultResponseBodyData {
	s.QualityCheckList = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetStatus(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetTaskId(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

type GetQualityCheckTaskResultResponseBodyDataConversationList struct {
	// example:
	//
	// 1
	CallType *string `json:"callType,omitempty" xml:"callType,omitempty"`
	// example:
	//
	// 234234
	CustomerId   *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	CustomerName *string `json:"customerName,omitempty" xml:"customerName,omitempty"`
	// example:
	//
	// 23984763826
	CustomerServiceId   *string                                                                  `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	CustomerServiceName *string                                                                  `json:"customerServiceName,omitempty" xml:"customerServiceName,omitempty"`
	DialogueList        []*GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtService *string `json:"gmtService,omitempty" xml:"gmtService,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBodyDataConversationList) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyDataConversationList) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCallType(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CallType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCustomerId(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CustomerId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCustomerName(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CustomerName = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCustomerServiceId(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CustomerServiceId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCustomerServiceName(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CustomerServiceName = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetDialogueList(v []*GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.DialogueList = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetGmtService(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.GmtService = &v
	return s
}

type GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList struct {
	// example:
	//
	// 0
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// example:
	//
	// 2024-09-27 11:23:20
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// null
	CustomerId        *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// 0
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetBegin(v int32) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Begin = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetBeginTime(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.BeginTime = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetContent(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Content = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetCustomerId(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.CustomerId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetCustomerServiceId(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.CustomerServiceId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetCustomerServiceType(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.CustomerServiceType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetEnd(v int32) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.End = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetId(v int32) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Id = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetRole(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Role = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetType(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Type = &v
	return s
}

type GetQualityCheckTaskResultResponseBodyDataQualityCheckList struct {
	BizType          *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CheckExplanation *string `json:"checkExplanation,omitempty" xml:"checkExplanation,omitempty"`
	// example:
	//
	// PASSED
	CheckPassed  *string `json:"checkPassed,omitempty" xml:"checkPassed,omitempty"`
	CheckProcess *string `json:"checkProcess,omitempty" xml:"checkProcess,omitempty"`
	// example:
	//
	// HIT
	Checked *string `json:"checked,omitempty" xml:"checked,omitempty"`
	// example:
	//
	// 2024-05-23 14:57:50
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// example:
	//
	// 2024-05-23 14:57:50
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// example:
	//
	// 0
	Mode           *string                                                                    `json:"mode,omitempty" xml:"mode,omitempty"`
	OriginDialogue []*GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue `json:"originDialogue,omitempty" xml:"originDialogue,omitempty" type:"Repeated"`
	// example:
	//
	// warning_customers
	QualityGroupId  *string `json:"qualityGroupId,omitempty" xml:"qualityGroupId,omitempty"`
	RuleDescription *string `json:"ruleDescription,omitempty" xml:"ruleDescription,omitempty"`
	// example:
	//
	// wcm_start
	RuleId     *string       `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	RuleType   *string       `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	SubNodeCol []interface{} `json:"subNodeCol,omitempty" xml:"subNodeCol,omitempty" type:"Repeated"`
}

func (s GetQualityCheckTaskResultResponseBodyDataQualityCheckList) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetBizType(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.BizType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetCheckExplanation(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.CheckExplanation = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetCheckPassed(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.CheckPassed = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetCheckProcess(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.CheckProcess = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetChecked(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.Checked = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetGmtEnd(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.GmtEnd = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetGmtStart(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.GmtStart = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetMode(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.Mode = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetOriginDialogue(v []*GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.OriginDialogue = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetQualityGroupId(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.QualityGroupId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetRuleDescription(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.RuleDescription = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetRuleId(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.RuleId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetRuleType(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.RuleType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetSubNodeCol(v []interface{}) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.SubNodeCol = v
	return s
}

type GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue struct {
	// example:
	//
	// 0
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// example:
	//
	// 2024-05-23 14:57:50
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// xxx
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// 23876432
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// 0
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetBegin(v int32) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Begin = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetBeginTime(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.BeginTime = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetContent(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Content = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetCustomerId(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.CustomerId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetCustomerServiceId(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.CustomerServiceId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetCustomerServiceType(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.CustomerServiceType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetEnd(v int32) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.End = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetId(v int32) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Id = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetRole(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Role = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetType(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Type = &v
	return s
}

type GetQualityCheckTaskResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityCheckTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityCheckTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQualityCheckTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponse) SetHeaders(v map[string]*string) *GetQualityCheckTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetQualityCheckTaskResultResponse) SetStatusCode(v int32) *GetQualityCheckTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityCheckTaskResultResponse) SetBody(v *GetQualityCheckTaskResultResponseBody) *GetQualityCheckTaskResultResponse {
	s.Body = v
	return s
}

type GetSummaryTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetSummaryTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultRequest) SetTaskId(v string) *GetSummaryTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetSummaryTaskResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetSummaryTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 0bc13a9517168617617186457e401f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetSummaryTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBody) SetCost(v int64) *GetSummaryTaskResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetData(v *GetSummaryTaskResultResponseBodyData) *GetSummaryTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetDataType(v string) *GetSummaryTaskResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetErrCode(v string) *GetSummaryTaskResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetMessage(v string) *GetSummaryTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetRequestId(v string) *GetSummaryTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetSuccess(v bool) *GetSummaryTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetTime(v string) *GetSummaryTaskResultResponseBody {
	s.Time = &v
	return s
}

type GetSummaryTaskResultResponseBodyData struct {
	Choices []*GetSummaryTaskResultResponseBodyDataChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1726285125915
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 1202
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 0bc13a9517168617617186457e401f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 300
	TotalTokens *int32                                     `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	Usage       *GetSummaryTaskResultResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetSummaryTaskResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyData) SetChoices(v []*GetSummaryTaskResultResponseBodyDataChoices) *GetSummaryTaskResultResponseBodyData {
	s.Choices = v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetCreated(v int64) *GetSummaryTaskResultResponseBodyData {
	s.Created = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetId(v string) *GetSummaryTaskResultResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetModelId(v string) *GetSummaryTaskResultResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetRequestId(v string) *GetSummaryTaskResultResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetTime(v string) *GetSummaryTaskResultResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetTotalTokens(v int32) *GetSummaryTaskResultResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetUsage(v *GetSummaryTaskResultResponseBodyDataUsage) *GetSummaryTaskResultResponseBodyData {
	s.Usage = v
	return s
}

type GetSummaryTaskResultResponseBodyDataChoices struct {
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                              `json:"index,omitempty" xml:"index,omitempty"`
	Message *GetSummaryTaskResultResponseBodyDataChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s GetSummaryTaskResultResponseBodyDataChoices) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataChoices) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetFinishReason(v string) *GetSummaryTaskResultResponseBodyDataChoices {
	s.FinishReason = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetIndex(v int32) *GetSummaryTaskResultResponseBodyDataChoices {
	s.Index = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetMessage(v *GetSummaryTaskResultResponseBodyDataChoicesMessage) *GetSummaryTaskResultResponseBodyDataChoices {
	s.Message = v
	return s
}

type GetSummaryTaskResultResponseBodyDataChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role      *string                  `json:"role,omitempty" xml:"role,omitempty"`
	ToolCalls []map[string]interface{} `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s GetSummaryTaskResultResponseBodyDataChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataChoicesMessage) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetContent(v string) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.Content = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetRole(v string) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.Role = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetToolCalls(v []map[string]interface{}) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.ToolCalls = v
	return s
}

type GetSummaryTaskResultResponseBodyDataUsage struct {
	// example:
	//
	// 0
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 0
	ImageTokens *int32 `json:"imageTokens,omitempty" xml:"imageTokens,omitempty"`
	// example:
	//
	// 100
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 200
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 300
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetSummaryTaskResultResponseBodyDataUsage) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetImageCount(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.ImageCount = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetImageTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.ImageTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetInputTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetOutputTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetTotalTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

type GetSummaryTaskResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSummaryTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSummaryTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponse) SetHeaders(v map[string]*string) *GetSummaryTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetSummaryTaskResultResponse) SetStatusCode(v int32) *GetSummaryTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSummaryTaskResultResponse) SetBody(v *GetSummaryTaskResultResponseBody) *GetSummaryTaskResultResponse {
	s.Body = v
	return s
}

type GetTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetTaskResultRequest) SetTaskId(v string) *GetTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetTaskResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// {
	//
	//   "file_url": "https://finllmworks.oss-cn-zhangjiakou.aliyuncs.com/render_pdf/5336180997111160501.pdf"
	//
	// }
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
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
	// 9D5D6BB5-BEAE-53C8-A70A-7275CC1F856C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBody) SetCost(v int64) *GetTaskResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetTaskResultResponseBody) SetData(v map[string]interface{}) *GetTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResultResponseBody) SetDataType(v string) *GetTaskResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetTaskResultResponseBody) SetErrCode(v string) *GetTaskResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetTaskResultResponseBody) SetMessage(v string) *GetTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResultResponseBody) SetRequestId(v string) *GetTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResultResponseBody) SetSuccess(v bool) *GetTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskResultResponseBody) SetTime(v string) *GetTaskResultResponseBody {
	s.Time = &v
	return s
}

type GetTaskResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponse) SetHeaders(v map[string]*string) *GetTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResultResponse) SetStatusCode(v int32) *GetTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResultResponse) SetBody(v *GetTaskResultResponseBody) *GetTaskResultResponse {
	s.Body = v
	return s
}

type GetTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponseBody struct {
	Cost      *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	DataType  *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	ErrCode   *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	Time      *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetCost(v int64) *GetTaskStatusResponseBody {
	s.Cost = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetData(v string) *GetTaskStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetDataType(v string) *GetTaskStatusResponseBody {
	s.DataType = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetErrCode(v string) *GetTaskStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccess(v bool) *GetTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetTime(v string) *GetTaskStatusResponseBody {
	s.Time = &v
	return s
}

type GetTaskStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) SetHeaders(v map[string]*string) *GetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusResponse) SetStatusCode(v int32) *GetTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

type InvokePluginRequest struct {
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// 3mj87da7zr
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s InvokePluginRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokePluginRequest) GoString() string {
	return s.String()
}

func (s *InvokePluginRequest) SetParams(v map[string]interface{}) *InvokePluginRequest {
	s.Params = v
	return s
}

func (s *InvokePluginRequest) SetPluginId(v string) *InvokePluginRequest {
	s.PluginId = &v
	return s
}

type InvokePluginResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// {\\"jobWaiting\\": [0, 0], \\"timestamps\\": [1713383820, 1713383880], \\"jobUsage\\": [0, 0], \\"quotaUsage\\": [123, 32]}
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
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
	// 915AAAB9-4908-5224-9E53-9E9D7D0AA94B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s InvokePluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokePluginResponseBody) GoString() string {
	return s.String()
}

func (s *InvokePluginResponseBody) SetCost(v int64) *InvokePluginResponseBody {
	s.Cost = &v
	return s
}

func (s *InvokePluginResponseBody) SetData(v map[string]interface{}) *InvokePluginResponseBody {
	s.Data = v
	return s
}

func (s *InvokePluginResponseBody) SetDataType(v string) *InvokePluginResponseBody {
	s.DataType = &v
	return s
}

func (s *InvokePluginResponseBody) SetErrCode(v string) *InvokePluginResponseBody {
	s.ErrCode = &v
	return s
}

func (s *InvokePluginResponseBody) SetMessage(v string) *InvokePluginResponseBody {
	s.Message = &v
	return s
}

func (s *InvokePluginResponseBody) SetRequestId(v string) *InvokePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokePluginResponseBody) SetSuccess(v bool) *InvokePluginResponseBody {
	s.Success = &v
	return s
}

func (s *InvokePluginResponseBody) SetTime(v string) *InvokePluginResponseBody {
	s.Time = &v
	return s
}

type InvokePluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokePluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokePluginResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokePluginResponse) GoString() string {
	return s.String()
}

func (s *InvokePluginResponse) SetHeaders(v map[string]*string) *InvokePluginResponse {
	s.Headers = v
	return s
}

func (s *InvokePluginResponse) SetStatusCode(v int32) *InvokePluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokePluginResponse) SetBody(v *InvokePluginResponseBody) *InvokePluginResponse {
	s.Body = v
	return s
}

type PreviewDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8326472354762354
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s PreviewDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewDocumentRequest) GoString() string {
	return s.String()
}

func (s *PreviewDocumentRequest) SetDocumentId(v string) *PreviewDocumentRequest {
	s.DocumentId = &v
	return s
}

type PreviewDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *PreviewDocumentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// ff551395-1c8a-4f30-8ffd-ef7e87c70b4c
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s PreviewDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviewDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponseBody) SetCost(v int64) *PreviewDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetData(v *PreviewDocumentResponseBodyData) *PreviewDocumentResponseBody {
	s.Data = v
	return s
}

func (s *PreviewDocumentResponseBody) SetDataType(v string) *PreviewDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetErrCode(v string) *PreviewDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetMessage(v string) *PreviewDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetRequestId(v string) *PreviewDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetSuccess(v bool) *PreviewDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetTime(v string) *PreviewDocumentResponseBody {
	s.Time = &v
	return s
}

type PreviewDocumentResponseBodyData struct {
	// example:
	//
	// pdf
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	UploadTime *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	// example:
	//
	// https://agi.alicdn.com/user/d0o/d3c1f50d-a6c2-49b3-b0c8-3e613c3f20ee_16872_3236784461.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PreviewDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PreviewDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponseBodyData) SetPreviewType(v string) *PreviewDocumentResponseBodyData {
	s.PreviewType = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetTitle(v string) *PreviewDocumentResponseBodyData {
	s.Title = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetUploadTime(v string) *PreviewDocumentResponseBodyData {
	s.UploadTime = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetUrl(v string) *PreviewDocumentResponseBodyData {
	s.Url = &v
	return s
}

type PreviewDocumentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviewDocumentResponse) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponse) SetHeaders(v map[string]*string) *PreviewDocumentResponse {
	s.Headers = v
	return s
}

func (s *PreviewDocumentResponse) SetStatusCode(v int32) *PreviewDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewDocumentResponse) SetBody(v *PreviewDocumentResponseBody) *PreviewDocumentResponse {
	s.Body = v
	return s
}

type ReIndexRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8326472354762354
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s ReIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s ReIndexRequest) GoString() string {
	return s.String()
}

func (s *ReIndexRequest) SetDocumentId(v string) *ReIndexRequest {
	s.DocumentId = &v
	return s
}

type ReIndexResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// True
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// 32FFC91D-0A9F-585A-B84F-8A54C5187035
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s ReIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReIndexResponseBody) GoString() string {
	return s.String()
}

func (s *ReIndexResponseBody) SetCost(v int64) *ReIndexResponseBody {
	s.Cost = &v
	return s
}

func (s *ReIndexResponseBody) SetData(v string) *ReIndexResponseBody {
	s.Data = &v
	return s
}

func (s *ReIndexResponseBody) SetDataType(v string) *ReIndexResponseBody {
	s.DataType = &v
	return s
}

func (s *ReIndexResponseBody) SetErrCode(v string) *ReIndexResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ReIndexResponseBody) SetMessage(v string) *ReIndexResponseBody {
	s.Message = &v
	return s
}

func (s *ReIndexResponseBody) SetRequestId(v string) *ReIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReIndexResponseBody) SetSuccess(v bool) *ReIndexResponseBody {
	s.Success = &v
	return s
}

func (s *ReIndexResponseBody) SetTime(v string) *ReIndexResponseBody {
	s.Time = &v
	return s
}

type ReIndexResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s ReIndexResponse) GoString() string {
	return s.String()
}

func (s *ReIndexResponse) SetHeaders(v map[string]*string) *ReIndexResponse {
	s.Headers = v
	return s
}

func (s *ReIndexResponse) SetStatusCode(v int32) *ReIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *ReIndexResponse) SetBody(v *ReIndexResponseBody) *ReIndexResponse {
	s.Body = v
	return s
}

type RealTimeDialogRequest struct {
	// example:
	//
	// false
	Analysis *bool `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// example:
	//
	// mixIntentChat
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	ConversationModel []*RealTimeDialogRequestConversationModel `json:"conversationModel,omitempty" xml:"conversationModel,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	DialogMemoryTurns *int32                 `json:"dialogMemoryTurns,omitempty" xml:"dialogMemoryTurns,omitempty"`
	MetaData          map[string]interface{} `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// example:
	//
	// common
	OpType *string `json:"opType,omitempty" xml:"opType,omitempty"`
	// example:
	//
	// false
	Recommend           *bool   `json:"recommend,omitempty" xml:"recommend,omitempty"`
	ScriptContentPlayed *string `json:"scriptContentPlayed,omitempty" xml:"scriptContentPlayed,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// false
	Stream  *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	UserVad *bool `json:"userVad,omitempty" xml:"userVad,omitempty"`
}

func (s RealTimeDialogRequest) String() string {
	return tea.Prettify(s)
}

func (s RealTimeDialogRequest) GoString() string {
	return s.String()
}

func (s *RealTimeDialogRequest) SetAnalysis(v bool) *RealTimeDialogRequest {
	s.Analysis = &v
	return s
}

func (s *RealTimeDialogRequest) SetBizType(v string) *RealTimeDialogRequest {
	s.BizType = &v
	return s
}

func (s *RealTimeDialogRequest) SetConversationModel(v []*RealTimeDialogRequestConversationModel) *RealTimeDialogRequest {
	s.ConversationModel = v
	return s
}

func (s *RealTimeDialogRequest) SetDialogMemoryTurns(v int32) *RealTimeDialogRequest {
	s.DialogMemoryTurns = &v
	return s
}

func (s *RealTimeDialogRequest) SetMetaData(v map[string]interface{}) *RealTimeDialogRequest {
	s.MetaData = v
	return s
}

func (s *RealTimeDialogRequest) SetOpType(v string) *RealTimeDialogRequest {
	s.OpType = &v
	return s
}

func (s *RealTimeDialogRequest) SetRecommend(v bool) *RealTimeDialogRequest {
	s.Recommend = &v
	return s
}

func (s *RealTimeDialogRequest) SetScriptContentPlayed(v string) *RealTimeDialogRequest {
	s.ScriptContentPlayed = &v
	return s
}

func (s *RealTimeDialogRequest) SetSessionId(v string) *RealTimeDialogRequest {
	s.SessionId = &v
	return s
}

func (s *RealTimeDialogRequest) SetStream(v bool) *RealTimeDialogRequest {
	s.Stream = &v
	return s
}

func (s *RealTimeDialogRequest) SetUserVad(v bool) *RealTimeDialogRequest {
	s.UserVad = &v
	return s
}

type RealTimeDialogRequestConversationModel struct {
	// example:
	//
	// 5
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// example:
	//
	// 2024-11-08 09:51:16
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 98457834685635
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// 1374683645635
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// 10
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Role *int32 `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// audio
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RealTimeDialogRequestConversationModel) String() string {
	return tea.Prettify(s)
}

func (s RealTimeDialogRequestConversationModel) GoString() string {
	return s.String()
}

func (s *RealTimeDialogRequestConversationModel) SetBegin(v int32) *RealTimeDialogRequestConversationModel {
	s.Begin = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetBeginTime(v string) *RealTimeDialogRequestConversationModel {
	s.BeginTime = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetContent(v string) *RealTimeDialogRequestConversationModel {
	s.Content = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetCustomerId(v string) *RealTimeDialogRequestConversationModel {
	s.CustomerId = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetCustomerServiceId(v string) *RealTimeDialogRequestConversationModel {
	s.CustomerServiceId = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetCustomerServiceType(v string) *RealTimeDialogRequestConversationModel {
	s.CustomerServiceType = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetEnd(v int32) *RealTimeDialogRequestConversationModel {
	s.End = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetRole(v int32) *RealTimeDialogRequestConversationModel {
	s.Role = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetType(v string) *RealTimeDialogRequestConversationModel {
	s.Type = &v
	return s
}

type RealTimeDialogResponseBody struct {
	Choices []*RealTimeDialogResponseBodyChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1735139569523
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RealTimeDialogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RealTimeDialogResponseBody) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponseBody) SetChoices(v []*RealTimeDialogResponseBodyChoices) *RealTimeDialogResponseBody {
	s.Choices = v
	return s
}

func (s *RealTimeDialogResponseBody) SetCreated(v string) *RealTimeDialogResponseBody {
	s.Created = &v
	return s
}

func (s *RealTimeDialogResponseBody) SetId(v string) *RealTimeDialogResponseBody {
	s.Id = &v
	return s
}

func (s *RealTimeDialogResponseBody) SetRequestId(v string) *RealTimeDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *RealTimeDialogResponseBody) SetSuccess(v bool) *RealTimeDialogResponseBody {
	s.Success = &v
	return s
}

type RealTimeDialogResponseBodyChoices struct {
	Delta *RealTimeDialogResponseBodyChoicesDelta `json:"delta,omitempty" xml:"delta,omitempty" type:"Struct"`
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                    `json:"index,omitempty" xml:"index,omitempty"`
	Message *RealTimeDialogResponseBodyChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s RealTimeDialogResponseBodyChoices) String() string {
	return tea.Prettify(s)
}

func (s RealTimeDialogResponseBodyChoices) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponseBodyChoices) SetDelta(v *RealTimeDialogResponseBodyChoicesDelta) *RealTimeDialogResponseBodyChoices {
	s.Delta = v
	return s
}

func (s *RealTimeDialogResponseBodyChoices) SetFinishReason(v string) *RealTimeDialogResponseBodyChoices {
	s.FinishReason = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoices) SetIndex(v int32) *RealTimeDialogResponseBodyChoices {
	s.Index = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoices) SetMessage(v *RealTimeDialogResponseBodyChoicesMessage) *RealTimeDialogResponseBodyChoices {
	s.Message = v
	return s
}

type RealTimeDialogResponseBodyChoicesDelta struct {
	// example:
	//
	// null
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// time
	//
	// example:
	//
	// null
	CallTime *string `json:"callTime,omitempty" xml:"callTime,omitempty"`
	// example:
	//
	// false
	HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
	// example:
	//
	// 1853360771162058752
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionName   *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
	Interrupt       *bool   `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
	// example:
	//
	// null
	RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
	// example:
	//
	// null
	RecommendScript               *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
	SelfDirectedScript            *string `json:"selfDirectedScript,omitempty" xml:"selfDirectedScript,omitempty"`
	SelfDirectedScriptFullContent *string `json:"selfDirectedScriptFullContent,omitempty" xml:"selfDirectedScriptFullContent,omitempty"`
}

func (s RealTimeDialogResponseBodyChoicesDelta) String() string {
	return tea.Prettify(s)
}

func (s RealTimeDialogResponseBodyChoicesDelta) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetAnalysisProcess(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.AnalysisProcess = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetCallTime(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.CallTime = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetHangUpDialog(v bool) *RealTimeDialogResponseBodyChoicesDelta {
	s.HangUpDialog = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetIntentionCode(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.IntentionCode = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetIntentionName(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.IntentionName = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetIntentionScript(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.IntentionScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetInterrupt(v bool) *RealTimeDialogResponseBodyChoicesDelta {
	s.Interrupt = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetRecommendIntention(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.RecommendIntention = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetRecommendScript(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.RecommendScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetSelfDirectedScript(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.SelfDirectedScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesDelta) SetSelfDirectedScriptFullContent(v string) *RealTimeDialogResponseBodyChoicesDelta {
	s.SelfDirectedScriptFullContent = &v
	return s
}

type RealTimeDialogResponseBodyChoicesMessage struct {
	// example:
	//
	// null
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// time
	//
	// example:
	//
	// 1735139569523
	CallTime *string `json:"callTime,omitempty" xml:"callTime,omitempty"`
	// example:
	//
	// false
	HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
	// example:
	//
	// 1853360771162058752
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionName   *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
	Interrupt       *bool   `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
	// example:
	//
	// null
	RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
	// example:
	//
	// null
	RecommendScript *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
	// example:
	//
	// null
	SelfDirectedScript            *string `json:"selfDirectedScript,omitempty" xml:"selfDirectedScript,omitempty"`
	SelfDirectedScriptFullContent *string `json:"selfDirectedScriptFullContent,omitempty" xml:"selfDirectedScriptFullContent,omitempty"`
}

func (s RealTimeDialogResponseBodyChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s RealTimeDialogResponseBodyChoicesMessage) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetAnalysisProcess(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.AnalysisProcess = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetCallTime(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.CallTime = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetHangUpDialog(v bool) *RealTimeDialogResponseBodyChoicesMessage {
	s.HangUpDialog = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetIntentionCode(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.IntentionCode = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetIntentionName(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.IntentionName = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetIntentionScript(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.IntentionScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetInterrupt(v bool) *RealTimeDialogResponseBodyChoicesMessage {
	s.Interrupt = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetRecommendIntention(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.RecommendIntention = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetRecommendScript(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.RecommendScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetSelfDirectedScript(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.SelfDirectedScript = &v
	return s
}

func (s *RealTimeDialogResponseBodyChoicesMessage) SetSelfDirectedScriptFullContent(v string) *RealTimeDialogResponseBodyChoicesMessage {
	s.SelfDirectedScriptFullContent = &v
	return s
}

type RealTimeDialogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RealTimeDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RealTimeDialogResponse) String() string {
	return tea.Prettify(s)
}

func (s RealTimeDialogResponse) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponse) SetHeaders(v map[string]*string) *RealTimeDialogResponse {
	s.Headers = v
	return s
}

func (s *RealTimeDialogResponse) SetStatusCode(v int32) *RealTimeDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *RealTimeDialogResponse) SetBody(v *RealTimeDialogResponseBody) *RealTimeDialogResponse {
	s.Body = v
	return s
}

type RebuildTaskRequest struct {
	// This parameter is required.
	TaskIds []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s RebuildTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildTaskRequest) GoString() string {
	return s.String()
}

func (s *RebuildTaskRequest) SetTaskIds(v []*string) *RebuildTaskRequest {
	s.TaskIds = v
	return s
}

type RebuildTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64                   `json:"cost,omitempty" xml:"cost,omitempty"`
	Data []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RebuildTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildTaskResponseBody) SetCost(v int64) *RebuildTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *RebuildTaskResponseBody) SetData(v []map[string]interface{}) *RebuildTaskResponseBody {
	s.Data = v
	return s
}

func (s *RebuildTaskResponseBody) SetDataType(v string) *RebuildTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *RebuildTaskResponseBody) SetErrCode(v string) *RebuildTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RebuildTaskResponseBody) SetMessage(v string) *RebuildTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RebuildTaskResponseBody) SetRequestId(v string) *RebuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildTaskResponseBody) SetSuccess(v bool) *RebuildTaskResponseBody {
	s.Success = &v
	return s
}

func (s *RebuildTaskResponseBody) SetTime(v string) *RebuildTaskResponseBody {
	s.Time = &v
	return s
}

type RebuildTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebuildTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildTaskResponse) GoString() string {
	return s.String()
}

func (s *RebuildTaskResponse) SetHeaders(v map[string]*string) *RebuildTaskResponse {
	s.Headers = v
	return s
}

func (s *RebuildTaskResponse) SetStatusCode(v int32) *RebuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildTaskResponse) SetBody(v *RebuildTaskResponseBody) *RebuildTaskResponse {
	s.Body = v
	return s
}

type RecallDocumentRequest struct {
	Filters []*RecallDocumentRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// false
	Rearrangement *bool `json:"rearrangement,omitempty" xml:"rearrangement,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s RecallDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentRequest) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequest) SetFilters(v []*RecallDocumentRequestFilters) *RecallDocumentRequest {
	s.Filters = v
	return s
}

func (s *RecallDocumentRequest) SetQuery(v string) *RecallDocumentRequest {
	s.Query = &v
	return s
}

func (s *RecallDocumentRequest) SetRearrangement(v bool) *RecallDocumentRequest {
	s.Rearrangement = &v
	return s
}

func (s *RecallDocumentRequest) SetTopK(v int32) *RecallDocumentRequest {
	s.TopK = &v
	return s
}

type RecallDocumentRequestFilters struct {
	And []*RecallDocumentRequestFiltersAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	// example:
	//
	// Text
	ChunkType *string   `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	DocIdList []*string `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sdbjhvs
	LibraryId *string                           `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	Or        []*RecallDocumentRequestFiltersOr `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
	Status    []*string                         `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
}

func (s RecallDocumentRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentRequestFilters) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFilters) SetAnd(v []*RecallDocumentRequestFiltersAnd) *RecallDocumentRequestFilters {
	s.And = v
	return s
}

func (s *RecallDocumentRequestFilters) SetChunkType(v string) *RecallDocumentRequestFilters {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentRequestFilters) SetDocIdList(v []*string) *RecallDocumentRequestFilters {
	s.DocIdList = v
	return s
}

func (s *RecallDocumentRequestFilters) SetLibraryId(v string) *RecallDocumentRequestFilters {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentRequestFilters) SetOr(v []*RecallDocumentRequestFiltersOr) *RecallDocumentRequestFilters {
	s.Or = v
	return s
}

func (s *RecallDocumentRequestFilters) SetStatus(v []*string) *RecallDocumentRequestFilters {
	s.Status = v
	return s
}

type RecallDocumentRequestFiltersAnd struct {
	// example:
	//
	// 20
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// docType
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// contains
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RecallDocumentRequestFiltersAnd) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentRequestFiltersAnd) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFiltersAnd) SetBoost(v float32) *RecallDocumentRequestFiltersAnd {
	s.Boost = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetKey(v string) *RecallDocumentRequestFiltersAnd {
	s.Key = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetOperator(v string) *RecallDocumentRequestFiltersAnd {
	s.Operator = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetValue(v string) *RecallDocumentRequestFiltersAnd {
	s.Value = &v
	return s
}

type RecallDocumentRequestFiltersOr struct {
	// example:
	//
	// 30
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// researcher
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RecallDocumentRequestFiltersOr) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentRequestFiltersOr) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFiltersOr) SetBoost(v float32) *RecallDocumentRequestFiltersOr {
	s.Boost = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetKey(v string) *RecallDocumentRequestFiltersOr {
	s.Key = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetOperator(v string) *RecallDocumentRequestFiltersOr {
	s.Operator = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetValue(v string) *RecallDocumentRequestFiltersOr {
	s.Value = &v
	return s
}

type RecallDocumentResponseBody struct {
	// example:
	//
	// 0
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RecallDocumentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 0bc13a9517168617617186457e401f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RecallDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBody) SetCost(v int64) *RecallDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *RecallDocumentResponseBody) SetData(v *RecallDocumentResponseBodyData) *RecallDocumentResponseBody {
	s.Data = v
	return s
}

func (s *RecallDocumentResponseBody) SetDataType(v string) *RecallDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *RecallDocumentResponseBody) SetErrCode(v string) *RecallDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RecallDocumentResponseBody) SetMessage(v string) *RecallDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *RecallDocumentResponseBody) SetRequestId(v string) *RecallDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecallDocumentResponseBody) SetSuccess(v bool) *RecallDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *RecallDocumentResponseBody) SetTime(v string) *RecallDocumentResponseBody {
	s.Time = &v
	return s
}

type RecallDocumentResponseBodyData struct {
	ChunkList     []*RecallDocumentResponseBodyDataChunkList     `json:"chunkList,omitempty" xml:"chunkList,omitempty" type:"Repeated"`
	ChunkPartList []*RecallDocumentResponseBodyDataChunkPartList `json:"chunkPartList,omitempty" xml:"chunkPartList,omitempty" type:"Repeated"`
	ChunkTextList []*string                                      `json:"chunkTextList,omitempty" xml:"chunkTextList,omitempty" type:"Repeated"`
	Documents     []*RecallDocumentResponseBodyDataDocuments     `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	EmbeddingElapsedMs *int64                                         `json:"embeddingElapsedMs,omitempty" xml:"embeddingElapsedMs,omitempty"`
	TextChunkList      []*RecallDocumentResponseBodyDataTextChunkList `json:"textChunkList,omitempty" xml:"textChunkList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TextSearchElapsedMs *int64 `json:"textSearchElapsedMs,omitempty" xml:"textSearchElapsedMs,omitempty"`
	// example:
	//
	// 400
	TotalElapsedMs  *int64                                           `json:"totalElapsedMs,omitempty" xml:"totalElapsedMs,omitempty"`
	VectorChunkList []*RecallDocumentResponseBodyDataVectorChunkList `json:"vectorChunkList,omitempty" xml:"vectorChunkList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	VectorSearchElapsedMs *int64 `json:"vectorSearchElapsedMs,omitempty" xml:"vectorSearchElapsedMs,omitempty"`
}

func (s RecallDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyData) SetChunkList(v []*RecallDocumentResponseBodyDataChunkList) *RecallDocumentResponseBodyData {
	s.ChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetChunkPartList(v []*RecallDocumentResponseBodyDataChunkPartList) *RecallDocumentResponseBodyData {
	s.ChunkPartList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetChunkTextList(v []*string) *RecallDocumentResponseBodyData {
	s.ChunkTextList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetDocuments(v []*RecallDocumentResponseBodyDataDocuments) *RecallDocumentResponseBodyData {
	s.Documents = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetEmbeddingElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.EmbeddingElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTextChunkList(v []*RecallDocumentResponseBodyDataTextChunkList) *RecallDocumentResponseBodyData {
	s.TextChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTextSearchElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.TextSearchElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTotalElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.TotalElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetVectorChunkList(v []*RecallDocumentResponseBodyDataVectorChunkList) *RecallDocumentResponseBodyData {
	s.VectorChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetVectorSearchElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.VectorSearchElapsedMs = &v
	return s
}

type RecallDocumentResponseBodyDataChunkList struct {
	// example:
	//
	// 823746762354
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 839468263472
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// dscsbdsk
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 982374872364
	NextChunkId *string                                       `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 827364827364832
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataChunkList) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetPos(v []*RecallDocumentResponseBodyDataChunkListPos) *RecallDocumentResponseBodyDataChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataChunkList {
	s.Title = &v
	return s
}

type RecallDocumentResponseBodyDataChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataChunkListPos) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataChunkListPos {
	s.TextHighlightArea = v
	return s
}

type RecallDocumentResponseBodyDataChunkPartList struct {
	// example:
	//
	// 98327482364
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 92837482364
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// sjdhgjsd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 2387648263542
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataChunkPartListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 32874682764
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	Title *string  `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataChunkPartList) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkPartList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkText(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkType(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetDocId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetFileType(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetLibraryId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetLibraryName(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetPos(v []*RecallDocumentResponseBodyDataChunkPartListPos) *RecallDocumentResponseBodyDataChunkPartList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetScore(v float32) *RecallDocumentResponseBodyDataChunkPartList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetTitle(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.Title = &v
	return s
}

type RecallDocumentResponseBodyDataChunkPartListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataChunkPartListPos) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkPartListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetPage(v int32) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.TextHighlightArea = v
	return s
}

type RecallDocumentResponseBodyDataDocuments struct {
	// example:
	//
	// 92837482364
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// {"a":"1"}
	DocumentMeta map[string]interface{} `json:"documentMeta,omitempty" xml:"documentMeta,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// sjdhgjsd
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/test.pdf
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RecallDocumentResponseBodyDataDocuments) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataDocuments) SetDocId(v string) *RecallDocumentResponseBodyDataDocuments {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetDocumentMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataDocuments {
	s.DocumentMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetFileType(v string) *RecallDocumentResponseBodyDataDocuments {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetGmtCreate(v string) *RecallDocumentResponseBodyDataDocuments {
	s.GmtCreate = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetLibraryId(v string) *RecallDocumentResponseBodyDataDocuments {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetTitle(v string) *RecallDocumentResponseBodyDataDocuments {
	s.Title = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetUrl(v string) *RecallDocumentResponseBodyDataDocuments {
	s.Url = &v
	return s
}

type RecallDocumentResponseBodyDataTextChunkList struct {
	// example:
	//
	// 32874682364
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 8372467263542
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// djsgfsjd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 23874682432
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataTextChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 89473868346
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	Title *string  `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataTextChunkList) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataTextChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetPos(v []*RecallDocumentResponseBodyDataTextChunkListPos) *RecallDocumentResponseBodyDataTextChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataTextChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.Title = &v
	return s
}

type RecallDocumentResponseBodyDataTextChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataTextChunkListPos) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataTextChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.TextHighlightArea = v
	return s
}

type RecallDocumentResponseBodyDataVectorChunkList struct {
	// example:
	//
	// 8723642345276
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// https://oss-xxxx-hangzhou.com/test.pdf
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 78326476235675372
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// djsgfsjd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 293846872343
	NextChunkId *string                                             `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataVectorChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 873647326542
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataVectorChunkList) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataVectorChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetPos(v []*RecallDocumentResponseBodyDataVectorChunkListPos) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Title = &v
	return s
}

type RecallDocumentResponseBodyDataVectorChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataVectorChunkListPos) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataVectorChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.TextHighlightArea = v
	return s
}

type RecallDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecallDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecallDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponse) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponse) SetHeaders(v map[string]*string) *RecallDocumentResponse {
	s.Headers = v
	return s
}

func (s *RecallDocumentResponse) SetStatusCode(v int32) *RecallDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *RecallDocumentResponse) SetBody(v *RecallDocumentResponseBody) *RecallDocumentResponse {
	s.Body = v
	return s
}

type RecognizeIntentionRequest struct {
	// example:
	//
	// false
	Analysis *bool `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// common
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	Conversation              *string                                               `json:"conversation,omitempty" xml:"conversation,omitempty"`
	GlobalIntentionList       []*RecognizeIntentionRequestGlobalIntentionList       `json:"globalIntentionList,omitempty" xml:"globalIntentionList,omitempty" type:"Repeated"`
	HierarchicalIntentionList []*RecognizeIntentionRequestHierarchicalIntentionList `json:"hierarchicalIntentionList,omitempty" xml:"hierarchicalIntentionList,omitempty" type:"Repeated"`
	IntentionDomainCode       *string                                               `json:"intentionDomainCode,omitempty" xml:"intentionDomainCode,omitempty"`
	IntentionList             []*RecognizeIntentionRequestIntentionList             `json:"intentionList,omitempty" xml:"intentionList,omitempty" type:"Repeated"`
	// example:
	//
	// common
	OpType *string `json:"opType,omitempty" xml:"opType,omitempty"`
	// example:
	//
	// false
	Recommend *bool `json:"recommend,omitempty" xml:"recommend,omitempty"`
}

func (s RecognizeIntentionRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequest) SetAnalysis(v bool) *RecognizeIntentionRequest {
	s.Analysis = &v
	return s
}

func (s *RecognizeIntentionRequest) SetBizType(v string) *RecognizeIntentionRequest {
	s.BizType = &v
	return s
}

func (s *RecognizeIntentionRequest) SetConversation(v string) *RecognizeIntentionRequest {
	s.Conversation = &v
	return s
}

func (s *RecognizeIntentionRequest) SetGlobalIntentionList(v []*RecognizeIntentionRequestGlobalIntentionList) *RecognizeIntentionRequest {
	s.GlobalIntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetHierarchicalIntentionList(v []*RecognizeIntentionRequestHierarchicalIntentionList) *RecognizeIntentionRequest {
	s.HierarchicalIntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetIntentionDomainCode(v string) *RecognizeIntentionRequest {
	s.IntentionDomainCode = &v
	return s
}

func (s *RecognizeIntentionRequest) SetIntentionList(v []*RecognizeIntentionRequestIntentionList) *RecognizeIntentionRequest {
	s.IntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetOpType(v string) *RecognizeIntentionRequest {
	s.OpType = &v
	return s
}

func (s *RecognizeIntentionRequest) SetRecommend(v bool) *RecognizeIntentionRequest {
	s.Recommend = &v
	return s
}

type RecognizeIntentionRequestGlobalIntentionList struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1810566978021232640
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
}

func (s RecognizeIntentionRequestGlobalIntentionList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionRequestGlobalIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetDescription(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetIntention(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetIntentionScript(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.IntentionScript = &v
	return s
}

type RecognizeIntentionRequestHierarchicalIntentionList struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1810929291010150400
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
}

func (s RecognizeIntentionRequestHierarchicalIntentionList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionRequestHierarchicalIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetDescription(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetIntention(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetIntentionScript(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.IntentionScript = &v
	return s
}

type RecognizeIntentionRequestIntentionList struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1808766224000262144
	IntentionCode   *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionScript *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
}

func (s RecognizeIntentionRequestIntentionList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionRequestIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestIntentionList) SetDescription(v string) *RecognizeIntentionRequestIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) SetIntention(v string) *RecognizeIntentionRequestIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestIntentionList {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) SetIntentionScript(v string) *RecognizeIntentionRequestIntentionList {
	s.IntentionScript = &v
	return s
}

type RecognizeIntentionResponseBody struct {
	// example:
	//
	// null
	Cost *int64                              `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RecognizeIntentionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 003D019A-1BB3-53EC-A0D2-CE76DA5D73B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RecognizeIntentionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponseBody) SetCost(v int64) *RecognizeIntentionResponseBody {
	s.Cost = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetData(v *RecognizeIntentionResponseBodyData) *RecognizeIntentionResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeIntentionResponseBody) SetDataType(v string) *RecognizeIntentionResponseBody {
	s.DataType = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetErrCode(v string) *RecognizeIntentionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetMessage(v string) *RecognizeIntentionResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetRequestId(v string) *RecognizeIntentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetSuccess(v bool) *RecognizeIntentionResponseBody {
	s.Success = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetTime(v string) *RecognizeIntentionResponseBody {
	s.Time = &v
	return s
}

type RecognizeIntentionResponseBodyData struct {
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// example:
	//
	// 1
	IntentionCode      *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionName      *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	IntentionScript    *string `json:"intentionScript,omitempty" xml:"intentionScript,omitempty"`
	RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
	RecommendScript    *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
}

func (s RecognizeIntentionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponseBodyData) SetAnalysisProcess(v string) *RecognizeIntentionResponseBodyData {
	s.AnalysisProcess = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetIntentionCode(v string) *RecognizeIntentionResponseBodyData {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetIntentionName(v string) *RecognizeIntentionResponseBodyData {
	s.IntentionName = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetIntentionScript(v string) *RecognizeIntentionResponseBodyData {
	s.IntentionScript = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetRecommendIntention(v string) *RecognizeIntentionResponseBodyData {
	s.RecommendIntention = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetRecommendScript(v string) *RecognizeIntentionResponseBodyData {
	s.RecommendScript = &v
	return s
}

type RecognizeIntentionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeIntentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeIntentionResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponse) SetHeaders(v map[string]*string) *RecognizeIntentionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeIntentionResponse) SetStatusCode(v int32) *RecognizeIntentionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeIntentionResponse) SetBody(v *RecognizeIntentionResponseBody) *RecognizeIntentionResponse {
	s.Body = v
	return s
}

type RunAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d6zxykawk9
	BotId *string `json:"botId,omitempty" xml:"botId,omitempty"`
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// example:
	//
	// 4vlag5ken3
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// false
	UseDraft *bool `json:"useDraft,omitempty" xml:"useDraft,omitempty"`
	// This parameter is required.
	UserContent *string `json:"userContent,omitempty" xml:"userContent,omitempty"`
	// example:
	//
	// w4paqoezm2
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s RunAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s RunAgentRequest) GoString() string {
	return s.String()
}

func (s *RunAgentRequest) SetBotId(v string) *RunAgentRequest {
	s.BotId = &v
	return s
}

func (s *RunAgentRequest) SetModelId(v string) *RunAgentRequest {
	s.ModelId = &v
	return s
}

func (s *RunAgentRequest) SetStream(v bool) *RunAgentRequest {
	s.Stream = &v
	return s
}

func (s *RunAgentRequest) SetThreadId(v string) *RunAgentRequest {
	s.ThreadId = &v
	return s
}

func (s *RunAgentRequest) SetUseDraft(v bool) *RunAgentRequest {
	s.UseDraft = &v
	return s
}

func (s *RunAgentRequest) SetUserContent(v string) *RunAgentRequest {
	s.UserContent = &v
	return s
}

func (s *RunAgentRequest) SetVersionId(v string) *RunAgentRequest {
	s.VersionId = &v
	return s
}

type RunAgentResponseBody struct {
	// example:
	//
	// null
	Cost *int64                    `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RunAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
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

func (s RunAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunAgentResponseBody) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBody) SetCost(v int64) *RunAgentResponseBody {
	s.Cost = &v
	return s
}

func (s *RunAgentResponseBody) SetData(v *RunAgentResponseBodyData) *RunAgentResponseBody {
	s.Data = v
	return s
}

func (s *RunAgentResponseBody) SetDataType(v string) *RunAgentResponseBody {
	s.DataType = &v
	return s
}

func (s *RunAgentResponseBody) SetErrCode(v string) *RunAgentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RunAgentResponseBody) SetMessage(v string) *RunAgentResponseBody {
	s.Message = &v
	return s
}

func (s *RunAgentResponseBody) SetRequestId(v string) *RunAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunAgentResponseBody) SetSuccess(v bool) *RunAgentResponseBody {
	s.Success = &v
	return s
}

func (s *RunAgentResponseBody) SetTime(v string) *RunAgentResponseBody {
	s.Time = &v
	return s
}

type RunAgentResponseBodyData struct {
	FunctionCallResponses []*RunAgentResponseBodyDataFunctionCallResponses `json:"functionCallResponses,omitempty" xml:"functionCallResponses,omitempty" type:"Repeated"`
	// example:
	//
	// 766
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 988
	OutputTokens *int32                            `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	Response     *RunAgentResponseBodyDataResponse `json:"response,omitempty" xml:"response,omitempty" type:"Struct"`
	// example:
	//
	// 4vlag5ken3
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// 5bdb9809856c58acb92001f8ae65773c
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
	// example:
	//
	// w4paqoezm2
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s RunAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyData) SetFunctionCallResponses(v []*RunAgentResponseBodyDataFunctionCallResponses) *RunAgentResponseBodyData {
	s.FunctionCallResponses = v
	return s
}

func (s *RunAgentResponseBodyData) SetInputTokens(v int32) *RunAgentResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *RunAgentResponseBodyData) SetOutputTokens(v int32) *RunAgentResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *RunAgentResponseBodyData) SetResponse(v *RunAgentResponseBodyDataResponse) *RunAgentResponseBodyData {
	s.Response = v
	return s
}

func (s *RunAgentResponseBodyData) SetThreadId(v string) *RunAgentResponseBodyData {
	s.ThreadId = &v
	return s
}

func (s *RunAgentResponseBodyData) SetTraceId(v string) *RunAgentResponseBodyData {
	s.TraceId = &v
	return s
}

func (s *RunAgentResponseBodyData) SetVersionId(v string) *RunAgentResponseBodyData {
	s.VersionId = &v
	return s
}

type RunAgentResponseBodyDataFunctionCallResponses struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 2025-01-21 16:37:14
	EndTime      *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FunctionArgs *string `json:"functionArgs,omitempty" xml:"functionArgs,omitempty"`
	// example:
	//
	// web_search
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	Result       *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 2025-01-21 16:37:14
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s RunAgentResponseBodyDataFunctionCallResponses) String() string {
	return tea.Prettify(s)
}

func (s RunAgentResponseBodyDataFunctionCallResponses) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetDisplayName(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.DisplayName = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetEndTime(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.EndTime = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetFunctionArgs(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.FunctionArgs = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetFunctionName(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.FunctionName = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetResult(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.Result = &v
	return s
}

func (s *RunAgentResponseBodyDataFunctionCallResponses) SetStartTime(v string) *RunAgentResponseBodyDataFunctionCallResponses {
	s.StartTime = &v
	return s
}

type RunAgentResponseBodyDataResponse struct {
	Choices []*RunAgentResponseBodyDataResponseChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1737448637
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// d91d9afa-7cfc-4235-b012-a6f8e6ffa443
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 2025-01-21T16:37:17.497206762
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RunAgentResponseBodyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s RunAgentResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyDataResponse) SetChoices(v []*RunAgentResponseBodyDataResponseChoices) *RunAgentResponseBodyDataResponse {
	s.Choices = v
	return s
}

func (s *RunAgentResponseBodyDataResponse) SetCreated(v int64) *RunAgentResponseBodyDataResponse {
	s.Created = &v
	return s
}

func (s *RunAgentResponseBodyDataResponse) SetId(v string) *RunAgentResponseBodyDataResponse {
	s.Id = &v
	return s
}

func (s *RunAgentResponseBodyDataResponse) SetModelId(v string) *RunAgentResponseBodyDataResponse {
	s.ModelId = &v
	return s
}

func (s *RunAgentResponseBodyDataResponse) SetTime(v string) *RunAgentResponseBodyDataResponse {
	s.Time = &v
	return s
}

type RunAgentResponseBodyDataResponseChoices struct {
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                          `json:"index,omitempty" xml:"index,omitempty"`
	Message *RunAgentResponseBodyDataResponseChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s RunAgentResponseBodyDataResponseChoices) String() string {
	return tea.Prettify(s)
}

func (s RunAgentResponseBodyDataResponseChoices) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyDataResponseChoices) SetFinishReason(v string) *RunAgentResponseBodyDataResponseChoices {
	s.FinishReason = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoices) SetIndex(v int32) *RunAgentResponseBodyDataResponseChoices {
	s.Index = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoices) SetMessage(v *RunAgentResponseBodyDataResponseChoicesMessage) *RunAgentResponseBodyDataResponseChoices {
	s.Message = v
	return s
}

type RunAgentResponseBodyDataResponseChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// assistant
	RoleDisplayName *string `json:"roleDisplayName,omitempty" xml:"roleDisplayName,omitempty"`
}

func (s RunAgentResponseBodyDataResponseChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s RunAgentResponseBodyDataResponseChoicesMessage) GoString() string {
	return s.String()
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) SetContent(v string) *RunAgentResponseBodyDataResponseChoicesMessage {
	s.Content = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) SetRole(v string) *RunAgentResponseBodyDataResponseChoicesMessage {
	s.Role = &v
	return s
}

func (s *RunAgentResponseBodyDataResponseChoicesMessage) SetRoleDisplayName(v string) *RunAgentResponseBodyDataResponseChoicesMessage {
	s.RoleDisplayName = &v
	return s
}

type RunAgentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s RunAgentResponse) GoString() string {
	return s.String()
}

func (s *RunAgentResponse) SetHeaders(v map[string]*string) *RunAgentResponse {
	s.Headers = v
	return s
}

func (s *RunAgentResponse) SetStatusCode(v int32) *RunAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *RunAgentResponse) SetBody(v *RunAgentResponseBody) *RunAgentResponse {
	s.Body = v
	return s
}

type RunChatResultGenerationRequest struct {
	// example:
	//
	// {"topP": 0.8}
	InferenceParameters map[string]interface{} `json:"inferenceParameters,omitempty" xml:"inferenceParameters,omitempty"`
	// This parameter is required.
	Messages []*RunChatResultGenerationRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// false
	Stream *bool                                  `json:"stream,omitempty" xml:"stream,omitempty"`
	Tools  []*RunChatResultGenerationRequestTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequest) SetInferenceParameters(v map[string]interface{}) *RunChatResultGenerationRequest {
	s.InferenceParameters = v
	return s
}

func (s *RunChatResultGenerationRequest) SetMessages(v []*RunChatResultGenerationRequestMessages) *RunChatResultGenerationRequest {
	s.Messages = v
	return s
}

func (s *RunChatResultGenerationRequest) SetModelId(v string) *RunChatResultGenerationRequest {
	s.ModelId = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetSessionId(v string) *RunChatResultGenerationRequest {
	s.SessionId = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetStream(v bool) *RunChatResultGenerationRequest {
	s.Stream = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetTools(v []*RunChatResultGenerationRequestTools) *RunChatResultGenerationRequest {
	s.Tools = v
	return s
}

type RunChatResultGenerationRequestMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunChatResultGenerationRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequestMessages) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestMessages) SetContent(v string) *RunChatResultGenerationRequestMessages {
	s.Content = &v
	return s
}

func (s *RunChatResultGenerationRequestMessages) SetRole(v string) *RunChatResultGenerationRequestMessages {
	s.Role = &v
	return s
}

type RunChatResultGenerationRequestTools struct {
	Function *RunChatResultGenerationRequestToolsFunction `json:"function,omitempty" xml:"function,omitempty" type:"Struct"`
	// example:
	//
	// function
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RunChatResultGenerationRequestTools) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequestTools) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestTools) SetFunction(v *RunChatResultGenerationRequestToolsFunction) *RunChatResultGenerationRequestTools {
	s.Function = v
	return s
}

func (s *RunChatResultGenerationRequestTools) SetType(v string) *RunChatResultGenerationRequestTools {
	s.Type = &v
	return s
}

type RunChatResultGenerationRequestToolsFunction struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// get_time
	Name       *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	Parameters *RunChatResultGenerationRequestToolsFunctionParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	Required   []*string                                              `json:"required,omitempty" xml:"required,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationRequestToolsFunction) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequestToolsFunction) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestToolsFunction) SetDescription(v string) *RunChatResultGenerationRequestToolsFunction {
	s.Description = &v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetName(v string) *RunChatResultGenerationRequestToolsFunction {
	s.Name = &v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetParameters(v *RunChatResultGenerationRequestToolsFunctionParameters) *RunChatResultGenerationRequestToolsFunction {
	s.Parameters = v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetRequired(v []*string) *RunChatResultGenerationRequestToolsFunction {
	s.Required = v
	return s
}

type RunChatResultGenerationRequestToolsFunctionParameters struct {
	// example:
	//
	// {
	//
	//                             "location": {
	//
	//                                 "type": "string",
	//
	//                                 "description": "The city and state, e.g. San Francisco, CA"
	//
	//                             },
	//
	//                             "unit": {
	//
	//                                 "type": "string",
	//
	//                                 "enum": [
	//
	//                                     "celsius",
	//
	//                                     "fahrenheit"
	//
	//                                 ]
	//
	//                             }
	//
	//                         }
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// example:
	//
	// object
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RunChatResultGenerationRequestToolsFunctionParameters) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequestToolsFunctionParameters) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) SetProperties(v map[string]interface{}) *RunChatResultGenerationRequestToolsFunctionParameters {
	s.Properties = v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) SetType(v string) *RunChatResultGenerationRequestToolsFunctionParameters {
	s.Type = &v
	return s
}

type RunChatResultGenerationResponseBody struct {
	Choices []*RunChatResultGenerationResponseBodyChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1720602203
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 500
	TotalTokens *int32                                    `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	Usage       *RunChatResultGenerationResponseBodyUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunChatResultGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBody) SetChoices(v []*RunChatResultGenerationResponseBodyChoices) *RunChatResultGenerationResponseBody {
	s.Choices = v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetCreated(v int64) *RunChatResultGenerationResponseBody {
	s.Created = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetId(v string) *RunChatResultGenerationResponseBody {
	s.Id = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetModelId(v string) *RunChatResultGenerationResponseBody {
	s.ModelId = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetRequestId(v string) *RunChatResultGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetTime(v string) *RunChatResultGenerationResponseBody {
	s.Time = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetTotalTokens(v int32) *RunChatResultGenerationResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetUsage(v *RunChatResultGenerationResponseBodyUsage) *RunChatResultGenerationResponseBody {
	s.Usage = v
	return s
}

type RunChatResultGenerationResponseBodyChoices struct {
	// example:
	//
	// null
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                             `json:"index,omitempty" xml:"index,omitempty"`
	Message *RunChatResultGenerationResponseBodyChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s RunChatResultGenerationResponseBodyChoices) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyChoices) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyChoices) SetFinishReason(v string) *RunChatResultGenerationResponseBodyChoices {
	s.FinishReason = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoices) SetIndex(v int32) *RunChatResultGenerationResponseBodyChoices {
	s.Index = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoices) SetMessage(v *RunChatResultGenerationResponseBodyChoicesMessage) *RunChatResultGenerationResponseBodyChoices {
	s.Message = v
	return s
}

type RunChatResultGenerationResponseBodyChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role      *string                  `json:"role,omitempty" xml:"role,omitempty"`
	ToolCalls []map[string]interface{} `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationResponseBodyChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyChoicesMessage) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetContent(v string) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.Content = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetRole(v string) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.Role = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetToolCalls(v []map[string]interface{}) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.ToolCalls = v
	return s
}

type RunChatResultGenerationResponseBodyUsage struct {
	// example:
	//
	// 0
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 0
	ImageTokens *int32 `json:"imageTokens,omitempty" xml:"imageTokens,omitempty"`
	// example:
	//
	// 200
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 300
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 500
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunChatResultGenerationResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyUsage) SetImageCount(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.ImageCount = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetImageTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.ImageTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetInputTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetOutputTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetTotalTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

type RunChatResultGenerationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunChatResultGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunChatResultGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponse) SetHeaders(v map[string]*string) *RunChatResultGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunChatResultGenerationResponse) SetStatusCode(v int32) *RunChatResultGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunChatResultGenerationResponse) SetBody(v *RunChatResultGenerationResponseBody) *RunChatResultGenerationResponse {
	s.Body = v
	return s
}

type RunLibraryChatGenerationRequest struct {
	DocIdList []*string `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// false
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// false
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// qwen-max
	FollowUpLlm *string `json:"followUpLlm,omitempty" xml:"followUpLlm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	LlmType *string `json:"llmType,omitempty" xml:"llmType,omitempty"`
	// example:
	//
	// qwen-max
	MultiQueryLlm *string `json:"multiQueryLlm,omitempty" xml:"multiQueryLlm,omitempty"`
	// This parameter is required.
	Query         *string                                       `json:"query,omitempty" xml:"query,omitempty"`
	QueryCriteria *RunLibraryChatGenerationRequestQueryCriteria `json:"queryCriteria,omitempty" xml:"queryCriteria,omitempty" type:"Struct"`
	// example:
	//
	// linear
	RerankType *string `json:"rerankType,omitempty" xml:"rerankType,omitempty"`
	// sessionId
	//
	// example:
	//
	// null
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// false
	Stream              *bool                                               `json:"stream,omitempty" xml:"stream,omitempty"`
	SubQueryList        []*string                                           `json:"subQueryList,omitempty" xml:"subQueryList,omitempty" type:"Repeated"`
	TextSearchParameter *RunLibraryChatGenerationRequestTextSearchParameter `json:"textSearchParameter,omitempty" xml:"textSearchParameter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	TopK                  *int32                                                `json:"topK,omitempty" xml:"topK,omitempty"`
	VectorSearchParameter *RunLibraryChatGenerationRequestVectorSearchParameter `json:"vectorSearchParameter,omitempty" xml:"vectorSearchParameter,omitempty" type:"Struct"`
	// example:
	//
	// false
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s RunLibraryChatGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequest) SetDocIdList(v []*string) *RunLibraryChatGenerationRequest {
	s.DocIdList = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableFollowUp(v bool) *RunLibraryChatGenerationRequest {
	s.EnableFollowUp = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableMultiQuery(v bool) *RunLibraryChatGenerationRequest {
	s.EnableMultiQuery = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableOpenQa(v bool) *RunLibraryChatGenerationRequest {
	s.EnableOpenQa = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetFollowUpLlm(v string) *RunLibraryChatGenerationRequest {
	s.FollowUpLlm = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetLibraryId(v string) *RunLibraryChatGenerationRequest {
	s.LibraryId = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetLlmType(v string) *RunLibraryChatGenerationRequest {
	s.LlmType = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetMultiQueryLlm(v string) *RunLibraryChatGenerationRequest {
	s.MultiQueryLlm = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetQuery(v string) *RunLibraryChatGenerationRequest {
	s.Query = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetQueryCriteria(v *RunLibraryChatGenerationRequestQueryCriteria) *RunLibraryChatGenerationRequest {
	s.QueryCriteria = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetRerankType(v string) *RunLibraryChatGenerationRequest {
	s.RerankType = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetSessionId(v string) *RunLibraryChatGenerationRequest {
	s.SessionId = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetStream(v bool) *RunLibraryChatGenerationRequest {
	s.Stream = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetSubQueryList(v []*string) *RunLibraryChatGenerationRequest {
	s.SubQueryList = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetTextSearchParameter(v *RunLibraryChatGenerationRequestTextSearchParameter) *RunLibraryChatGenerationRequest {
	s.TextSearchParameter = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetTopK(v int32) *RunLibraryChatGenerationRequest {
	s.TopK = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetVectorSearchParameter(v *RunLibraryChatGenerationRequestVectorSearchParameter) *RunLibraryChatGenerationRequest {
	s.VectorSearchParameter = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetWithDocumentReference(v bool) *RunLibraryChatGenerationRequest {
	s.WithDocumentReference = &v
	return s
}

type RunLibraryChatGenerationRequestQueryCriteria struct {
	And []*RunLibraryChatGenerationRequestQueryCriteriaAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	Or  []*RunLibraryChatGenerationRequestQueryCriteriaOr  `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
}

func (s RunLibraryChatGenerationRequestQueryCriteria) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteria) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) SetAnd(v []*RunLibraryChatGenerationRequestQueryCriteriaAnd) *RunLibraryChatGenerationRequestQueryCriteria {
	s.And = v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) SetOr(v []*RunLibraryChatGenerationRequestQueryCriteriaOr) *RunLibraryChatGenerationRequestQueryCriteria {
	s.Or = v
	return s
}

type RunLibraryChatGenerationRequestQueryCriteriaAnd struct {
	// example:
	//
	// 0.5
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// city
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RunLibraryChatGenerationRequestQueryCriteriaAnd) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteriaAnd) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetBoost(v float32) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Boost = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetKey(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Key = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetOperator(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Operator = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetValue(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Value = &v
	return s
}

type RunLibraryChatGenerationRequestQueryCriteriaOr struct {
	// example:
	//
	// 0.5
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// city
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RunLibraryChatGenerationRequestQueryCriteriaOr) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteriaOr) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetBoost(v float32) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Boost = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetKey(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Key = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetOperator(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Operator = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetValue(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Value = &v
	return s
}

type RunLibraryChatGenerationRequestTextSearchParameter struct {
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// IkMaxWord
	SearchAnalyzerType *string `json:"searchAnalyzerType,omitempty" xml:"searchAnalyzerType,omitempty"`
}

func (s RunLibraryChatGenerationRequestTextSearchParameter) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestTextSearchParameter) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) SetLimit(v int32) *RunLibraryChatGenerationRequestTextSearchParameter {
	s.Limit = &v
	return s
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) SetSearchAnalyzerType(v string) *RunLibraryChatGenerationRequestTextSearchParameter {
	s.SearchAnalyzerType = &v
	return s
}

type RunLibraryChatGenerationRequestVectorSearchParameter struct {
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s RunLibraryChatGenerationRequestVectorSearchParameter) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestVectorSearchParameter) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestVectorSearchParameter) SetLimit(v int32) *RunLibraryChatGenerationRequestVectorSearchParameter {
	s.Limit = &v
	return s
}

type RunLibraryChatGenerationResponseBody struct {
	// example:
	//
	// null
	Cost *int64      `json:"cost,omitempty" xml:"cost,omitempty"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
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
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RunLibraryChatGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationResponseBody) SetCost(v int64) *RunLibraryChatGenerationResponseBody {
	s.Cost = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetData(v interface{}) *RunLibraryChatGenerationResponseBody {
	s.Data = v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetDataType(v string) *RunLibraryChatGenerationResponseBody {
	s.DataType = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetErrCode(v string) *RunLibraryChatGenerationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetMessage(v string) *RunLibraryChatGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetRequestId(v string) *RunLibraryChatGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetSuccess(v bool) *RunLibraryChatGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetTime(v string) *RunLibraryChatGenerationResponseBody {
	s.Time = &v
	return s
}

type RunLibraryChatGenerationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunLibraryChatGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunLibraryChatGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationResponse) SetHeaders(v map[string]*string) *RunLibraryChatGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunLibraryChatGenerationResponse) SetStatusCode(v int32) *RunLibraryChatGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunLibraryChatGenerationResponse) SetBody(v *RunLibraryChatGenerationResponseBody) *RunLibraryChatGenerationResponse {
	s.Body = v
	return s
}

type SubmitChatQuestionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-09-27 11:23:20
	GmtService *string `json:"gmtService,omitempty" xml:"gmtService,omitempty"`
	// This parameter is required.
	LiveScriptContent *string `json:"liveScriptContent,omitempty" xml:"liveScriptContent,omitempty"`
	// example:
	//
	// true
	OpenSmallTalk *bool `json:"openSmallTalk,omitempty" xml:"openSmallTalk,omitempty"`
	// This parameter is required.
	QuestionList []*SubmitChatQuestionRequestQuestionList `json:"questionList,omitempty" xml:"questionList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s SubmitChatQuestionRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitChatQuestionRequest) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionRequest) SetGmtService(v string) *SubmitChatQuestionRequest {
	s.GmtService = &v
	return s
}

func (s *SubmitChatQuestionRequest) SetLiveScriptContent(v string) *SubmitChatQuestionRequest {
	s.LiveScriptContent = &v
	return s
}

func (s *SubmitChatQuestionRequest) SetOpenSmallTalk(v bool) *SubmitChatQuestionRequest {
	s.OpenSmallTalk = &v
	return s
}

func (s *SubmitChatQuestionRequest) SetQuestionList(v []*SubmitChatQuestionRequestQuestionList) *SubmitChatQuestionRequest {
	s.QuestionList = v
	return s
}

func (s *SubmitChatQuestionRequest) SetRequestId(v string) *SubmitChatQuestionRequest {
	s.RequestId = &v
	return s
}

func (s *SubmitChatQuestionRequest) SetSessionId(v string) *SubmitChatQuestionRequest {
	s.SessionId = &v
	return s
}

type SubmitChatQuestionRequestQuestionList struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-17 10:05:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Reply     *string `json:"reply,omitempty" xml:"reply,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1869300950603128834
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// PRODUCT_QA
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 39485783475638465
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s SubmitChatQuestionRequestQuestionList) String() string {
	return tea.Prettify(s)
}

func (s SubmitChatQuestionRequestQuestionList) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionRequestQuestionList) SetContent(v string) *SubmitChatQuestionRequestQuestionList {
	s.Content = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetGmtCreate(v string) *SubmitChatQuestionRequestQuestionList {
	s.GmtCreate = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetReply(v string) *SubmitChatQuestionRequestQuestionList {
	s.Reply = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetSessionId(v string) *SubmitChatQuestionRequestQuestionList {
	s.SessionId = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetType(v string) *SubmitChatQuestionRequestQuestionList {
	s.Type = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetUserId(v string) *SubmitChatQuestionRequestQuestionList {
	s.UserId = &v
	return s
}

func (s *SubmitChatQuestionRequestQuestionList) SetUserName(v string) *SubmitChatQuestionRequestQuestionList {
	s.UserName = &v
	return s
}

type SubmitChatQuestionResponseBody struct {
	// example:
	//
	// null
	Cost *int64                              `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *SubmitChatQuestionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 915AAAB9-4908-5224-9E53-9E9D7D0AA94B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s SubmitChatQuestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitChatQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionResponseBody) SetCost(v int64) *SubmitChatQuestionResponseBody {
	s.Cost = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetData(v *SubmitChatQuestionResponseBodyData) *SubmitChatQuestionResponseBody {
	s.Data = v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetDataType(v string) *SubmitChatQuestionResponseBody {
	s.DataType = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetErrCode(v string) *SubmitChatQuestionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetMessage(v string) *SubmitChatQuestionResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetRequestId(v string) *SubmitChatQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetSuccess(v bool) *SubmitChatQuestionResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitChatQuestionResponseBody) SetTime(v string) *SubmitChatQuestionResponseBody {
	s.Time = &v
	return s
}

type SubmitChatQuestionResponseBodyData struct {
	// example:
	//
	// 1869307330227937280
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
}

func (s SubmitChatQuestionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitChatQuestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionResponseBodyData) SetBatchId(v string) *SubmitChatQuestionResponseBodyData {
	s.BatchId = &v
	return s
}

type SubmitChatQuestionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitChatQuestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitChatQuestionResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitChatQuestionResponse) GoString() string {
	return s.String()
}

func (s *SubmitChatQuestionResponse) SetHeaders(v map[string]*string) *SubmitChatQuestionResponse {
	s.Headers = v
	return s
}

func (s *SubmitChatQuestionResponse) SetStatusCode(v int32) *SubmitChatQuestionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitChatQuestionResponse) SetBody(v *SubmitChatQuestionResponseBody) *SubmitChatQuestionResponse {
	s.Body = v
	return s
}

type UpdateDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// {
	//
	//         "businessId": "12321"
	//
	//     }
	Meta map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocumentRequest) SetDocId(v string) *UpdateDocumentRequest {
	s.DocId = &v
	return s
}

func (s *UpdateDocumentRequest) SetLibraryId(v string) *UpdateDocumentRequest {
	s.LibraryId = &v
	return s
}

func (s *UpdateDocumentRequest) SetMeta(v map[string]interface{}) *UpdateDocumentRequest {
	s.Meta = v
	return s
}

func (s *UpdateDocumentRequest) SetTitle(v string) *UpdateDocumentRequest {
	s.Title = &v
	return s
}

type UpdateDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// null
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s UpdateDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocumentResponseBody) SetCost(v int64) *UpdateDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetData(v string) *UpdateDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetDataType(v string) *UpdateDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetErrCode(v string) *UpdateDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetMessage(v string) *UpdateDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetRequestId(v string) *UpdateDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetSuccess(v bool) *UpdateDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetTime(v string) *UpdateDocumentResponseBody {
	s.Time = &v
	return s
}

type UpdateDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocumentResponse) SetHeaders(v map[string]*string) *UpdateDocumentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocumentResponse) SetStatusCode(v int32) *UpdateDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocumentResponse) SetBody(v *UpdateDocumentResponseBody) *UpdateDocumentResponse {
	s.Body = v
	return s
}

type UpdateDocumentChunkRequest struct {
	// This parameter is required.
	Chunks []*UpdateDocumentChunkRequestChunks `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sjdgdsfg
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s UpdateDocumentChunkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentChunkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocumentChunkRequest) SetChunks(v []*UpdateDocumentChunkRequestChunks) *UpdateDocumentChunkRequest {
	s.Chunks = v
	return s
}

func (s *UpdateDocumentChunkRequest) SetLibraryId(v string) *UpdateDocumentChunkRequest {
	s.LibraryId = &v
	return s
}

type UpdateDocumentChunkRequestChunks struct {
	// This parameter is required.
	//
	// example:
	//
	// 1987834755763847
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// This parameter is required.
	ChunkText *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
}

func (s UpdateDocumentChunkRequestChunks) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentChunkRequestChunks) GoString() string {
	return s.String()
}

func (s *UpdateDocumentChunkRequestChunks) SetChunkId(v string) *UpdateDocumentChunkRequestChunks {
	s.ChunkId = &v
	return s
}

func (s *UpdateDocumentChunkRequestChunks) SetChunkText(v string) *UpdateDocumentChunkRequestChunks {
	s.ChunkText = &v
	return s
}

type UpdateDocumentChunkResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// SUCCESS
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// 003D019A-1BB3-53EC-A0D2-CE76DA5D73B1
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

func (s UpdateDocumentChunkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentChunkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocumentChunkResponseBody) SetCost(v int64) *UpdateDocumentChunkResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetData(v string) *UpdateDocumentChunkResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetDataType(v string) *UpdateDocumentChunkResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetErrCode(v string) *UpdateDocumentChunkResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetMessage(v string) *UpdateDocumentChunkResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetRequestId(v string) *UpdateDocumentChunkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetSuccess(v bool) *UpdateDocumentChunkResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDocumentChunkResponseBody) SetTime(v string) *UpdateDocumentChunkResponseBody {
	s.Time = &v
	return s
}

type UpdateDocumentChunkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDocumentChunkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDocumentChunkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentChunkResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocumentChunkResponse) SetHeaders(v map[string]*string) *UpdateDocumentChunkResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocumentChunkResponse) SetStatusCode(v int32) *UpdateDocumentChunkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocumentChunkResponse) SetBody(v *UpdateDocumentChunkResponseBody) *UpdateDocumentChunkResponse {
	s.Body = v
	return s
}

type UpdateLibraryRequest struct {
	Description  *string                           `json:"description,omitempty" xml:"description,omitempty"`
	IndexSetting *UpdateLibraryRequestIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// dsfbashdbb
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s UpdateLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequest) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequest) SetDescription(v string) *UpdateLibraryRequest {
	s.Description = &v
	return s
}

func (s *UpdateLibraryRequest) SetIndexSetting(v *UpdateLibraryRequestIndexSetting) *UpdateLibraryRequest {
	s.IndexSetting = v
	return s
}

func (s *UpdateLibraryRequest) SetLibraryId(v string) *UpdateLibraryRequest {
	s.LibraryId = &v
	return s
}

func (s *UpdateLibraryRequest) SetLibraryName(v string) *UpdateLibraryRequest {
	s.LibraryName = &v
	return s
}

type UpdateLibraryRequestIndexSetting struct {
	ChunkStrategy      *UpdateLibraryRequestIndexSettingChunkStrategy      `json:"chunkStrategy,omitempty" xml:"chunkStrategy,omitempty" type:"Struct"`
	ModelConfig        *UpdateLibraryRequestIndexSettingModelConfig        `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	PromptRoleStyle    *string                                             `json:"promptRoleStyle,omitempty" xml:"promptRoleStyle,omitempty"`
	QueryEnhancer      *UpdateLibraryRequestIndexSettingQueryEnhancer      `json:"queryEnhancer,omitempty" xml:"queryEnhancer,omitempty" type:"Struct"`
	RecallStrategy     *UpdateLibraryRequestIndexSettingRecallStrategy     `json:"recallStrategy,omitempty" xml:"recallStrategy,omitempty" type:"Struct"`
	TextIndexSetting   *UpdateLibraryRequestIndexSettingTextIndexSetting   `json:"textIndexSetting,omitempty" xml:"textIndexSetting,omitempty" type:"Struct"`
	VectorIndexSetting *UpdateLibraryRequestIndexSettingVectorIndexSetting `json:"vectorIndexSetting,omitempty" xml:"vectorIndexSetting,omitempty" type:"Struct"`
}

func (s UpdateLibraryRequestIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSetting) SetChunkStrategy(v *UpdateLibraryRequestIndexSettingChunkStrategy) *UpdateLibraryRequestIndexSetting {
	s.ChunkStrategy = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetModelConfig(v *UpdateLibraryRequestIndexSettingModelConfig) *UpdateLibraryRequestIndexSetting {
	s.ModelConfig = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetPromptRoleStyle(v string) *UpdateLibraryRequestIndexSetting {
	s.PromptRoleStyle = &v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetQueryEnhancer(v *UpdateLibraryRequestIndexSettingQueryEnhancer) *UpdateLibraryRequestIndexSetting {
	s.QueryEnhancer = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetRecallStrategy(v *UpdateLibraryRequestIndexSettingRecallStrategy) *UpdateLibraryRequestIndexSetting {
	s.RecallStrategy = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetTextIndexSetting(v *UpdateLibraryRequestIndexSettingTextIndexSetting) *UpdateLibraryRequestIndexSetting {
	s.TextIndexSetting = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetVectorIndexSetting(v *UpdateLibraryRequestIndexSettingVectorIndexSetting) *UpdateLibraryRequestIndexSetting {
	s.VectorIndexSetting = v
	return s
}

type UpdateLibraryRequestIndexSettingChunkStrategy struct {
	// example:
	//
	// true
	DocTreeSplit *bool `json:"docTreeSplit,omitempty" xml:"docTreeSplit,omitempty"`
	// example:
	//
	// 160
	DocTreeSplitSize *int32 `json:"docTreeSplitSize,omitempty" xml:"docTreeSplitSize,omitempty"`
	// example:
	//
	// true
	EnhanceGraph *bool `json:"enhanceGraph,omitempty" xml:"enhanceGraph,omitempty"`
	// example:
	//
	// true
	EnhanceTable *bool `json:"enhanceTable,omitempty" xml:"enhanceTable,omitempty"`
	// example:
	//
	// 20
	Overlap *int32 `json:"overlap,omitempty" xml:"overlap,omitempty"`
	// example:
	//
	// true
	SentenceSplit *bool `json:"sentenceSplit,omitempty" xml:"sentenceSplit,omitempty"`
	// example:
	//
	// 160
	SentenceSplitSize *int32 `json:"sentenceSplitSize,omitempty" xml:"sentenceSplitSize,omitempty"`
	// example:
	//
	// 256
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// true
	Split *bool `json:"split,omitempty" xml:"split,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingChunkStrategy) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetDocTreeSplit(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.DocTreeSplit = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetDocTreeSplitSize(v int32) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.DocTreeSplitSize = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetEnhanceGraph(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.EnhanceGraph = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetEnhanceTable(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.EnhanceTable = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetOverlap(v int32) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.Overlap = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetSentenceSplit(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.SentenceSplit = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetSentenceSplitSize(v int32) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.SentenceSplitSize = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetSize(v int32) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.Size = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetSplit(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.Split = &v
	return s
}

type UpdateLibraryRequestIndexSettingModelConfig struct {
	// example:
	//
	// 0.8
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// topP
	//
	// example:
	//
	// 0.8
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingModelConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) SetTemperature(v float64) *UpdateLibraryRequestIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) SetTopP(v float64) *UpdateLibraryRequestIndexSettingModelConfig {
	s.TopP = &v
	return s
}

type UpdateLibraryRequestIndexSettingQueryEnhancer struct {
	// example:
	//
	// true
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// true
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// true
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// true
	EnableQueryRewrite *bool `json:"enableQueryRewrite,omitempty" xml:"enableQueryRewrite,omitempty"`
	// example:
	//
	// true
	EnableSession *bool `json:"enableSession,omitempty" xml:"enableSession,omitempty"`
	// example:
	//
	// sjdhgfc
	LocalKnowledgeId *string `json:"localKnowledgeId,omitempty" xml:"localKnowledgeId,omitempty"`
	// example:
	//
	// true
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingQueryEnhancer) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableFollowUp(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableFollowUp = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableMultiQuery(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableMultiQuery = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableOpenQa(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableOpenQa = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableQueryRewrite(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableQueryRewrite = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableSession(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableSession = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetLocalKnowledgeId(v string) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.LocalKnowledgeId = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetWithDocumentReference(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.WithDocumentReference = &v
	return s
}

type UpdateLibraryRequestIndexSettingRecallStrategy struct {
	// example:
	//
	// model
	DocumentRankType *string `json:"documentRankType,omitempty" xml:"documentRankType,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingRecallStrategy) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) SetDocumentRankType(v string) *UpdateLibraryRequestIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) SetLimit(v int32) *UpdateLibraryRequestIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

type UpdateLibraryRequestIndexSettingTextIndexSetting struct {
	// example:
	//
	// ElasticSearch
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// Standard
	IndexAnalyzer *string `json:"indexAnalyzer,omitempty" xml:"indexAnalyzer,omitempty"`
	// example:
	//
	// 0.5
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// Standard
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" xml:"searchAnalyzer,omitempty"`
	// example:
	//
	// 50
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingTextIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetCategory(v string) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.Category = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetEnable(v bool) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.Enable = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetIndexAnalyzer(v string) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.IndexAnalyzer = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetRankThreshold(v float64) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetSearchAnalyzer(v string) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.SearchAnalyzer = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetTopK(v int32) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.TopK = &v
	return s
}

type UpdateLibraryRequestIndexSettingVectorIndexSetting struct {
	// example:
	//
	// ADB
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// DashScope
	EmbeddingType *string `json:"embeddingType,omitempty" xml:"embeddingType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 0.5
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingVectorIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetCategory(v string) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.Category = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetEmbeddingType(v string) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.EmbeddingType = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetEnable(v bool) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.Enable = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetRankThreshold(v float64) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetTopK(v int32) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.TopK = &v
	return s
}

type UpdateLibraryResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// null
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s UpdateLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLibraryResponseBody) SetCost(v int64) *UpdateLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetData(v string) *UpdateLibraryResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetDataType(v string) *UpdateLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetErrCode(v string) *UpdateLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetMessage(v string) *UpdateLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetRequestId(v string) *UpdateLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetSuccess(v bool) *UpdateLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetTime(v string) *UpdateLibraryResponseBody {
	s.Time = &v
	return s
}

type UpdateLibraryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryResponse) GoString() string {
	return s.String()
}

func (s *UpdateLibraryResponse) SetHeaders(v map[string]*string) *UpdateLibraryResponse {
	s.Headers = v
	return s
}

func (s *UpdateLibraryResponse) SetStatusCode(v int32) *UpdateLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLibraryResponse) SetBody(v *UpdateLibraryResponseBody) *UpdateLibraryResponse {
	s.Body = v
	return s
}

type UpdateQaLibraryRequest struct {
	// This parameter is required.
	ParseQaResults []*UpdateQaLibraryRequestParseQaResults `json:"parseQaResults,omitempty" xml:"parseQaResults,omitempty" type:"Repeated"`
	// example:
	//
	// 6jh378d
	QaLibraryId *string `json:"qaLibraryId,omitempty" xml:"qaLibraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQaLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQaLibraryRequest) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryRequest) SetParseQaResults(v []*UpdateQaLibraryRequestParseQaResults) *UpdateQaLibraryRequest {
	s.ParseQaResults = v
	return s
}

func (s *UpdateQaLibraryRequest) SetQaLibraryId(v string) *UpdateQaLibraryRequest {
	s.QaLibraryId = &v
	return s
}

func (s *UpdateQaLibraryRequest) SetRequestId(v string) *UpdateQaLibraryRequest {
	s.RequestId = &v
	return s
}

type UpdateQaLibraryRequestParseQaResults struct {
	// This parameter is required.
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	// This parameter is required.
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
}

func (s UpdateQaLibraryRequestParseQaResults) String() string {
	return tea.Prettify(s)
}

func (s UpdateQaLibraryRequestParseQaResults) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryRequestParseQaResults) SetAnswer(v string) *UpdateQaLibraryRequestParseQaResults {
	s.Answer = &v
	return s
}

func (s *UpdateQaLibraryRequestParseQaResults) SetQuestion(v string) *UpdateQaLibraryRequestParseQaResults {
	s.Question = &v
	return s
}

type UpdateQaLibraryResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *UpdateQaLibraryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s UpdateQaLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQaLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryResponseBody) SetCost(v int64) *UpdateQaLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetData(v *UpdateQaLibraryResponseBodyData) *UpdateQaLibraryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetDataType(v string) *UpdateQaLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetErrCode(v string) *UpdateQaLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetMessage(v string) *UpdateQaLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetRequestId(v string) *UpdateQaLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetSuccess(v bool) *UpdateQaLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQaLibraryResponseBody) SetTime(v string) *UpdateQaLibraryResponseBody {
	s.Time = &v
	return s
}

type UpdateQaLibraryResponseBodyData struct {
	// example:
	//
	// 6jh378d
	QaLibraryId *string `json:"qaLibraryId,omitempty" xml:"qaLibraryId,omitempty"`
}

func (s UpdateQaLibraryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateQaLibraryResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryResponseBodyData) SetQaLibraryId(v string) *UpdateQaLibraryResponseBodyData {
	s.QaLibraryId = &v
	return s
}

type UpdateQaLibraryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQaLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQaLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQaLibraryResponse) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryResponse) SetHeaders(v map[string]*string) *UpdateQaLibraryResponse {
	s.Headers = v
	return s
}

func (s *UpdateQaLibraryResponse) SetStatusCode(v int32) *UpdateQaLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQaLibraryResponse) SetBody(v *UpdateQaLibraryResponseBody) *UpdateQaLibraryResponse {
	s.Body = v
	return s
}

type UploadDocumentRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss-xxx.hangzhou.com/test.pdf
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdhbcsj
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s UploadDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentRequest) SetData(v string) *UploadDocumentRequest {
	s.Data = &v
	return s
}

func (s *UploadDocumentRequest) SetFileName(v string) *UploadDocumentRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentRequest) SetFileUrl(v string) *UploadDocumentRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadDocumentRequest) SetLibraryId(v string) *UploadDocumentRequest {
	s.LibraryId = &v
	return s
}

type UploadDocumentAdvanceRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss-xxx.hangzhou.com/test.pdf
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdhbcsj
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s UploadDocumentAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentAdvanceRequest) SetData(v string) *UploadDocumentAdvanceRequest {
	s.Data = &v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetFileName(v string) *UploadDocumentAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetFileUrlObject(v io.Reader) *UploadDocumentAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetLibraryId(v string) *UploadDocumentAdvanceRequest {
	s.LibraryId = &v
	return s
}

type UploadDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 1782981430906818562
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
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
	// ff3fef67-48d9-4379-a237-9ba8143fe739
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s UploadDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponseBody) SetCost(v int64) *UploadDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *UploadDocumentResponseBody) SetData(v string) *UploadDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UploadDocumentResponseBody) SetDataType(v string) *UploadDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *UploadDocumentResponseBody) SetErrCode(v string) *UploadDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UploadDocumentResponseBody) SetMessage(v string) *UploadDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDocumentResponseBody) SetRequestId(v string) *UploadDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDocumentResponseBody) SetSuccess(v bool) *UploadDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDocumentResponseBody) SetTime(v string) *UploadDocumentResponseBody {
	s.Time = &v
	return s
}

type UploadDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentResponse) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponse) SetHeaders(v map[string]*string) *UploadDocumentResponse {
	s.Headers = v
	return s
}

func (s *UploadDocumentResponse) SetStatusCode(v int32) *UploadDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDocumentResponse) SetBody(v *UploadDocumentResponseBody) *UploadDocumentResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dianjin"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建按年文档总结任务
//
// @param request - CreateAnnualDocSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnnualDocSummaryTaskResponse
func (client *Client) CreateAnnualDocSummaryTaskWithOptions(workspaceId *string, request *CreateAnnualDocSummaryTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAnnualDocSummaryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnaYears)) {
		body["anaYears"] = request.AnaYears
	}

	if !tea.BoolValue(util.IsUnset(request.DocInfos)) {
		body["docInfos"] = request.DocInfos
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTable)) {
		body["enableTable"] = request.EnableTable
	}

	if !tea.BoolValue(util.IsUnset(request.Instruction)) {
		body["instruction"] = request.Instruction
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAnnualDocSummaryTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/summary/doc/annual"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAnnualDocSummaryTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAnnualDocSummaryTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建按年文档总结任务
//
// @param request - CreateAnnualDocSummaryTaskRequest
//
// @return CreateAnnualDocSummaryTaskResponse
func (client *Client) CreateAnnualDocSummaryTask(workspaceId *string, request *CreateAnnualDocSummaryTaskRequest) (_result *CreateAnnualDocSummaryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAnnualDocSummaryTaskResponse{}
	_body, _err := client.CreateAnnualDocSummaryTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建外呼会话
//
// @param request - CreateDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDialogResponse
func (client *Client) CreateDialogWithOptions(workspaceId *string, request *CreateDialogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.EnableLibrary)) {
		body["enableLibrary"] = request.EnableLibrary
	}

	if !tea.BoolValue(util.IsUnset(request.MetaData)) {
		body["metaData"] = request.MetaData
	}

	if !tea.BoolValue(util.IsUnset(request.PlayCode)) {
		body["playCode"] = request.PlayCode
	}

	if !tea.BoolValue(util.IsUnset(request.QaLibraryList)) {
		body["qaLibraryList"] = request.QaLibraryList
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SelfDirected)) {
		body["selfDirected"] = request.SelfDirected
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDialog"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/virtualHuman/dialog/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDialogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDialogResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建外呼会话
//
// @param request - CreateDialogRequest
//
// @return CreateDialogResponse
func (client *Client) CreateDialog(workspaceId *string, request *CreateDialogRequest) (_result *CreateDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDialogResponse{}
	_body, _err := client.CreateDialogWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建会话分析任务
//
// @param request - CreateDialogAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDialogAnalysisTaskResponse
func (client *Client) CreateDialogAnalysisTaskWithOptions(workspaceId *string, request *CreateDialogAnalysisTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDialogAnalysisTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalysisNodes)) {
		body["analysisNodes"] = request.AnalysisNodes
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationList)) {
		body["conversationList"] = request.ConversationList
	}

	if !tea.BoolValue(util.IsUnset(request.MetaData)) {
		body["metaData"] = request.MetaData
	}

	if !tea.BoolValue(util.IsUnset(request.PlayCode)) {
		body["playCode"] = request.PlayCode
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDialogAnalysisTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/virtualHuman/dialog/analysis/submit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDialogAnalysisTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDialogAnalysisTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建会话分析任务
//
// @param request - CreateDialogAnalysisTaskRequest
//
// @return CreateDialogAnalysisTaskResponse
func (client *Client) CreateDialogAnalysisTask(workspaceId *string, request *CreateDialogAnalysisTaskRequest) (_result *CreateDialogAnalysisTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDialogAnalysisTaskResponse{}
	_body, _err := client.CreateDialogAnalysisTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateDocsSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocsSummaryTaskResponse
func (client *Client) CreateDocsSummaryTaskWithOptions(workspaceId *string, request *CreateDocsSummaryTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDocsSummaryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocInfos)) {
		body["docInfos"] = request.DocInfos
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTable)) {
		body["enableTable"] = request.EnableTable
	}

	if !tea.BoolValue(util.IsUnset(request.Instruction)) {
		body["instruction"] = request.Instruction
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDocsSummaryTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/summary/docs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDocsSummaryTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDocsSummaryTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateDocsSummaryTaskRequest
//
// @return CreateDocsSummaryTaskResponse
func (client *Client) CreateDocsSummaryTask(workspaceId *string, request *CreateDocsSummaryTaskRequest) (_result *CreateDocsSummaryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDocsSummaryTaskResponse{}
	_body, _err := client.CreateDocsSummaryTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateFinReportSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFinReportSummaryTaskResponse
func (client *Client) CreateFinReportSummaryTaskWithOptions(workspaceId *string, request *CreateFinReportSummaryTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFinReportSummaryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTable)) {
		body["enableTable"] = request.EnableTable
	}

	if !tea.BoolValue(util.IsUnset(request.EndPage)) {
		body["endPage"] = request.EndPage
	}

	if !tea.BoolValue(util.IsUnset(request.Instruction)) {
		body["instruction"] = request.Instruction
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.StartPage)) {
		body["startPage"] = request.StartPage
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFinReportSummaryTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/summary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateFinReportSummaryTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateFinReportSummaryTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateFinReportSummaryTaskRequest
//
// @return CreateFinReportSummaryTaskResponse
func (client *Client) CreateFinReportSummaryTask(workspaceId *string, request *CreateFinReportSummaryTaskRequest) (_result *CreateFinReportSummaryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFinReportSummaryTaskResponse{}
	_body, _err := client.CreateFinReportSummaryTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文档库
//
// @param request - CreateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLibraryResponse
func (client *Client) CreateLibraryWithOptions(workspaceId *string, request *CreateLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IndexSetting)) {
		body["indexSetting"] = request.IndexSetting
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryName)) {
		body["libraryName"] = request.LibraryName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLibraryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLibraryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建文档库
//
// @param request - CreateLibraryRequest
//
// @return CreateLibraryResponse
func (client *Client) CreateLibrary(workspaceId *string, request *CreateLibraryRequest) (_result *CreateLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLibraryResponse{}
	_body, _err := client.CreateLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建PDF翻译任务
//
// @param request - CreatePdfTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdfTranslateTaskResponse
func (client *Client) CreatePdfTranslateTaskWithOptions(workspaceId *string, request *CreatePdfTranslateTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePdfTranslateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.Knowledge)) {
		body["knowledge"] = request.Knowledge
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.TranslateTo)) {
		body["translateTo"] = request.TranslateTo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePdfTranslateTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/pdfTranslate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePdfTranslateTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePdfTranslateTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建PDF翻译任务
//
// @param request - CreatePdfTranslateTaskRequest
//
// @return CreatePdfTranslateTaskResponse
func (client *Client) CreatePdfTranslateTask(workspaceId *string, request *CreatePdfTranslateTaskRequest) (_result *CreatePdfTranslateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePdfTranslateTaskResponse{}
	_body, _err := client.CreatePdfTranslateTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建预定义文档
//
// @param request - CreatePredefinedDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePredefinedDocumentResponse
func (client *Client) CreatePredefinedDocumentWithOptions(workspaceId *string, request *CreatePredefinedDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePredefinedDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Chunks)) {
		body["chunks"] = request.Chunks
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Metadata)) {
		body["metadata"] = request.Metadata
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePredefinedDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/createPredefinedDocument"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePredefinedDocumentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePredefinedDocumentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建预定义文档
//
// @param request - CreatePredefinedDocumentRequest
//
// @return CreatePredefinedDocumentResponse
func (client *Client) CreatePredefinedDocument(workspaceId *string, request *CreatePredefinedDocumentRequest) (_result *CreatePredefinedDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePredefinedDocumentResponse{}
	_body, _err := client.CreatePredefinedDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建财报总结的任务
//
// @param request - CreateQualityCheckTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityCheckTaskResponse
func (client *Client) CreateQualityCheckTaskWithOptions(workspaceId *string, request *CreateQualityCheckTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateQualityCheckTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationList)) {
		body["conversationList"] = request.ConversationList
	}

	if !tea.BoolValue(util.IsUnset(request.GmtService)) {
		body["gmtService"] = request.GmtService
	}

	if !tea.BoolValue(util.IsUnset(request.MetaData)) {
		body["metaData"] = request.MetaData
	}

	if !tea.BoolValue(util.IsUnset(request.QualityGroup)) {
		body["qualityGroup"] = request.QualityGroup
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQualityCheckTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/qualitycheck/task/submit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateQualityCheckTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateQualityCheckTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建财报总结的任务
//
// @param request - CreateQualityCheckTaskRequest
//
// @return CreateQualityCheckTaskResponse
func (client *Client) CreateQualityCheckTask(workspaceId *string, request *CreateQualityCheckTaskRequest) (_result *CreateQualityCheckTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQualityCheckTaskResponse{}
	_body, _err := client.CreateQualityCheckTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @param request - DeleteDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocumentWithOptions(workspaceId *string, request *DeleteDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocIds)) {
		body["docIds"] = request.DocIds
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteDocumentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteDocumentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除文档
//
// @param request - DeleteDocumentRequest
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocument(workspaceId *string, request *DeleteDocumentRequest) (_result *DeleteDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDocumentResponse{}
	_body, _err := client.DeleteDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文档库
//
// @param request - DeleteLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLibraryResponse
func (client *Client) DeleteLibraryWithOptions(workspaceId *string, request *DeleteLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		query["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteLibraryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteLibraryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除文档库
//
// @param request - DeleteLibraryRequest
//
// @return DeleteLibraryResponse
func (client *Client) DeleteLibrary(workspaceId *string, request *DeleteLibraryRequest) (_result *DeleteLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLibraryResponse{}
	_body, _err := client.DeleteLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 中断任务
//
// @param request - EvictTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvictTaskResponse
func (client *Client) EvictTaskWithOptions(workspaceId *string, request *EvictTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EvictTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EvictTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/evict"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EvictTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EvictTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 中断任务
//
// @param request - EvictTaskRequest
//
// @return EvictTaskResponse
func (client *Client) EvictTask(workspaceId *string, request *EvictTaskRequest) (_result *EvictTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EvictTaskResponse{}
	_body, _err := client.EvictTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据文档解析问答QA
//
// @param request - GenDocQaResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenDocQaResultResponse
func (client *Client) GenDocQaResultWithOptions(workspaceId *string, request *GenDocQaResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenDocQaResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenDocQaResult"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/virtualHuman/qa/parse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenDocQaResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenDocQaResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 根据文档解析问答QA
//
// @param request - GenDocQaResultRequest
//
// @return GenDocQaResultResponse
func (client *Client) GenDocQaResult(workspaceId *string, request *GenDocQaResultRequest) (_result *GenDocQaResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenDocQaResultResponse{}
	_body, _err := client.GenDocQaResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取app配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppConfigResponse
func (client *Client) GetAppConfigWithOptions(workspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppConfig"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/app/config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAppConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAppConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取app配置
//
// @return GetAppConfigResponse
func (client *Client) GetAppConfig(workspaceId *string) (_result *GetAppConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppConfigResponse{}
	_body, _err := client.GetAppConfigWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取问答结果
//
// @param request - GetChatQuestionRespRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatQuestionRespResponse
func (client *Client) GetChatQuestionRespWithOptions(workspaceId *string, request *GetChatQuestionRespRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetChatQuestionRespResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		body["batchId"] = request.BatchId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChatQuestionResp"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/virtualHuman/chat/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetChatQuestionRespResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetChatQuestionRespResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取问答结果
//
// @param request - GetChatQuestionRespRequest
//
// @return GetChatQuestionRespResponse
func (client *Client) GetChatQuestionResp(workspaceId *string, request *GetChatQuestionRespRequest) (_result *GetChatQuestionRespResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatQuestionRespResponse{}
	_body, _err := client.GetChatQuestionRespWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取外呼会话分析结果
//
// @param request - GetDialogAnalysisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDialogAnalysisResultResponse
func (client *Client) GetDialogAnalysisResultWithOptions(workspaceId *string, request *GetDialogAnalysisResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDialogAnalysisResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Asc)) {
		body["asc"] = request.Asc
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.SessionIds)) {
		body["sessionIds"] = request.SessionIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UseUrl)) {
		body["useUrl"] = request.UseUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDialogAnalysisResult"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/virtualHuman/dialog/analysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDialogAnalysisResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDialogAnalysisResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取外呼会话分析结果
//
// @param request - GetDialogAnalysisResultRequest
//
// @return GetDialogAnalysisResultResponse
func (client *Client) GetDialogAnalysisResult(workspaceId *string, request *GetDialogAnalysisResultRequest) (_result *GetDialogAnalysisResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDialogAnalysisResultResponse{}
	_body, _err := client.GetDialogAnalysisResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetDialogDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDialogDetailResponse
func (client *Client) GetDialogDetailWithOptions(workspaceId *string, request *GetDialogDetailRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDialogDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDialogDetail"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/virtualHuman/dialog/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDialogDetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDialogDetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetDialogDetailRequest
//
// @return GetDialogDetailResponse
func (client *Client) GetDialogDetail(workspaceId *string, request *GetDialogDetailRequest) (_result *GetDialogDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDialogDetailResponse{}
	_body, _err := client.GetDialogDetailWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档的chunk列表
//
// @param request - GetDocumentChunkListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentChunkListResponse
func (client *Client) GetDocumentChunkListWithOptions(workspaceId *string, request *GetDocumentChunkListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentChunkListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChunkIdList)) {
		body["chunkIdList"] = request.ChunkIdList
	}

	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchQuery)) {
		body["searchQuery"] = request.SearchQuery
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentChunkList"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/getDocumentChunk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentChunkListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentChunkListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取文档的chunk列表
//
// @param request - GetDocumentChunkListRequest
//
// @return GetDocumentChunkListResponse
func (client *Client) GetDocumentChunkList(workspaceId *string, request *GetDocumentChunkListRequest) (_result *GetDocumentChunkListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentChunkListResponse{}
	_body, _err := client.GetDocumentChunkListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库的文档列表
//
// @param request - GetDocumentListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentListResponse
func (client *Client) GetDocumentListWithOptions(workspaceId *string, request *GetDocumentListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		query["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentList"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/listDocument"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 分页查询文档库的文档列表
//
// @param request - GetDocumentListRequest
//
// @return GetDocumentListResponse
func (client *Client) GetDocumentList(workspaceId *string, request *GetDocumentListRequest) (_result *GetDocumentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentListResponse{}
	_body, _err := client.GetDocumentListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档URL
//
// @param request - GetDocumentUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentUrlResponse
func (client *Client) GetDocumentUrlWithOptions(workspaceId *string, request *GetDocumentUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentId)) {
		query["documentId"] = request.DocumentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentUrl"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/url"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDocumentUrlResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDocumentUrlResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取文档URL
//
// @param request - GetDocumentUrlRequest
//
// @return GetDocumentUrlResponse
func (client *Client) GetDocumentUrl(workspaceId *string, request *GetDocumentUrlRequest) (_result *GetDocumentUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentUrlResponse{}
	_body, _err := client.GetDocumentUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 带条件的分页查询文档库的文档列表
//
// @param request - GetFilterDocumentListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFilterDocumentListResponse
func (client *Client) GetFilterDocumentListWithOptions(workspaceId *string, request *GetFilterDocumentListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFilterDocumentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.And)) {
		body["and"] = request.And
	}

	if !tea.BoolValue(util.IsUnset(request.DocIdList)) {
		body["docIdList"] = request.DocIdList
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Or)) {
		body["or"] = request.Or
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFilterDocumentList"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/filterDocument"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetFilterDocumentListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetFilterDocumentListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 带条件的分页查询文档库的文档列表
//
// @param request - GetFilterDocumentListRequest
//
// @return GetFilterDocumentListResponse
func (client *Client) GetFilterDocumentList(workspaceId *string, request *GetFilterDocumentListRequest) (_result *GetFilterDocumentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFilterDocumentListResponse{}
	_body, _err := client.GetFilterDocumentListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetHistoryListByBizTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoryListByBizTypeResponse
func (client *Client) GetHistoryListByBizTypeWithOptions(workspaceId *string, request *GetHistoryListByBizTypeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHistoryListByBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistoryListByBizType"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/history/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHistoryListByBizTypeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHistoryListByBizTypeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetHistoryListByBizTypeRequest
//
// @return GetHistoryListByBizTypeResponse
func (client *Client) GetHistoryListByBizType(workspaceId *string, request *GetHistoryListByBizTypeRequest) (_result *GetHistoryListByBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHistoryListByBizTypeResponse{}
	_body, _err := client.GetHistoryListByBizTypeWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档库配置详情
//
// @param request - GetLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryResponse
func (client *Client) GetLibraryWithOptions(workspaceId *string, request *GetLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		query["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/get"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLibraryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLibraryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取文档库配置详情
//
// @param request - GetLibraryRequest
//
// @return GetLibraryResponse
func (client *Client) GetLibrary(workspaceId *string, request *GetLibraryRequest) (_result *GetLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLibraryResponse{}
	_body, _err := client.GetLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetLibraryListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryListResponse
func (client *Client) GetLibraryListWithOptions(workspaceId *string, request *GetLibraryListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLibraryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLibraryList"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLibraryListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLibraryListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetLibraryListRequest
//
// @return GetLibraryListResponse
func (client *Client) GetLibraryList(workspaceId *string, request *GetLibraryListRequest) (_result *GetLibraryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLibraryListResponse{}
	_body, _err := client.GetLibraryListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取解析结果
//
// @param request - GetParseResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParseResultResponse
func (client *Client) GetParseResultWithOptions(workspaceId *string, request *GetParseResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetParseResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.UseUrlResult)) {
		body["useUrlResult"] = request.UseUrlResult
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetParseResult"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/getParseResult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetParseResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetParseResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取解析结果
//
// @param request - GetParseResultRequest
//
// @return GetParseResultResponse
func (client *Client) GetParseResult(workspaceId *string, request *GetParseResultRequest) (_result *GetParseResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetParseResultResponse{}
	_body, _err := client.GetParseResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetQualityCheckTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityCheckTaskResultResponse
func (client *Client) GetQualityCheckTaskResultWithOptions(workspaceId *string, request *GetQualityCheckTaskResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQualityCheckTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQualityCheckTaskResult"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/qualitycheck/task/query"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetQualityCheckTaskResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetQualityCheckTaskResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetQualityCheckTaskResultRequest
//
// @return GetQualityCheckTaskResultResponse
func (client *Client) GetQualityCheckTaskResult(workspaceId *string, request *GetQualityCheckTaskResultRequest) (_result *GetQualityCheckTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQualityCheckTaskResultResponse{}
	_body, _err := client.GetQualityCheckTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetSummaryTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSummaryTaskResultResponse
func (client *Client) GetSummaryTaskResultWithOptions(workspaceId *string, request *GetSummaryTaskResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSummaryTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSummaryTaskResult"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/summary/result"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetSummaryTaskResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetSummaryTaskResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetSummaryTaskResultRequest
//
// @return GetSummaryTaskResultResponse
func (client *Client) GetSummaryTaskResult(workspaceId *string, request *GetSummaryTaskResultRequest) (_result *GetSummaryTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSummaryTaskResultResponse{}
	_body, _err := client.GetSummaryTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务结果
//
// @param request - GetTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResultResponse
func (client *Client) GetTaskResultWithOptions(workspaceId *string, request *GetTaskResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskResult"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/result"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTaskResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTaskResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取异步任务结果
//
// @param request - GetTaskResultRequest
//
// @return GetTaskResultResponse
func (client *Client) GetTaskResult(workspaceId *string, request *GetTaskResultRequest) (_result *GetTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResultResponse{}
	_body, _err := client.GetTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithOptions(workspaceId *string, request *GetTaskStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatus"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTaskStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTaskStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetTaskStatusRequest
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatus(workspaceId *string, request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 插件调试接口
//
// @param request - InvokePluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokePluginResponse
func (client *Client) InvokePluginWithOptions(workspaceId *string, request *InvokePluginRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InvokePluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.PluginId)) {
		body["pluginId"] = request.PluginId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokePlugin"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/plugin/invoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &InvokePluginResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &InvokePluginResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 插件调试接口
//
// @param request - InvokePluginRequest
//
// @return InvokePluginResponse
func (client *Client) InvokePlugin(workspaceId *string, request *InvokePluginRequest) (_result *InvokePluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokePluginResponse{}
	_body, _err := client.InvokePluginWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档预览
//
// @param request - PreviewDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewDocumentResponse
func (client *Client) PreviewDocumentWithOptions(workspaceId *string, request *PreviewDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PreviewDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentId)) {
		query["documentId"] = request.DocumentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreviewDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/preview"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PreviewDocumentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PreviewDocumentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取文档预览
//
// @param request - PreviewDocumentRequest
//
// @return PreviewDocumentResponse
func (client *Client) PreviewDocument(workspaceId *string, request *PreviewDocumentRequest) (_result *PreviewDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PreviewDocumentResponse{}
	_body, _err := client.PreviewDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新索引
//
// @param request - ReIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReIndexResponse
func (client *Client) ReIndexWithOptions(workspaceId *string, request *ReIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentId)) {
		query["documentId"] = request.DocumentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReIndex"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/reIndex"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReIndexResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReIndexResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 重新索引
//
// @param request - ReIndexRequest
//
// @return ReIndexResponse
func (client *Client) ReIndex(workspaceId *string, request *ReIndexRequest) (_result *ReIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReIndexResponse{}
	_body, _err := client.ReIndexWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实时对话
//
// @param request - RealTimeDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RealTimeDialogResponse
func (client *Client) RealTimeDialogWithOptions(workspaceId *string, request *RealTimeDialogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RealTimeDialogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Analysis)) {
		body["analysis"] = request.Analysis
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationModel)) {
		body["conversationModel"] = request.ConversationModel
	}

	if !tea.BoolValue(util.IsUnset(request.DialogMemoryTurns)) {
		body["dialogMemoryTurns"] = request.DialogMemoryTurns
	}

	if !tea.BoolValue(util.IsUnset(request.MetaData)) {
		body["metaData"] = request.MetaData
	}

	if !tea.BoolValue(util.IsUnset(request.OpType)) {
		body["opType"] = request.OpType
	}

	if !tea.BoolValue(util.IsUnset(request.Recommend)) {
		body["recommend"] = request.Recommend
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptContentPlayed)) {
		body["scriptContentPlayed"] = request.ScriptContentPlayed
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.UserVad)) {
		body["userVad"] = request.UserVad
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RealTimeDialog"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/realtime/dialog/chat"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RealTimeDialogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RealTimeDialogResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 实时对话
//
// @param request - RealTimeDialogRequest
//
// @return RealTimeDialogResponse
func (client *Client) RealTimeDialog(workspaceId *string, request *RealTimeDialogRequest) (_result *RealTimeDialogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RealTimeDialogResponse{}
	_body, _err := client.RealTimeDialogWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重建任务
//
// @param request - RebuildTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildTaskResponse
func (client *Client) RebuildTaskWithOptions(workspaceId *string, request *RebuildTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RebuildTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		body["taskIds"] = request.TaskIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RebuildTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/rebuild"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RebuildTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RebuildTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 重建任务
//
// @param request - RebuildTaskRequest
//
// @return RebuildTaskResponse
func (client *Client) RebuildTask(workspaceId *string, request *RebuildTaskRequest) (_result *RebuildTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RebuildTaskResponse{}
	_body, _err := client.RebuildTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档召回。
//
// @param request - RecallDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallDocumentResponse
func (client *Client) RecallDocumentWithOptions(workspaceId *string, request *RecallDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecallDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		body["filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Rearrangement)) {
		body["rearrangement"] = request.Rearrangement
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		body["topK"] = request.TopK
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecallDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/recallDocument"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RecallDocumentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RecallDocumentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 文档召回。
//
// @param request - RecallDocumentRequest
//
// @return RecallDocumentResponse
func (client *Client) RecallDocument(workspaceId *string, request *RecallDocumentRequest) (_result *RecallDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecallDocumentResponse{}
	_body, _err := client.RecallDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图识别
//
// @param request - RecognizeIntentionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeIntentionResponse
func (client *Client) RecognizeIntentionWithOptions(workspaceId *string, request *RecognizeIntentionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecognizeIntentionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Analysis)) {
		body["analysis"] = request.Analysis
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Conversation)) {
		body["conversation"] = request.Conversation
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalIntentionList)) {
		body["globalIntentionList"] = request.GlobalIntentionList
	}

	if !tea.BoolValue(util.IsUnset(request.HierarchicalIntentionList)) {
		body["hierarchicalIntentionList"] = request.HierarchicalIntentionList
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionDomainCode)) {
		body["intentionDomainCode"] = request.IntentionDomainCode
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionList)) {
		body["intentionList"] = request.IntentionList
	}

	if !tea.BoolValue(util.IsUnset(request.OpType)) {
		body["opType"] = request.OpType
	}

	if !tea.BoolValue(util.IsUnset(request.Recommend)) {
		body["recommend"] = request.Recommend
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeIntention"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/recog/intent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RecognizeIntentionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RecognizeIntentionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 意图识别
//
// @param request - RecognizeIntentionRequest
//
// @return RecognizeIntentionResponse
func (client *Client) RecognizeIntention(workspaceId *string, request *RecognizeIntentionRequest) (_result *RecognizeIntentionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecognizeIntentionResponse{}
	_body, _err := client.RecognizeIntentionWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行智能体
//
// @param request - RunAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAgentResponse
func (client *Client) RunAgentWithOptions(workspaceId *string, request *RunAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BotId)) {
		body["botId"] = request.BotId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.ThreadId)) {
		body["threadId"] = request.ThreadId
	}

	if !tea.BoolValue(util.IsUnset(request.UseDraft)) {
		body["useDraft"] = request.UseDraft
	}

	if !tea.BoolValue(util.IsUnset(request.UserContent)) {
		body["userContent"] = request.UserContent
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunAgent"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/bot/thread/run"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunAgentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunAgentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 运行智能体
//
// @param request - RunAgentRequest
//
// @return RunAgentResponse
func (client *Client) RunAgent(workspaceId *string, request *RunAgentRequest) (_result *RunAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunAgentResponse{}
	_body, _err := client.RunAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGenerationWithOptions(workspaceId *string, request *RunChatResultGenerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunChatResultGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InferenceParameters)) {
		body["inferenceParameters"] = request.InferenceParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Messages)) {
		body["messages"] = request.Messages
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.Tools)) {
		body["tools"] = request.Tools
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunChatResultGeneration"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/run/chat/generation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunChatResultGenerationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunChatResultGenerationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGeneration(workspaceId *string, request *RunChatResultGenerationRequest) (_result *RunChatResultGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunChatResultGenerationResponse{}
	_body, _err := client.RunChatResultGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGenerationWithOptions(workspaceId *string, request *RunLibraryChatGenerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunLibraryChatGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocIdList)) {
		body["docIdList"] = request.DocIdList
	}

	if !tea.BoolValue(util.IsUnset(request.EnableFollowUp)) {
		body["enableFollowUp"] = request.EnableFollowUp
	}

	if !tea.BoolValue(util.IsUnset(request.EnableMultiQuery)) {
		body["enableMultiQuery"] = request.EnableMultiQuery
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOpenQa)) {
		body["enableOpenQa"] = request.EnableOpenQa
	}

	if !tea.BoolValue(util.IsUnset(request.FollowUpLlm)) {
		body["followUpLlm"] = request.FollowUpLlm
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.LlmType)) {
		body["llmType"] = request.LlmType
	}

	if !tea.BoolValue(util.IsUnset(request.MultiQueryLlm)) {
		body["multiQueryLlm"] = request.MultiQueryLlm
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCriteria)) {
		body["queryCriteria"] = request.QueryCriteria
	}

	if !tea.BoolValue(util.IsUnset(request.RerankType)) {
		body["rerankType"] = request.RerankType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.SubQueryList)) {
		body["subQueryList"] = request.SubQueryList
	}

	if !tea.BoolValue(util.IsUnset(request.TextSearchParameter)) {
		body["textSearchParameter"] = request.TextSearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		body["topK"] = request.TopK
	}

	if !tea.BoolValue(util.IsUnset(request.VectorSearchParameter)) {
		body["vectorSearchParameter"] = request.VectorSearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.WithDocumentReference)) {
		body["withDocumentReference"] = request.WithDocumentReference
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunLibraryChatGeneration"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/run/library/chat/generation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunLibraryChatGenerationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunLibraryChatGenerationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGeneration(workspaceId *string, request *RunLibraryChatGenerationRequest) (_result *RunLibraryChatGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunLibraryChatGenerationResponse{}
	_body, _err := client.RunLibraryChatGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交问题列表
//
// @param request - SubmitChatQuestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitChatQuestionResponse
func (client *Client) SubmitChatQuestionWithOptions(workspaceId *string, request *SubmitChatQuestionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitChatQuestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GmtService)) {
		body["gmtService"] = request.GmtService
	}

	if !tea.BoolValue(util.IsUnset(request.LiveScriptContent)) {
		body["liveScriptContent"] = request.LiveScriptContent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenSmallTalk)) {
		body["openSmallTalk"] = request.OpenSmallTalk
	}

	if !tea.BoolValue(util.IsUnset(request.QuestionList)) {
		body["questionList"] = request.QuestionList
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitChatQuestion"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/virtualHuman/chat/submit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SubmitChatQuestionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SubmitChatQuestionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 提交问题列表
//
// @param request - SubmitChatQuestionRequest
//
// @return SubmitChatQuestionResponse
func (client *Client) SubmitChatQuestion(workspaceId *string, request *SubmitChatQuestionRequest) (_result *SubmitChatQuestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitChatQuestionResponse{}
	_body, _err := client.SubmitChatQuestionWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档
//
// @param request - UpdateDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocumentResponse
func (client *Client) UpdateDocumentWithOptions(workspaceId *string, request *UpdateDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Meta)) {
		body["meta"] = request.Meta
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/updateDocument"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateDocumentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateDocumentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新文档
//
// @param request - UpdateDocumentRequest
//
// @return UpdateDocumentResponse
func (client *Client) UpdateDocument(workspaceId *string, request *UpdateDocumentRequest) (_result *UpdateDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDocumentResponse{}
	_body, _err := client.UpdateDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档的chunk
//
// @param request - UpdateDocumentChunkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocumentChunkResponse
func (client *Client) UpdateDocumentChunkWithOptions(workspaceId *string, request *UpdateDocumentChunkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDocumentChunkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Chunks)) {
		body["chunks"] = request.Chunks
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDocumentChunk"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/updateDocumentChunk"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateDocumentChunkResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateDocumentChunkResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新文档的chunk
//
// @param request - UpdateDocumentChunkRequest
//
// @return UpdateDocumentChunkResponse
func (client *Client) UpdateDocumentChunk(workspaceId *string, request *UpdateDocumentChunkRequest) (_result *UpdateDocumentChunkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDocumentChunkResponse{}
	_body, _err := client.UpdateDocumentChunkWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档库配置
//
// @param request - UpdateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLibraryResponse
func (client *Client) UpdateLibraryWithOptions(workspaceId *string, request *UpdateLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IndexSetting)) {
		body["indexSetting"] = request.IndexSetting
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryName)) {
		body["libraryName"] = request.LibraryName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/update"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateLibraryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateLibraryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新文档库配置
//
// @param request - UpdateLibraryRequest
//
// @return UpdateLibraryResponse
func (client *Client) UpdateLibrary(workspaceId *string, request *UpdateLibraryRequest) (_result *UpdateLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLibraryResponse{}
	_body, _err := client.UpdateLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新QA问答库
//
// @param request - UpdateQaLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQaLibraryResponse
func (client *Client) UpdateQaLibraryWithOptions(workspaceId *string, request *UpdateQaLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateQaLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParseQaResults)) {
		body["parseQaResults"] = request.ParseQaResults
	}

	if !tea.BoolValue(util.IsUnset(request.QaLibraryId)) {
		body["qaLibraryId"] = request.QaLibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQaLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/virtualHuman/qa/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateQaLibraryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateQaLibraryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新QA问答库
//
// @param request - UpdateQaLibraryRequest
//
// @return UpdateQaLibraryResponse
func (client *Client) UpdateQaLibrary(workspaceId *string, request *UpdateQaLibraryRequest) (_result *UpdateQaLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQaLibraryResponse{}
	_body, _err := client.UpdateQaLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传文档到文档库
//
// @param request - UploadDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDocumentResponse
func (client *Client) UploadDocumentWithOptions(workspaceId *string, request *UploadDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UploadDocumentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UploadDocumentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 上传文档到文档库
//
// @param request - UploadDocumentRequest
//
// @return UploadDocumentResponse
func (client *Client) UploadDocument(workspaceId *string, request *UploadDocumentRequest) (_result *UploadDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadDocumentResponse{}
	_body, _err := client.UploadDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDocumentAdvance(workspaceId *string, request *UploadDocumentAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("DianJin"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	uploadDocumentReq := &UploadDocumentRequest{}
	openapiutil.Convert(request, uploadDocumentReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		uploadDocumentReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	uploadDocumentResp, _err := client.UploadDocumentWithOptions(workspaceId, uploadDocumentReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadDocumentResp
	return _result, _err
}
