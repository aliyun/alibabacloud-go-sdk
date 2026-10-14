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
	// The database engine type. This value must match the EngineType of the target instance, such as MySQL or PostgreSQL.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The database instance identity. You can obtain this value from Items[].InstanceId in the response of ListDataMaskingInstances.
	//
	// example:
	//
	// rm-2ze1abcdefgh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the response description. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The cloud product code. Column encryption supports RDS, POLARDB, and PolarDBX2. If ProductId is also specified, the two must correspond to each other.
	//
	// example:
	//
	// RDS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The cloud product type identifier. Valid values:
	//
	// - 5: RDS.
	//
	// - 8: POLARDB.
	//
	// - 27: PolarDBX2.
	//
	// If both ProductId and ProductCode are specified, they must correspond to each other.
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The list of column encryption targets grouped by database and table.
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
