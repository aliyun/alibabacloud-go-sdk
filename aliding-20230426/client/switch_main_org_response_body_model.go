// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchMainOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *SwitchMainOrgResponseBodyContent) *SwitchMainOrgResponseBody
	GetContent() *SwitchMainOrgResponseBodyContent
	SetErrorCode(v string) *SwitchMainOrgResponseBody
	GetErrorCode() *string
	SetErrorCtx(v map[string]interface{}) *SwitchMainOrgResponseBody
	GetErrorCtx() map[string]interface{}
	SetErrorMsg(v string) *SwitchMainOrgResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *SwitchMainOrgResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *SwitchMainOrgResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SwitchMainOrgResponseBody
	GetSuccess() *bool
}

type SwitchMainOrgResponseBody struct {
	Content *SwitchMainOrgResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// 0
	ErrorCode *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorCtx  map[string]interface{} `json:"errorCtx,omitempty" xml:"errorCtx,omitempty"`
	ErrorMsg  *string                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SwitchMainOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgResponseBody) GetContent() *SwitchMainOrgResponseBodyContent {
	return s.Content
}

func (s *SwitchMainOrgResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SwitchMainOrgResponseBody) GetErrorCtx() map[string]interface{} {
	return s.ErrorCtx
}

func (s *SwitchMainOrgResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SwitchMainOrgResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SwitchMainOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchMainOrgResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SwitchMainOrgResponseBody) SetContent(v *SwitchMainOrgResponseBodyContent) *SwitchMainOrgResponseBody {
	s.Content = v
	return s
}

func (s *SwitchMainOrgResponseBody) SetErrorCode(v string) *SwitchMainOrgResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SwitchMainOrgResponseBody) SetErrorCtx(v map[string]interface{}) *SwitchMainOrgResponseBody {
	s.ErrorCtx = v
	return s
}

func (s *SwitchMainOrgResponseBody) SetErrorMsg(v string) *SwitchMainOrgResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SwitchMainOrgResponseBody) SetHttpStatusCode(v int32) *SwitchMainOrgResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SwitchMainOrgResponseBody) SetRequestId(v string) *SwitchMainOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchMainOrgResponseBody) SetSuccess(v bool) *SwitchMainOrgResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchMainOrgResponseBody) Validate() error {
	return dara.Validate(s)
}

type SwitchMainOrgResponseBodyContent struct {
	// example:
	//
	// null
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SwitchMainOrgResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgResponseBodyContent) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgResponseBodyContent) GetData() interface{} {
	return s.Data
}

func (s *SwitchMainOrgResponseBodyContent) SetData(v interface{}) *SwitchMainOrgResponseBodyContent {
	s.Data = v
	return s
}

func (s *SwitchMainOrgResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
