// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDynamicImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateDynamicImageResponseBodyData) *GenerateDynamicImageResponseBody
	GetData() *GenerateDynamicImageResponseBodyData
	SetRequestId(v string) *GenerateDynamicImageResponseBody
	GetRequestId() *string
}

type GenerateDynamicImageResponseBody struct {
	Data *GenerateDynamicImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 3C047FB7-AE01-42CF-948F-D57F224620ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDynamicImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDynamicImageResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageResponseBody) GetData() *GenerateDynamicImageResponseBodyData {
	return s.Data
}

func (s *GenerateDynamicImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDynamicImageResponseBody) SetData(v *GenerateDynamicImageResponseBodyData) *GenerateDynamicImageResponseBody {
	s.Data = v
	return s
}

func (s *GenerateDynamicImageResponseBody) SetRequestId(v string) *GenerateDynamicImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDynamicImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateDynamicImageResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/photo-style-imitation/dynamicPhotoResult/dySkyHair_0d0465eb-b1c9-472a-a7dc-cdff1333f794.avi?Expires=1601196370&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=QejSaTZOR2MB2lVHOUH%2Fj8l3P4****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateDynamicImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateDynamicImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GenerateDynamicImageResponseBodyData) SetUrl(v string) *GenerateDynamicImageResponseBodyData {
	s.Url = &v
	return s
}

func (s *GenerateDynamicImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
