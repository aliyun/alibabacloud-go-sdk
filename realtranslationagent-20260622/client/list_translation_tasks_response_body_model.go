// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranslationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTranslationTasksResponseBody
	GetCode() *string
	SetData(v *ListTranslationTasksResponseBodyData) *ListTranslationTasksResponseBody
	GetData() *ListTranslationTasksResponseBodyData
	SetMessage(v string) *ListTranslationTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTranslationTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTranslationTasksResponseBody
	GetSuccess() *bool
}

type ListTranslationTasksResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *ListTranslationTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 50ABF118-2F9D-51DF-B1FB-1E389817DC47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTranslationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTranslationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTranslationTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTranslationTasksResponseBody) GetData() *ListTranslationTasksResponseBodyData {
	return s.Data
}

func (s *ListTranslationTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTranslationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTranslationTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTranslationTasksResponseBody) SetCode(v string) *ListTranslationTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListTranslationTasksResponseBody) SetData(v *ListTranslationTasksResponseBodyData) *ListTranslationTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListTranslationTasksResponseBody) SetMessage(v string) *ListTranslationTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListTranslationTasksResponseBody) SetRequestId(v string) *ListTranslationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTranslationTasksResponseBody) SetSuccess(v bool) *ListTranslationTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListTranslationTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTranslationTasksResponseBodyData struct {
	// The data list.
	List []*ListTranslationTasksResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The maximum number of results returned per request when using the NextToken-based pagination.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates whether a token exists for the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no next query exists.
	//
	// - If **NextToken*	- has a value, the value is the token for the next query.
	//
	// example:
	//
	// AAAAAVpfrV4aVmra0dxbtRB74lmSGzegoejeIqxIET/WdX50
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTranslationTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTranslationTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTranslationTasksResponseBodyData) GetList() []*ListTranslationTasksResponseBodyDataList {
	return s.List
}

func (s *ListTranslationTasksResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTranslationTasksResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTranslationTasksResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListTranslationTasksResponseBodyData) SetList(v []*ListTranslationTasksResponseBodyDataList) *ListTranslationTasksResponseBodyData {
	s.List = v
	return s
}

func (s *ListTranslationTasksResponseBodyData) SetMaxResults(v int32) *ListTranslationTasksResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListTranslationTasksResponseBodyData) SetNextToken(v string) *ListTranslationTasksResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTranslationTasksResponseBodyData) SetTotal(v int64) *ListTranslationTasksResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTranslationTasksResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTranslationTasksResponseBodyDataList struct {
	// The task completion time, in 13-digit timestamp format.
	//
	// example:
	//
	// 1782459562000
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The credits consumed by this task.
	//
	// example:
	//
	// 81.2992
	CostCredits *float64 `json:"CostCredits,omitempty" xml:"CostCredits,omitempty"`
	// The time consumed, in milliseconds.
	//
	// example:
	//
	// 196
	CostTime *int64 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// acc_93****c936
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The creator name.
	//
	// example:
	//
	// tes_account@test.com
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The error message when the task fails.
	//
	// example:
	//
	// device offline
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The parsed file format.
	//
	// example:
	//
	// PPTX
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// The file name.
	//
	// example:
	//
	// translated_a_file.pptx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The task creation time, in 13-digit timestamp format.
	//
	// example:
	//
	// 1782459562000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// org_c6******cdc2ce7
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The source file address.
	//
	// example:
	//
	// translated_a_file.pptx
	OriginalFileName *string `json:"OriginalFileName,omitempty" xml:"OriginalFileName,omitempty"`
	// The page count of the uploaded file.
	//
	// example:
	//
	// 21
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The task progress.
	//
	// example:
	//
	// 61
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The language of the source file.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The task start time, in 13-digit timestamp format.
	//
	// example:
	//
	// 1782459562000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - CANCELLED: Cancelled.
	//
	// - COMPLETED: Completed.
	//
	// - FAILED: Failed.
	//
	// - PROCESSING: Processing.
	//
	// - PENDING: Pending.
	//
	// - ANALYZED: Analyzed.
	//
	// example:
	//
	// CANCELLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The target language.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The translation task ID.
	//
	// example:
	//
	// f9c35b0453b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// - DOCUMENT: document type.
	//
	// example:
	//
	// DOCUMENT
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The translation template. Valid values:
	//
	// - common: General.
	//
	// example:
	//
	// common
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The word count of the uploaded file.
	//
	// example:
	//
	// 1600
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 6458351*****0cc5
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" xml:"WorkSpaceId,omitempty"`
}

func (s ListTranslationTasksResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListTranslationTasksResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTranslationTasksResponseBodyDataList) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *ListTranslationTasksResponseBodyDataList) GetCostCredits() *float64 {
	return s.CostCredits
}

func (s *ListTranslationTasksResponseBodyDataList) GetCostTime() *int64 {
	return s.CostTime
}

func (s *ListTranslationTasksResponseBodyDataList) GetCreator() *string {
	return s.Creator
}

func (s *ListTranslationTasksResponseBodyDataList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListTranslationTasksResponseBodyDataList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTranslationTasksResponseBodyDataList) GetFileFormat() *string {
	return s.FileFormat
}

func (s *ListTranslationTasksResponseBodyDataList) GetFileName() *string {
	return s.FileName
}

func (s *ListTranslationTasksResponseBodyDataList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListTranslationTasksResponseBodyDataList) GetOrgId() *string {
	return s.OrgId
}

func (s *ListTranslationTasksResponseBodyDataList) GetOriginalFileName() *string {
	return s.OriginalFileName
}

func (s *ListTranslationTasksResponseBodyDataList) GetPageCount() *int64 {
	return s.PageCount
}

func (s *ListTranslationTasksResponseBodyDataList) GetProgress() *int32 {
	return s.Progress
}

func (s *ListTranslationTasksResponseBodyDataList) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *ListTranslationTasksResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTranslationTasksResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListTranslationTasksResponseBodyDataList) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *ListTranslationTasksResponseBodyDataList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTranslationTasksResponseBodyDataList) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTranslationTasksResponseBodyDataList) GetTemplate() *string {
	return s.Template
}

func (s *ListTranslationTasksResponseBodyDataList) GetWordCount() *int64 {
	return s.WordCount
}

func (s *ListTranslationTasksResponseBodyDataList) GetWorkSpaceId() *string {
	return s.WorkSpaceId
}

func (s *ListTranslationTasksResponseBodyDataList) SetCompleteTime(v string) *ListTranslationTasksResponseBodyDataList {
	s.CompleteTime = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetCostCredits(v float64) *ListTranslationTasksResponseBodyDataList {
	s.CostCredits = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetCostTime(v int64) *ListTranslationTasksResponseBodyDataList {
	s.CostTime = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetCreator(v string) *ListTranslationTasksResponseBodyDataList {
	s.Creator = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetCreatorName(v string) *ListTranslationTasksResponseBodyDataList {
	s.CreatorName = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetErrorMessage(v string) *ListTranslationTasksResponseBodyDataList {
	s.ErrorMessage = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetFileFormat(v string) *ListTranslationTasksResponseBodyDataList {
	s.FileFormat = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetFileName(v string) *ListTranslationTasksResponseBodyDataList {
	s.FileName = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetGmtCreate(v string) *ListTranslationTasksResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetOrgId(v string) *ListTranslationTasksResponseBodyDataList {
	s.OrgId = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetOriginalFileName(v string) *ListTranslationTasksResponseBodyDataList {
	s.OriginalFileName = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetPageCount(v int64) *ListTranslationTasksResponseBodyDataList {
	s.PageCount = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetProgress(v int32) *ListTranslationTasksResponseBodyDataList {
	s.Progress = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetSourceLanguage(v string) *ListTranslationTasksResponseBodyDataList {
	s.SourceLanguage = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetStartTime(v string) *ListTranslationTasksResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetStatus(v string) *ListTranslationTasksResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetTargetLanguage(v string) *ListTranslationTasksResponseBodyDataList {
	s.TargetLanguage = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetTaskId(v string) *ListTranslationTasksResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetTaskType(v string) *ListTranslationTasksResponseBodyDataList {
	s.TaskType = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetTemplate(v string) *ListTranslationTasksResponseBodyDataList {
	s.Template = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetWordCount(v int64) *ListTranslationTasksResponseBodyDataList {
	s.WordCount = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) SetWorkSpaceId(v string) *ListTranslationTasksResponseBodyDataList {
	s.WorkSpaceId = &v
	return s
}

func (s *ListTranslationTasksResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
