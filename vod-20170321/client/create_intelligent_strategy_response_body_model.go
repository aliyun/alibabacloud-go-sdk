// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntelligentStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateIntelligentStrategyResponseBody
	GetRequestId() *string
	SetStrategyId(v string) *CreateIntelligentStrategyResponseBody
	GetStrategyId() *string
}

type CreateIntelligentStrategyResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s CreateIntelligentStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIntelligentStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntelligentStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIntelligentStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *CreateIntelligentStrategyResponseBody) SetRequestId(v string) *CreateIntelligentStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntelligentStrategyResponseBody) SetStrategyId(v string) *CreateIntelligentStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *CreateIntelligentStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
