// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentSkyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentSkyResponseBodyData) *SegmentSkyResponseBody
	GetData() *SegmentSkyResponseBodyData
	SetRequestId(v string) *SegmentSkyResponseBody
	GetRequestId() *string
}

type SegmentSkyResponseBody struct {
	Data *SegmentSkyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 80E9D0A0-0330-4210-9986-CAC50C922FF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentSkyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentSkyResponseBody) GetData() *SegmentSkyResponseBodyData {
	return s.Data
}

func (s *SegmentSkyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentSkyResponseBody) SetData(v *SegmentSkyResponseBodyData) *SegmentSkyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentSkyResponseBody) SetRequestId(v string) *SegmentSkyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentSkyResponseBody) Validate() error {
	return dara.Validate(s)
}

type SegmentSkyResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/skysegmentation-2020-05-18-10-44-16-5bc8dc79f9-92b7z/2020-5-18/invi_skySegmentator_015897703560961000003_SqZLDv.png?Expires=1589772156&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=gXrzAUl%2BvIdYbQ9XKdho54MlkX****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentSkyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentSkyResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentSkyResponseBodyData) SetImageURL(v string) *SegmentSkyResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *SegmentSkyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
