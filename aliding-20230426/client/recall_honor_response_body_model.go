// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallHonorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *RecallHonorResponseBodyContent) *RecallHonorResponseBody
	GetContent() *RecallHonorResponseBodyContent
	SetErrorCode(v string) *RecallHonorResponseBody
	GetErrorCode() *string
	SetErrorCtx(v map[string]interface{}) *RecallHonorResponseBody
	GetErrorCtx() map[string]interface{}
	SetErrorMsg(v string) *RecallHonorResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *RecallHonorResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RecallHonorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RecallHonorResponseBody
	GetSuccess() *bool
}

type RecallHonorResponseBody struct {
	Content *RecallHonorResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// 040008
	ErrorCode *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorCtx  map[string]interface{} `json:"errorCtx,omitempty" xml:"errorCtx,omitempty"`
	ErrorMsg  *string                `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RecallHonorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorResponseBody) GoString() string {
	return s.String()
}

func (s *RecallHonorResponseBody) GetContent() *RecallHonorResponseBodyContent {
	return s.Content
}

func (s *RecallHonorResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RecallHonorResponseBody) GetErrorCtx() map[string]interface{} {
	return s.ErrorCtx
}

func (s *RecallHonorResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RecallHonorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RecallHonorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecallHonorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RecallHonorResponseBody) SetContent(v *RecallHonorResponseBodyContent) *RecallHonorResponseBody {
	s.Content = v
	return s
}

func (s *RecallHonorResponseBody) SetErrorCode(v string) *RecallHonorResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RecallHonorResponseBody) SetErrorCtx(v map[string]interface{}) *RecallHonorResponseBody {
	s.ErrorCtx = v
	return s
}

func (s *RecallHonorResponseBody) SetErrorMsg(v string) *RecallHonorResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RecallHonorResponseBody) SetHttpStatusCode(v int32) *RecallHonorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RecallHonorResponseBody) SetRequestId(v string) *RecallHonorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecallHonorResponseBody) SetSuccess(v bool) *RecallHonorResponseBody {
	s.Success = &v
	return s
}

func (s *RecallHonorResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecallHonorResponseBodyContent struct {
	// example:
	//
	// success
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s RecallHonorResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorResponseBodyContent) GoString() string {
	return s.String()
}

func (s *RecallHonorResponseBodyContent) GetData() interface{} {
	return s.Data
}

func (s *RecallHonorResponseBodyContent) SetData(v interface{}) *RecallHonorResponseBodyContent {
	s.Data = v
	return s
}

func (s *RecallHonorResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
