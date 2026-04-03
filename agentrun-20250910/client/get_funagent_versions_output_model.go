// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunagentVersionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*FunagentVersionItem) *GetFunagentVersionsOutput
	GetItems() []*FunagentVersionItem
}

type GetFunagentVersionsOutput struct {
	Items []*FunagentVersionItem `json:"items" xml:"items" type:"Repeated"`
}

func (s GetFunagentVersionsOutput) String() string {
	return dara.Prettify(s)
}

func (s GetFunagentVersionsOutput) GoString() string {
	return s.String()
}

func (s *GetFunagentVersionsOutput) GetItems() []*FunagentVersionItem {
	return s.Items
}

func (s *GetFunagentVersionsOutput) SetItems(v []*FunagentVersionItem) *GetFunagentVersionsOutput {
	s.Items = v
	return s
}

func (s *GetFunagentVersionsOutput) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
