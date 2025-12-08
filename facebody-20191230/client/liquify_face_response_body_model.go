// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiquifyFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *LiquifyFaceResponseBodyData) *LiquifyFaceResponseBody
	GetData() *LiquifyFaceResponseBodyData
	SetRequestId(v string) *LiquifyFaceResponseBody
	GetRequestId() *string
}

type LiquifyFaceResponseBody struct {
	Data *LiquifyFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CA1C8937-B30C-15E6-B804-41C357BA413B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LiquifyFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LiquifyFaceResponseBody) GoString() string {
	return s.String()
}

func (s *LiquifyFaceResponseBody) GetData() *LiquifyFaceResponseBodyData {
	return s.Data
}

func (s *LiquifyFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LiquifyFaceResponseBody) SetData(v *LiquifyFaceResponseBodyData) *LiquifyFaceResponseBody {
	s.Data = v
	return s
}

func (s *LiquifyFaceResponseBody) SetRequestId(v string) *LiquifyFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LiquifyFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LiquifyFaceResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/face-liquify/CA1C8937-B30C-15E6-B804-41C357BA413B_5eca_20210923-093411.jpg?Expires=1632391451&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=46apJQVNBz%2FFzDLEYn2M1w9MKA****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s LiquifyFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s LiquifyFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *LiquifyFaceResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *LiquifyFaceResponseBodyData) SetImageURL(v string) *LiquifyFaceResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *LiquifyFaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
