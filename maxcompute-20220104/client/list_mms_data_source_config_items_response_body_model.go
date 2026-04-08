// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDataSourceConfigItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMmsDataSourceConfigItemsResponseBodyData) *ListMmsDataSourceConfigItemsResponseBody
	GetData() []*ListMmsDataSourceConfigItemsResponseBodyData
	SetRequestId(v string) *ListMmsDataSourceConfigItemsResponseBody
	GetRequestId() *string
}

type ListMmsDataSourceConfigItemsResponseBody struct {
	Data []*ListMmsDataSourceConfigItemsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 688003E1-D1B4-5468-957E-2FFB3AC8D79B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsDataSourceConfigItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourceConfigItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourceConfigItemsResponseBody) GetData() []*ListMmsDataSourceConfigItemsResponseBodyData {
	return s.Data
}

func (s *ListMmsDataSourceConfigItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsDataSourceConfigItemsResponseBody) SetData(v []*ListMmsDataSourceConfigItemsResponseBodyData) *ListMmsDataSourceConfigItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBody) SetRequestId(v string) *ListMmsDataSourceConfigItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMmsDataSourceConfigItemsResponseBodyData struct {
	// example:
	//
	// MaxCompute Default Project
	Desc  *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	Enums []*string `json:"enums,omitempty" xml:"enums,omitempty" type:"Repeated"`
	// example:
	//
	// group.basic
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// example:
	//
	// mc.default.project
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// MaxCompute Default Project
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// project
	PlaceHolder *string `json:"placeHolder,omitempty" xml:"placeHolder,omitempty"`
	// example:
	//
	// true
	Required *bool                  `json:"required,omitempty" xml:"required,omitempty"`
	SubItems map[string]interface{} `json:"subItems,omitempty" xml:"subItems,omitempty"`
	// example:
	//
	// .keytab
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsDataSourceConfigItemsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourceConfigItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetDesc() *string {
	return s.Desc
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetEnums() []*string {
	return s.Enums
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetGroup() *string {
	return s.Group
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetPlaceHolder() *string {
	return s.PlaceHolder
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetRequired() *bool {
	return s.Required
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetSubItems() map[string]interface{} {
	return s.SubItems
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetSubType() *string {
	return s.SubType
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) GetValue() interface{} {
	return s.Value
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetDesc(v string) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.Desc = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetEnums(v []*string) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.Enums = v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetGroup(v string) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.Group = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetKey(v string) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.Key = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetName(v string) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetPlaceHolder(v string) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.PlaceHolder = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetRequired(v bool) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.Required = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetSubItems(v map[string]interface{}) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.SubItems = v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetSubType(v string) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.SubType = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetType(v string) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) SetValue(v interface{}) *ListMmsDataSourceConfigItemsResponseBodyData {
	s.Value = v
	return s
}

func (s *ListMmsDataSourceConfigItemsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
