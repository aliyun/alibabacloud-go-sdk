// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordPostBackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *RecordPostBackResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *RecordPostBackResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *RecordPostBackResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *RecordPostBackResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *RecordPostBackResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *RecordPostBackResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *RecordPostBackResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *RecordPostBackResponseBody
	GetHttpStatusCode() *int32
	SetModule(v bool) *RecordPostBackResponseBody
	GetModule() *bool
	SetRequestId(v string) *RecordPostBackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RecordPostBackResponseBody
	GetSuccess() *bool
}

type RecordPostBackResponseBody struct {
	// example:
	//
	// false
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// esp-core-aliyun-com
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// PARAMETER.ILLEGAL
	DynamicCode    *string       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// PARAMETER.ILLEGAL
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// True
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// B8E5CC4C-7563-19BD-B02F-DFFFD4E51D4A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RecordPostBackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecordPostBackResponseBody) GoString() string {
	return s.String()
}

func (s *RecordPostBackResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *RecordPostBackResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *RecordPostBackResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *RecordPostBackResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *RecordPostBackResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *RecordPostBackResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RecordPostBackResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RecordPostBackResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RecordPostBackResponseBody) GetModule() *bool {
	return s.Module
}

func (s *RecordPostBackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecordPostBackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RecordPostBackResponseBody) SetAllowRetry(v bool) *RecordPostBackResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *RecordPostBackResponseBody) SetAppName(v string) *RecordPostBackResponseBody {
	s.AppName = &v
	return s
}

func (s *RecordPostBackResponseBody) SetDynamicCode(v string) *RecordPostBackResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RecordPostBackResponseBody) SetDynamicMessage(v string) *RecordPostBackResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RecordPostBackResponseBody) SetErrorArgs(v []interface{}) *RecordPostBackResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *RecordPostBackResponseBody) SetErrorCode(v string) *RecordPostBackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RecordPostBackResponseBody) SetErrorMsg(v string) *RecordPostBackResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RecordPostBackResponseBody) SetHttpStatusCode(v int32) *RecordPostBackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RecordPostBackResponseBody) SetModule(v bool) *RecordPostBackResponseBody {
	s.Module = &v
	return s
}

func (s *RecordPostBackResponseBody) SetRequestId(v string) *RecordPostBackResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecordPostBackResponseBody) SetSuccess(v bool) *RecordPostBackResponseBody {
	s.Success = &v
	return s
}

func (s *RecordPostBackResponseBody) Validate() error {
	return dara.Validate(s)
}
