// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceImageTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddFaceImageTemplateResponseBodyData) *AddFaceImageTemplateResponseBody
	GetData() *AddFaceImageTemplateResponseBodyData
	SetRequestId(v string) *AddFaceImageTemplateResponseBody
	GetRequestId() *string
}

type AddFaceImageTemplateResponseBody struct {
	Data *AddFaceImageTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 798A721D-7C7F-4D87-A125-1D04B3055C2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceImageTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponseBody) GetData() *AddFaceImageTemplateResponseBodyData {
	return s.Data
}

func (s *AddFaceImageTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFaceImageTemplateResponseBody) SetData(v *AddFaceImageTemplateResponseBodyData) *AddFaceImageTemplateResponseBody {
	s.Data = v
	return s
}

func (s *AddFaceImageTemplateResponseBody) SetRequestId(v string) *AddFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceImageTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFaceImageTemplateResponseBodyData struct {
	FaceInfos []*AddFaceImageTemplateResponseBodyDataFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddFaceImageTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddFaceImageTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponseBodyData) GetFaceInfos() []*AddFaceImageTemplateResponseBodyDataFaceInfos {
	return s.FaceInfos
}

func (s *AddFaceImageTemplateResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddFaceImageTemplateResponseBodyData) SetFaceInfos(v []*AddFaceImageTemplateResponseBodyDataFaceInfos) *AddFaceImageTemplateResponseBodyData {
	s.FaceInfos = v
	return s
}

func (s *AddFaceImageTemplateResponseBodyData) SetTemplateId(v string) *AddFaceImageTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *AddFaceImageTemplateResponseBodyData) Validate() error {
	if s.FaceInfos != nil {
		for _, item := range s.FaceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddFaceImageTemplateResponseBodyDataFaceInfos struct {
	FaceRect *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect `json:"FaceRect,omitempty" xml:"FaceRect,omitempty" type:"Struct"`
	// example:
	//
	// string 6cd509ea-54fa-4730-8e9d-c94cadcda048_0
	TemplateFaceID *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
}

func (s AddFaceImageTemplateResponseBodyDataFaceInfos) String() string {
	return dara.Prettify(s)
}

func (s AddFaceImageTemplateResponseBodyDataFaceInfos) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfos) GetFaceRect() *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect {
	return s.FaceRect
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfos) GetTemplateFaceID() *string {
	return s.TemplateFaceID
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfos) SetFaceRect(v *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) *AddFaceImageTemplateResponseBodyDataFaceInfos {
	s.FaceRect = v
	return s
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfos) SetTemplateFaceID(v string) *AddFaceImageTemplateResponseBodyDataFaceInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfos) Validate() error {
	if s.FaceRect != nil {
		if err := s.FaceRect.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect struct {
	// example:
	//
	// 94
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 89
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 254
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 318
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) String() string {
	return dara.Prettify(s)
}

func (s AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) GoString() string {
	return s.String()
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) GetHeight() *string {
	return s.Height
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) GetWidth() *string {
	return s.Width
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) GetX() *string {
	return s.X
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) GetY() *string {
	return s.Y
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) SetHeight(v string) *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect {
	s.Height = &v
	return s
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) SetWidth(v string) *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect {
	s.Width = &v
	return s
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) SetX(v string) *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect {
	s.X = &v
	return s
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) SetY(v string) *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect {
	s.Y = &v
	return s
}

func (s *AddFaceImageTemplateResponseBodyDataFaceInfosFaceRect) Validate() error {
	return dara.Validate(s)
}
