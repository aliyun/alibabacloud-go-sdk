// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DeleteEventResponseBodyContent) *DeleteEventResponseBody
	GetContent() *DeleteEventResponseBodyContent
	SetErrorCode(v string) *DeleteEventResponseBody
	GetErrorCode() *string
	SetErrorCtx(v map[string]interface{}) *DeleteEventResponseBody
	GetErrorCtx() map[string]interface{}
	SetErrorMsg(v string) *DeleteEventResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *DeleteEventResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DeleteEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventResponseBody
	GetSuccess() *bool
}

type DeleteEventResponseBody struct {
	Content *DeleteEventResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string                `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorCtx  map[string]interface{} `json:"errorCtx,omitempty" xml:"errorCtx,omitempty"`
	// example:
	//
	// ""
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// requestId
	//
	// example:
	//
	// 4248DCC9-785F-5A14-8BE0-830FD52E1261
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventResponseBody) GetContent() *DeleteEventResponseBodyContent {
	return s.Content
}

func (s *DeleteEventResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteEventResponseBody) GetErrorCtx() map[string]interface{} {
	return s.ErrorCtx
}

func (s *DeleteEventResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteEventResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventResponseBody) SetContent(v *DeleteEventResponseBodyContent) *DeleteEventResponseBody {
	s.Content = v
	return s
}

func (s *DeleteEventResponseBody) SetErrorCode(v string) *DeleteEventResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteEventResponseBody) SetErrorCtx(v map[string]interface{}) *DeleteEventResponseBody {
	s.ErrorCtx = v
	return s
}

func (s *DeleteEventResponseBody) SetErrorMsg(v string) *DeleteEventResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteEventResponseBody) SetHttpStatusCode(v int32) *DeleteEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteEventResponseBody) SetRequestId(v string) *DeleteEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventResponseBody) SetSuccess(v bool) *DeleteEventResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteEventResponseBodyContent struct {
	// example:
	//
	// []
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeleteEventResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DeleteEventResponseBodyContent) GetData() interface{} {
	return s.Data
}

func (s *DeleteEventResponseBodyContent) SetData(v interface{}) *DeleteEventResponseBodyContent {
	s.Data = v
	return s
}

func (s *DeleteEventResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
