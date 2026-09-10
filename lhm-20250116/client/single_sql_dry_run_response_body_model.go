// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSqlDryRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SingleSqlDryRunResponseBody
	GetData() *string
	SetErrCode(v string) *SingleSqlDryRunResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SingleSqlDryRunResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SingleSqlDryRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SingleSqlDryRunResponseBody
	GetSuccess() *bool
}

type SingleSqlDryRunResponseBody struct {
	// The business data returned by the operation (in string format). The specific content varies by operation.
	//
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code.
	//
	// example:
	//
	// None
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// not supported.pos 14327, line 452, column 10, token IDENTIFIER dialect_type
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CFD7C81E-1A53-5C7C-846D-381BC1C4385F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SingleSqlDryRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SingleSqlDryRunResponseBody) GoString() string {
	return s.String()
}

func (s *SingleSqlDryRunResponseBody) GetData() *string {
	return s.Data
}

func (s *SingleSqlDryRunResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SingleSqlDryRunResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SingleSqlDryRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SingleSqlDryRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SingleSqlDryRunResponseBody) SetData(v string) *SingleSqlDryRunResponseBody {
	s.Data = &v
	return s
}

func (s *SingleSqlDryRunResponseBody) SetErrCode(v string) *SingleSqlDryRunResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SingleSqlDryRunResponseBody) SetErrMessage(v string) *SingleSqlDryRunResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SingleSqlDryRunResponseBody) SetRequestId(v string) *SingleSqlDryRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleSqlDryRunResponseBody) SetSuccess(v bool) *SingleSqlDryRunResponseBody {
	s.Success = &v
	return s
}

func (s *SingleSqlDryRunResponseBody) Validate() error {
	return dara.Validate(s)
}
