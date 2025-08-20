// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentSkinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentSkinResponseBodyData) *SegmentSkinResponseBody
	GetData() *SegmentSkinResponseBodyData
	SetRequestId(v string) *SegmentSkinResponseBody
	GetRequestId() *string
}

type SegmentSkinResponseBody struct {
	Data *SegmentSkinResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DA007354-6CF5-45BE-8333-E06318D848C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentSkinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkinResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentSkinResponseBody) GetData() *SegmentSkinResponseBodyData {
	return s.Data
}

func (s *SegmentSkinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentSkinResponseBody) SetData(v *SegmentSkinResponseBodyData) *SegmentSkinResponseBody {
	s.Data = v
	return s
}

func (s *SegmentSkinResponseBody) SetRequestId(v string) *SegmentSkinResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentSkinResponseBody) Validate() error {
	return dara.Validate(s)
}

type SegmentSkinResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_skinsegmenter/2020-9-27/invi_skinsegmenter_016011971641871000001_wQbLq9.jpg?Expires=1601198964&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=xjKc%2BScprmB86cxtI%2B1T0R6TlE****
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SegmentSkinResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkinResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentSkinResponseBodyData) GetURL() *string {
	return s.URL
}

func (s *SegmentSkinResponseBodyData) SetURL(v string) *SegmentSkinResponseBodyData {
	s.URL = &v
	return s
}

func (s *SegmentSkinResponseBodyData) Validate() error {
	return dara.Validate(s)
}
