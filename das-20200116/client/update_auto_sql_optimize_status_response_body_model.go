// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoSqlOptimizeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAutoSqlOptimizeStatusResponseBody
	GetCode() *string
	SetData(v *UpdateAutoSqlOptimizeStatusResponseBodyData) *UpdateAutoSqlOptimizeStatusResponseBody
	GetData() *UpdateAutoSqlOptimizeStatusResponseBodyData
	SetMessage(v string) *UpdateAutoSqlOptimizeStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAutoSqlOptimizeStatusResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateAutoSqlOptimizeStatusResponseBody
	GetSuccess() *string
}

type UpdateAutoSqlOptimizeStatusResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UpdateAutoSqlOptimizeStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A52AD37C-35ED-581A-AC23-2232BE54****
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

func (s UpdateAutoSqlOptimizeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoSqlOptimizeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) GetData() *UpdateAutoSqlOptimizeStatusResponseBodyData {
	return s.Data
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) SetCode(v string) *UpdateAutoSqlOptimizeStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) SetData(v *UpdateAutoSqlOptimizeStatusResponseBodyData) *UpdateAutoSqlOptimizeStatusResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) SetMessage(v string) *UpdateAutoSqlOptimizeStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) SetRequestId(v string) *UpdateAutoSqlOptimizeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) SetSuccess(v string) *UpdateAutoSqlOptimizeStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateAutoSqlOptimizeStatusResponseBodyData struct {
	// The error code. Valid values:
	//
	// 	- **-1001**: indicates that the specified parameter is invalid.
	//
	// 	- **-91029**: indicates that a system error occurred.
	//
	// example:
	//
	// -1001
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// invalid param
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Indicates whether the request initiated to configure the automatic SQL optimization feature was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAutoSqlOptimizeStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoSqlOptimizeStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAutoSqlOptimizeStatusResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAutoSqlOptimizeStatusResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateAutoSqlOptimizeStatusResponseBodyData) GetSuccess() *string {
	return s.Success
}

func (s *UpdateAutoSqlOptimizeStatusResponseBodyData) SetErrorCode(v string) *UpdateAutoSqlOptimizeStatusResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponseBodyData) SetErrorMsg(v string) *UpdateAutoSqlOptimizeStatusResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponseBodyData) SetSuccess(v string) *UpdateAutoSqlOptimizeStatusResponseBodyData {
	s.Success = &v
	return s
}

func (s *UpdateAutoSqlOptimizeStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
