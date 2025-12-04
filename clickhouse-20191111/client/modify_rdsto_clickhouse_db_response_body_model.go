// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRDSToClickhouseDbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *ModifyRDSToClickhouseDbResponseBody
	GetErrorCode() *int64
	SetErrorMsg(v string) *ModifyRDSToClickhouseDbResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ModifyRDSToClickhouseDbResponseBody
	GetRequestId() *string
	SetStatus(v int64) *ModifyRDSToClickhouseDbResponseBody
	GetStatus() *int64
}

type ModifyRDSToClickhouseDbResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 	- If the value **1*	- is returned for the **Status*	- parameter, the system does not return the ErrorMsg parameter.
	//
	// 	- If the value **0*	- is returned for the **Status*	- parameter, the ErrorMsg parameter returns the cause for the modification failure.
	//
	// example:
	//
	// ClickHouse exception, code: 49, host: 100.100.118.132, port: 49670; Code: 49, e.displayText() = DB::Exception: Logical error: there is no global context (version 20.8.17.25)n
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 746CD303-0B82-5E8D-886D-93A9FAF3A876
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the modification was successful. Valid values:
	//
	// 	- **1**: The modification was successful.
	//
	// 	- **0**: The modification failed.
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyRDSToClickhouseDbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRDSToClickhouseDbResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRDSToClickhouseDbResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *ModifyRDSToClickhouseDbResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ModifyRDSToClickhouseDbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRDSToClickhouseDbResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *ModifyRDSToClickhouseDbResponseBody) SetErrorCode(v int64) *ModifyRDSToClickhouseDbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponseBody) SetErrorMsg(v string) *ModifyRDSToClickhouseDbResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponseBody) SetRequestId(v string) *ModifyRDSToClickhouseDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponseBody) SetStatus(v int64) *ModifyRDSToClickhouseDbResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyRDSToClickhouseDbResponseBody) Validate() error {
	return dara.Validate(s)
}
