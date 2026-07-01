// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOperationRecordDetailResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetOperationRecordDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOperationRecordDetailResponseBody
	GetMessage() *string
	SetOperationRecordDetailResponse(v *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) *GetOperationRecordDetailResponseBody
	GetOperationRecordDetailResponse() *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse
	SetRequestId(v string) *GetOperationRecordDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOperationRecordDetailResponseBody
	GetSuccess() *bool
}

type GetOperationRecordDetailResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The execution record detail response.
	OperationRecordDetailResponse *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse `json:"OperationRecordDetailResponse,omitempty" xml:"OperationRecordDetailResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOperationRecordDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationRecordDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOperationRecordDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOperationRecordDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOperationRecordDetailResponseBody) GetOperationRecordDetailResponse() *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse {
	return s.OperationRecordDetailResponse
}

func (s *GetOperationRecordDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOperationRecordDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOperationRecordDetailResponseBody) SetCode(v string) *GetOperationRecordDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetOperationRecordDetailResponseBody) SetHttpStatusCode(v int32) *GetOperationRecordDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOperationRecordDetailResponseBody) SetMessage(v string) *GetOperationRecordDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetOperationRecordDetailResponseBody) SetOperationRecordDetailResponse(v *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) *GetOperationRecordDetailResponseBody {
	s.OperationRecordDetailResponse = v
	return s
}

func (s *GetOperationRecordDetailResponseBody) SetRequestId(v string) *GetOperationRecordDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationRecordDetailResponseBody) SetSuccess(v bool) *GetOperationRecordDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetOperationRecordDetailResponseBody) Validate() error {
	if s.OperationRecordDetailResponse != nil {
		if err := s.OperationRecordDetailResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOperationRecordDetailResponseBodyOperationRecordDetailResponse struct {
	// The file ID.
	//
	// example:
	//
	// 12113111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The log content.
	//
	// example:
	//
	// Task started...
	//
	// Executing SQL...
	//
	// Task finished successfully.
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	// The operation record ID.
	//
	// example:
	//
	// 987654321
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The list of execution results.
	Results []*GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) GetFileId() *int64 {
	return s.FileId
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) GetLogContent() *string {
	return s.LogContent
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) GetOperationId() *string {
	return s.OperationId
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) GetResults() []*GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults {
	return s.Results
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) SetFileId(v int64) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse {
	s.FileId = &v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) SetLogContent(v string) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse {
	s.LogContent = &v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) SetOperationId(v string) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse {
	s.OperationId = &v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) SetResults(v []*GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse {
	s.Results = v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponse) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults struct {
	// The result index.
	//
	// example:
	//
	// 0
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The result content.
	//
	// example:
	//
	// {"count":100,"data":[{"id":1}]}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The result SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM test_table
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The task ID.
	//
	// example:
	//
	// task_123456
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The result title.
	//
	// example:
	//
	// 查询结果
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) GoString() string {
	return s.String()
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) GetIndex() *int32 {
	return s.Index
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) GetResult() *string {
	return s.Result
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) GetSql() *string {
	return s.Sql
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) GetTitle() *string {
	return s.Title
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) SetIndex(v int32) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults {
	s.Index = &v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) SetResult(v string) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults {
	s.Result = &v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) SetSql(v string) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults {
	s.Sql = &v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) SetTaskId(v string) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults {
	s.TaskId = &v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) SetTitle(v string) *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults {
	s.Title = &v
	return s
}

func (s *GetOperationRecordDetailResponseBodyOperationRecordDetailResponseResults) Validate() error {
	return dara.Validate(s)
}
