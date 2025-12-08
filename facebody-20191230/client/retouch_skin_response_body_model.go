// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetouchSkinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RetouchSkinResponseBodyData) *RetouchSkinResponseBody
	GetData() *RetouchSkinResponseBodyData
	SetRequestId(v string) *RetouchSkinResponseBody
	GetRequestId() *string
}

type RetouchSkinResponseBody struct {
	Data *RetouchSkinResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FD2BF3DF-3D98-1D5D-85BB-C8EC3A9FE347
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RetouchSkinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetouchSkinResponseBody) GoString() string {
	return s.String()
}

func (s *RetouchSkinResponseBody) GetData() *RetouchSkinResponseBodyData {
	return s.Data
}

func (s *RetouchSkinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetouchSkinResponseBody) SetData(v *RetouchSkinResponseBodyData) *RetouchSkinResponseBody {
	s.Data = v
	return s
}

func (s *RetouchSkinResponseBody) SetRequestId(v string) *RetouchSkinResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetouchSkinResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetouchSkinResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-gd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/skin-retouch/FD2BF3DF-3D98-1D5D-85BB-C8EC3A9FE347_3467_20210923-094238.jpg?Expires=1632391959&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=jXJcFTv3no7Gx%2BLuPvANhRSnc2****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RetouchSkinResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetouchSkinResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetouchSkinResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *RetouchSkinResponseBodyData) SetImageURL(v string) *RetouchSkinResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *RetouchSkinResponseBodyData) Validate() error {
	return dara.Validate(s)
}
