// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateUploadConfigResponseBody
	GetCode() *string
	SetData(v *GenerateUploadConfigResponseBodyData) *GenerateUploadConfigResponseBody
	GetData() *GenerateUploadConfigResponseBodyData
	SetHttpStatusCode(v int32) *GenerateUploadConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateUploadConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateUploadConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateUploadConfigResponseBody
	GetSuccess() *bool
}

type GenerateUploadConfigResponseBody struct {
	// example:
	//
	// NoData
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GenerateUploadConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateUploadConfigResponseBody) GetData() *GenerateUploadConfigResponseBodyData {
	return s.Data
}

func (s *GenerateUploadConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateUploadConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateUploadConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateUploadConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateUploadConfigResponseBody) SetCode(v string) *GenerateUploadConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetData(v *GenerateUploadConfigResponseBodyData) *GenerateUploadConfigResponseBody {
	s.Data = v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetHttpStatusCode(v int32) *GenerateUploadConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetMessage(v string) *GenerateUploadConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetRequestId(v string) *GenerateUploadConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetSuccess(v bool) *GenerateUploadConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateUploadConfigResponseBodyData struct {
	// example:
	//
	// oss://default/oss-bucket-name/aimiaobi/2021/07/01/1625126400000/1.docx
	FileKey   *string            `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FormDatas map[string]*string `json:"FormDatas,omitempty" xml:"FormDatas,omitempty"`
	// example:
	//
	// https://bucket-name.oss-cn-hangzhou.aliyuncs.com
	PostUrl *string `json:"PostUrl,omitempty" xml:"PostUrl,omitempty"`
}

func (s GenerateUploadConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponseBodyData) GetFileKey() *string {
	return s.FileKey
}

func (s *GenerateUploadConfigResponseBodyData) GetFormDatas() map[string]*string {
	return s.FormDatas
}

func (s *GenerateUploadConfigResponseBodyData) GetPostUrl() *string {
	return s.PostUrl
}

func (s *GenerateUploadConfigResponseBodyData) SetFileKey(v string) *GenerateUploadConfigResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *GenerateUploadConfigResponseBodyData) SetFormDatas(v map[string]*string) *GenerateUploadConfigResponseBodyData {
	s.FormDatas = v
	return s
}

func (s *GenerateUploadConfigResponseBodyData) SetPostUrl(v string) *GenerateUploadConfigResponseBodyData {
	s.PostUrl = &v
	return s
}

func (s *GenerateUploadConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
