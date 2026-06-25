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
	// The descriptions of all returned products.
	Auctions []*SearchImageByNameResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	// The error code. Valid values:
	//
	// - 0: success.
	//
	// - Non-zero: failure.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The overview of the search results.
	Head *SearchImageByNameResponseBodyHead `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The information such as category prediction and subject identification results.
	PicInfo *SearchImageByNameResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	// The image category.
	//
	// example:
	//
	// 20
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The custom content defined by the user.
	//
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The integer attribute.
	//
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// The integer attribute.
	//
	// example:
	//
	// 20
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	// The integer attribute. This field can be used for filtering during queries and is returned in query results.
	//
	// example:
	//
	// 1
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	// The integer attribute. This field can be used for filtering during queries and is returned in query results.
	//
	// example:
	//
	// 1
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The image name.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The product ID.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The image similarity score. Valid values: 0 to 1.
	//
	// > You must upgrade to V3.1.1 or later to use this field.
	//
	// example:
	//
	// 1
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The system scoring information.
	//
	// > - This field is deprecated. Use Score instead.
	//
	// - SortExprValues is a semicolon-separated pair. The first value indicates the relevance score of the image. A higher value indicates higher relevance to the query image. The scoring varies depending on the algorithm model.
	//
	// - When the category is 0 to 2, the value range of SortExprValues is 0 to 7.33136443711219e+24.
	//
	// - For other category values, the value range of SortExprValues is 0 to 5.37633353624177e+24. This score reaches its maximum when two images are identical.
	//
	// example:
	//
	// 5.37633353624177e+24;0
	SortExprValues *string `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	// The string attribute.
	//
	// example:
	//
	// ss
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// The string attribute.
	//
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	// The string attribute. The maximum length is 128 characters. This field can be used for filtering during queries and is returned in query results.
	//
	// example:
	//
	// test
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	// The string attribute. The maximum length is 128 characters. This field can be used for filtering during queries and is returned in query results.
	//
	// example:
	//
	// test
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
	// The number of results returned.
	//
	// example:
	//
	// 10
	DocsFound *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	// The number of matched results in the instance.
	//
	// example:
	//
	// 10000
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	// The search duration, in milliseconds.
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
	// The information about all categories supported by the system.
	AllCategories []*SearchImageByNameResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	// The category prediction result.
	//
	// If the user specifies a category in the request, the specified category is used.
	//
	// example:
	//
	// 20
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The subject identification collection.
	MultiRegion []*SearchImageByNameResponseBodyPicInfoMultiRegion `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	// The subject identification result.
	//
	// The subject region of the image, in the format of x1,x2,y1,y2, where x1,y1 is the upper-left point and x2,y2 is the lower-right point. If the user specifies a subject region in the request, the specified region is used.
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
	// other
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
	// The subject identification result.
	//
	// The subject region of the image, in the format of x1,x2,y1,y2, where x1,y1 is the upper-left point and x2,y2 is the lower-right point. If the user specifies a subject region in the request, the specified region is used.
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
