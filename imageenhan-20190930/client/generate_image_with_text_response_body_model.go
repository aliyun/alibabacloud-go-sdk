// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageWithTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateImageWithTextResponseBodyData) *GenerateImageWithTextResponseBody
	GetData() *GenerateImageWithTextResponseBodyData
	SetMessage(v string) *GenerateImageWithTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateImageWithTextResponseBody
	GetRequestId() *string
}

type GenerateImageWithTextResponseBody struct {
	Data    *GenerateImageWithTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7574ee8f-38a3-4b1e-9280-11c33ab46e51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateImageWithTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextResponseBody) GetData() *GenerateImageWithTextResponseBodyData {
	return s.Data
}

func (s *GenerateImageWithTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateImageWithTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateImageWithTextResponseBody) SetData(v *GenerateImageWithTextResponseBodyData) *GenerateImageWithTextResponseBody {
	s.Data = v
	return s
}

func (s *GenerateImageWithTextResponseBody) SetMessage(v string) *GenerateImageWithTextResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateImageWithTextResponseBody) SetRequestId(v string) *GenerateImageWithTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateImageWithTextResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateImageWithTextResponseBodyData struct {
	ImageUrls []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
}

func (s GenerateImageWithTextResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextResponseBodyData) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *GenerateImageWithTextResponseBodyData) SetImageUrls(v []*string) *GenerateImageWithTextResponseBodyData {
	s.ImageUrls = v
	return s
}

func (s *GenerateImageWithTextResponseBodyData) Validate() error {
	return dara.Validate(s)
}
