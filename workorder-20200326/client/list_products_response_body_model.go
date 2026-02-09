// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListProductsResponseBody
	GetCode() *int32
	SetData(v *ListProductsResponseBodyData) *ListProductsResponseBody
	GetData() *ListProductsResponseBodyData
	SetMessage(v string) *ListProductsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProductsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProductsResponseBody
	GetSuccess() *bool
}

type ListProductsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListProductsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CA6204AC-6AA9-4CFA-9310-7DFD20C19EBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListProductsResponseBody) GetData() *ListProductsResponseBodyData {
	return s.Data
}

func (s *ListProductsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProductsResponseBody) SetCode(v int32) *ListProductsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductsResponseBody) SetData(v *ListProductsResponseBodyData) *ListProductsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductsResponseBody) SetMessage(v string) *ListProductsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetSuccess(v bool) *ListProductsResponseBody {
	s.Success = &v
	return s
}

func (s *ListProductsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductsResponseBodyData struct {
	ConsultationMore []*ListProductsResponseBodyDataConsultationMore `json:"ConsultationMore,omitempty" xml:"ConsultationMore,omitempty" type:"Repeated"`
	HotConsultation  []*ListProductsResponseBodyDataHotConsultation  `json:"HotConsultation,omitempty" xml:"HotConsultation,omitempty" type:"Repeated"`
	HotTech          []*ListProductsResponseBodyDataHotTech          `json:"HotTech,omitempty" xml:"HotTech,omitempty" type:"Repeated"`
	TechMore         []*ListProductsResponseBodyDataTechMore         `json:"TechMore,omitempty" xml:"TechMore,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyData) GetConsultationMore() []*ListProductsResponseBodyDataConsultationMore {
	return s.ConsultationMore
}

func (s *ListProductsResponseBodyData) GetHotConsultation() []*ListProductsResponseBodyDataHotConsultation {
	return s.HotConsultation
}

func (s *ListProductsResponseBodyData) GetHotTech() []*ListProductsResponseBodyDataHotTech {
	return s.HotTech
}

func (s *ListProductsResponseBodyData) GetTechMore() []*ListProductsResponseBodyDataTechMore {
	return s.TechMore
}

func (s *ListProductsResponseBodyData) SetConsultationMore(v []*ListProductsResponseBodyDataConsultationMore) *ListProductsResponseBodyData {
	s.ConsultationMore = v
	return s
}

func (s *ListProductsResponseBodyData) SetHotConsultation(v []*ListProductsResponseBodyDataHotConsultation) *ListProductsResponseBodyData {
	s.HotConsultation = v
	return s
}

func (s *ListProductsResponseBodyData) SetHotTech(v []*ListProductsResponseBodyDataHotTech) *ListProductsResponseBodyData {
	s.HotTech = v
	return s
}

func (s *ListProductsResponseBodyData) SetTechMore(v []*ListProductsResponseBodyDataTechMore) *ListProductsResponseBodyData {
	s.TechMore = v
	return s
}

func (s *ListProductsResponseBodyData) Validate() error {
	if s.ConsultationMore != nil {
		for _, item := range s.ConsultationMore {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HotConsultation != nil {
		for _, item := range s.HotConsultation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HotTech != nil {
		for _, item := range s.HotTech {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TechMore != nil {
		for _, item := range s.TechMore {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsResponseBodyDataConsultationMore struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// account
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyDataConsultationMore) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyDataConsultationMore) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataConsultationMore) GetDescription() *string {
	return s.Description
}

func (s *ListProductsResponseBodyDataConsultationMore) GetName() *string {
	return s.Name
}

func (s *ListProductsResponseBodyDataConsultationMore) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductsResponseBodyDataConsultationMore) SetDescription(v string) *ListProductsResponseBodyDataConsultationMore {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataConsultationMore) SetName(v string) *ListProductsResponseBodyDataConsultationMore {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataConsultationMore) SetProductCode(v string) *ListProductsResponseBodyDataConsultationMore {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyDataConsultationMore) Validate() error {
	return dara.Validate(s)
}

type ListProductsResponseBodyDataHotConsultation struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// agent
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyDataHotConsultation) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyDataHotConsultation) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataHotConsultation) GetDescription() *string {
	return s.Description
}

func (s *ListProductsResponseBodyDataHotConsultation) GetName() *string {
	return s.Name
}

func (s *ListProductsResponseBodyDataHotConsultation) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductsResponseBodyDataHotConsultation) SetDescription(v string) *ListProductsResponseBodyDataHotConsultation {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataHotConsultation) SetName(v string) *ListProductsResponseBodyDataHotConsultation {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataHotConsultation) SetProductCode(v string) *ListProductsResponseBodyDataHotConsultation {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyDataHotConsultation) Validate() error {
	return dara.Validate(s)
}

type ListProductsResponseBodyDataHotTech struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyDataHotTech) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyDataHotTech) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataHotTech) GetDescription() *string {
	return s.Description
}

func (s *ListProductsResponseBodyDataHotTech) GetName() *string {
	return s.Name
}

func (s *ListProductsResponseBodyDataHotTech) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductsResponseBodyDataHotTech) SetDescription(v string) *ListProductsResponseBodyDataHotTech {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataHotTech) SetName(v string) *ListProductsResponseBodyDataHotTech {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataHotTech) SetProductCode(v string) *ListProductsResponseBodyDataHotTech {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyDataHotTech) Validate() error {
	return dara.Validate(s)
}

type ListProductsResponseBodyDataTechMore struct {
	GroupName   *string                                            `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProductList []*ListProductsResponseBodyDataTechMoreProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyDataTechMore) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyDataTechMore) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataTechMore) GetGroupName() *string {
	return s.GroupName
}

func (s *ListProductsResponseBodyDataTechMore) GetProductList() []*ListProductsResponseBodyDataTechMoreProductList {
	return s.ProductList
}

func (s *ListProductsResponseBodyDataTechMore) SetGroupName(v string) *ListProductsResponseBodyDataTechMore {
	s.GroupName = &v
	return s
}

func (s *ListProductsResponseBodyDataTechMore) SetProductList(v []*ListProductsResponseBodyDataTechMoreProductList) *ListProductsResponseBodyDataTechMore {
	s.ProductList = v
	return s
}

func (s *ListProductsResponseBodyDataTechMore) Validate() error {
	if s.ProductList != nil {
		for _, item := range s.ProductList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductsResponseBodyDataTechMoreProductList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyDataTechMoreProductList) String() string {
	return dara.Prettify(s)
}

func (s ListProductsResponseBodyDataTechMoreProductList) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyDataTechMoreProductList) GetDescription() *string {
	return s.Description
}

func (s *ListProductsResponseBodyDataTechMoreProductList) GetName() *string {
	return s.Name
}

func (s *ListProductsResponseBodyDataTechMoreProductList) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductsResponseBodyDataTechMoreProductList) SetDescription(v string) *ListProductsResponseBodyDataTechMoreProductList {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyDataTechMoreProductList) SetName(v string) *ListProductsResponseBodyDataTechMoreProductList {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyDataTechMoreProductList) SetProductCode(v string) *ListProductsResponseBodyDataTechMoreProductList {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyDataTechMoreProductList) Validate() error {
	return dara.Validate(s)
}
