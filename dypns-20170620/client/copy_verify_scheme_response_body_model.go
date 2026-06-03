// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyVerifySchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CopyVerifySchemeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CopyVerifySchemeResponseBody
	GetCode() *string
	SetData(v *CopyVerifySchemeResponseBodyData) *CopyVerifySchemeResponseBody
	GetData() *CopyVerifySchemeResponseBodyData
	SetMessage(v string) *CopyVerifySchemeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CopyVerifySchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CopyVerifySchemeResponseBody
	GetSuccess() *bool
}

type CopyVerifySchemeResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// 错误码
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 结果数据
	Data *CopyVerifySchemeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误消息
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2E7CA885-8D97-C497-A02C-2D31214D3285
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CopyVerifySchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CopyVerifySchemeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CopyVerifySchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CopyVerifySchemeResponseBody) GetData() *CopyVerifySchemeResponseBodyData {
	return s.Data
}

func (s *CopyVerifySchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CopyVerifySchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyVerifySchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CopyVerifySchemeResponseBody) SetAccessDeniedDetail(v string) *CopyVerifySchemeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CopyVerifySchemeResponseBody) SetCode(v string) *CopyVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CopyVerifySchemeResponseBody) SetData(v *CopyVerifySchemeResponseBodyData) *CopyVerifySchemeResponseBody {
	s.Data = v
	return s
}

func (s *CopyVerifySchemeResponseBody) SetMessage(v string) *CopyVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CopyVerifySchemeResponseBody) SetRequestId(v string) *CopyVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyVerifySchemeResponseBody) SetSuccess(v bool) *CopyVerifySchemeResponseBody {
	s.Success = &v
	return s
}

func (s *CopyVerifySchemeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyVerifySchemeResponseBodyData struct {
	// 方案Code
	//
	// example:
	//
	// FC100********212
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s CopyVerifySchemeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CopyVerifySchemeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CopyVerifySchemeResponseBodyData) GetSceneCode() *string {
	return s.SceneCode
}

func (s *CopyVerifySchemeResponseBodyData) SetSceneCode(v string) *CopyVerifySchemeResponseBodyData {
	s.SceneCode = &v
	return s
}

func (s *CopyVerifySchemeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
