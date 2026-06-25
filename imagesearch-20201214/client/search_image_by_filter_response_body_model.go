// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuctions(v []*SearchImageByFilterResponseBodyAuctions) *SearchImageByFilterResponseBody
	GetAuctions() []*SearchImageByFilterResponseBodyAuctions
	SetCode(v int32) *SearchImageByFilterResponseBody
	GetCode() *int32
	SetMsg(v string) *SearchImageByFilterResponseBody
	GetMsg() *string
	SetRequestId(v string) *SearchImageByFilterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchImageByFilterResponseBody
	GetSuccess() *bool
}

type SearchImageByFilterResponseBody struct {
	// The product description information returned.
	Auctions []*SearchImageByFilterResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
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
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
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

func (s SearchImageByFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByFilterResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByFilterResponseBody) GetAuctions() []*SearchImageByFilterResponseBodyAuctions {
	return s.Auctions
}

func (s *SearchImageByFilterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SearchImageByFilterResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *SearchImageByFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchImageByFilterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchImageByFilterResponseBody) SetAuctions(v []*SearchImageByFilterResponseBodyAuctions) *SearchImageByFilterResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByFilterResponseBody) SetCode(v int32) *SearchImageByFilterResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByFilterResponseBody) SetMsg(v string) *SearchImageByFilterResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByFilterResponseBody) SetRequestId(v string) *SearchImageByFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByFilterResponseBody) SetSuccess(v bool) *SearchImageByFilterResponseBody {
	s.Success = &v
	return s
}

func (s *SearchImageByFilterResponseBody) Validate() error {
	if s.Auctions != nil {
		for _, item := range s.Auctions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchImageByFilterResponseBodyAuctions struct {
	// The image category.
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
	// The integer type attribute.
	//
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// The integer type attribute.
	//
	// example:
	//
	// 2
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	// The integer type attribute.
	//
	// example:
	//
	// 2
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	// The integer type attribute.
	//
	// example:
	//
	// 2
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
	// The string type attribute.
	//
	// example:
	//
	// test
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// The string type attribute.
	//
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	// The string type attribute.
	//
	// example:
	//
	// test
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	// The string type attribute.
	//
	// example:
	//
	// test
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s SearchImageByFilterResponseBodyAuctions) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByFilterResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByFilterResponseBodyAuctions) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *SearchImageByFilterResponseBodyAuctions) GetCustomContent() *string {
	return s.CustomContent
}

func (s *SearchImageByFilterResponseBodyAuctions) GetIntAttr() *int32 {
	return s.IntAttr
}

func (s *SearchImageByFilterResponseBodyAuctions) GetIntAttr2() *int32 {
	return s.IntAttr2
}

func (s *SearchImageByFilterResponseBodyAuctions) GetIntAttr3() *int32 {
	return s.IntAttr3
}

func (s *SearchImageByFilterResponseBodyAuctions) GetIntAttr4() *int32 {
	return s.IntAttr4
}

func (s *SearchImageByFilterResponseBodyAuctions) GetPicName() *string {
	return s.PicName
}

func (s *SearchImageByFilterResponseBodyAuctions) GetProductId() *string {
	return s.ProductId
}

func (s *SearchImageByFilterResponseBodyAuctions) GetStrAttr() *string {
	return s.StrAttr
}

func (s *SearchImageByFilterResponseBodyAuctions) GetStrAttr2() *string {
	return s.StrAttr2
}

func (s *SearchImageByFilterResponseBodyAuctions) GetStrAttr3() *string {
	return s.StrAttr3
}

func (s *SearchImageByFilterResponseBodyAuctions) GetStrAttr4() *string {
	return s.StrAttr4
}

func (s *SearchImageByFilterResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByFilterResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetCustomContent(v string) *SearchImageByFilterResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByFilterResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetIntAttr2(v int32) *SearchImageByFilterResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetIntAttr3(v int32) *SearchImageByFilterResponseBodyAuctions {
	s.IntAttr3 = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetIntAttr4(v int32) *SearchImageByFilterResponseBodyAuctions {
	s.IntAttr4 = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetPicName(v string) *SearchImageByFilterResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetProductId(v string) *SearchImageByFilterResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetStrAttr(v string) *SearchImageByFilterResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetStrAttr2(v string) *SearchImageByFilterResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetStrAttr3(v string) *SearchImageByFilterResponseBodyAuctions {
	s.StrAttr3 = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) SetStrAttr4(v string) *SearchImageByFilterResponseBodyAuctions {
	s.StrAttr4 = &v
	return s
}

func (s *SearchImageByFilterResponseBodyAuctions) Validate() error {
	return dara.Validate(s)
}
