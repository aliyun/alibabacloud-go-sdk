// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAttendeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *AddAttendeeResponseBodyContent) *AddAttendeeResponseBody
	GetContent() *AddAttendeeResponseBodyContent
	SetErrorCode(v string) *AddAttendeeResponseBody
	GetErrorCode() *string
	SetErrorCtx(v map[string]interface{}) *AddAttendeeResponseBody
	GetErrorCtx() map[string]interface{}
	SetErrorMsg(v string) *AddAttendeeResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *AddAttendeeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *AddAttendeeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAttendeeResponseBody
	GetSuccess() *bool
}

type AddAttendeeResponseBody struct {
	Content *AddAttendeeResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// 200
	ErrorCode *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorCtx  map[string]interface{} `json:"ErrorCtx,omitempty" xml:"ErrorCtx,omitempty"`
	// example:
	//
	// error check permissions
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// RequestId
	//
	// example:
	//
	// 9BCC17ED-0187-54A0-BD31-56FDBE865447
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAttendeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeResponseBody) GoString() string {
	return s.String()
}

func (s *AddAttendeeResponseBody) GetContent() *AddAttendeeResponseBodyContent {
	return s.Content
}

func (s *AddAttendeeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddAttendeeResponseBody) GetErrorCtx() map[string]interface{} {
	return s.ErrorCtx
}

func (s *AddAttendeeResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *AddAttendeeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddAttendeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAttendeeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAttendeeResponseBody) SetContent(v *AddAttendeeResponseBodyContent) *AddAttendeeResponseBody {
	s.Content = v
	return s
}

func (s *AddAttendeeResponseBody) SetErrorCode(v string) *AddAttendeeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddAttendeeResponseBody) SetErrorCtx(v map[string]interface{}) *AddAttendeeResponseBody {
	s.ErrorCtx = v
	return s
}

func (s *AddAttendeeResponseBody) SetErrorMsg(v string) *AddAttendeeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddAttendeeResponseBody) SetHttpStatusCode(v int32) *AddAttendeeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddAttendeeResponseBody) SetRequestId(v string) *AddAttendeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAttendeeResponseBody) SetSuccess(v bool) *AddAttendeeResponseBody {
	s.Success = &v
	return s
}

func (s *AddAttendeeResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAttendeeResponseBodyContent struct {
	// example:
	//
	// []
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s AddAttendeeResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s AddAttendeeResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AddAttendeeResponseBodyContent) GetData() interface{} {
	return s.Data
}

func (s *AddAttendeeResponseBodyContent) SetData(v interface{}) *AddAttendeeResponseBodyContent {
	s.Data = v
	return s
}

func (s *AddAttendeeResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
