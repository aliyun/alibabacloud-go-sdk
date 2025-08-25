// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceVideoTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryFaceVideoTemplateResponseBodyData) *QueryFaceVideoTemplateResponseBody
	GetData() *QueryFaceVideoTemplateResponseBodyData
	SetRequestId(v string) *QueryFaceVideoTemplateResponseBody
	GetRequestId() *string
}

type QueryFaceVideoTemplateResponseBody struct {
	Data *QueryFaceVideoTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A06E3D21-890D-500B-97DA-D8B99D2DDDC4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryFaceVideoTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceVideoTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponseBody) GetData() *QueryFaceVideoTemplateResponseBodyData {
	return s.Data
}

func (s *QueryFaceVideoTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceVideoTemplateResponseBody) SetData(v *QueryFaceVideoTemplateResponseBodyData) *QueryFaceVideoTemplateResponseBody {
	s.Data = v
	return s
}

func (s *QueryFaceVideoTemplateResponseBody) SetRequestId(v string) *QueryFaceVideoTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryFaceVideoTemplateResponseBodyData struct {
	Elements []*QueryFaceVideoTemplateResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Total    *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryFaceVideoTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceVideoTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponseBodyData) GetElements() []*QueryFaceVideoTemplateResponseBodyDataElements {
	return s.Elements
}

func (s *QueryFaceVideoTemplateResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryFaceVideoTemplateResponseBodyData) SetElements(v []*QueryFaceVideoTemplateResponseBodyDataElements) *QueryFaceVideoTemplateResponseBodyData {
	s.Elements = v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyData) SetTotal(v int64) *QueryFaceVideoTemplateResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryFaceVideoTemplateResponseBodyDataElements struct {
	// example:
	//
	// 2021-09-06 15:17:19
	CreateTime *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FaceInfos  []*QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/image-face-fusion/OriginUrl/user/2021-09-06/e4d6ecf6-8dc8-4dac-acb5-56a737ccbc06?Expires=1630914551&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=kBcLm66y7%2FZ3eIMgwXJg1zNP7k****
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// example:
	//
	// 2021-09-06 15:18:15
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// zhangsan
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFaceVideoTemplateResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceVideoTemplateResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) GetFaceInfos() []*QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos {
	return s.FaceInfos
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) GetUserId() *string {
	return s.UserId
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetCreateTime(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.CreateTime = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetFaceInfos(v []*QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.FaceInfos = v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetTemplateId(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.TemplateId = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetTemplateURL(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.TemplateURL = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetUpdateTime(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.UpdateTime = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) SetUserId(v string) *QueryFaceVideoTemplateResponseBodyDataElements {
	s.UserId = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}

type QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos struct {
	TemplateFaceID  *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
	TemplateFaceURL *string `json:"TemplateFaceURL,omitempty" xml:"TemplateFaceURL,omitempty"`
}

func (s QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) GoString() string {
	return s.String()
}

func (s *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) GetTemplateFaceID() *string {
	return s.TemplateFaceID
}

func (s *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) GetTemplateFaceURL() *string {
	return s.TemplateFaceURL
}

func (s *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) SetTemplateFaceID(v string) *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) SetTemplateFaceURL(v string) *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos {
	s.TemplateFaceURL = &v
	return s
}

func (s *QueryFaceVideoTemplateResponseBodyDataElementsFaceInfos) Validate() error {
	return dara.Validate(s)
}
