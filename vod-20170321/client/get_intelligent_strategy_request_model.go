// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntelligentStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v string) *GetIntelligentStrategyRequest
	GetStrategyId() *string
}

type GetIntelligentStrategyRequest struct {
	// This parameter is required.
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s GetIntelligentStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIntelligentStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetIntelligentStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetIntelligentStrategyRequest) SetStrategyId(v string) *GetIntelligentStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *GetIntelligentStrategyRequest) Validate() error {
	return dara.Validate(s)
}
