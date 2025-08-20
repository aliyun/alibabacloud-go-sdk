// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ParseFaceResponseBodyData) *ParseFaceResponseBody
	GetData() *ParseFaceResponseBodyData
	SetRequestId(v string) *ParseFaceResponseBody
	GetRequestId() *string
}

type ParseFaceResponseBody struct {
	Data *ParseFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6C24839-91A7-41DA-B31F-98F08EF80CC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ParseFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ParseFaceResponseBody) GoString() string {
	return s.String()
}

func (s *ParseFaceResponseBody) GetData() *ParseFaceResponseBodyData {
	return s.Data
}

func (s *ParseFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ParseFaceResponseBody) SetData(v *ParseFaceResponseBodyData) *ParseFaceResponseBody {
	s.Data = v
	return s
}

func (s *ParseFaceResponseBody) SetRequestId(v string) *ParseFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ParseFaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ParseFaceResponseBodyData struct {
	Elements []*ParseFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageseg/ParseFace/ParseFace1.jpg
	OriginImageURL *string `json:"OriginImageURL,omitempty" xml:"OriginImageURL,omitempty"`
}

func (s ParseFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ParseFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ParseFaceResponseBodyData) GetElements() []*ParseFaceResponseBodyDataElements {
	return s.Elements
}

func (s *ParseFaceResponseBodyData) GetOriginImageURL() *string {
	return s.OriginImageURL
}

func (s *ParseFaceResponseBodyData) SetElements(v []*ParseFaceResponseBodyDataElements) *ParseFaceResponseBodyData {
	s.Elements = v
	return s
}

func (s *ParseFaceResponseBodyData) SetOriginImageURL(v string) *ParseFaceResponseBodyData {
	s.OriginImageURL = &v
	return s
}

func (s *ParseFaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ParseFaceResponseBodyDataElements struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/fivesensesegmenter/prod/560FA2E7-FDC6-59A5-ABDD-D62A05146734/skin/_18dd_20211231-040658.png?Expires=1640925418&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=2g0M88wZl%2Bn4t4gzQX%2BTIskpWB****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// skin
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ParseFaceResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s ParseFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *ParseFaceResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *ParseFaceResponseBodyDataElements) GetName() *string {
	return s.Name
}

func (s *ParseFaceResponseBodyDataElements) SetImageURL(v string) *ParseFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *ParseFaceResponseBodyDataElements) SetName(v string) *ParseFaceResponseBodyDataElements {
	s.Name = &v
	return s
}

func (s *ParseFaceResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
