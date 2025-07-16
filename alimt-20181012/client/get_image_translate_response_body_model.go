// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTranslateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetImageTranslateResponseBody
	GetCode() *int32
	SetData(v *GetImageTranslateResponseBodyData) *GetImageTranslateResponseBody
	GetData() *GetImageTranslateResponseBodyData
	SetMessage(v string) *GetImageTranslateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageTranslateResponseBody
	GetRequestId() *string
}

type GetImageTranslateResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetImageTranslateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A41F6E25-8520-4AF0-90EF-AF7E32840108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageTranslateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageTranslateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetImageTranslateResponseBody) GetData() *GetImageTranslateResponseBodyData {
	return s.Data
}

func (s *GetImageTranslateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageTranslateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageTranslateResponseBody) SetCode(v int32) *GetImageTranslateResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageTranslateResponseBody) SetData(v *GetImageTranslateResponseBodyData) *GetImageTranslateResponseBody {
	s.Data = v
	return s
}

func (s *GetImageTranslateResponseBody) SetMessage(v string) *GetImageTranslateResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageTranslateResponseBody) SetRequestId(v string) *GetImageTranslateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageTranslateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetImageTranslateResponseBodyData struct {
	Orc           *string `json:"Orc,omitempty" xml:"Orc,omitempty"`
	PictureEditor *string `json:"PictureEditor,omitempty" xml:"PictureEditor,omitempty"`
	// example:
	//
	// https://ae01.alicdn.com/kf/xxxxx.jpeg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImageTranslateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageTranslateResponseBodyData) GetOrc() *string {
	return s.Orc
}

func (s *GetImageTranslateResponseBodyData) GetPictureEditor() *string {
	return s.PictureEditor
}

func (s *GetImageTranslateResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetImageTranslateResponseBodyData) SetOrc(v string) *GetImageTranslateResponseBodyData {
	s.Orc = &v
	return s
}

func (s *GetImageTranslateResponseBodyData) SetPictureEditor(v string) *GetImageTranslateResponseBodyData {
	s.PictureEditor = &v
	return s
}

func (s *GetImageTranslateResponseBodyData) SetUrl(v string) *GetImageTranslateResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetImageTranslateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
