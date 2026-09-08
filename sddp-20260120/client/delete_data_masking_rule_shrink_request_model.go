// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataMaskingRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngineType(v string) *DeleteDataMaskingRuleShrinkRequest
	GetEngineType() *string
	SetInstanceId(v string) *DeleteDataMaskingRuleShrinkRequest
	GetInstanceId() *string
	SetLang(v string) *DeleteDataMaskingRuleShrinkRequest
	GetLang() *string
	SetProductCode(v string) *DeleteDataMaskingRuleShrinkRequest
	GetProductCode() *string
	SetProductId(v int64) *DeleteDataMaskingRuleShrinkRequest
	GetProductId() *int64
	SetSubRuleListShrink(v string) *DeleteDataMaskingRuleShrinkRequest
	GetSubRuleListShrink() *string
}

type DeleteDataMaskingRuleShrinkRequest struct {
	EngineType        *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductId         *int64  `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	SubRuleListShrink *string `json:"SubRuleList,omitempty" xml:"SubRuleList,omitempty"`
}

func (s DeleteDataMaskingRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataMaskingRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataMaskingRuleShrinkRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DeleteDataMaskingRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDataMaskingRuleShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDataMaskingRuleShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DeleteDataMaskingRuleShrinkRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DeleteDataMaskingRuleShrinkRequest) GetSubRuleListShrink() *string {
	return s.SubRuleListShrink
}

func (s *DeleteDataMaskingRuleShrinkRequest) SetEngineType(v string) *DeleteDataMaskingRuleShrinkRequest {
	s.EngineType = &v
	return s
}

func (s *DeleteDataMaskingRuleShrinkRequest) SetInstanceId(v string) *DeleteDataMaskingRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDataMaskingRuleShrinkRequest) SetLang(v string) *DeleteDataMaskingRuleShrinkRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataMaskingRuleShrinkRequest) SetProductCode(v string) *DeleteDataMaskingRuleShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *DeleteDataMaskingRuleShrinkRequest) SetProductId(v int64) *DeleteDataMaskingRuleShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *DeleteDataMaskingRuleShrinkRequest) SetSubRuleListShrink(v string) *DeleteDataMaskingRuleShrinkRequest {
	s.SubRuleListShrink = &v
	return s
}

func (s *DeleteDataMaskingRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
