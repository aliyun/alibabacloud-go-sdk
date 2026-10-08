// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthorizedResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *ListUserAuthorizedResourcesResponseBodyContent) *ListUserAuthorizedResourcesResponseBody
	GetContent() *ListUserAuthorizedResourcesResponseBodyContent
	SetErrorCode(v string) *ListUserAuthorizedResourcesResponseBody
	GetErrorCode() *string
	SetErrorCtx(v map[string]interface{}) *ListUserAuthorizedResourcesResponseBody
	GetErrorCtx() map[string]interface{}
	SetErrorMsg(v string) *ListUserAuthorizedResourcesResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *ListUserAuthorizedResourcesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListUserAuthorizedResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserAuthorizedResourcesResponseBody
	GetSuccess() *bool
}

type ListUserAuthorizedResourcesResponseBody struct {
	Content        *ListUserAuthorizedResourcesResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	ErrorCode      *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorCtx       map[string]interface{}                          `json:"ErrorCtx,omitempty" xml:"ErrorCtx,omitempty"`
	ErrorMsg       *string                                         `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpStatusCode *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserAuthorizedResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthorizedResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAuthorizedResourcesResponseBody) GetContent() *ListUserAuthorizedResourcesResponseBodyContent {
	return s.Content
}

func (s *ListUserAuthorizedResourcesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListUserAuthorizedResourcesResponseBody) GetErrorCtx() map[string]interface{} {
	return s.ErrorCtx
}

func (s *ListUserAuthorizedResourcesResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListUserAuthorizedResourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUserAuthorizedResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserAuthorizedResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserAuthorizedResourcesResponseBody) SetContent(v *ListUserAuthorizedResourcesResponseBodyContent) *ListUserAuthorizedResourcesResponseBody {
	s.Content = v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBody) SetErrorCode(v string) *ListUserAuthorizedResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBody) SetErrorCtx(v map[string]interface{}) *ListUserAuthorizedResourcesResponseBody {
	s.ErrorCtx = v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBody) SetErrorMsg(v string) *ListUserAuthorizedResourcesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBody) SetHttpStatusCode(v int32) *ListUserAuthorizedResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBody) SetRequestId(v string) *ListUserAuthorizedResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBody) SetSuccess(v bool) *ListUserAuthorizedResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserAuthorizedResourcesResponseBodyContent struct {
	Data     interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
}

func (s ListUserAuthorizedResourcesResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthorizedResourcesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListUserAuthorizedResourcesResponseBodyContent) GetData() interface{} {
	return s.Data
}

func (s *ListUserAuthorizedResourcesResponseBodyContent) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ListUserAuthorizedResourcesResponseBodyContent) SetData(v interface{}) *ListUserAuthorizedResourcesResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBodyContent) SetMetadata(v interface{}) *ListUserAuthorizedResourcesResponseBodyContent {
	s.Metadata = v
	return s
}

func (s *ListUserAuthorizedResourcesResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
