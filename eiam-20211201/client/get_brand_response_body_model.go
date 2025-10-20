// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBrand(v *GetBrandResponseBodyBrand) *GetBrandResponseBody
	GetBrand() *GetBrandResponseBodyBrand
	SetRequestId(v string) *GetBrandResponseBody
	GetRequestId() *string
}

type GetBrandResponseBody struct {
	Brand *GetBrandResponseBodyBrand `json:"Brand,omitempty" xml:"Brand,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBrandResponseBody) GoString() string {
	return s.String()
}

func (s *GetBrandResponseBody) GetBrand() *GetBrandResponseBodyBrand {
	return s.Brand
}

func (s *GetBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBrandResponseBody) SetBrand(v *GetBrandResponseBodyBrand) *GetBrandResponseBody {
	s.Brand = v
	return s
}

func (s *GetBrandResponseBody) SetRequestId(v string) *GetBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBrandResponseBody) Validate() error {
	if s.Brand != nil {
		if err := s.Brand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBrandResponseBodyBrand struct {
	// 品牌ID
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// 品牌名称
	//
	// example:
	//
	// Custom Brand
	BrandName *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	// 品牌类型
	//
	// example:
	//
	// user_custom
	BrandType *string `json:"BrandType,omitempty" xml:"BrandType,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 品牌状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBrandResponseBodyBrand) String() string {
	return dara.Prettify(s)
}

func (s GetBrandResponseBodyBrand) GoString() string {
	return s.String()
}

func (s *GetBrandResponseBodyBrand) GetBrandId() *string {
	return s.BrandId
}

func (s *GetBrandResponseBodyBrand) GetBrandName() *string {
	return s.BrandName
}

func (s *GetBrandResponseBodyBrand) GetBrandType() *string {
	return s.BrandType
}

func (s *GetBrandResponseBodyBrand) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBrandResponseBodyBrand) GetStatus() *string {
	return s.Status
}

func (s *GetBrandResponseBodyBrand) SetBrandId(v string) *GetBrandResponseBodyBrand {
	s.BrandId = &v
	return s
}

func (s *GetBrandResponseBodyBrand) SetBrandName(v string) *GetBrandResponseBodyBrand {
	s.BrandName = &v
	return s
}

func (s *GetBrandResponseBodyBrand) SetBrandType(v string) *GetBrandResponseBodyBrand {
	s.BrandType = &v
	return s
}

func (s *GetBrandResponseBodyBrand) SetInstanceId(v string) *GetBrandResponseBodyBrand {
	s.InstanceId = &v
	return s
}

func (s *GetBrandResponseBodyBrand) SetStatus(v string) *GetBrandResponseBodyBrand {
	s.Status = &v
	return s
}

func (s *GetBrandResponseBodyBrand) Validate() error {
	return dara.Validate(s)
}
