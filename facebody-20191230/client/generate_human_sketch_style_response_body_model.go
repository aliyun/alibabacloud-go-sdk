// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanSketchStyleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateHumanSketchStyleResponseBodyData) *GenerateHumanSketchStyleResponseBody
	GetData() *GenerateHumanSketchStyleResponseBodyData
	SetRequestId(v string) *GenerateHumanSketchStyleResponseBody
	GetRequestId() *string
}

type GenerateHumanSketchStyleResponseBody struct {
	Data *GenerateHumanSketchStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E499788C-22DA-4F0E-B9C0-0E9D30A25716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateHumanSketchStyleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanSketchStyleResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponseBody) GetData() *GenerateHumanSketchStyleResponseBodyData {
	return s.Data
}

func (s *GenerateHumanSketchStyleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateHumanSketchStyleResponseBody) SetData(v *GenerateHumanSketchStyleResponseBodyData) *GenerateHumanSketchStyleResponseBody {
	s.Data = v
	return s
}

func (s *GenerateHumanSketchStyleResponseBody) SetRequestId(v string) *GenerateHumanSketchStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateHumanSketchStyleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateHumanSketchStyleResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/person-image-pencil/fd49dd0a-9e24-4bb5-b303-0745cdb453e0_7aa6_20210128-073045.jpg?Expires=1611820849&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=1oLYG%2FPe%2BNRIK7XcsUQYaKF%2B1C****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanSketchStyleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanSketchStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *GenerateHumanSketchStyleResponseBodyData) SetImageURL(v string) *GenerateHumanSketchStyleResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *GenerateHumanSketchStyleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
