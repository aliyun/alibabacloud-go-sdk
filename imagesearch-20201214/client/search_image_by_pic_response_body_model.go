// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByPicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuctions(v []*SearchImageByPicResponseBodyAuctions) *SearchImageByPicResponseBody
	GetAuctions() []*SearchImageByPicResponseBodyAuctions
	SetCode(v int32) *SearchImageByPicResponseBody
	GetCode() *int32
	SetHead(v *SearchImageByPicResponseBodyHead) *SearchImageByPicResponseBody
	GetHead() *SearchImageByPicResponseBodyHead
	SetMsg(v string) *SearchImageByPicResponseBody
	GetMsg() *string
	SetPicInfo(v *SearchImageByPicResponseBodyPicInfo) *SearchImageByPicResponseBody
	GetPicInfo() *SearchImageByPicResponseBodyPicInfo
	SetRequestId(v string) *SearchImageByPicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchImageByPicResponseBody
	GetSuccess() *bool
}

type SearchImageByPicResponseBody struct {
	// The product descriptions returned.
	Auctions []*SearchImageByPicResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
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
	Head *SearchImageByPicResponseBodyHead `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The information such as the system-selected category and result of subject recognition.
	PicInfo *SearchImageByPicResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByPicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByPicResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBody) GetAuctions() []*SearchImageByPicResponseBodyAuctions {
	return s.Auctions
}

func (s *SearchImageByPicResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SearchImageByPicResponseBody) GetHead() *SearchImageByPicResponseBodyHead {
	return s.Head
}

func (s *SearchImageByPicResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SearchImageByPicResponseBody) GetPicInfo() *SearchImageByPicResponseBodyPicInfo {
	return s.PicInfo
}

func (s *SearchImageByPicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchImageByPicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchImageByPicResponseBody) SetAuctions(v []*SearchImageByPicResponseBodyAuctions) *SearchImageByPicResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByPicResponseBody) SetCode(v int32) *SearchImageByPicResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetHead(v *SearchImageByPicResponseBodyHead) *SearchImageByPicResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByPicResponseBody) SetMsg(v string) *SearchImageByPicResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetPicInfo(v *SearchImageByPicResponseBodyPicInfo) *SearchImageByPicResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchImageByPicResponseBody) SetRequestId(v string) *SearchImageByPicResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetSuccess(v bool) *SearchImageByPicResponseBody {
	s.Success = &v
	return s
}

func (s *SearchImageByPicResponseBody) Validate() error {
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

type SearchImageByPicResponseBodyAuctions struct {
	// The category of the image.
	//
	// example:
	//
	// 8888888
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
	// The similarity score of the searched image. Valid values: 0 to 1.
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
	// 2
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s SearchImageByPicResponseBodyAuctions) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByPicResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyAuctions) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *SearchImageByPicResponseBodyAuctions) GetCustomContent() *string {
	return s.CustomContent
}

func (s *SearchImageByPicResponseBodyAuctions) GetIntAttr() *int32 {
	return s.IntAttr
}

func (s *SearchImageByPicResponseBodyAuctions) GetIntAttr2() *int32 {
	return s.IntAttr2
}

func (s *SearchImageByPicResponseBodyAuctions) GetIntAttr3() *int32 {
	return s.IntAttr3
}

func (s *SearchImageByPicResponseBodyAuctions) GetIntAttr4() *int32 {
	return s.IntAttr4
}

func (s *SearchImageByPicResponseBodyAuctions) GetPicName() *string {
	return s.PicName
}

func (s *SearchImageByPicResponseBodyAuctions) GetProductId() *string {
	return s.ProductId
}

func (s *SearchImageByPicResponseBodyAuctions) GetScore() *float32 {
	return s.Score
}

func (s *SearchImageByPicResponseBodyAuctions) GetSortExprValues() *string {
	return s.SortExprValues
}

func (s *SearchImageByPicResponseBodyAuctions) GetStrAttr() *string {
	return s.StrAttr
}

func (s *SearchImageByPicResponseBodyAuctions) GetStrAttr2() *string {
	return s.StrAttr2
}

func (s *SearchImageByPicResponseBodyAuctions) GetStrAttr3() *string {
	return s.StrAttr3
}

func (s *SearchImageByPicResponseBodyAuctions) GetStrAttr4() *string {
	return s.StrAttr4
}

func (s *SearchImageByPicResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByPicResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetCustomContent(v string) *SearchImageByPicResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr2(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr3(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr3 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr4(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr4 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetPicName(v string) *SearchImageByPicResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetProductId(v string) *SearchImageByPicResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetScore(v float32) *SearchImageByPicResponseBodyAuctions {
	s.Score = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetSortExprValues(v string) *SearchImageByPicResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr2(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr3(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr3 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr4(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr4 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) Validate() error {
	return dara.Validate(s)
}

type SearchImageByPicResponseBodyHead struct {
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

func (s SearchImageByPicResponseBodyHead) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByPicResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyHead) GetDocsFound() *int32 {
	return s.DocsFound
}

func (s *SearchImageByPicResponseBodyHead) GetDocsReturn() *int32 {
	return s.DocsReturn
}

func (s *SearchImageByPicResponseBodyHead) GetSearchTime() *int32 {
	return s.SearchTime
}

func (s *SearchImageByPicResponseBodyHead) SetDocsFound(v int32) *SearchImageByPicResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchImageByPicResponseBodyHead) SetDocsReturn(v int32) *SearchImageByPicResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchImageByPicResponseBodyHead) SetSearchTime(v int32) *SearchImageByPicResponseBodyHead {
	s.SearchTime = &v
	return s
}

func (s *SearchImageByPicResponseBodyHead) Validate() error {
	return dara.Validate(s)
}

type SearchImageByPicResponseBodyPicInfo struct {
	// The categories that are supported by the system.
	AllCategories []*SearchImageByPicResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	// The category selected by the system. If a category is specified in the request, the specified category prevails.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The recognized subjects.
	//
	// >  To use this feature, you must upgrade the SDK to version 3.1.1.
	MultiRegion []*SearchImageByPicResponseBodyPicInfoMultiRegion `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	// The result of subject recognition. The subject area of the image, in the format of x1,x2,y1,y2. Specifically, x1 and y1 specify the upper-left pixel, and x2 and y2 specify the lower-right pixel. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfo) GetAllCategories() []*SearchImageByPicResponseBodyPicInfoAllCategories {
	return s.AllCategories
}

func (s *SearchImageByPicResponseBodyPicInfo) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *SearchImageByPicResponseBodyPicInfo) GetMultiRegion() []*SearchImageByPicResponseBodyPicInfoMultiRegion {
	return s.MultiRegion
}

func (s *SearchImageByPicResponseBodyPicInfo) GetRegion() *string {
	return s.Region
}

func (s *SearchImageByPicResponseBodyPicInfo) SetAllCategories(v []*SearchImageByPicResponseBodyPicInfoAllCategories) *SearchImageByPicResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetCategoryId(v int32) *SearchImageByPicResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetMultiRegion(v []*SearchImageByPicResponseBodyPicInfoMultiRegion) *SearchImageByPicResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetRegion(v string) *SearchImageByPicResponseBodyPicInfo {
	s.Region = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) Validate() error {
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

type SearchImageByPicResponseBodyPicInfoAllCategories struct {
	// The ID of the category.
	//
	// example:
	//
	// 88888888
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// other
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfoAllCategories) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) GetId() *int32 {
	return s.Id
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) GetName() *string {
	return s.Name
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByPicResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByPicResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) Validate() error {
	return dara.Validate(s)
}

type SearchImageByPicResponseBodyPicInfoMultiRegion struct {
	// The result of subject recognition. The subject area of the image, in the format of x1,x2,y1,y2. Specifically, x1 and y1 specify the upper-left pixel, and x2 and y2 specify the lower-right pixel. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfoMultiRegion) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfoMultiRegion) GetRegion() *string {
	return s.Region
}

func (s *SearchImageByPicResponseBodyPicInfoMultiRegion) SetRegion(v string) *SearchImageByPicResponseBodyPicInfoMultiRegion {
	s.Region = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfoMultiRegion) Validate() error {
	return dara.Validate(s)
}
