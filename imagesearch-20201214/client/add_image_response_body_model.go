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
	// The error code.
	//
	// - 0: success.
	//
	// - Non-zero: failure.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// > No data is returned for successful requests. Error messages are returned for failed requests.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The category prediction and subject identification results.
	PicInfo *AddImageResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E0845DE6-52AF-4B50-9F15-51ED4044E6AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
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
	// The information about all categories supported by the system.
	AllCategories []*AddImageResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	// The category prediction result. If the user specifies a category in the request, the specified category is used.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The collection of subject identification results.
	//
	// > Upgrade to V3.1.1 or later to use this feature.
	MultiRegion []*AddImageResponseBodyPicInfoMultiRegion `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	// The subject identification result. The subject region of the image, in the format of `x1,x2,y1,y2`, where `x1,y1` is the upper-left point and `x2,y2` is the lower-right point. If the user specifies a subject region in the request, the specified region is used.
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

func (s *AddImageResponseBodyPicInfo) GetAllCategories() []*AddImageResponseBodyPicInfoAllCategories {
	return s.AllCategories
}

func (s *AddImageResponseBodyPicInfo) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *AddImageResponseBodyPicInfo) GetMultiRegion() []*AddImageResponseBodyPicInfoMultiRegion {
	return s.MultiRegion
}

func (s *AddImageResponseBodyPicInfo) GetRegion() *string {
	return s.Region
}

func (s *AddImageResponseBodyPicInfo) SetAllCategories(v []*AddImageResponseBodyPicInfoAllCategories) *AddImageResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *AddImageResponseBodyPicInfo) SetCategoryId(v int32) *AddImageResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *AddImageResponseBodyPicInfo) SetMultiRegion(v []*AddImageResponseBodyPicInfoMultiRegion) *AddImageResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

func (s *AddImageResponseBodyPicInfo) SetRegion(v string) *AddImageResponseBodyPicInfo {
	s.Region = &v
	return s
}

func (s *AddImageResponseBodyPicInfo) Validate() error {
	if s.AllCategories != nil {
		for _, item := range s.AllCategories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MultiRegion != nil {
		for _, item := range s.MultiRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddImageResponseBodyPicInfoAllCategories struct {
	// The category ID.
	//
	// example:
	//
	// 88888888
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The category name.
	//
	// example:
	//
	// Other
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AddImageResponseBodyPicInfoAllCategories) String() string {
	return dara.Prettify(s)
}

func (s AddImageResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *AddImageResponseBodyPicInfoAllCategories) GetId() *int32 {
	return s.Id
}

func (s *AddImageResponseBodyPicInfoAllCategories) GetName() *string {
	return s.Name
}

func (s *AddImageResponseBodyPicInfoAllCategories) SetId(v int32) *AddImageResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *AddImageResponseBodyPicInfoAllCategories) SetName(v string) *AddImageResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

func (s *AddImageResponseBodyPicInfoAllCategories) Validate() error {
	return dara.Validate(s)
}

type AddImageResponseBodyPicInfoMultiRegion struct {
	// The subject identification result. The subject region of the image, in the format of x1,x2,y1,y2, where x1,y1 is the upper-left point and x2,y2 is the lower-right point. If the user specifies a subject region in the request, the specified region is used.
	//
	// example:
	//
	// 94,691,206,650
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s AddImageResponseBodyPicInfoMultiRegion) String() string {
	return dara.Prettify(s)
}

func (s AddImageResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *AddImageResponseBodyPicInfoMultiRegion) GetRegion() *string {
	return s.Region
}

func (s *AddImageResponseBodyPicInfoMultiRegion) SetRegion(v string) *AddImageResponseBodyPicInfoMultiRegion {
	s.Region = &v
	return s
}

func (s *AddImageResponseBodyPicInfoMultiRegion) Validate() error {
	return dara.Validate(s)
}
