// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeImageFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *MergeImageFaceResponseBodyData) *MergeImageFaceResponseBody
	GetData() *MergeImageFaceResponseBodyData
	SetRequestId(v string) *MergeImageFaceResponseBody
	GetRequestId() *string
}

type MergeImageFaceResponseBody struct {
	Data *MergeImageFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 8473A5E5-488E-4612-800E-F95B42381570
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MergeImageFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MergeImageFaceResponseBody) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponseBody) GetData() *MergeImageFaceResponseBodyData {
	return s.Data
}

func (s *MergeImageFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MergeImageFaceResponseBody) SetData(v *MergeImageFaceResponseBodyData) *MergeImageFaceResponseBody {
	s.Data = v
	return s
}

func (s *MergeImageFaceResponseBody) SetRequestId(v string) *MergeImageFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeImageFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MergeImageFaceResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/image-face-fusion/09dedc78-c355-442a-9312-7ab074d6de49_54dc_20210129-025407.jpg?Expires=1611890647&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=2mk2OgIrwMqrvA%2FvDdp0MNaoVU****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s MergeImageFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MergeImageFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *MergeImageFaceResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *MergeImageFaceResponseBodyData) SetImageURL(v string) *MergeImageFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *MergeImageFaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
