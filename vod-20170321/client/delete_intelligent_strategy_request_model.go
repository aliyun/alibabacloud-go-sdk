// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntelligentStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v string) *DeleteIntelligentStrategyRequest
	GetStrategyId() *string
}

type DeleteIntelligentStrategyRequest struct {
	// This parameter is required.
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DeleteIntelligentStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntelligentStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntelligentStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DeleteIntelligentStrategyRequest) SetStrategyId(v string) *DeleteIntelligentStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *DeleteIntelligentStrategyRequest) Validate() error {
	return dara.Validate(s)
}
