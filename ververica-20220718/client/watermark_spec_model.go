// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWatermarkSpec interface {
	dara.Model
	String() string
	GoString() string
	SetColumn(v string) *WatermarkSpec
	GetColumn() *string
	SetWatermarkExpression(v string) *WatermarkSpec
	GetWatermarkExpression() *string
	SetWatermarkType(v string) *WatermarkSpec
	GetWatermarkType() *string
}

type WatermarkSpec struct {
	Column              *string `json:"column,omitempty" xml:"column,omitempty"`
	WatermarkExpression *string `json:"watermarkExpression,omitempty" xml:"watermarkExpression,omitempty"`
	WatermarkType       *string `json:"watermarkType,omitempty" xml:"watermarkType,omitempty"`
}

func (s WatermarkSpec) String() string {
	return dara.Prettify(s)
}

func (s WatermarkSpec) GoString() string {
	return s.String()
}

func (s *WatermarkSpec) GetColumn() *string {
	return s.Column
}

func (s *WatermarkSpec) GetWatermarkExpression() *string {
	return s.WatermarkExpression
}

func (s *WatermarkSpec) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *WatermarkSpec) SetColumn(v string) *WatermarkSpec {
	s.Column = &v
	return s
}

func (s *WatermarkSpec) SetWatermarkExpression(v string) *WatermarkSpec {
	s.WatermarkExpression = &v
	return s
}

func (s *WatermarkSpec) SetWatermarkType(v string) *WatermarkSpec {
	s.WatermarkType = &v
	return s
}

func (s *WatermarkSpec) Validate() error {
	return dara.Validate(s)
}
