// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddImageResponseBody
	GetCode() *int32
	SetMessage(v string) *AddImageResponseBody
	GetMessage() *string
	SetPicInfo(v *AddImageResponseBodyPicInfo) *AddImageResponseBody
	GetPicInfo() *AddImageResponseBodyPicInfo
	SetRequestId(v string) *AddImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddImageResponseBody
	GetSuccess() *bool
}

type AddImageResponseBody struct {
	// The code returned.
	//
	// 	- A value of 0 indicates that the request was successful.
	//
	// 	- Values other than 0 indicate that the request failed.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request failed.
	//
	// > No value is returned if the request was successful, and an error message is returned if the request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The results of category prediction and subject identification.
	PicInfo *AddImageResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E0845DE6-52AF-4B50-9F15-51ED4044E6AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddImageResponseBody) GetPicInfo() *AddImageResponseBodyPicInfo {
	return s.PicInfo
}

func (s *AddImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddImageResponseBody) SetCode(v int32) *AddImageResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageResponseBody) SetMessage(v string) *AddImageResponseBody {
	s.Message = &v
	return s
}

func (s *AddImageResponseBody) SetPicInfo(v *AddImageResponseBodyPicInfo) *AddImageResponseBody {
	s.PicInfo = v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetSuccess(v bool) *AddImageResponseBody {
	s.Success = &v
	return s
}

func (s *AddImageResponseBody) Validate() error {
	if s.PicInfo != nil {
		if err := s.PicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddImageResponseBodyPicInfo struct {
	// The result of category prediction. If a category is specified in the request, the specified category prevails.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The result of subject identification. The subject area of the image is in the format of `x1,x2,y1,y2`. `x1 and y1` represent the position in the upper-left corner, in pixels. `x2 and y2` represent the position in the lower-right corner, in pixels. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 94,691,206,650
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s AddImageResponseBodyPicInfo) String() string {
	return dara.Prettify(s)
}

func (s AddImageResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *AddImageResponseBodyPicInfo) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *AddImageResponseBodyPicInfo) GetRegion() *string {
	return s.Region
}

func (s *AddImageResponseBodyPicInfo) SetCategoryId(v int32) *AddImageResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *AddImageResponseBodyPicInfo) SetRegion(v string) *AddImageResponseBodyPicInfo {
	s.Region = &v
	return s
}

func (s *AddImageResponseBodyPicInfo) Validate() error {
	return dara.Validate(s)
}
