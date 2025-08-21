// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v *CreateCategoryResponseBodyCategory) *CreateCategoryResponseBody
	GetCategory() *CreateCategoryResponseBodyCategory
	SetRequestId(v string) *CreateCategoryResponseBody
	GetRequestId() *string
}

type CreateCategoryResponseBody struct {
	Category *CreateCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// example:
	//
	// A629A28F-F25E-5572-A679-FA46FB0151D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponseBody) GetCategory() *CreateCategoryResponseBodyCategory {
	return s.Category
}

func (s *CreateCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCategoryResponseBody) SetCategory(v *CreateCategoryResponseBodyCategory) *CreateCategoryResponseBody {
	s.Category = v
	return s
}

func (s *CreateCategoryResponseBody) SetRequestId(v string) *CreateCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateCategoryResponseBodyCategory struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 30000049006
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateCategoryResponseBodyCategory) String() string {
	return dara.Prettify(s)
}

func (s CreateCategoryResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *CreateCategoryResponseBodyCategory) GetBizCode() *string {
	return s.BizCode
}

func (s *CreateCategoryResponseBodyCategory) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreateCategoryResponseBodyCategory) GetName() *string {
	return s.Name
}

func (s *CreateCategoryResponseBodyCategory) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *CreateCategoryResponseBodyCategory) GetStatus() *int32 {
	return s.Status
}

func (s *CreateCategoryResponseBodyCategory) SetBizCode(v string) *CreateCategoryResponseBodyCategory {
	s.BizCode = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) SetCategoryId(v int64) *CreateCategoryResponseBodyCategory {
	s.CategoryId = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) SetName(v string) *CreateCategoryResponseBodyCategory {
	s.Name = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) SetParentCategoryId(v int64) *CreateCategoryResponseBodyCategory {
	s.ParentCategoryId = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) SetStatus(v int32) *CreateCategoryResponseBodyCategory {
	s.Status = &v
	return s
}

func (s *CreateCategoryResponseBodyCategory) Validate() error {
	return dara.Validate(s)
}
