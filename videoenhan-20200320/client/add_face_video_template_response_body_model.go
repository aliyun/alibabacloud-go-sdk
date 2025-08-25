// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceVideoTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v *AddFaceVideoTemplateResponseBodyDate) *AddFaceVideoTemplateResponseBody
	GetDate() *AddFaceVideoTemplateResponseBodyDate
	SetMessage(v string) *AddFaceVideoTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddFaceVideoTemplateResponseBody
	GetRequestId() *string
}

type AddFaceVideoTemplateResponseBody struct {
	Date    *AddFaceVideoTemplateResponseBodyDate `json:"Date,omitempty" xml:"Date,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F61182AE-515B-5B0A-A877-1C9AE908FF15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFaceVideoTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFaceVideoTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateResponseBody) GetDate() *AddFaceVideoTemplateResponseBodyDate {
	return s.Date
}

func (s *AddFaceVideoTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddFaceVideoTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFaceVideoTemplateResponseBody) SetDate(v *AddFaceVideoTemplateResponseBodyDate) *AddFaceVideoTemplateResponseBody {
	s.Date = v
	return s
}

func (s *AddFaceVideoTemplateResponseBody) SetMessage(v string) *AddFaceVideoTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBody) SetRequestId(v string) *AddFaceVideoTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddFaceVideoTemplateResponseBodyDate struct {
	FaceInfos []*AddFaceVideoTemplateResponseBodyDateFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TransResult *string `json:"TransResult,omitempty" xml:"TransResult,omitempty"`
}

func (s AddFaceVideoTemplateResponseBodyDate) String() string {
	return dara.Prettify(s)
}

func (s AddFaceVideoTemplateResponseBodyDate) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateResponseBodyDate) GetFaceInfos() []*AddFaceVideoTemplateResponseBodyDateFaceInfos {
	return s.FaceInfos
}

func (s *AddFaceVideoTemplateResponseBodyDate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AddFaceVideoTemplateResponseBodyDate) GetTransResult() *string {
	return s.TransResult
}

func (s *AddFaceVideoTemplateResponseBodyDate) SetFaceInfos(v []*AddFaceVideoTemplateResponseBodyDateFaceInfos) *AddFaceVideoTemplateResponseBodyDate {
	s.FaceInfos = v
	return s
}

func (s *AddFaceVideoTemplateResponseBodyDate) SetTemplateId(v string) *AddFaceVideoTemplateResponseBodyDate {
	s.TemplateId = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBodyDate) SetTransResult(v string) *AddFaceVideoTemplateResponseBodyDate {
	s.TransResult = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBodyDate) Validate() error {
	return dara.Validate(s)
}

type AddFaceVideoTemplateResponseBodyDateFaceInfos struct {
	TemplateFaceID  *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
	TemplateFaceURL *string `json:"TemplateFaceURL,omitempty" xml:"TemplateFaceURL,omitempty"`
}

func (s AddFaceVideoTemplateResponseBodyDateFaceInfos) String() string {
	return dara.Prettify(s)
}

func (s AddFaceVideoTemplateResponseBodyDateFaceInfos) GoString() string {
	return s.String()
}

func (s *AddFaceVideoTemplateResponseBodyDateFaceInfos) GetTemplateFaceID() *string {
	return s.TemplateFaceID
}

func (s *AddFaceVideoTemplateResponseBodyDateFaceInfos) GetTemplateFaceURL() *string {
	return s.TemplateFaceURL
}

func (s *AddFaceVideoTemplateResponseBodyDateFaceInfos) SetTemplateFaceID(v string) *AddFaceVideoTemplateResponseBodyDateFaceInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBodyDateFaceInfos) SetTemplateFaceURL(v string) *AddFaceVideoTemplateResponseBodyDateFaceInfos {
	s.TemplateFaceURL = &v
	return s
}

func (s *AddFaceVideoTemplateResponseBodyDateFaceInfos) Validate() error {
	return dara.Validate(s)
}
