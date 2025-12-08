// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSuperResolutionImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateSuperResolutionImageResponseBodyData) *GenerateSuperResolutionImageResponseBody
	GetData() *GenerateSuperResolutionImageResponseBodyData
	SetMessage(v string) *GenerateSuperResolutionImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateSuperResolutionImageResponseBody
	GetRequestId() *string
}

type GenerateSuperResolutionImageResponseBody struct {
	Data    *GenerateSuperResolutionImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4ad5c3ef-5ac4-4e1c-b14f-90d939aa73eb
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateSuperResolutionImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateSuperResolutionImageResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateSuperResolutionImageResponseBody) GetData() *GenerateSuperResolutionImageResponseBodyData {
	return s.Data
}

func (s *GenerateSuperResolutionImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateSuperResolutionImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateSuperResolutionImageResponseBody) SetData(v *GenerateSuperResolutionImageResponseBodyData) *GenerateSuperResolutionImageResponseBody {
	s.Data = v
	return s
}

func (s *GenerateSuperResolutionImageResponseBody) SetMessage(v string) *GenerateSuperResolutionImageResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateSuperResolutionImageResponseBody) SetRequestId(v string) *GenerateSuperResolutionImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateSuperResolutionImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateSuperResolutionImageResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-damo-eas-cn-shanghai.oss-cn-shanghai.aliyuncs.com/diffusion-sr/2023-02-07/d01cede5-28bf-4719-96d9-77198d51c2f2/20230207_150650515242_3dbctnjy5f.jpg?Expires=1675755681&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=4FeDXpp0DilXsHdt7qc%2Ffh3zoC****
	ResultUrl *string `json:"ResultUrl,omitempty" xml:"ResultUrl,omitempty"`
}

func (s GenerateSuperResolutionImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateSuperResolutionImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateSuperResolutionImageResponseBodyData) GetResultUrl() *string {
	return s.ResultUrl
}

func (s *GenerateSuperResolutionImageResponseBodyData) SetResultUrl(v string) *GenerateSuperResolutionImageResponseBodyData {
	s.ResultUrl = &v
	return s
}

func (s *GenerateSuperResolutionImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
