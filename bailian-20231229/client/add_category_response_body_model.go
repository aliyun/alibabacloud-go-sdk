// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddCategoryResponseBody
	GetCode() *string
	SetData(v *AddCategoryResponseBodyData) *AddCategoryResponseBody
	GetData() *AddCategoryResponseBodyData
	SetMessage(v string) *AddCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCategoryResponseBody
	GetRequestId() *string
	SetStatus(v string) *AddCategoryResponseBody
	GetStatus() *string
	SetSuccess(v bool) *AddCategoryResponseBody
	GetSuccess() *bool
}

type AddCategoryResponseBody struct {
	// example:
	//
	// success
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 778C0B3B-03C1-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddCategoryResponseBody) GetData() *AddCategoryResponseBodyData {
	return s.Data
}

func (s *AddCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCategoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *AddCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddCategoryResponseBody) SetCode(v string) *AddCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *AddCategoryResponseBody) SetData(v *AddCategoryResponseBodyData) *AddCategoryResponseBody {
	s.Data = v
	return s
}

func (s *AddCategoryResponseBody) SetMessage(v string) *AddCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *AddCategoryResponseBody) SetRequestId(v string) *AddCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCategoryResponseBody) SetStatus(v string) *AddCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *AddCategoryResponseBody) SetSuccess(v bool) *AddCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *AddCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddCategoryResponseBodyData struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s AddCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddCategoryResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *AddCategoryResponseBodyData) GetCategoryName() *string {
	return s.CategoryName
}

func (s *AddCategoryResponseBodyData) SetCategoryId(v string) *AddCategoryResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *AddCategoryResponseBodyData) SetCategoryName(v string) *AddCategoryResponseBodyData {
	s.CategoryName = &v
	return s
}

func (s *AddCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
