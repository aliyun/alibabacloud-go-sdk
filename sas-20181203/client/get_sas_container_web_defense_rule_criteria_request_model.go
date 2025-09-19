// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *GetSasContainerWebDefenseRuleCriteriaRequest
	GetValue() *string
}

type GetSasContainerWebDefenseRuleCriteriaRequest struct {
	// The value of the search condition. Fuzzy match is supported.
	//
	// example:
	//
	// 525
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSasContainerWebDefenseRuleCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleCriteriaRequest) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *GetSasContainerWebDefenseRuleCriteriaRequest) SetValue(v string) *GetSasContainerWebDefenseRuleCriteriaRequest {
	s.Value = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleCriteriaRequest) Validate() error {
	return dara.Validate(s)
}
