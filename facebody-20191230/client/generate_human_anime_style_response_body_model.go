// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanAnimeStyleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateHumanAnimeStyleResponseBodyData) *GenerateHumanAnimeStyleResponseBody
	GetData() *GenerateHumanAnimeStyleResponseBodyData
	SetRequestId(v string) *GenerateHumanAnimeStyleResponseBody
	GetRequestId() *string
}

type GenerateHumanAnimeStyleResponseBody struct {
	Data *GenerateHumanAnimeStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 59697D68-2A6E-4553-89BD-0FADD07881E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateHumanAnimeStyleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponseBody) GetData() *GenerateHumanAnimeStyleResponseBodyData {
	return s.Data
}

func (s *GenerateHumanAnimeStyleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateHumanAnimeStyleResponseBody) SetData(v *GenerateHumanAnimeStyleResponseBodyData) *GenerateHumanAnimeStyleResponseBody {
	s.Data = v
	return s
}

func (s *GenerateHumanAnimeStyleResponseBody) SetRequestId(v string) *GenerateHumanAnimeStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateHumanAnimeStyleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateHumanAnimeStyleResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/person-image-cartoonizer/59697D68-2A6E-4553-89BD-0FADD07881E8_7ee5_20201027-070958.jpg?Expires=1603784400&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSR****&Signature=ut2kn46Lz%2FRwqJ9jWJ0RBDut12****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanAnimeStyleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *GenerateHumanAnimeStyleResponseBodyData) SetImageURL(v string) *GenerateHumanAnimeStyleResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *GenerateHumanAnimeStyleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
