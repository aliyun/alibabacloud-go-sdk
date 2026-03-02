// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStrategyItem interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *StrategyItem
	GetAction() *string
	SetResource(v string) *StrategyItem
	GetResource() *string
}

type StrategyItem struct {
	Action   *string `json:"action,omitempty" xml:"action,omitempty"`
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s StrategyItem) String() string {
	return dara.Prettify(s)
}

func (s StrategyItem) GoString() string {
	return s.String()
}

func (s *StrategyItem) GetAction() *string {
	return s.Action
}

func (s *StrategyItem) GetResource() *string {
	return s.Resource
}

func (s *StrategyItem) SetAction(v string) *StrategyItem {
	s.Action = &v
	return s
}

func (s *StrategyItem) SetResource(v string) *StrategyItem {
	s.Resource = &v
	return s
}

func (s *StrategyItem) Validate() error {
	return dara.Validate(s)
}
