// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFullSampleItem interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *FullSampleItem
	GetId() *string
	SetModifyTime(v string) *FullSampleItem
	GetModifyTime() *string
	SetOperator(v string) *FullSampleItem
	GetOperator() *string
}

type FullSampleItem struct {
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Operator   *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s FullSampleItem) String() string {
	return dara.Prettify(s)
}

func (s FullSampleItem) GoString() string {
	return s.String()
}

func (s *FullSampleItem) GetId() *string {
	return s.Id
}

func (s *FullSampleItem) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *FullSampleItem) GetOperator() *string {
	return s.Operator
}

func (s *FullSampleItem) SetId(v string) *FullSampleItem {
	s.Id = &v
	return s
}

func (s *FullSampleItem) SetModifyTime(v string) *FullSampleItem {
	s.ModifyTime = &v
	return s
}

func (s *FullSampleItem) SetOperator(v string) *FullSampleItem {
	s.Operator = &v
	return s
}

func (s *FullSampleItem) Validate() error {
	return dara.Validate(s)
}
