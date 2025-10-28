// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeExecErrorSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeExecErrorSampleResponseBody
	GetCode() *string
	SetData(v *GetQueryOptimizeExecErrorSampleResponseBodyData) *GetQueryOptimizeExecErrorSampleResponseBody
	GetData() *GetQueryOptimizeExecErrorSampleResponseBodyData
	SetMessage(v string) *GetQueryOptimizeExecErrorSampleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeExecErrorSampleResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeExecErrorSampleResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeExecErrorSampleResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *GetQueryOptimizeExecErrorSampleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
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

func (s GetQueryOptimizeExecErrorSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) GetData() *GetQueryOptimizeExecErrorSampleResponseBodyData {
	return s.Data
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetCode(v string) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetData(v *GetQueryOptimizeExecErrorSampleResponseBodyData) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetMessage(v string) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetRequestId(v string) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) SetSuccess(v string) *GetQueryOptimizeExecErrorSampleResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueryOptimizeExecErrorSampleResponseBodyData struct {
	// A reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The queried data.
	List []*GetQueryOptimizeExecErrorSampleResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeExecErrorSampleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) GetList() []*GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	return s.List
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetExtra(v string) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetList(v []*GetQueryOptimizeExecErrorSampleResponseBodyDataList) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) SetTotal(v int64) *GetQueryOptimizeExecErrorSampleResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyData) Validate() error {
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

type GetQueryOptimizeExecErrorSampleResponseBodyDataList struct {
	// The name of the database.
	//
	// example:
	//
	// testdb01
	Dbname *string `json:"Dbname,omitempty" xml:"Dbname,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// 1146
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The IP address of the client that executes the SQL statement.
	//
	// example:
	//
	// 100.104.XX.XX
	OrigHost *string `json:"OrigHost,omitempty" xml:"OrigHost,omitempty"`
	// The SQL template ID.
	//
	// example:
	//
	// 2e8147b5ca2dfc640dfd5e43d96a****
	SqlId *string `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The content of the SQL statement that failed to be executed.
	//
	// example:
	//
	// select 	- from test1
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// The point in time when the failed SQL statement was executed. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1643020306739
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The username of the client that executes the SQL statement.
	//
	// example:
	//
	// test01
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetQueryOptimizeExecErrorSampleResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeExecErrorSampleResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) GetDbname() *string {
	return s.Dbname
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) GetOrigHost() *string {
	return s.OrigHost
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) GetSqlId() *string {
	return s.SqlId
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) GetSqlText() *string {
	return s.SqlText
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) GetUser() *string {
	return s.User
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetDbname(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.Dbname = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetErrorCode(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.ErrorCode = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetOrigHost(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.OrigHost = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetSqlId(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetSqlText(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.SqlText = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetTimestamp(v int64) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.Timestamp = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) SetUser(v string) *GetQueryOptimizeExecErrorSampleResponseBodyDataList {
	s.User = &v
	return s
}

func (s *GetQueryOptimizeExecErrorSampleResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
