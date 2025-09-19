// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStrategyTemplateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetStrategyId(v int64) *GetStrategyTemplateDetailRequest
	GetStrategyId() *int64
}

type GetStrategyTemplateDetailRequest struct {
	// The ID of the template.
	//
	// >  You can call the [GetOpaStrategyTemplateSummary](~~GetOpaStrategyTemplateSummary~~) operation to query the IDs of templates.
	//
	// example:
	//
	// 2
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s GetStrategyTemplateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStrategyTemplateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetStrategyTemplateDetailRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *GetStrategyTemplateDetailRequest) SetStrategyId(v int64) *GetStrategyTemplateDetailRequest {
	s.StrategyId = &v
	return s
}

func (s *GetStrategyTemplateDetailRequest) Validate() error {
	return dara.Validate(s)
}
