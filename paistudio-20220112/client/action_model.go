// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAction interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *Action
	GetActionType() *string
}

type Action struct {
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
}

func (s Action) String() string {
	return dara.Prettify(s)
}

func (s Action) GoString() string {
	return s.String()
}

func (s *Action) GetActionType() *string {
	return s.ActionType
}

func (s *Action) SetActionType(v string) *Action {
	s.ActionType = &v
	return s
}

func (s *Action) Validate() error {
	return dara.Validate(s)
}
