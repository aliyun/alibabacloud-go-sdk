// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBTaskSQLJobLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDBTaskSQLJobLogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDBTaskSQLJobLogResponseBody
	GetErrorMessage() *string
	SetLog(v string) *GetDBTaskSQLJobLogResponseBody
	GetLog() *string
	SetRequestId(v string) *GetDBTaskSQLJobLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDBTaskSQLJobLogResponseBody
	GetSuccess() *bool
}

type GetDBTaskSQLJobLogResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// MissingJobId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// JobId is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The log that records the scheduling details.
	//
	// example:
	//
	// log_****
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4AC23904-55DE-550B-9676-E8946F07****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDBTaskSQLJobLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDBTaskSQLJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBTaskSQLJobLogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDBTaskSQLJobLogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDBTaskSQLJobLogResponseBody) GetLog() *string {
	return s.Log
}

func (s *GetDBTaskSQLJobLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDBTaskSQLJobLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDBTaskSQLJobLogResponseBody) SetErrorCode(v string) *GetDBTaskSQLJobLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDBTaskSQLJobLogResponseBody) SetErrorMessage(v string) *GetDBTaskSQLJobLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDBTaskSQLJobLogResponseBody) SetLog(v string) *GetDBTaskSQLJobLogResponseBody {
	s.Log = &v
	return s
}

func (s *GetDBTaskSQLJobLogResponseBody) SetRequestId(v string) *GetDBTaskSQLJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBTaskSQLJobLogResponseBody) SetSuccess(v bool) *GetDBTaskSQLJobLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetDBTaskSQLJobLogResponseBody) Validate() error {
	return dara.Validate(s)
}
