// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDomainSpecialBizCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *SubmitDomainSpecialBizCredentialsResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *SubmitDomainSpecialBizCredentialsResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *SubmitDomainSpecialBizCredentialsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *SubmitDomainSpecialBizCredentialsResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *SubmitDomainSpecialBizCredentialsResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *SubmitDomainSpecialBizCredentialsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SubmitDomainSpecialBizCredentialsResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *SubmitDomainSpecialBizCredentialsResponseBody
	GetHttpStatusCode() *int32
	SetModule(v interface{}) *SubmitDomainSpecialBizCredentialsResponseBody
	GetModule() interface{}
	SetRequestId(v string) *SubmitDomainSpecialBizCredentialsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitDomainSpecialBizCredentialsResponseBody
	GetSuccess() *bool
	SetSynchro(v bool) *SubmitDomainSpecialBizCredentialsResponseBody
	GetSynchro() *bool
}

type SubmitDomainSpecialBizCredentialsResponseBody struct {
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
	// A83E1D74-E46B-505C-947A-8C6B7E41C011
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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

func (s SubmitDomainSpecialBizCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDomainSpecialBizCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetModule() interface{} {
	return s.Module
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetAllowRetry(v bool) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetAppName(v string) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.AppName = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetDynamicCode(v string) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetDynamicMessage(v string) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetErrorArgs(v []interface{}) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetErrorCode(v string) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetErrorMsg(v string) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetHttpStatusCode(v int32) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetModule(v interface{}) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.Module = v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetRequestId(v string) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetSuccess(v bool) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) SetSynchro(v bool) *SubmitDomainSpecialBizCredentialsResponseBody {
	s.Synchro = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsResponseBody) Validate() error {
	return dara.Validate(s)
}
