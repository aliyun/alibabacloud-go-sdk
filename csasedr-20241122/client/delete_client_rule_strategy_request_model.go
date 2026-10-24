// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientRuleStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdList(v []*int64) *DeleteClientRuleStrategyRequest
	GetIdList() []*int64
}

type DeleteClientRuleStrategyRequest struct {
	// The list of IDs of the policies to delete. Use the policy IDs returned by the create operation or obtained from the policy list or details operations.
	IdList []*int64 `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
}

func (s DeleteClientRuleStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientRuleStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientRuleStrategyRequest) GetIdList() []*int64 {
	return s.IdList
}

func (s *DeleteClientRuleStrategyRequest) SetIdList(v []*int64) *DeleteClientRuleStrategyRequest {
	s.IdList = v
	return s
}

func (s *DeleteClientRuleStrategyRequest) Validate() error {
	return dara.Validate(s)
}
