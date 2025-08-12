// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertStrategyRelation interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *AlertStrategyRelation
	GetAlertName() *string
	SetStrategyUuid(v string) *AlertStrategyRelation
	GetStrategyUuid() *string
}

type AlertStrategyRelation struct {
	AlertName    *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	StrategyUuid *string `json:"StrategyUuid,omitempty" xml:"StrategyUuid,omitempty"`
}

func (s AlertStrategyRelation) String() string {
	return dara.Prettify(s)
}

func (s AlertStrategyRelation) GoString() string {
	return s.String()
}

func (s *AlertStrategyRelation) GetAlertName() *string {
	return s.AlertName
}

func (s *AlertStrategyRelation) GetStrategyUuid() *string {
	return s.StrategyUuid
}

func (s *AlertStrategyRelation) SetAlertName(v string) *AlertStrategyRelation {
	s.AlertName = &v
	return s
}

func (s *AlertStrategyRelation) SetStrategyUuid(v string) *AlertStrategyRelation {
	s.StrategyUuid = &v
	return s
}

func (s *AlertStrategyRelation) Validate() error {
	return dara.Validate(s)
}
