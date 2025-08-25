// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbstractEcommerceVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AbstractEcommerceVideoResponseBodyData) *AbstractEcommerceVideoResponseBody
	GetData() *AbstractEcommerceVideoResponseBodyData
	SetMessage(v string) *AbstractEcommerceVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *AbstractEcommerceVideoResponseBody
	GetRequestId() *string
}

type AbstractEcommerceVideoResponseBody struct {
	Data    *AbstractEcommerceVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 50B33B81-CCB8-42BC-8A73-AC838618936E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbstractEcommerceVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbstractEcommerceVideoResponseBody) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponseBody) GetData() *AbstractEcommerceVideoResponseBodyData {
	return s.Data
}

func (s *AbstractEcommerceVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbstractEcommerceVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbstractEcommerceVideoResponseBody) SetData(v *AbstractEcommerceVideoResponseBodyData) *AbstractEcommerceVideoResponseBody {
	s.Data = v
	return s
}

func (s *AbstractEcommerceVideoResponseBody) SetMessage(v string) *AbstractEcommerceVideoResponseBody {
	s.Message = &v
	return s
}

func (s *AbstractEcommerceVideoResponseBody) SetRequestId(v string) *AbstractEcommerceVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbstractEcommerceVideoResponseBody) Validate() error {
	return dara.Validate(s)
}

type AbstractEcommerceVideoResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/shop-video-abs/2020-03-20-19/YVgDynxB.jpg?Expires=1584707249&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=e5Q3O%2ByA6H7UhYJeMZxz4p70de****
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/shop-video-abs/2020-03-20-19/YVgDynxB.mp4?Expires=1584707249&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=KErufmbHvTUYYLRj6i42wY7Tew****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AbstractEcommerceVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AbstractEcommerceVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoResponseBodyData) GetVideoCoverUrl() *string {
	return s.VideoCoverUrl
}

func (s *AbstractEcommerceVideoResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AbstractEcommerceVideoResponseBodyData) SetVideoCoverUrl(v string) *AbstractEcommerceVideoResponseBodyData {
	s.VideoCoverUrl = &v
	return s
}

func (s *AbstractEcommerceVideoResponseBodyData) SetVideoUrl(v string) *AbstractEcommerceVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *AbstractEcommerceVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
