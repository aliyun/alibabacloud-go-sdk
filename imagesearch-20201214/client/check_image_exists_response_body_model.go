// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckImageExistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuctions(v *CheckImageExistsResponseBodyAuctions) *CheckImageExistsResponseBody
	GetAuctions() *CheckImageExistsResponseBodyAuctions
	SetCode(v int32) *CheckImageExistsResponseBody
	GetCode() *int32
	SetExists(v bool) *CheckImageExistsResponseBody
	GetExists() *bool
	SetMsg(v string) *CheckImageExistsResponseBody
	GetMsg() *string
	SetRequestId(v string) *CheckImageExistsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckImageExistsResponseBody
	GetSuccess() *bool
}

type CheckImageExistsResponseBody struct {
	Auctions *CheckImageExistsResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Exists *bool `json:"Exists,omitempty" xml:"Exists,omitempty"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckImageExistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckImageExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckImageExistsResponseBody) GetAuctions() *CheckImageExistsResponseBodyAuctions {
	return s.Auctions
}

func (s *CheckImageExistsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CheckImageExistsResponseBody) GetExists() *bool {
	return s.Exists
}

func (s *CheckImageExistsResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CheckImageExistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckImageExistsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckImageExistsResponseBody) SetAuctions(v *CheckImageExistsResponseBodyAuctions) *CheckImageExistsResponseBody {
	s.Auctions = v
	return s
}

func (s *CheckImageExistsResponseBody) SetCode(v int32) *CheckImageExistsResponseBody {
	s.Code = &v
	return s
}

func (s *CheckImageExistsResponseBody) SetExists(v bool) *CheckImageExistsResponseBody {
	s.Exists = &v
	return s
}

func (s *CheckImageExistsResponseBody) SetMsg(v string) *CheckImageExistsResponseBody {
	s.Msg = &v
	return s
}

func (s *CheckImageExistsResponseBody) SetRequestId(v string) *CheckImageExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckImageExistsResponseBody) SetSuccess(v bool) *CheckImageExistsResponseBody {
	s.Success = &v
	return s
}

func (s *CheckImageExistsResponseBody) Validate() error {
	if s.Auctions != nil {
		if err := s.Auctions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckImageExistsResponseBodyAuctions struct {
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// example:
	//
	// 2
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	// example:
	//
	// 2
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	// example:
	//
	// 2
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// test
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	// example:
	//
	// test
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	// example:
	//
	// test
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s CheckImageExistsResponseBodyAuctions) String() string {
	return dara.Prettify(s)
}

func (s CheckImageExistsResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *CheckImageExistsResponseBodyAuctions) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *CheckImageExistsResponseBodyAuctions) GetCustomContent() *string {
	return s.CustomContent
}

func (s *CheckImageExistsResponseBodyAuctions) GetIntAttr() *int32 {
	return s.IntAttr
}

func (s *CheckImageExistsResponseBodyAuctions) GetIntAttr2() *int32 {
	return s.IntAttr2
}

func (s *CheckImageExistsResponseBodyAuctions) GetIntAttr3() *int32 {
	return s.IntAttr3
}

func (s *CheckImageExistsResponseBodyAuctions) GetIntAttr4() *int32 {
	return s.IntAttr4
}

func (s *CheckImageExistsResponseBodyAuctions) GetPicName() *string {
	return s.PicName
}

func (s *CheckImageExistsResponseBodyAuctions) GetProductId() *string {
	return s.ProductId
}

func (s *CheckImageExistsResponseBodyAuctions) GetStrAttr() *string {
	return s.StrAttr
}

func (s *CheckImageExistsResponseBodyAuctions) GetStrAttr2() *string {
	return s.StrAttr2
}

func (s *CheckImageExistsResponseBodyAuctions) GetStrAttr3() *string {
	return s.StrAttr3
}

func (s *CheckImageExistsResponseBodyAuctions) GetStrAttr4() *string {
	return s.StrAttr4
}

func (s *CheckImageExistsResponseBodyAuctions) SetCategoryId(v int32) *CheckImageExistsResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetCustomContent(v string) *CheckImageExistsResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetIntAttr(v int32) *CheckImageExistsResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetIntAttr2(v int32) *CheckImageExistsResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetIntAttr3(v int32) *CheckImageExistsResponseBodyAuctions {
	s.IntAttr3 = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetIntAttr4(v int32) *CheckImageExistsResponseBodyAuctions {
	s.IntAttr4 = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetPicName(v string) *CheckImageExistsResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetProductId(v string) *CheckImageExistsResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetStrAttr(v string) *CheckImageExistsResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetStrAttr2(v string) *CheckImageExistsResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetStrAttr3(v string) *CheckImageExistsResponseBodyAuctions {
	s.StrAttr3 = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) SetStrAttr4(v string) *CheckImageExistsResponseBodyAuctions {
	s.StrAttr4 = &v
	return s
}

func (s *CheckImageExistsResponseBodyAuctions) Validate() error {
	return dara.Validate(s)
}
