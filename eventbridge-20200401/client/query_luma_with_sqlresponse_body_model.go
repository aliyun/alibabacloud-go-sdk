// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLumaWithSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryLumaWithSQLResponseBody
	GetCode() *string
	SetData(v *ExecutionResult) *QueryLumaWithSQLResponseBody
	GetData() *ExecutionResult
	SetMessage(v string) *QueryLumaWithSQLResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryLumaWithSQLResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryLumaWithSQLResponseBody
	GetSuccess() *bool
}

type QueryLumaWithSQLResponseBody struct {
	// The response code of the operation. A value of Success indicates that the call succeeded. Otherwise, a specific error code is returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result of the SQL query, including column definitions and row data.
	Data *ExecutionResult `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. Use this ID for troubleshooting and when submitting a ticket.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryLumaWithSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLumaWithSQLResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLumaWithSQLResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryLumaWithSQLResponseBody) GetData() *ExecutionResult {
	return s.Data
}

func (s *QueryLumaWithSQLResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryLumaWithSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLumaWithSQLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryLumaWithSQLResponseBody) SetCode(v string) *QueryLumaWithSQLResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLumaWithSQLResponseBody) SetData(v *ExecutionResult) *QueryLumaWithSQLResponseBody {
	s.Data = v
	return s
}

func (s *QueryLumaWithSQLResponseBody) SetMessage(v string) *QueryLumaWithSQLResponseBody {
	s.Message = &v
	return s
}

func (s *QueryLumaWithSQLResponseBody) SetRequestId(v string) *QueryLumaWithSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLumaWithSQLResponseBody) SetSuccess(v bool) *QueryLumaWithSQLResponseBody {
	s.Success = &v
	return s
}

func (s *QueryLumaWithSQLResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
