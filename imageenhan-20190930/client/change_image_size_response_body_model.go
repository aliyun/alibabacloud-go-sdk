// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeImageSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ChangeImageSizeResponseBodyData) *ChangeImageSizeResponseBody
	GetData() *ChangeImageSizeResponseBodyData
	SetRequestId(v string) *ChangeImageSizeResponseBody
	GetRequestId() *string
}

type ChangeImageSizeResponseBody struct {
	Data *ChangeImageSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2833446F-A431-40EB-A502-6EC9DFEEEEB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeImageSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeImageSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBody) GetData() *ChangeImageSizeResponseBodyData {
	return s.Data
}

func (s *ChangeImageSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeImageSizeResponseBody) SetData(v *ChangeImageSizeResponseBodyData) *ChangeImageSizeResponseBody {
	s.Data = v
	return s
}

func (s *ChangeImageSizeResponseBody) SetRequestId(v string) *ChangeImageSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeImageSizeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeImageSizeResponseBodyData struct {
	RetainLocation *ChangeImageSizeResponseBodyDataRetainLocation `json:"RetainLocation,omitempty" xml:"RetainLocation,omitempty" type:"Struct"`
	// example:
	//
	// http://ivpd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/upload/result_filter/2019-11-21/invi_filter_015743271470661000112_NVKmET.png?Expires=1574586347&OSSAccessKeyId=LTAI4FeJ8qKkYn6SrHhQ****&Signature=QqRAiqvyXsVlZ77M8yFc5QKJDE****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ChangeImageSizeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeImageSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBodyData) GetRetainLocation() *ChangeImageSizeResponseBodyDataRetainLocation {
	return s.RetainLocation
}

func (s *ChangeImageSizeResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ChangeImageSizeResponseBodyData) SetRetainLocation(v *ChangeImageSizeResponseBodyDataRetainLocation) *ChangeImageSizeResponseBodyData {
	s.RetainLocation = v
	return s
}

func (s *ChangeImageSizeResponseBodyData) SetUrl(v string) *ChangeImageSizeResponseBodyData {
	s.Url = &v
	return s
}

func (s *ChangeImageSizeResponseBodyData) Validate() error {
	if s.RetainLocation != nil {
		if err := s.RetainLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeImageSizeResponseBodyDataRetainLocation struct {
	// example:
	//
	// 224
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 298
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ChangeImageSizeResponseBodyDataRetainLocation) String() string {
	return dara.Prettify(s)
}

func (s ChangeImageSizeResponseBodyDataRetainLocation) GoString() string {
	return s.String()
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) GetHeight() *int32 {
	return s.Height
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) GetWidth() *int32 {
	return s.Width
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) GetX() *int32 {
	return s.X
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) GetY() *int32 {
	return s.Y
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetHeight(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Height = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetWidth(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Width = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetX(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.X = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) SetY(v int32) *ChangeImageSizeResponseBodyDataRetainLocation {
	s.Y = &v
	return s
}

func (s *ChangeImageSizeResponseBodyDataRetainLocation) Validate() error {
	return dara.Validate(s)
}
