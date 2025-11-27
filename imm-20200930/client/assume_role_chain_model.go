// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleChain interface {
	dara.Model
	String() string
	GoString() string
	SetChain(v []*AssumeRoleChainNode) *AssumeRoleChain
	GetChain() []*AssumeRoleChainNode
	SetPolicy(v string) *AssumeRoleChain
	GetPolicy() *string
}

type AssumeRoleChain struct {
	Chain  []*AssumeRoleChainNode `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Repeated"`
	Policy *string                `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s AssumeRoleChain) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleChain) GoString() string {
	return s.String()
}

func (s *AssumeRoleChain) GetChain() []*AssumeRoleChainNode {
	return s.Chain
}

func (s *AssumeRoleChain) GetPolicy() *string {
	return s.Policy
}

func (s *AssumeRoleChain) SetChain(v []*AssumeRoleChainNode) *AssumeRoleChain {
	s.Chain = v
	return s
}

func (s *AssumeRoleChain) SetPolicy(v string) *AssumeRoleChain {
	s.Policy = &v
	return s
}

func (s *AssumeRoleChain) Validate() error {
	if s.Chain != nil {
		for _, item := range s.Chain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
