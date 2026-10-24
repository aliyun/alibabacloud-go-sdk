// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientRuleStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdListShrink(v string) *DeleteClientRuleStrategyShrinkRequest
	GetIdListShrink() *string
}

type DeleteClientRuleStrategyShrinkRequest struct {
	// The list of IDs of the policies to delete. Use the policy IDs returned by the create operation or obtained from the policy list or details operations.
	IdListShrink *string `json:"IdList,omitempty" xml:"IdList,omitempty"`
}

func (s DeleteClientRuleStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientRuleStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientRuleStrategyShrinkRequest) GetIdListShrink() *string {
	return s.IdListShrink
}

func (s *DeleteClientRuleStrategyShrinkRequest) SetIdListShrink(v string) *DeleteClientRuleStrategyShrinkRequest {
	s.IdListShrink = &v
	return s
}

func (s *DeleteClientRuleStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
