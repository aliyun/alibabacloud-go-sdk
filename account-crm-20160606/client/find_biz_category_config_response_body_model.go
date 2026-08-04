// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindBizCategoryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindBizCategoryConfigResponseBody
	GetCode() *string
	SetData(v *FindBizCategoryConfigResponseBodyData) *FindBizCategoryConfigResponseBody
	GetData() *FindBizCategoryConfigResponseBodyData
	SetMessage(v string) *FindBizCategoryConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindBizCategoryConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindBizCategoryConfigResponseBody
	GetSuccess() *bool
}

type FindBizCategoryConfigResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *FindBizCategoryConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindBizCategoryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindBizCategoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *FindBizCategoryConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindBizCategoryConfigResponseBody) GetData() *FindBizCategoryConfigResponseBodyData {
	return s.Data
}

func (s *FindBizCategoryConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindBizCategoryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindBizCategoryConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindBizCategoryConfigResponseBody) SetCode(v string) *FindBizCategoryConfigResponseBody {
	s.Code = &v
	return s
}

func (s *FindBizCategoryConfigResponseBody) SetData(v *FindBizCategoryConfigResponseBodyData) *FindBizCategoryConfigResponseBody {
	s.Data = v
	return s
}

func (s *FindBizCategoryConfigResponseBody) SetMessage(v string) *FindBizCategoryConfigResponseBody {
	s.Message = &v
	return s
}

func (s *FindBizCategoryConfigResponseBody) SetRequestId(v string) *FindBizCategoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindBizCategoryConfigResponseBody) SetSuccess(v bool) *FindBizCategoryConfigResponseBody {
	s.Success = &v
	return s
}

func (s *FindBizCategoryConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindBizCategoryConfigResponseBodyData struct {
	BizCategory []*FindBizCategoryConfigResponseBodyDataBizCategory `json:"BizCategory,omitempty" xml:"BizCategory,omitempty" type:"Repeated"`
}

func (s FindBizCategoryConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FindBizCategoryConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindBizCategoryConfigResponseBodyData) GetBizCategory() []*FindBizCategoryConfigResponseBodyDataBizCategory {
	return s.BizCategory
}

func (s *FindBizCategoryConfigResponseBodyData) SetBizCategory(v []*FindBizCategoryConfigResponseBodyDataBizCategory) *FindBizCategoryConfigResponseBodyData {
	s.BizCategory = v
	return s
}

func (s *FindBizCategoryConfigResponseBodyData) Validate() error {
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

type FindBizCategoryConfigResponseBodyDataBizCategory struct {
	Code       *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	IsCheck    *bool                                                       `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	MainBiz    *bool                                                       `json:"MainBiz,omitempty" xml:"MainBiz,omitempty"`
	Name       *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Other      *string                                                     `json:"Other,omitempty" xml:"Other,omitempty"`
	SubConfigs *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs `json:"SubConfigs,omitempty" xml:"SubConfigs,omitempty" type:"Struct"`
}

func (s FindBizCategoryConfigResponseBodyDataBizCategory) String() string {
	return dara.Prettify(s)
}

func (s FindBizCategoryConfigResponseBodyDataBizCategory) GoString() string {
	return s.String()
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) GetCode() *string {
	return s.Code
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) GetIsCheck() *bool {
	return s.IsCheck
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) GetMainBiz() *bool {
	return s.MainBiz
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) GetName() *string {
	return s.Name
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) GetOther() *string {
	return s.Other
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) GetSubConfigs() *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs {
	return s.SubConfigs
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) SetCode(v string) *FindBizCategoryConfigResponseBodyDataBizCategory {
	s.Code = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) SetIsCheck(v bool) *FindBizCategoryConfigResponseBodyDataBizCategory {
	s.IsCheck = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) SetMainBiz(v bool) *FindBizCategoryConfigResponseBodyDataBizCategory {
	s.MainBiz = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) SetName(v string) *FindBizCategoryConfigResponseBodyDataBizCategory {
	s.Name = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) SetOther(v string) *FindBizCategoryConfigResponseBodyDataBizCategory {
	s.Other = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) SetSubConfigs(v *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs) *FindBizCategoryConfigResponseBodyDataBizCategory {
	s.SubConfigs = v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategory) Validate() error {
	if s.SubConfigs != nil {
		if err := s.SubConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs struct {
	BizSubCategory []*FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory `json:"BizSubCategory,omitempty" xml:"BizSubCategory,omitempty" type:"Repeated"`
}

func (s FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs) String() string {
	return dara.Prettify(s)
}

func (s FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs) GoString() string {
	return s.String()
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs) GetBizSubCategory() []*FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory {
	return s.BizSubCategory
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs) SetBizSubCategory(v []*FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs {
	s.BizSubCategory = v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigs) Validate() error {
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

type FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsCheck *bool   `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	MainBiz *bool   `json:"MainBiz,omitempty" xml:"MainBiz,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Other   *string `json:"Other,omitempty" xml:"Other,omitempty"`
}

func (s FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) String() string {
	return dara.Prettify(s)
}

func (s FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) GoString() string {
	return s.String()
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) GetCode() *string {
	return s.Code
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) GetIsCheck() *bool {
	return s.IsCheck
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) GetMainBiz() *bool {
	return s.MainBiz
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) GetName() *string {
	return s.Name
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) GetOther() *string {
	return s.Other
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) SetCode(v string) *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.Code = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) SetIsCheck(v bool) *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.IsCheck = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) SetMainBiz(v bool) *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.MainBiz = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) SetName(v string) *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.Name = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) SetOther(v string) *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory {
	s.Other = &v
	return s
}

func (s *FindBizCategoryConfigResponseBodyDataBizCategorySubConfigsBizSubCategory) Validate() error {
	return dara.Validate(s)
}
