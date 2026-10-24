// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientRuleStrategyDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetClientRuleStrategyDetailRequest
	GetId() *int64
}

type GetClientRuleStrategyDetailRequest struct {
	// The ID of the client rule policy. You can obtain the ID from the response of the create operation or the policy list.
	//
	// example:
	//
	// 14808
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetClientRuleStrategyDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientRuleStrategyDetailRequest) GoString() string {
	return s.String()
}

func (s *GetClientRuleStrategyDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *GetClientRuleStrategyDetailRequest) SetId(v int64) *GetClientRuleStrategyDetailRequest {
	s.Id = &v
	return s
}

func (s *GetClientRuleStrategyDetailRequest) Validate() error {
	return dara.Validate(s)
}
