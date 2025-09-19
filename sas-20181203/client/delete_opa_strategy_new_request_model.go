// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpaStrategyNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyIds(v []*int64) *DeleteOpaStrategyNewRequest
	GetStrategyIds() []*int64
}

type DeleteOpaStrategyNewRequest struct {
	// The IDs of rules.
	StrategyIds []*int64 `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty" type:"Repeated"`
}

func (s DeleteOpaStrategyNewRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpaStrategyNewRequest) GoString() string {
	return s.String()
}

func (s *DeleteOpaStrategyNewRequest) GetStrategyIds() []*int64 {
	return s.StrategyIds
}

func (s *DeleteOpaStrategyNewRequest) SetStrategyIds(v []*int64) *DeleteOpaStrategyNewRequest {
	s.StrategyIds = v
	return s
}

func (s *DeleteOpaStrategyNewRequest) Validate() error {
	return dara.Validate(s)
}
