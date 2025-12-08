// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlurFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BlurFaceResponseBodyData) *BlurFaceResponseBody
	GetData() *BlurFaceResponseBodyData
	SetRequestId(v string) *BlurFaceResponseBody
	GetRequestId() *string
}

type BlurFaceResponseBody struct {
	Data *BlurFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4C6080B0-64C4-488E-BBA6-245749CA11AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BlurFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BlurFaceResponseBody) GoString() string {
	return s.String()
}

func (s *BlurFaceResponseBody) GetData() *BlurFaceResponseBodyData {
	return s.Data
}

func (s *BlurFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BlurFaceResponseBody) SetData(v *BlurFaceResponseBodyData) *BlurFaceResponseBody {
	s.Data = v
	return s
}

func (s *BlurFaceResponseBody) SetRequestId(v string) *BlurFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlurFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BlurFaceResponseBodyData struct {
	// example:
	//
	// http://viapi-cn-shanghai-dha-segmenter.oss-cn-shanghai.aliyuncs.com/upload/result_FaceBlur/2020-8-5/invi_FaceBlur_015966077063461060948_IBdDsq.jpg?Expires=1596609506&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=x8n3jq1X91Sq7BKWE4vRHSu6g9****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	MaskURL  *string `json:"MaskURL,omitempty" xml:"MaskURL,omitempty"`
}

func (s BlurFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BlurFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *BlurFaceResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *BlurFaceResponseBodyData) GetMaskURL() *string {
	return s.MaskURL
}

func (s *BlurFaceResponseBodyData) SetImageURL(v string) *BlurFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *BlurFaceResponseBodyData) SetMaskURL(v string) *BlurFaceResponseBodyData {
	s.MaskURL = &v
	return s
}

func (s *BlurFaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
