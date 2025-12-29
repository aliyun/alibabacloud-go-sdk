// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelServiceCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelServiceCategoryResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelServiceCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelServiceCategoryResponseBody
	GetRequestId() *string
	SetResult(v []*ListHotelServiceCategoryResponseBodyResult) *ListHotelServiceCategoryResponseBody
	GetResult() []*ListHotelServiceCategoryResponseBodyResult
}

type ListHotelServiceCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5373C821-65D2-1764-B9F9-951914937FF5
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelServiceCategoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelServiceCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelServiceCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelServiceCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelServiceCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelServiceCategoryResponseBody) GetResult() []*ListHotelServiceCategoryResponseBodyResult {
	return s.Result
}

func (s *ListHotelServiceCategoryResponseBody) SetCode(v int32) *ListHotelServiceCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetMessage(v string) *ListHotelServiceCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetRequestId(v string) *ListHotelServiceCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) SetResult(v []*ListHotelServiceCategoryResponseBodyResult) *ListHotelServiceCategoryResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelServiceCategoryResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelServiceCategoryResponseBodyResult struct {
	// example:
	//
	// GOODS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 生活及洗漱用品
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingfenlei/wupintianjia.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 物品添加
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// HOTEL_SERVICE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelServiceCategoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelServiceCategoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelServiceCategoryResponseBodyResult) GetCode() *string {
	return s.Code
}

func (s *ListHotelServiceCategoryResponseBodyResult) GetDesc() *string {
	return s.Desc
}

func (s *ListHotelServiceCategoryResponseBodyResult) GetIcon() *string {
	return s.Icon
}

func (s *ListHotelServiceCategoryResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListHotelServiceCategoryResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetCode(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetDesc(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Desc = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetIcon(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetName(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) SetType(v string) *ListHotelServiceCategoryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListHotelServiceCategoryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
