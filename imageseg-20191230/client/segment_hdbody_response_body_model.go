// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDBodyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SegmentHDBodyResponseBodyData) *SegmentHDBodyResponseBody
	GetData() *SegmentHDBodyResponseBodyData
	SetRequestId(v string) *SegmentHDBodyResponseBody
	GetRequestId() *string
}

type SegmentHDBodyResponseBody struct {
	Data *SegmentHDBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// A8D3F5C3-E414-4981-8D84-E2CADF0B7CBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SegmentHDBodyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyResponseBody) GetData() *SegmentHDBodyResponseBodyData {
	return s.Data
}

func (s *SegmentHDBodyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SegmentHDBodyResponseBody) SetData(v *SegmentHDBodyResponseBodyData) *SegmentHDBodyResponseBody {
	s.Data = v
	return s
}

func (s *SegmentHDBodyResponseBody) SetRequestId(v string) *SegmentHDBodyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentHDBodyResponseBody) Validate() error {
	return dara.Validate(s)
}

type SegmentHDBodyResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/segmenthdbody-2020-05-18-16-27-45-675d9884d7-kd9dz/2020-5-18/invi_humansegmenter_015897914589851000001_wQbLq9.png?Expires=1589793259&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=Lx6xSS0t7lqEvy5Qd1keccIAjL****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s SegmentHDBodyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *SegmentHDBodyResponseBodyData) SetImageURL(v string) *SegmentHDBodyResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *SegmentHDBodyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
