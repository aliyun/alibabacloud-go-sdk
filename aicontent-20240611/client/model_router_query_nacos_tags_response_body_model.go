// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ModelRouterQueryNacosTagsResponseBodyData) *ModelRouterQueryNacosTagsResponseBody
	GetData() []*ModelRouterQueryNacosTagsResponseBodyData
	SetErrCode(v string) *ModelRouterQueryNacosTagsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryNacosTagsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryNacosTagsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryNacosTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryNacosTagsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryNacosTagsResponseBody struct {
	// example:
	//
	// []
	Data []*ModelRouterQueryNacosTagsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryNacosTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetData() []*ModelRouterQueryNacosTagsResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetData(v []*ModelRouterQueryNacosTagsResponseBodyData) *ModelRouterQueryNacosTagsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetErrCode(v string) *ModelRouterQueryNacosTagsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetErrMessage(v string) *ModelRouterQueryNacosTagsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryNacosTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetRequestId(v string) *ModelRouterQueryNacosTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetSuccess(v bool) *ModelRouterQueryNacosTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModelRouterQueryNacosTagsResponseBodyData struct {
	// example:
	//
	// 文本生成
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Tag     *string `json:"tag,omitempty" xml:"tag,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// example:
	//
	// NLP
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ModelRouterQueryNacosTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) GetTagName() *string {
	return s.TagName
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) SetLabel(v string) *ModelRouterQueryNacosTagsResponseBodyData {
	s.Label = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) SetTag(v string) *ModelRouterQueryNacosTagsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) SetTagName(v string) *ModelRouterQueryNacosTagsResponseBodyData {
	s.TagName = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) SetValue(v string) *ModelRouterQueryNacosTagsResponseBodyData {
	s.Value = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
