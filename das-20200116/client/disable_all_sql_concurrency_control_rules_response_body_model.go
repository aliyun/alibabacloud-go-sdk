// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAllSqlConcurrencyControlRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableAllSqlConcurrencyControlRulesResponseBody
	GetCode() *string
	SetData(v string) *DisableAllSqlConcurrencyControlRulesResponseBody
	GetData() *string
	SetMessage(v string) *DisableAllSqlConcurrencyControlRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableAllSqlConcurrencyControlRulesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DisableAllSqlConcurrencyControlRulesResponseBody
	GetSuccess() *string
}

type DisableAllSqlConcurrencyControlRulesResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message that contains information such as an error code is returned.
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

func (s DisableAllSqlConcurrencyControlRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAllSqlConcurrencyControlRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) GetData() *string {
	return s.Data
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetCode(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetData(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.Data = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetMessage(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetRequestId(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) SetSuccess(v string) *DisableAllSqlConcurrencyControlRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
