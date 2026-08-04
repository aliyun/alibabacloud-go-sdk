// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomerCategoryResponseBody
	GetCode() *string
	SetData(v *GetCustomerCategoryResponseBodyData) *GetCustomerCategoryResponseBody
	GetData() *GetCustomerCategoryResponseBodyData
	SetMessage(v string) *GetCustomerCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomerCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomerCategoryResponseBody
	GetSuccess() *bool
}

type GetCustomerCategoryResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCustomerCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomerCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerCategoryResponseBody) GetData() *GetCustomerCategoryResponseBodyData {
	return s.Data
}

func (s *GetCustomerCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomerCategoryResponseBody) SetCode(v string) *GetCustomerCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerCategoryResponseBody) SetData(v *GetCustomerCategoryResponseBodyData) *GetCustomerCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerCategoryResponseBody) SetMessage(v string) *GetCustomerCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerCategoryResponseBody) SetRequestId(v string) *GetCustomerCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerCategoryResponseBody) SetSuccess(v bool) *GetCustomerCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomerCategoryResponseBodyData struct {
	BizCategory []*GetCustomerCategoryResponseBodyDataBizCategory `json:"BizCategory,omitempty" xml:"BizCategory,omitempty" type:"Repeated"`
}

func (s GetCustomerCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryResponseBodyData) GetBizCategory() []*GetCustomerCategoryResponseBodyDataBizCategory {
	return s.BizCategory
}

func (s *GetCustomerCategoryResponseBodyData) SetBizCategory(v []*GetCustomerCategoryResponseBodyDataBizCategory) *GetCustomerCategoryResponseBodyData {
	s.BizCategory = v
	return s
}

func (s *GetCustomerCategoryResponseBodyData) Validate() error {
	if s.BizCategory != nil {
		for _, item := range s.BizCategory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCustomerCategoryResponseBodyDataBizCategory struct {
	Code       *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	IsCheck    *bool                                                     `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	MainBiz    *bool                                                     `json:"MainBiz,omitempty" xml:"MainBiz,omitempty"`
	Name       *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Other      *string                                                   `json:"Other,omitempty" xml:"Other,omitempty"`
	SubConfigs *GetCustomerCategoryResponseBodyDataBizCategorySubConfigs `json:"SubConfigs,omitempty" xml:"SubConfigs,omitempty" type:"Struct"`
}

func (s GetCustomerCategoryResponseBodyDataBizCategory) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryResponseBodyDataBizCategory) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) GetCode() *string {
	return s.Code
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) GetIsCheck() *bool {
	return s.IsCheck
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) GetMainBiz() *bool {
	return s.MainBiz
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) GetName() *string {
	return s.Name
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) GetOther() *string {
	return s.Other
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) GetSubConfigs() *GetCustomerCategoryResponseBodyDataBizCategorySubConfigs {
	return s.SubConfigs
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) SetCode(v string) *GetCustomerCategoryResponseBodyDataBizCategory {
	s.Code = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) SetIsCheck(v bool) *GetCustomerCategoryResponseBodyDataBizCategory {
	s.IsCheck = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) SetMainBiz(v bool) *GetCustomerCategoryResponseBodyDataBizCategory {
	s.MainBiz = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) SetName(v string) *GetCustomerCategoryResponseBodyDataBizCategory {
	s.Name = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) SetOther(v string) *GetCustomerCategoryResponseBodyDataBizCategory {
	s.Other = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) SetSubConfigs(v *GetCustomerCategoryResponseBodyDataBizCategorySubConfigs) *GetCustomerCategoryResponseBodyDataBizCategory {
	s.SubConfigs = v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategory) Validate() error {
	if s.SubConfigs != nil {
		if err := s.SubConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomerCategoryResponseBodyDataBizCategorySubConfigs struct {
	BizSubCategory []*GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory `json:"BizSubCategory,omitempty" xml:"BizSubCategory,omitempty" type:"Repeated"`
}

func (s GetCustomerCategoryResponseBodyDataBizCategorySubConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryResponseBodyDataBizCategorySubConfigs) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigs) GetBizSubCategory() []*GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory {
	return s.BizSubCategory
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigs) SetBizSubCategory(v []*GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) *GetCustomerCategoryResponseBodyDataBizCategorySubConfigs {
	s.BizSubCategory = v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigs) Validate() error {
	if s.BizSubCategory != nil {
		for _, item := range s.BizSubCategory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsCheck *bool   `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	MainBiz *bool   `json:"MainBiz,omitempty" xml:"MainBiz,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Other   *string `json:"Other,omitempty" xml:"Other,omitempty"`
}

func (s GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) GetCode() *string {
	return s.Code
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) GetIsCheck() *bool {
	return s.IsCheck
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) GetMainBiz() *bool {
	return s.MainBiz
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) GetName() *string {
	return s.Name
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) GetOther() *string {
	return s.Other
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) SetCode(v string) *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.Code = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) SetIsCheck(v bool) *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.IsCheck = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) SetMainBiz(v bool) *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.MainBiz = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) SetName(v string) *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.Name = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) SetOther(v string) *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.Other = &v
	return s
}

func (s *GetCustomerCategoryResponseBodyDataBizCategorySubConfigsBizSubCategory) Validate() error {
	return dara.Validate(s)
}
