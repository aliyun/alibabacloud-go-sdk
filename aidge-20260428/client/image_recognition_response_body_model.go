// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRecognitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageRecognitionResponseBody
	GetCode() *string
	SetData(v *ImageRecognitionResponseBodyData) *ImageRecognitionResponseBody
	GetData() *ImageRecognitionResponseBodyData
	SetMessage(v string) *ImageRecognitionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageRecognitionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageRecognitionResponseBody
	GetSuccess() *bool
}

type ImageRecognitionResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The intelligent element recognition result.
	Data *ImageRecognitionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageRecognitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionResponseBody) GoString() string {
	return s.String()
}

func (s *ImageRecognitionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageRecognitionResponseBody) GetData() *ImageRecognitionResponseBodyData {
	return s.Data
}

func (s *ImageRecognitionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageRecognitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageRecognitionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageRecognitionResponseBody) SetCode(v string) *ImageRecognitionResponseBody {
	s.Code = &v
	return s
}

func (s *ImageRecognitionResponseBody) SetData(v *ImageRecognitionResponseBodyData) *ImageRecognitionResponseBody {
	s.Data = v
	return s
}

func (s *ImageRecognitionResponseBody) SetMessage(v string) *ImageRecognitionResponseBody {
	s.Message = &v
	return s
}

func (s *ImageRecognitionResponseBody) SetRequestId(v string) *ImageRecognitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageRecognitionResponseBody) SetSuccess(v bool) *ImageRecognitionResponseBody {
	s.Success = &v
	return s
}

func (s *ImageRecognitionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageRecognitionResponseBodyData struct {
	// The border pixel information, returned as a comma-separated string.
	//
	// example:
	//
	// 8,10,0,1,330,330
	BorderPixel *string `json:"BorderPixel,omitempty" xml:"BorderPixel,omitempty"`
	// Indicates whether the non-subject area contains text.
	//
	// example:
	//
	// true
	NoobjCharacter *bool `json:"NoobjCharacter,omitempty" xml:"NoobjCharacter,omitempty"`
	// Indicates whether the non-subject area contains a logo.
	//
	// example:
	//
	// false
	NoobjLogo *bool `json:"NoobjLogo,omitempty" xml:"NoobjLogo,omitempty"`
	// Indicates whether the non-subject area contains irrelevant pixels or noise.
	//
	// example:
	//
	// false
	NoobjNpx *bool `json:"NoobjNpx,omitempty" xml:"NoobjNpx,omitempty"`
	// Indicates whether the non-subject area contains a watermark.
	//
	// example:
	//
	// false
	NoobjWatermark *bool `json:"NoobjWatermark,omitempty" xml:"NoobjWatermark,omitempty"`
	// Indicates whether the subject area contains text.
	//
	// example:
	//
	// true
	ObjCharacter *bool `json:"ObjCharacter,omitempty" xml:"ObjCharacter,omitempty"`
	// Indicates whether the subject area contains a logo.
	//
	// example:
	//
	// false
	ObjLogo *bool `json:"ObjLogo,omitempty" xml:"ObjLogo,omitempty"`
	// Indicates whether the subject area contains irrelevant pixels or noise.
	//
	// example:
	//
	// false
	ObjNpx *bool `json:"ObjNpx,omitempty" xml:"ObjNpx,omitempty"`
	// Indicates whether the subject area contains a watermark.
	//
	// example:
	//
	// false
	ObjWatermark *bool `json:"ObjWatermark,omitempty" xml:"ObjWatermark,omitempty"`
	// The product count.
	//
	// example:
	//
	// 2
	PdNum *int32 `json:"PdNum,omitempty" xml:"PdNum,omitempty"`
	// The product proportion.
	//
	// example:
	//
	// 74.15%
	PdProp *string `json:"PdProp,omitempty" xml:"PdProp,omitempty"`
	// The list of recognized text.
	RecText []*string `json:"RecText,omitempty" xml:"RecText,omitempty" type:"Repeated"`
	// The text proportion.
	//
	// example:
	//
	// 7.52%
	TextProp *string `json:"TextProp,omitempty" xml:"TextProp,omitempty"`
	// The usage information.
	//
	// example:
	//
	// {"ProcessedImageCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s ImageRecognitionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageRecognitionResponseBodyData) GetBorderPixel() *string {
	return s.BorderPixel
}

func (s *ImageRecognitionResponseBodyData) GetNoobjCharacter() *bool {
	return s.NoobjCharacter
}

func (s *ImageRecognitionResponseBodyData) GetNoobjLogo() *bool {
	return s.NoobjLogo
}

func (s *ImageRecognitionResponseBodyData) GetNoobjNpx() *bool {
	return s.NoobjNpx
}

func (s *ImageRecognitionResponseBodyData) GetNoobjWatermark() *bool {
	return s.NoobjWatermark
}

func (s *ImageRecognitionResponseBodyData) GetObjCharacter() *bool {
	return s.ObjCharacter
}

func (s *ImageRecognitionResponseBodyData) GetObjLogo() *bool {
	return s.ObjLogo
}

func (s *ImageRecognitionResponseBodyData) GetObjNpx() *bool {
	return s.ObjNpx
}

func (s *ImageRecognitionResponseBodyData) GetObjWatermark() *bool {
	return s.ObjWatermark
}

func (s *ImageRecognitionResponseBodyData) GetPdNum() *int32 {
	return s.PdNum
}

func (s *ImageRecognitionResponseBodyData) GetPdProp() *string {
	return s.PdProp
}

func (s *ImageRecognitionResponseBodyData) GetRecText() []*string {
	return s.RecText
}

func (s *ImageRecognitionResponseBodyData) GetTextProp() *string {
	return s.TextProp
}

func (s *ImageRecognitionResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *ImageRecognitionResponseBodyData) SetBorderPixel(v string) *ImageRecognitionResponseBodyData {
	s.BorderPixel = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetNoobjCharacter(v bool) *ImageRecognitionResponseBodyData {
	s.NoobjCharacter = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetNoobjLogo(v bool) *ImageRecognitionResponseBodyData {
	s.NoobjLogo = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetNoobjNpx(v bool) *ImageRecognitionResponseBodyData {
	s.NoobjNpx = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetNoobjWatermark(v bool) *ImageRecognitionResponseBodyData {
	s.NoobjWatermark = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetObjCharacter(v bool) *ImageRecognitionResponseBodyData {
	s.ObjCharacter = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetObjLogo(v bool) *ImageRecognitionResponseBodyData {
	s.ObjLogo = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetObjNpx(v bool) *ImageRecognitionResponseBodyData {
	s.ObjNpx = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetObjWatermark(v bool) *ImageRecognitionResponseBodyData {
	s.ObjWatermark = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetPdNum(v int32) *ImageRecognitionResponseBodyData {
	s.PdNum = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetPdProp(v string) *ImageRecognitionResponseBodyData {
	s.PdProp = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetRecText(v []*string) *ImageRecognitionResponseBodyData {
	s.RecText = v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetTextProp(v string) *ImageRecognitionResponseBodyData {
	s.TextProp = &v
	return s
}

func (s *ImageRecognitionResponseBodyData) SetUsageMap(v map[string]*int64) *ImageRecognitionResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageRecognitionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
