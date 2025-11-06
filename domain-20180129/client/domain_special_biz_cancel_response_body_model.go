// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainSpecialBizCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *DomainSpecialBizCancelResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *DomainSpecialBizCancelResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *DomainSpecialBizCancelResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DomainSpecialBizCancelResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *DomainSpecialBizCancelResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *DomainSpecialBizCancelResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *DomainSpecialBizCancelResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *DomainSpecialBizCancelResponseBody
	GetHttpStatusCode() *int32
	SetModule(v interface{}) *DomainSpecialBizCancelResponseBody
	GetModule() interface{}
	SetRequestId(v string) *DomainSpecialBizCancelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DomainSpecialBizCancelResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *DomainSpecialBizCancelResponseBody
	GetSynchro() *bool
}

type DomainSpecialBizCancelResponseBody struct {
	// Indicates whether retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The name of the application for which the error code is returned.
	//
	// example:
	//
	// test-com
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// -
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// -
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The array of error parameters that are returned.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The error code.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 110001
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The HTTP status code that is directly returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned data.
	//
	// example:
	//
	// -
	Module interface{} `json:"Module,omitempty" xml:"Module,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 096DCF79-A89A-5CED-B7DE-1B04761023B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates whether to perform synchronous processing.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DomainSpecialBizCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DomainSpecialBizCancelResponseBody) GoString() string {
	return s.String()
}

func (s *DomainSpecialBizCancelResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *DomainSpecialBizCancelResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DomainSpecialBizCancelResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DomainSpecialBizCancelResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DomainSpecialBizCancelResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *DomainSpecialBizCancelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DomainSpecialBizCancelResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DomainSpecialBizCancelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DomainSpecialBizCancelResponseBody) GetModule() interface{} {
	return s.Module
}

func (s *DomainSpecialBizCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DomainSpecialBizCancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DomainSpecialBizCancelResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *DomainSpecialBizCancelResponseBody) SetAllowRetry(v bool) *DomainSpecialBizCancelResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetAppName(v string) *DomainSpecialBizCancelResponseBody {
	s.AppName = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetDynamicCode(v string) *DomainSpecialBizCancelResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetDynamicMessage(v string) *DomainSpecialBizCancelResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetErrorArgs(v []interface{}) *DomainSpecialBizCancelResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetErrorCode(v string) *DomainSpecialBizCancelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetErrorMsg(v string) *DomainSpecialBizCancelResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetHttpStatusCode(v int32) *DomainSpecialBizCancelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetModule(v interface{}) *DomainSpecialBizCancelResponseBody {
	s.Module = v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetRequestId(v string) *DomainSpecialBizCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetSuccess(v bool) *DomainSpecialBizCancelResponseBody {
	s.Success = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) SetSynchro(v bool) *DomainSpecialBizCancelResponseBody {
	s.Synchro = &v
	return s
}

func (s *DomainSpecialBizCancelResponseBody) Validate() error {
	return dara.Validate(s)
}
