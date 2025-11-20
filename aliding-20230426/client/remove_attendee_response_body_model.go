// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAttendeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *RemoveAttendeeResponseBodyContent) *RemoveAttendeeResponseBody
	GetContent() *RemoveAttendeeResponseBodyContent
	SetErrorCode(v string) *RemoveAttendeeResponseBody
	GetErrorCode() *string
	SetErrorCtx(v map[string]interface{}) *RemoveAttendeeResponseBody
	GetErrorCtx() map[string]interface{}
	SetErrorMsg(v string) *RemoveAttendeeResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *RemoveAttendeeResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RemoveAttendeeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveAttendeeResponseBody
	GetSuccess() *bool
}

type RemoveAttendeeResponseBody struct {
	Content *RemoveAttendeeResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorCtx  map[string]interface{} `json:"ErrorCtx,omitempty" xml:"ErrorCtx,omitempty"`
	ErrorMsg  *string                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// RequestId
	//
	// example:
	//
	// A348BA5D-FFD4-57E4-9450-23A14D72F331
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveAttendeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAttendeeResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeResponseBody) GetContent() *RemoveAttendeeResponseBodyContent {
	return s.Content
}

func (s *RemoveAttendeeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemoveAttendeeResponseBody) GetErrorCtx() map[string]interface{} {
	return s.ErrorCtx
}

func (s *RemoveAttendeeResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RemoveAttendeeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveAttendeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAttendeeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveAttendeeResponseBody) SetContent(v *RemoveAttendeeResponseBodyContent) *RemoveAttendeeResponseBody {
	s.Content = v
	return s
}

func (s *RemoveAttendeeResponseBody) SetErrorCode(v string) *RemoveAttendeeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemoveAttendeeResponseBody) SetErrorCtx(v map[string]interface{}) *RemoveAttendeeResponseBody {
	s.ErrorCtx = v
	return s
}

func (s *RemoveAttendeeResponseBody) SetErrorMsg(v string) *RemoveAttendeeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RemoveAttendeeResponseBody) SetHttpStatusCode(v int32) *RemoveAttendeeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveAttendeeResponseBody) SetRequestId(v string) *RemoveAttendeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAttendeeResponseBody) SetSuccess(v bool) *RemoveAttendeeResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveAttendeeResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveAttendeeResponseBodyContent struct {
	// example:
	//
	// []
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s RemoveAttendeeResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s RemoveAttendeeResponseBodyContent) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeResponseBodyContent) GetData() interface{} {
	return s.Data
}

func (s *RemoveAttendeeResponseBodyContent) SetData(v interface{}) *RemoveAttendeeResponseBodyContent {
	s.Data = v
	return s
}

func (s *RemoveAttendeeResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
