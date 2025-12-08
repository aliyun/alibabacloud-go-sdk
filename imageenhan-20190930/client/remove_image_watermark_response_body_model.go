// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RemoveImageWatermarkResponseBodyData) *RemoveImageWatermarkResponseBody
	GetData() *RemoveImageWatermarkResponseBodyData
	SetRequestId(v string) *RemoveImageWatermarkResponseBody
	GetRequestId() *string
}

type RemoveImageWatermarkResponseBody struct {
	Data *RemoveImageWatermarkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 885070A7-E721-4062-99A0-80C0EBBF9406
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveImageWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkResponseBody) GetData() *RemoveImageWatermarkResponseBodyData {
	return s.Data
}

func (s *RemoveImageWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveImageWatermarkResponseBody) SetData(v *RemoveImageWatermarkResponseBodyData) *RemoveImageWatermarkResponseBody {
	s.Data = v
	return s
}

func (s *RemoveImageWatermarkResponseBody) SetRequestId(v string) *RemoveImageWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveImageWatermarkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveImageWatermarkResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/image-delogo/2020-03-27-03/00%3A06-5a6f0a0f-c940-4955-af75-79e8267f1699.jpg?Expires=1585279806&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=R4OC2R5%2Fkw08XYFXmCWjgk7Y9N****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageWatermarkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageWatermarkResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveImageWatermarkResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *RemoveImageWatermarkResponseBodyData) SetImageURL(v string) *RemoveImageWatermarkResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *RemoveImageWatermarkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
