// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceImageTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryFaceImageTemplateResponseBodyData) *QueryFaceImageTemplateResponseBody
	GetData() *QueryFaceImageTemplateResponseBodyData
	SetRequestId(v string) *QueryFaceImageTemplateResponseBody
	GetRequestId() *string
}

type QueryFaceImageTemplateResponseBody struct {
	Data *QueryFaceImageTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 7C29675C-751F-4D2F-86FB-2BD8D69AC860
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFaceImageTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBody) GetData() *QueryFaceImageTemplateResponseBodyData {
	return s.Data
}

func (s *QueryFaceImageTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceImageTemplateResponseBody) SetData(v *QueryFaceImageTemplateResponseBodyData) *QueryFaceImageTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceImageTemplateResponseBody) SetRequestId(v string) *QueryFaceImageTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceImageTemplateResponseBodyData struct {
	Elements []*QueryFaceImageTemplateResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Total    *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceImageTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBodyData) GetElements() []*QueryFaceImageTemplateResponseBodyDataElements {
	return s.Elements
}

func (s *QueryFaceImageTemplateResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryFaceImageTemplateResponseBodyData) SetElements(v []*QueryFaceImageTemplateResponseBodyDataElements) *QueryFaceImageTemplateResponseBodyData {
	s.Elements = v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyData) SetTotal(v int64) *QueryFaceImageTemplateResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyData) Validate() error {
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFaceImageTemplateResponseBodyDataElements struct {
	// example:
	//
	// 2021-01-29 10:19:05
	CreateTime *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FaceInfos  []*QueryFaceImageTemplateResponseBodyDataElementsFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/MergeImageFace/MergeImageFace-1.png
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// example:
	//
	// 2021-01-29 10:19:05
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// zhangsan
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceImageTemplateResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) GetFaceInfos() []*QueryFaceImageTemplateResponseBodyDataElementsFaceInfos {
	return s.FaceInfos
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) GetUserId() *string {
	return s.UserId
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetCreateTime(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.CreateTime = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetFaceInfos(v []*QueryFaceImageTemplateResponseBodyDataElementsFaceInfos) *QueryFaceImageTemplateResponseBodyDataElements {
	s.FaceInfos = v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetTemplateId(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.TemplateId = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetTemplateURL(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.TemplateURL = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetUpdateTime(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.UpdateTime = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) SetUserId(v string) *QueryFaceImageTemplateResponseBodyDataElements {
	s.UserId = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElements) Validate() error {
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

type QueryFaceImageTemplateResponseBodyDataElementsFaceInfos struct {
	FaceRect *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect `json:"FaceRect,omitempty" xml:"FaceRect,omitempty" type:"Struct"`
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048
	TemplateFaceID *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
}

func (s QueryFaceImageTemplateResponseBodyDataElementsFaceInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBodyDataElementsFaceInfos) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfos) GetFaceRect() *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect {
	return s.FaceRect
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfos) GetTemplateFaceID() *string {
	return s.TemplateFaceID
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfos) SetFaceRect(v *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) *QueryFaceImageTemplateResponseBodyDataElementsFaceInfos {
	s.FaceRect = v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfos) SetTemplateFaceID(v string) *QueryFaceImageTemplateResponseBodyDataElementsFaceInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfos) Validate() error {
	if s.FaceRect != nil {
		if err := s.FaceRect.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect struct {
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

func (s QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) GoString() string {
	return s.String()
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) GetHeight() *string {
	return s.Height
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) GetWidth() *string {
	return s.Width
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) GetX() *string {
	return s.X
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) GetY() *string {
	return s.Y
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) SetHeight(v string) *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect {
	s.Height = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) SetWidth(v string) *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect {
	s.Width = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) SetX(v string) *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect {
	s.X = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) SetY(v string) *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect {
	s.Y = &v
	return s
}

func (s *QueryFaceImageTemplateResponseBodyDataElementsFaceInfosFaceRect) Validate() error {
	return dara.Validate(s)
}
