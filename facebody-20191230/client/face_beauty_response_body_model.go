// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceBeautyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *FaceBeautyResponseBodyData) *FaceBeautyResponseBody
	GetData() *FaceBeautyResponseBodyData
	SetRequestId(v string) *FaceBeautyResponseBody
	GetRequestId() *string
}

type FaceBeautyResponseBody struct {
	Data *FaceBeautyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 35C284E1-C5F5-4E5E-B7AD-97BBF485CDC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FaceBeautyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FaceBeautyResponseBody) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponseBody) GetData() *FaceBeautyResponseBodyData {
	return s.Data
}

func (s *FaceBeautyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FaceBeautyResponseBody) SetData(v *FaceBeautyResponseBodyData) *FaceBeautyResponseBody {
	s.Data = v
	return s
}

func (s *FaceBeautyResponseBody) SetRequestId(v string) *FaceBeautyResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceBeautyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FaceBeautyResponseBodyData struct {
	// example:
	//
	// http://algo-app-taobao-mm-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/pixelai-portrait-beauty/2020_03_24/2bbbdb7907901412_facebeauty.jpg?Expires=1585277923&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=myOPfNQcRyijgrWBU%2BZ4dPuV5Q****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s FaceBeautyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FaceBeautyResponseBodyData) GoString() string {
	return s.String()
}

func (s *FaceBeautyResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *FaceBeautyResponseBodyData) SetImageURL(v string) *FaceBeautyResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *FaceBeautyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
