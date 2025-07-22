// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlOptimizeAdviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSqlOptimizeAdviceResponseBody
	GetCode() *string
	SetData(v *GetSqlOptimizeAdviceResponseBodyData) *GetSqlOptimizeAdviceResponseBody
	GetData() *GetSqlOptimizeAdviceResponseBodyData
	SetMessage(v string) *GetSqlOptimizeAdviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSqlOptimizeAdviceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetSqlOptimizeAdviceResponseBody
	GetSuccess() *string
}

type GetSqlOptimizeAdviceResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *GetSqlOptimizeAdviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSqlOptimizeAdviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlOptimizeAdviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlOptimizeAdviceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSqlOptimizeAdviceResponseBody) GetData() *GetSqlOptimizeAdviceResponseBodyData {
	return s.Data
}

func (s *GetSqlOptimizeAdviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSqlOptimizeAdviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlOptimizeAdviceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetSqlOptimizeAdviceResponseBody) SetCode(v string) *GetSqlOptimizeAdviceResponseBody {
	s.Code = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) SetData(v *GetSqlOptimizeAdviceResponseBodyData) *GetSqlOptimizeAdviceResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) SetMessage(v string) *GetSqlOptimizeAdviceResponseBody {
	s.Message = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) SetRequestId(v string) *GetSqlOptimizeAdviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) SetSuccess(v string) *GetSqlOptimizeAdviceResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSqlOptimizeAdviceResponseBodyData struct {
	// The time when the task was created. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1632303861000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The URL that is used to download the file.
	//
	// example:
	//
	// https://das-sql-optimize.oss-cn-shanghai.aliyuncs.com/adb/oss_sql_optimize_advice/1083*******
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The time when the file expires. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The file expires three days after the task is created.
	//
	// example:
	//
	// 1632563061000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **INIT**: The task is being initialized.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **FINISH**: The task is complete.
	//
	// 	- **FAILED**: The task failed.
	//
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status code of the task. Valid values:
	//
	// 	- **NO_DATA**: No data is returned.
	//
	// 	- **INTERNAL_ERROR**: An internal error occurred.
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// example:
	//
	// SUCCESS
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2021091710461519216****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSqlOptimizeAdviceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSqlOptimizeAdviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlOptimizeAdviceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSqlOptimizeAdviceResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetSqlOptimizeAdviceResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetSqlOptimizeAdviceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSqlOptimizeAdviceResponseBodyData) GetStatusCode() *string {
	return s.StatusCode
}

func (s *GetSqlOptimizeAdviceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetCreateTime(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetDownloadUrl(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetExpireTime(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetStatus(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetStatusCode(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) SetTaskId(v string) *GetSqlOptimizeAdviceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
