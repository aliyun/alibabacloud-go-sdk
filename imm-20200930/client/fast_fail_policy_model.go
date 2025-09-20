// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFastFailPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *FastFailPolicy
	GetAction() *string
}

type FastFailPolicy struct {
	// example:
	//
	// abort
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
}

func (s FastFailPolicy) String() string {
	return dara.Prettify(s)
}

func (s FastFailPolicy) GoString() string {
	return s.String()
}

func (s *FastFailPolicy) GetAction() *string {
	return s.Action
}

func (s *FastFailPolicy) SetAction(v string) *FastFailPolicy {
	s.Action = &v
	return s
}

func (s *FastFailPolicy) Validate() error {
	return dara.Validate(s)
}
