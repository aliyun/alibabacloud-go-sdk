// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageWithTextAndImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateImageWithTextAndImageResponseBodyData) *GenerateImageWithTextAndImageResponseBody
	GetData() *GenerateImageWithTextAndImageResponseBodyData
	SetMessage(v string) *GenerateImageWithTextAndImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateImageWithTextAndImageResponseBody
	GetRequestId() *string
}

type GenerateImageWithTextAndImageResponseBody struct {
	Data    *GenerateImageWithTextAndImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 141fb6d1-28e8-4d93-8165-d966f7092e6a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateImageWithTextAndImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextAndImageResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextAndImageResponseBody) GetData() *GenerateImageWithTextAndImageResponseBodyData {
	return s.Data
}

func (s *GenerateImageWithTextAndImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateImageWithTextAndImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateImageWithTextAndImageResponseBody) SetData(v *GenerateImageWithTextAndImageResponseBodyData) *GenerateImageWithTextAndImageResponseBody {
	s.Data = v
	return s
}

func (s *GenerateImageWithTextAndImageResponseBody) SetMessage(v string) *GenerateImageWithTextAndImageResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateImageWithTextAndImageResponseBody) SetRequestId(v string) *GenerateImageWithTextAndImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateImageWithTextAndImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateImageWithTextAndImageResponseBodyData struct {
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
}

func (s GenerateImageWithTextAndImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextAndImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextAndImageResponseBodyData) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *GenerateImageWithTextAndImageResponseBodyData) SetImageUrls(v []*string) *GenerateImageWithTextAndImageResponseBodyData {
	s.ImageUrls = v
	return s
}

func (s *GenerateImageWithTextAndImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
