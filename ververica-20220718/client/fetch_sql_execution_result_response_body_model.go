// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchSqlExecutionResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SqlExecutionFetchResult) *FetchSqlExecutionResultResponseBody
	GetData() *SqlExecutionFetchResult
	SetErrorCode(v string) *FetchSqlExecutionResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *FetchSqlExecutionResultResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *FetchSqlExecutionResultResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *FetchSqlExecutionResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchSqlExecutionResultResponseBody
	GetSuccess() *bool
}

type FetchSqlExecutionResultResponseBody struct {
	// The execution result details returned when the request is successful.
	//
	// example:
	//
	// { "jobs": [ { "jid": "4df35f8e54554b23bf7dcd38a151****", "name": "69d001d5-419a-4bfc-9c2e-849cacd3****", "state": "RUNNING", "start-time": 1659154942068, "end-time": -1, "duration": 188161756, "last-modification": 1659154968305, "tasks": { "total": 2, "created": 0, "scheduled": 0, "deploying": 0, "running": 2, "finished": 0, "canceling": 0, "canceled": 0, "failed": 0, "reconciling": 0, "initializing": 0 } } ] }
	Data *SqlExecutionFetchResult `json:"data,omitempty" xml:"data,omitempty"`
	// The business error code. This value is not empty when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business error message. This value is not empty when success is false. This value is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code. The value is always 200. Use success to determine whether the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the business request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FetchSqlExecutionResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchSqlExecutionResultResponseBody) GoString() string {
	return s.String()
}

func (s *FetchSqlExecutionResultResponseBody) GetData() *SqlExecutionFetchResult {
	return s.Data
}

func (s *FetchSqlExecutionResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FetchSqlExecutionResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *FetchSqlExecutionResultResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *FetchSqlExecutionResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchSqlExecutionResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchSqlExecutionResultResponseBody) SetData(v *SqlExecutionFetchResult) *FetchSqlExecutionResultResponseBody {
	s.Data = v
	return s
}

func (s *FetchSqlExecutionResultResponseBody) SetErrorCode(v string) *FetchSqlExecutionResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FetchSqlExecutionResultResponseBody) SetErrorMessage(v string) *FetchSqlExecutionResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FetchSqlExecutionResultResponseBody) SetHttpCode(v int32) *FetchSqlExecutionResultResponseBody {
	s.HttpCode = &v
	return s
}

func (s *FetchSqlExecutionResultResponseBody) SetRequestId(v string) *FetchSqlExecutionResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchSqlExecutionResultResponseBody) SetSuccess(v bool) *FetchSqlExecutionResultResponseBody {
	s.Success = &v
	return s
}

func (s *FetchSqlExecutionResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
