// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoCoverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateVideoCoverResponseBodyData) *GenerateVideoCoverResponseBody
	GetData() *GenerateVideoCoverResponseBodyData
	SetMessage(v string) *GenerateVideoCoverResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateVideoCoverResponseBody
	GetRequestId() *string
}

type GenerateVideoCoverResponseBody struct {
	Data    *GenerateVideoCoverResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5B95B724-C5B9-4F77-A743-0CA4EA95CC82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateVideoCoverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoCoverResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBody) GetData() *GenerateVideoCoverResponseBodyData {
	return s.Data
}

func (s *GenerateVideoCoverResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateVideoCoverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateVideoCoverResponseBody) SetData(v *GenerateVideoCoverResponseBodyData) *GenerateVideoCoverResponseBody {
	s.Data = v
	return s
}

func (s *GenerateVideoCoverResponseBody) SetMessage(v string) *GenerateVideoCoverResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateVideoCoverResponseBody) SetRequestId(v string) *GenerateVideoCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateVideoCoverResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateVideoCoverResponseBodyData struct {
	Outputs []*GenerateVideoCoverResponseBodyDataOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s GenerateVideoCoverResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoCoverResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBodyData) GetOutputs() []*GenerateVideoCoverResponseBodyDataOutputs {
	return s.Outputs
}

func (s *GenerateVideoCoverResponseBodyData) SetOutputs(v []*GenerateVideoCoverResponseBodyDataOutputs) *GenerateVideoCoverResponseBodyData {
	s.Outputs = v
	return s
}

func (s *GenerateVideoCoverResponseBodyData) Validate() error {
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateVideoCoverResponseBodyDataOutputs struct {
	// example:
	//
	// 6.1819260887924425
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-cover/2020-05-11-07/pic_lOyxGGAqQYSANGxP.mp4_202_544_960_c9f88b2a5f75e17c093d1a65f5edff4d_beautified.png?Expires=1589185385&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=PAalKsfeZC4UQzYDTU%2F3D1G7Xt****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateVideoCoverResponseBodyDataOutputs) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoCoverResponseBodyDataOutputs) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) GetConfidence() *float32 {
	return s.Confidence
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) GetImageURL() *string {
	return s.ImageURL
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) SetConfidence(v float32) *GenerateVideoCoverResponseBodyDataOutputs {
	s.Confidence = &v
	return s
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) SetImageURL(v string) *GenerateVideoCoverResponseBodyDataOutputs {
	s.ImageURL = &v
	return s
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) Validate() error {
	return dara.Validate(s)
}
