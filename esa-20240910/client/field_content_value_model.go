// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFieldContentValue interface {
	dara.Model
	String() string
	GoString() string
	SetSortOrder(v int64) *FieldContentValue
	GetSortOrder() *int64
	SetFieldList(v []*FieldContentValueFieldList) *FieldContentValue
	GetFieldList() []*FieldContentValueFieldList
}

type FieldContentValue struct {
	SortOrder *int64                        `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	FieldList []*FieldContentValueFieldList `json:"FieldList,omitempty" xml:"FieldList,omitempty" type:"Repeated"`
}

func (s FieldContentValue) String() string {
	return dara.Prettify(s)
}

func (s FieldContentValue) GoString() string {
	return s.String()
}

func (s *FieldContentValue) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *FieldContentValue) GetFieldList() []*FieldContentValueFieldList {
	return s.FieldList
}

func (s *FieldContentValue) SetSortOrder(v int64) *FieldContentValue {
	s.SortOrder = &v
	return s
}

func (s *FieldContentValue) SetFieldList(v []*FieldContentValueFieldList) *FieldContentValue {
	s.FieldList = v
	return s
}

func (s *FieldContentValue) Validate() error {
	if s.FieldList != nil {
		for _, item := range s.FieldList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FieldContentValueFieldList struct {
	FieldName     *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DescriptionCn *string `json:"DescriptionCn,omitempty" xml:"DescriptionCn,omitempty"`
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DataType      *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	SortOrder     *int64  `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	IsDefault     *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
}

func (s FieldContentValueFieldList) String() string {
	return dara.Prettify(s)
}

func (s FieldContentValueFieldList) GoString() string {
	return s.String()
}

func (s *FieldContentValueFieldList) GetFieldName() *string {
	return s.FieldName
}

func (s *FieldContentValueFieldList) GetDescription() *string {
	return s.Description
}

func (s *FieldContentValueFieldList) GetDescriptionCn() *string {
	return s.DescriptionCn
}

func (s *FieldContentValueFieldList) GetCategory() *string {
	return s.Category
}

func (s *FieldContentValueFieldList) GetDataType() *string {
	return s.DataType
}

func (s *FieldContentValueFieldList) GetSortOrder() *int64 {
	return s.SortOrder
}

func (s *FieldContentValueFieldList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *FieldContentValueFieldList) SetFieldName(v string) *FieldContentValueFieldList {
	s.FieldName = &v
	return s
}

func (s *FieldContentValueFieldList) SetDescription(v string) *FieldContentValueFieldList {
	s.Description = &v
	return s
}

func (s *FieldContentValueFieldList) SetDescriptionCn(v string) *FieldContentValueFieldList {
	s.DescriptionCn = &v
	return s
}

func (s *FieldContentValueFieldList) SetCategory(v string) *FieldContentValueFieldList {
	s.Category = &v
	return s
}

func (s *FieldContentValueFieldList) SetDataType(v string) *FieldContentValueFieldList {
	s.DataType = &v
	return s
}

func (s *FieldContentValueFieldList) SetSortOrder(v int64) *FieldContentValueFieldList {
	s.SortOrder = &v
	return s
}

func (s *FieldContentValueFieldList) SetIsDefault(v bool) *FieldContentValueFieldList {
	s.IsDefault = &v
	return s
}

func (s *FieldContentValueFieldList) Validate() error {
	return dara.Validate(s)
}
