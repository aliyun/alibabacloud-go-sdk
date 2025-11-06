// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZkTrackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListZkTrackResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *ListZkTrackResponseBody
	GetHttpCode() *string
	SetMessage(v string) *ListZkTrackResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *ListZkTrackResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListZkTrackResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListZkTrackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListZkTrackResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListZkTrackResponseBody
	GetTotalCount() *int64
	SetTraces(v []*ListZkTrackResponseBodyTraces) *ListZkTrackResponseBody
	GetTraces() []*ListZkTrackResponseBodyTraces
}

type ListZkTrackResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The track data.
	Traces []*ListZkTrackResponseBodyTraces `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Repeated"`
}

func (s ListZkTrackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListZkTrackResponseBody) GoString() string {
	return s.String()
}

func (s *ListZkTrackResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListZkTrackResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *ListZkTrackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListZkTrackResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListZkTrackResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListZkTrackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListZkTrackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListZkTrackResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListZkTrackResponseBody) GetTraces() []*ListZkTrackResponseBodyTraces {
	return s.Traces
}

func (s *ListZkTrackResponseBody) SetErrorCode(v string) *ListZkTrackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListZkTrackResponseBody) SetHttpCode(v string) *ListZkTrackResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListZkTrackResponseBody) SetMessage(v string) *ListZkTrackResponseBody {
	s.Message = &v
	return s
}

func (s *ListZkTrackResponseBody) SetPageNumber(v int64) *ListZkTrackResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListZkTrackResponseBody) SetPageSize(v int64) *ListZkTrackResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListZkTrackResponseBody) SetRequestId(v string) *ListZkTrackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZkTrackResponseBody) SetSuccess(v bool) *ListZkTrackResponseBody {
	s.Success = &v
	return s
}

func (s *ListZkTrackResponseBody) SetTotalCount(v int64) *ListZkTrackResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListZkTrackResponseBody) SetTraces(v []*ListZkTrackResponseBodyTraces) *ListZkTrackResponseBody {
	s.Traces = v
	return s
}

func (s *ListZkTrackResponseBody) Validate() error {
	if s.Traces != nil {
		for _, item := range s.Traces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListZkTrackResponseBodyTraces struct {
	// The access control list (ACL).
	//
	// example:
	//
	// world:anyone:cdrwa
	Acl *string `json:"Acl,omitempty" xml:"Acl,omitempty"`
	// The data type. Valid values:
	//
	// 	- persist
	//
	// 	- ephemeral
	//
	// example:
	//
	// persist
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The type of the event. For trajectory of the Notify type:
	//
	// 	- NodeCreated
	//
	// 	- NodeDeleted
	//
	// 	- NodeDataChanged
	//
	// 	- NodeChildrenChanged
	//
	// example:
	//
	// NodeCreated
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Indicates whether the transaction ended.
	//
	// example:
	//
	// true
	Finished *bool `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// The logging time.
	//
	// example:
	//
	// 2022-11-28 15:09:15,606
	LogDate *string `json:"LogDate,omitempty" xml:"LogDate,omitempty"`
	// The transaction size.
	//
	// example:
	//
	// 3
	MultiSize *int64 `json:"MultiSize,omitempty" xml:"MultiSize,omitempty"`
	// The type of the operation. For trajectory of the Push type:
	//
	// 	- Create
	//
	// 	- Update
	//
	// 	- Delete
	//
	// 	- SetAcl
	//
	// 	- Multi
	//
	// For trajectory of the Pull type:
	//
	// 	- GetData
	//
	// 	- GetChild
	//
	// 	- GetStat
	//
	// example:
	//
	// Create
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// The path.
	//
	// example:
	//
	// /path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The returned result.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 0x301fdfbdbf00***
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The timestamp. It is not available.
	//
	// example:
	//
	// 1669619383000
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The type of the trajectory. Valid values:
	//
	// 	- Push
	//
	// 	- Pull
	//
	// 	- Notify
	//
	// example:
	//
	// Push
	TraceType *string `json:"TraceType,omitempty" xml:"TraceType,omitempty"`
	// The time to live (TTL).
	//
	// example:
	//
	// 0
	Ttl *int64 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// Indicates whether the monitoring feature is enabled.
	//
	// example:
	//
	// true
	Watch *bool `json:"Watch,omitempty" xml:"Watch,omitempty"`
}

func (s ListZkTrackResponseBodyTraces) String() string {
	return dara.Prettify(s)
}

func (s ListZkTrackResponseBodyTraces) GoString() string {
	return s.String()
}

func (s *ListZkTrackResponseBodyTraces) GetAcl() *string {
	return s.Acl
}

func (s *ListZkTrackResponseBodyTraces) GetDataType() *string {
	return s.DataType
}

func (s *ListZkTrackResponseBodyTraces) GetEventType() *string {
	return s.EventType
}

func (s *ListZkTrackResponseBodyTraces) GetFinished() *bool {
	return s.Finished
}

func (s *ListZkTrackResponseBodyTraces) GetLogDate() *string {
	return s.LogDate
}

func (s *ListZkTrackResponseBodyTraces) GetMultiSize() *int64 {
	return s.MultiSize
}

func (s *ListZkTrackResponseBodyTraces) GetOpType() *string {
	return s.OpType
}

func (s *ListZkTrackResponseBodyTraces) GetPath() *string {
	return s.Path
}

func (s *ListZkTrackResponseBodyTraces) GetResult() *string {
	return s.Result
}

func (s *ListZkTrackResponseBodyTraces) GetSessionId() *string {
	return s.SessionId
}

func (s *ListZkTrackResponseBodyTraces) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListZkTrackResponseBodyTraces) GetTraceType() *string {
	return s.TraceType
}

func (s *ListZkTrackResponseBodyTraces) GetTtl() *int64 {
	return s.Ttl
}

func (s *ListZkTrackResponseBodyTraces) GetWatch() *bool {
	return s.Watch
}

func (s *ListZkTrackResponseBodyTraces) SetAcl(v string) *ListZkTrackResponseBodyTraces {
	s.Acl = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetDataType(v string) *ListZkTrackResponseBodyTraces {
	s.DataType = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetEventType(v string) *ListZkTrackResponseBodyTraces {
	s.EventType = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetFinished(v bool) *ListZkTrackResponseBodyTraces {
	s.Finished = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetLogDate(v string) *ListZkTrackResponseBodyTraces {
	s.LogDate = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetMultiSize(v int64) *ListZkTrackResponseBodyTraces {
	s.MultiSize = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetOpType(v string) *ListZkTrackResponseBodyTraces {
	s.OpType = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetPath(v string) *ListZkTrackResponseBodyTraces {
	s.Path = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetResult(v string) *ListZkTrackResponseBodyTraces {
	s.Result = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetSessionId(v string) *ListZkTrackResponseBodyTraces {
	s.SessionId = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetTimestamp(v string) *ListZkTrackResponseBodyTraces {
	s.Timestamp = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetTraceType(v string) *ListZkTrackResponseBodyTraces {
	s.TraceType = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetTtl(v int64) *ListZkTrackResponseBodyTraces {
	s.Ttl = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) SetWatch(v bool) *ListZkTrackResponseBodyTraces {
	s.Watch = &v
	return s
}

func (s *ListZkTrackResponseBodyTraces) Validate() error {
	return dara.Validate(s)
}
