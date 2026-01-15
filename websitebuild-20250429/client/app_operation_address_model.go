// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppOperationAddress interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*AppOperateAction) *AppOperationAddress
	GetActions() []*AppOperateAction
}

type AppOperationAddress struct {
	Actions []*AppOperateAction `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
}

func (s AppOperationAddress) String() string {
	return dara.Prettify(s)
}

func (s AppOperationAddress) GoString() string {
	return s.String()
}

func (s *AppOperationAddress) GetActions() []*AppOperateAction {
	return s.Actions
}

func (s *AppOperationAddress) SetActions(v []*AppOperateAction) *AppOperationAddress {
	s.Actions = v
	return s
}

func (s *AppOperationAddress) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
