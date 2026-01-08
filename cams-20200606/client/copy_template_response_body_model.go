// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CopyTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CopyTemplateResponseBody
	GetCode() *string
	SetData(v *CopyTemplateResponseBodyData) *CopyTemplateResponseBody
	GetData() *CopyTemplateResponseBodyData
	SetMessage(v string) *CopyTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CopyTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CopyTemplateResponseBody
	GetSuccess() *bool
}

type CopyTemplateResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CopyTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CopyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CopyTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CopyTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CopyTemplateResponseBody) GetData() *CopyTemplateResponseBodyData {
	return s.Data
}

func (s *CopyTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CopyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CopyTemplateResponseBody) SetAccessDeniedDetail(v string) *CopyTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CopyTemplateResponseBody) SetCode(v string) *CopyTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CopyTemplateResponseBody) SetData(v *CopyTemplateResponseBodyData) *CopyTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CopyTemplateResponseBody) SetMessage(v string) *CopyTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CopyTemplateResponseBody) SetRequestId(v string) *CopyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyTemplateResponseBody) SetSuccess(v bool) *CopyTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CopyTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyTemplateResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	SceneTemplateCode *string `json:"SceneTemplateCode,omitempty" xml:"SceneTemplateCode,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	SceneTemplateName *string `json:"SceneTemplateName,omitempty" xml:"SceneTemplateName,omitempty"`
	// example:
	//
	// 示例值示例值
	WhatsappCatagory *string `json:"WhatsappCatagory,omitempty" xml:"WhatsappCatagory,omitempty"`
}

func (s CopyTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CopyTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CopyTemplateResponseBodyData) GetSceneTemplateCode() *string {
	return s.SceneTemplateCode
}

func (s *CopyTemplateResponseBodyData) GetSceneTemplateName() *string {
	return s.SceneTemplateName
}

func (s *CopyTemplateResponseBodyData) GetWhatsappCatagory() *string {
	return s.WhatsappCatagory
}

func (s *CopyTemplateResponseBodyData) SetSceneTemplateCode(v string) *CopyTemplateResponseBodyData {
	s.SceneTemplateCode = &v
	return s
}

func (s *CopyTemplateResponseBodyData) SetSceneTemplateName(v string) *CopyTemplateResponseBodyData {
	s.SceneTemplateName = &v
	return s
}

func (s *CopyTemplateResponseBodyData) SetWhatsappCatagory(v string) *CopyTemplateResponseBodyData {
	s.WhatsappCatagory = &v
	return s
}

func (s *CopyTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
