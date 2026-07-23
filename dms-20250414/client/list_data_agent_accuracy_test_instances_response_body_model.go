// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataAgentAccuracyTestInstancesResponseBodyData) *ListDataAgentAccuracyTestInstancesResponseBody
	GetData() []*ListDataAgentAccuracyTestInstancesResponseBodyData
	SetErrorCode(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataAgentAccuracyTestInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetNextToken() *string
	SetPageNumber(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetSuccess() *string
	SetTimestamp(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetTimestamp() *string
	SetTotal(v string) *ListDataAgentAccuracyTestInstancesResponseBody
	GetTotal() *string
}

type ListDataAgentAccuracyTestInstancesResponseBody struct {
	// The response struct.
	Data []*ListDataAgentAccuracyTestInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The maximum number of entries returned per page. You can use this parameter together with NextToken to implement paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// zCXS*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-xxx-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The operation timestamp.
	//
	// example:
	//
	// 1768270172
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 3
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDataAgentAccuracyTestInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetData() []*ListDataAgentAccuracyTestInstancesResponseBodyData {
	return s.Data
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) GetTotal() *string {
	return s.Total
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetData(v []*ListDataAgentAccuracyTestInstancesResponseBodyData) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetErrorCode(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetErrorMessage(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetMaxResults(v int32) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetNextToken(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetPageNumber(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetPageSize(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetRequestId(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetSuccess(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetTimestamp(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.Timestamp = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) SetTotal(v string) *ListDataAgentAccuracyTestInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBody) Validate() error {
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

type ListDataAgentAccuracyTestInstancesResponseBodyData struct {
	// The accuracy test instance ID.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
	// The custom agent ID.
	//
	// example:
	//
	// ca-4x8uzp5wjqu4xxxxxxxxxx
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The UID of the workspace creator.
	//
	// example:
	//
	// 20282*****7591
	Creator    *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Datasource *string `json:"Datasource,omitempty" xml:"Datasource,omitempty"`
	// The ID of the test set file.
	//
	// example:
	//
	// f-8*******01m
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-06-30T07:31:09.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-06-30T07:31:09.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The language used for the analysis task.
	//
	// example:
	//
	// CHINESE
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The maximum number of concurrent sessions during the test.
	//
	// example:
	//
	// 5
	MaxConcurrent *string `json:"MaxConcurrent,omitempty" xml:"MaxConcurrent,omitempty"`
	// The analysis mode to be tested.
	//
	// example:
	//
	// 0
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The custom agent name.
	//
	// example:
	//
	// Agent测试名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether sessions are displayed after analysis. This parameter is not supported.
	//
	// example:
	//
	// true
	NeedDelete *string `json:"NeedDelete,omitempty" xml:"NeedDelete,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentAccuracyTestInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetDatasource() *string {
	return s.Datasource
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetMaxConcurrent() *string {
	return s.MaxConcurrent
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetMode() *int32 {
	return s.Mode
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetNeedDelete() *string {
	return s.NeedDelete
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetAccuracyTestInsId(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.AccuracyTestInsId = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetAgentId(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetCreator(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.Creator = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetDatasource(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.Datasource = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetFileId(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.FileId = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetGmtCreate(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetGmtModified(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetLanguage(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.Language = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetMaxConcurrent(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.MaxConcurrent = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetMode(v int32) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.Mode = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetName(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetNeedDelete(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.NeedDelete = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) SetWorkspaceId(v string) *ListDataAgentAccuracyTestInstancesResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
