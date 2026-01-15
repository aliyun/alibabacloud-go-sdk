// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuctions(v []*SearchImageByNameResponseBodyAuctions) *SearchImageByNameResponseBody
	GetAuctions() []*SearchImageByNameResponseBodyAuctions
	SetCode(v int32) *SearchImageByNameResponseBody
	GetCode() *int32
	SetHead(v *SearchImageByNameResponseBodyHead) *SearchImageByNameResponseBody
	GetHead() *SearchImageByNameResponseBodyHead
	SetMsg(v string) *SearchImageByNameResponseBody
	GetMsg() *string
	SetPicInfo(v *SearchImageByNameResponseBodyPicInfo) *SearchImageByNameResponseBody
	GetPicInfo() *SearchImageByNameResponseBodyPicInfo
	SetRequestId(v string) *SearchImageByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchImageByNameResponseBody
	GetSuccess() *bool
}

type SearchImageByNameResponseBody struct {
	// The product descriptions returned.
	Auctions []*SearchImageByNameResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The summary of the search result.
	Head *SearchImageByNameResponseBodyHead `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The information such as the system-selected category and result of subject recognition.
	PicInfo *SearchImageByNameResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBody) GetAuctions() []*SearchImageByNameResponseBodyAuctions {
	return s.Auctions
}

func (s *SearchImageByNameResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SearchImageByNameResponseBody) GetHead() *SearchImageByNameResponseBodyHead {
	return s.Head
}

func (s *SearchImageByNameResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SearchImageByNameResponseBody) GetPicInfo() *SearchImageByNameResponseBodyPicInfo {
	return s.PicInfo
}

func (s *SearchImageByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchImageByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchImageByNameResponseBody) SetAuctions(v []*SearchImageByNameResponseBodyAuctions) *SearchImageByNameResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByNameResponseBody) SetCode(v int32) *SearchImageByNameResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetHead(v *SearchImageByNameResponseBodyHead) *SearchImageByNameResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByNameResponseBody) SetMsg(v string) *SearchImageByNameResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetPicInfo(v *SearchImageByNameResponseBodyPicInfo) *SearchImageByNameResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchImageByNameResponseBody) SetRequestId(v string) *SearchImageByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetSuccess(v bool) *SearchImageByNameResponseBody {
	s.Success = &v
	return s
}

func (s *SearchImageByNameResponseBody) Validate() error {
	if s.Auctions != nil {
		for _, item := range s.Auctions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Head != nil {
		if err := s.Head.Validate(); err != nil {
			return err
		}
	}
	if s.PicInfo != nil {
		if err := s.PicInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchImageByNameResponseBodyAuctions struct {
	// The category of the image.
	//
	// example:
	//
	// 20
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The user-defined content.
	//
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The attribute, which is an integer.
	//
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// example:
	//
	// 20
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The similarity score of the returned image. Valid values: 0 to 1.
	//
	// >  To use this feature, you must upgrade the SDK to version 3.1.1.
	//
	// example:
	//
	// 1
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The score information about the image.
	//
	// > 	- This parameter is not supported. We recommend that you use the Score parameter.
	//
	// >	- The SortExprValues parameter indicates a 2-tuple in which values are separated by a semicolon (;). The first value indicates the correlation score of the returned image. A greater value indicates a higher correlation with the sample image. Different algorithms are used.
	//
	// >	- If the value of CategoryId is within the value range from 0 to 2, the value range of SortExprValues is from 0 to 7.33136443711219e+24.
	//
	// >	- If the value of CategoryId is not within the value range from 0 to 2, the value range of SortExprValues is from 0 to 5.37633353624177e+24. If the returned image is identical with the sample image, the highest correlation score is generated.
	//
	// example:
	//
	// 5.37633353624177e+24;0
	SortExprValues *string `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	// The attribute, which is a string.
	//
	// example:
	//
	// ss
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s SearchImageByNameResponseBodyAuctions) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByNameResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyAuctions) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *SearchImageByNameResponseBodyAuctions) GetCustomContent() *string {
	return s.CustomContent
}

func (s *SearchImageByNameResponseBodyAuctions) GetIntAttr() *int32 {
	return s.IntAttr
}

func (s *SearchImageByNameResponseBodyAuctions) GetIntAttr2() *int32 {
	return s.IntAttr2
}

func (s *SearchImageByNameResponseBodyAuctions) GetIntAttr3() *int32 {
	return s.IntAttr3
}

func (s *SearchImageByNameResponseBodyAuctions) GetIntAttr4() *int32 {
	return s.IntAttr4
}

func (s *SearchImageByNameResponseBodyAuctions) GetPicName() *string {
	return s.PicName
}

func (s *SearchImageByNameResponseBodyAuctions) GetProductId() *string {
	return s.ProductId
}

func (s *SearchImageByNameResponseBodyAuctions) GetScore() *float32 {
	return s.Score
}

func (s *SearchImageByNameResponseBodyAuctions) GetSortExprValues() *string {
	return s.SortExprValues
}

func (s *SearchImageByNameResponseBodyAuctions) GetStrAttr() *string {
	return s.StrAttr
}

func (s *SearchImageByNameResponseBodyAuctions) GetStrAttr2() *string {
	return s.StrAttr2
}

func (s *SearchImageByNameResponseBodyAuctions) GetStrAttr3() *string {
	return s.StrAttr3
}

func (s *SearchImageByNameResponseBodyAuctions) GetStrAttr4() *string {
	return s.StrAttr4
}

func (s *SearchImageByNameResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByNameResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetCustomContent(v string) *SearchImageByNameResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr2(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr3(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr3 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr4(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr4 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetPicName(v string) *SearchImageByNameResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetProductId(v string) *SearchImageByNameResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetScore(v float32) *SearchImageByNameResponseBodyAuctions {
	s.Score = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetSortExprValues(v string) *SearchImageByNameResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr2(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr3(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr3 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr4(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr4 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) Validate() error {
	return dara.Validate(s)
}

type SearchImageByNameResponseBodyHead struct {
	// The number of images returned.
	//
	// example:
	//
	// 10
	DocsFound *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	// The number of images that match the search conditions on the Image Search instance.
	//
	// example:
	//
	// 10000
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	// The time it takes to complete the search process. Unit: milliseconds.
	//
	// example:
	//
	// 95
	SearchTime *int32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
}

func (s SearchImageByNameResponseBodyHead) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByNameResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyHead) GetDocsFound() *int32 {
	return s.DocsFound
}

func (s *SearchImageByNameResponseBodyHead) GetDocsReturn() *int32 {
	return s.DocsReturn
}

func (s *SearchImageByNameResponseBodyHead) GetSearchTime() *int32 {
	return s.SearchTime
}

func (s *SearchImageByNameResponseBodyHead) SetDocsFound(v int32) *SearchImageByNameResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchImageByNameResponseBodyHead) SetDocsReturn(v int32) *SearchImageByNameResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchImageByNameResponseBodyHead) SetSearchTime(v int32) *SearchImageByNameResponseBodyHead {
	s.SearchTime = &v
	return s
}

func (s *SearchImageByNameResponseBodyHead) Validate() error {
	return dara.Validate(s)
}

type SearchImageByNameResponseBodyPicInfo struct {
	// The categories that are supported by the system.
	AllCategories []*SearchImageByNameResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	// The category selected by the system.
	//
	// If a category is specified in the request, the specified category prevails.
	//
	// example:
	//
	// 20
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The recognized subjects.
	MultiRegion []*SearchImageByNameResponseBodyPicInfoMultiRegion `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	// The result of subject recognition.
	//
	// The subject area of the image, in the format of x1,x2,y1,y2. Specifically, x1 and y1 specify the upper-left pixel, and x2 and y2 specify the lower-right pixel. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfo) GetAllCategories() []*SearchImageByNameResponseBodyPicInfoAllCategories {
	return s.AllCategories
}

func (s *SearchImageByNameResponseBodyPicInfo) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *SearchImageByNameResponseBodyPicInfo) GetMultiRegion() []*SearchImageByNameResponseBodyPicInfoMultiRegion {
	return s.MultiRegion
}

func (s *SearchImageByNameResponseBodyPicInfo) GetRegion() *string {
	return s.Region
}

func (s *SearchImageByNameResponseBodyPicInfo) SetAllCategories(v []*SearchImageByNameResponseBodyPicInfoAllCategories) *SearchImageByNameResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetCategoryId(v int32) *SearchImageByNameResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetMultiRegion(v []*SearchImageByNameResponseBodyPicInfoMultiRegion) *SearchImageByNameResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetRegion(v string) *SearchImageByNameResponseBodyPicInfo {
	s.Region = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) Validate() error {
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

type SearchImageByNameResponseBodyPicInfoAllCategories struct {
	// The ID of the category.
	//
	// example:
	//
	// other
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// 88888888
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfoAllCategories) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) GetId() *int32 {
	return s.Id
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) GetName() *string {
	return s.Name
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByNameResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByNameResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) Validate() error {
	return dara.Validate(s)
}

type SearchImageByNameResponseBodyPicInfoMultiRegion struct {
	// The result of subject recognition.
	//
	// The subject area of the image, in the format of x1,x2,y1,y2. Specifically, x1 and y1 specify the upper-left pixel, and x2 and y2 specify the lower-right pixel. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfoMultiRegion) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfoMultiRegion) GetRegion() *string {
	return s.Region
}

func (s *SearchImageByNameResponseBodyPicInfoMultiRegion) SetRegion(v string) *SearchImageByNameResponseBodyPicInfoMultiRegion {
	s.Region = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfoMultiRegion) Validate() error {
	return dara.Validate(s)
}
